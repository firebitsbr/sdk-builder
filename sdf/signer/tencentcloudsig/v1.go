package tencentcloudsig

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/credentials"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/sdfutil"
)

// Signer 腾讯云API签名
type Signer struct {
	// The authentication credentials the request will be signed against.
	// This value must be set to sign requests.
	Credentials *credentials.Credentials

	// Sets the log level the signer should use when reporting information to
	// the logger. If the logger is nil nothing will be logged. See
	// sdf.LogLevelType for more information on available logging levels
	//
	// By default nothing will be logged.
	Debug sdf.LogLevelType

	// The logger loging information will be written to. If there the logger
	// is nil, nothing will be logged.
	Logger sdf.Logger

	// currentTimeFn returns the time value which represents the current time.
	// This value should only be used for testing. If it is nil the default
	// time.Now will be used.
	currentTimeFn func() time.Time
}

// NewSigner returns a Signer
func NewSigner(credentials *credentials.Credentials, options ...func(*Signer)) *Signer {
	v1 := &Signer{
		Credentials: credentials,
	}

	for _, option := range options {
		option(v1)
	}

	return v1
}

type signingCtx struct {
	ServiceName string
	Region      string
	Request     *http.Request
	Query       url.Values
	Time        time.Time
	ExpireTime  time.Duration

	credValues credentials.Value

	signature string
}

func (ctx *signingCtx) buildStringToSign() string {

	stringToSign := strings.ToUpper(ctx.Request.Method)
	stringToSign += ctx.Request.URL.Host
	stringToSign += ctx.Request.URL.Path
	stringToSign += "?"

	params := ctx.Query

	// 排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var plainParms string
	for i := range keys {
		k := keys[i]
		plainParms += "&" + fmt.Sprintf("%v", k) + "=" + fmt.Sprintf("%v", params[k][0])
	}

	stringToSign += plainParms[1:]

	return stringToSign
}

func (ctx *signingCtx) completeSignParams(apiVersion, action, region string) {
	params := ctx.Query
	params.Set("SecretId", ctx.credValues.AccessKeyID)
	params.Set("Timestamp", fmt.Sprintf("%v", time.Now().Unix()))
	params.Set("Nonce", sdfutil.GenNonce())
	params.Set("Version", apiVersion)
	params.Set("Action", action)
	params.Set("Region", region)
}

func (ctx *signingCtx) buildSignature() string {
	secretKey := ctx.credValues.SecretAccessKey

	stringToSign := ctx.buildStringToSign()

	return ShaHmac1(stringToSign, secretKey)
}

func ShaHmac1(source, secret string) string {
	key := []byte(secret)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(source))
	signedBytes := hmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

// SignRequestHandler is a named request handler the SDK will use to sign
// service client request with using the V1 signature.
var SignRequestHandler = request.NamedHandler{
	Name: "tencentv1.SignRequestHandler", Fn: SignSDKRequest,
}

// SignSDKRequest signs an tencent request with the V1 signature. This
// request handler should only be used with the SDK's built in service client's
// API operation requests.
func SignSDKRequest(req *request.Request) {
	signSDKRequestWithCurrTime(req, time.Now)
}

func signSDKRequestWithCurrTime(req *request.Request, curTimeFn func() time.Time) {
	// If the request does not need to be signed ignore the signing of the
	// request if the AnonymousCredentials object is used.
	if req.Config.Credentials == credentials.AnonymousCredentials {
		return
	}

	region := req.ClientInfo.SigningRegion
	if region == "" {
		region = sdf.StringValue(req.Config.Region)
	}

	name := req.ClientInfo.SigningName
	if name == "" {
		name = req.ClientInfo.ServiceName
	}

	apiVersion := req.ClientInfo.APIVersion
	action := req.Operation.Name

	v1 := NewSigner(req.Config.Credentials, func(v1 *Signer) {
		v1.Debug = req.Config.LogLevel.Value()
		v1.Logger = req.Config.Logger
		v1.currentTimeFn = curTimeFn
		v1.Credentials = req.Config.Credentials
	})

	signingTime := req.Time
	if !req.LastSignedAt.IsZero() {
		signingTime = req.LastSignedAt
	}

	err := v1.sign(req.HTTPRequest, name, region, apiVersion, action,
		req.ExpireTime, signingTime,
	)
	if err != nil {
		req.Error = err
		return
	}

	req.LastSignedAt = curTimeFn()
}

func (v1 *Signer) sign(r *http.Request, service, region, apiVersion, action string, exp time.Duration, signTime time.Time) error {
	currentTimeFn := v1.currentTimeFn
	if currentTimeFn == nil {
		currentTimeFn = time.Now
	}

	ctx := &signingCtx{
		Request:     r,
		Query:       r.URL.Query(),
		Time:        signTime,
		ExpireTime:  exp,
		ServiceName: service,
		Region:      region,
	}

	var err error
	ctx.credValues, err = v1.Credentials.Get()
	if err != nil {
		return err
	}

	if ctx.isRequestSigned() {
		ctx.Time = currentTimeFn()
		ctx.handlePresignRemoval()
	}

	sig := ctx.build(region, apiVersion, action)
	ctx.Query.Set("Signature", sig)
	ctx.Request.URL.RawQuery = ctx.Query.Encode()
	if v1.Debug.Matches(sdf.LogDebugWithSigning) {
		v1.logSigningInfo(ctx)
	}

	return nil
}

const logSignInfoMsg = `DEBUG: Request Signature:
---[ HOST STRING  ]-----------------------------
%s
---[ SIGN TO SIGN ]--------------------------------
%s
%s
-----------------------------------------------------`
const logSignedURLMsg = `
---[ SIGNED URL ]------------------------------------
%s`

func (v1 *Signer) logSigningInfo(ctx *signingCtx) {
	msg := fmt.Sprintf(logSignInfoMsg, ctx.Request.Host, ctx.Request.URL.Path, ctx.Request.URL.RawQuery)
	v1.Logger.Log(msg)
}

func (ctx *signingCtx) handlePresignRemoval() {
	// The credentials have expired for this request. The current signing
	// is invalid, and needs to be request because the request will fail.
	ctx.removePresign()
}

// isRequestSigned returns if the request is currently signed or presigned
func (ctx *signingCtx) isRequestSigned() bool {
	if ctx.Query.Get("Signature") != "" {
		return true
	}

	return false
}

// unsign removes signing flags for both signed and presigned requests.
func (ctx *signingCtx) removePresign() {
	ctx.Query.Del("Signature")
}

func (ctx *signingCtx) build(region, apiVersion, action string) string {
	ctx.completeSignParams(apiVersion, action, region)
	return ctx.buildSignature()
}

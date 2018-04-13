package alicloudsig

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/credentials"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/sdfutil"
)

// Signer 阿里云API签名
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

func (v1 *Signer) logSigningInfo(ctx *signingCtx) {
	msg := fmt.Sprintf(logSignInfoMsg, ctx.Request.Host, ctx.Request.URL.Path, ctx.Request.URL.RawQuery)
	v1.Logger.Log(msg)
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

func replace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func (s *signingCtx) SignatureVersion() string {
	return "1.0"
}

func (s *signingCtx) SignatureMethod() string {
	return "HMAC-SHA1"
}

func (s *signingCtx) SignatureFormat() string {
	return "JSON"
}

func (s *signingCtx) completeSignParams(apiVersion, action, region string) {
	params := s.Query
	params.Set("AccessKeyId", s.credValues.AccessKeyID)
	params.Set("SignatureMethod", s.SignatureMethod())
	params.Set("SignatureVersion", s.SignatureVersion())
	params.Set("Timestamp", sdfutil.GetTimeInFormatISO8601())
	params.Set("SignatureNonce", sdfutil.GetUUIDV4())
	params.Set("Format", s.SignatureFormat())
	params.Set("Version", apiVersion)
	params.Set("Action", action)
	params.Set("RegionId", region)
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

func (s *signingCtx) handlePresignRemoval() {
	// The credentials have expired for this request. The current signing
	// is invalid, and needs to be request because the request will fail.
	s.removePresign()
}

// isRequestSigned returns if the request is currently signed or presigned
func (s *signingCtx) isRequestSigned() bool {
	if s.Query.Get("Signature") != "" {
		return true
	}

	return false
}

// unsign removes signing flags for both signed and presigned requests.
func (s *signingCtx) removePresign() {
	s.Query.Del("Signature")
}

// SignRequestHandler is a named request handler the SDK will use to sign
// service client request with using the V1 signature.
var SignRequestHandler = request.NamedHandler{
	Name: "aliv1.SignRequestHandler", Fn: SignSDKRequest,
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

	err := v1.signRequest(req.HTTPRequest, name, region, apiVersion, action,
		req.ExpireTime, signingTime,
	)
	if err != nil {
		req.Error = err
		return
	}

	req.LastSignedAt = curTimeFn()
}

func (v1 *Signer) signRequest(r *http.Request, service, region, apiVersion, action string, exp time.Duration, signTime time.Time) error {
	currentTimeFn := v1.currentTimeFn
	if currentTimeFn == nil {
		currentTimeFn = time.Now
	}

	s := &signingCtx{
		Request:     r,
		Query:       r.URL.Query(),
		Time:        signTime,
		ExpireTime:  exp,
		ServiceName: service,
		Region:      region,
	}

	var err error
	s.credValues, err = v1.Credentials.Get()
	if err != nil {
		return err
	}

	if s.isRequestSigned() {
		s.Time = currentTimeFn()
		s.handlePresignRemoval()
	}

	sig := s.build(region, apiVersion, action)
	s.Query.Set("Signature", sig)
	s.Request.URL.RawQuery = s.Query.Encode()

	if v1.Debug.Matches(sdf.LogDebugWithSigning) {
		v1.logSigningInfo(s)
	}

	return nil
}

func (s *signingCtx) buildCredentialString(canonicalizedQueryString string) string {
	return strings.ToUpper(s.Request.Method) + "&%2F&" + canonicalizedQueryString
}

func (s *signingCtx) build(region, apiVersion, action string) string {
	s.completeSignParams(apiVersion, action, region)
	stringToSign := s.buildStringToSign()
	return s.sign(stringToSign, "&")
}

func (s *signingCtx) buildStringToSign() string {
	stringToSign := s.Query.Encode()
	stringToSign = replace(stringToSign)
	stringToSign = url.QueryEscape(stringToSign)
	return s.buildCredentialString(stringToSign)
}

func (s *signingCtx) sign(stringToSign, secretSuffix string) string {
	secret := s.credValues.SecretAccessKey + secretSuffix
	return ShaHmac1(stringToSign, secret)
}

func ShaHmac1(source, secret string) string {
	key := []byte(secret)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(source))
	signedBytes := hmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

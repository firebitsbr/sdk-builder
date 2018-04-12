package alicloudsig

import (
	"crypto/hmac"
	srand "crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/credentials"
	"github.com/polefishu/sdk-builder/sdf/request"
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

type signingCtx struct {
	ServiceName      string
	Region           string
	Request          *http.Request
	Body             io.ReadSeeker
	Query            url.Values
	Time             time.Time
	ExpireTime       time.Duration
	SignedHeaderVals http.Header

	credValues credentials.Value

	signedHeaders string
	signature     string
}

const dictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

//createRandomString create random string
func createRandomString() string {
	b := make([]byte, 32)
	l := len(dictionary)

	_, err := srand.Read(b)

	if err != nil {
		// fail back to insecure rand
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}

func (ctx *signingCtx) makePlainText() (plainText string, err error) {
	params := ctx.Query
	params.Set("AccessKeyId", ctx.credValues.AccessKeyID)
	params.Set("SignatureMethod", "HMAC-SHA1")
	params.Set("SignatureVersion", "1.0")
	// params.Set("Timestamp", "2017-06-29T11:12:42Z")
	params.Set("Timestamp", NewISO6801Time(time.Now().UTC()).String())
	// params.Set("SignatureNonce", "bUhobV5P67XalU6ET1_cOp8GCjCdhF0K")
	params.Set("SignatureNonce", createRandomString())

	// 排序
	// keys := make([]string, 0, len(params))
	// for k := range params {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)

	// var plainParms string
	// for i := range keys {
	// 	k := keys[i]
	// 	plainParms += "&" + fmt.Sprintf("%v", k) + "=" + fmt.Sprintf("%v", params[k][0])
	// }
	// plainText += plainParms[1:]

	// return plainText, nil
	return params.Encode(), nil
}

//createSignature creates signature for string following Aliyun rules
func createSignature(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha1.Write([]byte(stringToSignature))
	sign := hmacSha1.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

func percentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	// str = strings.Replace(str, ":", "%3A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}

func (ctx *signingCtx) buildSignature() {
	secretKey := ctx.credValues.SecretAccessKey + "&"
	var source string

	source, err := ctx.makePlainText()
	if err != nil {
		log.Fatalln("Make PlainText error.", err)
	}

	canonicalizedQueryString := percentReplace(source)
	stringToSign := strings.ToUpper(ctx.Request.Method) + "&%2F&" + url.QueryEscape(canonicalizedQueryString)

	sig := createSignature(stringToSign, secretKey)

	ctx.Query.Set("Signature", sig)
	// ctx.Query.Set("Signature", url.QueryEscape(sig))
	ctx.Request.URL.RawQuery = ctx.Query.Encode()
}

// SignRequestHandler is a named request handler the SDK will use to sign
// service client request with using the V1 signature.
var SignRequestHandler = request.NamedHandler{
	Name: "tv1.SignRequestHandler", Fn: SignSDKRequest,
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

	signedHeaders, err := v1.sign(req.HTTPRequest, req.GetBody(),
		name, region, req.ExpireTime, signingTime,
	)
	if err != nil {
		req.Error = err
		req.SignedHeaderVals = nil
		return
	}

	req.SignedHeaderVals = signedHeaders
	req.LastSignedAt = curTimeFn()
}

func (v1 *Signer) sign(r *http.Request, body io.ReadSeeker, service, region string, exp time.Duration, signTime time.Time) (http.Header, error) {
	currentTimeFn := v1.currentTimeFn
	if currentTimeFn == nil {
		currentTimeFn = time.Now
	}

	ctx := &signingCtx{
		Request:     r,
		Body:        body,
		Query:       r.URL.Query(),
		Time:        signTime,
		ExpireTime:  exp,
		ServiceName: service,
		Region:      region,
	}

	var err error
	ctx.credValues, err = v1.Credentials.Get()
	if err != nil {
		return http.Header{}, err
	}

	if ctx.isRequestSigned() {
		ctx.Time = currentTimeFn()
		ctx.handlePresignRemoval()
	}

	ctx.build()

	if v1.Debug.Matches(sdf.LogDebugWithSigning) {
		v1.logSigningInfo(ctx)
	}

	return ctx.SignedHeaderVals, nil
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

func (ctx *signingCtx) build() {
	ctx.buildSignature() // depends on string to sign
}

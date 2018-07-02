package tencentcloudsig

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/polefishu/sdk-builder/sdf/credentials"
)

func TestSignWithRequestBody(t *testing.T) {
	creds := credentials.NewStaticCredentials("AKIDz8krbsJ5yKBZQpn74WFkmLPx3EXAMPLE", "Gu5t9xGARNpq86cd98joQYCN3EXAMPLE", "")
	signer := NewSigner(creds)

	req, err := http.NewRequest("GET", "https://cvm.tencentcloudapi.com/?InstanceIds.0=ins-09dx96dg&Offset=0&Limit=20", nil)

	name := "vpc"
	region := "ap-guangzhou"
	action := "DescribeInstances"
	apiVersion := "2017-03-12"

	currentTimeFn := signer.currentTimeFn
	if currentTimeFn == nil {
		currentTimeFn = time.Now
	}

	ctx := &signingCtx{
		Request:     req,
		Query:       req.URL.Query(),
		Time:        time.Now(),
		ExpireTime:  time.Duration(200),
		ServiceName: name,
		Region:      region,
	}

	ctx.credValues, err = signer.Credentials.Get()
	if err != nil {
		t.Fatal(err)
	}

	if ctx.isRequestSigned() {
		ctx.Time = currentTimeFn()
		ctx.handlePresignRemoval()
	}

	params := ctx.Query
	params.Set("SecretId", ctx.credValues.AccessKeyID)
	params.Set("Timestamp", fmt.Sprintf("%v", 1465185768))
	params.Set("Nonce", "11886")
	params.Set("Version", apiVersion)
	params.Set("Action", action)
	params.Set("Region", region)

	secretKey := ctx.credValues.SecretAccessKey

	stringToSign := ctx.buildStringToSign()

	t.Log("stringToSign", stringToSign)
	sig := ShaHmac1(stringToSign, secretKey)
	t.Log("signStr", sig)
	ctx.Query.Set("Signature", sig)
	t.Log(ctx.Query.Encode())
}

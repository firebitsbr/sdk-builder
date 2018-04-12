package rest_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/client"
	"github.com/polefishu/sdk-builder/sdf/client/metadata"
	"github.com/polefishu/sdk-builder/sdf/protocol/rest"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/signer/v4"
	"github.com/polefishu/sdk-builder/sdftesting/unit"
)

func TestUnsetHeaders(t *testing.T) {
	cfg := &sdf.Config{Region: sdf.String("us-west-2")}
	c := unit.Session.ClientConfig("testService", cfg)
	svc := client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName:   "testService",
			SigningName:   c.SigningName,
			SigningRegion: c.SigningRegion,
			Endpoint:      c.Endpoint,
			APIVersion:    "",
		},
		c.Handlers,
	)

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(rest.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(rest.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(rest.UnmarshalMetaHandler)
	op := &request.Operation{
		Name:     "test-operation",
		HTTPPath: "/",
	}

	input := &struct {
		Foo sdf.JSONValue `location:"header" locationName:"x-amz-foo" type:"jsonvalue"`
		Bar sdf.JSONValue `location:"header" locationName:"x-amz-bar" type:"jsonvalue"`
	}{}

	output := &struct {
		Foo sdf.JSONValue `location:"header" locationName:"x-amz-foo" type:"jsonvalue"`
		Bar sdf.JSONValue `location:"header" locationName:"x-amz-bar" type:"jsonvalue"`
	}{}

	req := svc.NewRequest(op, input, output)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(nil)), Header: http.Header{}}
	req.HTTPResponse.Header.Set("X-Amz-Foo", "e30=")

	// unmarshal response
	rest.UnmarshalMeta(req)
	rest.Unmarshal(req)
	if req.Error != nil {
		t.Fatal(req.Error)
	}
}

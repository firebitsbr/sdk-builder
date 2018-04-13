package alicloudquery

//go:generate go run -tags codegen ../../../models/protocol_tests/generate.go ../../../models/protocol_tests/input/ec2.json build_test.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"time"

	"github.com/polefishu/sdk-builder/sdf/protocol/json/jsonutil"
	"github.com/polefishu/sdk-builder/sdf/protocol/query/queryutil"
	"github.com/polefishu/sdk-builder/sdf/protocol/rest"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
)

// BuildHandler is a named request handler for building aliyunquery protocol requests
var BuildHandler = request.NamedHandler{Name: "alicloudsdk.aliyunquery.Build", Fn: Build}

// Build builds a request for the aliyun protocol.
func Build(r *request.Request) {
	body := r.HTTPRequest.URL.Query()
	// body.Set("Action", r.Operation.Name)
	// body.Set("Format", "JSON")

	if err := queryutil.Parse(body, r.Params, true); err != nil {
		r.Error = sdferr.New("SerializationError", "failed encoding EC2 Query request", err)
	}

	r.HTTPRequest.Method = "GET"
	r.HTTPRequest.URL.RawQuery = body.Encode()
}

func genNonce() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Int())
}

var emptyJSON = []byte("{}")

// UnmarshalHandler is a named request handler for unmarshaling jsonrpc protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "alicloudsdk.jsonrpc.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling jsonrpc protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "alicloudsdk.jsonrpc.UnmarshalMeta", Fn: UnmarshalMeta}

// UnmarshalErrorHandler is a named request handler for unmarshaling jsonrpc protocol request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "alicloudsdk.jsonrpc.UnmarshalError", Fn: UnmarshalError}

// Unmarshal unmarshals a response for a JSON RPC service.
func Unmarshal(req *request.Request) {
	defer req.HTTPResponse.Body.Close()
	if req.DataFilled() {
		err := jsonutil.UnmarshalJSON(req.Data, req.HTTPResponse.Body)
		if err != nil {
			req.Error = sdferr.New("SerializationError", "failed decoding JSON RPC response", err)
		}
	}
	return
}

// UnmarshalMeta unmarshals headers from a response for a JSON RPC service.
func UnmarshalMeta(req *request.Request) {
	rest.UnmarshalMeta(req)
}

// UnmarshalError unmarshals an error response for a JSON RPC service.
func UnmarshalError(req *request.Request) {
	defer req.HTTPResponse.Body.Close()
	bodyBytes, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil {
		req.Error = sdferr.New("SerializationError", "failed reading JSON RPC error response", err)
		return
	}
	if len(bodyBytes) == 0 {
		req.Error = sdferr.NewRequestFailure(
			sdferr.New("SerializationError", req.HTTPResponse.Status, nil),
			req.HTTPResponse.StatusCode,
			"",
		)
		return
	}
	var jsonErr jsonErrorResponse
	if err := json.Unmarshal(bodyBytes, &jsonErr); err != nil {
		req.Error = sdferr.New("SerializationError", "failed decoding JSON RPC error response", err)
		return
	}

	req.Error = sdferr.NewRequestFailure(
		sdferr.New(jsonErr.Code, jsonErr.Message, nil),
		req.HTTPResponse.StatusCode,
		jsonErr.RequestID,
	)
}

type jsonErrorResponse struct {
	Code      string `json:"Code"`
	Message   string `json:"message"`
	HostID    string `json:"HostId"`
	RequestID string `json:"RequestId"`
}

// BuildVersionHandler is a named request handler  add Version to url query
var BuildVersionHandler = request.NamedHandler{Name: "alicloudsdk.aliyunquery.BuildVersionHandler", Fn: BuildVersion}

// BuildVersion add Version=2014-05-26 to url query
func BuildVersion(req *request.Request) {
	body := url.Values{
		"Version": {"2014-05-26"},
	}
	req.HTTPRequest.URL.RawQuery = body.Encode()
}

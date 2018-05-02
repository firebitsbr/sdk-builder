package tencentquery

//go:generate go run -tags codegen ../../../models/protocol_tests/generate.go ../../../models/protocol_tests/input/ec2.json build_test.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/polefishu/sdk-builder/sdf/protocol/json/jsonutil"
	"github.com/polefishu/sdk-builder/sdf/protocol/query/queryutil"
	"github.com/polefishu/sdk-builder/sdf/protocol/rest"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
)

// BuildHandler is a named request handler for building cvmquery protocol requests
var BuildHandler = request.NamedHandler{Name: "tencentsdk.tencentquery.Build", Fn: Build}

// Build builds a request for the CVM protocol.
func Build(r *request.Request) {
	body := r.HTTPRequest.URL.Query()

	if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = sdferr.New("SerializationError", "failed encoding tencent Query request", err)
	}

	r.HTTPRequest.Method = "GET"
	r.HTTPRequest.URL.RawQuery = body.Encode()
}

var emptyJSON = []byte("{}")

// UnmarshalHandler is a named request handler for unmarshaling jsonrpc protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "tencentsdk.jsonrpc.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling jsonrpc protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "tencentsdk.jsonrpc.UnmarshalMeta", Fn: UnmarshalMeta}

// UnmarshalErrorHandler is a named request handler for unmarshaling jsonrpc protocol request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "tencentsdk.jsonrpc.UnmarshalError", Fn: UnmarshalError}

// Unmarshal unmarshals a response for a JSON RPC service.
func Unmarshal(req *request.Request) {
	defer req.HTTPResponse.Body.Close()
	if req.DataFilled() {
		bodyBytes, err := ioutil.ReadAll(req.HTTPResponse.Body)

		if err != nil {
			req.Error = sdferr.New("SerializationError", "failed reading JSON RPC response", err)
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
		codes := jsonErr.Response.Error.Code
		if len(codes) > 0 {
			req.Error = sdferr.NewRequestFailure(
				sdferr.New(codes, jsonErr.Response.Error.Message, nil),
				req.HTTPResponse.StatusCode,
				jsonErr.Response.RequestID,
			)
			return
		}
		var jsonOldErr jsonErrorOldResponse
		if err := json.Unmarshal(bodyBytes, &jsonOldErr); err != nil {
			req.Error = sdferr.New("SerializationError", "failed decoding JSON RPC error response", err)
			return
		}
		if jsonOldErr.Code > 0 {
			req.Error = sdferr.NewRequestFailure(
				sdferr.New(fmt.Sprintf("%d", jsonOldErr.Code), jsonOldErr.Message, nil),
				req.HTTPResponse.StatusCode,
				jsonOldErr.CodeDesc,
			)
			return
		}

		err = jsonutil.UnmarshalJSON2(req.Data, bodyBytes)

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

	codes := jsonErr.Response.Error.Code
	req.Error = sdferr.NewRequestFailure(
		sdferr.New(codes, jsonErr.Response.Error.Message, nil),
		req.HTTPResponse.StatusCode,
		jsonErr.Response.RequestID,
	)
}

type errRes struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}
type reponse struct {
	Error     errRes
	RequestID string `json:"RequestId"`
}
type jsonErrorResponse struct {
	Response reponse
}

type jsonErrorOldResponse struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

// BuildVersionHandler is a named request handler  add Version to url query
var BuildVersionHandler = request.NamedHandler{Name: "tencentsdk.tencentquery.BuildVersionHandler", Fn: BuildVersion}

// BuildVersion add Version=2017-03-12 to url query
func BuildVersion(req *request.Request) {
	body := url.Values{
		//"Version": {"2017-03-12"},
		"Region": {req.ClientInfo.SigningRegion},
	}
	req.HTTPRequest.URL.RawQuery = body.Encode()
}

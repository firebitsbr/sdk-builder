// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dfw

import (
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/client"
	"github.com/polefishu/sdk-builder/sdf/client/metadata"
	"github.com/polefishu/sdk-builder/sdf/protocol/tencentquery"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/signer/tencentcloudsig"
)

// DFW provides the API operation methods for making requests to
// SecurityGroup. See this package's package overview docs
// for details on the service.
//
// DFW methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type DFW struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "dfw"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the DFW client with a session.
// If additional configuration is needed for the client instance use the optional
// sdf.Config parameter to add your extra config.
//
// Example:
//     // Create a DFW client from just a session.
//     svc := dfw.New(mySession)
//
//     // Create a DFW client with additional configuration
//     svc := dfw.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*sdf.Config) *DFW {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg sdf.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *DFW {
	svc := &DFW{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2017-06-20",
				JSONVersion:   "1.1",
			},
			handlers,
		),
	}

	// Handlers

	svc.Handlers.Sign.PushBackNamed(tencentcloudsig.SignRequestHandler)

	svc.Handlers.Build.PushBackNamed(tencentquery.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(tencentquery.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(tencentquery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(tencentquery.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DFW operation and runs any
// custom request initialization.
func (c *DFW) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

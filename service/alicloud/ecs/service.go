package ecs

import (
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/client"
	"github.com/polefishu/sdk-builder/sdf/client/metadata"
	"github.com/polefishu/sdk-builder/sdf/protocol/alicloudquery"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/signer/alicloudsig"
)

// ECS provides the API operation methods for making requests to
// Amazon Elastic Compute Cloud. See this package's package overview docs
// for details on the service.
//
// ECS methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ECS struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "ecs"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

func New(p client.ConfigProvider, cfgs ...*sdf.Config) *ECS {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg sdf.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ECS {
	svc := &ECS{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-05-26",
			},
			handlers,
		),
	}

	// Handlers
	/*
		svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
		svc.Handlers.Build.PushBackNamed(ec2query.BuildHandler)
		svc.Handlers.Unmarshal.PushBackNamed(ec2query.UnmarshalHandler)
		svc.Handlers.UnmarshalMeta.PushBackNamed(ec2query.UnmarshalMetaHandler)
		svc.Handlers.UnmarshalError.PushBackNamed(ec2query.UnmarshalErrorHandler)
	*/
	svc.Handlers.Sign.PushBackNamed(alicloudsig.SignRequestHandler)
	svc.Handlers.Build.Clear()
	// svc.Handlers.Build.PushBackNamed(alicloudquery.BuildVersionHandler)

	svc.Handlers.Build.PushBackNamed(alicloudquery.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(alicloudquery.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(alicloudquery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(alicloudquery.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

func (c *ECS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

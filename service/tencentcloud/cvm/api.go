package cvm

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/sdfutil"
)

const opDescribeInstances = "DescribeInstances"

func (c *CVM) DescribeInstancesRequest(input *DescribeInstancesInput) (req *request.Request, output *DescribeInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/index.php",
		Paginator: &request.Paginator{
			InputTokens:     []string{"Offset"},
			OutputTokens:    []string{"Response.TotalCount"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeInstancesInput{}
	}

	output = &DescribeInstancesOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *CVM) DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	return out, req.Send()
}

func (c *CVM) DescribeInstancesWithContext(ctx sdf.Context, input *DescribeInstancesInput, opts ...request.Option) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (c *CVM) DescribeInstancesPages(input *DescribeInstancesInput, fn func(*DescribeInstancesOutput, bool) bool) error {
	return c.DescribeInstancesPagesWithContext(sdf.BackgroundContext(), input, fn)
}

func (c *CVM) DescribeInstancesPagesWithContext(ctx sdf.Context, input *DescribeInstancesInput, fn func(*DescribeInstancesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *DescribeInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for cont && p.Next() {
		cont = fn(p.Page().(*DescribeInstancesOutput), !p.HasNextPage())
	}
	return p.Err()
}

type DescribeInstancesInput struct {
	_ struct{} `type:"structure"`

	InstanceIds []*string `locationName:"InstanceIds" locationNameList:"InstanceIds" type:"list" flattened:"true"`

	// test
	Limit *int64 `type:"integer"`

	// test
	Offset *int64 `type:"integer"`

	Region *string `type:"string" enum:"RegionEnum"`
}

// String returns the string representation
func (s DescribeInstancesInput) String() string {
	return sdfutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesInput) GoString() string {
	return s.String()
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeInstancesInput) SetInstanceIds(v []*string) *DescribeInstancesInput {
	s.InstanceIds = v
	return s
}

// SetLimit sets the Limit field's value.
func (s *DescribeInstancesInput) SetLimit(v int64) *DescribeInstancesInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *DescribeInstancesInput) SetOffset(v int64) *DescribeInstancesInput {
	s.Offset = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DescribeInstancesInput) SetRegion(v string) *DescribeInstancesInput {
	s.Region = &v
	return s
}

// test
type DescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	// test
	Response *Response `type:"structure"`
}

// String returns the string representation
func (s DescribeInstancesOutput) String() string {
	return sdfutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetResponse sets the Response field's value.
func (s *DescribeInstancesOutput) SetResponse(v *Response) *DescribeInstancesOutput {
	s.Response = v
	return s
}

// Response item
type Response struct {
	_ struct{} `type:"structure"`

	// test
	InstanceSet []*Instance `locationName:"InstanceSet" type:"list"`

	RequestID *string `locationName:"RequestId" type:"string"`

	TotalCount *int64 `locationName:"TotalCount" type:"integer"`
}

// String returns the string representation
func (s Response) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Response) GoString() string {
	return s.String()
}

// SetInstanceSet sets the InstanceSet field's value.
func (s *Response) SetInstanceSet(v []*Instance) *Response {
	s.InstanceSet = v
	return s
}

// SetRequestID sets the RequestID field's value.
func (s *Response) SetRequestID(v string) *Response {
	s.RequestID = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *Response) SetTotalCount(v int64) *Response {
	s.TotalCount = &v
	return s
}

type ResponseError struct {
	_ struct{} `type:"structure"`

	Code *string `locationName:"Code" type:"string"`

	Message *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s ResponseError) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponseError) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ResponseError) SetCode(v string) *ResponseError {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ResponseError) SetMessage(v string) *ResponseError {
	s.Message = &v
	return s
}

// item
type Instance struct {
	_ struct{} `type:"structure"`

	// test
	CPU *int64 `locationName:"CPU" type:"integer"`

	CreatedTime *string `locationName:"CreatedTime" type:"string"`

	ExpiredTime *string `locationName:"ExpiredTime" type:"string"`

	ImageID *string `locationName:"ImageId" type:"string"`

	InstanceChargeType *string `locationName:"InstanceChargeType" type:"string"`

	InstanceID *string `locationName:"InstanceId" type:"string"`

	InstanceName *string `locationName:"InstanceName" type:"string"`

	// test
	InstanceType *string `locationName:"InstanceType" type:"string"`

	// test
	Memory *int64 `locationName:"Memory" type:"integer"`

	OsName *string `locationName:"OsName" type:"string"`

	// test
	PrivateIpAddresses []*string `locationName:"PrivateIpAddresses" type:"list"`

	// test
	PublicIpAddresses []*string `locationName:"PublicIpAddresses" type:"list"`
}

// String returns the string representation
func (s Instance) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Instance) GoString() string {
	return s.String()
}

// SetCPU sets the CPU field's value.
func (s *Instance) SetCPU(v int64) *Instance {
	s.CPU = &v
	return s
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *Instance) SetCreatedTime(v string) *Instance {
	s.CreatedTime = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *Instance) SetExpiredTime(v string) *Instance {
	s.ExpiredTime = &v
	return s
}

// SetImageID sets the ImageID field's value.
func (s *Instance) SetImageID(v string) *Instance {
	s.ImageID = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *Instance) SetInstanceChargeType(v string) *Instance {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceID sets the InstanceID field's value.
func (s *Instance) SetInstanceID(v string) *Instance {
	s.InstanceID = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *Instance) SetInstanceName(v string) *Instance {
	s.InstanceName = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *Instance) SetInstanceType(v string) *Instance {
	s.InstanceType = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *Instance) SetMemory(v int64) *Instance {
	s.Memory = &v
	return s
}

// SetOsName sets the OsName field's value.
func (s *Instance) SetOsName(v string) *Instance {
	s.OsName = &v
	return s
}

// SetPrivateIpAddresses sets the PrivateIpAddresses field's value.
func (s *Instance) SetPrivateIpAddresses(v []*string) *Instance {
	s.PrivateIpAddresses = v
	return s
}

// SetPublicIpAddresses sets the PublicIpAddresses field's value.
func (s *Instance) SetPublicIpAddresses(v []*string) *Instance {
	s.PublicIpAddresses = v
	return s
}

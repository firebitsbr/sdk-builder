package ecs

import (
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdf/sdfutil"
)

const opDescribeInstances = "DescribeInstances"

func (c *ECS) DescribeInstancesRequest(input *DescribeInstancesInput) (req *request.Request, output *DescribeInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"PageNumber"},
			OutputTokens:    []string{"PageNumber"},
			LimitToken:      "PageSize",
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

func (c *ECS) DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	return out, req.Send()
}

func (c *ECS) DescribeInstancesWithContext(ctx sdf.Context, input *DescribeInstancesInput, opts ...request.Option) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (c *ECS) DescribeInstancesPages(input *DescribeInstancesInput, fn func(*DescribeInstancesOutput, bool) bool) error {
	return c.DescribeInstancesPagesWithContext(sdf.BackgroundContext(), input, fn)
}

func (c *ECS) DescribeInstancesPagesWithContext(ctx sdf.Context, input *DescribeInstancesInput, fn func(*DescribeInstancesOutput, bool) bool, opts ...request.Option) error {
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

	InstanceIds *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RegionId *string `locationName:"RegionId" type:"string" enum:"RegionEnum"`
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
func (s *DescribeInstancesInput) SetInstanceIds(v string) *DescribeInstancesInput {
	s.InstanceIds = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeInstancesInput) SetPageNumber(v int64) *DescribeInstancesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeInstancesInput) SetPageSize(v int64) *DescribeInstancesInput {
	s.PageSize = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeInstancesInput) SetRegionId(v string) *DescribeInstancesInput {
	s.RegionId = &v
	return s
}

// Contains the output of DescribeInstances.
type DescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Instances *Instances `type:"structure"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestID *string `locationName:"RequestId" type:"string"`

	// test
	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeInstancesOutput) String() string {
	return sdfutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetInstances sets the Instances field's value.
func (s *DescribeInstancesOutput) SetInstances(v *Instances) *DescribeInstancesOutput {
	s.Instances = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeInstancesOutput) SetPageNumber(v int64) *DescribeInstancesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeInstancesOutput) SetPageSize(v int64) *DescribeInstancesOutput {
	s.PageSize = &v
	return s
}

// SetRequestID sets the RequestID field's value.
func (s *DescribeInstancesOutput) SetRequestID(v string) *DescribeInstancesOutput {
	s.RequestID = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeInstancesOutput) SetTotalCount(v int64) *DescribeInstancesOutput {
	s.TotalCount = &v
	return s
}

type Instances struct {
	_ struct{} `type:"structure"`

	// test
	InstanceList []*Instance `locationName:"Instance" type:"list"`
}

// String returns the string representation
func (s Instances) String() string {
	return sdfutil.Prettify(s)
}

// GoString returns the string representation
func (s Instances) GoString() string {
	return s.String()
}

// SetInstanceList sets the InstanceList field's value.
func (s *Instances) SetInstanceList(v []*Instance) *Instances {
	s.InstanceList = v
	return s
}

// item
type Instance struct {
	_ struct{} `type:"structure"`

	// test
	CPU *int64 `locationName:"Cpu" type:"integer"`

	CreationTime *string `locationName:"CreationTime" type:"string"`

	Description *string `locationName:"Description" type:"string"`

	//EipAddress *EipAddress `locationName:"EipAddress" type:"structure"`

	ExpiredTime *string `locationName:"ExpiredTime" type:"string"`

	HostName *string `locationName:"HostName" type:"string"`

	ImageId *string `locationName:"ImageId" type:"string"`

	//InnerIpAddress *InnerIpAddressList `locationName:"InnerIpAddress" type:"structure"`

	InstanceChargeType *string `type:"string"`

	// test
	InstanceId *string `locationName:"InstanceId" type:"string"`

	InstanceName *string `locationName:"InstanceName" type:"string"`

	InstanceNetworkType *string `type:"string"`

	// test
	InstanceType *string `locationName:"InstanceType" type:"string"`

	InternetChargeType *string `locationName:"InternetChargeType" type:"string"`

	InternetMaxBandwidthIn *int64 `type:"integer"`

	InternetMaxBandwidthOut *int64 `type:"integer"`

	// test
	Memory *int64 `locationName:"Memory" type:"integer"`

	OSName *string `locationName:"OSName" type:"string"`

	OSType *string `locationName:"OSType" type:"string"`

	//PublicIpAddress *PublicIpAddressList `locationName:"PublicIpAddress" type:"structure"`

	RegionId *string `locationName:"RegionId" type:"string" enum:"RegionEnum"`

	// test
	Status *string `locationName:"Status" type:"string"`

	//VpcAttributes *VpcAttributes `type:"structure"`

	ZoneId *string `locationName:"ZoneId" type:"string"`
}

// String returns the string representation
func (s Instance) String() string {
	return sdfutil.Prettify(s)
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

// SetCreationTime sets the CreationTime field's value.
func (s *Instance) SetCreationTime(v string) *Instance {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *Instance) SetDescription(v string) *Instance {
	s.Description = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *Instance) SetExpiredTime(v string) *Instance {
	s.ExpiredTime = &v
	return s
}

// SetHostName sets the HostName field's value.
func (s *Instance) SetHostName(v string) *Instance {
	s.HostName = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *Instance) SetImageId(v string) *Instance {
	s.ImageId = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *Instance) SetInstanceChargeType(v string) *Instance {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *Instance) SetInstanceId(v string) *Instance {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *Instance) SetInstanceName(v string) *Instance {
	s.InstanceName = &v
	return s
}

// SetInstanceNetworkType sets the InstanceNetworkType field's value.
func (s *Instance) SetInstanceNetworkType(v string) *Instance {
	s.InstanceNetworkType = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *Instance) SetInstanceType(v string) *Instance {
	s.InstanceType = &v
	return s
}

// SetInternetChargeType sets the InternetChargeType field's value.
func (s *Instance) SetInternetChargeType(v string) *Instance {
	s.InternetChargeType = &v
	return s
}

// SetInternetMaxBandwidthIn sets the InternetMaxBandwidthIn field's value.
func (s *Instance) SetInternetMaxBandwidthIn(v int64) *Instance {
	s.InternetMaxBandwidthIn = &v
	return s
}

// SetInternetMaxBandwidthOut sets the InternetMaxBandwidthOut field's value.
func (s *Instance) SetInternetMaxBandwidthOut(v int64) *Instance {
	s.InternetMaxBandwidthOut = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *Instance) SetMemory(v int64) *Instance {
	s.Memory = &v
	return s
}

// SetOSName sets the OSName field's value.
func (s *Instance) SetOSName(v string) *Instance {
	s.OSName = &v
	return s
}

// SetOSType sets the OSType field's value.
func (s *Instance) SetOSType(v string) *Instance {
	s.OSType = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *Instance) SetRegionId(v string) *Instance {
	s.RegionId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Instance) SetStatus(v string) *Instance {
	s.Status = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *Instance) SetZoneId(v string) *Instance {
	s.ZoneId = &v
	return s
}

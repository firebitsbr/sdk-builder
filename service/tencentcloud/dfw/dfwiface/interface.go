// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dfwiface provides an interface to enable mocking the SecurityGroup service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dfwiface

import (
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/service/dfw"
)

// DFWAPI provides an interface to enable mocking the
// dfw.DFW service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // SecurityGroup.
//    func myFunc(svc dfwiface.DFWAPI) bool {
//        // Make svc.CreateSecurityGroup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := dfw.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDFWClient struct {
//        dfwiface.DFWAPI
//    }
//    func (m *mockDFWClient) CreateSecurityGroup(input *dfw.CreateSecurityGroupInput) (*dfw.CreateSecurityGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDFWClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DFWAPI interface {
	CreateSecurityGroup(*dfw.CreateSecurityGroupInput) (*dfw.CreateSecurityGroupOutput, error)
	CreateSecurityGroupWithContext(sdf.Context, *dfw.CreateSecurityGroupInput, ...request.Option) (*dfw.CreateSecurityGroupOutput, error)
	CreateSecurityGroupRequest(*dfw.CreateSecurityGroupInput) (*request.Request, *dfw.CreateSecurityGroupOutput)

	CreateSecurityGroupPolicy(*dfw.CreateSecurityGroupPolicyInput) (*dfw.CreateSecurityGroupPolicyOutput, error)
	CreateSecurityGroupPolicyWithContext(sdf.Context, *dfw.CreateSecurityGroupPolicyInput, ...request.Option) (*dfw.CreateSecurityGroupPolicyOutput, error)
	CreateSecurityGroupPolicyRequest(*dfw.CreateSecurityGroupPolicyInput) (*request.Request, *dfw.CreateSecurityGroupPolicyOutput)

	DeleteSecurityGroup(*dfw.DeleteSecurityGroupInput) (*dfw.DeleteSecurityGroupOutput, error)
	DeleteSecurityGroupWithContext(sdf.Context, *dfw.DeleteSecurityGroupInput, ...request.Option) (*dfw.DeleteSecurityGroupOutput, error)
	DeleteSecurityGroupRequest(*dfw.DeleteSecurityGroupInput) (*request.Request, *dfw.DeleteSecurityGroupOutput)

	DeleteSecurityGroupPolicy(*dfw.DeleteSecurityGroupPolicyInput) (*dfw.DeleteSecurityGroupPolicyOutput, error)
	DeleteSecurityGroupPolicyWithContext(sdf.Context, *dfw.DeleteSecurityGroupPolicyInput, ...request.Option) (*dfw.DeleteSecurityGroupPolicyOutput, error)
	DeleteSecurityGroupPolicyRequest(*dfw.DeleteSecurityGroupPolicyInput) (*request.Request, *dfw.DeleteSecurityGroupPolicyOutput)

	DescribeSecurityGroupPolicys(*dfw.DescribeSecurityGroupPolicysInput) (*dfw.DescribeSecurityGroupPolicysOutput, error)
	DescribeSecurityGroupPolicysWithContext(sdf.Context, *dfw.DescribeSecurityGroupPolicysInput, ...request.Option) (*dfw.DescribeSecurityGroupPolicysOutput, error)
	DescribeSecurityGroupPolicysRequest(*dfw.DescribeSecurityGroupPolicysInput) (*request.Request, *dfw.DescribeSecurityGroupPolicysOutput)

	DescribeSecurityGroups(*dfw.DescribeSecurityGroupsInput) (*dfw.DescribeSecurityGroupsOutput, error)
	DescribeSecurityGroupsWithContext(sdf.Context, *dfw.DescribeSecurityGroupsInput, ...request.Option) (*dfw.DescribeSecurityGroupsOutput, error)
	DescribeSecurityGroupsRequest(*dfw.DescribeSecurityGroupsInput) (*request.Request, *dfw.DescribeSecurityGroupsOutput)

	DescribeSecurityGroupsPages(*dfw.DescribeSecurityGroupsInput, func(*dfw.DescribeSecurityGroupsOutput, bool) bool) error
	DescribeSecurityGroupsPagesWithContext(sdf.Context, *dfw.DescribeSecurityGroupsInput, func(*dfw.DescribeSecurityGroupsOutput, bool) bool, ...request.Option) error
}

var _ DFWAPI = (*dfw.DFW)(nil)
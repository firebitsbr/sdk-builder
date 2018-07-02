// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cvmiface provides an interface to enable mocking the Cloud Virtual Machine service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cvmiface

import (
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/service/cvm"
)

// CVMAPI provides an interface to enable mocking the
// cvm.CVM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Cloud Virtual Machine.
//    func myFunc(svc cvmiface.CVMAPI) bool {
//        // Make svc.DescribeImages request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cvm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCVMClient struct {
//        cvmiface.CVMAPI
//    }
//    func (m *mockCVMClient) DescribeImages(input *cvm.DescribeImagesInput) (*cvm.DescribeImagesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCVMClient{}
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
type CVMAPI interface {
	DescribeImages(*cvm.DescribeImagesInput) (*cvm.DescribeImagesOutput, error)
	DescribeImagesWithContext(sdf.Context, *cvm.DescribeImagesInput, ...request.Option) (*cvm.DescribeImagesOutput, error)
	DescribeImagesRequest(*cvm.DescribeImagesInput) (*request.Request, *cvm.DescribeImagesOutput)

	DescribeInstances(*cvm.DescribeInstancesInput) (*cvm.DescribeInstancesOutput, error)
	DescribeInstancesWithContext(sdf.Context, *cvm.DescribeInstancesInput, ...request.Option) (*cvm.DescribeInstancesOutput, error)
	DescribeInstancesRequest(*cvm.DescribeInstancesInput) (*request.Request, *cvm.DescribeInstancesOutput)

	DescribeInstancesPages(*cvm.DescribeInstancesInput, func(*cvm.DescribeInstancesOutput, bool) bool) error
	DescribeInstancesPagesWithContext(sdf.Context, *cvm.DescribeInstancesInput, func(*cvm.DescribeInstancesOutput, bool) bool, ...request.Option) error

	DescribeInstancesStatus(*cvm.DescribeInstancesStatusInput) (*cvm.DescribeInstancesStatusOutput, error)
	DescribeInstancesStatusWithContext(sdf.Context, *cvm.DescribeInstancesStatusInput, ...request.Option) (*cvm.DescribeInstancesStatusOutput, error)
	DescribeInstancesStatusRequest(*cvm.DescribeInstancesStatusInput) (*request.Request, *cvm.DescribeInstancesStatusOutput)

	RebootInstances(*cvm.RebootInstancesInput) (*cvm.RebootInstancesOutput, error)
	RebootInstancesWithContext(sdf.Context, *cvm.RebootInstancesInput, ...request.Option) (*cvm.RebootInstancesOutput, error)
	RebootInstancesRequest(*cvm.RebootInstancesInput) (*request.Request, *cvm.RebootInstancesOutput)

	RunInstances(*cvm.RunInstancesInput) (*cvm.RunInstancesOutput, error)
	RunInstancesWithContext(sdf.Context, *cvm.RunInstancesInput, ...request.Option) (*cvm.RunInstancesOutput, error)
	RunInstancesRequest(*cvm.RunInstancesInput) (*request.Request, *cvm.RunInstancesOutput)

	StartInstances(*cvm.StartInstancesInput) (*cvm.StartInstancesOutput, error)
	StartInstancesWithContext(sdf.Context, *cvm.StartInstancesInput, ...request.Option) (*cvm.StartInstancesOutput, error)
	StartInstancesRequest(*cvm.StartInstancesInput) (*request.Request, *cvm.StartInstancesOutput)

	StopInstances(*cvm.StopInstancesInput) (*cvm.StopInstancesOutput, error)
	StopInstancesWithContext(sdf.Context, *cvm.StopInstancesInput, ...request.Option) (*cvm.StopInstancesOutput, error)
	StopInstancesRequest(*cvm.StopInstancesInput) (*request.Request, *cvm.StopInstancesOutput)

	TerminateInstances(*cvm.TerminateInstancesInput) (*cvm.TerminateInstancesOutput, error)
	TerminateInstancesWithContext(sdf.Context, *cvm.TerminateInstancesInput, ...request.Option) (*cvm.TerminateInstancesOutput, error)
	TerminateInstancesRequest(*cvm.TerminateInstancesInput) (*request.Request, *cvm.TerminateInstancesOutput)

	WaitUntilInstanceRestart(*cvm.DescribeInstancesStatusInput) error
	WaitUntilInstanceRestartWithContext(sdf.Context, *cvm.DescribeInstancesStatusInput, ...request.WaiterOption) error

	WaitUntilInstanceRunning(*cvm.DescribeInstancesStatusInput) error
	WaitUntilInstanceRunningWithContext(sdf.Context, *cvm.DescribeInstancesStatusInput, ...request.WaiterOption) error

	WaitUntilInstanceStopped(*cvm.DescribeInstancesStatusInput) error
	WaitUntilInstanceStoppedWithContext(sdf.Context, *cvm.DescribeInstancesStatusInput, ...request.WaiterOption) error
}

var _ CVMAPI = (*cvm.CVM)(nil)
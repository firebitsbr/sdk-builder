// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cvm

import (
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/request"
)

// WaitUntilInstanceRestart uses the CVM API operation
// DescribeInstancesStatus to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CVM) WaitUntilInstanceRestart(input *DescribeInstancesStatusInput) error {
	return c.WaitUntilInstanceRestartWithContext(sdf.BackgroundContext(), input)
}

// WaitUntilInstanceRestartWithContext is an extended version of WaitUntilInstanceRestart.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CVM) WaitUntilInstanceRestartWithContext(ctx sdf.Context, input *DescribeInstancesStatusInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceRestart",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "RUNNING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "PENDING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "EXPIRED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInstancesStatusInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstancesStatusRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceRunning uses the CVM API operation
// DescribeInstancesStatus to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CVM) WaitUntilInstanceRunning(input *DescribeInstancesStatusInput) error {
	return c.WaitUntilInstanceRunningWithContext(sdf.BackgroundContext(), input)
}

// WaitUntilInstanceRunningWithContext is an extended version of WaitUntilInstanceRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CVM) WaitUntilInstanceRunningWithContext(ctx sdf.Context, input *DescribeInstancesStatusInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceRunning",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "RUNNING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "TERMINATING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "STOPPING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "EXPIRED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInstancesStatusInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstancesStatusRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceStopped uses the CVM API operation
// DescribeInstancesStatus to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CVM) WaitUntilInstanceStopped(input *DescribeInstancesStatusInput) error {
	return c.WaitUntilInstanceStoppedWithContext(sdf.BackgroundContext(), input)
}

// WaitUntilInstanceStoppedWithContext is an extended version of WaitUntilInstanceStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CVM) WaitUntilInstanceStoppedWithContext(ctx sdf.Context, input *DescribeInstancesStatusInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceStopped",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "STOPPED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "TERMINATING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "STARTING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Response.InstanceStatusSet[].InstanceState",
				Expected: "REBOOTING",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInstancesStatusInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstancesStatusRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
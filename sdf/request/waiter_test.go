package request_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/client"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/polefishu/sdk-builder/sdftesting"
	"github.com/polefishu/sdk-builder/sdftesting/unit"
	"github.com/aws/aws-sdk-go/service/s3"
)

type mockClient struct {
	*client.Client
}
type MockInput struct{}
type MockOutput struct {
	States []*MockState
}
type MockState struct {
	State *string
}

func (c *mockClient) MockRequest(input *MockInput) (*request.Request, *MockOutput) {
	op := &request.Operation{
		Name:       "Mock",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MockInput{}
	}

	output := &MockOutput{}
	req := c.NewRequest(op, input, output)
	req.Data = output
	return req, output
}

func BuildNewMockRequest(c *mockClient, in *MockInput) func([]request.Option) (*request.Request, error) {
	return func(opts []request.Option) (*request.Request, error) {
		req, _ := c.MockRequest(in)
		req.ApplyOptions(opts...)
		return req, nil
	}
}

func TestWaiterPathAll(t *testing.T) {
	svc := &mockClient{Client: awstesting.NewClient(&sdf.Config{
		Region: sdf.String("mock-region"),
	})}
	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()

	reqNum := 0
	resps := []*MockOutput{
		{ // Request 1
			States: []*MockState{
				{State: sdf.String("pending")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 2
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 3
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("running")},
			},
		},
	}

	numBuiltReq := 0
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		numBuiltReq++
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		if reqNum >= len(resps) {
			assert.Fail(t, "too many polling requests made")
			return
		}
		r.Data = resps[reqNum]
		reqNum++
	})

	w := request.Waiter{
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(0),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.PathAllWaiterMatch,
				Argument: "States[].State",
				Expected: "running",
			},
		},
		NewRequest: BuildNewMockRequest(svc, &MockInput{}),
	}

	err := w.WaitWithContext(sdf.BackgroundContext())
	assert.NoError(t, err)
	assert.Equal(t, 3, numBuiltReq)
	assert.Equal(t, 3, reqNum)
}

func TestWaiterPath(t *testing.T) {
	svc := &mockClient{Client: awstesting.NewClient(&sdf.Config{
		Region: sdf.String("mock-region"),
	})}
	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()

	reqNum := 0
	resps := []*MockOutput{
		{ // Request 1
			States: []*MockState{
				{State: sdf.String("pending")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 2
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 3
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("running")},
			},
		},
	}

	numBuiltReq := 0
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		numBuiltReq++
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		if reqNum >= len(resps) {
			assert.Fail(t, "too many polling requests made")
			return
		}
		r.Data = resps[reqNum]
		reqNum++
	})

	w := request.Waiter{
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(0),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.PathWaiterMatch,
				Argument: "States[].State",
				Expected: "running",
			},
		},
		NewRequest: BuildNewMockRequest(svc, &MockInput{}),
	}

	err := w.WaitWithContext(sdf.BackgroundContext())
	assert.NoError(t, err)
	assert.Equal(t, 3, numBuiltReq)
	assert.Equal(t, 3, reqNum)
}

func TestWaiterFailure(t *testing.T) {
	svc := &mockClient{Client: awstesting.NewClient(&sdf.Config{
		Region: sdf.String("mock-region"),
	})}
	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()

	reqNum := 0
	resps := []*MockOutput{
		{ // Request 1
			States: []*MockState{
				{State: sdf.String("pending")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 2
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 3
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("stopping")},
			},
		},
	}

	numBuiltReq := 0
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		numBuiltReq++
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		if reqNum >= len(resps) {
			assert.Fail(t, "too many polling requests made")
			return
		}
		r.Data = resps[reqNum]
		reqNum++
	})

	w := request.Waiter{
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(0),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.PathAllWaiterMatch,
				Argument: "States[].State",
				Expected: "running",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.PathAnyWaiterMatch,
				Argument: "States[].State",
				Expected: "stopping",
			},
		},
		NewRequest: BuildNewMockRequest(svc, &MockInput{}),
	}

	err := w.WaitWithContext(sdf.BackgroundContext()).(sdferr.Error)
	assert.Error(t, err)
	assert.Equal(t, request.WaiterResourceNotReadyErrorCode, err.Code())
	assert.Equal(t, "failed waiting for successful resource state", err.Message())
	assert.Equal(t, 3, numBuiltReq)
	assert.Equal(t, 3, reqNum)
}

func TestWaiterError(t *testing.T) {
	svc := &mockClient{Client: awstesting.NewClient(&sdf.Config{
		Region: sdf.String("mock-region"),
	})}
	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.UnmarshalError.Clear()
	svc.Handlers.ValidateResponse.Clear()

	reqNum := 0
	resps := []*MockOutput{
		{ // Request 1
			States: []*MockState{
				{State: sdf.String("pending")},
				{State: sdf.String("pending")},
			},
		},
		{ // Request 1, error case retry
		},
		{ // Request 2, error case failure
		},
		{ // Request 3
			States: []*MockState{
				{State: sdf.String("running")},
				{State: sdf.String("running")},
			},
		},
	}
	reqErrs := make([]error, len(resps))
	reqErrs[1] = sdferr.New("MockException", "mock exception message", nil)
	reqErrs[2] = sdferr.New("FailureException", "mock failure exception message", nil)

	numBuiltReq := 0
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		numBuiltReq++
	})
	svc.Handlers.Send.PushBack(func(r *request.Request) {
		code := 200
		if reqNum == 1 {
			code = 400
		}
		r.HTTPResponse = &http.Response{
			StatusCode: code,
			Status:     http.StatusText(code),
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		if reqNum >= len(resps) {
			assert.Fail(t, "too many polling requests made")
			return
		}
		r.Data = resps[reqNum]
		reqNum++
	})
	svc.Handlers.UnmarshalMeta.PushBack(func(r *request.Request) {
		// If there was an error unmarshal error will be called instead of unmarshal
		// need to increment count here also
		if err := reqErrs[reqNum]; err != nil {
			r.Error = err
			reqNum++
		}
	})

	w := request.Waiter{
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(0),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.PathAllWaiterMatch,
				Argument: "States[].State",
				Expected: "running",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Argument: "",
				Expected: "MockException",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Argument: "",
				Expected: "FailureException",
			},
		},
		NewRequest: BuildNewMockRequest(svc, &MockInput{}),
	}

	err := w.WaitWithContext(sdf.BackgroundContext())
	if err == nil {
		t.Fatalf("expected error, but did not get one")
	}
	aerr := err.(sdferr.Error)
	if e, a := request.WaiterResourceNotReadyErrorCode, aerr.Code(); e != a {
		t.Errorf("expect %q error code, got %q", e, a)
	}
	if e, a := 3, numBuiltReq; e != a {
		t.Errorf("expect %d built requests got %d", e, a)
	}
	if e, a := 3, reqNum; e != a {
		t.Errorf("expect %d reqNum got %d", e, a)
	}
}

func TestWaiterStatus(t *testing.T) {
	svc := &mockClient{Client: awstesting.NewClient(&sdf.Config{
		Region: sdf.String("mock-region"),
	})}
	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()

	reqNum := 0
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		reqNum++
	})
	svc.Handlers.Send.PushBack(func(r *request.Request) {
		code := 200
		if reqNum == 3 {
			code = 404
			r.Error = sdferr.New("NotFound", "Not Found", nil)
		}
		r.HTTPResponse = &http.Response{
			StatusCode: code,
			Status:     http.StatusText(code),
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	})

	w := request.Waiter{
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(0),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Argument: "",
				Expected: 404,
			},
		},
		NewRequest: BuildNewMockRequest(svc, &MockInput{}),
	}

	err := w.WaitWithContext(sdf.BackgroundContext())
	assert.NoError(t, err)
	assert.Equal(t, 3, reqNum)
}

func TestWaiter_ApplyOptions(t *testing.T) {
	w := request.Waiter{}

	logger := sdf.NewDefaultLogger()

	w.ApplyOptions(
		request.WithWaiterLogger(logger),
		request.WithWaiterRequestOptions(request.WithLogLevel(sdf.LogDebug)),
		request.WithWaiterMaxAttempts(2),
		request.WithWaiterDelay(request.ConstantWaiterDelay(5*time.Second)),
	)

	if e, a := logger, w.Logger; e != a {
		t.Errorf("expect logger to be set, and match, was not, %v, %v", e, a)
	}

	if len(w.RequestOptions) != 1 {
		t.Fatalf("expect request options to be set to only a single option, %v", w.RequestOptions)
	}
	r := request.Request{}
	r.ApplyOptions(w.RequestOptions...)
	if e, a := sdf.LogDebug, r.Config.LogLevel.Value(); e != a {
		t.Errorf("expect %v loglevel got %v", e, a)
	}

	if e, a := 2, w.MaxAttempts; e != a {
		t.Errorf("expect %d retryer max attempts, got %d", e, a)
	}
	if e, a := 5*time.Second, w.Delay(0); e != a {
		t.Errorf("expect %d retryer delay, got %d", e, a)
	}
}

func TestWaiter_WithContextCanceled(t *testing.T) {
	c := awstesting.NewClient()

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	reqCount := 0

	w := request.Waiter{
		Name:             "TestWaiter",
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(1 * time.Millisecond),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
		},
		Logger: sdf.NewDefaultLogger(),
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req := c.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
			req.HTTPResponse = &http.Response{StatusCode: http.StatusNotFound}
			req.Handlers.Clear()
			req.Data = struct{}{}
			req.Handlers.Send.PushBack(func(r *request.Request) {
				if reqCount == 1 {
					ctx.Error = fmt.Errorf("context canceled")
					close(ctx.DoneCh)
				}
				reqCount++
			})

			return req, nil
		},
	}

	w.SleepWithContext = func(c sdf.Context, delay time.Duration) error {
		context := c.(*awstesting.FakeContext)
		select {
		case <-context.DoneCh:
			return context.Err()
		default:
			return nil
		}
	}

	err := w.WaitWithContext(ctx)

	if err == nil {
		t.Fatalf("expect waiter to be canceled.")
	}
	aerr := err.(sdferr.Error)
	if e, a := request.CanceledErrorCode, aerr.Code(); e != a {
		t.Errorf("expect %q error code, got %q", e, a)
	}
	if e, a := 2, reqCount; e != a {
		t.Errorf("expect %d requests, got %d", e, a)
	}
}

func TestWaiter_WithContext(t *testing.T) {
	c := awstesting.NewClient()

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	reqCount := 0

	statuses := []int{http.StatusNotFound, http.StatusOK}

	w := request.Waiter{
		Name:             "TestWaiter",
		MaxAttempts:      10,
		Delay:            request.ConstantWaiterDelay(1 * time.Millisecond),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
		},
		Logger: sdf.NewDefaultLogger(),
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req := c.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
			req.HTTPResponse = &http.Response{StatusCode: statuses[reqCount]}
			req.Handlers.Clear()
			req.Data = struct{}{}
			req.Handlers.Send.PushBack(func(r *request.Request) {
				if reqCount == 1 {
					ctx.Error = fmt.Errorf("context canceled")
					close(ctx.DoneCh)
				}
				reqCount++
			})

			return req, nil
		},
	}

	err := w.WaitWithContext(ctx)

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := 2, reqCount; e != a {
		t.Errorf("expect %d requests, got %d", e, a)
	}
}

func TestWaiter_AttemptsExpires(t *testing.T) {
	c := awstesting.NewClient()

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	reqCount := 0

	w := request.Waiter{
		Name:             "TestWaiter",
		MaxAttempts:      2,
		Delay:            request.ConstantWaiterDelay(1 * time.Millisecond),
		SleepWithContext: sdf.SleepWithContext,
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
		},
		Logger: sdf.NewDefaultLogger(),
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req := c.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
			req.HTTPResponse = &http.Response{StatusCode: http.StatusNotFound}
			req.Handlers.Clear()
			req.Data = struct{}{}
			req.Handlers.Send.PushBack(func(r *request.Request) {
				reqCount++
			})

			return req, nil
		},
	}

	err := w.WaitWithContext(ctx)

	if err == nil {
		t.Fatalf("expect error did not get one")
	}
	aerr := err.(sdferr.Error)
	if e, a := request.WaiterResourceNotReadyErrorCode, aerr.Code(); e != a {
		t.Errorf("expect %q error code, got %q", e, a)
	}
	if e, a := 2, reqCount; e != a {
		t.Errorf("expect %d requests, got %d", e, a)
	}
}

func TestWaiterNilInput(t *testing.T) {
	// Code generation doesn't have a great way to verify the code is correct
	// other than being run via unit tests in the SDK. This should be fixed
	// So code generation can be validated independently.

	client := s3.New(unit.Session)
	client.Handlers.Validate.Clear()
	client.Handlers.Send.Clear() // mock sending
	client.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: http.StatusOK,
		}
	})
	client.Handlers.Unmarshal.Clear()
	client.Handlers.UnmarshalMeta.Clear()
	client.Handlers.ValidateResponse.Clear()
	client.Config.SleepDelay = func(dur time.Duration) {}

	// Ensure waiters do not panic on nil input. It doesn't make sense to
	// call a waiter without an input, Validation will
	err := client.WaitUntilBucketExists(nil)
	if err != nil {
		t.Fatalf("expect no error, but got %v", err)
	}
}

func TestWaiterWithContextNilInput(t *testing.T) {
	// Code generation doesn't have a great way to verify the code is correct
	// other than being run via unit tests in the SDK. This should be fixed
	// So code generation can be validated independently.

	client := s3.New(unit.Session)
	client.Handlers.Validate.Clear()
	client.Handlers.Send.Clear() // mock sending
	client.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: http.StatusOK,
		}
	})
	client.Handlers.Unmarshal.Clear()
	client.Handlers.UnmarshalMeta.Clear()
	client.Handlers.ValidateResponse.Clear()

	// Ensure waiters do not panic on nil input
	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	err := client.WaitUntilBucketExistsWithContext(ctx, nil,
		request.WithWaiterDelay(request.ConstantWaiterDelay(0)),
		request.WithWaiterMaxAttempts(1),
	)
	if err != nil {
		t.Fatalf("expect no error, but got %v", err)
	}
}

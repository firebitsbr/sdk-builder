package sdf_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdftesting"
)

func TestSleepWithContext(t *testing.T) {
	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}

	err := sdf.SleepWithContext(ctx, 1*time.Millisecond)
	if err != nil {
		t.Errorf("expect context to not be canceled, got %v", err)
	}
}

func TestSleepWithContext_Canceled(t *testing.T) {
	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}

	expectErr := fmt.Errorf("context canceled")

	ctx.Error = expectErr
	close(ctx.DoneCh)

	err := sdf.SleepWithContext(ctx, 1*time.Millisecond)
	if err == nil {
		t.Fatalf("expect error, did not get one")
	}

	if e, a := expectErr, err; e != a {
		t.Errorf("expect %v error, got %v", e, a)
	}
}

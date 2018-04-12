package ecs_test

import (
	"testing"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/session"
	"github.com/polefishu/sdk-builder/service/alicloud/ecs"
)

func TestDescribeInstances(t *testing.T) {
	region := "cn-shanghai"
	config := sdf.NewConfig()
	config.MaxRetries = sdf.Int(1)
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)
	sess := session.Must(session.NewSession())

	svc := ecs.New(sess, config)
	input := &ecs.DescribeInstancesInput{
		RegionId: sdf.String(region),
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(sdferr.Error); ok {
			switch aerr.Code() {
			default:
				t.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to sdferr.Error to get the Code and
			// Message from an error.
			t.Error(err.Error())
		}
		return
	}

	t.Log(result)
}

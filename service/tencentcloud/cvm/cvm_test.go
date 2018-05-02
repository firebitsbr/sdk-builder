package cvm_test

import (
	"testing"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/session"
	"github.com/polefishu/sdk-builder/service/tencentcloud/cvm"
)

func TestDescribeInstances(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cvm.New(sess, config)

	input := &cvm.DescribeInstancesInput{
		Region: sdf.String(region),
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(sdferr.Error); ok {
			switch aerr.Code() {
			default:
				t.Error(aerr.Error())
			}
		} else {
			t.Error(err.Error())
		}
		return
	}

	t.Log(result)
}

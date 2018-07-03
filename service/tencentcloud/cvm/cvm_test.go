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
	config.Endpoint = sdf.String("cvm.tencentcloudapi.com")

	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cvm.New(sess, config)

	input := &cvm.DescribeInstancesInput{}

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

func TestDescribeInstancesStatus(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("cvm.tencentcloudapi.com")

	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cvm.New(sess, config)

	input := &cvm.DescribeInstancesStatusInput{}

	result, err := svc.DescribeInstancesStatus(input)
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

func TestRunInstances(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("cvm.tencentcloudapi.com")

	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cvm.New(sess, config)

	input := &cvm.RunInstancesInput{
		ImageId:   sdf.String("img-dkwyg6sr"),
		Placement: &cvm.Placement{Zone: sdf.String("ap-shanghai-1")},
	}

	result, err := svc.RunInstances(input)
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

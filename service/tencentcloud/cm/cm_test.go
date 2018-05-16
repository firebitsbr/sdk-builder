package cm_test

import (
	"testing"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/session"
	"github.com/polefishu/sdk-builder/service/tencentcloud/cm"
)

func TestDescribeBaseMetrics(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cm.New(sess, config)

	input := &cm.DescribeBaseMetricsInput{
		Region:    sdf.String(region),
		Namespace: sdf.String("qce/cvm"),
	}

	result, err := svc.DescribeBaseMetrics(input)
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

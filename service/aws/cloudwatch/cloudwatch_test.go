package cloudwatch_test

import (
	"testing"
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/session"
	"github.com/polefishu/sdk-builder/service/aws/cloudwatch"
)

func TestDescribeBaseMetrics(t *testing.T) {
	region := "ap-northeast-1"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cloudwatch.New(sess, config)
	start := time.Unix(1526515200, 0)
	end := time.Unix(1526526000, 0)
	input := &cloudwatch.GetMetricDataInput{
		StartTime: &start,
		EndTime:   &end,
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			&cloudwatch.MetricDataQuery{
				Id: sdf.String("m1"),
				MetricStat: &cloudwatch.MetricStat{
					Period: sdf.Int64(60),
					Stat:   sdf.String("Average"),
					// Unit:   sdf.String("Count"),
					Metric: &cloudwatch.Metric{
						Namespace:  sdf.String("AWS/EC2"),
						MetricName: sdf.String("CPUUtilization"),
						Dimensions: []*cloudwatch.Dimension{
							&cloudwatch.Dimension{
								Name:  sdf.String("InstanceId"),
								Value: sdf.String("i-0b2f996b647f961d3"),
							},
						},
					},
				},
			},
		},
	}

	result, err := svc.GetMetricData(input)
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

func TestGetMetricStatistics(t *testing.T) {
	region := "ap-northeast-1"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := cloudwatch.New(sess, config)
	start := time.Unix(1526515200, 0)
	end := time.Unix(1526526000, 0)
	input := &cloudwatch.GetMetricStatisticsInput{
		StartTime:  &start,
		EndTime:    &end,
		Period:     sdf.Int64(60),
		Statistics: []*string{sdf.String("Sum")},
		Namespace:  sdf.String("AWS/EC2"),
		MetricName: sdf.String("CPUUtilization"),
		Dimensions: []*cloudwatch.Dimension{
			&cloudwatch.Dimension{
				Name:  sdf.String("InstanceId"),
				Value: sdf.String("i-0b2f996b647f961d3"),
			},
		},
	}

	result, err := svc.GetMetricStatistics(input)
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

package vpc_test

import (
	"testing"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/session"
	"github.com/polefishu/sdk-builder/service/tencentcloud/vpc"
)

func TestDescribeSecurityGroups(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.DescribeSecurityGroupsInput{}

	result, err := svc.DescribeSecurityGroups(input)
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

func TestDescribeSecurityGroupPolicies(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.DescribeSecurityGroupPoliciesInput{
		SecurityGroupId: sdf.String("sg-bfkg4y2u"),
	}

	result, err := svc.DescribeSecurityGroupPolicies(input)
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

func TestCreateSecurityGroupPolicies(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.CreateSecurityGroupPoliciesInput{
		SecurityGroupId: sdf.String("sg-bfkg4y2u"),
		SecurityGroupPolicySet: &vpc.SecurityGroupPolicyData{
			// Version: sdf.Int64(10),
			Ingress: []*vpc.SecurityGroupPolicy{
				&vpc.SecurityGroupPolicy{
					Action:      sdf.String("accept"),
					PolicyIndex: sdf.Int64(0),
					CidrBlock:   sdf.String("127.0.0.1"),
					Port:        sdf.String("80"),
					Protocol:    sdf.String("TCP"),
				},
			},
		},
	}

	result, err := svc.CreateSecurityGroupPolicies(input)
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

func TestDescribeVpcs(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.DescribeVpcsInput{}

	result, err := svc.DescribeVpcs(input)
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

func TestCreateVpc(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.CreateVpcInput{
		CidrBlock: sdf.String("10.0.0.0/16"),
		VpcName:   sdf.String("delshu"),
	}

	result, err := svc.CreateVpc(input)
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

func TestDeleteVpc(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.DeleteVpcInput{
		VpcId: sdf.String("vpc-5ovrekkm"),
	}

	result, err := svc.DeleteVpc(input)
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

func TestDescribeSubnets(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.DescribeSubnetsInput{
		SubnetIds: []*string{sdf.String("subnet-fnlcxbar")},
	}

	result, err := svc.DescribeSubnets(input)
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

func TestCreateSubnet(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.CreateSubnetInput{
		VpcId:      sdf.String("vpc-1ihrk430"),
		CidrBlock:  sdf.String("10.0.1.0/24"),
		SubnetName: sdf.String("1ihrk430"),
		Zone:       sdf.String("ap-shanghai-1"),
	}

	result, err := svc.CreateSubnet(input)
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

func TestDeleteSubnet(t *testing.T) {
	region := "ap-shanghai"
	config := sdf.NewConfig()
	config.Region = sdf.String(region)
	config.MaxRetries = sdf.Int(1)
	config.Endpoint = sdf.String("vpc.tencentcloudapi.com")
	config.WithLogLevel(sdf.LogDebugWithHTTPBody)

	sess := session.Must(session.NewSession())

	svc := vpc.New(sess, config)

	input := &vpc.DeleteSubnetInput{
		SubnetId: sdf.String("subnet-9pby212z"),
	}

	result, err := svc.DeleteSubnet(input)
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

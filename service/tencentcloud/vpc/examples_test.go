// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdf/sdferr"
	"github.com/polefishu/sdk-builder/sdf/session"
	"github.com/polefishu/sdk-builder/service/vpc"
)

var _ time.Duration
var _ strings.Reader
var _ sdf.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To describe VirtualPrivateCloud
//
// This example DescribeVpc.
func ExampleVPC_DescribeVpc_shared00() {
	svc := vpc.New(session.New())
	input := &vpc.DescribeVpcInput{}

	result, err := svc.DescribeVpc(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

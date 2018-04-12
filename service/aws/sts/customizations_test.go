package sts_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/polefishu/sdk-builder/sdf"
	"github.com/polefishu/sdk-builder/sdftesting/unit"
)

var svc = sts.New(unit.Session, &sdf.Config{
	Region: sdf.String("mock-region"),
})

func TestUnsignedRequest_AssumeRoleWithSAML(t *testing.T) {
	req, _ := svc.AssumeRoleWithSAMLRequest(&sts.AssumeRoleWithSAMLInput{
		PrincipalArn:  sdf.String("ARN01234567890123456789"),
		RoleArn:       sdf.String("ARN01234567890123456789"),
		SAMLAssertion: sdf.String("ASSERT"),
	})

	err := req.Sign()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "", req.HTTPRequest.Header.Get("Authorization"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestUnsignedRequest_AssumeRoleWithWebIdentity(t *testing.T) {
	req, _ := svc.AssumeRoleWithWebIdentityRequest(&sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          sdf.String("ARN01234567890123456789"),
		RoleSessionName:  sdf.String("SESSION"),
		WebIdentityToken: sdf.String("TOKEN"),
	})

	err := req.Sign()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "", req.HTTPRequest.Header.Get("Authorization"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

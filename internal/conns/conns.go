package conns

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	awsbasev1 "github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/version"
)

// ServicePackage is the minimal interface exported from each AWS service package.
// Its methods return the Plugin SDK and Framework resources and data sources implemented in the package.
type ServicePackage interface {
	FrameworkDataSources(context.Context) []*types.ServicePackageFrameworkDataSource
	FrameworkResources(context.Context) []*types.ServicePackageFrameworkResource
	SDKDataSources(context.Context) []*types.ServicePackageSDKDataSource
	SDKResources(context.Context) []*types.ServicePackageSDKResource
	ServicePackageName() string
}

func NewSessionForRegion(cfg *aws.Config, region, terraformVersion string) (*session.Session, error) {
	session, err := session.NewSession(cfg)

	if err != nil {
		return nil, err
	}

	apnInfo := StdUserAgentProducts(terraformVersion)

	awsbasev1.SetSessionUserAgent(session, apnInfo, awsbase.UserAgentProducts{})

	return session.Copy(&aws.Config{Region: aws.String(region)}), nil
}

func StdUserAgentProducts(terraformVersion string) *awsbase.APNInfo {
	return &awsbase.APNInfo{
		PartnerName: "HashiCorp",
		Products: []awsbase.UserAgentProduct{
			{Name: "Terraform", Version: terraformVersion, Comment: "+https://www.terraform.io"},
			{Name: "terraform-provider-aws", Version: version.ProviderVersion, Comment: "+https://registry.terraform.io/providers/hashicorp/aws"},
		},
	}
}

// ReverseDNS switches a DNS hostname to reverse DNS and vice-versa.
func ReverseDNS(hostname string) string {
	parts := strings.Split(hostname, ".")

	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return strings.Join(parts, ".")
}

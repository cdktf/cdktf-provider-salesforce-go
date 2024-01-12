// Deprecated: HashiCorp is no longer publishing new versions of the prebuilt provider for salesforce.
// Previously-published versions of this prebuilt provider will still continue to be available as installable Go modules,
// but these will not be compatible with newer versions of CDK for Terraform and are not eligible for commercial support.
// You can continue to use the salesforce provider in your CDK for Terraform projects with newer versions of CDKTF,
// but you will need to generate the bindings locally. See https://cdk.tf/imports for details.
module github.com/cdktf/cdktf-provider-salesforce-go/salesforce/v8

go 1.18

require (
	github.com/aws/jsii-runtime-go v1.93.0
	github.com/hashicorp/terraform-cdk-go/cdktf v0.20.0
	github.com/aws/constructs-go/constructs/v10 v10.3.0
)

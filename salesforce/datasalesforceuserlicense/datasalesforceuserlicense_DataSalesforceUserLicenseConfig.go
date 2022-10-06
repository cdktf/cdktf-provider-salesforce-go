package datasalesforceuserlicense

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSalesforceUserLicenseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A string that uniquely identifies a particular user license.
	//
	// Valid options vary depending on organization type and configuration. For a complete list see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_userlicense.htm
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/salesforce/d/user_license#license_definition_key DataSalesforceUserLicense#license_definition_key}
	LicenseDefinitionKey *string `field:"required" json:"licenseDefinitionKey" yaml:"licenseDefinitionKey"`
}


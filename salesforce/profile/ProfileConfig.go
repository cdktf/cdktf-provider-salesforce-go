package profile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The name of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/profile#name Profile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the UserLicense associated with this profile. Forces replacement if updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/profile#user_license_id Profile#user_license_id}
	UserLicenseId *string `field:"required" json:"userLicenseId" yaml:"userLicenseId"`
	// Description of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/profile#description Profile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Map of permissions for the profile.
	//
	// At this time specific permissions can only be set, the comprehensive list will not be read from Salesforce. The keys should follow Salesforce 'SnakeCase' format however the 'Permissions' prefix should be omitted. Permissions will not import to state due to a technical limitation, you will need to run a subsequent apply if you have permissions set in config during import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/profile#permissions Profile#permissions}
	Permissions *map[string]interface{} `field:"optional" json:"permissions" yaml:"permissions"`
}


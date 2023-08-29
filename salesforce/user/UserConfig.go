// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// The user’s alias. For example, jsmith.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#alias User#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The user’s email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#email User#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The user’s last name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#last_name User#last_name}
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// ID of the user’s Profile. Use this value to cache metadata based on profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#profile_id User#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Contains the name that a user enters to log in to the API or the user interface.
	//
	// The value for this field must be in the form of an email address, using all lowercase characters. It must also be unique across all organizations. If you try to create or update a User with a duplicate value for this field, the operation is rejected. Each inserted User also counts as a license. Every organization has a maximum number of licenses. If you attempt to exceed the maximum number of licenses by inserting User records, the create request is rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#username User#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The email encoding for the user, such as ISO-8859-1 or UTF-8. Defaults to UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#email_encoding_key User#email_encoding_key}
	EmailEncodingKey *string `field:"optional" json:"emailEncodingKey" yaml:"emailEncodingKey"`
	// The user’s language. Defaults to en_US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#language_locale_key User#language_locale_key}
	LanguageLocaleKey *string `field:"optional" json:"languageLocaleKey" yaml:"languageLocaleKey"`
	// The value of the field affects formatting and parsing of values, especially numeric values, in the user interface.
	//
	// It doesn’t affect the API. The field values are named according to the language, and the country if necessary, using two-letter ISO codes. The set of names is based on the ISO standard. You can also manually set a user’s locale in the user interface, and then use that value for inserting or updating other users via the API. Defaults to en_US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#locale_sid_key User#locale_sid_key}
	LocaleSidKey *string `field:"optional" json:"localeSidKey" yaml:"localeSidKey"`
	// Reset password and send an email to the user.
	//
	// No reset is performed if this field is omitted, is false, or was true and remained true on subsequent apply. Please set to false and then true in subsequent applies, or have it set to true on create to trigger the reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#reset_password User#reset_password}
	ResetPassword interface{} `field:"optional" json:"resetPassword" yaml:"resetPassword"`
	// A User time zone affects the offset used when displaying or entering times in the user interface.
	//
	// But the API doesn’t use a User time zone when querying or setting values. Values for this field are named using region and key city, according to ISO standards. You can also manually set one User time zone in the user interface, and then use that value for creating or updating other User records via the API. Defaults to America/New_York.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#time_zone_sid_key User#time_zone_sid_key}
	TimeZoneSidKey *string `field:"optional" json:"timeZoneSidKey" yaml:"timeZoneSidKey"`
	// ID of the user’s UserRole.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs/resources/user#user_role_id User#user_role_id}
	UserRoleId *string `field:"optional" json:"userRoleId" yaml:"userRoleId"`
}


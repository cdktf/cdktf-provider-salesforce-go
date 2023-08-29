// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package profile

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Profile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Profile) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProfile_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateProfile_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetPermissionsParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Profile) validateSetUserLicenseIdParameters(val *string) error {
	return nil
}

func validateNewProfileParameters(scope constructs.Construct, id *string, config *ProfileConfig) error {
	return nil
}


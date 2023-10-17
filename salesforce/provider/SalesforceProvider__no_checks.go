// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SalesforceProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SalesforceProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSalesforceProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSalesforceProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSalesforceProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSalesforceProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewSalesforceProviderParameters(scope constructs.Construct, id *string, config *SalesforceProviderConfig) error {
	return nil
}


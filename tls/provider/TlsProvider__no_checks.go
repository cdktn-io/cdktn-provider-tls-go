// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TlsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TlsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTlsProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTlsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTlsProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTlsProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TlsProvider) validateSetProxyParameters(val interface{}) error {
	return nil
}

func validateNewTlsProviderParameters(scope constructs.Construct, id *string, config *TlsProviderConfig) error {
	return nil
}


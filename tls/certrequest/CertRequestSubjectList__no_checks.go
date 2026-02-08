// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package certrequest

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertRequestSubjectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CertRequestSubjectList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertRequestSubjectList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertRequestSubjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CertRequestSubjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertRequestSubjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertRequestSubjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertRequestSubjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


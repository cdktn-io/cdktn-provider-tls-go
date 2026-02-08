// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package selfsignedcert

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SelfSignedCertSubjectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SelfSignedCertSubjectList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SelfSignedCertSubjectList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SelfSignedCertSubjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SelfSignedCertSubjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SelfSignedCertSubjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SelfSignedCertSubjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSelfSignedCertSubjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


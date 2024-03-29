// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiAwsCmekSpec AwsCmekSpec
//
// AwsCmekSpec is the specification of the AWS CMEK.
//
// swagger:model openapiAwsCmekSpec
type OpenapiAwsCmekSpec struct {

	// The KMS ARN of the AWS CMEK.
	// Example: arn:aws:kms:example
	// Required: true
	KmsArn *string `json:"kms_arn"`

	// The region name of the AWS CMEK. The region value should match the cloud provider's region code.
	//
	// You can get the complete list of available regions from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// For the detailed information on each region, refer to the documentation of the [AWS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html) cloud provider.
	// Example: us-west-2
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this openapi aws cmek spec
func (m *OpenapiAwsCmekSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKmsArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiAwsCmekSpec) validateKmsArn(formats strfmt.Registry) error {

	if err := validate.Required("kms_arn", "body", m.KmsArn); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiAwsCmekSpec) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi aws cmek spec based on context it is used
func (m *OpenapiAwsCmekSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiAwsCmekSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiAwsCmekSpec) UnmarshalBinary(b []byte) error {
	var res OpenapiAwsCmekSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

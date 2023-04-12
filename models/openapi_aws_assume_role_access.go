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

// OpenapiAwsAssumeRoleAccess AwsAssumeRoleAccess
//
// AwsAssumeRoleAccess represents the settings to access the Amazon S3 source by assuming to a specific AWS role.
//
// swagger:model openapiAwsAssumeRoleAccess
type OpenapiAwsAssumeRoleAccess struct {

	// The specific AWS role ARN that needs to be assumed to access the Amazon S3 data source.
	// Example: arn:aws:iam::999999999999:role/sample-role
	// Required: true
	AssumeRole *string `json:"assume_role"`
}

// Validate validates this openapi aws assume role access
func (m *OpenapiAwsAssumeRoleAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssumeRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiAwsAssumeRoleAccess) validateAssumeRole(formats strfmt.Registry) error {

	if err := validate.Required("assume_role", "body", m.AssumeRole); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi aws assume role access based on context it is used
func (m *OpenapiAwsAssumeRoleAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiAwsAssumeRoleAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiAwsAssumeRoleAccess) UnmarshalBinary(b []byte) error {
	var res OpenapiAwsAssumeRoleAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
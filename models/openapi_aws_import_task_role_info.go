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

// OpenapiAwsImportTaskRoleInfo AwsImportTaskRoleInfo
//
// AwsImportTaskRoleInfo is the role information for import tasks on an AWS cluster.
//
// swagger:model openapiAwsImportTaskRoleInfo
type OpenapiAwsImportTaskRoleInfo struct {

	// The account ID under which the import tasks for this cluster are running.
	// Example: 999999999999
	// Required: true
	AccountID *string `json:"account_id"`

	// The unique external ID that binds to the cluster, which is a long string. When an import task starts and attempts to assume a specified role, it automatically attaches this external ID. This means that you can configure this external ID in the assumed role's trust relationship, so that only the import task of that specified cluster can access the data by assuming the role. This can provide additional security.
	// Example: abcdefghijklmnopqrstuvwxyz0123456789xxxxxxxxxxxxxx
	// Required: true
	ExternalID *string `json:"external_id"`
}

// Validate validates this openapi aws import task role info
func (m *OpenapiAwsImportTaskRoleInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiAwsImportTaskRoleInfo) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiAwsImportTaskRoleInfo) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi aws import task role info based on context it is used
func (m *OpenapiAwsImportTaskRoleInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiAwsImportTaskRoleInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiAwsImportTaskRoleInfo) UnmarshalBinary(b []byte) error {
	var res OpenapiAwsImportTaskRoleInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

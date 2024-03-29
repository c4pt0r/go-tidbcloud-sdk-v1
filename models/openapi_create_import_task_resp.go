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

// OpenapiCreateImportTaskResp CreateImportTaskResp
//
// CreateImportTaskResp is the response to the creation of an import task.
//
// swagger:model openapiCreateImportTaskResp
type OpenapiCreateImportTaskResp struct {

	// The ID of the import task.
	// Example: 12345
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this openapi create import task resp
func (m *OpenapiCreateImportTaskResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiCreateImportTaskResp) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi create import task resp based on context it is used
func (m *OpenapiCreateImportTaskResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiCreateImportTaskResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiCreateImportTaskResp) UnmarshalBinary(b []byte) error {
	var res OpenapiCreateImportTaskResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

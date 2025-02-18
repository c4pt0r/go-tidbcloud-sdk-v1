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

// OpenapiAzureTokenAccess AzureTokenAccess
//
// AzureTokenAccess represents the settings to access the Azure blob source by using sas token.
//
// swagger:model openapiAzureTokenAccess
type OpenapiAzureTokenAccess struct {

	// The sas token to access the data. This information will be redacted when it is retrieved to obtain the import task information.
	// Example: YOUR_SAS_TOKEN
	// Required: true
	SasToken *string `json:"sas_token"`
}

// Validate validates this openapi azure token access
func (m *OpenapiAzureTokenAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSasToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiAzureTokenAccess) validateSasToken(formats strfmt.Registry) error {

	if err := validate.Required("sas_token", "body", m.SasToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi azure token access based on context it is used
func (m *OpenapiAzureTokenAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiAzureTokenAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiAzureTokenAccess) UnmarshalBinary(b []byte) error {
	var res OpenapiAzureTokenAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

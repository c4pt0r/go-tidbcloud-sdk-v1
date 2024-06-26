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

// OpenapiTiDBComponent openapi ti d b component
//
// swagger:model openapiTiDBComponent
type OpenapiTiDBComponent struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// Example: 2
	// Required: true
	NodeQuantity *int32 `json:"node_quantity"`

	// The size of the TiDB component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 4, then the cluster does not support TiFlash.
	// Example: 8C16G
	// Required: true
	NodeSize *string `json:"node_size"`
}

// Validate validates this openapi ti d b component
func (m *OpenapiTiDBComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiTiDBComponent) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiTiDBComponent) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi ti d b component based on context it is used
func (m *OpenapiTiDBComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiTiDBComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiTiDBComponent) UnmarshalBinary(b []byte) error {
	var res OpenapiTiDBComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

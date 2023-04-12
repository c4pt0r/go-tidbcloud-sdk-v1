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

// OpenapiTiKVComponent openapi ti k v component
//
// swagger:model openapiTiKVComponent
type OpenapiTiKVComponent struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - The `node_quantity` of TiKV must be a multiple of 3.
	// Example: 3
	// Required: true
	NodeQuantity *int32 `json:"node_quantity"`

	// The size of the TiKV component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then the cluster does not support TiFlash.
	// Example: 8C64G
	// Required: true
	NodeSize *string `json:"node_size"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 1024
	// Required: true
	StorageSizeGib *int32 `json:"storage_size_gib"`
}

// Validate validates this openapi ti k v component
func (m *OpenapiTiKVComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageSizeGib(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiTiKVComponent) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiTiKVComponent) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiTiKVComponent) validateStorageSizeGib(formats strfmt.Registry) error {

	if err := validate.Required("storage_size_gib", "body", m.StorageSizeGib); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi ti k v component based on context it is used
func (m *OpenapiTiKVComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiTiKVComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiTiKVComponent) UnmarshalBinary(b []byte) error {
	var res OpenapiTiKVComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

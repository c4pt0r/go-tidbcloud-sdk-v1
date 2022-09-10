// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiUpdateTiKVComponent openapi update ti k v component
//
// swagger:model openapiUpdateTiKVComponent
type OpenapiUpdateTiKVComponent struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - You cannot decrease node quantity for TiKV.
	// - The `node_quantity` of TiKV must be a multiple of 3.
	// Example: 6
	NodeQuantity int32 `json:"node_quantity,omitempty"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - You cannot decrease storage size for TiKV.
	// - If your TiDB cluster is hosted by AWS, after changing the storage size of TiKV, you must wait at least six hours before you can change it again.
	// Example: 1024
	StorageSizeGib int32 `json:"storage_size_gib,omitempty"`
}

// Validate validates this openapi update ti k v component
func (m *OpenapiUpdateTiKVComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi update ti k v component based on context it is used
func (m *OpenapiUpdateTiKVComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateTiKVComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateTiKVComponent) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateTiKVComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
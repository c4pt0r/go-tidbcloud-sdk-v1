// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiUpdateTiFlashComponent openapi update ti flash component
//
// swagger:model openapiUpdateTiFlashComponent
type OpenapiUpdateTiFlashComponent struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 2
	NodeQuantity *int32 `json:"node_quantity,omitempty"`

	// The size of the TiFlash component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 4, then the cluster does not support TiFlash.
	//
	// **Limitations**:
	// - See [Change node size](https://docs.pingcap.com/tidbcloud/scale-tidb-cluster#change-node-size).
	// Example: 16C128G
	NodeSize *string `json:"node_size,omitempty"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - You cannot decrease storage size for TiFlash.
	// - If your TiDB cluster is hosted by AWS, after changing the storage size of TiFlash, you must wait at least six hours before you can change it again.
	// Example: 2048
	StorageSizeGib *int32 `json:"storage_size_gib,omitempty"`
}

// Validate validates this openapi update ti flash component
func (m *OpenapiUpdateTiFlashComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi update ti flash component based on context it is used
func (m *OpenapiUpdateTiFlashComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateTiFlashComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateTiFlashComponent) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateTiFlashComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

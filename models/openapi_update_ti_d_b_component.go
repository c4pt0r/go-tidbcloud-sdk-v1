// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiUpdateTiDBComponent openapi update ti d b component
//
// swagger:model openapiUpdateTiDBComponent
type OpenapiUpdateTiDBComponent struct {

	// The ID of the accesspoint.
	// Example: 1
	AccesspointID *string `json:"accesspoint_id,omitempty"`

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 3
	NodeQuantity *int32 `json:"node_quantity,omitempty"`

	// The size of the TiDB component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 4, then the cluster does not support TiFlash.
	//
	// **Limitations**:
	// - See [Change node size](https://docs.pingcap.com/tidbcloud/scale-tidb-cluster#change-node-size).
	// Example: 16C32G
	NodeSize *string `json:"node_size,omitempty"`
}

// Validate validates this openapi update ti d b component
func (m *OpenapiUpdateTiDBComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi update ti d b component based on context it is used
func (m *OpenapiUpdateTiDBComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateTiDBComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateTiDBComponent) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateTiDBComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

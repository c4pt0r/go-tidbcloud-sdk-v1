// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiUpdateClusterConfig UpdateClusterComponents
//
// UpdateClusterComponents is the request for updating cluster components.
//
// swagger:model openapiUpdateClusterConfig
type OpenapiUpdateClusterConfig struct {

	// components
	Components *OpenapiUpdateClusterConfigComponents `json:"components,omitempty"`

	// Flag that indicates whether the cluster is paused. `true` means to pause the cluster, and `false` means to resume the cluster. For more details, refer to [Pause or Resume a TiDB Cluster](https://docs.pingcap.com/tidbcloud/pause-or-resume-tidb-cluster).
	//
	// **Limitations:**
	//  - The cluster can be paused only when the `cluster_status` is `"AVAILABLE"`.
	// - The cluster can be resumed only when the `cluster_status` is `"PAUSED"`.
	Paused bool `json:"paused,omitempty"`
}

// Validate validates this openapi update cluster config
func (m *OpenapiUpdateClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiUpdateClusterConfig) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	if m.Components != nil {
		if err := m.Components.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi update cluster config based on the context it is used
func (m *OpenapiUpdateClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiUpdateClusterConfig) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.Components != nil {
		if err := m.Components.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfig) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiUpdateClusterConfigComponents The components of the cluster.
//
// swagger:model OpenapiUpdateClusterConfigComponents
type OpenapiUpdateClusterConfigComponents struct {

	// tidb
	Tidb *OpenapiUpdateClusterConfigComponentsTidb `json:"tidb,omitempty"`

	// tiflash
	Tiflash *OpenapiUpdateClusterConfigComponentsTiflash `json:"tiflash,omitempty"`

	// tikv
	Tikv *OpenapiUpdateClusterConfigComponentsTikv `json:"tikv,omitempty"`
}

// Validate validates this openapi update cluster config components
func (m *OpenapiUpdateClusterConfigComponents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTidb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTiflash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTikv(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiUpdateClusterConfigComponents) validateTidb(formats strfmt.Registry) error {
	if swag.IsZero(m.Tidb) { // not required
		return nil
	}

	if m.Tidb != nil {
		if err := m.Tidb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components" + "." + "tidb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components" + "." + "tidb")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiUpdateClusterConfigComponents) validateTiflash(formats strfmt.Registry) error {
	if swag.IsZero(m.Tiflash) { // not required
		return nil
	}

	if m.Tiflash != nil {
		if err := m.Tiflash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components" + "." + "tiflash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components" + "." + "tiflash")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiUpdateClusterConfigComponents) validateTikv(formats strfmt.Registry) error {
	if swag.IsZero(m.Tikv) { // not required
		return nil
	}

	if m.Tikv != nil {
		if err := m.Tikv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components" + "." + "tikv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components" + "." + "tikv")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi update cluster config components based on the context it is used
func (m *OpenapiUpdateClusterConfigComponents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTidb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTiflash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTikv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiUpdateClusterConfigComponents) contextValidateTidb(ctx context.Context, formats strfmt.Registry) error {

	if m.Tidb != nil {
		if err := m.Tidb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components" + "." + "tidb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components" + "." + "tidb")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiUpdateClusterConfigComponents) contextValidateTiflash(ctx context.Context, formats strfmt.Registry) error {

	if m.Tiflash != nil {
		if err := m.Tiflash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components" + "." + "tiflash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components" + "." + "tiflash")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiUpdateClusterConfigComponents) contextValidateTikv(ctx context.Context, formats strfmt.Registry) error {

	if m.Tikv != nil {
		if err := m.Tikv.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components" + "." + "tikv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components" + "." + "tikv")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponents) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateClusterConfigComponents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiUpdateClusterConfigComponentsTidb The TiDB component of the cluster.
//
// swagger:model OpenapiUpdateClusterConfigComponentsTidb
type OpenapiUpdateClusterConfigComponentsTidb struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 3
	NodeQuantity int32 `json:"node_quantity,omitempty"`

	// The size of the TiDB component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then the cluster does not support TiFlash.
	//
	// **Limitations**:
	// - See [Change node size](https://docs.pingcap.com/tidbcloud/scale-tidb-cluster#change-node-size).
	// Example: 16C32G
	NodeSize string `json:"node_size,omitempty"`
}

// Validate validates this openapi update cluster config components tidb
func (m *OpenapiUpdateClusterConfigComponentsTidb) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi update cluster config components tidb based on context it is used
func (m *OpenapiUpdateClusterConfigComponentsTidb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponentsTidb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponentsTidb) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateClusterConfigComponentsTidb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiUpdateClusterConfigComponentsTiflash The TiFlash component of the cluster.
//
// If you want to add TiFlash nodes to a cluster that does not have one before (increase the node_quantity of `"TIFLASH"` from 0), you must specify the `node_size`, `storage_size_gib` and `node_quantity` of TiFlash nodes.
//
// swagger:model OpenapiUpdateClusterConfigComponentsTiflash
type OpenapiUpdateClusterConfigComponentsTiflash struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 2
	NodeQuantity int32 `json:"node_quantity,omitempty"`

	// The size of the TiFlash component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then the cluster does not support TiFlash.
	//
	// **Limitations**:
	// - See [Change node size](https://docs.pingcap.com/tidbcloud/scale-tidb-cluster#change-node-size).
	// Example: 16C128G
	NodeSize string `json:"node_size,omitempty"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - You cannot decrease storage size for TiFlash.
	// - If your TiDB cluster is hosted by AWS, after changing the storage size of TiFlash, you must wait at least six hours before you can change it again.
	// Example: 2048
	StorageSizeGib int32 `json:"storage_size_gib,omitempty"`
}

// Validate validates this openapi update cluster config components tiflash
func (m *OpenapiUpdateClusterConfigComponentsTiflash) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi update cluster config components tiflash based on context it is used
func (m *OpenapiUpdateClusterConfigComponentsTiflash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponentsTiflash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponentsTiflash) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateClusterConfigComponentsTiflash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiUpdateClusterConfigComponentsTikv The TiKV component of the cluster.
//
// swagger:model OpenapiUpdateClusterConfigComponentsTikv
type OpenapiUpdateClusterConfigComponentsTikv struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - The `node_quantity` of TiKV must be a multiple of 3.
	// Example: 6
	NodeQuantity int32 `json:"node_quantity,omitempty"`

	// The size of the TiKV component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 2 or 4, then the cluster does not support TiFlash.
	//
	// **Limitations**:
	// - See [Change node size](https://docs.pingcap.com/tidbcloud/scale-tidb-cluster#change-node-size).
	// Example: 16C64G
	NodeSize string `json:"node_size,omitempty"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Limitations**:
	// - You cannot decrease storage size for TiKV.
	// - If your TiDB cluster is hosted by AWS, after changing the storage size of TiKV, you must wait at least six hours before you can change it again.
	// Example: 2048
	StorageSizeGib int32 `json:"storage_size_gib,omitempty"`
}

// Validate validates this openapi update cluster config components tikv
func (m *OpenapiUpdateClusterConfigComponentsTikv) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi update cluster config components tikv based on context it is used
func (m *OpenapiUpdateClusterConfigComponentsTikv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponentsTikv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiUpdateClusterConfigComponentsTikv) UnmarshalBinary(b []byte) error {
	var res OpenapiUpdateClusterConfigComponentsTikv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

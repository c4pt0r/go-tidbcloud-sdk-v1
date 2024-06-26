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

// OpenapiClusterComponents openapi cluster components
//
// swagger:model openapiClusterComponents
type OpenapiClusterComponents struct {

	// tidb
	// Required: true
	Tidb *OpenapiClusterComponentsTidb `json:"tidb"`

	// tiflash
	Tiflash *OpenapiClusterComponentsTiflash `json:"tiflash,omitempty"`

	// tikv
	// Required: true
	Tikv *OpenapiClusterComponentsTikv `json:"tikv"`
}

// Validate validates this openapi cluster components
func (m *OpenapiClusterComponents) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterComponents) validateTidb(formats strfmt.Registry) error {

	if err := validate.Required("tidb", "body", m.Tidb); err != nil {
		return err
	}

	if m.Tidb != nil {
		if err := m.Tidb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tidb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tidb")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiClusterComponents) validateTiflash(formats strfmt.Registry) error {
	if swag.IsZero(m.Tiflash) { // not required
		return nil
	}

	if m.Tiflash != nil {
		if err := m.Tiflash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tiflash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tiflash")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiClusterComponents) validateTikv(formats strfmt.Registry) error {

	if err := validate.Required("tikv", "body", m.Tikv); err != nil {
		return err
	}

	if m.Tikv != nil {
		if err := m.Tikv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tikv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tikv")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi cluster components based on the context it is used
func (m *OpenapiClusterComponents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OpenapiClusterComponents) contextValidateTidb(ctx context.Context, formats strfmt.Registry) error {

	if m.Tidb != nil {

		if err := m.Tidb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tidb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tidb")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiClusterComponents) contextValidateTiflash(ctx context.Context, formats strfmt.Registry) error {

	if m.Tiflash != nil {

		if swag.IsZero(m.Tiflash) { // not required
			return nil
		}

		if err := m.Tiflash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tiflash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tiflash")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiClusterComponents) contextValidateTikv(ctx context.Context, formats strfmt.Registry) error {

	if m.Tikv != nil {

		if err := m.Tikv.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tikv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tikv")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterComponents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterComponents) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterComponents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterComponentsTidb The TiDB component of the cluster.
//
// swagger:model OpenapiClusterComponentsTidb
type OpenapiClusterComponentsTidb struct {

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

// Validate validates this openapi cluster components tidb
func (m *OpenapiClusterComponentsTidb) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterComponentsTidb) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("tidb"+"."+"node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterComponentsTidb) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("tidb"+"."+"node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster components tidb based on context it is used
func (m *OpenapiClusterComponentsTidb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterComponentsTidb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterComponentsTidb) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterComponentsTidb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterComponentsTiflash The TiFlash component of the cluster.
//
// swagger:model OpenapiClusterComponentsTiflash
type OpenapiClusterComponentsTiflash struct {

	// The number of nodes in the cluster. You can get the minimum and step of a node quantity from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 1
	// Required: true
	NodeQuantity *int32 `json:"node_quantity"`

	// The size of the TiFlash component in the cluster. You can get the available node size of each region from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// **Additional combination rules**:
	// - If the vCPUs of TiDB or TiKV component is 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 4, then the cluster does not support TiFlash.
	// Example: 8C64G
	// Required: true
	NodeSize *string `json:"node_size"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// Example: 1024
	// Required: true
	StorageSizeGib *int32 `json:"storage_size_gib"`
}

// Validate validates this openapi cluster components tiflash
func (m *OpenapiClusterComponentsTiflash) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterComponentsTiflash) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("tiflash"+"."+"node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterComponentsTiflash) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("tiflash"+"."+"node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterComponentsTiflash) validateStorageSizeGib(formats strfmt.Registry) error {

	if err := validate.Required("tiflash"+"."+"storage_size_gib", "body", m.StorageSizeGib); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster components tiflash based on context it is used
func (m *OpenapiClusterComponentsTiflash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterComponentsTiflash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterComponentsTiflash) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterComponentsTiflash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterComponentsTikv The TiKV component of the cluster.
//
// swagger:model OpenapiClusterComponentsTikv
type OpenapiClusterComponentsTikv struct {

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
	// - If the vCPUs of TiDB or TiKV component is 4, then their vCPUs need to be the same.
	// - If the vCPUs of TiDB or TiKV component is 4, then the cluster does not support TiFlash.
	// Example: 8C64G
	// Required: true
	NodeSize *string `json:"node_size"`

	// The storage size of a node in the cluster. You can get the minimum and maximum of storage size from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	// Example: 1024
	// Required: true
	StorageSizeGib *int32 `json:"storage_size_gib"`
}

// Validate validates this openapi cluster components tikv
func (m *OpenapiClusterComponentsTikv) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterComponentsTikv) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("tikv"+"."+"node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterComponentsTikv) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("tikv"+"."+"node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterComponentsTikv) validateStorageSizeGib(formats strfmt.Registry) error {

	if err := validate.Required("tikv"+"."+"storage_size_gib", "body", m.StorageSizeGib); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster components tikv based on context it is used
func (m *OpenapiClusterComponentsTikv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterComponentsTikv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterComponentsTikv) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterComponentsTikv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

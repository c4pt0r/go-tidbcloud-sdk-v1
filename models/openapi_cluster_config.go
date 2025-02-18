// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiClusterConfig openapi cluster config
//
// swagger:model openapiClusterConfig
type OpenapiClusterConfig struct {

	// components
	Components *OpenapiClusterConfigComponents `json:"components,omitempty"`

	// A list of IP addresses and Classless Inter-Domain Routing (CIDR) addresses that are allowed to access the TiDB Cloud cluster via [standard connection](https://docs.pingcap.com/tidbcloud/connect-to-tidb-cluster#connect-via-standard-connection).
	IPAccessList []*OpenapiClusterConfigIPAccessListItems0 `json:"ip_access_list"`

	// The TiDB port for connection. The port must be in the range of 1024-65535 except 10080.
	//
	// **Limitations**:
	// - For a TiDB Cloud Serverless cluster, only port `4000` is available.
	// Example: 4000
	// Maximum: 65535
	// Minimum: 1024
	Port int32 `json:"port,omitempty"`

	// The root password to access the cluster. It must be 8-64 characters.
	// Example: password_example
	// Required: true
	// Max Length: 64
	// Min Length: 8
	RootPassword *string `json:"root_password"`
}

// Validate validates this openapi cluster config
func (m *OpenapiClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAccessList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiClusterConfig) validateComponents(formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfig) validateIPAccessList(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAccessList) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAccessList); i++ {
		if swag.IsZero(m.IPAccessList[i]) { // not required
			continue
		}

		if m.IPAccessList[i] != nil {
			if err := m.IPAccessList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_access_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_access_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiClusterConfig) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterConfig) validateRootPassword(formats strfmt.Registry) error {

	if err := validate.Required("root_password", "body", m.RootPassword); err != nil {
		return err
	}

	if err := validate.MinLength("root_password", "body", *m.RootPassword, 8); err != nil {
		return err
	}

	if err := validate.MaxLength("root_password", "body", *m.RootPassword, 64); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi cluster config based on the context it is used
func (m *OpenapiClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAccessList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiClusterConfig) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.Components != nil {

		if swag.IsZero(m.Components) { // not required
			return nil
		}

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

func (m *OpenapiClusterConfig) contextValidateIPAccessList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAccessList); i++ {

		if m.IPAccessList[i] != nil {

			if swag.IsZero(m.IPAccessList[i]) { // not required
				return nil
			}

			if err := m.IPAccessList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_access_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_access_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterConfig) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterConfigComponents The components of the cluster.
//
// **Limitations**:
// - For a TiDB Cloud Dedicated cluster, the `components` parameter is **required**.
// - For a TiDB Cloud Serverless cluster, the `components` value is **ignored**. Setting this configuration does not have any effects.
// Example: {"tidb":{"node_quantity":2,"node_size":"8C16G"},"tikv":{"node_quantity":3,"node_size":"8C32G","storage_size_gib":1024}}
//
// swagger:model OpenapiClusterConfigComponents
type OpenapiClusterConfigComponents struct {

	// tidb
	// Required: true
	Tidb *OpenapiClusterConfigComponentsTidb `json:"tidb"`

	// tiflash
	Tiflash *OpenapiClusterConfigComponentsTiflash `json:"tiflash,omitempty"`

	// tikv
	// Required: true
	Tikv *OpenapiClusterConfigComponentsTikv `json:"tikv"`
}

// Validate validates this openapi cluster config components
func (m *OpenapiClusterConfigComponents) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfigComponents) validateTidb(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tidb", "body", m.Tidb); err != nil {
		return err
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

func (m *OpenapiClusterConfigComponents) validateTiflash(formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfigComponents) validateTikv(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tikv", "body", m.Tikv); err != nil {
		return err
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

// ContextValidate validate this openapi cluster config components based on the context it is used
func (m *OpenapiClusterConfigComponents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfigComponents) contextValidateTidb(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OpenapiClusterConfigComponents) contextValidateTiflash(ctx context.Context, formats strfmt.Registry) error {

	if m.Tiflash != nil {

		if swag.IsZero(m.Tiflash) { // not required
			return nil
		}

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

func (m *OpenapiClusterConfigComponents) contextValidateTikv(ctx context.Context, formats strfmt.Registry) error {

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
func (m *OpenapiClusterConfigComponents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterConfigComponents) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterConfigComponents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterConfigComponentsTidb The TiDB component of the cluster.
//
// swagger:model OpenapiClusterConfigComponentsTidb
type OpenapiClusterConfigComponentsTidb struct {

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

// Validate validates this openapi cluster config components tidb
func (m *OpenapiClusterConfigComponentsTidb) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfigComponentsTidb) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tidb"+"."+"node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterConfigComponentsTidb) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tidb"+"."+"node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster config components tidb based on context it is used
func (m *OpenapiClusterConfigComponentsTidb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterConfigComponentsTidb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterConfigComponentsTidb) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterConfigComponentsTidb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterConfigComponentsTiflash The TiFlash component of the cluster.
//
// swagger:model OpenapiClusterConfigComponentsTiflash
type OpenapiClusterConfigComponentsTiflash struct {

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

// Validate validates this openapi cluster config components tiflash
func (m *OpenapiClusterConfigComponentsTiflash) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfigComponentsTiflash) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tiflash"+"."+"node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterConfigComponentsTiflash) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tiflash"+"."+"node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterConfigComponentsTiflash) validateStorageSizeGib(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tiflash"+"."+"storage_size_gib", "body", m.StorageSizeGib); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster config components tiflash based on context it is used
func (m *OpenapiClusterConfigComponentsTiflash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterConfigComponentsTiflash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterConfigComponentsTiflash) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterConfigComponentsTiflash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterConfigComponentsTikv The TiKV component of the cluster.
//
// swagger:model OpenapiClusterConfigComponentsTikv
type OpenapiClusterConfigComponentsTikv struct {

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

// Validate validates this openapi cluster config components tikv
func (m *OpenapiClusterConfigComponentsTikv) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterConfigComponentsTikv) validateNodeQuantity(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tikv"+"."+"node_quantity", "body", m.NodeQuantity); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterConfigComponentsTikv) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tikv"+"."+"node_size", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiClusterConfigComponentsTikv) validateStorageSizeGib(formats strfmt.Registry) error {

	if err := validate.Required("components"+"."+"tikv"+"."+"storage_size_gib", "body", m.StorageSizeGib); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster config components tikv based on context it is used
func (m *OpenapiClusterConfigComponentsTikv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterConfigComponentsTikv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterConfigComponentsTikv) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterConfigComponentsTikv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterConfigIPAccessListItems0 openapi cluster config IP access list items0
//
// swagger:model OpenapiClusterConfigIPAccessListItems0
type OpenapiClusterConfigIPAccessListItems0 struct {

	// The IP address or CIDR range that you want to add to the cluster's IP access list.
	// Example: 8.8.8.8/32
	// Required: true
	Cidr *string `json:"cidr"`

	// Description that explains the purpose of the entry.
	// Example: My Current IP Address
	Description string `json:"description,omitempty"`
}

// Validate validates this openapi cluster config IP access list items0
func (m *OpenapiClusterConfigIPAccessListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiClusterConfigIPAccessListItems0) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster config IP access list items0 based on context it is used
func (m *OpenapiClusterConfigIPAccessListItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterConfigIPAccessListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterConfigIPAccessListItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterConfigIPAccessListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

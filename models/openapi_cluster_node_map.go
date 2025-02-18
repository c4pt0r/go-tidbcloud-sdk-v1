// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiClusterNodeMap openapi cluster node map
//
// swagger:model openapiClusterNodeMap
type OpenapiClusterNodeMap struct {

	// TiDB node map.
	// Example: [{"availability_zone":"us-west-2a","node_name":"tidb-0","node_size":"8C16G","ram_bytes":"17179869184","status":"NODE_STATUS_AVAILABLE","vcpu_num":8},{"availability_zone":"us-west-2b","node_name":"tidb-1","node_size":"8C16G","ram_bytes":"17179869184","status":"NODE_STATUS_AVAILABLE","vcpu_num":8}]
	// Required: true
	Tidb []*OpenapiClusterNodeMapTidbItems0 `json:"tidb"`

	// TiFlash node map.
	// Example: [{"availability_zone":"us-west-2a","node_name":"tiflash-0","node_size":"8C64G","ram_bytes":"68719476736","status":"NODE_STATUS_AVAILABLE","storage_size_gib":1024,"vcpu_num":8},{"availability_zone":"us-west-2b","node_name":"tiflash-1","node_size":"8C64G","ram_bytes":"68719476736","status":"NODE_STATUS_AVAILABLE","storage_size_gib":1024,"vcpu_num":8}]
	Tiflash []*OpenapiClusterNodeMapTiflashItems0 `json:"tiflash"`

	// TiKV node map.
	// Example: [{"availability_zone":"us-west-2a","node_name":"tikv-0","node_size":"8C32G","ram_bytes":"68719476736","status":"NODE_STATUS_AVAILABLE","storage_size_gib":1024,"vcpu_num":8},{"availability_zone":"us-west-2b","node_name":"tikv-1","node_size":"8C64G","ram_bytes":"68719476736","status":"NODE_STATUS_AVAILABLE","storage_size_gib":1024,"vcpu_num":8},{"availability_zone":"us-west-2c","node_name":"tikv-2","node_size":"8C64G","ram_bytes":"68719476736","status":"NODE_STATUS_AVAILABLE","storage_size_gib":1024,"vcpu_num":8}]
	// Required: true
	Tikv []*OpenapiClusterNodeMapTikvItems0 `json:"tikv"`
}

// Validate validates this openapi cluster node map
func (m *OpenapiClusterNodeMap) Validate(formats strfmt.Registry) error {
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

func (m *OpenapiClusterNodeMap) validateTidb(formats strfmt.Registry) error {

	if err := validate.Required("tidb", "body", m.Tidb); err != nil {
		return err
	}

	for i := 0; i < len(m.Tidb); i++ {
		if swag.IsZero(m.Tidb[i]) { // not required
			continue
		}

		if m.Tidb[i] != nil {
			if err := m.Tidb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tidb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tidb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiClusterNodeMap) validateTiflash(formats strfmt.Registry) error {
	if swag.IsZero(m.Tiflash) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiflash); i++ {
		if swag.IsZero(m.Tiflash[i]) { // not required
			continue
		}

		if m.Tiflash[i] != nil {
			if err := m.Tiflash[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiflash" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiflash" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiClusterNodeMap) validateTikv(formats strfmt.Registry) error {

	if err := validate.Required("tikv", "body", m.Tikv); err != nil {
		return err
	}

	for i := 0; i < len(m.Tikv); i++ {
		if swag.IsZero(m.Tikv[i]) { // not required
			continue
		}

		if m.Tikv[i] != nil {
			if err := m.Tikv[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tikv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tikv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this openapi cluster node map based on the context it is used
func (m *OpenapiClusterNodeMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OpenapiClusterNodeMap) contextValidateTidb(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tidb); i++ {

		if m.Tidb[i] != nil {

			if swag.IsZero(m.Tidb[i]) { // not required
				return nil
			}

			if err := m.Tidb[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tidb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tidb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiClusterNodeMap) contextValidateTiflash(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tiflash); i++ {

		if m.Tiflash[i] != nil {

			if swag.IsZero(m.Tiflash[i]) { // not required
				return nil
			}

			if err := m.Tiflash[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiflash" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiflash" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiClusterNodeMap) contextValidateTikv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tikv); i++ {

		if m.Tikv[i] != nil {

			if swag.IsZero(m.Tikv[i]) { // not required
				return nil
			}

			if err := m.Tikv[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tikv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tikv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterNodeMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterNodeMap) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterNodeMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterNodeMapTidbItems0 openapi cluster node map tidb items0
//
// swagger:model OpenapiClusterNodeMapTidbItems0
type OpenapiClusterNodeMapTidbItems0 struct {

	// The availability zone of a node in the cluster.
	// Example: us-west-2a
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// The name of a node in the cluster.
	// Example: tidb-0
	NodeName string `json:"node_name,omitempty"`

	// The size of the TiDB component in the cluster.
	// Example: 8C16G
	NodeSize string `json:"node_size,omitempty"`

	// The RAM size of a node in the cluster. If the `cluster_type` is `"DEVELOPER"`, `ram_bytes` is always 0.
	// Example: 17179869184
	RAMBytes string `json:"ram_bytes,omitempty"`

	// The status of a node in the cluster.
	// Example: NODE_STATUS_AVAILABLE
	// Enum: ["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]
	Status string `json:"status,omitempty"`

	// The total vCPUs of a node in the cluster. If the `cluster_type` is `"DEVELOPER"`, `vcpu_num` is always 0.
	// Example: 8
	VcpuNum int32 `json:"vcpu_num,omitempty"`
}

// Validate validates this openapi cluster node map tidb items0
func (m *OpenapiClusterNodeMapTidbItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiClusterNodeMapTidbItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiClusterNodeMapTidbItems0TypeStatusPropEnum = append(openapiClusterNodeMapTidbItems0TypeStatusPropEnum, v)
	}
}

const (

	// OpenapiClusterNodeMapTidbItems0StatusNODESTATUSAVAILABLE captures enum value "NODE_STATUS_AVAILABLE"
	OpenapiClusterNodeMapTidbItems0StatusNODESTATUSAVAILABLE string = "NODE_STATUS_AVAILABLE"

	// OpenapiClusterNodeMapTidbItems0StatusNODESTATUSUNAVAILABLE captures enum value "NODE_STATUS_UNAVAILABLE"
	OpenapiClusterNodeMapTidbItems0StatusNODESTATUSUNAVAILABLE string = "NODE_STATUS_UNAVAILABLE"

	// OpenapiClusterNodeMapTidbItems0StatusNODESTATUSCREATING captures enum value "NODE_STATUS_CREATING"
	OpenapiClusterNodeMapTidbItems0StatusNODESTATUSCREATING string = "NODE_STATUS_CREATING"

	// OpenapiClusterNodeMapTidbItems0StatusNODESTATUSDELETING captures enum value "NODE_STATUS_DELETING"
	OpenapiClusterNodeMapTidbItems0StatusNODESTATUSDELETING string = "NODE_STATUS_DELETING"
)

// prop value enum
func (m *OpenapiClusterNodeMapTidbItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiClusterNodeMapTidbItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiClusterNodeMapTidbItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster node map tidb items0 based on context it is used
func (m *OpenapiClusterNodeMapTidbItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterNodeMapTidbItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterNodeMapTidbItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterNodeMapTidbItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterNodeMapTiflashItems0 openapi cluster node map tiflash items0
//
// swagger:model OpenapiClusterNodeMapTiflashItems0
type OpenapiClusterNodeMapTiflashItems0 struct {

	// The availability zone of a node in the cluster.
	// Example: us-west-2a
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// The name of a node in the cluster.
	// Example: tiflash-0
	NodeName string `json:"node_name,omitempty"`

	// The size of the TiFlash component in the cluster.
	// Example: 8C64G
	NodeSize string `json:"node_size,omitempty"`

	// The RAM size of a node in the cluster. If the `cluster_type` is `"DEVELOPER"`, `ram_bytes` is always 0.
	// Example: 68719476736
	RAMBytes string `json:"ram_bytes,omitempty"`

	// The status of a node in the cluster.
	// Example: NODE_STATUS_AVAILABLE
	// Enum: ["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]
	Status string `json:"status,omitempty"`

	// The storage size of a node in the cluster.
	// Example: 1024
	StorageSizeGib int32 `json:"storage_size_gib,omitempty"`

	// The total vCPUs of a node in the cluster. If the `cluster_type` is `"DEVELOPER"`, `vcpu_num` is always 0.
	// Example: 8
	VcpuNum int32 `json:"vcpu_num,omitempty"`
}

// Validate validates this openapi cluster node map tiflash items0
func (m *OpenapiClusterNodeMapTiflashItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiClusterNodeMapTiflashItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiClusterNodeMapTiflashItems0TypeStatusPropEnum = append(openapiClusterNodeMapTiflashItems0TypeStatusPropEnum, v)
	}
}

const (

	// OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSAVAILABLE captures enum value "NODE_STATUS_AVAILABLE"
	OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSAVAILABLE string = "NODE_STATUS_AVAILABLE"

	// OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSUNAVAILABLE captures enum value "NODE_STATUS_UNAVAILABLE"
	OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSUNAVAILABLE string = "NODE_STATUS_UNAVAILABLE"

	// OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSCREATING captures enum value "NODE_STATUS_CREATING"
	OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSCREATING string = "NODE_STATUS_CREATING"

	// OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSDELETING captures enum value "NODE_STATUS_DELETING"
	OpenapiClusterNodeMapTiflashItems0StatusNODESTATUSDELETING string = "NODE_STATUS_DELETING"
)

// prop value enum
func (m *OpenapiClusterNodeMapTiflashItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiClusterNodeMapTiflashItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiClusterNodeMapTiflashItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster node map tiflash items0 based on context it is used
func (m *OpenapiClusterNodeMapTiflashItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterNodeMapTiflashItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterNodeMapTiflashItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterNodeMapTiflashItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiClusterNodeMapTikvItems0 openapi cluster node map tikv items0
//
// swagger:model OpenapiClusterNodeMapTikvItems0
type OpenapiClusterNodeMapTikvItems0 struct {

	// The availability zone of a node in the cluster.
	// Example: us-west-2a
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// The name of a node in the cluster.
	// Example: tikv-0
	NodeName string `json:"node_name,omitempty"`

	// The size of the TiKV component in the cluster.
	// Example: 8C32G
	NodeSize string `json:"node_size,omitempty"`

	// The RAM size of a node in the cluster. If the `cluster_type` is `"DEVELOPER"`, `ram_bytes` is always 0.
	// Example: 68719476736
	RAMBytes string `json:"ram_bytes,omitempty"`

	// The status of a node in the cluster.
	// Example: NODE_STATUS_AVAILABLE
	// Enum: ["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]
	Status string `json:"status,omitempty"`

	// The storage size of a node in the cluster.
	// Example: 1024
	StorageSizeGib int32 `json:"storage_size_gib,omitempty"`

	// The total vCPUs of a node in the cluster. If the `cluster_type` is `"DEVELOPER"`, `vcpu_num` is always 0.
	// Example: 8
	VcpuNum int32 `json:"vcpu_num,omitempty"`
}

// Validate validates this openapi cluster node map tikv items0
func (m *OpenapiClusterNodeMapTikvItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiClusterNodeMapTikvItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiClusterNodeMapTikvItems0TypeStatusPropEnum = append(openapiClusterNodeMapTikvItems0TypeStatusPropEnum, v)
	}
}

const (

	// OpenapiClusterNodeMapTikvItems0StatusNODESTATUSAVAILABLE captures enum value "NODE_STATUS_AVAILABLE"
	OpenapiClusterNodeMapTikvItems0StatusNODESTATUSAVAILABLE string = "NODE_STATUS_AVAILABLE"

	// OpenapiClusterNodeMapTikvItems0StatusNODESTATUSUNAVAILABLE captures enum value "NODE_STATUS_UNAVAILABLE"
	OpenapiClusterNodeMapTikvItems0StatusNODESTATUSUNAVAILABLE string = "NODE_STATUS_UNAVAILABLE"

	// OpenapiClusterNodeMapTikvItems0StatusNODESTATUSCREATING captures enum value "NODE_STATUS_CREATING"
	OpenapiClusterNodeMapTikvItems0StatusNODESTATUSCREATING string = "NODE_STATUS_CREATING"

	// OpenapiClusterNodeMapTikvItems0StatusNODESTATUSDELETING captures enum value "NODE_STATUS_DELETING"
	OpenapiClusterNodeMapTikvItems0StatusNODESTATUSDELETING string = "NODE_STATUS_DELETING"
)

// prop value enum
func (m *OpenapiClusterNodeMapTikvItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiClusterNodeMapTikvItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiClusterNodeMapTikvItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi cluster node map tikv items0 based on context it is used
func (m *OpenapiClusterNodeMapTikvItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiClusterNodeMapTikvItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiClusterNodeMapTikvItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiClusterNodeMapTikvItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

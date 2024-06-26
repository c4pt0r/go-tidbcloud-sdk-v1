// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OpenapiPrivateEndpointServiceStatus openapi private endpoint service status
//
// swagger:model openapiPrivateEndpointServiceStatus
type OpenapiPrivateEndpointServiceStatus string

func NewOpenapiPrivateEndpointServiceStatus(value OpenapiPrivateEndpointServiceStatus) *OpenapiPrivateEndpointServiceStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OpenapiPrivateEndpointServiceStatus.
func (m OpenapiPrivateEndpointServiceStatus) Pointer() *OpenapiPrivateEndpointServiceStatus {
	return &m
}

const (

	// OpenapiPrivateEndpointServiceStatusCREATING captures enum value "CREATING"
	OpenapiPrivateEndpointServiceStatusCREATING OpenapiPrivateEndpointServiceStatus = "CREATING"

	// OpenapiPrivateEndpointServiceStatusACTIVE captures enum value "ACTIVE"
	OpenapiPrivateEndpointServiceStatusACTIVE OpenapiPrivateEndpointServiceStatus = "ACTIVE"

	// OpenapiPrivateEndpointServiceStatusDELETING captures enum value "DELETING"
	OpenapiPrivateEndpointServiceStatusDELETING OpenapiPrivateEndpointServiceStatus = "DELETING"
)

// for schema
var openapiPrivateEndpointServiceStatusEnum []interface{}

func init() {
	var res []OpenapiPrivateEndpointServiceStatus
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiPrivateEndpointServiceStatusEnum = append(openapiPrivateEndpointServiceStatusEnum, v)
	}
}

func (m OpenapiPrivateEndpointServiceStatus) validateOpenapiPrivateEndpointServiceStatusEnum(path, location string, value OpenapiPrivateEndpointServiceStatus) error {
	if err := validate.EnumCase(path, location, value, openapiPrivateEndpointServiceStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openapi private endpoint service status
func (m OpenapiPrivateEndpointServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenapiPrivateEndpointServiceStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openapi private endpoint service status based on context it is used
func (m OpenapiPrivateEndpointServiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

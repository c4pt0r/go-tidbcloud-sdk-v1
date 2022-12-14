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

// OpenapiGetRestoreRespStatusEnum openapi get restore resp status enum
//
// swagger:model openapiGetRestoreRespStatusEnum
type OpenapiGetRestoreRespStatusEnum string

func NewOpenapiGetRestoreRespStatusEnum(value OpenapiGetRestoreRespStatusEnum) *OpenapiGetRestoreRespStatusEnum {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OpenapiGetRestoreRespStatusEnum.
func (m OpenapiGetRestoreRespStatusEnum) Pointer() *OpenapiGetRestoreRespStatusEnum {
	return &m
}

const (

	// OpenapiGetRestoreRespStatusEnumPENDING captures enum value "PENDING"
	OpenapiGetRestoreRespStatusEnumPENDING OpenapiGetRestoreRespStatusEnum = "PENDING"

	// OpenapiGetRestoreRespStatusEnumRUNNING captures enum value "RUNNING"
	OpenapiGetRestoreRespStatusEnumRUNNING OpenapiGetRestoreRespStatusEnum = "RUNNING"

	// OpenapiGetRestoreRespStatusEnumFAILED captures enum value "FAILED"
	OpenapiGetRestoreRespStatusEnumFAILED OpenapiGetRestoreRespStatusEnum = "FAILED"

	// OpenapiGetRestoreRespStatusEnumSUCCESS captures enum value "SUCCESS"
	OpenapiGetRestoreRespStatusEnumSUCCESS OpenapiGetRestoreRespStatusEnum = "SUCCESS"
)

// for schema
var openapiGetRestoreRespStatusEnumEnum []interface{}

func init() {
	var res []OpenapiGetRestoreRespStatusEnum
	if err := json.Unmarshal([]byte(`["PENDING","RUNNING","FAILED","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiGetRestoreRespStatusEnumEnum = append(openapiGetRestoreRespStatusEnumEnum, v)
	}
}

func (m OpenapiGetRestoreRespStatusEnum) validateOpenapiGetRestoreRespStatusEnumEnum(path, location string, value OpenapiGetRestoreRespStatusEnum) error {
	if err := validate.EnumCase(path, location, value, openapiGetRestoreRespStatusEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openapi get restore resp status enum
func (m OpenapiGetRestoreRespStatusEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenapiGetRestoreRespStatusEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openapi get restore resp status enum based on context it is used
func (m OpenapiGetRestoreRespStatusEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

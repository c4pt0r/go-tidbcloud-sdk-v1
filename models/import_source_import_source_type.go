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

// ImportSourceImportSourceType import source import source type
//
// swagger:model ImportSourceImportSourceType
type ImportSourceImportSourceType string

func NewImportSourceImportSourceType(value ImportSourceImportSourceType) *ImportSourceImportSourceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ImportSourceImportSourceType.
func (m ImportSourceImportSourceType) Pointer() *ImportSourceImportSourceType {
	return &m
}

const (

	// ImportSourceImportSourceTypeS3 captures enum value "S3"
	ImportSourceImportSourceTypeS3 ImportSourceImportSourceType = "S3"

	// ImportSourceImportSourceTypeGCS captures enum value "GCS"
	ImportSourceImportSourceTypeGCS ImportSourceImportSourceType = "GCS"

	// ImportSourceImportSourceTypeLOCALFILE captures enum value "LOCAL_FILE"
	ImportSourceImportSourceTypeLOCALFILE ImportSourceImportSourceType = "LOCAL_FILE"

	// ImportSourceImportSourceTypeAZBLOB captures enum value "AZBLOB"
	ImportSourceImportSourceTypeAZBLOB ImportSourceImportSourceType = "AZBLOB"
)

// for schema
var importSourceImportSourceTypeEnum []interface{}

func init() {
	var res []ImportSourceImportSourceType
	if err := json.Unmarshal([]byte(`["S3","GCS","LOCAL_FILE","AZBLOB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		importSourceImportSourceTypeEnum = append(importSourceImportSourceTypeEnum, v)
	}
}

func (m ImportSourceImportSourceType) validateImportSourceImportSourceTypeEnum(path, location string, value ImportSourceImportSourceType) error {
	if err := validate.EnumCase(path, location, value, importSourceImportSourceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this import source import source type
func (m ImportSourceImportSourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateImportSourceImportSourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this import source import source type based on context it is used
func (m ImportSourceImportSourceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

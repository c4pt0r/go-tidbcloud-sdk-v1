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

// OpenapiImportSpec ImportSpec
//
// ImportSpec represents the settings and specifications of an import task.
//
// swagger:model openapiImportSpec
type OpenapiImportSpec struct {

	// source
	// Required: true
	Source *OpenapiImportSpecSource `json:"source"`

	// target
	// Required: true
	Target *OpenapiImportSpecTarget `json:"target"`
}

// Validate validates this openapi import spec
func (m *OpenapiImportSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpec) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpec) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi import spec based on the context it is used
func (m *OpenapiImportSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpec) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpec) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpec) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecSource ImportSource
//
// The data source settings of the import task.
//
// swagger:model OpenapiImportSpecSource
type OpenapiImportSpecSource struct {

	// aws assume role access
	AwsAssumeRoleAccess *OpenapiImportSpecSourceAwsAssumeRoleAccess `json:"aws_assume_role_access,omitempty"`

	// aws key access
	AwsKeyAccess *OpenapiImportSpecSourceAwsKeyAccess `json:"aws_key_access,omitempty"`

	// azure token access
	AzureTokenAccess *OpenapiImportSpecSourceAzureTokenAccess `json:"azure_token_access,omitempty"`

	// format
	// Required: true
	Format *OpenapiImportSpecSourceFormat `json:"format"`

	// The data source type of an import task.
	//
	// - `"S3"`: import data from Amazon S3
	// - `"GCS"`: import data from Google Cloud Storage
	// - `"LOCAL_FILE"`: import data from a local file (only available for [TiDB Cloud Serverless](https://docs.pingcap.com/tidbcloud/select-cluster-tier#tidb-cloud-serverless) clusters). Before you import from a local file, you need to first upload the file using the [Upload a local file for an import task](#tag/Import/operation/UploadLocalFile) endpoint.
	//
	// **Note:** Currently, if this import spec is used for a [preview](#tag/Import/operation/PreviewImportData) request, only the `LOCAL_FILE` source type is supported.
	// Example: S3
	// Required: true
	// Enum: ["S3","GCS","LOCAL_FILE","AZBLOB"]
	Type *string `json:"type"`

	// The data source URI of an import task. The URI scheme must match the data source type. Here are the scheme of each source type:
	// * `S3`: `s3://`
	// * `GCS`: `gs://`
	// * `LOCAL_FILE`: `file://`.
	//
	// **Note:** If the import source type is `LOCAL_FILE`, just provide the `upload_stub_id` of the uploaded file from the response of [Upload a local file for an import task](#tag/Import/operation/UploadLocalFile), and make it as the data source folder. For example: `file://12345/`.
	//
	// **Limitation**: If the import source type is `LOCAL_FILE`, only the `CSV` source format type is supported.
	// Example: s3://example-bucket/example-source-data/
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this openapi import spec source
func (m *OpenapiImportSpecSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsAssumeRoleAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsKeyAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureTokenAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSource) validateAwsAssumeRoleAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsAssumeRoleAccess) { // not required
		return nil
	}

	if m.AwsAssumeRoleAccess != nil {
		if err := m.AwsAssumeRoleAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "aws_assume_role_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "aws_assume_role_access")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpecSource) validateAwsKeyAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsKeyAccess) { // not required
		return nil
	}

	if m.AwsKeyAccess != nil {
		if err := m.AwsKeyAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "aws_key_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "aws_key_access")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpecSource) validateAzureTokenAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureTokenAccess) { // not required
		return nil
	}

	if m.AzureTokenAccess != nil {
		if err := m.AzureTokenAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "azure_token_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "azure_token_access")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpecSource) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"format", "body", m.Format); err != nil {
		return err
	}

	if m.Format != nil {
		if err := m.Format.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "format")
			}
			return err
		}
	}

	return nil
}

var openapiImportSpecSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["S3","GCS","LOCAL_FILE","AZBLOB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiImportSpecSourceTypeTypePropEnum = append(openapiImportSpecSourceTypeTypePropEnum, v)
	}
}

const (

	// OpenapiImportSpecSourceTypeS3 captures enum value "S3"
	OpenapiImportSpecSourceTypeS3 string = "S3"

	// OpenapiImportSpecSourceTypeGCS captures enum value "GCS"
	OpenapiImportSpecSourceTypeGCS string = "GCS"

	// OpenapiImportSpecSourceTypeLOCALFILE captures enum value "LOCAL_FILE"
	OpenapiImportSpecSourceTypeLOCALFILE string = "LOCAL_FILE"

	// OpenapiImportSpecSourceTypeAZBLOB captures enum value "AZBLOB"
	OpenapiImportSpecSourceTypeAZBLOB string = "AZBLOB"
)

// prop value enum
func (m *OpenapiImportSpecSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiImportSpecSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiImportSpecSource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("source"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiImportSpecSource) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"uri", "body", m.URI); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi import spec source based on the context it is used
func (m *OpenapiImportSpecSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsAssumeRoleAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsKeyAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureTokenAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSource) contextValidateAwsAssumeRoleAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsAssumeRoleAccess != nil {

		if swag.IsZero(m.AwsAssumeRoleAccess) { // not required
			return nil
		}

		if err := m.AwsAssumeRoleAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "aws_assume_role_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "aws_assume_role_access")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpecSource) contextValidateAwsKeyAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsKeyAccess != nil {

		if swag.IsZero(m.AwsKeyAccess) { // not required
			return nil
		}

		if err := m.AwsKeyAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "aws_key_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "aws_key_access")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpecSource) contextValidateAzureTokenAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureTokenAccess != nil {

		if swag.IsZero(m.AzureTokenAccess) { // not required
			return nil
		}

		if err := m.AzureTokenAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "azure_token_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "azure_token_access")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiImportSpecSource) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.Format != nil {

		if err := m.Format.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "format")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecSource) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecSourceAwsAssumeRoleAccess AwsAssumeRoleAccess
//
// The settings to access the S3 data by assuming a specific AWS role. This field is only needed if you need to access S3 data by assuming an AWS role.
//
// **Note:** Provide only one of `aws_assume_role_access` and `aws_key_access`. If both `aws_assume_role_access` and `aws_key_access` are provided, an error will be reported.
//
// swagger:model OpenapiImportSpecSourceAwsAssumeRoleAccess
type OpenapiImportSpecSourceAwsAssumeRoleAccess struct {

	// The specific AWS role ARN that needs to be assumed to access the Amazon S3 data source.
	// Example: arn:aws:iam::999999999999:role/sample-role
	// Required: true
	AssumeRole *string `json:"assume_role"`
}

// Validate validates this openapi import spec source aws assume role access
func (m *OpenapiImportSpecSourceAwsAssumeRoleAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssumeRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSourceAwsAssumeRoleAccess) validateAssumeRole(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"aws_assume_role_access"+"."+"assume_role", "body", m.AssumeRole); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi import spec source aws assume role access based on context it is used
func (m *OpenapiImportSpecSourceAwsAssumeRoleAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecSourceAwsAssumeRoleAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecSourceAwsAssumeRoleAccess) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecSourceAwsAssumeRoleAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecSourceAwsKeyAccess AwsKeyAccess
//
// The settings to access the S3 data with an access key. This field is only needed if you want to access the S3 data with an access key.
//
// **Note:** Provide only one of `aws_assume_role_access` and `aws_key_access`. If both `aws_assume_role_access` and `aws_key_access` are provided, an error will be reported.
//
// swagger:model OpenapiImportSpecSourceAwsKeyAccess
type OpenapiImportSpecSourceAwsKeyAccess struct {

	// The access key ID of the account to access the data. This information will be redacted when it is retrieved to obtain the import task information.
	// Example: YOUR_ACCESS_KEY_ID
	// Required: true
	AccessKeyID *string `json:"access_key_id"`

	// The secret access key for the account to access the data. This information will be redacted when it is retrieved to obtain the import task information.
	// Example: YOUR_SECRET_ACCESS_KEY
	// Required: true
	SecretAccessKey *string `json:"secret_access_key"`
}

// Validate validates this openapi import spec source aws key access
func (m *OpenapiImportSpecSourceAwsKeyAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSourceAwsKeyAccess) validateAccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"aws_key_access"+"."+"access_key_id", "body", m.AccessKeyID); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiImportSpecSourceAwsKeyAccess) validateSecretAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"aws_key_access"+"."+"secret_access_key", "body", m.SecretAccessKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi import spec source aws key access based on context it is used
func (m *OpenapiImportSpecSourceAwsKeyAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecSourceAwsKeyAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecSourceAwsKeyAccess) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecSourceAwsKeyAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecSourceAzureTokenAccess AzureTokenAccess
//
// The settings to access the Azblob data with an sas token. This field is only needed if you want to access the Azblob data with an sas token.
//
// swagger:model OpenapiImportSpecSourceAzureTokenAccess
type OpenapiImportSpecSourceAzureTokenAccess struct {

	// The sas token to access the data. This information will be redacted when it is retrieved to obtain the import task information.
	// Example: YOUR_SAS_TOKEN
	// Required: true
	SasToken *string `json:"sas_token"`
}

// Validate validates this openapi import spec source azure token access
func (m *OpenapiImportSpecSourceAzureTokenAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSasToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSourceAzureTokenAccess) validateSasToken(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"azure_token_access"+"."+"sas_token", "body", m.SasToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi import spec source azure token access based on context it is used
func (m *OpenapiImportSpecSourceAzureTokenAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecSourceAzureTokenAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecSourceAzureTokenAccess) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecSourceAzureTokenAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecSourceFormat ImportSourceFormat
//
// The format settings of the import data source.
//
// swagger:model OpenapiImportSpecSourceFormat
type OpenapiImportSpecSourceFormat struct {

	// csv config
	CsvConfig *OpenapiImportSpecSourceFormatCsvConfig `json:"csv_config,omitempty"`

	// The format type of an import source.
	// Example: CSV
	// Required: true
	// Enum: ["CSV","PARQUET","SQL","AURORA_SNAPSHOT"]
	Type *string `json:"type"`
}

// Validate validates this openapi import spec source format
func (m *OpenapiImportSpecSourceFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCsvConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSourceFormat) validateCsvConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.CsvConfig) { // not required
		return nil
	}

	if m.CsvConfig != nil {
		if err := m.CsvConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "format" + "." + "csv_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "format" + "." + "csv_config")
			}
			return err
		}
	}

	return nil
}

var openapiImportSpecSourceFormatTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CSV","PARQUET","SQL","AURORA_SNAPSHOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiImportSpecSourceFormatTypeTypePropEnum = append(openapiImportSpecSourceFormatTypeTypePropEnum, v)
	}
}

const (

	// OpenapiImportSpecSourceFormatTypeCSV captures enum value "CSV"
	OpenapiImportSpecSourceFormatTypeCSV string = "CSV"

	// OpenapiImportSpecSourceFormatTypePARQUET captures enum value "PARQUET"
	OpenapiImportSpecSourceFormatTypePARQUET string = "PARQUET"

	// OpenapiImportSpecSourceFormatTypeSQL captures enum value "SQL"
	OpenapiImportSpecSourceFormatTypeSQL string = "SQL"

	// OpenapiImportSpecSourceFormatTypeAURORASNAPSHOT captures enum value "AURORA_SNAPSHOT"
	OpenapiImportSpecSourceFormatTypeAURORASNAPSHOT string = "AURORA_SNAPSHOT"
)

// prop value enum
func (m *OpenapiImportSpecSourceFormat) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiImportSpecSourceFormatTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiImportSpecSourceFormat) validateType(formats strfmt.Registry) error {

	if err := validate.Required("source"+"."+"format"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("source"+"."+"format"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi import spec source format based on the context it is used
func (m *OpenapiImportSpecSourceFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCsvConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecSourceFormat) contextValidateCsvConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CsvConfig != nil {

		if swag.IsZero(m.CsvConfig) { // not required
			return nil
		}

		if err := m.CsvConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source" + "." + "format" + "." + "csv_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source" + "." + "format" + "." + "csv_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecSourceFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecSourceFormat) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecSourceFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecSourceFormatCsvConfig ImportSourceCSVConfig
//
// The CSV format settings to parse the source CSV files. This field is only needed if the source format is CSV.
//
// swagger:model OpenapiImportSpecSourceFormatCsvConfig
type OpenapiImportSpecSourceFormatCsvConfig struct {

	// Whether a backslash (`\`) symbol followed by a character should be combined as a whole and treated as an escape sequence in a CSV field. For example, if this parameter is set to `true`, `\n` will be treated as a 'new-line' character. If it is set to `false`, `\n` will be treated as two separate characters: backslash and `n`.
	//
	// Currently, these are several supported escape sequences: `\0`, `\b`, `\n`, `\r`, `\t`, and `\Z`. If the parameter is set to `true`, but the backslash escape sequence is not recognized, the backslash character is ignored.
	BackslashEscape *bool `json:"backslash_escape,omitempty"`

	// The delimiter character used to separate fields in the CSV data.
	Delimiter *string `json:"delimiter,omitempty"`

	// Whether the CSV data has a header row, which is not part of the data. If it is set to `true`, the import task will use the column names in the header row to match the column names in the target table.
	HasHeaderRow *bool `json:"has_header_row,omitempty"`

	// The character used to quote the fields in the CSV data.
	Quote *string `json:"quote,omitempty"`
}

// Validate validates this openapi import spec source format csv config
func (m *OpenapiImportSpecSourceFormatCsvConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi import spec source format csv config based on context it is used
func (m *OpenapiImportSpecSourceFormatCsvConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecSourceFormatCsvConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecSourceFormatCsvConfig) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecSourceFormatCsvConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecTarget ImportTarget
//
// The target settings of the import task.
//
// swagger:model OpenapiImportSpecTarget
type OpenapiImportSpecTarget struct {

	// The settings for each target table that is being imported for the import task. If you leave it empty, the system will scan all the files in the data source using the default file patterns and collect all the tables to import. The files include data files, table schema files, and DB schema files. If you provide a list of tables, only those tables will be imported. For more information about the default file pattern, see [Import CSV Files from Amazon S3 or GCS into TiDB Cloud](https://docs.pingcap.com/tidbcloud/import-csv-files).
	//
	// **Limitations:**
	// * Currently, if you want to use a custom filename pattern, you can only specify one table. If all the tables use the default filename pattern, you can specify more than one target table in `tables`.
	// * It is recommended that you pre-create the target tables before creating an import task. You can do this either by executing the `CREATE TABLE` statement in the cluster or by specifying the table definition in the table creation options.
	// * If a target table is not created, the import module tries to find a **TABLE SCHEMA FILE** containing the `CREATE TABLE` statement of the table in the data source folder with the name `${db_name}.${table_name}-schema.sql` (for example, `db01.tbl01-schema.sql`). If this file is found, the `CREATE TABLE` statement is automatically executed if the table doesn't exist before the actual import process starts. If the table is still missing after this pre-create step, an error will occur.
	Tables []*OpenapiImportSpecTargetTablesItems0 `json:"tables"`
}

// Validate validates this openapi import spec target
func (m *OpenapiImportSpecTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecTarget) validateTables(formats strfmt.Registry) error {
	if swag.IsZero(m.Tables) { // not required
		return nil
	}

	for i := 0; i < len(m.Tables); i++ {
		if swag.IsZero(m.Tables[i]) { // not required
			continue
		}

		if m.Tables[i] != nil {
			if err := m.Tables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target" + "." + "tables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target" + "." + "tables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this openapi import spec target based on the context it is used
func (m *OpenapiImportSpecTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecTarget) contextValidateTables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tables); i++ {

		if m.Tables[i] != nil {

			if swag.IsZero(m.Tables[i]) { // not required
				return nil
			}

			if err := m.Tables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target" + "." + "tables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target" + "." + "tables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecTarget) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSpecTargetTablesItems0 ImportTargetTable
//
// ImportTargetTable represents the settings for importing source data into a single target table of an import task.
//
// swagger:model OpenapiImportSpecTargetTablesItems0
type OpenapiImportSpecTargetTablesItems0 struct {

	// The target database name.
	// Example: db01
	// Required: true
	DatabaseName *string `json:"database_name"`

	// The filename pattern used to map the files in the data source to this target table. The pattern should be a simple glob pattern. Here are some examples:
	// * `my-data?.csv`: all CSV files starting with `my-data` and one character (such as `my-data1.csv` and `my-data2.csv`) will be imported into the same target table.
	// * `my-data*.csv`: all CSV files starting with `my-data` will be imported into the same target table.
	//
	// If no pattern is specified, a default pattern is used. The default pattern will try to find files with this naming convention as the data files for this table: `${db_name}.${table_name}.[numeric_index].${format_suffix}`.
	//
	// Here are some examples of filenames that can be matched as data files for the table `db01.table01`: `db01.table01.csv`, `db01.table01.00001.csv`.
	//
	// For more information about the custom file pattern and the default pattern, refer to [Import CSV Files from Amazon S3 or GCS into TiDB Cloud](https://docs.pingcap.com/tidbcloud/import-csv-files).
	//
	// **Note:** For `LOCAL_FILE` import tasks, use the local file name for this field. The local file name must match the local file name in [Upload a local file for an import task](#tag/Import/operation/UploadLocalFile).
	// Example: data/db01/table01.*.csv
	FileNamePattern string `json:"file_name_pattern,omitempty"`

	// The target table name.
	// Example: table01
	// Required: true
	TableName *string `json:"table_name"`
}

// Validate validates this openapi import spec target tables items0
func (m *OpenapiImportSpecTargetTablesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSpecTargetTablesItems0) validateDatabaseName(formats strfmt.Registry) error {

	if err := validate.Required("database_name", "body", m.DatabaseName); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiImportSpecTargetTablesItems0) validateTableName(formats strfmt.Registry) error {

	if err := validate.Required("table_name", "body", m.TableName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi import spec target tables items0 based on context it is used
func (m *OpenapiImportSpecTargetTablesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSpecTargetTablesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSpecTargetTablesItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSpecTargetTablesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

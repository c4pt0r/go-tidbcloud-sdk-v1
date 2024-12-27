// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiCreatePrivateEndpointServiceResp CreatePrivateEndpointServiceResp
//
// CreatePrivateEndpointServiceResp is the response for creating private endpoint service for a cluster.
//
// swagger:model openapiCreatePrivateEndpointServiceResp
type OpenapiCreatePrivateEndpointServiceResp struct {

	// private endpoint service
	// Required: true
	PrivateEndpointService *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService `json:"private_endpoint_service"`
}

// Validate validates this openapi create private endpoint service resp
func (m *OpenapiCreatePrivateEndpointServiceResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateEndpointService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointServiceResp) validatePrivateEndpointService(formats strfmt.Registry) error {

	if err := validate.Required("private_endpoint_service", "body", m.PrivateEndpointService); err != nil {
		return err
	}

	if m.PrivateEndpointService != nil {
		if err := m.PrivateEndpointService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_endpoint_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("private_endpoint_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi create private endpoint service resp based on the context it is used
func (m *OpenapiCreatePrivateEndpointServiceResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateEndpointService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointServiceResp) contextValidatePrivateEndpointService(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivateEndpointService != nil {

		if err := m.PrivateEndpointService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_endpoint_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("private_endpoint_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointServiceResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointServiceResp) UnmarshalBinary(b []byte) error {
	var res OpenapiCreatePrivateEndpointServiceResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService The newly created private endpoint service resource.
//
// swagger:model OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService
type OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService struct {

	// Availability zones for the private endpoint service. This field is only applicable when the `cloud_provider` is `"AWS"`.
	// Example: ["usw2-az2","usw2-az1"]
	AzIds []string `json:"az_ids"`

	// The cloud provider on which the private endpoint service is hosted.
	// - `"AWS"`: the Amazon Web Services cloud provider
	// - `"GCP"`: the Google Cloud cloud provider
	// Example: AWS
	// Enum: ["AWS","GCP"]
	CloudProvider string `json:"cloud_provider,omitempty"`

	// The DNS name of the private endpoint service.
	// Example: privatelink-tidb.01234567890.clusters.tidb-cloud.com
	DNSName string `json:"dns_name,omitempty"`

	// The name of the private endpoint service, which is used for connection.
	// Example: com.amazonaws.vpce.us-west-2.vpce-svc-01234567890123456
	Name string `json:"name,omitempty"`

	// The port of the private endpoint service.
	// Example: 4000
	Port int32 `json:"port,omitempty"`

	// The status of the private endpoint service.
	// Example: ACTIVE
	// Enum: ["CREATING","ACTIVE","DELETING"]
	Status string `json:"status,omitempty"`
}

// Validate validates this openapi create private endpoint service resp private endpoint service
func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeCloudProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeCloudProviderPropEnum = append(openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeCloudProviderPropEnum, v)
	}
}

const (

	// OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceCloudProviderAWS captures enum value "AWS"
	OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceCloudProviderAWS string = "AWS"

	// OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceCloudProviderGCP captures enum value "GCP"
	OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceCloudProviderGCP string = "GCP"
)

// prop value enum
func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) validateCloudProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeCloudProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) validateCloudProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudProviderEnum("private_endpoint_service"+"."+"cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

var openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeStatusPropEnum = append(openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeStatusPropEnum, v)
	}
}

const (

	// OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceStatusCREATING captures enum value "CREATING"
	OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceStatusCREATING string = "CREATING"

	// OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceStatusACTIVE captures enum value "ACTIVE"
	OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceStatusACTIVE string = "ACTIVE"

	// OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceStatusDELETING captures enum value "DELETING"
	OpenapiCreatePrivateEndpointServiceRespPrivateEndpointServiceStatusDELETING string = "DELETING"
)

// prop value enum
func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiCreatePrivateEndpointServiceRespPrivateEndpointServiceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("private_endpoint_service"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi create private endpoint service resp private endpoint service based on context it is used
func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService) UnmarshalBinary(b []byte) error {
	var res OpenapiCreatePrivateEndpointServiceRespPrivateEndpointService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

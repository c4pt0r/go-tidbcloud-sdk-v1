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

// PrivateEndpointPrivateLinkService private endpoint private link service
//
// swagger:model PrivateEndpointPrivateLinkService
type PrivateEndpointPrivateLinkService struct {

	// The name of the private endpoint service, which is used for connection.
	// Example: com.amazonaws.vpce.us-west-2.vpce-svc-01234567890123456
	ServiceName string `json:"service_name,omitempty"`

	// The status of the private endpoint service.
	// Example: ACTIVE
	// Enum: ["CREATING","ACTIVE","DELETING"]
	Status string `json:"status,omitempty"`
}

// Validate validates this private endpoint private link service
func (m *PrivateEndpointPrivateLinkService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var privateEndpointPrivateLinkServiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		privateEndpointPrivateLinkServiceTypeStatusPropEnum = append(privateEndpointPrivateLinkServiceTypeStatusPropEnum, v)
	}
}

const (

	// PrivateEndpointPrivateLinkServiceStatusCREATING captures enum value "CREATING"
	PrivateEndpointPrivateLinkServiceStatusCREATING string = "CREATING"

	// PrivateEndpointPrivateLinkServiceStatusACTIVE captures enum value "ACTIVE"
	PrivateEndpointPrivateLinkServiceStatusACTIVE string = "ACTIVE"

	// PrivateEndpointPrivateLinkServiceStatusDELETING captures enum value "DELETING"
	PrivateEndpointPrivateLinkServiceStatusDELETING string = "DELETING"
)

// prop value enum
func (m *PrivateEndpointPrivateLinkService) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, privateEndpointPrivateLinkServiceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PrivateEndpointPrivateLinkService) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this private endpoint private link service based on context it is used
func (m *PrivateEndpointPrivateLinkService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateEndpointPrivateLinkService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateEndpointPrivateLinkService) UnmarshalBinary(b []byte) error {
	var res PrivateEndpointPrivateLinkService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

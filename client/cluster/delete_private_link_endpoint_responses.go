// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeletePrivateLinkEndpointReader is a Reader for the DeletePrivateLinkEndpoint structure.
type DeletePrivateLinkEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateLinkEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePrivateLinkEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePrivateLinkEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePrivateLinkEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePrivateLinkEndpointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePrivateLinkEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePrivateLinkEndpointTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePrivateLinkEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePrivateLinkEndpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePrivateLinkEndpointOK creates a DeletePrivateLinkEndpointOK with default headers values
func NewDeletePrivateLinkEndpointOK() *DeletePrivateLinkEndpointOK {
	return &DeletePrivateLinkEndpointOK{}
}

/*
DeletePrivateLinkEndpointOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeletePrivateLinkEndpointOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete private link endpoint o k response has a 2xx status code
func (o *DeletePrivateLinkEndpointOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete private link endpoint o k response has a 3xx status code
func (o *DeletePrivateLinkEndpointOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint o k response has a 4xx status code
func (o *DeletePrivateLinkEndpointOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete private link endpoint o k response has a 5xx status code
func (o *DeletePrivateLinkEndpointOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private link endpoint o k response a status code equal to that given
func (o *DeletePrivateLinkEndpointOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete private link endpoint o k response
func (o *DeletePrivateLinkEndpointOK) Code() int {
	return 200
}

func (o *DeletePrivateLinkEndpointOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointOK %s", 200, payload)
}

func (o *DeletePrivateLinkEndpointOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointOK %s", 200, payload)
}

func (o *DeletePrivateLinkEndpointOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointBadRequest creates a DeletePrivateLinkEndpointBadRequest with default headers values
func NewDeletePrivateLinkEndpointBadRequest() *DeletePrivateLinkEndpointBadRequest {
	return &DeletePrivateLinkEndpointBadRequest{}
}

/*
DeletePrivateLinkEndpointBadRequest describes a response with status code 400, with default header values.

A request field is invalid.
*/
type DeletePrivateLinkEndpointBadRequest struct {
	Payload *DeletePrivateLinkEndpointBadRequestBody
}

// IsSuccess returns true when this delete private link endpoint bad request response has a 2xx status code
func (o *DeletePrivateLinkEndpointBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private link endpoint bad request response has a 3xx status code
func (o *DeletePrivateLinkEndpointBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint bad request response has a 4xx status code
func (o *DeletePrivateLinkEndpointBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private link endpoint bad request response has a 5xx status code
func (o *DeletePrivateLinkEndpointBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private link endpoint bad request response a status code equal to that given
func (o *DeletePrivateLinkEndpointBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete private link endpoint bad request response
func (o *DeletePrivateLinkEndpointBadRequest) Code() int {
	return 400
}

func (o *DeletePrivateLinkEndpointBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointBadRequest %s", 400, payload)
}

func (o *DeletePrivateLinkEndpointBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointBadRequest %s", 400, payload)
}

func (o *DeletePrivateLinkEndpointBadRequest) GetPayload() *DeletePrivateLinkEndpointBadRequestBody {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateLinkEndpointBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointUnauthorized creates a DeletePrivateLinkEndpointUnauthorized with default headers values
func NewDeletePrivateLinkEndpointUnauthorized() *DeletePrivateLinkEndpointUnauthorized {
	return &DeletePrivateLinkEndpointUnauthorized{}
}

/*
DeletePrivateLinkEndpointUnauthorized describes a response with status code 401, with default header values.

The API key cannot be authenticated.
*/
type DeletePrivateLinkEndpointUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this delete private link endpoint unauthorized response has a 2xx status code
func (o *DeletePrivateLinkEndpointUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private link endpoint unauthorized response has a 3xx status code
func (o *DeletePrivateLinkEndpointUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint unauthorized response has a 4xx status code
func (o *DeletePrivateLinkEndpointUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private link endpoint unauthorized response has a 5xx status code
func (o *DeletePrivateLinkEndpointUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private link endpoint unauthorized response a status code equal to that given
func (o *DeletePrivateLinkEndpointUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete private link endpoint unauthorized response
func (o *DeletePrivateLinkEndpointUnauthorized) Code() int {
	return 401
}

func (o *DeletePrivateLinkEndpointUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointUnauthorized %s", 401, payload)
}

func (o *DeletePrivateLinkEndpointUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointUnauthorized %s", 401, payload)
}

func (o *DeletePrivateLinkEndpointUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointForbidden creates a DeletePrivateLinkEndpointForbidden with default headers values
func NewDeletePrivateLinkEndpointForbidden() *DeletePrivateLinkEndpointForbidden {
	return &DeletePrivateLinkEndpointForbidden{}
}

/*
DeletePrivateLinkEndpointForbidden describes a response with status code 403, with default header values.

The API key does not have permission to access the resource.
*/
type DeletePrivateLinkEndpointForbidden struct {
	Payload *DeletePrivateLinkEndpointForbiddenBody
}

// IsSuccess returns true when this delete private link endpoint forbidden response has a 2xx status code
func (o *DeletePrivateLinkEndpointForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private link endpoint forbidden response has a 3xx status code
func (o *DeletePrivateLinkEndpointForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint forbidden response has a 4xx status code
func (o *DeletePrivateLinkEndpointForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private link endpoint forbidden response has a 5xx status code
func (o *DeletePrivateLinkEndpointForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private link endpoint forbidden response a status code equal to that given
func (o *DeletePrivateLinkEndpointForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete private link endpoint forbidden response
func (o *DeletePrivateLinkEndpointForbidden) Code() int {
	return 403
}

func (o *DeletePrivateLinkEndpointForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointForbidden %s", 403, payload)
}

func (o *DeletePrivateLinkEndpointForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointForbidden %s", 403, payload)
}

func (o *DeletePrivateLinkEndpointForbidden) GetPayload() *DeletePrivateLinkEndpointForbiddenBody {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateLinkEndpointForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointNotFound creates a DeletePrivateLinkEndpointNotFound with default headers values
func NewDeletePrivateLinkEndpointNotFound() *DeletePrivateLinkEndpointNotFound {
	return &DeletePrivateLinkEndpointNotFound{}
}

/*
DeletePrivateLinkEndpointNotFound describes a response with status code 404, with default header values.

The requested resource does not exist.
*/
type DeletePrivateLinkEndpointNotFound struct {
	Payload *DeletePrivateLinkEndpointNotFoundBody
}

// IsSuccess returns true when this delete private link endpoint not found response has a 2xx status code
func (o *DeletePrivateLinkEndpointNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private link endpoint not found response has a 3xx status code
func (o *DeletePrivateLinkEndpointNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint not found response has a 4xx status code
func (o *DeletePrivateLinkEndpointNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private link endpoint not found response has a 5xx status code
func (o *DeletePrivateLinkEndpointNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private link endpoint not found response a status code equal to that given
func (o *DeletePrivateLinkEndpointNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete private link endpoint not found response
func (o *DeletePrivateLinkEndpointNotFound) Code() int {
	return 404
}

func (o *DeletePrivateLinkEndpointNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointNotFound %s", 404, payload)
}

func (o *DeletePrivateLinkEndpointNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointNotFound %s", 404, payload)
}

func (o *DeletePrivateLinkEndpointNotFound) GetPayload() *DeletePrivateLinkEndpointNotFoundBody {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateLinkEndpointNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointTooManyRequests creates a DeletePrivateLinkEndpointTooManyRequests with default headers values
func NewDeletePrivateLinkEndpointTooManyRequests() *DeletePrivateLinkEndpointTooManyRequests {
	return &DeletePrivateLinkEndpointTooManyRequests{}
}

/*
DeletePrivateLinkEndpointTooManyRequests describes a response with status code 429, with default header values.

You have exceed the rate limit.
*/
type DeletePrivateLinkEndpointTooManyRequests struct {
	Payload *DeletePrivateLinkEndpointTooManyRequestsBody
}

// IsSuccess returns true when this delete private link endpoint too many requests response has a 2xx status code
func (o *DeletePrivateLinkEndpointTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private link endpoint too many requests response has a 3xx status code
func (o *DeletePrivateLinkEndpointTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint too many requests response has a 4xx status code
func (o *DeletePrivateLinkEndpointTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private link endpoint too many requests response has a 5xx status code
func (o *DeletePrivateLinkEndpointTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private link endpoint too many requests response a status code equal to that given
func (o *DeletePrivateLinkEndpointTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete private link endpoint too many requests response
func (o *DeletePrivateLinkEndpointTooManyRequests) Code() int {
	return 429
}

func (o *DeletePrivateLinkEndpointTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointTooManyRequests %s", 429, payload)
}

func (o *DeletePrivateLinkEndpointTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointTooManyRequests %s", 429, payload)
}

func (o *DeletePrivateLinkEndpointTooManyRequests) GetPayload() *DeletePrivateLinkEndpointTooManyRequestsBody {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateLinkEndpointTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointInternalServerError creates a DeletePrivateLinkEndpointInternalServerError with default headers values
func NewDeletePrivateLinkEndpointInternalServerError() *DeletePrivateLinkEndpointInternalServerError {
	return &DeletePrivateLinkEndpointInternalServerError{}
}

/*
DeletePrivateLinkEndpointInternalServerError describes a response with status code 500, with default header values.

Server error.
*/
type DeletePrivateLinkEndpointInternalServerError struct {
	Payload *DeletePrivateLinkEndpointInternalServerErrorBody
}

// IsSuccess returns true when this delete private link endpoint internal server error response has a 2xx status code
func (o *DeletePrivateLinkEndpointInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private link endpoint internal server error response has a 3xx status code
func (o *DeletePrivateLinkEndpointInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private link endpoint internal server error response has a 4xx status code
func (o *DeletePrivateLinkEndpointInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete private link endpoint internal server error response has a 5xx status code
func (o *DeletePrivateLinkEndpointInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete private link endpoint internal server error response a status code equal to that given
func (o *DeletePrivateLinkEndpointInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete private link endpoint internal server error response
func (o *DeletePrivateLinkEndpointInternalServerError) Code() int {
	return 500
}

func (o *DeletePrivateLinkEndpointInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointInternalServerError %s", 500, payload)
}

func (o *DeletePrivateLinkEndpointInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] deletePrivateLinkEndpointInternalServerError %s", 500, payload)
}

func (o *DeletePrivateLinkEndpointInternalServerError) GetPayload() *DeletePrivateLinkEndpointInternalServerErrorBody {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateLinkEndpointInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateLinkEndpointDefault creates a DeletePrivateLinkEndpointDefault with default headers values
func NewDeletePrivateLinkEndpointDefault(code int) *DeletePrivateLinkEndpointDefault {
	return &DeletePrivateLinkEndpointDefault{
		_statusCode: code,
	}
}

/*
DeletePrivateLinkEndpointDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeletePrivateLinkEndpointDefault struct {
	_statusCode int

	Payload *DeletePrivateLinkEndpointDefaultBody
}

// IsSuccess returns true when this delete private link endpoint default response has a 2xx status code
func (o *DeletePrivateLinkEndpointDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete private link endpoint default response has a 3xx status code
func (o *DeletePrivateLinkEndpointDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete private link endpoint default response has a 4xx status code
func (o *DeletePrivateLinkEndpointDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete private link endpoint default response has a 5xx status code
func (o *DeletePrivateLinkEndpointDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete private link endpoint default response a status code equal to that given
func (o *DeletePrivateLinkEndpointDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete private link endpoint default response
func (o *DeletePrivateLinkEndpointDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateLinkEndpointDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] DeletePrivateLinkEndpoint default %s", o._statusCode, payload)
}

func (o *DeletePrivateLinkEndpointDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/accesspoints/{accesspoint_id}/privateLinkEndpoints/{private_link_endpoint_id}][%d] DeletePrivateLinkEndpoint default %s", o._statusCode, payload)
}

func (o *DeletePrivateLinkEndpointDefault) GetPayload() *DeletePrivateLinkEndpointDefaultBody {
	return o.Payload
}

func (o *DeletePrivateLinkEndpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateLinkEndpointDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeletePrivateLinkEndpointBadRequestBody delete private link endpoint bad request body
swagger:model DeletePrivateLinkEndpointBadRequestBody
*/
type DeletePrivateLinkEndpointBadRequestBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private link endpoint bad request body
func (o *DeletePrivateLinkEndpointBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private link endpoint bad request body based on context it is used
func (o *DeletePrivateLinkEndpointBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateLinkEndpointDefaultBody delete private link endpoint default body
swagger:model DeletePrivateLinkEndpointDefaultBody
*/
type DeletePrivateLinkEndpointDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*DeletePrivateLinkEndpointDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private link endpoint default body
func (o *DeletePrivateLinkEndpointDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePrivateLinkEndpointDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeletePrivateLinkEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeletePrivateLinkEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete private link endpoint default body based on the context it is used
func (o *DeletePrivateLinkEndpointDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePrivateLinkEndpointDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeletePrivateLinkEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeletePrivateLinkEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateLinkEndpointDefaultBodyDetailsItems0 delete private link endpoint default body details items0
swagger:model DeletePrivateLinkEndpointDefaultBodyDetailsItems0
*/
type DeletePrivateLinkEndpointDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// delete private link endpoint default body details items0
	DeletePrivateLinkEndpointDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *DeletePrivateLinkEndpointDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DeletePrivateLinkEndpointDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.DeletePrivateLinkEndpointDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o DeletePrivateLinkEndpointDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.DeletePrivateLinkEndpointDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.DeletePrivateLinkEndpointDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this delete private link endpoint default body details items0
func (o *DeletePrivateLinkEndpointDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private link endpoint default body details items0 based on context it is used
func (o *DeletePrivateLinkEndpointDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateLinkEndpointForbiddenBody delete private link endpoint forbidden body
swagger:model DeletePrivateLinkEndpointForbiddenBody
*/
type DeletePrivateLinkEndpointForbiddenBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private link endpoint forbidden body
func (o *DeletePrivateLinkEndpointForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private link endpoint forbidden body based on context it is used
func (o *DeletePrivateLinkEndpointForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateLinkEndpointInternalServerErrorBody delete private link endpoint internal server error body
swagger:model DeletePrivateLinkEndpointInternalServerErrorBody
*/
type DeletePrivateLinkEndpointInternalServerErrorBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private link endpoint internal server error body
func (o *DeletePrivateLinkEndpointInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private link endpoint internal server error body based on context it is used
func (o *DeletePrivateLinkEndpointInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateLinkEndpointNotFoundBody delete private link endpoint not found body
swagger:model DeletePrivateLinkEndpointNotFoundBody
*/
type DeletePrivateLinkEndpointNotFoundBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private link endpoint not found body
func (o *DeletePrivateLinkEndpointNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private link endpoint not found body based on context it is used
func (o *DeletePrivateLinkEndpointNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateLinkEndpointTooManyRequestsBody delete private link endpoint too many requests body
swagger:model DeletePrivateLinkEndpointTooManyRequestsBody
*/
type DeletePrivateLinkEndpointTooManyRequestsBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private link endpoint too many requests body
func (o *DeletePrivateLinkEndpointTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private link endpoint too many requests body based on context it is used
func (o *DeletePrivateLinkEndpointTooManyRequestsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateLinkEndpointTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateLinkEndpointTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

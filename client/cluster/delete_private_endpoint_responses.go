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

// DeletePrivateEndpointReader is a Reader for the DeletePrivateEndpoint structure.
type DeletePrivateEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePrivateEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePrivateEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePrivateEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePrivateEndpointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePrivateEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePrivateEndpointTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePrivateEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePrivateEndpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePrivateEndpointOK creates a DeletePrivateEndpointOK with default headers values
func NewDeletePrivateEndpointOK() *DeletePrivateEndpointOK {
	return &DeletePrivateEndpointOK{}
}

/*
DeletePrivateEndpointOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeletePrivateEndpointOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete private endpoint o k response has a 2xx status code
func (o *DeletePrivateEndpointOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete private endpoint o k response has a 3xx status code
func (o *DeletePrivateEndpointOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint o k response has a 4xx status code
func (o *DeletePrivateEndpointOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete private endpoint o k response has a 5xx status code
func (o *DeletePrivateEndpointOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private endpoint o k response a status code equal to that given
func (o *DeletePrivateEndpointOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete private endpoint o k response
func (o *DeletePrivateEndpointOK) Code() int {
	return 200
}

func (o *DeletePrivateEndpointOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointOK %s", 200, payload)
}

func (o *DeletePrivateEndpointOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointOK %s", 200, payload)
}

func (o *DeletePrivateEndpointOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeletePrivateEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointBadRequest creates a DeletePrivateEndpointBadRequest with default headers values
func NewDeletePrivateEndpointBadRequest() *DeletePrivateEndpointBadRequest {
	return &DeletePrivateEndpointBadRequest{}
}

/*
DeletePrivateEndpointBadRequest describes a response with status code 400, with default header values.

A request field is invalid.
*/
type DeletePrivateEndpointBadRequest struct {
	Payload *DeletePrivateEndpointBadRequestBody
}

// IsSuccess returns true when this delete private endpoint bad request response has a 2xx status code
func (o *DeletePrivateEndpointBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private endpoint bad request response has a 3xx status code
func (o *DeletePrivateEndpointBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint bad request response has a 4xx status code
func (o *DeletePrivateEndpointBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private endpoint bad request response has a 5xx status code
func (o *DeletePrivateEndpointBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private endpoint bad request response a status code equal to that given
func (o *DeletePrivateEndpointBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete private endpoint bad request response
func (o *DeletePrivateEndpointBadRequest) Code() int {
	return 400
}

func (o *DeletePrivateEndpointBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointBadRequest %s", 400, payload)
}

func (o *DeletePrivateEndpointBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointBadRequest %s", 400, payload)
}

func (o *DeletePrivateEndpointBadRequest) GetPayload() *DeletePrivateEndpointBadRequestBody {
	return o.Payload
}

func (o *DeletePrivateEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateEndpointBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointUnauthorized creates a DeletePrivateEndpointUnauthorized with default headers values
func NewDeletePrivateEndpointUnauthorized() *DeletePrivateEndpointUnauthorized {
	return &DeletePrivateEndpointUnauthorized{}
}

/*
DeletePrivateEndpointUnauthorized describes a response with status code 401, with default header values.

The API key cannot be authenticated.
*/
type DeletePrivateEndpointUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this delete private endpoint unauthorized response has a 2xx status code
func (o *DeletePrivateEndpointUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private endpoint unauthorized response has a 3xx status code
func (o *DeletePrivateEndpointUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint unauthorized response has a 4xx status code
func (o *DeletePrivateEndpointUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private endpoint unauthorized response has a 5xx status code
func (o *DeletePrivateEndpointUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private endpoint unauthorized response a status code equal to that given
func (o *DeletePrivateEndpointUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete private endpoint unauthorized response
func (o *DeletePrivateEndpointUnauthorized) Code() int {
	return 401
}

func (o *DeletePrivateEndpointUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointUnauthorized %s", 401, payload)
}

func (o *DeletePrivateEndpointUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointUnauthorized %s", 401, payload)
}

func (o *DeletePrivateEndpointUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DeletePrivateEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointForbidden creates a DeletePrivateEndpointForbidden with default headers values
func NewDeletePrivateEndpointForbidden() *DeletePrivateEndpointForbidden {
	return &DeletePrivateEndpointForbidden{}
}

/*
DeletePrivateEndpointForbidden describes a response with status code 403, with default header values.

The API key does not have permission to access the resource.
*/
type DeletePrivateEndpointForbidden struct {
	Payload *DeletePrivateEndpointForbiddenBody
}

// IsSuccess returns true when this delete private endpoint forbidden response has a 2xx status code
func (o *DeletePrivateEndpointForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private endpoint forbidden response has a 3xx status code
func (o *DeletePrivateEndpointForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint forbidden response has a 4xx status code
func (o *DeletePrivateEndpointForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private endpoint forbidden response has a 5xx status code
func (o *DeletePrivateEndpointForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private endpoint forbidden response a status code equal to that given
func (o *DeletePrivateEndpointForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete private endpoint forbidden response
func (o *DeletePrivateEndpointForbidden) Code() int {
	return 403
}

func (o *DeletePrivateEndpointForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointForbidden %s", 403, payload)
}

func (o *DeletePrivateEndpointForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointForbidden %s", 403, payload)
}

func (o *DeletePrivateEndpointForbidden) GetPayload() *DeletePrivateEndpointForbiddenBody {
	return o.Payload
}

func (o *DeletePrivateEndpointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateEndpointForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointNotFound creates a DeletePrivateEndpointNotFound with default headers values
func NewDeletePrivateEndpointNotFound() *DeletePrivateEndpointNotFound {
	return &DeletePrivateEndpointNotFound{}
}

/*
DeletePrivateEndpointNotFound describes a response with status code 404, with default header values.

The requested resource does not exist.
*/
type DeletePrivateEndpointNotFound struct {
	Payload *DeletePrivateEndpointNotFoundBody
}

// IsSuccess returns true when this delete private endpoint not found response has a 2xx status code
func (o *DeletePrivateEndpointNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private endpoint not found response has a 3xx status code
func (o *DeletePrivateEndpointNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint not found response has a 4xx status code
func (o *DeletePrivateEndpointNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private endpoint not found response has a 5xx status code
func (o *DeletePrivateEndpointNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private endpoint not found response a status code equal to that given
func (o *DeletePrivateEndpointNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete private endpoint not found response
func (o *DeletePrivateEndpointNotFound) Code() int {
	return 404
}

func (o *DeletePrivateEndpointNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointNotFound %s", 404, payload)
}

func (o *DeletePrivateEndpointNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointNotFound %s", 404, payload)
}

func (o *DeletePrivateEndpointNotFound) GetPayload() *DeletePrivateEndpointNotFoundBody {
	return o.Payload
}

func (o *DeletePrivateEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateEndpointNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointTooManyRequests creates a DeletePrivateEndpointTooManyRequests with default headers values
func NewDeletePrivateEndpointTooManyRequests() *DeletePrivateEndpointTooManyRequests {
	return &DeletePrivateEndpointTooManyRequests{}
}

/*
DeletePrivateEndpointTooManyRequests describes a response with status code 429, with default header values.

You have exceed the rate limit.
*/
type DeletePrivateEndpointTooManyRequests struct {
	Payload *DeletePrivateEndpointTooManyRequestsBody
}

// IsSuccess returns true when this delete private endpoint too many requests response has a 2xx status code
func (o *DeletePrivateEndpointTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private endpoint too many requests response has a 3xx status code
func (o *DeletePrivateEndpointTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint too many requests response has a 4xx status code
func (o *DeletePrivateEndpointTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete private endpoint too many requests response has a 5xx status code
func (o *DeletePrivateEndpointTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete private endpoint too many requests response a status code equal to that given
func (o *DeletePrivateEndpointTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete private endpoint too many requests response
func (o *DeletePrivateEndpointTooManyRequests) Code() int {
	return 429
}

func (o *DeletePrivateEndpointTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointTooManyRequests %s", 429, payload)
}

func (o *DeletePrivateEndpointTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointTooManyRequests %s", 429, payload)
}

func (o *DeletePrivateEndpointTooManyRequests) GetPayload() *DeletePrivateEndpointTooManyRequestsBody {
	return o.Payload
}

func (o *DeletePrivateEndpointTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateEndpointTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointInternalServerError creates a DeletePrivateEndpointInternalServerError with default headers values
func NewDeletePrivateEndpointInternalServerError() *DeletePrivateEndpointInternalServerError {
	return &DeletePrivateEndpointInternalServerError{}
}

/*
DeletePrivateEndpointInternalServerError describes a response with status code 500, with default header values.

Server error.
*/
type DeletePrivateEndpointInternalServerError struct {
	Payload *DeletePrivateEndpointInternalServerErrorBody
}

// IsSuccess returns true when this delete private endpoint internal server error response has a 2xx status code
func (o *DeletePrivateEndpointInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete private endpoint internal server error response has a 3xx status code
func (o *DeletePrivateEndpointInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete private endpoint internal server error response has a 4xx status code
func (o *DeletePrivateEndpointInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete private endpoint internal server error response has a 5xx status code
func (o *DeletePrivateEndpointInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete private endpoint internal server error response a status code equal to that given
func (o *DeletePrivateEndpointInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete private endpoint internal server error response
func (o *DeletePrivateEndpointInternalServerError) Code() int {
	return 500
}

func (o *DeletePrivateEndpointInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointInternalServerError %s", 500, payload)
}

func (o *DeletePrivateEndpointInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] deletePrivateEndpointInternalServerError %s", 500, payload)
}

func (o *DeletePrivateEndpointInternalServerError) GetPayload() *DeletePrivateEndpointInternalServerErrorBody {
	return o.Payload
}

func (o *DeletePrivateEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateEndpointInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateEndpointDefault creates a DeletePrivateEndpointDefault with default headers values
func NewDeletePrivateEndpointDefault(code int) *DeletePrivateEndpointDefault {
	return &DeletePrivateEndpointDefault{
		_statusCode: code,
	}
}

/*
DeletePrivateEndpointDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeletePrivateEndpointDefault struct {
	_statusCode int

	Payload *DeletePrivateEndpointDefaultBody
}

// IsSuccess returns true when this delete private endpoint default response has a 2xx status code
func (o *DeletePrivateEndpointDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete private endpoint default response has a 3xx status code
func (o *DeletePrivateEndpointDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete private endpoint default response has a 4xx status code
func (o *DeletePrivateEndpointDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete private endpoint default response has a 5xx status code
func (o *DeletePrivateEndpointDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete private endpoint default response a status code equal to that given
func (o *DeletePrivateEndpointDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete private endpoint default response
func (o *DeletePrivateEndpointDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateEndpointDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] DeletePrivateEndpoint default %s", o._statusCode, payload)
}

func (o *DeletePrivateEndpointDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1beta/projects/{project_id}/clusters/{cluster_id}/private_endpoints/{endpoint_id}][%d] DeletePrivateEndpoint default %s", o._statusCode, payload)
}

func (o *DeletePrivateEndpointDefault) GetPayload() *DeletePrivateEndpointDefaultBody {
	return o.Payload
}

func (o *DeletePrivateEndpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePrivateEndpointDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeletePrivateEndpointBadRequestBody delete private endpoint bad request body
swagger:model DeletePrivateEndpointBadRequestBody
*/
type DeletePrivateEndpointBadRequestBody struct {

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

// Validate validates this delete private endpoint bad request body
func (o *DeletePrivateEndpointBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private endpoint bad request body based on context it is used
func (o *DeletePrivateEndpointBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateEndpointDefaultBody delete private endpoint default body
swagger:model DeletePrivateEndpointDefaultBody
*/
type DeletePrivateEndpointDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*DeletePrivateEndpointDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete private endpoint default body
func (o *DeletePrivateEndpointDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePrivateEndpointDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("DeletePrivateEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeletePrivateEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete private endpoint default body based on the context it is used
func (o *DeletePrivateEndpointDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePrivateEndpointDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeletePrivateEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeletePrivateEndpoint default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateEndpointDefaultBodyDetailsItems0 delete private endpoint default body details items0
swagger:model DeletePrivateEndpointDefaultBodyDetailsItems0
*/
type DeletePrivateEndpointDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// delete private endpoint default body details items0
	DeletePrivateEndpointDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *DeletePrivateEndpointDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DeletePrivateEndpointDefaultBodyDetailsItems0

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
		o.DeletePrivateEndpointDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o DeletePrivateEndpointDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.DeletePrivateEndpointDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.DeletePrivateEndpointDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this delete private endpoint default body details items0
func (o *DeletePrivateEndpointDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private endpoint default body details items0 based on context it is used
func (o *DeletePrivateEndpointDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateEndpointForbiddenBody delete private endpoint forbidden body
swagger:model DeletePrivateEndpointForbiddenBody
*/
type DeletePrivateEndpointForbiddenBody struct {

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

// Validate validates this delete private endpoint forbidden body
func (o *DeletePrivateEndpointForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private endpoint forbidden body based on context it is used
func (o *DeletePrivateEndpointForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateEndpointInternalServerErrorBody delete private endpoint internal server error body
swagger:model DeletePrivateEndpointInternalServerErrorBody
*/
type DeletePrivateEndpointInternalServerErrorBody struct {

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

// Validate validates this delete private endpoint internal server error body
func (o *DeletePrivateEndpointInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private endpoint internal server error body based on context it is used
func (o *DeletePrivateEndpointInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateEndpointNotFoundBody delete private endpoint not found body
swagger:model DeletePrivateEndpointNotFoundBody
*/
type DeletePrivateEndpointNotFoundBody struct {

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

// Validate validates this delete private endpoint not found body
func (o *DeletePrivateEndpointNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private endpoint not found body based on context it is used
func (o *DeletePrivateEndpointNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeletePrivateEndpointTooManyRequestsBody delete private endpoint too many requests body
swagger:model DeletePrivateEndpointTooManyRequestsBody
*/
type DeletePrivateEndpointTooManyRequestsBody struct {

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

// Validate validates this delete private endpoint too many requests body
func (o *DeletePrivateEndpointTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete private endpoint too many requests body based on context it is used
func (o *DeletePrivateEndpointTooManyRequestsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePrivateEndpointTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePrivateEndpointTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res DeletePrivateEndpointTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

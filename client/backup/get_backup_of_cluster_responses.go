// Code generated by go-swagger; DO NOT EDIT.

package backup

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
	"github.com/go-openapi/validate"
)

// GetBackupOfClusterReader is a Reader for the GetBackupOfCluster structure.
type GetBackupOfClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupOfClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupOfClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupOfClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBackupOfClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBackupOfClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBackupOfClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBackupOfClusterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackupOfClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBackupOfClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupOfClusterOK creates a GetBackupOfClusterOK with default headers values
func NewGetBackupOfClusterOK() *GetBackupOfClusterOK {
	return &GetBackupOfClusterOK{}
}

/*
GetBackupOfClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetBackupOfClusterOK struct {
	Payload *GetBackupOfClusterOKBody
}

// IsSuccess returns true when this get backup of cluster o k response has a 2xx status code
func (o *GetBackupOfClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup of cluster o k response has a 3xx status code
func (o *GetBackupOfClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster o k response has a 4xx status code
func (o *GetBackupOfClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup of cluster o k response has a 5xx status code
func (o *GetBackupOfClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup of cluster o k response a status code equal to that given
func (o *GetBackupOfClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBackupOfClusterOK) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterOK  %+v", 200, o.Payload)
}

func (o *GetBackupOfClusterOK) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterOK  %+v", 200, o.Payload)
}

func (o *GetBackupOfClusterOK) GetPayload() *GetBackupOfClusterOKBody {
	return o.Payload
}

func (o *GetBackupOfClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterBadRequest creates a GetBackupOfClusterBadRequest with default headers values
func NewGetBackupOfClusterBadRequest() *GetBackupOfClusterBadRequest {
	return &GetBackupOfClusterBadRequest{}
}

/*
GetBackupOfClusterBadRequest describes a response with status code 400, with default header values.

A request field is invalid.
*/
type GetBackupOfClusterBadRequest struct {
	Payload *GetBackupOfClusterBadRequestBody
}

// IsSuccess returns true when this get backup of cluster bad request response has a 2xx status code
func (o *GetBackupOfClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup of cluster bad request response has a 3xx status code
func (o *GetBackupOfClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster bad request response has a 4xx status code
func (o *GetBackupOfClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup of cluster bad request response has a 5xx status code
func (o *GetBackupOfClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup of cluster bad request response a status code equal to that given
func (o *GetBackupOfClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetBackupOfClusterBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterBadRequest  %+v", 400, o.Payload)
}

func (o *GetBackupOfClusterBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterBadRequest  %+v", 400, o.Payload)
}

func (o *GetBackupOfClusterBadRequest) GetPayload() *GetBackupOfClusterBadRequestBody {
	return o.Payload
}

func (o *GetBackupOfClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterUnauthorized creates a GetBackupOfClusterUnauthorized with default headers values
func NewGetBackupOfClusterUnauthorized() *GetBackupOfClusterUnauthorized {
	return &GetBackupOfClusterUnauthorized{}
}

/*
GetBackupOfClusterUnauthorized describes a response with status code 401, with default header values.

The API key cannot be authenticated.
*/
type GetBackupOfClusterUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this get backup of cluster unauthorized response has a 2xx status code
func (o *GetBackupOfClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup of cluster unauthorized response has a 3xx status code
func (o *GetBackupOfClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster unauthorized response has a 4xx status code
func (o *GetBackupOfClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup of cluster unauthorized response has a 5xx status code
func (o *GetBackupOfClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup of cluster unauthorized response a status code equal to that given
func (o *GetBackupOfClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBackupOfClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBackupOfClusterUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBackupOfClusterUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *GetBackupOfClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterForbidden creates a GetBackupOfClusterForbidden with default headers values
func NewGetBackupOfClusterForbidden() *GetBackupOfClusterForbidden {
	return &GetBackupOfClusterForbidden{}
}

/*
GetBackupOfClusterForbidden describes a response with status code 403, with default header values.

The API key does not have permission to access the resource.
*/
type GetBackupOfClusterForbidden struct {
	Payload *GetBackupOfClusterForbiddenBody
}

// IsSuccess returns true when this get backup of cluster forbidden response has a 2xx status code
func (o *GetBackupOfClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup of cluster forbidden response has a 3xx status code
func (o *GetBackupOfClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster forbidden response has a 4xx status code
func (o *GetBackupOfClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup of cluster forbidden response has a 5xx status code
func (o *GetBackupOfClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup of cluster forbidden response a status code equal to that given
func (o *GetBackupOfClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBackupOfClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterForbidden  %+v", 403, o.Payload)
}

func (o *GetBackupOfClusterForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterForbidden  %+v", 403, o.Payload)
}

func (o *GetBackupOfClusterForbidden) GetPayload() *GetBackupOfClusterForbiddenBody {
	return o.Payload
}

func (o *GetBackupOfClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterNotFound creates a GetBackupOfClusterNotFound with default headers values
func NewGetBackupOfClusterNotFound() *GetBackupOfClusterNotFound {
	return &GetBackupOfClusterNotFound{}
}

/*
GetBackupOfClusterNotFound describes a response with status code 404, with default header values.

The requested resource does not exist.
*/
type GetBackupOfClusterNotFound struct {
	Payload *GetBackupOfClusterNotFoundBody
}

// IsSuccess returns true when this get backup of cluster not found response has a 2xx status code
func (o *GetBackupOfClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup of cluster not found response has a 3xx status code
func (o *GetBackupOfClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster not found response has a 4xx status code
func (o *GetBackupOfClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup of cluster not found response has a 5xx status code
func (o *GetBackupOfClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup of cluster not found response a status code equal to that given
func (o *GetBackupOfClusterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBackupOfClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetBackupOfClusterNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetBackupOfClusterNotFound) GetPayload() *GetBackupOfClusterNotFoundBody {
	return o.Payload
}

func (o *GetBackupOfClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterTooManyRequests creates a GetBackupOfClusterTooManyRequests with default headers values
func NewGetBackupOfClusterTooManyRequests() *GetBackupOfClusterTooManyRequests {
	return &GetBackupOfClusterTooManyRequests{}
}

/*
GetBackupOfClusterTooManyRequests describes a response with status code 429, with default header values.

You have exceed the rate limit.
*/
type GetBackupOfClusterTooManyRequests struct {
	Payload *GetBackupOfClusterTooManyRequestsBody
}

// IsSuccess returns true when this get backup of cluster too many requests response has a 2xx status code
func (o *GetBackupOfClusterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup of cluster too many requests response has a 3xx status code
func (o *GetBackupOfClusterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster too many requests response has a 4xx status code
func (o *GetBackupOfClusterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup of cluster too many requests response has a 5xx status code
func (o *GetBackupOfClusterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup of cluster too many requests response a status code equal to that given
func (o *GetBackupOfClusterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetBackupOfClusterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBackupOfClusterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBackupOfClusterTooManyRequests) GetPayload() *GetBackupOfClusterTooManyRequestsBody {
	return o.Payload
}

func (o *GetBackupOfClusterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterInternalServerError creates a GetBackupOfClusterInternalServerError with default headers values
func NewGetBackupOfClusterInternalServerError() *GetBackupOfClusterInternalServerError {
	return &GetBackupOfClusterInternalServerError{}
}

/*
GetBackupOfClusterInternalServerError describes a response with status code 500, with default header values.

Server error.
*/
type GetBackupOfClusterInternalServerError struct {
	Payload *GetBackupOfClusterInternalServerErrorBody
}

// IsSuccess returns true when this get backup of cluster internal server error response has a 2xx status code
func (o *GetBackupOfClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup of cluster internal server error response has a 3xx status code
func (o *GetBackupOfClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup of cluster internal server error response has a 4xx status code
func (o *GetBackupOfClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup of cluster internal server error response has a 5xx status code
func (o *GetBackupOfClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get backup of cluster internal server error response a status code equal to that given
func (o *GetBackupOfClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetBackupOfClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBackupOfClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] getBackupOfClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBackupOfClusterInternalServerError) GetPayload() *GetBackupOfClusterInternalServerErrorBody {
	return o.Payload
}

func (o *GetBackupOfClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupOfClusterDefault creates a GetBackupOfClusterDefault with default headers values
func NewGetBackupOfClusterDefault(code int) *GetBackupOfClusterDefault {
	return &GetBackupOfClusterDefault{
		_statusCode: code,
	}
}

/*
GetBackupOfClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetBackupOfClusterDefault struct {
	_statusCode int

	Payload *GetBackupOfClusterDefaultBody
}

// Code gets the status code for the get backup of cluster default response
func (o *GetBackupOfClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get backup of cluster default response has a 2xx status code
func (o *GetBackupOfClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get backup of cluster default response has a 3xx status code
func (o *GetBackupOfClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get backup of cluster default response has a 4xx status code
func (o *GetBackupOfClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get backup of cluster default response has a 5xx status code
func (o *GetBackupOfClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get backup of cluster default response a status code equal to that given
func (o *GetBackupOfClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetBackupOfClusterDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] GetBackupOfCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackupOfClusterDefault) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects/{project_id}/clusters/{cluster_id}/backups/{backup_id}][%d] GetBackupOfCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackupOfClusterDefault) GetPayload() *GetBackupOfClusterDefaultBody {
	return o.Payload
}

func (o *GetBackupOfClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBackupOfClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetBackupOfClusterBadRequestBody get backup of cluster bad request body
swagger:model GetBackupOfClusterBadRequestBody
*/
type GetBackupOfClusterBadRequestBody struct {

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

// Validate validates this get backup of cluster bad request body
func (o *GetBackupOfClusterBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup of cluster bad request body based on context it is used
func (o *GetBackupOfClusterBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterDefaultBody get backup of cluster default body
swagger:model GetBackupOfClusterDefaultBody
*/
type GetBackupOfClusterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*GetBackupOfClusterDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get backup of cluster default body
func (o *GetBackupOfClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackupOfClusterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetBackupOfCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetBackupOfCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get backup of cluster default body based on the context it is used
func (o *GetBackupOfClusterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackupOfClusterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetBackupOfCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetBackupOfCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterDefaultBodyDetailsItems0 get backup of cluster default body details items0
swagger:model GetBackupOfClusterDefaultBodyDetailsItems0
*/
type GetBackupOfClusterDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get backup of cluster default body details items0
func (o *GetBackupOfClusterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup of cluster default body details items0 based on context it is used
func (o *GetBackupOfClusterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterForbiddenBody get backup of cluster forbidden body
swagger:model GetBackupOfClusterForbiddenBody
*/
type GetBackupOfClusterForbiddenBody struct {

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

// Validate validates this get backup of cluster forbidden body
func (o *GetBackupOfClusterForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup of cluster forbidden body based on context it is used
func (o *GetBackupOfClusterForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterInternalServerErrorBody get backup of cluster internal server error body
swagger:model GetBackupOfClusterInternalServerErrorBody
*/
type GetBackupOfClusterInternalServerErrorBody struct {

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

// Validate validates this get backup of cluster internal server error body
func (o *GetBackupOfClusterInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup of cluster internal server error body based on context it is used
func (o *GetBackupOfClusterInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterNotFoundBody get backup of cluster not found body
swagger:model GetBackupOfClusterNotFoundBody
*/
type GetBackupOfClusterNotFoundBody struct {

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

// Validate validates this get backup of cluster not found body
func (o *GetBackupOfClusterNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup of cluster not found body based on context it is used
func (o *GetBackupOfClusterNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterOKBody GetBackupOfClusterResp
//
// This response for getting backup of a cluster.
swagger:model GetBackupOfClusterOKBody
*/
type GetBackupOfClusterOKBody struct {

	// The creation time of the backup in UTC. The time format follows the [ISO8601](http://en.wikipedia.org/wiki/ISO_8601) standard, which is `YYYY-MM-DD` (year-month-day) + T +`HH:MM:SS` (hour-minutes-seconds) + Z. For example, `2020-01-01T00:00:00Z`.
	// Example: 2020-01-01T00:00:00Z
	// Format: date-time
	CreateTimestamp strfmt.DateTime `json:"create_timestamp,omitempty"`

	// The description of the backup. It is specified by the user when taking a manual type backup. It helps you add additional information to the backup.
	// Example: backup for cluster upgrade in 2022/06/07
	Description string `json:"description,omitempty"`

	// The ID of the backup.
	// Example: 1
	ID string `json:"id,omitempty"`

	// The name of the backup.
	// Example: backup-1
	Name string `json:"name,omitempty"`

	// The bytes of the backup.
	// Example: 102400
	Size string `json:"size,omitempty"`

	// The status of backup.
	// Example: SUCCESS
	// Enum: [PENDING RUNNING FAILED SUCCESS]
	Status string `json:"status,omitempty"`

	// The type of backup. TiDB Cloud only supports manual and auto backup. For more information, see [TiDB Cloud Documentation](https://docs.pingcap.com/tidbcloud/backup-and-restore#backup).
	// Example: MANUAL
	// Enum: [MANUAL AUTO]
	Type string `json:"type,omitempty"`
}

// Validate validates this get backup of cluster o k body
func (o *GetBackupOfClusterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackupOfClusterOKBody) validateCreateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.CreateTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("getBackupOfClusterOK"+"."+"create_timestamp", "body", "date-time", o.CreateTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var getBackupOfClusterOKBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","RUNNING","FAILED","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getBackupOfClusterOKBodyTypeStatusPropEnum = append(getBackupOfClusterOKBodyTypeStatusPropEnum, v)
	}
}

const (

	// GetBackupOfClusterOKBodyStatusPENDING captures enum value "PENDING"
	GetBackupOfClusterOKBodyStatusPENDING string = "PENDING"

	// GetBackupOfClusterOKBodyStatusRUNNING captures enum value "RUNNING"
	GetBackupOfClusterOKBodyStatusRUNNING string = "RUNNING"

	// GetBackupOfClusterOKBodyStatusFAILED captures enum value "FAILED"
	GetBackupOfClusterOKBodyStatusFAILED string = "FAILED"

	// GetBackupOfClusterOKBodyStatusSUCCESS captures enum value "SUCCESS"
	GetBackupOfClusterOKBodyStatusSUCCESS string = "SUCCESS"
)

// prop value enum
func (o *GetBackupOfClusterOKBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getBackupOfClusterOKBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetBackupOfClusterOKBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getBackupOfClusterOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

var getBackupOfClusterOKBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getBackupOfClusterOKBodyTypeTypePropEnum = append(getBackupOfClusterOKBodyTypeTypePropEnum, v)
	}
}

const (

	// GetBackupOfClusterOKBodyTypeMANUAL captures enum value "MANUAL"
	GetBackupOfClusterOKBodyTypeMANUAL string = "MANUAL"

	// GetBackupOfClusterOKBodyTypeAUTO captures enum value "AUTO"
	GetBackupOfClusterOKBodyTypeAUTO string = "AUTO"
)

// prop value enum
func (o *GetBackupOfClusterOKBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getBackupOfClusterOKBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetBackupOfClusterOKBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("getBackupOfClusterOK"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get backup of cluster o k body based on context it is used
func (o *GetBackupOfClusterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetBackupOfClusterTooManyRequestsBody get backup of cluster too many requests body
swagger:model GetBackupOfClusterTooManyRequestsBody
*/
type GetBackupOfClusterTooManyRequestsBody struct {

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

// Validate validates this get backup of cluster too many requests body
func (o *GetBackupOfClusterTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup of cluster too many requests body based on context it is used
func (o *GetBackupOfClusterTooManyRequestsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetBackupOfClusterTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackupOfClusterTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res GetBackupOfClusterTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
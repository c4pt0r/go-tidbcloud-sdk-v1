// Code generated by go-swagger; DO NOT EDIT.

package import_operations

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

// UpdateImportTaskReader is a Reader for the UpdateImportTask structure.
type UpdateImportTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateImportTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateImportTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateImportTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateImportTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateImportTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateImportTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateImportTaskTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateImportTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateImportTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateImportTaskOK creates a UpdateImportTaskOK with default headers values
func NewUpdateImportTaskOK() *UpdateImportTaskOK {
	return &UpdateImportTaskOK{}
}

/*
UpdateImportTaskOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateImportTaskOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update import task o k response has a 2xx status code
func (o *UpdateImportTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update import task o k response has a 3xx status code
func (o *UpdateImportTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task o k response has a 4xx status code
func (o *UpdateImportTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update import task o k response has a 5xx status code
func (o *UpdateImportTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update import task o k response a status code equal to that given
func (o *UpdateImportTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update import task o k response
func (o *UpdateImportTaskOK) Code() int {
	return 200
}

func (o *UpdateImportTaskOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskOK %s", 200, payload)
}

func (o *UpdateImportTaskOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskOK %s", 200, payload)
}

func (o *UpdateImportTaskOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateImportTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskBadRequest creates a UpdateImportTaskBadRequest with default headers values
func NewUpdateImportTaskBadRequest() *UpdateImportTaskBadRequest {
	return &UpdateImportTaskBadRequest{}
}

/*
UpdateImportTaskBadRequest describes a response with status code 400, with default header values.

A request field is invalid.
*/
type UpdateImportTaskBadRequest struct {
	Payload *UpdateImportTaskBadRequestBody
}

// IsSuccess returns true when this update import task bad request response has a 2xx status code
func (o *UpdateImportTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update import task bad request response has a 3xx status code
func (o *UpdateImportTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task bad request response has a 4xx status code
func (o *UpdateImportTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update import task bad request response has a 5xx status code
func (o *UpdateImportTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update import task bad request response a status code equal to that given
func (o *UpdateImportTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update import task bad request response
func (o *UpdateImportTaskBadRequest) Code() int {
	return 400
}

func (o *UpdateImportTaskBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskBadRequest %s", 400, payload)
}

func (o *UpdateImportTaskBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskBadRequest %s", 400, payload)
}

func (o *UpdateImportTaskBadRequest) GetPayload() *UpdateImportTaskBadRequestBody {
	return o.Payload
}

func (o *UpdateImportTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateImportTaskBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskUnauthorized creates a UpdateImportTaskUnauthorized with default headers values
func NewUpdateImportTaskUnauthorized() *UpdateImportTaskUnauthorized {
	return &UpdateImportTaskUnauthorized{}
}

/*
UpdateImportTaskUnauthorized describes a response with status code 401, with default header values.

The API key cannot be authenticated.
*/
type UpdateImportTaskUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this update import task unauthorized response has a 2xx status code
func (o *UpdateImportTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update import task unauthorized response has a 3xx status code
func (o *UpdateImportTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task unauthorized response has a 4xx status code
func (o *UpdateImportTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update import task unauthorized response has a 5xx status code
func (o *UpdateImportTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update import task unauthorized response a status code equal to that given
func (o *UpdateImportTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update import task unauthorized response
func (o *UpdateImportTaskUnauthorized) Code() int {
	return 401
}

func (o *UpdateImportTaskUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskUnauthorized %s", 401, payload)
}

func (o *UpdateImportTaskUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskUnauthorized %s", 401, payload)
}

func (o *UpdateImportTaskUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateImportTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskForbidden creates a UpdateImportTaskForbidden with default headers values
func NewUpdateImportTaskForbidden() *UpdateImportTaskForbidden {
	return &UpdateImportTaskForbidden{}
}

/*
UpdateImportTaskForbidden describes a response with status code 403, with default header values.

The API key does not have permission to access the resource.
*/
type UpdateImportTaskForbidden struct {
	Payload *UpdateImportTaskForbiddenBody
}

// IsSuccess returns true when this update import task forbidden response has a 2xx status code
func (o *UpdateImportTaskForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update import task forbidden response has a 3xx status code
func (o *UpdateImportTaskForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task forbidden response has a 4xx status code
func (o *UpdateImportTaskForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update import task forbidden response has a 5xx status code
func (o *UpdateImportTaskForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update import task forbidden response a status code equal to that given
func (o *UpdateImportTaskForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update import task forbidden response
func (o *UpdateImportTaskForbidden) Code() int {
	return 403
}

func (o *UpdateImportTaskForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskForbidden %s", 403, payload)
}

func (o *UpdateImportTaskForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskForbidden %s", 403, payload)
}

func (o *UpdateImportTaskForbidden) GetPayload() *UpdateImportTaskForbiddenBody {
	return o.Payload
}

func (o *UpdateImportTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateImportTaskForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskNotFound creates a UpdateImportTaskNotFound with default headers values
func NewUpdateImportTaskNotFound() *UpdateImportTaskNotFound {
	return &UpdateImportTaskNotFound{}
}

/*
UpdateImportTaskNotFound describes a response with status code 404, with default header values.

The requested resource does not exist.
*/
type UpdateImportTaskNotFound struct {
	Payload *UpdateImportTaskNotFoundBody
}

// IsSuccess returns true when this update import task not found response has a 2xx status code
func (o *UpdateImportTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update import task not found response has a 3xx status code
func (o *UpdateImportTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task not found response has a 4xx status code
func (o *UpdateImportTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update import task not found response has a 5xx status code
func (o *UpdateImportTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update import task not found response a status code equal to that given
func (o *UpdateImportTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update import task not found response
func (o *UpdateImportTaskNotFound) Code() int {
	return 404
}

func (o *UpdateImportTaskNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskNotFound %s", 404, payload)
}

func (o *UpdateImportTaskNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskNotFound %s", 404, payload)
}

func (o *UpdateImportTaskNotFound) GetPayload() *UpdateImportTaskNotFoundBody {
	return o.Payload
}

func (o *UpdateImportTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateImportTaskNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskTooManyRequests creates a UpdateImportTaskTooManyRequests with default headers values
func NewUpdateImportTaskTooManyRequests() *UpdateImportTaskTooManyRequests {
	return &UpdateImportTaskTooManyRequests{}
}

/*
UpdateImportTaskTooManyRequests describes a response with status code 429, with default header values.

You have exceed the rate limit.
*/
type UpdateImportTaskTooManyRequests struct {
	Payload *UpdateImportTaskTooManyRequestsBody
}

// IsSuccess returns true when this update import task too many requests response has a 2xx status code
func (o *UpdateImportTaskTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update import task too many requests response has a 3xx status code
func (o *UpdateImportTaskTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task too many requests response has a 4xx status code
func (o *UpdateImportTaskTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update import task too many requests response has a 5xx status code
func (o *UpdateImportTaskTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update import task too many requests response a status code equal to that given
func (o *UpdateImportTaskTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update import task too many requests response
func (o *UpdateImportTaskTooManyRequests) Code() int {
	return 429
}

func (o *UpdateImportTaskTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskTooManyRequests %s", 429, payload)
}

func (o *UpdateImportTaskTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskTooManyRequests %s", 429, payload)
}

func (o *UpdateImportTaskTooManyRequests) GetPayload() *UpdateImportTaskTooManyRequestsBody {
	return o.Payload
}

func (o *UpdateImportTaskTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateImportTaskTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskInternalServerError creates a UpdateImportTaskInternalServerError with default headers values
func NewUpdateImportTaskInternalServerError() *UpdateImportTaskInternalServerError {
	return &UpdateImportTaskInternalServerError{}
}

/*
UpdateImportTaskInternalServerError describes a response with status code 500, with default header values.

Server error.
*/
type UpdateImportTaskInternalServerError struct {
	Payload *UpdateImportTaskInternalServerErrorBody
}

// IsSuccess returns true when this update import task internal server error response has a 2xx status code
func (o *UpdateImportTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update import task internal server error response has a 3xx status code
func (o *UpdateImportTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update import task internal server error response has a 4xx status code
func (o *UpdateImportTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update import task internal server error response has a 5xx status code
func (o *UpdateImportTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update import task internal server error response a status code equal to that given
func (o *UpdateImportTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update import task internal server error response
func (o *UpdateImportTaskInternalServerError) Code() int {
	return 500
}

func (o *UpdateImportTaskInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskInternalServerError %s", 500, payload)
}

func (o *UpdateImportTaskInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] updateImportTaskInternalServerError %s", 500, payload)
}

func (o *UpdateImportTaskInternalServerError) GetPayload() *UpdateImportTaskInternalServerErrorBody {
	return o.Payload
}

func (o *UpdateImportTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateImportTaskInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImportTaskDefault creates a UpdateImportTaskDefault with default headers values
func NewUpdateImportTaskDefault(code int) *UpdateImportTaskDefault {
	return &UpdateImportTaskDefault{
		_statusCode: code,
	}
}

/*
UpdateImportTaskDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateImportTaskDefault struct {
	_statusCode int

	Payload *UpdateImportTaskDefaultBody
}

// IsSuccess returns true when this update import task default response has a 2xx status code
func (o *UpdateImportTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update import task default response has a 3xx status code
func (o *UpdateImportTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update import task default response has a 4xx status code
func (o *UpdateImportTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update import task default response has a 5xx status code
func (o *UpdateImportTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update import task default response a status code equal to that given
func (o *UpdateImportTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update import task default response
func (o *UpdateImportTaskDefault) Code() int {
	return o._statusCode
}

func (o *UpdateImportTaskDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] UpdateImportTask default %s", o._statusCode, payload)
}

func (o *UpdateImportTaskDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1beta/projects/{project_id}/clusters/{cluster_id}/imports/{import_id}][%d] UpdateImportTask default %s", o._statusCode, payload)
}

func (o *UpdateImportTaskDefault) GetPayload() *UpdateImportTaskDefaultBody {
	return o.Payload
}

func (o *UpdateImportTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateImportTaskDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateImportTaskBadRequestBody update import task bad request body
swagger:model UpdateImportTaskBadRequestBody
*/
type UpdateImportTaskBadRequestBody struct {

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

// Validate validates this update import task bad request body
func (o *UpdateImportTaskBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update import task bad request body based on context it is used
func (o *UpdateImportTaskBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskBody UpdateImportTaskReq
//
// UpdateImportTaskReq is the request to update an import task.
swagger:model UpdateImportTaskBody
*/
type UpdateImportTaskBody struct {

	// The action to apply to the import task.
	//
	// **Limitation:**
	// Currently, only `CANCEL` is supported when the import task is in the `PREPARING` or `IMPORTING` phase, meaning that the import task will be cancelled.
	// Example: CANCEL
	// Required: true
	// Enum: ["CANCEL"]
	Action *string `json:"action"`
}

// Validate validates this update import task body
func (o *UpdateImportTaskBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateImportTaskBodyTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CANCEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateImportTaskBodyTypeActionPropEnum = append(updateImportTaskBodyTypeActionPropEnum, v)
	}
}

const (

	// UpdateImportTaskBodyActionCANCEL captures enum value "CANCEL"
	UpdateImportTaskBodyActionCANCEL string = "CANCEL"
)

// prop value enum
func (o *UpdateImportTaskBody) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateImportTaskBodyTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateImportTaskBody) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"action", "body", o.Action); err != nil {
		return err
	}

	// value enum
	if err := o.validateActionEnum("body"+"."+"action", "body", *o.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update import task body based on context it is used
func (o *UpdateImportTaskBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskDefaultBody update import task default body
swagger:model UpdateImportTaskDefaultBody
*/
type UpdateImportTaskDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*UpdateImportTaskDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update import task default body
func (o *UpdateImportTaskDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateImportTaskDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UpdateImportTask default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateImportTask default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update import task default body based on the context it is used
func (o *UpdateImportTaskDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateImportTaskDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateImportTask default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateImportTask default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskDefaultBodyDetailsItems0 update import task default body details items0
swagger:model UpdateImportTaskDefaultBodyDetailsItems0
*/
type UpdateImportTaskDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// update import task default body details items0
	UpdateImportTaskDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *UpdateImportTaskDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv UpdateImportTaskDefaultBodyDetailsItems0

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
		o.UpdateImportTaskDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o UpdateImportTaskDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.UpdateImportTaskDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.UpdateImportTaskDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this update import task default body details items0
func (o *UpdateImportTaskDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update import task default body details items0 based on context it is used
func (o *UpdateImportTaskDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskForbiddenBody update import task forbidden body
swagger:model UpdateImportTaskForbiddenBody
*/
type UpdateImportTaskForbiddenBody struct {

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

// Validate validates this update import task forbidden body
func (o *UpdateImportTaskForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update import task forbidden body based on context it is used
func (o *UpdateImportTaskForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskForbiddenBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskInternalServerErrorBody update import task internal server error body
swagger:model UpdateImportTaskInternalServerErrorBody
*/
type UpdateImportTaskInternalServerErrorBody struct {

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

// Validate validates this update import task internal server error body
func (o *UpdateImportTaskInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update import task internal server error body based on context it is used
func (o *UpdateImportTaskInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskNotFoundBody update import task not found body
swagger:model UpdateImportTaskNotFoundBody
*/
type UpdateImportTaskNotFoundBody struct {

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

// Validate validates this update import task not found body
func (o *UpdateImportTaskNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update import task not found body based on context it is used
func (o *UpdateImportTaskNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateImportTaskTooManyRequestsBody update import task too many requests body
swagger:model UpdateImportTaskTooManyRequestsBody
*/
type UpdateImportTaskTooManyRequestsBody struct {

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

// Validate validates this update import task too many requests body
func (o *UpdateImportTaskTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update import task too many requests body based on context it is used
func (o *UpdateImportTaskTooManyRequestsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateImportTaskTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateImportTaskTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res UpdateImportTaskTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

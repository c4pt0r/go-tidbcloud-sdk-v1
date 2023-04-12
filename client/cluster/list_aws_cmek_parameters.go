// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListAwsCmekParams creates a new ListAwsCmekParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAwsCmekParams() *ListAwsCmekParams {
	return &ListAwsCmekParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAwsCmekParamsWithTimeout creates a new ListAwsCmekParams object
// with the ability to set a timeout on a request.
func NewListAwsCmekParamsWithTimeout(timeout time.Duration) *ListAwsCmekParams {
	return &ListAwsCmekParams{
		timeout: timeout,
	}
}

// NewListAwsCmekParamsWithContext creates a new ListAwsCmekParams object
// with the ability to set a context for a request.
func NewListAwsCmekParamsWithContext(ctx context.Context) *ListAwsCmekParams {
	return &ListAwsCmekParams{
		Context: ctx,
	}
}

// NewListAwsCmekParamsWithHTTPClient creates a new ListAwsCmekParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAwsCmekParamsWithHTTPClient(client *http.Client) *ListAwsCmekParams {
	return &ListAwsCmekParams{
		HTTPClient: client,
	}
}

/*
ListAwsCmekParams contains all the parameters to send to the API endpoint

	for the list aws cmek operation.

	Typically these are written to a http.Request.
*/
type ListAwsCmekParams struct {

	/* ProjectID.

	   The ID of the project. You can get the project ID from the response of [List all accessible projects](#tag/Project/operation/ListProjects).

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list aws cmek params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAwsCmekParams) WithDefaults() *ListAwsCmekParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list aws cmek params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAwsCmekParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list aws cmek params
func (o *ListAwsCmekParams) WithTimeout(timeout time.Duration) *ListAwsCmekParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list aws cmek params
func (o *ListAwsCmekParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list aws cmek params
func (o *ListAwsCmekParams) WithContext(ctx context.Context) *ListAwsCmekParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list aws cmek params
func (o *ListAwsCmekParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list aws cmek params
func (o *ListAwsCmekParams) WithHTTPClient(client *http.Client) *ListAwsCmekParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list aws cmek params
func (o *ListAwsCmekParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list aws cmek params
func (o *ListAwsCmekParams) WithProjectID(projectID string) *ListAwsCmekParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list aws cmek params
func (o *ListAwsCmekParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAwsCmekParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
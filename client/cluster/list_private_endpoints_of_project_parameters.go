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

// NewListPrivateEndpointsOfProjectParams creates a new ListPrivateEndpointsOfProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPrivateEndpointsOfProjectParams() *ListPrivateEndpointsOfProjectParams {
	return &ListPrivateEndpointsOfProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPrivateEndpointsOfProjectParamsWithTimeout creates a new ListPrivateEndpointsOfProjectParams object
// with the ability to set a timeout on a request.
func NewListPrivateEndpointsOfProjectParamsWithTimeout(timeout time.Duration) *ListPrivateEndpointsOfProjectParams {
	return &ListPrivateEndpointsOfProjectParams{
		timeout: timeout,
	}
}

// NewListPrivateEndpointsOfProjectParamsWithContext creates a new ListPrivateEndpointsOfProjectParams object
// with the ability to set a context for a request.
func NewListPrivateEndpointsOfProjectParamsWithContext(ctx context.Context) *ListPrivateEndpointsOfProjectParams {
	return &ListPrivateEndpointsOfProjectParams{
		Context: ctx,
	}
}

// NewListPrivateEndpointsOfProjectParamsWithHTTPClient creates a new ListPrivateEndpointsOfProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPrivateEndpointsOfProjectParamsWithHTTPClient(client *http.Client) *ListPrivateEndpointsOfProjectParams {
	return &ListPrivateEndpointsOfProjectParams{
		HTTPClient: client,
	}
}

/*
ListPrivateEndpointsOfProjectParams contains all the parameters to send to the API endpoint

	for the list private endpoints of project operation.

	Typically these are written to a http.Request.
*/
type ListPrivateEndpointsOfProjectParams struct {

	/* ProjectID.

	   The ID of the project. You can get the project ID from the response of [List all accessible projects](#tag/Project/operation/ListProjects).

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list private endpoints of project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPrivateEndpointsOfProjectParams) WithDefaults() *ListPrivateEndpointsOfProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list private endpoints of project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPrivateEndpointsOfProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) WithTimeout(timeout time.Duration) *ListPrivateEndpointsOfProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) WithContext(ctx context.Context) *ListPrivateEndpointsOfProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) WithHTTPClient(client *http.Client) *ListPrivateEndpointsOfProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) WithProjectID(projectID string) *ListPrivateEndpointsOfProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list private endpoints of project params
func (o *ListPrivateEndpointsOfProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPrivateEndpointsOfProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewUpdateAccesspointParams creates a new UpdateAccesspointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAccesspointParams() *UpdateAccesspointParams {
	return &UpdateAccesspointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccesspointParamsWithTimeout creates a new UpdateAccesspointParams object
// with the ability to set a timeout on a request.
func NewUpdateAccesspointParamsWithTimeout(timeout time.Duration) *UpdateAccesspointParams {
	return &UpdateAccesspointParams{
		timeout: timeout,
	}
}

// NewUpdateAccesspointParamsWithContext creates a new UpdateAccesspointParams object
// with the ability to set a context for a request.
func NewUpdateAccesspointParamsWithContext(ctx context.Context) *UpdateAccesspointParams {
	return &UpdateAccesspointParams{
		Context: ctx,
	}
}

// NewUpdateAccesspointParamsWithHTTPClient creates a new UpdateAccesspointParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAccesspointParamsWithHTTPClient(client *http.Client) *UpdateAccesspointParams {
	return &UpdateAccesspointParams{
		HTTPClient: client,
	}
}

/*
UpdateAccesspointParams contains all the parameters to send to the API endpoint

	for the update accesspoint operation.

	Typically these are written to a http.Request.
*/
type UpdateAccesspointParams struct {

	/* AccesspointID.

	   The ID of the accesspoint to be updated.

	   Format: uint64
	*/
	AccesspointID string

	// Body.
	Body UpdateAccesspointBody

	/* ClusterID.

	   The ID of the cluster.

	   Format: uint64
	*/
	ClusterID string

	/* ProjectID.

	   The ID of the project.

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update accesspoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccesspointParams) WithDefaults() *UpdateAccesspointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update accesspoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccesspointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update accesspoint params
func (o *UpdateAccesspointParams) WithTimeout(timeout time.Duration) *UpdateAccesspointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update accesspoint params
func (o *UpdateAccesspointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update accesspoint params
func (o *UpdateAccesspointParams) WithContext(ctx context.Context) *UpdateAccesspointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update accesspoint params
func (o *UpdateAccesspointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update accesspoint params
func (o *UpdateAccesspointParams) WithHTTPClient(client *http.Client) *UpdateAccesspointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update accesspoint params
func (o *UpdateAccesspointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccesspointID adds the accesspointID to the update accesspoint params
func (o *UpdateAccesspointParams) WithAccesspointID(accesspointID string) *UpdateAccesspointParams {
	o.SetAccesspointID(accesspointID)
	return o
}

// SetAccesspointID adds the accesspointId to the update accesspoint params
func (o *UpdateAccesspointParams) SetAccesspointID(accesspointID string) {
	o.AccesspointID = accesspointID
}

// WithBody adds the body to the update accesspoint params
func (o *UpdateAccesspointParams) WithBody(body UpdateAccesspointBody) *UpdateAccesspointParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update accesspoint params
func (o *UpdateAccesspointParams) SetBody(body UpdateAccesspointBody) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update accesspoint params
func (o *UpdateAccesspointParams) WithClusterID(clusterID string) *UpdateAccesspointParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update accesspoint params
func (o *UpdateAccesspointParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the update accesspoint params
func (o *UpdateAccesspointParams) WithProjectID(projectID string) *UpdateAccesspointParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update accesspoint params
func (o *UpdateAccesspointParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccesspointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accesspoint_id
	if err := r.SetPathParam("accesspoint_id", o.AccesspointID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

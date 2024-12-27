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

// NewDeletePrivateEndpointParams creates a new DeletePrivateEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePrivateEndpointParams() *DeletePrivateEndpointParams {
	return &DeletePrivateEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateEndpointParamsWithTimeout creates a new DeletePrivateEndpointParams object
// with the ability to set a timeout on a request.
func NewDeletePrivateEndpointParamsWithTimeout(timeout time.Duration) *DeletePrivateEndpointParams {
	return &DeletePrivateEndpointParams{
		timeout: timeout,
	}
}

// NewDeletePrivateEndpointParamsWithContext creates a new DeletePrivateEndpointParams object
// with the ability to set a context for a request.
func NewDeletePrivateEndpointParamsWithContext(ctx context.Context) *DeletePrivateEndpointParams {
	return &DeletePrivateEndpointParams{
		Context: ctx,
	}
}

// NewDeletePrivateEndpointParamsWithHTTPClient creates a new DeletePrivateEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePrivateEndpointParamsWithHTTPClient(client *http.Client) *DeletePrivateEndpointParams {
	return &DeletePrivateEndpointParams{
		HTTPClient: client,
	}
}

/*
DeletePrivateEndpointParams contains all the parameters to send to the API endpoint

	for the delete private endpoint operation.

	Typically these are written to a http.Request.
*/
type DeletePrivateEndpointParams struct {

	/* ClusterID.

	   The ID of the cluster.

	   Format: uint64
	*/
	ClusterID string

	/* EndpointID.

	   The ID of the private endpoint to be deleted. You can get the ID from the `endpoints.id` field in the response of [List all private endpoints in a project](#tag/Cluster/operation/ListPrivateEndpoints).

	   Format: uint64
	*/
	EndpointID string

	/* ProjectID.

	   The ID of the project. You can get the project ID from the response of [List all accessible projects](#tag/Project/operation/ListProjects).

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete private endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateEndpointParams) WithDefaults() *DeletePrivateEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete private endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete private endpoint params
func (o *DeletePrivateEndpointParams) WithTimeout(timeout time.Duration) *DeletePrivateEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private endpoint params
func (o *DeletePrivateEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private endpoint params
func (o *DeletePrivateEndpointParams) WithContext(ctx context.Context) *DeletePrivateEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private endpoint params
func (o *DeletePrivateEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private endpoint params
func (o *DeletePrivateEndpointParams) WithHTTPClient(client *http.Client) *DeletePrivateEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private endpoint params
func (o *DeletePrivateEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete private endpoint params
func (o *DeletePrivateEndpointParams) WithClusterID(clusterID string) *DeletePrivateEndpointParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete private endpoint params
func (o *DeletePrivateEndpointParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEndpointID adds the endpointID to the delete private endpoint params
func (o *DeletePrivateEndpointParams) WithEndpointID(endpointID string) *DeletePrivateEndpointParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the delete private endpoint params
func (o *DeletePrivateEndpointParams) SetEndpointID(endpointID string) {
	o.EndpointID = endpointID
}

// WithProjectID adds the projectID to the delete private endpoint params
func (o *DeletePrivateEndpointParams) WithProjectID(projectID string) *DeletePrivateEndpointParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete private endpoint params
func (o *DeletePrivateEndpointParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param endpoint_id
	if err := r.SetPathParam("endpoint_id", o.EndpointID); err != nil {
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

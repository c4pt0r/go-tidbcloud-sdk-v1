// Code generated by go-swagger; DO NOT EDIT.

package import_operations

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

// NewCreateImportTaskParams creates a new CreateImportTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateImportTaskParams() *CreateImportTaskParams {
	return &CreateImportTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateImportTaskParamsWithTimeout creates a new CreateImportTaskParams object
// with the ability to set a timeout on a request.
func NewCreateImportTaskParamsWithTimeout(timeout time.Duration) *CreateImportTaskParams {
	return &CreateImportTaskParams{
		timeout: timeout,
	}
}

// NewCreateImportTaskParamsWithContext creates a new CreateImportTaskParams object
// with the ability to set a context for a request.
func NewCreateImportTaskParamsWithContext(ctx context.Context) *CreateImportTaskParams {
	return &CreateImportTaskParams{
		Context: ctx,
	}
}

// NewCreateImportTaskParamsWithHTTPClient creates a new CreateImportTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateImportTaskParamsWithHTTPClient(client *http.Client) *CreateImportTaskParams {
	return &CreateImportTaskParams{
		HTTPClient: client,
	}
}

/*
CreateImportTaskParams contains all the parameters to send to the API endpoint

	for the create import task operation.

	Typically these are written to a http.Request.
*/
type CreateImportTaskParams struct {

	// Body.
	Body CreateImportTaskBody

	/* ClusterID.

	   The ID of your cluster that you want to start an import job. You can get the cluster ID from the response of [List all clusters in a project](#tag/Cluster/operation/ListClustersOfProject).

	   Format: uint64
	*/
	ClusterID string

	/* ProjectID.

	   The ID of your project. You can get the project ID from the response of [List all accessible projects](#tag/Project/operation/ListProjects).

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create import task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateImportTaskParams) WithDefaults() *CreateImportTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create import task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateImportTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create import task params
func (o *CreateImportTaskParams) WithTimeout(timeout time.Duration) *CreateImportTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create import task params
func (o *CreateImportTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create import task params
func (o *CreateImportTaskParams) WithContext(ctx context.Context) *CreateImportTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create import task params
func (o *CreateImportTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create import task params
func (o *CreateImportTaskParams) WithHTTPClient(client *http.Client) *CreateImportTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create import task params
func (o *CreateImportTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create import task params
func (o *CreateImportTaskParams) WithBody(body CreateImportTaskBody) *CreateImportTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create import task params
func (o *CreateImportTaskParams) SetBody(body CreateImportTaskBody) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create import task params
func (o *CreateImportTaskParams) WithClusterID(clusterID string) *CreateImportTaskParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create import task params
func (o *CreateImportTaskParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create import task params
func (o *CreateImportTaskParams) WithProjectID(projectID string) *CreateImportTaskParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create import task params
func (o *CreateImportTaskParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateImportTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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

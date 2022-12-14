// Code generated by go-swagger; DO NOT EDIT.

package backup

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
	"github.com/go-openapi/swag"
)

// NewListBackUpOfClusterParams creates a new ListBackUpOfClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBackUpOfClusterParams() *ListBackUpOfClusterParams {
	return &ListBackUpOfClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBackUpOfClusterParamsWithTimeout creates a new ListBackUpOfClusterParams object
// with the ability to set a timeout on a request.
func NewListBackUpOfClusterParamsWithTimeout(timeout time.Duration) *ListBackUpOfClusterParams {
	return &ListBackUpOfClusterParams{
		timeout: timeout,
	}
}

// NewListBackUpOfClusterParamsWithContext creates a new ListBackUpOfClusterParams object
// with the ability to set a context for a request.
func NewListBackUpOfClusterParamsWithContext(ctx context.Context) *ListBackUpOfClusterParams {
	return &ListBackUpOfClusterParams{
		Context: ctx,
	}
}

// NewListBackUpOfClusterParamsWithHTTPClient creates a new ListBackUpOfClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBackUpOfClusterParamsWithHTTPClient(client *http.Client) *ListBackUpOfClusterParams {
	return &ListBackUpOfClusterParams{
		HTTPClient: client,
	}
}

/*
ListBackUpOfClusterParams contains all the parameters to send to the API endpoint

	for the list back up of cluster operation.

	Typically these are written to a http.Request.
*/
type ListBackUpOfClusterParams struct {

	/* ClusterID.

	   The ID of your cluster. You can get the cluster ID from the response of [Get all clusters in a project](#tag/Cluster/operation/ListClustersOfProject).

	   Format: uint64
	*/
	ClusterID string

	/* Page.

	   The number of pages.

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* PageSize.

	   The size of a page.

	   Format: int32
	   Default: 10
	*/
	PageSize *int32

	/* ProjectID.

	   The ID of your project. You can get the project ID from the response of [List all accessible projects.](#tag/Project/operation/ListProjects).

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list back up of cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBackUpOfClusterParams) WithDefaults() *ListBackUpOfClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list back up of cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBackUpOfClusterParams) SetDefaults() {
	var (
		pageDefault = int32(1)

		pageSizeDefault = int32(10)
	)

	val := ListBackUpOfClusterParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithTimeout(timeout time.Duration) *ListBackUpOfClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithContext(ctx context.Context) *ListBackUpOfClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithHTTPClient(client *http.Client) *ListBackUpOfClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithClusterID(clusterID string) *ListBackUpOfClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithPage adds the page to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithPage(page *int32) *ListBackUpOfClusterParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithPageSize(pageSize *int32) *ListBackUpOfClusterParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the list back up of cluster params
func (o *ListBackUpOfClusterParams) WithProjectID(projectID string) *ListBackUpOfClusterParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list back up of cluster params
func (o *ListBackUpOfClusterParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListBackUpOfClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
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

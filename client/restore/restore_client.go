// Code generated by go-swagger; DO NOT EDIT.

package restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new restore API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for restore API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRestoreTask(params *CreateRestoreTaskParams, opts ...ClientOption) (*CreateRestoreTaskOK, error)

	GetRestoreTask(params *GetRestoreTaskParams, opts ...ClientOption) (*GetRestoreTaskOK, error)

	ListRestoreTasks(params *ListRestoreTasksParams, opts ...ClientOption) (*ListRestoreTasksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateRestoreTask creates a restore task

	You can use this endpoint to restore data from a previously created backup file to a new cluster. In this endpoint, you must specify the configuration of the new cluster you want to restore data to.

**Limitations:**

- For Dedicated Tier, you can only restore data from a smaller node size to a larger node size.

- You cannot restore data from a Dedicated Tier cluster to a Developer Tier cluster.

For Developer Tier clusters, you cannot create restore tasks via API.
*/
func (a *Client) CreateRestoreTask(params *CreateRestoreTaskParams, opts ...ClientOption) (*CreateRestoreTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRestoreTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRestoreTask",
		Method:             "POST",
		PathPattern:        "/api/v1beta/projects/{project_id}/restores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRestoreTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRestoreTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRestoreTaskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetRestoreTask gets a restore task

For Developer Tier clusters, you cannot manage restore tasks via API.
*/
func (a *Client) GetRestoreTask(params *GetRestoreTaskParams, opts ...ClientOption) (*GetRestoreTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestoreTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRestoreTask",
		Method:             "GET",
		PathPattern:        "/api/v1beta/projects/{project_id}/restores/{restore_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestoreTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRestoreTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRestoreTaskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	ListRestoreTasks lists the restore tasks in a project

For Developer Tier clusters, you cannot create or manage restore tasks via API.
*/
func (a *Client) ListRestoreTasks(params *ListRestoreTasksParams, opts ...ClientOption) (*ListRestoreTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRestoreTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRestoreTasks",
		Method:             "GET",
		PathPattern:        "/api/v1beta/projects/{project_id}/restores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRestoreTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRestoreTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRestoreTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
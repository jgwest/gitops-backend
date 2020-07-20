// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMixin8(params *CreateMixin8Params) (*CreateMixin8OK, error)

	DeleteMixin8(params *DeleteMixin8Params) (*DeleteMixin8OK, error)

	DeleteResource(params *DeleteResourceParams) (*DeleteResourceOK, error)

	GetApplicationSyncWindows(params *GetApplicationSyncWindowsParams) (*GetApplicationSyncWindowsOK, error)

	GetManifests(params *GetManifestsParams) (*GetManifestsOK, error)

	GetMixin8(params *GetMixin8Params) (*GetMixin8OK, error)

	GetResource(params *GetResourceParams) (*GetResourceOK, error)

	ListMixin8(params *ListMixin8Params) (*ListMixin8OK, error)

	ListResourceActions(params *ListResourceActionsParams) (*ListResourceActionsOK, error)

	ListResourceEvents(params *ListResourceEventsParams) (*ListResourceEventsOK, error)

	ManagedResources(params *ManagedResourcesParams) (*ManagedResourcesOK, error)

	Patch(params *PatchParams) (*PatchOK, error)

	PatchResource(params *PatchResourceParams) (*PatchResourceOK, error)

	PodLogs(params *PodLogsParams) (*PodLogsOK, error)

	ResourceTree(params *ResourceTreeParams) (*ResourceTreeOK, error)

	RevisionMetadata(params *RevisionMetadataParams) (*RevisionMetadataOK, error)

	Rollback(params *RollbackParams) (*RollbackOK, error)

	RunResourceAction(params *RunResourceActionParams) (*RunResourceActionOK, error)

	Sync(params *SyncParams) (*SyncOK, error)

	TerminateOperation(params *TerminateOperationParams) (*TerminateOperationOK, error)

	UpdateMixin8(params *UpdateMixin8Params) (*UpdateMixin8OK, error)

	UpdateSpec(params *UpdateSpecParams) (*UpdateSpecOK, error)

	Watch(params *WatchParams) (*WatchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateMixin8 creates creates an application
*/
func (a *Client) CreateMixin8(params *CreateMixin8Params) (*CreateMixin8OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMixin8Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateMixin8",
		Method:             "POST",
		PathPattern:        "/api/v1/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMixin8Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMixin8OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateMixin8: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteMixin8 deletes deletes an application
*/
func (a *Client) DeleteMixin8(params *DeleteMixin8Params) (*DeleteMixin8OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMixin8Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMixin8",
		Method:             "DELETE",
		PathPattern:        "/api/v1/applications/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMixin8Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteMixin8OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteMixin8: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteResource deletes resource deletes a single application resource
*/
func (a *Client) DeleteResource(params *DeleteResourceParams) (*DeleteResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteResource",
		Method:             "DELETE",
		PathPattern:        "/api/v1/applications/{name}/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplicationSyncWindows gets returns an application by name
*/
func (a *Client) GetApplicationSyncWindows(params *GetApplicationSyncWindowsParams) (*GetApplicationSyncWindowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationSyncWindowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetApplicationSyncWindows",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/syncwindows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationSyncWindowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationSyncWindowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetApplicationSyncWindows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetManifests gets manifests returns application manifests
*/
func (a *Client) GetManifests(params *GetManifestsParams) (*GetManifestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManifestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetManifests",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetManifestsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetManifestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetManifests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMixin8 gets returns an application by name
*/
func (a *Client) GetMixin8(params *GetMixin8Params) (*GetMixin8OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMixin8Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMixin8",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMixin8Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMixin8OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMixin8: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetResource gets resource returns single application resource
*/
func (a *Client) GetResource(params *GetResourceParams) (*GetResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetResource",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListMixin8 lists returns list of applications
*/
func (a *Client) ListMixin8(params *ListMixin8Params) (*ListMixin8OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMixin8Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListMixin8",
		Method:             "GET",
		PathPattern:        "/api/v1/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListMixin8Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMixin8OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListMixin8: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListResourceActions list resource actions API
*/
func (a *Client) ListResourceActions(params *ListResourceActionsParams) (*ListResourceActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListResourceActionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListResourceActions",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/resource/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListResourceActionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListResourceActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListResourceActions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListResourceEvents lists resource events returns a list of event resources
*/
func (a *Client) ListResourceEvents(params *ListResourceEventsParams) (*ListResourceEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListResourceEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListResourceEvents",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListResourceEventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListResourceEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListResourceEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ManagedResources managed resources API
*/
func (a *Client) ManagedResources(params *ManagedResourcesParams) (*ManagedResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagedResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ManagedResources",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{applicationName}/managed-resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagedResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ManagedResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ManagedResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Patch patches patch an application
*/
func (a *Client) Patch(params *PatchParams) (*PatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Patch",
		Method:             "PATCH",
		PathPattern:        "/api/v1/applications/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Patch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchResource patches resource patch single application resource
*/
func (a *Client) PatchResource(params *PatchResourceParams) (*PatchResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchResource",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PodLogs pods logs returns stream of log entries for the specified pod pod
*/
func (a *Client) PodLogs(params *PodLogsParams) (*PodLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PodLogs",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/pods/{podName}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodLogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PodLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResourceTree resource tree API
*/
func (a *Client) ResourceTree(params *ResourceTreeParams) (*ResourceTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceTreeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ResourceTree",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{applicationName}/resource-tree",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResourceTreeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResourceTreeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResourceTree: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RevisionMetadata gets the meta data author date tags message for a specific revision of the application
*/
func (a *Client) RevisionMetadata(params *RevisionMetadataParams) (*RevisionMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevisionMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RevisionMetadata",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/revisions/{revision}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RevisionMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevisionMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RevisionMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Rollback rollbacks syncs an application to its target state
*/
func (a *Client) Rollback(params *RollbackParams) (*RollbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRollbackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Rollback",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/rollback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RollbackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RollbackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Rollback: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RunResourceAction run resource action API
*/
func (a *Client) RunResourceAction(params *RunResourceActionParams) (*RunResourceActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunResourceActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunResourceAction",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/resource/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunResourceActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunResourceActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RunResourceAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Sync syncs syncs an application to its target state
*/
func (a *Client) Sync(params *SyncParams) (*SyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Sync",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Sync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TerminateOperation terminates operation terminates the currently running operation
*/
func (a *Client) TerminateOperation(params *TerminateOperationParams) (*TerminateOperationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateOperationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TerminateOperation",
		Method:             "DELETE",
		PathPattern:        "/api/v1/applications/{name}/operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TerminateOperationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TerminateOperationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TerminateOperation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateMixin8 updates updates an application
*/
func (a *Client) UpdateMixin8(params *UpdateMixin8Params) (*UpdateMixin8OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMixin8Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateMixin8",
		Method:             "PUT",
		PathPattern:        "/api/v1/applications/{application.metadata.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateMixin8Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMixin8OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateMixin8: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSpec updates spec updates an application spec
*/
func (a *Client) UpdateSpec(params *UpdateSpecParams) (*UpdateSpecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSpecParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateSpec",
		Method:             "PUT",
		PathPattern:        "/api/v1/applications/{name}/spec",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateSpecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSpecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSpec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Watch watches returns stream of application change events
*/
func (a *Client) Watch(params *WatchParams) (*WatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Watch",
		Method:             "GET",
		PathPattern:        "/api/v1/stream/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Watch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
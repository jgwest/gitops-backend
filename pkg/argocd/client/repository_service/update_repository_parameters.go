// Code generated by go-swagger; DO NOT EDIT.

package repository_service

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

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// NewUpdateRepositoryParams creates a new UpdateRepositoryParams object
// with the default values initialized.
func NewUpdateRepositoryParams() *UpdateRepositoryParams {
	var ()
	return &UpdateRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepositoryParamsWithTimeout creates a new UpdateRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRepositoryParamsWithTimeout(timeout time.Duration) *UpdateRepositoryParams {
	var ()
	return &UpdateRepositoryParams{

		timeout: timeout,
	}
}

// NewUpdateRepositoryParamsWithContext creates a new UpdateRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRepositoryParamsWithContext(ctx context.Context) *UpdateRepositoryParams {
	var ()
	return &UpdateRepositoryParams{

		Context: ctx,
	}
}

// NewUpdateRepositoryParamsWithHTTPClient creates a new UpdateRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRepositoryParamsWithHTTPClient(client *http.Client) *UpdateRepositoryParams {
	var ()
	return &UpdateRepositoryParams{
		HTTPClient: client,
	}
}

/*UpdateRepositoryParams contains all the parameters to send to the API endpoint
for the update repository operation typically these are written to a http.Request
*/
type UpdateRepositoryParams struct {

	/*Body*/
	Body *models.V1alpha1Repository
	/*RepoRepo*/
	RepoRepo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update repository params
func (o *UpdateRepositoryParams) WithTimeout(timeout time.Duration) *UpdateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository params
func (o *UpdateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository params
func (o *UpdateRepositoryParams) WithContext(ctx context.Context) *UpdateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository params
func (o *UpdateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository params
func (o *UpdateRepositoryParams) WithHTTPClient(client *http.Client) *UpdateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository params
func (o *UpdateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository params
func (o *UpdateRepositoryParams) WithBody(body *models.V1alpha1Repository) *UpdateRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository params
func (o *UpdateRepositoryParams) SetBody(body *models.V1alpha1Repository) {
	o.Body = body
}

// WithRepoRepo adds the repoRepo to the update repository params
func (o *UpdateRepositoryParams) WithRepoRepo(repoRepo string) *UpdateRepositoryParams {
	o.SetRepoRepo(repoRepo)
	return o
}

// SetRepoRepo adds the repoRepo to the update repository params
func (o *UpdateRepositoryParams) SetRepoRepo(repoRepo string) {
	o.RepoRepo = repoRepo
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repo.repo
	if err := r.SetPathParam("repo.repo", o.RepoRepo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package project

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

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewCreateClusterRoleParams creates a new CreateClusterRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterRoleParams() *CreateClusterRoleParams {
	return &CreateClusterRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterRoleParamsWithTimeout creates a new CreateClusterRoleParams object
// with the ability to set a timeout on a request.
func NewCreateClusterRoleParamsWithTimeout(timeout time.Duration) *CreateClusterRoleParams {
	return &CreateClusterRoleParams{
		timeout: timeout,
	}
}

// NewCreateClusterRoleParamsWithContext creates a new CreateClusterRoleParams object
// with the ability to set a context for a request.
func NewCreateClusterRoleParamsWithContext(ctx context.Context) *CreateClusterRoleParams {
	return &CreateClusterRoleParams{
		Context: ctx,
	}
}

// NewCreateClusterRoleParamsWithHTTPClient creates a new CreateClusterRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterRoleParamsWithHTTPClient(client *http.Client) *CreateClusterRoleParams {
	return &CreateClusterRoleParams{
		HTTPClient: client,
	}
}

/* CreateClusterRoleParams contains all the parameters to send to the API endpoint
   for the create cluster role operation.

   Typically these are written to a http.Request.
*/
type CreateClusterRoleParams struct {

	// Body.
	Body *models.ClusterRole

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterRoleParams) WithDefaults() *CreateClusterRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster role params
func (o *CreateClusterRoleParams) WithTimeout(timeout time.Duration) *CreateClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster role params
func (o *CreateClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster role params
func (o *CreateClusterRoleParams) WithContext(ctx context.Context) *CreateClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster role params
func (o *CreateClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster role params
func (o *CreateClusterRoleParams) WithHTTPClient(client *http.Client) *CreateClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster role params
func (o *CreateClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster role params
func (o *CreateClusterRoleParams) WithBody(body *models.ClusterRole) *CreateClusterRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster role params
func (o *CreateClusterRoleParams) SetBody(body *models.ClusterRole) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create cluster role params
func (o *CreateClusterRoleParams) WithClusterID(clusterID string) *CreateClusterRoleParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create cluster role params
func (o *CreateClusterRoleParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the create cluster role params
func (o *CreateClusterRoleParams) WithDC(dc string) *CreateClusterRoleParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the create cluster role params
func (o *CreateClusterRoleParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the create cluster role params
func (o *CreateClusterRoleParams) WithProjectID(projectID string) *CreateClusterRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create cluster role params
func (o *CreateClusterRoleParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
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

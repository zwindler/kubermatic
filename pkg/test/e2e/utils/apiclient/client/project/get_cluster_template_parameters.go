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
)

// NewGetClusterTemplateParams creates a new GetClusterTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterTemplateParams() *GetClusterTemplateParams {
	return &GetClusterTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTemplateParamsWithTimeout creates a new GetClusterTemplateParams object
// with the ability to set a timeout on a request.
func NewGetClusterTemplateParamsWithTimeout(timeout time.Duration) *GetClusterTemplateParams {
	return &GetClusterTemplateParams{
		timeout: timeout,
	}
}

// NewGetClusterTemplateParamsWithContext creates a new GetClusterTemplateParams object
// with the ability to set a context for a request.
func NewGetClusterTemplateParamsWithContext(ctx context.Context) *GetClusterTemplateParams {
	return &GetClusterTemplateParams{
		Context: ctx,
	}
}

// NewGetClusterTemplateParamsWithHTTPClient creates a new GetClusterTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterTemplateParamsWithHTTPClient(client *http.Client) *GetClusterTemplateParams {
	return &GetClusterTemplateParams{
		HTTPClient: client,
	}
}

/* GetClusterTemplateParams contains all the parameters to send to the API endpoint
   for the get cluster template operation.

   Typically these are written to a http.Request.
*/
type GetClusterTemplateParams struct {

	// ProjectID.
	ProjectID string

	// TemplateID.
	ClusterTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTemplateParams) WithDefaults() *GetClusterTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster template params
func (o *GetClusterTemplateParams) WithTimeout(timeout time.Duration) *GetClusterTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster template params
func (o *GetClusterTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster template params
func (o *GetClusterTemplateParams) WithContext(ctx context.Context) *GetClusterTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster template params
func (o *GetClusterTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster template params
func (o *GetClusterTemplateParams) WithHTTPClient(client *http.Client) *GetClusterTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster template params
func (o *GetClusterTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get cluster template params
func (o *GetClusterTemplateParams) WithProjectID(projectID string) *GetClusterTemplateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get cluster template params
func (o *GetClusterTemplateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithClusterTemplateID adds the templateID to the get cluster template params
func (o *GetClusterTemplateParams) WithClusterTemplateID(templateID string) *GetClusterTemplateParams {
	o.SetClusterTemplateID(templateID)
	return o
}

// SetClusterTemplateID adds the templateId to the get cluster template params
func (o *GetClusterTemplateParams) SetClusterTemplateID(templateID string) {
	o.ClusterTemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.ClusterTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

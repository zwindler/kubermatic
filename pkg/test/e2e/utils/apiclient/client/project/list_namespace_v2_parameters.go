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

// NewListNamespaceV2Params creates a new ListNamespaceV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNamespaceV2Params() *ListNamespaceV2Params {
	return &ListNamespaceV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNamespaceV2ParamsWithTimeout creates a new ListNamespaceV2Params object
// with the ability to set a timeout on a request.
func NewListNamespaceV2ParamsWithTimeout(timeout time.Duration) *ListNamespaceV2Params {
	return &ListNamespaceV2Params{
		timeout: timeout,
	}
}

// NewListNamespaceV2ParamsWithContext creates a new ListNamespaceV2Params object
// with the ability to set a context for a request.
func NewListNamespaceV2ParamsWithContext(ctx context.Context) *ListNamespaceV2Params {
	return &ListNamespaceV2Params{
		Context: ctx,
	}
}

// NewListNamespaceV2ParamsWithHTTPClient creates a new ListNamespaceV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListNamespaceV2ParamsWithHTTPClient(client *http.Client) *ListNamespaceV2Params {
	return &ListNamespaceV2Params{
		HTTPClient: client,
	}
}

/* ListNamespaceV2Params contains all the parameters to send to the API endpoint
   for the list namespace v2 operation.

   Typically these are written to a http.Request.
*/
type ListNamespaceV2Params struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list namespace v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNamespaceV2Params) WithDefaults() *ListNamespaceV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list namespace v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNamespaceV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list namespace v2 params
func (o *ListNamespaceV2Params) WithTimeout(timeout time.Duration) *ListNamespaceV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list namespace v2 params
func (o *ListNamespaceV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list namespace v2 params
func (o *ListNamespaceV2Params) WithContext(ctx context.Context) *ListNamespaceV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list namespace v2 params
func (o *ListNamespaceV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list namespace v2 params
func (o *ListNamespaceV2Params) WithHTTPClient(client *http.Client) *ListNamespaceV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list namespace v2 params
func (o *ListNamespaceV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list namespace v2 params
func (o *ListNamespaceV2Params) WithClusterID(clusterID string) *ListNamespaceV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list namespace v2 params
func (o *ListNamespaceV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list namespace v2 params
func (o *ListNamespaceV2Params) WithProjectID(projectID string) *ListNamespaceV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list namespace v2 params
func (o *ListNamespaceV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListNamespaceV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

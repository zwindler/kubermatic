// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

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

// NewGetRuleGroupParams creates a new GetRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuleGroupParams() *GetRuleGroupParams {
	return &GetRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuleGroupParamsWithTimeout creates a new GetRuleGroupParams object
// with the ability to set a timeout on a request.
func NewGetRuleGroupParamsWithTimeout(timeout time.Duration) *GetRuleGroupParams {
	return &GetRuleGroupParams{
		timeout: timeout,
	}
}

// NewGetRuleGroupParamsWithContext creates a new GetRuleGroupParams object
// with the ability to set a context for a request.
func NewGetRuleGroupParamsWithContext(ctx context.Context) *GetRuleGroupParams {
	return &GetRuleGroupParams{
		Context: ctx,
	}
}

// NewGetRuleGroupParamsWithHTTPClient creates a new GetRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuleGroupParamsWithHTTPClient(client *http.Client) *GetRuleGroupParams {
	return &GetRuleGroupParams{
		HTTPClient: client,
	}
}

/* GetRuleGroupParams contains all the parameters to send to the API endpoint
   for the get rule group operation.

   Typically these are written to a http.Request.
*/
type GetRuleGroupParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	// RulegroupID.
	RuleGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuleGroupParams) WithDefaults() *GetRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get rule group params
func (o *GetRuleGroupParams) WithTimeout(timeout time.Duration) *GetRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rule group params
func (o *GetRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rule group params
func (o *GetRuleGroupParams) WithContext(ctx context.Context) *GetRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rule group params
func (o *GetRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rule group params
func (o *GetRuleGroupParams) WithHTTPClient(client *http.Client) *GetRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rule group params
func (o *GetRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get rule group params
func (o *GetRuleGroupParams) WithClusterID(clusterID string) *GetRuleGroupParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get rule group params
func (o *GetRuleGroupParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get rule group params
func (o *GetRuleGroupParams) WithProjectID(projectID string) *GetRuleGroupParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get rule group params
func (o *GetRuleGroupParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRuleGroupID adds the rulegroupID to the get rule group params
func (o *GetRuleGroupParams) WithRuleGroupID(rulegroupID string) *GetRuleGroupParams {
	o.SetRuleGroupID(rulegroupID)
	return o
}

// SetRuleGroupID adds the rulegroupId to the get rule group params
func (o *GetRuleGroupParams) SetRuleGroupID(rulegroupID string) {
	o.RuleGroupID = rulegroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param rulegroup_id
	if err := r.SetPathParam("rulegroup_id", o.RuleGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

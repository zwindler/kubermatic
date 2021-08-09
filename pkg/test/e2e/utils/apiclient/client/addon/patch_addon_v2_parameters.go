// Code generated by go-swagger; DO NOT EDIT.

package addon

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

// NewPatchAddonV2Params creates a new PatchAddonV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAddonV2Params() *PatchAddonV2Params {
	return &PatchAddonV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAddonV2ParamsWithTimeout creates a new PatchAddonV2Params object
// with the ability to set a timeout on a request.
func NewPatchAddonV2ParamsWithTimeout(timeout time.Duration) *PatchAddonV2Params {
	return &PatchAddonV2Params{
		timeout: timeout,
	}
}

// NewPatchAddonV2ParamsWithContext creates a new PatchAddonV2Params object
// with the ability to set a context for a request.
func NewPatchAddonV2ParamsWithContext(ctx context.Context) *PatchAddonV2Params {
	return &PatchAddonV2Params{
		Context: ctx,
	}
}

// NewPatchAddonV2ParamsWithHTTPClient creates a new PatchAddonV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAddonV2ParamsWithHTTPClient(client *http.Client) *PatchAddonV2Params {
	return &PatchAddonV2Params{
		HTTPClient: client,
	}
}

/* PatchAddonV2Params contains all the parameters to send to the API endpoint
   for the patch addon v2 operation.

   Typically these are written to a http.Request.
*/
type PatchAddonV2Params struct {

	// Body.
	Body *models.Addon

	// AddonID.
	AddonID string

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch addon v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAddonV2Params) WithDefaults() *PatchAddonV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch addon v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAddonV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch addon v2 params
func (o *PatchAddonV2Params) WithTimeout(timeout time.Duration) *PatchAddonV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch addon v2 params
func (o *PatchAddonV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch addon v2 params
func (o *PatchAddonV2Params) WithContext(ctx context.Context) *PatchAddonV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch addon v2 params
func (o *PatchAddonV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch addon v2 params
func (o *PatchAddonV2Params) WithHTTPClient(client *http.Client) *PatchAddonV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch addon v2 params
func (o *PatchAddonV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch addon v2 params
func (o *PatchAddonV2Params) WithBody(body *models.Addon) *PatchAddonV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch addon v2 params
func (o *PatchAddonV2Params) SetBody(body *models.Addon) {
	o.Body = body
}

// WithAddonID adds the addonID to the patch addon v2 params
func (o *PatchAddonV2Params) WithAddonID(addonID string) *PatchAddonV2Params {
	o.SetAddonID(addonID)
	return o
}

// SetAddonID adds the addonId to the patch addon v2 params
func (o *PatchAddonV2Params) SetAddonID(addonID string) {
	o.AddonID = addonID
}

// WithClusterID adds the clusterID to the patch addon v2 params
func (o *PatchAddonV2Params) WithClusterID(clusterID string) *PatchAddonV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the patch addon v2 params
func (o *PatchAddonV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the patch addon v2 params
func (o *PatchAddonV2Params) WithProjectID(projectID string) *PatchAddonV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch addon v2 params
func (o *PatchAddonV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAddonV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param addon_id
	if err := r.SetPathParam("addon_id", o.AddonID); err != nil {
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

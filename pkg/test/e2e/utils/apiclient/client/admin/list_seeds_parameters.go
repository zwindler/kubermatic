// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewListSeedsParams creates a new ListSeedsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSeedsParams() *ListSeedsParams {
	return &ListSeedsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSeedsParamsWithTimeout creates a new ListSeedsParams object
// with the ability to set a timeout on a request.
func NewListSeedsParamsWithTimeout(timeout time.Duration) *ListSeedsParams {
	return &ListSeedsParams{
		timeout: timeout,
	}
}

// NewListSeedsParamsWithContext creates a new ListSeedsParams object
// with the ability to set a context for a request.
func NewListSeedsParamsWithContext(ctx context.Context) *ListSeedsParams {
	return &ListSeedsParams{
		Context: ctx,
	}
}

// NewListSeedsParamsWithHTTPClient creates a new ListSeedsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSeedsParamsWithHTTPClient(client *http.Client) *ListSeedsParams {
	return &ListSeedsParams{
		HTTPClient: client,
	}
}

/* ListSeedsParams contains all the parameters to send to the API endpoint
   for the list seeds operation.

   Typically these are written to a http.Request.
*/
type ListSeedsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list seeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSeedsParams) WithDefaults() *ListSeedsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list seeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSeedsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list seeds params
func (o *ListSeedsParams) WithTimeout(timeout time.Duration) *ListSeedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list seeds params
func (o *ListSeedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list seeds params
func (o *ListSeedsParams) WithContext(ctx context.Context) *ListSeedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list seeds params
func (o *ListSeedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list seeds params
func (o *ListSeedsParams) WithHTTPClient(client *http.Client) *ListSeedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list seeds params
func (o *ListSeedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSeedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

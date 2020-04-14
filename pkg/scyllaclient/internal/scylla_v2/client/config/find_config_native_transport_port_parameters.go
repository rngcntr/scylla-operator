// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigNativeTransportPortParams creates a new FindConfigNativeTransportPortParams object
// with the default values initialized.
func NewFindConfigNativeTransportPortParams() *FindConfigNativeTransportPortParams {

	return &FindConfigNativeTransportPortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigNativeTransportPortParamsWithTimeout creates a new FindConfigNativeTransportPortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigNativeTransportPortParamsWithTimeout(timeout time.Duration) *FindConfigNativeTransportPortParams {

	return &FindConfigNativeTransportPortParams{

		timeout: timeout,
	}
}

// NewFindConfigNativeTransportPortParamsWithContext creates a new FindConfigNativeTransportPortParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigNativeTransportPortParamsWithContext(ctx context.Context) *FindConfigNativeTransportPortParams {

	return &FindConfigNativeTransportPortParams{

		Context: ctx,
	}
}

// NewFindConfigNativeTransportPortParamsWithHTTPClient creates a new FindConfigNativeTransportPortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigNativeTransportPortParamsWithHTTPClient(client *http.Client) *FindConfigNativeTransportPortParams {

	return &FindConfigNativeTransportPortParams{
		HTTPClient: client,
	}
}

/*FindConfigNativeTransportPortParams contains all the parameters to send to the API endpoint
for the find config native transport port operation typically these are written to a http.Request
*/
type FindConfigNativeTransportPortParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config native transport port params
func (o *FindConfigNativeTransportPortParams) WithTimeout(timeout time.Duration) *FindConfigNativeTransportPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config native transport port params
func (o *FindConfigNativeTransportPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config native transport port params
func (o *FindConfigNativeTransportPortParams) WithContext(ctx context.Context) *FindConfigNativeTransportPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config native transport port params
func (o *FindConfigNativeTransportPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config native transport port params
func (o *FindConfigNativeTransportPortParams) WithHTTPClient(client *http.Client) *FindConfigNativeTransportPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config native transport port params
func (o *FindConfigNativeTransportPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigNativeTransportPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

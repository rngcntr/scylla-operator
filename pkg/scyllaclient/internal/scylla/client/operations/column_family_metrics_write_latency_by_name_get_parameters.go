// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewColumnFamilyMetricsWriteLatencyByNameGetParams creates a new ColumnFamilyMetricsWriteLatencyByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsWriteLatencyByNameGetParams() *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsWriteLatencyByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsWriteLatencyByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyByNameGetParamsWithContext creates a new ColumnFamilyMetricsWriteLatencyByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsWriteLatencyByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsWriteLatencyByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsWriteLatencyByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsWriteLatencyByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsWriteLatencyByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics write latency by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsWriteLatencyByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) WithName(name string) *ColumnFamilyMetricsWriteLatencyByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics write latency by name get params
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsWriteLatencyByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/agent/models"
)

// NewOperationsDeletefileParams creates a new OperationsDeletefileParams object
// with the default values initialized.
func NewOperationsDeletefileParams() *OperationsDeletefileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsDeletefileParams{
		Async: asyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewOperationsDeletefileParamsWithTimeout creates a new OperationsDeletefileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperationsDeletefileParamsWithTimeout(timeout time.Duration) *OperationsDeletefileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsDeletefileParams{
		Async: asyncDefault,

		timeout: timeout,
	}
}

// NewOperationsDeletefileParamsWithContext creates a new OperationsDeletefileParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperationsDeletefileParamsWithContext(ctx context.Context) *OperationsDeletefileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsDeletefileParams{
		Async: asyncDefault,

		Context: ctx,
	}
}

// NewOperationsDeletefileParamsWithHTTPClient creates a new OperationsDeletefileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperationsDeletefileParamsWithHTTPClient(client *http.Client) *OperationsDeletefileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsDeletefileParams{
		Async:      asyncDefault,
		HTTPClient: client,
	}
}

/*OperationsDeletefileParams contains all the parameters to send to the API endpoint
for the operations deletefile operation typically these are written to a http.Request
*/
type OperationsDeletefileParams struct {

	/*Async
	  Async request

	*/
	Async bool
	/*Group
	  Place this operation under this stat group

	*/
	Group string
	/*Deletefile
	  deletefile

	*/
	Deletefile *models.RemotePath

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operations deletefile params
func (o *OperationsDeletefileParams) WithTimeout(timeout time.Duration) *OperationsDeletefileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operations deletefile params
func (o *OperationsDeletefileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operations deletefile params
func (o *OperationsDeletefileParams) WithContext(ctx context.Context) *OperationsDeletefileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operations deletefile params
func (o *OperationsDeletefileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operations deletefile params
func (o *OperationsDeletefileParams) WithHTTPClient(client *http.Client) *OperationsDeletefileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operations deletefile params
func (o *OperationsDeletefileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsync adds the async to the operations deletefile params
func (o *OperationsDeletefileParams) WithAsync(async bool) *OperationsDeletefileParams {
	o.SetAsync(async)
	return o
}

// SetAsync adds the async to the operations deletefile params
func (o *OperationsDeletefileParams) SetAsync(async bool) {
	o.Async = async
}

// WithGroup adds the group to the operations deletefile params
func (o *OperationsDeletefileParams) WithGroup(group string) *OperationsDeletefileParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the operations deletefile params
func (o *OperationsDeletefileParams) SetGroup(group string) {
	o.Group = group
}

// WithDeletefile adds the deletefile to the operations deletefile params
func (o *OperationsDeletefileParams) WithDeletefile(deletefile *models.RemotePath) *OperationsDeletefileParams {
	o.SetDeletefile(deletefile)
	return o
}

// SetDeletefile adds the deletefile to the operations deletefile params
func (o *OperationsDeletefileParams) SetDeletefile(deletefile *models.RemotePath) {
	o.Deletefile = deletefile
}

// WriteToRequest writes these params to a swagger request
func (o *OperationsDeletefileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param _async
	qrAsync := o.Async
	qAsync := swag.FormatBool(qrAsync)
	if qAsync != "" {
		if err := r.SetQueryParam("_async", qAsync); err != nil {
			return err
		}
	}

	// query param _group
	qrGroup := o.Group
	qGroup := qrGroup
	if qGroup != "" {
		if err := r.SetQueryParam("_group", qGroup); err != nil {
			return err
		}
	}

	if o.Deletefile != nil {
		if err := r.SetBodyParam(o.Deletefile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

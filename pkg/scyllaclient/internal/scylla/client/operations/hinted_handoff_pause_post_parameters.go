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
)

// NewHintedHandoffPausePostParams creates a new HintedHandoffPausePostParams object
// with the default values initialized.
func NewHintedHandoffPausePostParams() *HintedHandoffPausePostParams {
	var ()
	return &HintedHandoffPausePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHintedHandoffPausePostParamsWithTimeout creates a new HintedHandoffPausePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHintedHandoffPausePostParamsWithTimeout(timeout time.Duration) *HintedHandoffPausePostParams {
	var ()
	return &HintedHandoffPausePostParams{

		timeout: timeout,
	}
}

// NewHintedHandoffPausePostParamsWithContext creates a new HintedHandoffPausePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewHintedHandoffPausePostParamsWithContext(ctx context.Context) *HintedHandoffPausePostParams {
	var ()
	return &HintedHandoffPausePostParams{

		Context: ctx,
	}
}

// NewHintedHandoffPausePostParamsWithHTTPClient creates a new HintedHandoffPausePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHintedHandoffPausePostParamsWithHTTPClient(client *http.Client) *HintedHandoffPausePostParams {
	var ()
	return &HintedHandoffPausePostParams{
		HTTPClient: client,
	}
}

/*HintedHandoffPausePostParams contains all the parameters to send to the API endpoint
for the hinted handoff pause post operation typically these are written to a http.Request
*/
type HintedHandoffPausePostParams struct {

	/*Pause
	  pause status

	*/
	Pause bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) WithTimeout(timeout time.Duration) *HintedHandoffPausePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) WithContext(ctx context.Context) *HintedHandoffPausePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) WithHTTPClient(client *http.Client) *HintedHandoffPausePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPause adds the pause to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) WithPause(pause bool) *HintedHandoffPausePostParams {
	o.SetPause(pause)
	return o
}

// SetPause adds the pause to the hinted handoff pause post params
func (o *HintedHandoffPausePostParams) SetPause(pause bool) {
	o.Pause = pause
}

// WriteToRequest writes these params to a swagger request
func (o *HintedHandoffPausePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param pause
	qrPause := o.Pause
	qPause := swag.FormatBool(qrPause)
	if qPause != "" {
		if err := r.SetQueryParam("pause", qPause); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

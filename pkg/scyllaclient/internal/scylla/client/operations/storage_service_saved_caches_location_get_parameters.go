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

// NewStorageServiceSavedCachesLocationGetParams creates a new StorageServiceSavedCachesLocationGetParams object
// with the default values initialized.
func NewStorageServiceSavedCachesLocationGetParams() *StorageServiceSavedCachesLocationGetParams {

	return &StorageServiceSavedCachesLocationGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceSavedCachesLocationGetParamsWithTimeout creates a new StorageServiceSavedCachesLocationGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceSavedCachesLocationGetParamsWithTimeout(timeout time.Duration) *StorageServiceSavedCachesLocationGetParams {

	return &StorageServiceSavedCachesLocationGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceSavedCachesLocationGetParamsWithContext creates a new StorageServiceSavedCachesLocationGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceSavedCachesLocationGetParamsWithContext(ctx context.Context) *StorageServiceSavedCachesLocationGetParams {

	return &StorageServiceSavedCachesLocationGetParams{

		Context: ctx,
	}
}

// NewStorageServiceSavedCachesLocationGetParamsWithHTTPClient creates a new StorageServiceSavedCachesLocationGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceSavedCachesLocationGetParamsWithHTTPClient(client *http.Client) *StorageServiceSavedCachesLocationGetParams {

	return &StorageServiceSavedCachesLocationGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceSavedCachesLocationGetParams contains all the parameters to send to the API endpoint
for the storage service saved caches location get operation typically these are written to a http.Request
*/
type StorageServiceSavedCachesLocationGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service saved caches location get params
func (o *StorageServiceSavedCachesLocationGetParams) WithTimeout(timeout time.Duration) *StorageServiceSavedCachesLocationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service saved caches location get params
func (o *StorageServiceSavedCachesLocationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service saved caches location get params
func (o *StorageServiceSavedCachesLocationGetParams) WithContext(ctx context.Context) *StorageServiceSavedCachesLocationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service saved caches location get params
func (o *StorageServiceSavedCachesLocationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service saved caches location get params
func (o *StorageServiceSavedCachesLocationGetParams) WithHTTPClient(client *http.Client) *StorageServiceSavedCachesLocationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service saved caches location get params
func (o *StorageServiceSavedCachesLocationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceSavedCachesLocationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

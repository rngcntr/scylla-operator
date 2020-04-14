// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigReadRequestTimeoutInMsReader is a Reader for the FindConfigReadRequestTimeoutInMs structure.
type FindConfigReadRequestTimeoutInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigReadRequestTimeoutInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigReadRequestTimeoutInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigReadRequestTimeoutInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigReadRequestTimeoutInMsOK creates a FindConfigReadRequestTimeoutInMsOK with default headers values
func NewFindConfigReadRequestTimeoutInMsOK() *FindConfigReadRequestTimeoutInMsOK {
	return &FindConfigReadRequestTimeoutInMsOK{}
}

/*FindConfigReadRequestTimeoutInMsOK handles this case with default header values.

Config value
*/
type FindConfigReadRequestTimeoutInMsOK struct {
	Payload int64
}

func (o *FindConfigReadRequestTimeoutInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigReadRequestTimeoutInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigReadRequestTimeoutInMsDefault creates a FindConfigReadRequestTimeoutInMsDefault with default headers values
func NewFindConfigReadRequestTimeoutInMsDefault(code int) *FindConfigReadRequestTimeoutInMsDefault {
	return &FindConfigReadRequestTimeoutInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigReadRequestTimeoutInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigReadRequestTimeoutInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config read request timeout in ms default response
func (o *FindConfigReadRequestTimeoutInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigReadRequestTimeoutInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigReadRequestTimeoutInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigReadRequestTimeoutInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

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

// FindConfigStoragePortReader is a Reader for the FindConfigStoragePort structure.
type FindConfigStoragePortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigStoragePortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigStoragePortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigStoragePortDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigStoragePortOK creates a FindConfigStoragePortOK with default headers values
func NewFindConfigStoragePortOK() *FindConfigStoragePortOK {
	return &FindConfigStoragePortOK{}
}

/*FindConfigStoragePortOK handles this case with default header values.

Config value
*/
type FindConfigStoragePortOK struct {
	Payload int64
}

func (o *FindConfigStoragePortOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigStoragePortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigStoragePortDefault creates a FindConfigStoragePortDefault with default headers values
func NewFindConfigStoragePortDefault(code int) *FindConfigStoragePortDefault {
	return &FindConfigStoragePortDefault{
		_statusCode: code,
	}
}

/*FindConfigStoragePortDefault handles this case with default header values.

unexpected error
*/
type FindConfigStoragePortDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config storage port default response
func (o *FindConfigStoragePortDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigStoragePortDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigStoragePortDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigStoragePortDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceDrainGetReader is a Reader for the StorageServiceDrainGet structure.
type StorageServiceDrainGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceDrainGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceDrainGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceDrainGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceDrainGetOK creates a StorageServiceDrainGetOK with default headers values
func NewStorageServiceDrainGetOK() *StorageServiceDrainGetOK {
	return &StorageServiceDrainGetOK{}
}

/*StorageServiceDrainGetOK handles this case with default header values.

StorageServiceDrainGetOK storage service drain get o k
*/
type StorageServiceDrainGetOK struct {
	Payload string
}

func (o *StorageServiceDrainGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceDrainGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceDrainGetDefault creates a StorageServiceDrainGetDefault with default headers values
func NewStorageServiceDrainGetDefault(code int) *StorageServiceDrainGetDefault {
	return &StorageServiceDrainGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceDrainGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceDrainGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service drain get default response
func (o *StorageServiceDrainGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceDrainGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceDrainGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceDrainGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

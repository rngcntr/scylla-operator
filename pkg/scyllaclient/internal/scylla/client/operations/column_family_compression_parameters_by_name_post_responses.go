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

// ColumnFamilyCompressionParametersByNamePostReader is a Reader for the ColumnFamilyCompressionParametersByNamePost structure.
type ColumnFamilyCompressionParametersByNamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyCompressionParametersByNamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyCompressionParametersByNamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyCompressionParametersByNamePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyCompressionParametersByNamePostOK creates a ColumnFamilyCompressionParametersByNamePostOK with default headers values
func NewColumnFamilyCompressionParametersByNamePostOK() *ColumnFamilyCompressionParametersByNamePostOK {
	return &ColumnFamilyCompressionParametersByNamePostOK{}
}

/*ColumnFamilyCompressionParametersByNamePostOK handles this case with default header values.

ColumnFamilyCompressionParametersByNamePostOK column family compression parameters by name post o k
*/
type ColumnFamilyCompressionParametersByNamePostOK struct {
}

func (o *ColumnFamilyCompressionParametersByNamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyCompressionParametersByNamePostDefault creates a ColumnFamilyCompressionParametersByNamePostDefault with default headers values
func NewColumnFamilyCompressionParametersByNamePostDefault(code int) *ColumnFamilyCompressionParametersByNamePostDefault {
	return &ColumnFamilyCompressionParametersByNamePostDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyCompressionParametersByNamePostDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyCompressionParametersByNamePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family compression parameters by name post default response
func (o *ColumnFamilyCompressionParametersByNamePostDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyCompressionParametersByNamePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyCompressionParametersByNamePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyCompressionParametersByNamePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

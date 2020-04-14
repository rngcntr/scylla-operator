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

// CompactionManagerStopCompactionPostReader is a Reader for the CompactionManagerStopCompactionPost structure.
type CompactionManagerStopCompactionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompactionManagerStopCompactionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompactionManagerStopCompactionPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCompactionManagerStopCompactionPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompactionManagerStopCompactionPostOK creates a CompactionManagerStopCompactionPostOK with default headers values
func NewCompactionManagerStopCompactionPostOK() *CompactionManagerStopCompactionPostOK {
	return &CompactionManagerStopCompactionPostOK{}
}

/*CompactionManagerStopCompactionPostOK handles this case with default header values.

CompactionManagerStopCompactionPostOK compaction manager stop compaction post o k
*/
type CompactionManagerStopCompactionPostOK struct {
}

func (o *CompactionManagerStopCompactionPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompactionManagerStopCompactionPostDefault creates a CompactionManagerStopCompactionPostDefault with default headers values
func NewCompactionManagerStopCompactionPostDefault(code int) *CompactionManagerStopCompactionPostDefault {
	return &CompactionManagerStopCompactionPostDefault{
		_statusCode: code,
	}
}

/*CompactionManagerStopCompactionPostDefault handles this case with default header values.

internal server error
*/
type CompactionManagerStopCompactionPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the compaction manager stop compaction post default response
func (o *CompactionManagerStopCompactionPostDefault) Code() int {
	return o._statusCode
}

func (o *CompactionManagerStopCompactionPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CompactionManagerStopCompactionPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CompactionManagerStopCompactionPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

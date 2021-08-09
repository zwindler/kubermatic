// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListGCPDiskTypesReader is a Reader for the ListGCPDiskTypes structure.
type ListGCPDiskTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPDiskTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPDiskTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPDiskTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPDiskTypesOK creates a ListGCPDiskTypesOK with default headers values
func NewListGCPDiskTypesOK() *ListGCPDiskTypesOK {
	return &ListGCPDiskTypesOK{}
}

/* ListGCPDiskTypesOK describes a response with status code 200, with default header values.

GCPDiskTypeList
*/
type ListGCPDiskTypesOK struct {
	Payload models.GCPDiskTypeList
}

func (o *ListGCPDiskTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/disktypes][%d] listGCPDiskTypesOK  %+v", 200, o.Payload)
}
func (o *ListGCPDiskTypesOK) GetPayload() models.GCPDiskTypeList {
	return o.Payload
}

func (o *ListGCPDiskTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPDiskTypesDefault creates a ListGCPDiskTypesDefault with default headers values
func NewListGCPDiskTypesDefault(code int) *ListGCPDiskTypesDefault {
	return &ListGCPDiskTypesDefault{
		_statusCode: code,
	}
}

/* ListGCPDiskTypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPDiskTypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p disk types default response
func (o *ListGCPDiskTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPDiskTypesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/disktypes][%d] listGCPDiskTypes default  %+v", o._statusCode, o.Payload)
}
func (o *ListGCPDiskTypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPDiskTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

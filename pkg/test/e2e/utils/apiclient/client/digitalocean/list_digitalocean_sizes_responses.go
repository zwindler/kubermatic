// Code generated by go-swagger; DO NOT EDIT.

package digitalocean

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListDigitaloceanSizesReader is a Reader for the ListDigitaloceanSizes structure.
type ListDigitaloceanSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDigitaloceanSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDigitaloceanSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDigitaloceanSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDigitaloceanSizesOK creates a ListDigitaloceanSizesOK with default headers values
func NewListDigitaloceanSizesOK() *ListDigitaloceanSizesOK {
	return &ListDigitaloceanSizesOK{}
}

/* ListDigitaloceanSizesOK describes a response with status code 200, with default header values.

DigitaloceanSizeList
*/
type ListDigitaloceanSizesOK struct {
	Payload *models.DigitaloceanSizeList
}

func (o *ListDigitaloceanSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/digitalocean/sizes][%d] listDigitaloceanSizesOK  %+v", 200, o.Payload)
}
func (o *ListDigitaloceanSizesOK) GetPayload() *models.DigitaloceanSizeList {
	return o.Payload
}

func (o *ListDigitaloceanSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DigitaloceanSizeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDigitaloceanSizesDefault creates a ListDigitaloceanSizesDefault with default headers values
func NewListDigitaloceanSizesDefault(code int) *ListDigitaloceanSizesDefault {
	return &ListDigitaloceanSizesDefault{
		_statusCode: code,
	}
}

/* ListDigitaloceanSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListDigitaloceanSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list digitalocean sizes default response
func (o *ListDigitaloceanSizesDefault) Code() int {
	return o._statusCode
}

func (o *ListDigitaloceanSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/digitalocean/sizes][%d] listDigitaloceanSizes default  %+v", o._statusCode, o.Payload)
}
func (o *ListDigitaloceanSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListDigitaloceanSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

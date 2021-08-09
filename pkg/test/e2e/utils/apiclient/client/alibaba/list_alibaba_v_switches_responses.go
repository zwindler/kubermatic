// Code generated by go-swagger; DO NOT EDIT.

package alibaba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAlibabaVSwitchesReader is a Reader for the ListAlibabaVSwitches structure.
type ListAlibabaVSwitchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAlibabaVSwitchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAlibabaVSwitchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAlibabaVSwitchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAlibabaVSwitchesOK creates a ListAlibabaVSwitchesOK with default headers values
func NewListAlibabaVSwitchesOK() *ListAlibabaVSwitchesOK {
	return &ListAlibabaVSwitchesOK{}
}

/* ListAlibabaVSwitchesOK describes a response with status code 200, with default header values.

AlibabaVSwitchList
*/
type ListAlibabaVSwitchesOK struct {
	Payload models.AlibabaVSwitchList
}

func (o *ListAlibabaVSwitchesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/alibaba/vswitches][%d] listAlibabaVSwitchesOK  %+v", 200, o.Payload)
}
func (o *ListAlibabaVSwitchesOK) GetPayload() models.AlibabaVSwitchList {
	return o.Payload
}

func (o *ListAlibabaVSwitchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlibabaVSwitchesDefault creates a ListAlibabaVSwitchesDefault with default headers values
func NewListAlibabaVSwitchesDefault(code int) *ListAlibabaVSwitchesDefault {
	return &ListAlibabaVSwitchesDefault{
		_statusCode: code,
	}
}

/* ListAlibabaVSwitchesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAlibabaVSwitchesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list alibaba v switches default response
func (o *ListAlibabaVSwitchesDefault) Code() int {
	return o._statusCode
}

func (o *ListAlibabaVSwitchesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/alibaba/vswitches][%d] listAlibabaVSwitches default  %+v", o._statusCode, o.Payload)
}
func (o *ListAlibabaVSwitchesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAlibabaVSwitchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

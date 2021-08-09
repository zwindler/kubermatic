// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAzureSizesNoCredentialsV2Reader is a Reader for the ListAzureSizesNoCredentialsV2 structure.
type ListAzureSizesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureSizesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureSizesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureSizesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureSizesNoCredentialsV2OK creates a ListAzureSizesNoCredentialsV2OK with default headers values
func NewListAzureSizesNoCredentialsV2OK() *ListAzureSizesNoCredentialsV2OK {
	return &ListAzureSizesNoCredentialsV2OK{}
}

/* ListAzureSizesNoCredentialsV2OK describes a response with status code 200, with default header values.

AzureSizeList
*/
type ListAzureSizesNoCredentialsV2OK struct {
	Payload models.AzureSizeList
}

func (o *ListAzureSizesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes][%d] listAzureSizesNoCredentialsV2OK  %+v", 200, o.Payload)
}
func (o *ListAzureSizesNoCredentialsV2OK) GetPayload() models.AzureSizeList {
	return o.Payload
}

func (o *ListAzureSizesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureSizesNoCredentialsV2Default creates a ListAzureSizesNoCredentialsV2Default with default headers values
func NewListAzureSizesNoCredentialsV2Default(code int) *ListAzureSizesNoCredentialsV2Default {
	return &ListAzureSizesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListAzureSizesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureSizesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure sizes no credentials v2 default response
func (o *ListAzureSizesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAzureSizesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes][%d] listAzureSizesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}
func (o *ListAzureSizesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureSizesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

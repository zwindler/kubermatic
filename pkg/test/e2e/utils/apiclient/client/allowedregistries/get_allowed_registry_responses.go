// Code generated by go-swagger; DO NOT EDIT.

package allowedregistries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetAllowedRegistryReader is a Reader for the GetAllowedRegistry structure.
type GetAllowedRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllowedRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllowedRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllowedRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllowedRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAllowedRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllowedRegistryOK creates a GetAllowedRegistryOK with default headers values
func NewGetAllowedRegistryOK() *GetAllowedRegistryOK {
	return &GetAllowedRegistryOK{}
}

/* GetAllowedRegistryOK describes a response with status code 200, with default header values.

AllowedRegistry
*/
type GetAllowedRegistryOK struct {
	Payload *models.AllowedRegistry
}

func (o *GetAllowedRegistryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries/{allowed_registry}][%d] getAllowedRegistryOK  %+v", 200, o.Payload)
}
func (o *GetAllowedRegistryOK) GetPayload() *models.AllowedRegistry {
	return o.Payload
}

func (o *GetAllowedRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllowedRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllowedRegistryUnauthorized creates a GetAllowedRegistryUnauthorized with default headers values
func NewGetAllowedRegistryUnauthorized() *GetAllowedRegistryUnauthorized {
	return &GetAllowedRegistryUnauthorized{}
}

/* GetAllowedRegistryUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetAllowedRegistryUnauthorized struct {
}

func (o *GetAllowedRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries/{allowed_registry}][%d] getAllowedRegistryUnauthorized ", 401)
}

func (o *GetAllowedRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllowedRegistryForbidden creates a GetAllowedRegistryForbidden with default headers values
func NewGetAllowedRegistryForbidden() *GetAllowedRegistryForbidden {
	return &GetAllowedRegistryForbidden{}
}

/* GetAllowedRegistryForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetAllowedRegistryForbidden struct {
}

func (o *GetAllowedRegistryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries/{allowed_registry}][%d] getAllowedRegistryForbidden ", 403)
}

func (o *GetAllowedRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllowedRegistryDefault creates a GetAllowedRegistryDefault with default headers values
func NewGetAllowedRegistryDefault(code int) *GetAllowedRegistryDefault {
	return &GetAllowedRegistryDefault{
		_statusCode: code,
	}
}

/* GetAllowedRegistryDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetAllowedRegistryDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get allowed registry default response
func (o *GetAllowedRegistryDefault) Code() int {
	return o._statusCode
}

func (o *GetAllowedRegistryDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries/{allowed_registry}][%d] getAllowedRegistry default  %+v", o._statusCode, o.Payload)
}
func (o *GetAllowedRegistryDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllowedRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

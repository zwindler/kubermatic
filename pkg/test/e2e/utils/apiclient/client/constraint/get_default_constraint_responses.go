// Code generated by go-swagger; DO NOT EDIT.

package constraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetDefaultConstraintReader is a Reader for the GetDefaultConstraint structure.
type GetDefaultConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefaultConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDefaultConstraintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDefaultConstraintForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDefaultConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDefaultConstraintOK creates a GetDefaultConstraintOK with default headers values
func NewGetDefaultConstraintOK() *GetDefaultConstraintOK {
	return &GetDefaultConstraintOK{}
}

/* GetDefaultConstraintOK describes a response with status code 200, with default header values.

Constraint
*/
type GetDefaultConstraintOK struct {
	Payload *models.Constraint
}

func (o *GetDefaultConstraintOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints/{constraint_name}][%d] getDefaultConstraintOK  %+v", 200, o.Payload)
}
func (o *GetDefaultConstraintOK) GetPayload() *models.Constraint {
	return o.Payload
}

func (o *GetDefaultConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Constraint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultConstraintUnauthorized creates a GetDefaultConstraintUnauthorized with default headers values
func NewGetDefaultConstraintUnauthorized() *GetDefaultConstraintUnauthorized {
	return &GetDefaultConstraintUnauthorized{}
}

/* GetDefaultConstraintUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetDefaultConstraintUnauthorized struct {
}

func (o *GetDefaultConstraintUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints/{constraint_name}][%d] getDefaultConstraintUnauthorized ", 401)
}

func (o *GetDefaultConstraintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDefaultConstraintForbidden creates a GetDefaultConstraintForbidden with default headers values
func NewGetDefaultConstraintForbidden() *GetDefaultConstraintForbidden {
	return &GetDefaultConstraintForbidden{}
}

/* GetDefaultConstraintForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetDefaultConstraintForbidden struct {
}

func (o *GetDefaultConstraintForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints/{constraint_name}][%d] getDefaultConstraintForbidden ", 403)
}

func (o *GetDefaultConstraintForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDefaultConstraintDefault creates a GetDefaultConstraintDefault with default headers values
func NewGetDefaultConstraintDefault(code int) *GetDefaultConstraintDefault {
	return &GetDefaultConstraintDefault{
		_statusCode: code,
	}
}

/* GetDefaultConstraintDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetDefaultConstraintDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get default constraint default response
func (o *GetDefaultConstraintDefault) Code() int {
	return o._statusCode
}

func (o *GetDefaultConstraintDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints/{constraint_name}][%d] getDefaultConstraint default  %+v", o._statusCode, o.Payload)
}
func (o *GetDefaultConstraintDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDefaultConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

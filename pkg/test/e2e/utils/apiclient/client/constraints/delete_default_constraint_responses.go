// Code generated by go-swagger; DO NOT EDIT.

package constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteDefaultConstraintReader is a Reader for the DeleteDefaultConstraint structure.
type DeleteDefaultConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDefaultConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDefaultConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDefaultConstraintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDefaultConstraintForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDefaultConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDefaultConstraintOK creates a DeleteDefaultConstraintOK with default headers values
func NewDeleteDefaultConstraintOK() *DeleteDefaultConstraintOK {
	return &DeleteDefaultConstraintOK{}
}

/* DeleteDefaultConstraintOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteDefaultConstraintOK struct {
}

func (o *DeleteDefaultConstraintOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constraints/{constraint_name}][%d] deleteDefaultConstraintOK ", 200)
}

func (o *DeleteDefaultConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDefaultConstraintUnauthorized creates a DeleteDefaultConstraintUnauthorized with default headers values
func NewDeleteDefaultConstraintUnauthorized() *DeleteDefaultConstraintUnauthorized {
	return &DeleteDefaultConstraintUnauthorized{}
}

/* DeleteDefaultConstraintUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteDefaultConstraintUnauthorized struct {
}

func (o *DeleteDefaultConstraintUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constraints/{constraint_name}][%d] deleteDefaultConstraintUnauthorized ", 401)
}

func (o *DeleteDefaultConstraintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDefaultConstraintForbidden creates a DeleteDefaultConstraintForbidden with default headers values
func NewDeleteDefaultConstraintForbidden() *DeleteDefaultConstraintForbidden {
	return &DeleteDefaultConstraintForbidden{}
}

/* DeleteDefaultConstraintForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteDefaultConstraintForbidden struct {
}

func (o *DeleteDefaultConstraintForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constraints/{constraint_name}][%d] deleteDefaultConstraintForbidden ", 403)
}

func (o *DeleteDefaultConstraintForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDefaultConstraintDefault creates a DeleteDefaultConstraintDefault with default headers values
func NewDeleteDefaultConstraintDefault(code int) *DeleteDefaultConstraintDefault {
	return &DeleteDefaultConstraintDefault{
		_statusCode: code,
	}
}

/* DeleteDefaultConstraintDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteDefaultConstraintDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete default constraint default response
func (o *DeleteDefaultConstraintDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDefaultConstraintDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constraints/{constraint_name}][%d] deleteDefaultConstraint default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteDefaultConstraintDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDefaultConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

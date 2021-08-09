// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListRoleV2Reader is a Reader for the ListRoleV2 structure.
type ListRoleV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRoleV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRoleV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRoleV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRoleV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRoleV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRoleV2OK creates a ListRoleV2OK with default headers values
func NewListRoleV2OK() *ListRoleV2OK {
	return &ListRoleV2OK{}
}

/* ListRoleV2OK describes a response with status code 200, with default header values.

Role
*/
type ListRoleV2OK struct {
	Payload []*models.Role
}

func (o *ListRoleV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/roles][%d] listRoleV2OK  %+v", 200, o.Payload)
}
func (o *ListRoleV2OK) GetPayload() []*models.Role {
	return o.Payload
}

func (o *ListRoleV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRoleV2Unauthorized creates a ListRoleV2Unauthorized with default headers values
func NewListRoleV2Unauthorized() *ListRoleV2Unauthorized {
	return &ListRoleV2Unauthorized{}
}

/* ListRoleV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListRoleV2Unauthorized struct {
}

func (o *ListRoleV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/roles][%d] listRoleV2Unauthorized ", 401)
}

func (o *ListRoleV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleV2Forbidden creates a ListRoleV2Forbidden with default headers values
func NewListRoleV2Forbidden() *ListRoleV2Forbidden {
	return &ListRoleV2Forbidden{}
}

/* ListRoleV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListRoleV2Forbidden struct {
}

func (o *ListRoleV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/roles][%d] listRoleV2Forbidden ", 403)
}

func (o *ListRoleV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleV2Default creates a ListRoleV2Default with default headers values
func NewListRoleV2Default(code int) *ListRoleV2Default {
	return &ListRoleV2Default{
		_statusCode: code,
	}
}

/* ListRoleV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListRoleV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list role v2 default response
func (o *ListRoleV2Default) Code() int {
	return o._statusCode
}

func (o *ListRoleV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/roles][%d] listRoleV2 default  %+v", o._statusCode, o.Payload)
}
func (o *ListRoleV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListRoleV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

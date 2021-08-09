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

// ListClustersForProjectReader is a Reader for the ListClustersForProject structure.
type ListClustersForProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersForProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersForProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClustersForProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClustersForProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClustersForProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClustersForProjectOK creates a ListClustersForProjectOK with default headers values
func NewListClustersForProjectOK() *ListClustersForProjectOK {
	return &ListClustersForProjectOK{}
}

/* ListClustersForProjectOK describes a response with status code 200, with default header values.

ClusterList
*/
type ListClustersForProjectOK struct {
	Payload models.ClusterList
}

func (o *ListClustersForProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters][%d] listClustersForProjectOK  %+v", 200, o.Payload)
}
func (o *ListClustersForProjectOK) GetPayload() models.ClusterList {
	return o.Payload
}

func (o *ListClustersForProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersForProjectUnauthorized creates a ListClustersForProjectUnauthorized with default headers values
func NewListClustersForProjectUnauthorized() *ListClustersForProjectUnauthorized {
	return &ListClustersForProjectUnauthorized{}
}

/* ListClustersForProjectUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListClustersForProjectUnauthorized struct {
}

func (o *ListClustersForProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters][%d] listClustersForProjectUnauthorized ", 401)
}

func (o *ListClustersForProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersForProjectForbidden creates a ListClustersForProjectForbidden with default headers values
func NewListClustersForProjectForbidden() *ListClustersForProjectForbidden {
	return &ListClustersForProjectForbidden{}
}

/* ListClustersForProjectForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListClustersForProjectForbidden struct {
}

func (o *ListClustersForProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters][%d] listClustersForProjectForbidden ", 403)
}

func (o *ListClustersForProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersForProjectDefault creates a ListClustersForProjectDefault with default headers values
func NewListClustersForProjectDefault(code int) *ListClustersForProjectDefault {
	return &ListClustersForProjectDefault{
		_statusCode: code,
	}
}

/* ListClustersForProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListClustersForProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list clusters for project default response
func (o *ListClustersForProjectDefault) Code() int {
	return o._statusCode
}

func (o *ListClustersForProjectDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters][%d] listClustersForProject default  %+v", o._statusCode, o.Payload)
}
func (o *ListClustersForProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClustersForProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

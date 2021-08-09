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

// PatchNodeDeploymentReader is a Reader for the PatchNodeDeployment structure.
type PatchNodeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchNodeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchNodeDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchNodeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchNodeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchNodeDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchNodeDeploymentOK creates a PatchNodeDeploymentOK with default headers values
func NewPatchNodeDeploymentOK() *PatchNodeDeploymentOK {
	return &PatchNodeDeploymentOK{}
}

/* PatchNodeDeploymentOK describes a response with status code 200, with default header values.

NodeDeployment
*/
type PatchNodeDeploymentOK struct {
	Payload *models.NodeDeployment
}

func (o *PatchNodeDeploymentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] patchNodeDeploymentOK  %+v", 200, o.Payload)
}
func (o *PatchNodeDeploymentOK) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *PatchNodeDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNodeDeploymentUnauthorized creates a PatchNodeDeploymentUnauthorized with default headers values
func NewPatchNodeDeploymentUnauthorized() *PatchNodeDeploymentUnauthorized {
	return &PatchNodeDeploymentUnauthorized{}
}

/* PatchNodeDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchNodeDeploymentUnauthorized struct {
}

func (o *PatchNodeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] patchNodeDeploymentUnauthorized ", 401)
}

func (o *PatchNodeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchNodeDeploymentForbidden creates a PatchNodeDeploymentForbidden with default headers values
func NewPatchNodeDeploymentForbidden() *PatchNodeDeploymentForbidden {
	return &PatchNodeDeploymentForbidden{}
}

/* PatchNodeDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchNodeDeploymentForbidden struct {
}

func (o *PatchNodeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] patchNodeDeploymentForbidden ", 403)
}

func (o *PatchNodeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchNodeDeploymentDefault creates a PatchNodeDeploymentDefault with default headers values
func NewPatchNodeDeploymentDefault(code int) *PatchNodeDeploymentDefault {
	return &PatchNodeDeploymentDefault{
		_statusCode: code,
	}
}

/* PatchNodeDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchNodeDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch node deployment default response
func (o *PatchNodeDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *PatchNodeDeploymentDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] patchNodeDeployment default  %+v", o._statusCode, o.Payload)
}
func (o *PatchNodeDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchNodeDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

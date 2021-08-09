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

// PatchClusterV2Reader is a Reader for the PatchClusterV2 structure.
type PatchClusterV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchClusterV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchClusterV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchClusterV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchClusterV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchClusterV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchClusterV2OK creates a PatchClusterV2OK with default headers values
func NewPatchClusterV2OK() *PatchClusterV2OK {
	return &PatchClusterV2OK{}
}

/* PatchClusterV2OK describes a response with status code 200, with default header values.

Cluster
*/
type PatchClusterV2OK struct {
	Payload *models.Cluster
}

func (o *PatchClusterV2OK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] patchClusterV2OK  %+v", 200, o.Payload)
}
func (o *PatchClusterV2OK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *PatchClusterV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClusterV2Unauthorized creates a PatchClusterV2Unauthorized with default headers values
func NewPatchClusterV2Unauthorized() *PatchClusterV2Unauthorized {
	return &PatchClusterV2Unauthorized{}
}

/* PatchClusterV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterV2Unauthorized struct {
}

func (o *PatchClusterV2Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] patchClusterV2Unauthorized ", 401)
}

func (o *PatchClusterV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterV2Forbidden creates a PatchClusterV2Forbidden with default headers values
func NewPatchClusterV2Forbidden() *PatchClusterV2Forbidden {
	return &PatchClusterV2Forbidden{}
}

/* PatchClusterV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterV2Forbidden struct {
}

func (o *PatchClusterV2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] patchClusterV2Forbidden ", 403)
}

func (o *PatchClusterV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterV2Default creates a PatchClusterV2Default with default headers values
func NewPatchClusterV2Default(code int) *PatchClusterV2Default {
	return &PatchClusterV2Default{
		_statusCode: code,
	}
}

/* PatchClusterV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type PatchClusterV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch cluster v2 default response
func (o *PatchClusterV2Default) Code() int {
	return o._statusCode
}

func (o *PatchClusterV2Default) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] patchClusterV2 default  %+v", o._statusCode, o.Payload)
}
func (o *PatchClusterV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchClusterV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

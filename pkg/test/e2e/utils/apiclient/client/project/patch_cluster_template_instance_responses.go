// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchClusterTemplateInstanceReader is a Reader for the PatchClusterTemplateInstance structure.
type PatchClusterTemplateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchClusterTemplateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchClusterTemplateInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchClusterTemplateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchClusterTemplateInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchClusterTemplateInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchClusterTemplateInstanceOK creates a PatchClusterTemplateInstanceOK with default headers values
func NewPatchClusterTemplateInstanceOK() *PatchClusterTemplateInstanceOK {
	return &PatchClusterTemplateInstanceOK{}
}

/* PatchClusterTemplateInstanceOK describes a response with status code 200, with default header values.

ClusterTemplateInstance
*/
type PatchClusterTemplateInstanceOK struct {
	Payload *models.ClusterTemplateInstance
}

func (o *PatchClusterTemplateInstanceOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances/{instance_id}][%d] patchClusterTemplateInstanceOK  %+v", 200, o.Payload)
}
func (o *PatchClusterTemplateInstanceOK) GetPayload() *models.ClusterTemplateInstance {
	return o.Payload
}

func (o *PatchClusterTemplateInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTemplateInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClusterTemplateInstanceUnauthorized creates a PatchClusterTemplateInstanceUnauthorized with default headers values
func NewPatchClusterTemplateInstanceUnauthorized() *PatchClusterTemplateInstanceUnauthorized {
	return &PatchClusterTemplateInstanceUnauthorized{}
}

/* PatchClusterTemplateInstanceUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterTemplateInstanceUnauthorized struct {
}

func (o *PatchClusterTemplateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances/{instance_id}][%d] patchClusterTemplateInstanceUnauthorized ", 401)
}

func (o *PatchClusterTemplateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterTemplateInstanceForbidden creates a PatchClusterTemplateInstanceForbidden with default headers values
func NewPatchClusterTemplateInstanceForbidden() *PatchClusterTemplateInstanceForbidden {
	return &PatchClusterTemplateInstanceForbidden{}
}

/* PatchClusterTemplateInstanceForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterTemplateInstanceForbidden struct {
}

func (o *PatchClusterTemplateInstanceForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances/{instance_id}][%d] patchClusterTemplateInstanceForbidden ", 403)
}

func (o *PatchClusterTemplateInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterTemplateInstanceDefault creates a PatchClusterTemplateInstanceDefault with default headers values
func NewPatchClusterTemplateInstanceDefault(code int) *PatchClusterTemplateInstanceDefault {
	return &PatchClusterTemplateInstanceDefault{
		_statusCode: code,
	}
}

/* PatchClusterTemplateInstanceDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchClusterTemplateInstanceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch cluster template instance default response
func (o *PatchClusterTemplateInstanceDefault) Code() int {
	return o._statusCode
}

func (o *PatchClusterTemplateInstanceDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances/{instance_id}][%d] patchClusterTemplateInstance default  %+v", o._statusCode, o.Payload)
}
func (o *PatchClusterTemplateInstanceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchClusterTemplateInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchClusterTemplateInstanceBody patch cluster template instance body
swagger:model PatchClusterTemplateInstanceBody
*/
type PatchClusterTemplateInstanceBody struct {

	// replicas
	Replicas int64 `json:"replicas,omitempty"`
}

// Validate validates this patch cluster template instance body
func (o *PatchClusterTemplateInstanceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch cluster template instance body based on context it is used
func (o *PatchClusterTemplateInstanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchClusterTemplateInstanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchClusterTemplateInstanceBody) UnmarshalBinary(b []byte) error {
	var res PatchClusterTemplateInstanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

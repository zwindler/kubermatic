// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListOpenstackSecurityGroupsReader is a Reader for the ListOpenstackSecurityGroups structure.
type ListOpenstackSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackSecurityGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackSecurityGroupsOK creates a ListOpenstackSecurityGroupsOK with default headers values
func NewListOpenstackSecurityGroupsOK() *ListOpenstackSecurityGroupsOK {
	return &ListOpenstackSecurityGroupsOK{}
}

/* ListOpenstackSecurityGroupsOK describes a response with status code 200, with default header values.

OpenstackSecurityGroup
*/
type ListOpenstackSecurityGroupsOK struct {
	Payload []*models.OpenstackSecurityGroup
}

func (o *ListOpenstackSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/securitygroups][%d] listOpenstackSecurityGroupsOK  %+v", 200, o.Payload)
}
func (o *ListOpenstackSecurityGroupsOK) GetPayload() []*models.OpenstackSecurityGroup {
	return o.Payload
}

func (o *ListOpenstackSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackSecurityGroupsDefault creates a ListOpenstackSecurityGroupsDefault with default headers values
func NewListOpenstackSecurityGroupsDefault(code int) *ListOpenstackSecurityGroupsDefault {
	return &ListOpenstackSecurityGroupsDefault{
		_statusCode: code,
	}
}

/* ListOpenstackSecurityGroupsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackSecurityGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack security groups default response
func (o *ListOpenstackSecurityGroupsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackSecurityGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/securitygroups][%d] listOpenstackSecurityGroups default  %+v", o._statusCode, o.Payload)
}
func (o *ListOpenstackSecurityGroupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackSecurityGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

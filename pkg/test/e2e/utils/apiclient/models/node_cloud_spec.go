// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeCloudSpec NodeCloudSpec represents the collection of cloud provider specific settings. Only one must be set at a time.
//
// swagger:model NodeCloudSpec
type NodeCloudSpec struct {

	// alibaba
	Alibaba *AlibabaNodeSpec `json:"alibaba,omitempty"`

	// anexia
	Anexia *AnexiaNodeSpec `json:"anexia,omitempty"`

	// aws
	Aws *AWSNodeSpec `json:"aws,omitempty"`

	// azure
	Azure *AzureNodeSpec `json:"azure,omitempty"`

	// digitalocean
	Digitalocean *DigitaloceanNodeSpec `json:"digitalocean,omitempty"`

	// gcp
	Gcp *GCPNodeSpec `json:"gcp,omitempty"`

	// hetzner
	Hetzner *HetznerNodeSpec `json:"hetzner,omitempty"`

	// kubevirt
	Kubevirt *KubevirtNodeSpec `json:"kubevirt,omitempty"`

	// openstack
	Openstack *OpenstackNodeSpec `json:"openstack,omitempty"`

	// packet
	Packet *PacketNodeSpec `json:"packet,omitempty"`

	// vsphere
	Vsphere *VSphereNodeSpec `json:"vsphere,omitempty"`
}

// Validate validates this node cloud spec
func (m *NodeCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlibaba(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnexia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigitalocean(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHetzner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubevirt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeCloudSpec) validateAlibaba(formats strfmt.Registry) error {
	if swag.IsZero(m.Alibaba) { // not required
		return nil
	}

	if m.Alibaba != nil {
		if err := m.Alibaba.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alibaba")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateAnexia(formats strfmt.Registry) error {
	if swag.IsZero(m.Anexia) { // not required
		return nil
	}

	if m.Anexia != nil {
		if err := m.Anexia.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("anexia")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateDigitalocean(formats strfmt.Registry) error {
	if swag.IsZero(m.Digitalocean) { // not required
		return nil
	}

	if m.Digitalocean != nil {
		if err := m.Digitalocean.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateGcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateHetzner(formats strfmt.Registry) error {
	if swag.IsZero(m.Hetzner) { // not required
		return nil
	}

	if m.Hetzner != nil {
		if err := m.Hetzner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateKubevirt(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubevirt) { // not required
		return nil
	}

	if m.Kubevirt != nil {
		if err := m.Kubevirt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validatePacket(formats strfmt.Registry) error {
	if swag.IsZero(m.Packet) { // not required
		return nil
	}

	if m.Packet != nil {
		if err := m.Packet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packet")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) validateVsphere(formats strfmt.Registry) error {
	if swag.IsZero(m.Vsphere) { // not required
		return nil
	}

	if m.Vsphere != nil {
		if err := m.Vsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this node cloud spec based on the context it is used
func (m *NodeCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlibaba(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnexia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigitalocean(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHetzner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubevirt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePacket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeCloudSpec) contextValidateAlibaba(ctx context.Context, formats strfmt.Registry) error {

	if m.Alibaba != nil {
		if err := m.Alibaba.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alibaba")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateAnexia(ctx context.Context, formats strfmt.Registry) error {

	if m.Anexia != nil {
		if err := m.Anexia.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("anexia")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {
		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {
		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateDigitalocean(ctx context.Context, formats strfmt.Registry) error {

	if m.Digitalocean != nil {
		if err := m.Digitalocean.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {
		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateHetzner(ctx context.Context, formats strfmt.Registry) error {

	if m.Hetzner != nil {
		if err := m.Hetzner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateKubevirt(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubevirt != nil {
		if err := m.Kubevirt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	if m.Openstack != nil {
		if err := m.Openstack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidatePacket(ctx context.Context, formats strfmt.Registry) error {

	if m.Packet != nil {
		if err := m.Packet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packet")
			}
			return err
		}
	}

	return nil
}

func (m *NodeCloudSpec) contextValidateVsphere(ctx context.Context, formats strfmt.Registry) error {

	if m.Vsphere != nil {
		if err := m.Vsphere.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeCloudSpec) UnmarshalBinary(b []byte) error {
	var res NodeCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MachineDeploymentVMResourceQuota machine deployment VM resource quota
//
// swagger:model MachineDeploymentVMResourceQuota
type MachineDeploymentVMResourceQuota struct {

	// enable g p u
	EnableGPU bool `json:"enableGPU,omitempty"`

	// Maximal number of vCPU
	MaxCPU int64 `json:"maxCPU,omitempty"`

	// Maximum RAM size in GB
	MaxRAM int64 `json:"maxRAM,omitempty"`

	// Minimal number of vCPU
	MinCPU int64 `json:"minCPU,omitempty"`

	// Minimal RAM size in GB
	MinRAM int64 `json:"minRAM,omitempty"`
}

// Validate validates this machine deployment VM resource quota
func (m *MachineDeploymentVMResourceQuota) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this machine deployment VM resource quota based on context it is used
func (m *MachineDeploymentVMResourceQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MachineDeploymentVMResourceQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineDeploymentVMResourceQuota) UnmarshalBinary(b []byte) error {
	var res MachineDeploymentVMResourceQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

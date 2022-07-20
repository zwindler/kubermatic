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

// ResourceQuota resource quota
//
// swagger:model ResourceQuota
type ResourceQuota struct {

	// name
	Name string `json:"name,omitempty"`

	// SubjectHumanReadableName contains the human-readable name for the subject(if applicable). Just filled as information in get/list.
	SubjectHumanReadableName string `json:"subjectHumanReadableName,omitempty"`

	// subject kind
	SubjectKind string `json:"subjectKind,omitempty"`

	// subject name
	SubjectName string `json:"subjectName,omitempty"`

	// quota
	Quota *Quota `json:"quota,omitempty"`

	// status
	Status *ResourceQuotaStatus `json:"status,omitempty"`
}

// Validate validates this resource quota
func (m *ResourceQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceQuota) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceQuota) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resource quota based on the context it is used
func (m *ResourceQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceQuota) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.Quota != nil {
		if err := m.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceQuota) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceQuota) UnmarshalBinary(b []byte) error {
	var res ResourceQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

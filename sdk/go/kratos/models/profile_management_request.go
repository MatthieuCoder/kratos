// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProfileManagementRequest Request presents a profile management request
//
// This request is used when an identity wants to update profile information
// (especially traits) in a selfservice manner.
//
// For more information head over to: https://www.ory.sh/docs/kratos/selfservice/profile
// swagger:model profileManagementRequest
type ProfileManagementRequest struct {

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// form
	Form *Form `json:"form,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// identity
	Identity *Identity `json:"identity,omitempty"`

	// issued at
	// Format: date-time
	IssuedAt strfmt.DateTime `json:"issued_at,omitempty"`

	// request url
	RequestURL string `json:"request_url,omitempty"`
}

// Validate validates this profile management request
func (m *ProfileManagementRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProfileManagementRequest) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProfileManagementRequest) validateForm(formats strfmt.Registry) error {

	if swag.IsZero(m.Form) { // not required
		return nil
	}

	if m.Form != nil {
		if err := m.Form.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form")
			}
			return err
		}
	}

	return nil
}

func (m *ProfileManagementRequest) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *ProfileManagementRequest) validateIssuedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.IssuedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("issued_at", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProfileManagementRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProfileManagementRequest) UnmarshalBinary(b []byte) error {
	var res ProfileManagementRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
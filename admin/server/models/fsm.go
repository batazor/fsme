// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Fsm fsm
// swagger:model fsm
type Fsm struct {

	// callbacks
	Callbacks string `json:"callbacks,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// events
	Events string `json:"events,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// rules
	// Required: true
	Rules *string `json:"rules"`

	// state
	State string `json:"state,omitempty"`

	ClientDB *mongo.Client
}

// Validate validates this fsm
func (m *Fsm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fsm) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Fsm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Fsm) UnmarshalBinary(b []byte) error {
	var res Fsm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KeyValuePairObject Individual key-value pair
// swagger:model KeyValuePairObject
type KeyValuePairObject struct {

	// Name of the data
	KvpKey string `json:"kvpKey,omitempty"`

	// kvp type
	KvpType EnumKVPType `json:"kvpType,omitempty"`

	// Value of the data
	KvpValue string `json:"kvpValue,omitempty"`
}

// Validate validates this key value pair object
func (m *KeyValuePairObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKvpType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyValuePairObject) validateKvpType(formats strfmt.Registry) error {

	if swag.IsZero(m.KvpType) { // not required
		return nil
	}

	if err := m.KvpType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kvpType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyValuePairObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyValuePairObject) UnmarshalBinary(b []byte) error {
	var res KeyValuePairObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TypeInstanceID type_instance_id
//
// A type instance ID
// swagger:model type_instance_id
type TypeInstanceID struct {

	// The plugin ID
	Plugin string `json:"plugin,omitempty"`

	// The plugin instance
	PluginInstance string `json:"plugin_instance,omitempty"`

	// The plugin type
	Type string `json:"type,omitempty"`

	// The plugin type instance
	TypeInstance string `json:"type_instance,omitempty"`
}

// Validate validates this type instance id
func (m *TypeInstanceID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypeInstanceID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypeInstanceID) UnmarshalBinary(b []byte) error {
	var res TypeInstanceID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/M15t/go-swagger/examples/external-types/fred"
	alternate "github.com/M15t/go-swagger/examples/external-types/fred"
)

// MyReaderObject This object demonstrates several ways to refer to an external interface (here assumed akin to io.Reader).
// MarshalBinary() methods are generated. No validation is expected on binary format.
//
// swagger:model MyReaderObject
type MyReaderObject struct {

	// reader1
	// Format: binary
	Reader1 MyStreamer `json:"reader1,omitempty"`

	// In line definition of the external type.
	//
	// Notice that we have provided some information in the spec, so the generator
	// can infer we want it to be understood as an io.Reader, with no validation.
	//
	// Format: binary
	Reader2 alternate.MyAlternateStreamer `json:"reader2,omitempty"`

	// In line definition of the external type.
	//
	// Notice that we have provided some information in the spec, as a hint in the extension
	// rather than in the type definition.
	//
	// So this will be documented as an object, but the generated code knows this is a stream.
	//
	Reader3 fred.MyAlternateStreamer `json:"reader3,omitempty"`
}

// Validate validates this my reader object
func (m *MyReaderObject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this my reader object based on context it is used
func (m *MyReaderObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MyReaderObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MyReaderObject) UnmarshalBinary(b []byte) error {
	var res MyReaderObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

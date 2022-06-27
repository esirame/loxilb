// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoadbalanceEntry loadbalance entry
//
// swagger:model LoadbalanceEntry
type LoadbalanceEntry struct {

	// values of End point servers
	EndPoint []*LoadbalanceEntryEndPointItems0 `json:"end_point"`

	// IP address for externel access
	ExternelIPAddress string `json:"externel_ip_address,omitempty"`

	// port number for the access
	Port int64 `json:"port,omitempty"`

	// value for access protocol
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this loadbalance entry
func (m *LoadbalanceEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadbalanceEntry) validateEndPoint(formats strfmt.Registry) error {
	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPoint); i++ {
		if swag.IsZero(m.EndPoint[i]) { // not required
			continue
		}

		if m.EndPoint[i] != nil {
			if err := m.EndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("end_point" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("end_point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this loadbalance entry based on the context it is used
func (m *LoadbalanceEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndPoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadbalanceEntry) contextValidateEndPoint(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EndPoint); i++ {

		if m.EndPoint[i] != nil {
			if err := m.EndPoint[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("end_point" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("end_point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadbalanceEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadbalanceEntry) UnmarshalBinary(b []byte) error {
	var res LoadbalanceEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadbalanceEntryEndPointItems0 loadbalance entry end point items0
//
// swagger:model LoadbalanceEntryEndPointItems0
type LoadbalanceEntryEndPointItems0 struct {

	// IP address for externel access
	EndpointIPAddress string `json:"endpoint_ip_address,omitempty"`

	// port number for access service
	Port int64 `json:"port,omitempty"`

	// Weight for the load balancing
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this loadbalance entry end point items0
func (m *LoadbalanceEntryEndPointItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this loadbalance entry end point items0 based on context it is used
func (m *LoadbalanceEntryEndPointItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoadbalanceEntryEndPointItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadbalanceEntryEndPointItems0) UnmarshalBinary(b []byte) error {
	var res LoadbalanceEntryEndPointItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
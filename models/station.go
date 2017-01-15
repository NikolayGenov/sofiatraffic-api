package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Station station
// swagger:model Station
type Station struct {

	// The id that Urban Mobility Center uses - usually 4 digits
	UMCID float64 `json:"UMC_id,omitempty"`

	// Name of the station. Note that there could be multiple stations with that name
	Name string `json:"name,omitempty"`
}

// Validate validates this station
func (m *Station) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

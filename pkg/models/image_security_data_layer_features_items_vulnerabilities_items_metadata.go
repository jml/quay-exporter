// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata image security data layer features items vulnerabilities items metadata
// swagger:model imageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata
type ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata struct {

	// n v d
	NVD *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVD `json:"NVD,omitempty"`
}

// Validate validates this image security data layer features items vulnerabilities items metadata
func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNVD(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata) validateNVD(formats strfmt.Registry) error {

	if swag.IsZero(m.NVD) { // not required
		return nil
	}

	if m.NVD != nil {

		if err := m.NVD.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NVD")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata) UnmarshalBinary(b []byte) error {
	var res ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package visualization_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1beta1VisualizationType Type of visualization to be generated.
// This is required when creating the pipeline through CreateVisualization
// API.
// swagger:model v1beta1VisualizationType
type V1beta1VisualizationType string

const (

	// V1beta1VisualizationTypeROCCURVE captures enum value "ROC_CURVE"
	V1beta1VisualizationTypeROCCURVE V1beta1VisualizationType = "ROC_CURVE"

	// V1beta1VisualizationTypeTFDV captures enum value "TFDV"
	V1beta1VisualizationTypeTFDV V1beta1VisualizationType = "TFDV"

	// V1beta1VisualizationTypeTFMA captures enum value "TFMA"
	V1beta1VisualizationTypeTFMA V1beta1VisualizationType = "TFMA"

	// V1beta1VisualizationTypeTABLE captures enum value "TABLE"
	V1beta1VisualizationTypeTABLE V1beta1VisualizationType = "TABLE"

	// V1beta1VisualizationTypeCUSTOM captures enum value "CUSTOM"
	V1beta1VisualizationTypeCUSTOM V1beta1VisualizationType = "CUSTOM"
)

// for schema
var v1beta1VisualizationTypeEnum []interface{}

func init() {
	var res []V1beta1VisualizationType
	if err := json.Unmarshal([]byte(`["ROC_CURVE","TFDV","TFMA","TABLE","CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1beta1VisualizationTypeEnum = append(v1beta1VisualizationTypeEnum, v)
	}
}

func (m V1beta1VisualizationType) validateV1beta1VisualizationTypeEnum(path, location string, value V1beta1VisualizationType) error {
	if err := validate.Enum(path, location, value, v1beta1VisualizationTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1beta1 visualization type
func (m V1beta1VisualizationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1beta1VisualizationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

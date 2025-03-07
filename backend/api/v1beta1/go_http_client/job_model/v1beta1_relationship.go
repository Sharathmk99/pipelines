// Code generated by go-swagger; DO NOT EDIT.

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1beta1Relationship v1beta1 relationship
// swagger:model v1beta1Relationship
type V1beta1Relationship string

const (

	// V1beta1RelationshipUNKNOWNRELATIONSHIP captures enum value "UNKNOWN_RELATIONSHIP"
	V1beta1RelationshipUNKNOWNRELATIONSHIP V1beta1Relationship = "UNKNOWN_RELATIONSHIP"

	// V1beta1RelationshipOWNER captures enum value "OWNER"
	V1beta1RelationshipOWNER V1beta1Relationship = "OWNER"

	// V1beta1RelationshipCREATOR captures enum value "CREATOR"
	V1beta1RelationshipCREATOR V1beta1Relationship = "CREATOR"
)

// for schema
var v1beta1RelationshipEnum []interface{}

func init() {
	var res []V1beta1Relationship
	if err := json.Unmarshal([]byte(`["UNKNOWN_RELATIONSHIP","OWNER","CREATOR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1beta1RelationshipEnum = append(v1beta1RelationshipEnum, v)
	}
}

func (m V1beta1Relationship) validateV1beta1RelationshipEnum(path, location string, value V1beta1Relationship) error {
	if err := validate.Enum(path, location, value, v1beta1RelationshipEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1beta1 relationship
func (m V1beta1Relationship) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1beta1RelationshipEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

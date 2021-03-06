// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1WidgetSpec Widget spec definition
// swagger:model v1WidgetSpec
type V1WidgetSpec struct {

	// Widget kind
	Kind string `json:"kind,omitempty"`

	// Meta inforation
	Meta interface{} `json:"meta,omitempty"`

	// Search spec
	Search *V1SearchSpec `json:"search,omitempty"`
}

// Validate validates this v1 widget spec
func (m *V1WidgetSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WidgetSpec) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Search) { // not required
		return nil
	}

	if m.Search != nil {
		if err := m.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WidgetSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WidgetSpec) UnmarshalBinary(b []byte) error {
	var res V1WidgetSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

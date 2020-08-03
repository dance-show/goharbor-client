// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Project project
//
// swagger:model Project
type Project struct {

	// The total number of charts under this project.
	ChartCount int64 `json:"chart_count,omitempty"`

	// The creation time of the project.
	CreationTime string `json:"creation_time,omitempty"`

	// The role ID with highest permission of the current user who triggered the API (for UI)
	CurrentUserRoleID int64 `json:"current_user_role_id,omitempty"`

	// The list of role ID of the current user who triggered the API (for UI)
	CurrentUserRoleIds []int32 `json:"current_user_role_ids"`

	// The CVE whitelist of this project.
	CveWhitelist *CVEWhitelist `json:"cve_whitelist,omitempty"`

	// A deletion mark of the project.
	Deleted bool `json:"deleted,omitempty"`

	// The metadata of the project.
	Metadata *ProjectMetadata `json:"metadata,omitempty"`

	// The name of the project.
	Name string `json:"name,omitempty"`

	// The owner ID of the project always means the creator of the project.
	OwnerID int32 `json:"owner_id,omitempty"`

	// The owner name of the project.
	OwnerName string `json:"owner_name,omitempty"`

	// Project ID
	ProjectID int32 `json:"project_id,omitempty"`

	// The number of the repositories under this project.
	RepoCount int64 `json:"repo_count,omitempty"`

	// Correspond to the UI about whether the project's publicity is  updatable (for UI)
	Togglable bool `json:"togglable,omitempty"`

	// The update time of the project.
	UpdateTime string `json:"update_time,omitempty"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCveWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) validateCveWhitelist(formats strfmt.Registry) error {
	if swag.IsZero(m.CveWhitelist) { // not required
		return nil
	}

	if m.CveWhitelist != nil {
		if err := m.CveWhitelist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cve_whitelist")
			}
			return err
		}
	}

	return nil
}

func (m *Project) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

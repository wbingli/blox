package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*TaskDetailContainerModel task detail container model

swagger:model TaskDetailContainerModel
*/
type TaskDetailContainerModel struct {

	/* container arn

	Required: true
	*/
	ContainerArn *string `json:"containerArn"`

	/* exit code
	 */
	ExitCode int32 `json:"exitCode,omitempty"`

	/* last status

	Required: true
	*/
	LastStatus *string `json:"lastStatus"`

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* network bindings
	 */
	NetworkBindings []*TaskDetailNetworkBindingModel `json:"networkBindings,omitempty"`

	/* reason
	 */
	Reason string `json:"reason,omitempty"`
}

// Validate validates this task detail container model
func (m *TaskDetailContainerModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerArn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkBindings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDetailContainerModel) validateContainerArn(formats strfmt.Registry) error {

	if err := validate.Required("containerArn", "body", m.ContainerArn); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailContainerModel) validateLastStatus(formats strfmt.Registry) error {

	if err := validate.Required("lastStatus", "body", m.LastStatus); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailContainerModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailContainerModel) validateNetworkBindings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkBindings) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkBindings); i++ {

		if swag.IsZero(m.NetworkBindings[i]) { // not required
			continue
		}

		if m.NetworkBindings[i] != nil {

			if err := m.NetworkBindings[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

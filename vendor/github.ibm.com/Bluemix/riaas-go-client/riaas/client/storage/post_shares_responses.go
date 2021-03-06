// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostSharesReader is a Reader for the PostShares structure.
type PostSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSharesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostSharesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSharesCreated creates a PostSharesCreated with default headers values
func NewPostSharesCreated() *PostSharesCreated {
	return &PostSharesCreated{}
}

/*PostSharesCreated handles this case with default header values.

dummy
*/
type PostSharesCreated struct {
	Payload *models.Share
}

func (o *PostSharesCreated) Error() string {
	return fmt.Sprintf("[POST /shares][%d] postSharesCreated  %+v", 201, o.Payload)
}

func (o *PostSharesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Share)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSharesNotFound creates a PostSharesNotFound with default headers values
func NewPostSharesNotFound() *PostSharesNotFound {
	return &PostSharesNotFound{}
}

/*PostSharesNotFound handles this case with default header values.

error
*/
type PostSharesNotFound struct {
	Payload *models.Riaaserror
}

func (o *PostSharesNotFound) Error() string {
	return fmt.Sprintf("[POST /shares][%d] postSharesNotFound  %+v", 404, o.Payload)
}

func (o *PostSharesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostSharesBody post shares body
swagger:model PostSharesBody
*/
type PostSharesBody struct {

	// The user-defined name for this file share
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// profile
	Profile *PostSharesParamsBodyProfile `json:"profile,omitempty"`

	// resource group
	ResourceGroup *PostSharesParamsBodyResourceGroup `json:"resource_group,omitempty"`

	// Mount targets for the file share
	Targets []*models.ShareTargetWithZoneTemplate `json:"targets"`
}

// Validate validates this post shares body
func (o *PostSharesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSharesBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("ShareTemplate"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (o *PostSharesBody) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(o.Profile) { // not required
		return nil
	}

	if o.Profile != nil {
		if err := o.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShareTemplate" + "." + "profile")
			}
			return err
		}
	}

	return nil
}

func (o *PostSharesBody) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(o.ResourceGroup) { // not required
		return nil
	}

	if o.ResourceGroup != nil {
		if err := o.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShareTemplate" + "." + "resource_group")
			}
			return err
		}
	}

	return nil
}

func (o *PostSharesBody) validateTargets(formats strfmt.Registry) error {

	if swag.IsZero(o.Targets) { // not required
		return nil
	}

	for i := 0; i < len(o.Targets); i++ {
		if swag.IsZero(o.Targets[i]) { // not required
			continue
		}

		if o.Targets[i] != nil {
			if err := o.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ShareTemplate" + "." + "targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostSharesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSharesBody) UnmarshalBinary(b []byte) error {
	var res PostSharesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostSharesParamsBodyProfile profile
swagger:model PostSharesParamsBodyProfile
*/
type PostSharesParamsBodyProfile struct {

	// The CRN for this share profile
	Crn string `json:"crn,omitempty"`

	// The name for this share profile
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this post shares params body profile
func (o *PostSharesParamsBodyProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSharesParamsBodyProfile) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("ShareTemplate"+"."+"profile"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostSharesParamsBodyProfile) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSharesParamsBodyProfile) UnmarshalBinary(b []byte) error {
	var res PostSharesParamsBodyProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostSharesParamsBodyResourceGroup ResourceGroupIdentity
swagger:model PostSharesParamsBodyResourceGroup
*/
type PostSharesParamsBodyResourceGroup struct {

	// The unique identifier for this resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post shares params body resource group
func (o *PostSharesParamsBodyResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSharesParamsBodyResourceGroup) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ShareTemplate"+"."+"resource_group"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostSharesParamsBodyResourceGroup) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSharesParamsBodyResourceGroup) UnmarshalBinary(b []byte) error {
	var res PostSharesParamsBodyResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

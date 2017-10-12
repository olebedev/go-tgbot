// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// GetUserProfilePhotosReader is a Reader for the GetUserProfilePhotos structure.
type GetUserProfilePhotosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserProfilePhotosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserProfilePhotosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserProfilePhotosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUserProfilePhotosUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserProfilePhotosForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserProfilePhotosNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUserProfilePhotosEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUserProfilePhotosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserProfilePhotosOK creates a GetUserProfilePhotosOK with default headers values
func NewGetUserProfilePhotosOK() *GetUserProfilePhotosOK {
	return &GetUserProfilePhotosOK{}
}

/*GetUserProfilePhotosOK handles this case with default header values.

GetUserProfilePhotosOK get user profile photos o k
*/
type GetUserProfilePhotosOK struct {
	Payload GetUserProfilePhotosOKBody
}

func (o *GetUserProfilePhotosOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosOK  %+v", 200, o.Payload)
}

func (o *GetUserProfilePhotosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfilePhotosBadRequest creates a GetUserProfilePhotosBadRequest with default headers values
func NewGetUserProfilePhotosBadRequest() *GetUserProfilePhotosBadRequest {
	return &GetUserProfilePhotosBadRequest{}
}

/*GetUserProfilePhotosBadRequest handles this case with default header values.

Bad Request
*/
type GetUserProfilePhotosBadRequest struct {
	Payload *models.Error
}

func (o *GetUserProfilePhotosBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserProfilePhotosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfilePhotosUnauthorized creates a GetUserProfilePhotosUnauthorized with default headers values
func NewGetUserProfilePhotosUnauthorized() *GetUserProfilePhotosUnauthorized {
	return &GetUserProfilePhotosUnauthorized{}
}

/*GetUserProfilePhotosUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUserProfilePhotosUnauthorized struct {
	Payload *models.Error
}

func (o *GetUserProfilePhotosUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserProfilePhotosUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfilePhotosForbidden creates a GetUserProfilePhotosForbidden with default headers values
func NewGetUserProfilePhotosForbidden() *GetUserProfilePhotosForbidden {
	return &GetUserProfilePhotosForbidden{}
}

/*GetUserProfilePhotosForbidden handles this case with default header values.

Forbidden
*/
type GetUserProfilePhotosForbidden struct {
	Payload *models.Error
}

func (o *GetUserProfilePhotosForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosForbidden  %+v", 403, o.Payload)
}

func (o *GetUserProfilePhotosForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfilePhotosNotFound creates a GetUserProfilePhotosNotFound with default headers values
func NewGetUserProfilePhotosNotFound() *GetUserProfilePhotosNotFound {
	return &GetUserProfilePhotosNotFound{}
}

/*GetUserProfilePhotosNotFound handles this case with default header values.

Not Found
*/
type GetUserProfilePhotosNotFound struct {
	Payload *models.Error
}

func (o *GetUserProfilePhotosNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosNotFound  %+v", 404, o.Payload)
}

func (o *GetUserProfilePhotosNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfilePhotosEnhanceYourCalm creates a GetUserProfilePhotosEnhanceYourCalm with default headers values
func NewGetUserProfilePhotosEnhanceYourCalm() *GetUserProfilePhotosEnhanceYourCalm {
	return &GetUserProfilePhotosEnhanceYourCalm{}
}

/*GetUserProfilePhotosEnhanceYourCalm handles this case with default header values.

Flood
*/
type GetUserProfilePhotosEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *GetUserProfilePhotosEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUserProfilePhotosEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfilePhotosInternalServerError creates a GetUserProfilePhotosInternalServerError with default headers values
func NewGetUserProfilePhotosInternalServerError() *GetUserProfilePhotosInternalServerError {
	return &GetUserProfilePhotosInternalServerError{}
}

/*GetUserProfilePhotosInternalServerError handles this case with default header values.

Internal
*/
type GetUserProfilePhotosInternalServerError struct {
	Payload *models.Error
}

func (o *GetUserProfilePhotosInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getUserProfilePhotos][%d] getUserProfilePhotosInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserProfilePhotosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUserProfilePhotosOKBody get user profile photos o k body
swagger:model GetUserProfilePhotosOKBody
*/

type GetUserProfilePhotosOKBody struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// error code
	// Required: true
	ErrorCode *int64 `json:"error_code"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// result
	// Required: true
	Result *models.UserProfilePhotos `json:"result"`
}

/* polymorph GetUserProfilePhotosOKBody description false */

/* polymorph GetUserProfilePhotosOKBody error_code false */

/* polymorph GetUserProfilePhotosOKBody ok false */

/* polymorph GetUserProfilePhotosOKBody result false */

// Validate validates this get user profile photos o k body
func (o *GetUserProfilePhotosOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateErrorCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserProfilePhotosOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getUserProfilePhotosOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUserProfilePhotosOKBody) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("getUserProfilePhotosOK"+"."+"error_code", "body", o.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (o *GetUserProfilePhotosOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("getUserProfilePhotosOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *GetUserProfilePhotosOKBody) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("getUserProfilePhotosOK"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	if o.Result != nil {

		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUserProfilePhotosOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserProfilePhotosOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserProfilePhotosOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserProfilePhotosOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

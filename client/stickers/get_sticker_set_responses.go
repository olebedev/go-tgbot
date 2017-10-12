// Code generated by go-swagger; DO NOT EDIT.

package stickers

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

// GetStickerSetReader is a Reader for the GetStickerSet structure.
type GetStickerSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStickerSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStickerSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetStickerSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetStickerSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetStickerSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetStickerSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetStickerSetEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetStickerSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStickerSetOK creates a GetStickerSetOK with default headers values
func NewGetStickerSetOK() *GetStickerSetOK {
	return &GetStickerSetOK{}
}

/*GetStickerSetOK handles this case with default header values.

GetStickerSetOK get sticker set o k
*/
type GetStickerSetOK struct {
	Payload GetStickerSetOKBody
}

func (o *GetStickerSetOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetOK  %+v", 200, o.Payload)
}

func (o *GetStickerSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickerSetBadRequest creates a GetStickerSetBadRequest with default headers values
func NewGetStickerSetBadRequest() *GetStickerSetBadRequest {
	return &GetStickerSetBadRequest{}
}

/*GetStickerSetBadRequest handles this case with default header values.

Bad Request
*/
type GetStickerSetBadRequest struct {
	Payload *models.Error
}

func (o *GetStickerSetBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetBadRequest  %+v", 400, o.Payload)
}

func (o *GetStickerSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickerSetUnauthorized creates a GetStickerSetUnauthorized with default headers values
func NewGetStickerSetUnauthorized() *GetStickerSetUnauthorized {
	return &GetStickerSetUnauthorized{}
}

/*GetStickerSetUnauthorized handles this case with default header values.

Unauthorized
*/
type GetStickerSetUnauthorized struct {
	Payload *models.Error
}

func (o *GetStickerSetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStickerSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickerSetForbidden creates a GetStickerSetForbidden with default headers values
func NewGetStickerSetForbidden() *GetStickerSetForbidden {
	return &GetStickerSetForbidden{}
}

/*GetStickerSetForbidden handles this case with default header values.

Forbidden
*/
type GetStickerSetForbidden struct {
	Payload *models.Error
}

func (o *GetStickerSetForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetForbidden  %+v", 403, o.Payload)
}

func (o *GetStickerSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickerSetNotFound creates a GetStickerSetNotFound with default headers values
func NewGetStickerSetNotFound() *GetStickerSetNotFound {
	return &GetStickerSetNotFound{}
}

/*GetStickerSetNotFound handles this case with default header values.

Not Found
*/
type GetStickerSetNotFound struct {
	Payload *models.Error
}

func (o *GetStickerSetNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetNotFound  %+v", 404, o.Payload)
}

func (o *GetStickerSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickerSetEnhanceYourCalm creates a GetStickerSetEnhanceYourCalm with default headers values
func NewGetStickerSetEnhanceYourCalm() *GetStickerSetEnhanceYourCalm {
	return &GetStickerSetEnhanceYourCalm{}
}

/*GetStickerSetEnhanceYourCalm handles this case with default header values.

Flood
*/
type GetStickerSetEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *GetStickerSetEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetStickerSetEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickerSetInternalServerError creates a GetStickerSetInternalServerError with default headers values
func NewGetStickerSetInternalServerError() *GetStickerSetInternalServerError {
	return &GetStickerSetInternalServerError{}
}

/*GetStickerSetInternalServerError handles this case with default header values.

Internal
*/
type GetStickerSetInternalServerError struct {
	Payload *models.Error
}

func (o *GetStickerSetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getStickerSet][%d] getStickerSetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStickerSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetStickerSetOKBody get sticker set o k body
swagger:model GetStickerSetOKBody
*/

type GetStickerSetOKBody struct {

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
	Result *models.StickerSet `json:"result"`
}

/* polymorph GetStickerSetOKBody description false */

/* polymorph GetStickerSetOKBody error_code false */

/* polymorph GetStickerSetOKBody ok false */

/* polymorph GetStickerSetOKBody result false */

// Validate validates this get sticker set o k body
func (o *GetStickerSetOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetStickerSetOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getStickerSetOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetStickerSetOKBody) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("getStickerSetOK"+"."+"error_code", "body", o.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (o *GetStickerSetOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("getStickerSetOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *GetStickerSetOKBody) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("getStickerSetOK"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	if o.Result != nil {

		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStickerSetOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStickerSetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStickerSetOKBody) UnmarshalBinary(b []byte) error {
	var res GetStickerSetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

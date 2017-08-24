// Code generated by go-swagger; DO NOT EDIT.

package messages

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

// EditMessageCaptionReader is a Reader for the EditMessageCaption structure.
type EditMessageCaptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditMessageCaptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditMessageCaptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditMessageCaptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEditMessageCaptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEditMessageCaptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEditMessageCaptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewEditMessageCaptionEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEditMessageCaptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditMessageCaptionOK creates a EditMessageCaptionOK with default headers values
func NewEditMessageCaptionOK() *EditMessageCaptionOK {
	return &EditMessageCaptionOK{}
}

/*EditMessageCaptionOK handles this case with default header values.

EditMessageCaptionOK edit message caption o k
*/
type EditMessageCaptionOK struct {
	Payload EditMessageCaptionOKBody
}

func (o *EditMessageCaptionOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionOK  %+v", 200, o.Payload)
}

func (o *EditMessageCaptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionBadRequest creates a EditMessageCaptionBadRequest with default headers values
func NewEditMessageCaptionBadRequest() *EditMessageCaptionBadRequest {
	return &EditMessageCaptionBadRequest{}
}

/*EditMessageCaptionBadRequest handles this case with default header values.

Bad Request
*/
type EditMessageCaptionBadRequest struct {
	Payload *models.Error
}

func (o *EditMessageCaptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionBadRequest  %+v", 400, o.Payload)
}

func (o *EditMessageCaptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionUnauthorized creates a EditMessageCaptionUnauthorized with default headers values
func NewEditMessageCaptionUnauthorized() *EditMessageCaptionUnauthorized {
	return &EditMessageCaptionUnauthorized{}
}

/*EditMessageCaptionUnauthorized handles this case with default header values.

Unauthorized
*/
type EditMessageCaptionUnauthorized struct {
	Payload *models.Error
}

func (o *EditMessageCaptionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionUnauthorized  %+v", 401, o.Payload)
}

func (o *EditMessageCaptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionForbidden creates a EditMessageCaptionForbidden with default headers values
func NewEditMessageCaptionForbidden() *EditMessageCaptionForbidden {
	return &EditMessageCaptionForbidden{}
}

/*EditMessageCaptionForbidden handles this case with default header values.

Forbidden
*/
type EditMessageCaptionForbidden struct {
	Payload *models.Error
}

func (o *EditMessageCaptionForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionForbidden  %+v", 403, o.Payload)
}

func (o *EditMessageCaptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionNotFound creates a EditMessageCaptionNotFound with default headers values
func NewEditMessageCaptionNotFound() *EditMessageCaptionNotFound {
	return &EditMessageCaptionNotFound{}
}

/*EditMessageCaptionNotFound handles this case with default header values.

Not Found
*/
type EditMessageCaptionNotFound struct {
	Payload *models.Error
}

func (o *EditMessageCaptionNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionNotFound  %+v", 404, o.Payload)
}

func (o *EditMessageCaptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionEnhanceYourCalm creates a EditMessageCaptionEnhanceYourCalm with default headers values
func NewEditMessageCaptionEnhanceYourCalm() *EditMessageCaptionEnhanceYourCalm {
	return &EditMessageCaptionEnhanceYourCalm{}
}

/*EditMessageCaptionEnhanceYourCalm handles this case with default header values.

Flood
*/
type EditMessageCaptionEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *EditMessageCaptionEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *EditMessageCaptionEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionInternalServerError creates a EditMessageCaptionInternalServerError with default headers values
func NewEditMessageCaptionInternalServerError() *EditMessageCaptionInternalServerError {
	return &EditMessageCaptionInternalServerError{}
}

/*EditMessageCaptionInternalServerError handles this case with default header values.

Internal
*/
type EditMessageCaptionInternalServerError struct {
	Payload *models.Error
}

func (o *EditMessageCaptionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionInternalServerError  %+v", 500, o.Payload)
}

func (o *EditMessageCaptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EditMessageCaptionOKBody edit message caption o k body
swagger:model EditMessageCaptionOKBody
*/
type EditMessageCaptionOKBody struct {

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
	Result interface{} `json:"result"`
}

// Validate validates this edit message caption o k body
func (o *EditMessageCaptionOKBody) Validate(formats strfmt.Registry) error {
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

func (o *EditMessageCaptionOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("editMessageCaptionOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *EditMessageCaptionOKBody) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("editMessageCaptionOK"+"."+"error_code", "body", o.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (o *EditMessageCaptionOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("editMessageCaptionOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *EditMessageCaptionOKBody) validateResult(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (o *EditMessageCaptionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditMessageCaptionOKBody) UnmarshalBinary(b []byte) error {
	var res EditMessageCaptionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

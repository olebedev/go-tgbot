// Code generated by go-swagger; DO NOT EDIT.

package chats

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

// GetChatMemberReader is a Reader for the GetChatMember structure.
type GetChatMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChatMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChatMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetChatMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetChatMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetChatMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetChatMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetChatMemberEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetChatMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChatMemberOK creates a GetChatMemberOK with default headers values
func NewGetChatMemberOK() *GetChatMemberOK {
	return &GetChatMemberOK{}
}

/*GetChatMemberOK handles this case with default header values.

GetChatMemberOK get chat member o k
*/
type GetChatMemberOK struct {
	Payload GetChatMemberOKBody
}

func (o *GetChatMemberOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberOK  %+v", 200, o.Payload)
}

func (o *GetChatMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMemberBadRequest creates a GetChatMemberBadRequest with default headers values
func NewGetChatMemberBadRequest() *GetChatMemberBadRequest {
	return &GetChatMemberBadRequest{}
}

/*GetChatMemberBadRequest handles this case with default header values.

Bad Request
*/
type GetChatMemberBadRequest struct {
	Payload *models.Error
}

func (o *GetChatMemberBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberBadRequest  %+v", 400, o.Payload)
}

func (o *GetChatMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMemberUnauthorized creates a GetChatMemberUnauthorized with default headers values
func NewGetChatMemberUnauthorized() *GetChatMemberUnauthorized {
	return &GetChatMemberUnauthorized{}
}

/*GetChatMemberUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChatMemberUnauthorized struct {
	Payload *models.Error
}

func (o *GetChatMemberUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChatMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMemberForbidden creates a GetChatMemberForbidden with default headers values
func NewGetChatMemberForbidden() *GetChatMemberForbidden {
	return &GetChatMemberForbidden{}
}

/*GetChatMemberForbidden handles this case with default header values.

Forbidden
*/
type GetChatMemberForbidden struct {
	Payload *models.Error
}

func (o *GetChatMemberForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetChatMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMemberNotFound creates a GetChatMemberNotFound with default headers values
func NewGetChatMemberNotFound() *GetChatMemberNotFound {
	return &GetChatMemberNotFound{}
}

/*GetChatMemberNotFound handles this case with default header values.

Not Found
*/
type GetChatMemberNotFound struct {
	Payload *models.Error
}

func (o *GetChatMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetChatMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMemberEnhanceYourCalm creates a GetChatMemberEnhanceYourCalm with default headers values
func NewGetChatMemberEnhanceYourCalm() *GetChatMemberEnhanceYourCalm {
	return &GetChatMemberEnhanceYourCalm{}
}

/*GetChatMemberEnhanceYourCalm handles this case with default header values.

Flood
*/
type GetChatMemberEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *GetChatMemberEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetChatMemberEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMemberInternalServerError creates a GetChatMemberInternalServerError with default headers values
func NewGetChatMemberInternalServerError() *GetChatMemberInternalServerError {
	return &GetChatMemberInternalServerError{}
}

/*GetChatMemberInternalServerError handles this case with default header values.

Internal
*/
type GetChatMemberInternalServerError struct {
	Payload *models.Error
}

func (o *GetChatMemberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMember][%d] getChatMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChatMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetChatMemberOKBody get chat member o k body
swagger:model GetChatMemberOKBody
*/
type GetChatMemberOKBody struct {

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
	Result *models.ChatMember `json:"result"`
}

// Validate validates this get chat member o k body
func (o *GetChatMemberOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetChatMemberOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getChatMemberOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetChatMemberOKBody) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("getChatMemberOK"+"."+"error_code", "body", o.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (o *GetChatMemberOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("getChatMemberOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *GetChatMemberOKBody) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("getChatMemberOK"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	if o.Result != nil {

		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getChatMemberOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetChatMemberOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChatMemberOKBody) UnmarshalBinary(b []byte) error {
	var res GetChatMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

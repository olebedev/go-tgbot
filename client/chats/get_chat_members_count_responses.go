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

// GetChatMembersCountReader is a Reader for the GetChatMembersCount structure.
type GetChatMembersCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChatMembersCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChatMembersCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetChatMembersCountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetChatMembersCountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetChatMembersCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetChatMembersCountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetChatMembersCountEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetChatMembersCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChatMembersCountOK creates a GetChatMembersCountOK with default headers values
func NewGetChatMembersCountOK() *GetChatMembersCountOK {
	return &GetChatMembersCountOK{}
}

/*GetChatMembersCountOK handles this case with default header values.

GetChatMembersCountOK get chat members count o k
*/
type GetChatMembersCountOK struct {
	Payload GetChatMembersCountOKBody
}

func (o *GetChatMembersCountOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountOK  %+v", 200, o.Payload)
}

func (o *GetChatMembersCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMembersCountBadRequest creates a GetChatMembersCountBadRequest with default headers values
func NewGetChatMembersCountBadRequest() *GetChatMembersCountBadRequest {
	return &GetChatMembersCountBadRequest{}
}

/*GetChatMembersCountBadRequest handles this case with default header values.

Bad Request
*/
type GetChatMembersCountBadRequest struct {
	Payload *models.Error
}

func (o *GetChatMembersCountBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountBadRequest  %+v", 400, o.Payload)
}

func (o *GetChatMembersCountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMembersCountUnauthorized creates a GetChatMembersCountUnauthorized with default headers values
func NewGetChatMembersCountUnauthorized() *GetChatMembersCountUnauthorized {
	return &GetChatMembersCountUnauthorized{}
}

/*GetChatMembersCountUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChatMembersCountUnauthorized struct {
	Payload *models.Error
}

func (o *GetChatMembersCountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChatMembersCountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMembersCountForbidden creates a GetChatMembersCountForbidden with default headers values
func NewGetChatMembersCountForbidden() *GetChatMembersCountForbidden {
	return &GetChatMembersCountForbidden{}
}

/*GetChatMembersCountForbidden handles this case with default header values.

Forbidden
*/
type GetChatMembersCountForbidden struct {
	Payload *models.Error
}

func (o *GetChatMembersCountForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountForbidden  %+v", 403, o.Payload)
}

func (o *GetChatMembersCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMembersCountNotFound creates a GetChatMembersCountNotFound with default headers values
func NewGetChatMembersCountNotFound() *GetChatMembersCountNotFound {
	return &GetChatMembersCountNotFound{}
}

/*GetChatMembersCountNotFound handles this case with default header values.

Not Found
*/
type GetChatMembersCountNotFound struct {
	Payload *models.Error
}

func (o *GetChatMembersCountNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountNotFound  %+v", 404, o.Payload)
}

func (o *GetChatMembersCountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMembersCountEnhanceYourCalm creates a GetChatMembersCountEnhanceYourCalm with default headers values
func NewGetChatMembersCountEnhanceYourCalm() *GetChatMembersCountEnhanceYourCalm {
	return &GetChatMembersCountEnhanceYourCalm{}
}

/*GetChatMembersCountEnhanceYourCalm handles this case with default header values.

Flood
*/
type GetChatMembersCountEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *GetChatMembersCountEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetChatMembersCountEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatMembersCountInternalServerError creates a GetChatMembersCountInternalServerError with default headers values
func NewGetChatMembersCountInternalServerError() *GetChatMembersCountInternalServerError {
	return &GetChatMembersCountInternalServerError{}
}

/*GetChatMembersCountInternalServerError handles this case with default header values.

Internal
*/
type GetChatMembersCountInternalServerError struct {
	Payload *models.Error
}

func (o *GetChatMembersCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChatMembersCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetChatMembersCountOKBody get chat members count o k body
swagger:model GetChatMembersCountOKBody
*/
type GetChatMembersCountOKBody struct {

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
	Result *int64 `json:"result"`
}

// Validate validates this get chat members count o k body
func (o *GetChatMembersCountOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetChatMembersCountOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getChatMembersCountOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetChatMembersCountOKBody) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("getChatMembersCountOK"+"."+"error_code", "body", o.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (o *GetChatMembersCountOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("getChatMembersCountOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *GetChatMembersCountOKBody) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("getChatMembersCountOK"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetChatMembersCountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChatMembersCountOKBody) UnmarshalBinary(b []byte) error {
	var res GetChatMembersCountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

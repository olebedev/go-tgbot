package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// LeaveChatReader is a Reader for the LeaveChat structure.
type LeaveChatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LeaveChatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLeaveChatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLeaveChatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewLeaveChatUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewLeaveChatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLeaveChatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewLeaveChatEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLeaveChatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLeaveChatOK creates a LeaveChatOK with default headers values
func NewLeaveChatOK() *LeaveChatOK {
	return &LeaveChatOK{}
}

/*LeaveChatOK handles this case with default header values.

LeaveChatOK leave chat o k
*/
type LeaveChatOK struct {
	Payload LeaveChatOKBody
}

func (o *LeaveChatOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatOK  %+v", 200, o.Payload)
}

func (o *LeaveChatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveChatBadRequest creates a LeaveChatBadRequest with default headers values
func NewLeaveChatBadRequest() *LeaveChatBadRequest {
	return &LeaveChatBadRequest{}
}

/*LeaveChatBadRequest handles this case with default header values.

Bad Request
*/
type LeaveChatBadRequest struct {
	Payload *models.Error
}

func (o *LeaveChatBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatBadRequest  %+v", 400, o.Payload)
}

func (o *LeaveChatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveChatUnauthorized creates a LeaveChatUnauthorized with default headers values
func NewLeaveChatUnauthorized() *LeaveChatUnauthorized {
	return &LeaveChatUnauthorized{}
}

/*LeaveChatUnauthorized handles this case with default header values.

Unauthorized
*/
type LeaveChatUnauthorized struct {
	Payload *models.Error
}

func (o *LeaveChatUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatUnauthorized  %+v", 401, o.Payload)
}

func (o *LeaveChatUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveChatForbidden creates a LeaveChatForbidden with default headers values
func NewLeaveChatForbidden() *LeaveChatForbidden {
	return &LeaveChatForbidden{}
}

/*LeaveChatForbidden handles this case with default header values.

Forbidden
*/
type LeaveChatForbidden struct {
	Payload *models.Error
}

func (o *LeaveChatForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatForbidden  %+v", 403, o.Payload)
}

func (o *LeaveChatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveChatNotFound creates a LeaveChatNotFound with default headers values
func NewLeaveChatNotFound() *LeaveChatNotFound {
	return &LeaveChatNotFound{}
}

/*LeaveChatNotFound handles this case with default header values.

Not Found
*/
type LeaveChatNotFound struct {
	Payload *models.Error
}

func (o *LeaveChatNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatNotFound  %+v", 404, o.Payload)
}

func (o *LeaveChatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveChatEnhanceYourCalm creates a LeaveChatEnhanceYourCalm with default headers values
func NewLeaveChatEnhanceYourCalm() *LeaveChatEnhanceYourCalm {
	return &LeaveChatEnhanceYourCalm{}
}

/*LeaveChatEnhanceYourCalm handles this case with default header values.

Flood
*/
type LeaveChatEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *LeaveChatEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *LeaveChatEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveChatInternalServerError creates a LeaveChatInternalServerError with default headers values
func NewLeaveChatInternalServerError() *LeaveChatInternalServerError {
	return &LeaveChatInternalServerError{}
}

/*LeaveChatInternalServerError handles this case with default header values.

Internal
*/
type LeaveChatInternalServerError struct {
	Payload *models.Error
}

func (o *LeaveChatInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/leaveChat][%d] leaveChatInternalServerError  %+v", 500, o.Payload)
}

func (o *LeaveChatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LeaveChatOKBody leave chat o k body
swagger:model LeaveChatOKBody
*/
type LeaveChatOKBody struct {

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
	Result *bool `json:"result"`
}

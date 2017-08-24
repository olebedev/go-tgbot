// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendStickerReader is a Reader for the SendSticker structure.
type SendStickerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendStickerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendStickerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendStickerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendStickerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendStickerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendStickerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendStickerEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendStickerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendStickerOK creates a SendStickerOK with default headers values
func NewSendStickerOK() *SendStickerOK {
	return &SendStickerOK{}
}

/*SendStickerOK handles this case with default header values.

SendStickerOK send sticker o k
*/
type SendStickerOK struct {
	Payload *models.ResponseMessage
}

func (o *SendStickerOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerOK  %+v", 200, o.Payload)
}

func (o *SendStickerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendStickerBadRequest creates a SendStickerBadRequest with default headers values
func NewSendStickerBadRequest() *SendStickerBadRequest {
	return &SendStickerBadRequest{}
}

/*SendStickerBadRequest handles this case with default header values.

Bad Request
*/
type SendStickerBadRequest struct {
	Payload *models.Error
}

func (o *SendStickerBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerBadRequest  %+v", 400, o.Payload)
}

func (o *SendStickerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendStickerUnauthorized creates a SendStickerUnauthorized with default headers values
func NewSendStickerUnauthorized() *SendStickerUnauthorized {
	return &SendStickerUnauthorized{}
}

/*SendStickerUnauthorized handles this case with default header values.

Unauthorized
*/
type SendStickerUnauthorized struct {
	Payload *models.Error
}

func (o *SendStickerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerUnauthorized  %+v", 401, o.Payload)
}

func (o *SendStickerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendStickerForbidden creates a SendStickerForbidden with default headers values
func NewSendStickerForbidden() *SendStickerForbidden {
	return &SendStickerForbidden{}
}

/*SendStickerForbidden handles this case with default header values.

Forbidden
*/
type SendStickerForbidden struct {
	Payload *models.Error
}

func (o *SendStickerForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerForbidden  %+v", 403, o.Payload)
}

func (o *SendStickerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendStickerNotFound creates a SendStickerNotFound with default headers values
func NewSendStickerNotFound() *SendStickerNotFound {
	return &SendStickerNotFound{}
}

/*SendStickerNotFound handles this case with default header values.

Not Found
*/
type SendStickerNotFound struct {
	Payload *models.Error
}

func (o *SendStickerNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerNotFound  %+v", 404, o.Payload)
}

func (o *SendStickerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendStickerEnhanceYourCalm creates a SendStickerEnhanceYourCalm with default headers values
func NewSendStickerEnhanceYourCalm() *SendStickerEnhanceYourCalm {
	return &SendStickerEnhanceYourCalm{}
}

/*SendStickerEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendStickerEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendStickerEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendStickerEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendStickerInternalServerError creates a SendStickerInternalServerError with default headers values
func NewSendStickerInternalServerError() *SendStickerInternalServerError {
	return &SendStickerInternalServerError{}
}

/*SendStickerInternalServerError handles this case with default header values.

Internal
*/
type SendStickerInternalServerError struct {
	Payload *models.Error
}

func (o *SendStickerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendSticker][%d] sendStickerInternalServerError  %+v", 500, o.Payload)
}

func (o *SendStickerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

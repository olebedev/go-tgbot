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

// SendVideoReader is a Reader for the SendVideo structure.
type SendVideoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVideoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVideoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendVideoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendVideoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendVideoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendVideoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendVideoEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendVideoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVideoOK creates a SendVideoOK with default headers values
func NewSendVideoOK() *SendVideoOK {
	return &SendVideoOK{}
}

/*SendVideoOK handles this case with default header values.

SendVideoOK send video o k
*/
type SendVideoOK struct {
	Payload *models.ResponseMessage
}

func (o *SendVideoOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoOK  %+v", 200, o.Payload)
}

func (o *SendVideoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoBadRequest creates a SendVideoBadRequest with default headers values
func NewSendVideoBadRequest() *SendVideoBadRequest {
	return &SendVideoBadRequest{}
}

/*SendVideoBadRequest handles this case with default header values.

Bad Request
*/
type SendVideoBadRequest struct {
	Payload *models.Error
}

func (o *SendVideoBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoBadRequest  %+v", 400, o.Payload)
}

func (o *SendVideoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoUnauthorized creates a SendVideoUnauthorized with default headers values
func NewSendVideoUnauthorized() *SendVideoUnauthorized {
	return &SendVideoUnauthorized{}
}

/*SendVideoUnauthorized handles this case with default header values.

Unauthorized
*/
type SendVideoUnauthorized struct {
	Payload *models.Error
}

func (o *SendVideoUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoUnauthorized  %+v", 401, o.Payload)
}

func (o *SendVideoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoForbidden creates a SendVideoForbidden with default headers values
func NewSendVideoForbidden() *SendVideoForbidden {
	return &SendVideoForbidden{}
}

/*SendVideoForbidden handles this case with default header values.

Forbidden
*/
type SendVideoForbidden struct {
	Payload *models.Error
}

func (o *SendVideoForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoForbidden  %+v", 403, o.Payload)
}

func (o *SendVideoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNotFound creates a SendVideoNotFound with default headers values
func NewSendVideoNotFound() *SendVideoNotFound {
	return &SendVideoNotFound{}
}

/*SendVideoNotFound handles this case with default header values.

Not Found
*/
type SendVideoNotFound struct {
	Payload *models.Error
}

func (o *SendVideoNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoNotFound  %+v", 404, o.Payload)
}

func (o *SendVideoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoEnhanceYourCalm creates a SendVideoEnhanceYourCalm with default headers values
func NewSendVideoEnhanceYourCalm() *SendVideoEnhanceYourCalm {
	return &SendVideoEnhanceYourCalm{}
}

/*SendVideoEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendVideoEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendVideoEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendVideoEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoInternalServerError creates a SendVideoInternalServerError with default headers values
func NewSendVideoInternalServerError() *SendVideoInternalServerError {
	return &SendVideoInternalServerError{}
}

/*SendVideoInternalServerError handles this case with default header values.

Internal
*/
type SendVideoInternalServerError struct {
	Payload *models.Error
}

func (o *SendVideoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoInternalServerError  %+v", 500, o.Payload)
}

func (o *SendVideoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

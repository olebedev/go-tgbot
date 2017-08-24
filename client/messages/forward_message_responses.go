// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// ForwardMessageReader is a Reader for the ForwardMessage structure.
type ForwardMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForwardMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewForwardMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewForwardMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewForwardMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewForwardMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewForwardMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewForwardMessageEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewForwardMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewForwardMessageOK creates a ForwardMessageOK with default headers values
func NewForwardMessageOK() *ForwardMessageOK {
	return &ForwardMessageOK{}
}

/*ForwardMessageOK handles this case with default header values.

ForwardMessageOK forward message o k
*/
type ForwardMessageOK struct {
	Payload *models.ResponseMessage
}

func (o *ForwardMessageOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageOK  %+v", 200, o.Payload)
}

func (o *ForwardMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForwardMessageBadRequest creates a ForwardMessageBadRequest with default headers values
func NewForwardMessageBadRequest() *ForwardMessageBadRequest {
	return &ForwardMessageBadRequest{}
}

/*ForwardMessageBadRequest handles this case with default header values.

Bad Request
*/
type ForwardMessageBadRequest struct {
	Payload *models.Error
}

func (o *ForwardMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageBadRequest  %+v", 400, o.Payload)
}

func (o *ForwardMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForwardMessageUnauthorized creates a ForwardMessageUnauthorized with default headers values
func NewForwardMessageUnauthorized() *ForwardMessageUnauthorized {
	return &ForwardMessageUnauthorized{}
}

/*ForwardMessageUnauthorized handles this case with default header values.

Unauthorized
*/
type ForwardMessageUnauthorized struct {
	Payload *models.Error
}

func (o *ForwardMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *ForwardMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForwardMessageForbidden creates a ForwardMessageForbidden with default headers values
func NewForwardMessageForbidden() *ForwardMessageForbidden {
	return &ForwardMessageForbidden{}
}

/*ForwardMessageForbidden handles this case with default header values.

Forbidden
*/
type ForwardMessageForbidden struct {
	Payload *models.Error
}

func (o *ForwardMessageForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageForbidden  %+v", 403, o.Payload)
}

func (o *ForwardMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForwardMessageNotFound creates a ForwardMessageNotFound with default headers values
func NewForwardMessageNotFound() *ForwardMessageNotFound {
	return &ForwardMessageNotFound{}
}

/*ForwardMessageNotFound handles this case with default header values.

Not Found
*/
type ForwardMessageNotFound struct {
	Payload *models.Error
}

func (o *ForwardMessageNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageNotFound  %+v", 404, o.Payload)
}

func (o *ForwardMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForwardMessageEnhanceYourCalm creates a ForwardMessageEnhanceYourCalm with default headers values
func NewForwardMessageEnhanceYourCalm() *ForwardMessageEnhanceYourCalm {
	return &ForwardMessageEnhanceYourCalm{}
}

/*ForwardMessageEnhanceYourCalm handles this case with default header values.

Flood
*/
type ForwardMessageEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *ForwardMessageEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *ForwardMessageEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForwardMessageInternalServerError creates a ForwardMessageInternalServerError with default headers values
func NewForwardMessageInternalServerError() *ForwardMessageInternalServerError {
	return &ForwardMessageInternalServerError{}
}

/*ForwardMessageInternalServerError handles this case with default header values.

Internal
*/
type ForwardMessageInternalServerError struct {
	Payload *models.Error
}

func (o *ForwardMessageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/forwardMessage][%d] forwardMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *ForwardMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

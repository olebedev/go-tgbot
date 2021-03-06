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

// EditMessageReplyMarkupReader is a Reader for the EditMessageReplyMarkup structure.
type EditMessageReplyMarkupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditMessageReplyMarkupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditMessageReplyMarkupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditMessageReplyMarkupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEditMessageReplyMarkupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEditMessageReplyMarkupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEditMessageReplyMarkupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewEditMessageReplyMarkupEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEditMessageReplyMarkupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditMessageReplyMarkupOK creates a EditMessageReplyMarkupOK with default headers values
func NewEditMessageReplyMarkupOK() *EditMessageReplyMarkupOK {
	return &EditMessageReplyMarkupOK{}
}

/*EditMessageReplyMarkupOK handles this case with default header values.

EditMessageReplyMarkupOK edit message reply markup o k
*/
type EditMessageReplyMarkupOK struct {
	Payload *models.EditMessageReplyMarkupOKBody
}

func (o *EditMessageReplyMarkupOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupOK  %+v", 200, o.Payload)
}

func (o *EditMessageReplyMarkupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EditMessageReplyMarkupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageReplyMarkupBadRequest creates a EditMessageReplyMarkupBadRequest with default headers values
func NewEditMessageReplyMarkupBadRequest() *EditMessageReplyMarkupBadRequest {
	return &EditMessageReplyMarkupBadRequest{}
}

/*EditMessageReplyMarkupBadRequest handles this case with default header values.

Bad Request
*/
type EditMessageReplyMarkupBadRequest struct {
	Payload *models.Error
}

func (o *EditMessageReplyMarkupBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupBadRequest  %+v", 400, o.Payload)
}

func (o *EditMessageReplyMarkupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageReplyMarkupUnauthorized creates a EditMessageReplyMarkupUnauthorized with default headers values
func NewEditMessageReplyMarkupUnauthorized() *EditMessageReplyMarkupUnauthorized {
	return &EditMessageReplyMarkupUnauthorized{}
}

/*EditMessageReplyMarkupUnauthorized handles this case with default header values.

Unauthorized
*/
type EditMessageReplyMarkupUnauthorized struct {
	Payload *models.Error
}

func (o *EditMessageReplyMarkupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupUnauthorized  %+v", 401, o.Payload)
}

func (o *EditMessageReplyMarkupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageReplyMarkupForbidden creates a EditMessageReplyMarkupForbidden with default headers values
func NewEditMessageReplyMarkupForbidden() *EditMessageReplyMarkupForbidden {
	return &EditMessageReplyMarkupForbidden{}
}

/*EditMessageReplyMarkupForbidden handles this case with default header values.

Forbidden
*/
type EditMessageReplyMarkupForbidden struct {
	Payload *models.Error
}

func (o *EditMessageReplyMarkupForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupForbidden  %+v", 403, o.Payload)
}

func (o *EditMessageReplyMarkupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageReplyMarkupNotFound creates a EditMessageReplyMarkupNotFound with default headers values
func NewEditMessageReplyMarkupNotFound() *EditMessageReplyMarkupNotFound {
	return &EditMessageReplyMarkupNotFound{}
}

/*EditMessageReplyMarkupNotFound handles this case with default header values.

Not Found
*/
type EditMessageReplyMarkupNotFound struct {
	Payload *models.Error
}

func (o *EditMessageReplyMarkupNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupNotFound  %+v", 404, o.Payload)
}

func (o *EditMessageReplyMarkupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageReplyMarkupEnhanceYourCalm creates a EditMessageReplyMarkupEnhanceYourCalm with default headers values
func NewEditMessageReplyMarkupEnhanceYourCalm() *EditMessageReplyMarkupEnhanceYourCalm {
	return &EditMessageReplyMarkupEnhanceYourCalm{}
}

/*EditMessageReplyMarkupEnhanceYourCalm handles this case with default header values.

Flood
*/
type EditMessageReplyMarkupEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *EditMessageReplyMarkupEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *EditMessageReplyMarkupEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageReplyMarkupInternalServerError creates a EditMessageReplyMarkupInternalServerError with default headers values
func NewEditMessageReplyMarkupInternalServerError() *EditMessageReplyMarkupInternalServerError {
	return &EditMessageReplyMarkupInternalServerError{}
}

/*EditMessageReplyMarkupInternalServerError handles this case with default header values.

Internal
*/
type EditMessageReplyMarkupInternalServerError struct {
	Payload *models.Error
}

func (o *EditMessageReplyMarkupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageReplyMarkup][%d] editMessageReplyMarkupInternalServerError  %+v", 500, o.Payload)
}

func (o *EditMessageReplyMarkupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package callbacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// AnswerCallbackQueryReader is a Reader for the AnswerCallbackQuery structure.
type AnswerCallbackQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AnswerCallbackQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAnswerCallbackQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAnswerCallbackQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAnswerCallbackQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAnswerCallbackQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAnswerCallbackQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewAnswerCallbackQueryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAnswerCallbackQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAnswerCallbackQueryOK creates a AnswerCallbackQueryOK with default headers values
func NewAnswerCallbackQueryOK() *AnswerCallbackQueryOK {
	return &AnswerCallbackQueryOK{}
}

/*AnswerCallbackQueryOK handles this case with default header values.

AnswerCallbackQueryOK answer callback query o k
*/
type AnswerCallbackQueryOK struct {
	Payload *models.ResponseBool
}

func (o *AnswerCallbackQueryOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryOK  %+v", 200, o.Payload)
}

func (o *AnswerCallbackQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerCallbackQueryBadRequest creates a AnswerCallbackQueryBadRequest with default headers values
func NewAnswerCallbackQueryBadRequest() *AnswerCallbackQueryBadRequest {
	return &AnswerCallbackQueryBadRequest{}
}

/*AnswerCallbackQueryBadRequest handles this case with default header values.

Bad Request
*/
type AnswerCallbackQueryBadRequest struct {
	Payload *models.Error
}

func (o *AnswerCallbackQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryBadRequest  %+v", 400, o.Payload)
}

func (o *AnswerCallbackQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerCallbackQueryUnauthorized creates a AnswerCallbackQueryUnauthorized with default headers values
func NewAnswerCallbackQueryUnauthorized() *AnswerCallbackQueryUnauthorized {
	return &AnswerCallbackQueryUnauthorized{}
}

/*AnswerCallbackQueryUnauthorized handles this case with default header values.

Unauthorized
*/
type AnswerCallbackQueryUnauthorized struct {
	Payload *models.Error
}

func (o *AnswerCallbackQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *AnswerCallbackQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerCallbackQueryForbidden creates a AnswerCallbackQueryForbidden with default headers values
func NewAnswerCallbackQueryForbidden() *AnswerCallbackQueryForbidden {
	return &AnswerCallbackQueryForbidden{}
}

/*AnswerCallbackQueryForbidden handles this case with default header values.

Forbidden
*/
type AnswerCallbackQueryForbidden struct {
	Payload *models.Error
}

func (o *AnswerCallbackQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryForbidden  %+v", 403, o.Payload)
}

func (o *AnswerCallbackQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerCallbackQueryNotFound creates a AnswerCallbackQueryNotFound with default headers values
func NewAnswerCallbackQueryNotFound() *AnswerCallbackQueryNotFound {
	return &AnswerCallbackQueryNotFound{}
}

/*AnswerCallbackQueryNotFound handles this case with default header values.

Not Found
*/
type AnswerCallbackQueryNotFound struct {
	Payload *models.Error
}

func (o *AnswerCallbackQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryNotFound  %+v", 404, o.Payload)
}

func (o *AnswerCallbackQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerCallbackQueryEnhanceYourCalm creates a AnswerCallbackQueryEnhanceYourCalm with default headers values
func NewAnswerCallbackQueryEnhanceYourCalm() *AnswerCallbackQueryEnhanceYourCalm {
	return &AnswerCallbackQueryEnhanceYourCalm{}
}

/*AnswerCallbackQueryEnhanceYourCalm handles this case with default header values.

Flood
*/
type AnswerCallbackQueryEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *AnswerCallbackQueryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *AnswerCallbackQueryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerCallbackQueryInternalServerError creates a AnswerCallbackQueryInternalServerError with default headers values
func NewAnswerCallbackQueryInternalServerError() *AnswerCallbackQueryInternalServerError {
	return &AnswerCallbackQueryInternalServerError{}
}

/*AnswerCallbackQueryInternalServerError handles this case with default header values.

Internal
*/
type AnswerCallbackQueryInternalServerError struct {
	Payload *models.Error
}

func (o *AnswerCallbackQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerCallbackQuery][%d] answerCallbackQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *AnswerCallbackQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

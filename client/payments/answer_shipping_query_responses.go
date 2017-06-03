package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// AnswerShippingQueryReader is a Reader for the AnswerShippingQuery structure.
type AnswerShippingQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AnswerShippingQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAnswerShippingQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAnswerShippingQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAnswerShippingQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAnswerShippingQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAnswerShippingQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewAnswerShippingQueryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAnswerShippingQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAnswerShippingQueryOK creates a AnswerShippingQueryOK with default headers values
func NewAnswerShippingQueryOK() *AnswerShippingQueryOK {
	return &AnswerShippingQueryOK{}
}

/*AnswerShippingQueryOK handles this case with default header values.

AnswerShippingQueryOK answer shipping query o k
*/
type AnswerShippingQueryOK struct {
	Payload *models.ResponseMessage
}

func (o *AnswerShippingQueryOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryOK  %+v", 200, o.Payload)
}

func (o *AnswerShippingQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerShippingQueryBadRequest creates a AnswerShippingQueryBadRequest with default headers values
func NewAnswerShippingQueryBadRequest() *AnswerShippingQueryBadRequest {
	return &AnswerShippingQueryBadRequest{}
}

/*AnswerShippingQueryBadRequest handles this case with default header values.

Bad Request
*/
type AnswerShippingQueryBadRequest struct {
	Payload *models.Error
}

func (o *AnswerShippingQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryBadRequest  %+v", 400, o.Payload)
}

func (o *AnswerShippingQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerShippingQueryUnauthorized creates a AnswerShippingQueryUnauthorized with default headers values
func NewAnswerShippingQueryUnauthorized() *AnswerShippingQueryUnauthorized {
	return &AnswerShippingQueryUnauthorized{}
}

/*AnswerShippingQueryUnauthorized handles this case with default header values.

Unauthorized
*/
type AnswerShippingQueryUnauthorized struct {
	Payload *models.Error
}

func (o *AnswerShippingQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *AnswerShippingQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerShippingQueryForbidden creates a AnswerShippingQueryForbidden with default headers values
func NewAnswerShippingQueryForbidden() *AnswerShippingQueryForbidden {
	return &AnswerShippingQueryForbidden{}
}

/*AnswerShippingQueryForbidden handles this case with default header values.

Forbidden
*/
type AnswerShippingQueryForbidden struct {
	Payload *models.Error
}

func (o *AnswerShippingQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryForbidden  %+v", 403, o.Payload)
}

func (o *AnswerShippingQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerShippingQueryNotFound creates a AnswerShippingQueryNotFound with default headers values
func NewAnswerShippingQueryNotFound() *AnswerShippingQueryNotFound {
	return &AnswerShippingQueryNotFound{}
}

/*AnswerShippingQueryNotFound handles this case with default header values.

Not Found
*/
type AnswerShippingQueryNotFound struct {
	Payload *models.Error
}

func (o *AnswerShippingQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryNotFound  %+v", 404, o.Payload)
}

func (o *AnswerShippingQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerShippingQueryEnhanceYourCalm creates a AnswerShippingQueryEnhanceYourCalm with default headers values
func NewAnswerShippingQueryEnhanceYourCalm() *AnswerShippingQueryEnhanceYourCalm {
	return &AnswerShippingQueryEnhanceYourCalm{}
}

/*AnswerShippingQueryEnhanceYourCalm handles this case with default header values.

Flood
*/
type AnswerShippingQueryEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *AnswerShippingQueryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *AnswerShippingQueryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerShippingQueryInternalServerError creates a AnswerShippingQueryInternalServerError with default headers values
func NewAnswerShippingQueryInternalServerError() *AnswerShippingQueryInternalServerError {
	return &AnswerShippingQueryInternalServerError{}
}

/*AnswerShippingQueryInternalServerError handles this case with default header values.

Internal
*/
type AnswerShippingQueryInternalServerError struct {
	Payload *models.Error
}

func (o *AnswerShippingQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerShippingQuery][%d] answerShippingQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *AnswerShippingQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
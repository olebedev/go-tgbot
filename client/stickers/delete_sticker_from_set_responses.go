// Code generated by go-swagger; DO NOT EDIT.

package stickers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// DeleteStickerFromSetReader is a Reader for the DeleteStickerFromSet structure.
type DeleteStickerFromSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStickerFromSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteStickerFromSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteStickerFromSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteStickerFromSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteStickerFromSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteStickerFromSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewDeleteStickerFromSetEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteStickerFromSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteStickerFromSetOK creates a DeleteStickerFromSetOK with default headers values
func NewDeleteStickerFromSetOK() *DeleteStickerFromSetOK {
	return &DeleteStickerFromSetOK{}
}

/*DeleteStickerFromSetOK handles this case with default header values.

DeleteStickerFromSetOK delete sticker from set o k
*/
type DeleteStickerFromSetOK struct {
	Payload *models.ResponseBool
}

func (o *DeleteStickerFromSetOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetOK  %+v", 200, o.Payload)
}

func (o *DeleteStickerFromSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStickerFromSetBadRequest creates a DeleteStickerFromSetBadRequest with default headers values
func NewDeleteStickerFromSetBadRequest() *DeleteStickerFromSetBadRequest {
	return &DeleteStickerFromSetBadRequest{}
}

/*DeleteStickerFromSetBadRequest handles this case with default header values.

Bad Request
*/
type DeleteStickerFromSetBadRequest struct {
	Payload *models.Error
}

func (o *DeleteStickerFromSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteStickerFromSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStickerFromSetUnauthorized creates a DeleteStickerFromSetUnauthorized with default headers values
func NewDeleteStickerFromSetUnauthorized() *DeleteStickerFromSetUnauthorized {
	return &DeleteStickerFromSetUnauthorized{}
}

/*DeleteStickerFromSetUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteStickerFromSetUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteStickerFromSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteStickerFromSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStickerFromSetForbidden creates a DeleteStickerFromSetForbidden with default headers values
func NewDeleteStickerFromSetForbidden() *DeleteStickerFromSetForbidden {
	return &DeleteStickerFromSetForbidden{}
}

/*DeleteStickerFromSetForbidden handles this case with default header values.

Forbidden
*/
type DeleteStickerFromSetForbidden struct {
	Payload *models.Error
}

func (o *DeleteStickerFromSetForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetForbidden  %+v", 403, o.Payload)
}

func (o *DeleteStickerFromSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStickerFromSetNotFound creates a DeleteStickerFromSetNotFound with default headers values
func NewDeleteStickerFromSetNotFound() *DeleteStickerFromSetNotFound {
	return &DeleteStickerFromSetNotFound{}
}

/*DeleteStickerFromSetNotFound handles this case with default header values.

Not Found
*/
type DeleteStickerFromSetNotFound struct {
	Payload *models.Error
}

func (o *DeleteStickerFromSetNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStickerFromSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStickerFromSetEnhanceYourCalm creates a DeleteStickerFromSetEnhanceYourCalm with default headers values
func NewDeleteStickerFromSetEnhanceYourCalm() *DeleteStickerFromSetEnhanceYourCalm {
	return &DeleteStickerFromSetEnhanceYourCalm{}
}

/*DeleteStickerFromSetEnhanceYourCalm handles this case with default header values.

Flood
*/
type DeleteStickerFromSetEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *DeleteStickerFromSetEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteStickerFromSetEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStickerFromSetInternalServerError creates a DeleteStickerFromSetInternalServerError with default headers values
func NewDeleteStickerFromSetInternalServerError() *DeleteStickerFromSetInternalServerError {
	return &DeleteStickerFromSetInternalServerError{}
}

/*DeleteStickerFromSetInternalServerError handles this case with default header values.

Internal
*/
type DeleteStickerFromSetInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteStickerFromSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/deleteStickerFromSet][%d] deleteStickerFromSetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteStickerFromSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

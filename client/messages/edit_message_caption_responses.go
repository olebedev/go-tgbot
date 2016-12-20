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

// EditMessageCaptionReader is a Reader for the EditMessageCaption structure.
type EditMessageCaptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditMessageCaptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditMessageCaptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditMessageCaptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditMessageCaptionOK creates a EditMessageCaptionOK with default headers values
func NewEditMessageCaptionOK() *EditMessageCaptionOK {
	return &EditMessageCaptionOK{}
}

/*EditMessageCaptionOK handles this case with default header values.

EditMessageCaptionOK edit message caption o k
*/
type EditMessageCaptionOK struct {
	Payload EditMessageCaptionOKBody
}

func (o *EditMessageCaptionOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionOK  %+v", 200, o.Payload)
}

func (o *EditMessageCaptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageCaptionBadRequest creates a EditMessageCaptionBadRequest with default headers values
func NewEditMessageCaptionBadRequest() *EditMessageCaptionBadRequest {
	return &EditMessageCaptionBadRequest{}
}

/*EditMessageCaptionBadRequest handles this case with default header values.

Error
*/
type EditMessageCaptionBadRequest struct {
	Payload *models.Error
}

func (o *EditMessageCaptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageCaption][%d] editMessageCaptionBadRequest  %+v", 400, o.Payload)
}

func (o *EditMessageCaptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EditMessageCaptionOKBody edit message caption o k body
swagger:model EditMessageCaptionOKBody
*/
type EditMessageCaptionOKBody struct {

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
	Result interface{} `json:"result"`
}

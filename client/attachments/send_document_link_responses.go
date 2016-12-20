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

// SendDocumentLinkReader is a Reader for the SendDocumentLink structure.
type SendDocumentLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendDocumentLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendDocumentLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendDocumentLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendDocumentLinkOK creates a SendDocumentLinkOK with default headers values
func NewSendDocumentLinkOK() *SendDocumentLinkOK {
	return &SendDocumentLinkOK{}
}

/*SendDocumentLinkOK handles this case with default header values.

SendDocumentLinkOK send document link o k
*/
type SendDocumentLinkOK struct {
	Payload *models.ResponseMessage
}

func (o *SendDocumentLinkOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendDocument#link][%d] sendDocumentLinkOK  %+v", 200, o.Payload)
}

func (o *SendDocumentLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendDocumentLinkBadRequest creates a SendDocumentLinkBadRequest with default headers values
func NewSendDocumentLinkBadRequest() *SendDocumentLinkBadRequest {
	return &SendDocumentLinkBadRequest{}
}

/*SendDocumentLinkBadRequest handles this case with default header values.

Error
*/
type SendDocumentLinkBadRequest struct {
	Payload *models.Error
}

func (o *SendDocumentLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendDocument#link][%d] sendDocumentLinkBadRequest  %+v", 400, o.Payload)
}

func (o *SendDocumentLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

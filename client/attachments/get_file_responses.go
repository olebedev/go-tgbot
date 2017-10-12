// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// GetFileReader is a Reader for the GetFile structure.
type GetFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetFileEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFileOK creates a GetFileOK with default headers values
func NewGetFileOK() *GetFileOK {
	return &GetFileOK{}
}

/*GetFileOK handles this case with default header values.

GetFileOK get file o k
*/
type GetFileOK struct {
	Payload GetFileOKBody
}

func (o *GetFileOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileOK  %+v", 200, o.Payload)
}

func (o *GetFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileBadRequest creates a GetFileBadRequest with default headers values
func NewGetFileBadRequest() *GetFileBadRequest {
	return &GetFileBadRequest{}
}

/*GetFileBadRequest handles this case with default header values.

Bad Request
*/
type GetFileBadRequest struct {
	Payload *models.Error
}

func (o *GetFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileUnauthorized creates a GetFileUnauthorized with default headers values
func NewGetFileUnauthorized() *GetFileUnauthorized {
	return &GetFileUnauthorized{}
}

/*GetFileUnauthorized handles this case with default header values.

Unauthorized
*/
type GetFileUnauthorized struct {
	Payload *models.Error
}

func (o *GetFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileForbidden creates a GetFileForbidden with default headers values
func NewGetFileForbidden() *GetFileForbidden {
	return &GetFileForbidden{}
}

/*GetFileForbidden handles this case with default header values.

Forbidden
*/
type GetFileForbidden struct {
	Payload *models.Error
}

func (o *GetFileForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileForbidden  %+v", 403, o.Payload)
}

func (o *GetFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileNotFound creates a GetFileNotFound with default headers values
func NewGetFileNotFound() *GetFileNotFound {
	return &GetFileNotFound{}
}

/*GetFileNotFound handles this case with default header values.

Not Found
*/
type GetFileNotFound struct {
	Payload *models.Error
}

func (o *GetFileNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileNotFound  %+v", 404, o.Payload)
}

func (o *GetFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileEnhanceYourCalm creates a GetFileEnhanceYourCalm with default headers values
func NewGetFileEnhanceYourCalm() *GetFileEnhanceYourCalm {
	return &GetFileEnhanceYourCalm{}
}

/*GetFileEnhanceYourCalm handles this case with default header values.

Flood
*/
type GetFileEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *GetFileEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetFileEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileInternalServerError creates a GetFileInternalServerError with default headers values
func NewGetFileInternalServerError() *GetFileInternalServerError {
	return &GetFileInternalServerError{}
}

/*GetFileInternalServerError handles this case with default header values.

Internal
*/
type GetFileInternalServerError struct {
	Payload *models.Error
}

func (o *GetFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getFile][%d] getFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFileOKBody get file o k body
swagger:model GetFileOKBody
*/

type GetFileOKBody struct {

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
	Result *models.File `json:"result"`
}

/* polymorph GetFileOKBody description false */

/* polymorph GetFileOKBody error_code false */

/* polymorph GetFileOKBody ok false */

/* polymorph GetFileOKBody result false */

// Validate validates this get file o k body
func (o *GetFileOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateErrorCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFileOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getFileOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetFileOKBody) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("getFileOK"+"."+"error_code", "body", o.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (o *GetFileOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("getFileOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *GetFileOKBody) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("getFileOK"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	if o.Result != nil {

		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFileOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFileOKBody) UnmarshalBinary(b []byte) error {
	var res GetFileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

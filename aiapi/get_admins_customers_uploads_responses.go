// Code generated by go-swagger; DO NOT EDIT.

package aiapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Maritime-AI/mai-go-client/models"
)

// GetAdminsCustomersUploadsReader is a Reader for the GetAdminsCustomersUploads structure.
type GetAdminsCustomersUploadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminsCustomersUploadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminsCustomersUploadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAdminsCustomersUploadsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAdminsCustomersUploadsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAdminsCustomersUploadsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admins/customers/uploads] GetAdminsCustomersUploads", response, response.Code())
	}
}

// NewGetAdminsCustomersUploadsOK creates a GetAdminsCustomersUploadsOK with default headers values
func NewGetAdminsCustomersUploadsOK() *GetAdminsCustomersUploadsOK {
	return &GetAdminsCustomersUploadsOK{}
}

/*
GetAdminsCustomersUploadsOK describes a response with status code 200, with default header values.

Success
*/
type GetAdminsCustomersUploadsOK struct {
	Payload []*models.File
}

// IsSuccess returns true when this get admins customers uploads o k response has a 2xx status code
func (o *GetAdminsCustomersUploadsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admins customers uploads o k response has a 3xx status code
func (o *GetAdminsCustomersUploadsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins customers uploads o k response has a 4xx status code
func (o *GetAdminsCustomersUploadsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admins customers uploads o k response has a 5xx status code
func (o *GetAdminsCustomersUploadsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admins customers uploads o k response a status code equal to that given
func (o *GetAdminsCustomersUploadsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get admins customers uploads o k response
func (o *GetAdminsCustomersUploadsOK) Code() int {
	return 200
}

func (o *GetAdminsCustomersUploadsOK) Error() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsOK  %+v", 200, o.Payload)
}

func (o *GetAdminsCustomersUploadsOK) String() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsOK  %+v", 200, o.Payload)
}

func (o *GetAdminsCustomersUploadsOK) GetPayload() []*models.File {
	return o.Payload
}

func (o *GetAdminsCustomersUploadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminsCustomersUploadsBadRequest creates a GetAdminsCustomersUploadsBadRequest with default headers values
func NewGetAdminsCustomersUploadsBadRequest() *GetAdminsCustomersUploadsBadRequest {
	return &GetAdminsCustomersUploadsBadRequest{}
}

/*
GetAdminsCustomersUploadsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAdminsCustomersUploadsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get admins customers uploads bad request response has a 2xx status code
func (o *GetAdminsCustomersUploadsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admins customers uploads bad request response has a 3xx status code
func (o *GetAdminsCustomersUploadsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins customers uploads bad request response has a 4xx status code
func (o *GetAdminsCustomersUploadsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admins customers uploads bad request response has a 5xx status code
func (o *GetAdminsCustomersUploadsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get admins customers uploads bad request response a status code equal to that given
func (o *GetAdminsCustomersUploadsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get admins customers uploads bad request response
func (o *GetAdminsCustomersUploadsBadRequest) Code() int {
	return 400
}

func (o *GetAdminsCustomersUploadsBadRequest) Error() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAdminsCustomersUploadsBadRequest) String() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAdminsCustomersUploadsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAdminsCustomersUploadsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminsCustomersUploadsUnauthorized creates a GetAdminsCustomersUploadsUnauthorized with default headers values
func NewGetAdminsCustomersUploadsUnauthorized() *GetAdminsCustomersUploadsUnauthorized {
	return &GetAdminsCustomersUploadsUnauthorized{}
}

/*
GetAdminsCustomersUploadsUnauthorized describes a response with status code 401, with default header values.

Not Authorized
*/
type GetAdminsCustomersUploadsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get admins customers uploads unauthorized response has a 2xx status code
func (o *GetAdminsCustomersUploadsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admins customers uploads unauthorized response has a 3xx status code
func (o *GetAdminsCustomersUploadsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins customers uploads unauthorized response has a 4xx status code
func (o *GetAdminsCustomersUploadsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admins customers uploads unauthorized response has a 5xx status code
func (o *GetAdminsCustomersUploadsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get admins customers uploads unauthorized response a status code equal to that given
func (o *GetAdminsCustomersUploadsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get admins customers uploads unauthorized response
func (o *GetAdminsCustomersUploadsUnauthorized) Code() int {
	return 401
}

func (o *GetAdminsCustomersUploadsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAdminsCustomersUploadsUnauthorized) String() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAdminsCustomersUploadsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAdminsCustomersUploadsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminsCustomersUploadsInternalServerError creates a GetAdminsCustomersUploadsInternalServerError with default headers values
func NewGetAdminsCustomersUploadsInternalServerError() *GetAdminsCustomersUploadsInternalServerError {
	return &GetAdminsCustomersUploadsInternalServerError{}
}

/*
GetAdminsCustomersUploadsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAdminsCustomersUploadsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get admins customers uploads internal server error response has a 2xx status code
func (o *GetAdminsCustomersUploadsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admins customers uploads internal server error response has a 3xx status code
func (o *GetAdminsCustomersUploadsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admins customers uploads internal server error response has a 4xx status code
func (o *GetAdminsCustomersUploadsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admins customers uploads internal server error response has a 5xx status code
func (o *GetAdminsCustomersUploadsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get admins customers uploads internal server error response a status code equal to that given
func (o *GetAdminsCustomersUploadsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get admins customers uploads internal server error response
func (o *GetAdminsCustomersUploadsInternalServerError) Code() int {
	return 500
}

func (o *GetAdminsCustomersUploadsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAdminsCustomersUploadsInternalServerError) String() string {
	return fmt.Sprintf("[GET /admins/customers/uploads][%d] getAdminsCustomersUploadsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAdminsCustomersUploadsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAdminsCustomersUploadsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

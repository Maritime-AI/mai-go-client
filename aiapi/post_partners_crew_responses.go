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

// PostPartnersCrewReader is a Reader for the PostPartnersCrew structure.
type PostPartnersCrewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPartnersCrewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostPartnersCrewNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPartnersCrewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPartnersCrewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPartnersCrewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPartnersCrewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /partners/crew] PostPartnersCrew", response, response.Code())
	}
}

// NewPostPartnersCrewNoContent creates a PostPartnersCrewNoContent with default headers values
func NewPostPartnersCrewNoContent() *PostPartnersCrewNoContent {
	return &PostPartnersCrewNoContent{}
}

/*
PostPartnersCrewNoContent describes a response with status code 204, with default header values.

Success No Content
*/
type PostPartnersCrewNoContent struct {
}

// IsSuccess returns true when this post partners crew no content response has a 2xx status code
func (o *PostPartnersCrewNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post partners crew no content response has a 3xx status code
func (o *PostPartnersCrewNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post partners crew no content response has a 4xx status code
func (o *PostPartnersCrewNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post partners crew no content response has a 5xx status code
func (o *PostPartnersCrewNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post partners crew no content response a status code equal to that given
func (o *PostPartnersCrewNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post partners crew no content response
func (o *PostPartnersCrewNoContent) Code() int {
	return 204
}

func (o *PostPartnersCrewNoContent) Error() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewNoContent ", 204)
}

func (o *PostPartnersCrewNoContent) String() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewNoContent ", 204)
}

func (o *PostPartnersCrewNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartnersCrewBadRequest creates a PostPartnersCrewBadRequest with default headers values
func NewPostPartnersCrewBadRequest() *PostPartnersCrewBadRequest {
	return &PostPartnersCrewBadRequest{}
}

/*
PostPartnersCrewBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostPartnersCrewBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post partners crew bad request response has a 2xx status code
func (o *PostPartnersCrewBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post partners crew bad request response has a 3xx status code
func (o *PostPartnersCrewBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post partners crew bad request response has a 4xx status code
func (o *PostPartnersCrewBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post partners crew bad request response has a 5xx status code
func (o *PostPartnersCrewBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post partners crew bad request response a status code equal to that given
func (o *PostPartnersCrewBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post partners crew bad request response
func (o *PostPartnersCrewBadRequest) Code() int {
	return 400
}

func (o *PostPartnersCrewBadRequest) Error() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewBadRequest  %+v", 400, o.Payload)
}

func (o *PostPartnersCrewBadRequest) String() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewBadRequest  %+v", 400, o.Payload)
}

func (o *PostPartnersCrewBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostPartnersCrewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartnersCrewUnauthorized creates a PostPartnersCrewUnauthorized with default headers values
func NewPostPartnersCrewUnauthorized() *PostPartnersCrewUnauthorized {
	return &PostPartnersCrewUnauthorized{}
}

/*
PostPartnersCrewUnauthorized describes a response with status code 401, with default header values.

Not Authorized
*/
type PostPartnersCrewUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this post partners crew unauthorized response has a 2xx status code
func (o *PostPartnersCrewUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post partners crew unauthorized response has a 3xx status code
func (o *PostPartnersCrewUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post partners crew unauthorized response has a 4xx status code
func (o *PostPartnersCrewUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post partners crew unauthorized response has a 5xx status code
func (o *PostPartnersCrewUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post partners crew unauthorized response a status code equal to that given
func (o *PostPartnersCrewUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post partners crew unauthorized response
func (o *PostPartnersCrewUnauthorized) Code() int {
	return 401
}

func (o *PostPartnersCrewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPartnersCrewUnauthorized) String() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPartnersCrewUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostPartnersCrewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartnersCrewForbidden creates a PostPartnersCrewForbidden with default headers values
func NewPostPartnersCrewForbidden() *PostPartnersCrewForbidden {
	return &PostPartnersCrewForbidden{}
}

/*
PostPartnersCrewForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostPartnersCrewForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this post partners crew forbidden response has a 2xx status code
func (o *PostPartnersCrewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post partners crew forbidden response has a 3xx status code
func (o *PostPartnersCrewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post partners crew forbidden response has a 4xx status code
func (o *PostPartnersCrewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post partners crew forbidden response has a 5xx status code
func (o *PostPartnersCrewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post partners crew forbidden response a status code equal to that given
func (o *PostPartnersCrewForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post partners crew forbidden response
func (o *PostPartnersCrewForbidden) Code() int {
	return 403
}

func (o *PostPartnersCrewForbidden) Error() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewForbidden  %+v", 403, o.Payload)
}

func (o *PostPartnersCrewForbidden) String() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewForbidden  %+v", 403, o.Payload)
}

func (o *PostPartnersCrewForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostPartnersCrewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartnersCrewInternalServerError creates a PostPartnersCrewInternalServerError with default headers values
func NewPostPartnersCrewInternalServerError() *PostPartnersCrewInternalServerError {
	return &PostPartnersCrewInternalServerError{}
}

/*
PostPartnersCrewInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostPartnersCrewInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post partners crew internal server error response has a 2xx status code
func (o *PostPartnersCrewInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post partners crew internal server error response has a 3xx status code
func (o *PostPartnersCrewInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post partners crew internal server error response has a 4xx status code
func (o *PostPartnersCrewInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post partners crew internal server error response has a 5xx status code
func (o *PostPartnersCrewInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post partners crew internal server error response a status code equal to that given
func (o *PostPartnersCrewInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post partners crew internal server error response
func (o *PostPartnersCrewInternalServerError) Code() int {
	return 500
}

func (o *PostPartnersCrewInternalServerError) Error() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPartnersCrewInternalServerError) String() string {
	return fmt.Sprintf("[POST /partners/crew][%d] postPartnersCrewInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPartnersCrewInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostPartnersCrewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/*ListProjectsOK handles this case with default header values.

Return all matched projects.
*/
type ListProjectsOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of projects
	 */
	XTotalCount int64

	Payload []*model.Project
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /projects][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) GetPayload() []*model.Project {
	return o.Payload
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsUnauthorized creates a ListProjectsUnauthorized with default headers values
func NewListProjectsUnauthorized() *ListProjectsUnauthorized {
	return &ListProjectsUnauthorized{}
}

/*ListProjectsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListProjectsUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects][%d] listProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProjectsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsInternalServerError creates a ListProjectsInternalServerError with default headers values
func NewListProjectsInternalServerError() *ListProjectsInternalServerError {
	return &ListProjectsInternalServerError{}
}

/*ListProjectsInternalServerError handles this case with default header values.

Internal server error
*/
type ListProjectsInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects][%d] listProjectsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// SearchLdapUserReader is a Reader for the SearchLdapUser structure.
type SearchLdapUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLdapUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLdapUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchLdapUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchLdapUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchLdapUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchLdapUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchLdapUserOK creates a SearchLdapUserOK with default headers values
func NewSearchLdapUserOK() *SearchLdapUserOK {
	return &SearchLdapUserOK{}
}

/*SearchLdapUserOK handles this case with default header values.

Search ldap users successfully.
*/
type SearchLdapUserOK struct {
	Payload []*model.LdapUser
}

func (o *SearchLdapUserOK) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserOK  %+v", 200, o.Payload)
}

func (o *SearchLdapUserOK) GetPayload() []*model.LdapUser {
	return o.Payload
}

func (o *SearchLdapUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapUserBadRequest creates a SearchLdapUserBadRequest with default headers values
func NewSearchLdapUserBadRequest() *SearchLdapUserBadRequest {
	return &SearchLdapUserBadRequest{}
}

/*SearchLdapUserBadRequest handles this case with default header values.

Bad request
*/
type SearchLdapUserBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserBadRequest  %+v", 400, o.Payload)
}

func (o *SearchLdapUserBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapUserUnauthorized creates a SearchLdapUserUnauthorized with default headers values
func NewSearchLdapUserUnauthorized() *SearchLdapUserUnauthorized {
	return &SearchLdapUserUnauthorized{}
}

/*SearchLdapUserUnauthorized handles this case with default header values.

Unauthorized
*/
type SearchLdapUserUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchLdapUserUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapUserForbidden creates a SearchLdapUserForbidden with default headers values
func NewSearchLdapUserForbidden() *SearchLdapUserForbidden {
	return &SearchLdapUserForbidden{}
}

/*SearchLdapUserForbidden handles this case with default header values.

Forbidden
*/
type SearchLdapUserForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapUserForbidden) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserForbidden  %+v", 403, o.Payload)
}

func (o *SearchLdapUserForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapUserInternalServerError creates a SearchLdapUserInternalServerError with default headers values
func NewSearchLdapUserInternalServerError() *SearchLdapUserInternalServerError {
	return &SearchLdapUserInternalServerError{}
}

/*SearchLdapUserInternalServerError handles this case with default header values.

Internal server error
*/
type SearchLdapUserInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchLdapUserInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

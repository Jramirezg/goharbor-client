// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetHealthParams creates a new GetHealthParams object
// with the default values initialized.
func NewGetHealthParams() *GetHealthParams {
	var ()
	return &GetHealthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthParamsWithTimeout creates a new GetHealthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHealthParamsWithTimeout(timeout time.Duration) *GetHealthParams {
	var ()
	return &GetHealthParams{

		timeout: timeout,
	}
}

// NewGetHealthParamsWithContext creates a new GetHealthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHealthParamsWithContext(ctx context.Context) *GetHealthParams {
	var ()
	return &GetHealthParams{

		Context: ctx,
	}
}

// NewGetHealthParamsWithHTTPClient creates a new GetHealthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHealthParamsWithHTTPClient(client *http.Client) *GetHealthParams {
	var ()
	return &GetHealthParams{
		HTTPClient: client,
	}
}

/*GetHealthParams contains all the parameters to send to the API endpoint
for the get health operation typically these are written to a http.Request
*/
type GetHealthParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get health params
func (o *GetHealthParams) WithTimeout(timeout time.Duration) *GetHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get health params
func (o *GetHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get health params
func (o *GetHealthParams) WithContext(ctx context.Context) *GetHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get health params
func (o *GetHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get health params
func (o *GetHealthParams) WithHTTPClient(client *http.Client) *GetHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get health params
func (o *GetHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get health params
func (o *GetHealthParams) WithXRequestID(xRequestID *string) *GetHealthParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get health params
func (o *GetHealthParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
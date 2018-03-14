// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetChangeParentRevisionsParams creates a new GetChangeParentRevisionsParams object
// with the default values initialized.
func NewGetChangeParentRevisionsParams() *GetChangeParentRevisionsParams {
	var ()
	return &GetChangeParentRevisionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChangeParentRevisionsParamsWithTimeout creates a new GetChangeParentRevisionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChangeParentRevisionsParamsWithTimeout(timeout time.Duration) *GetChangeParentRevisionsParams {
	var ()
	return &GetChangeParentRevisionsParams{

		timeout: timeout,
	}
}

// NewGetChangeParentRevisionsParamsWithContext creates a new GetChangeParentRevisionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChangeParentRevisionsParamsWithContext(ctx context.Context) *GetChangeParentRevisionsParams {
	var ()
	return &GetChangeParentRevisionsParams{

		Context: ctx,
	}
}

// NewGetChangeParentRevisionsParamsWithHTTPClient creates a new GetChangeParentRevisionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChangeParentRevisionsParamsWithHTTPClient(client *http.Client) *GetChangeParentRevisionsParams {
	var ()
	return &GetChangeParentRevisionsParams{
		HTTPClient: client,
	}
}

/*GetChangeParentRevisionsParams contains all the parameters to send to the API endpoint
for the get change parent revisions operation typically these are written to a http.Request
*/
type GetChangeParentRevisionsParams struct {

	/*ChangeLocator*/
	ChangeLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) WithTimeout(timeout time.Duration) *GetChangeParentRevisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) WithContext(ctx context.Context) *GetChangeParentRevisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) WithHTTPClient(client *http.Client) *GetChangeParentRevisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeLocator adds the changeLocator to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) WithChangeLocator(changeLocator string) *GetChangeParentRevisionsParams {
	o.SetChangeLocator(changeLocator)
	return o
}

// SetChangeLocator adds the changeLocator to the get change parent revisions params
func (o *GetChangeParentRevisionsParams) SetChangeLocator(changeLocator string) {
	o.ChangeLocator = changeLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetChangeParentRevisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param changeLocator
	if err := r.SetPathParam("changeLocator", o.ChangeLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
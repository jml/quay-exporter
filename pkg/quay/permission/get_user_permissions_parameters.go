// Code generated by go-swagger; DO NOT EDIT.

package permission

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

// NewGetUserPermissionsParams creates a new GetUserPermissionsParams object
// with the default values initialized.
func NewGetUserPermissionsParams() *GetUserPermissionsParams {
	var ()
	return &GetUserPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserPermissionsParamsWithTimeout creates a new GetUserPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserPermissionsParamsWithTimeout(timeout time.Duration) *GetUserPermissionsParams {
	var ()
	return &GetUserPermissionsParams{

		timeout: timeout,
	}
}

// NewGetUserPermissionsParamsWithContext creates a new GetUserPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserPermissionsParamsWithContext(ctx context.Context) *GetUserPermissionsParams {
	var ()
	return &GetUserPermissionsParams{

		Context: ctx,
	}
}

// NewGetUserPermissionsParamsWithHTTPClient creates a new GetUserPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserPermissionsParamsWithHTTPClient(client *http.Client) *GetUserPermissionsParams {
	var ()
	return &GetUserPermissionsParams{
		HTTPClient: client,
	}
}

/*GetUserPermissionsParams contains all the parameters to send to the API endpoint
for the get user permissions operation typically these are written to a http.Request
*/
type GetUserPermissionsParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Username
	  The username of the user to which the permission applies

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user permissions params
func (o *GetUserPermissionsParams) WithTimeout(timeout time.Duration) *GetUserPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user permissions params
func (o *GetUserPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user permissions params
func (o *GetUserPermissionsParams) WithContext(ctx context.Context) *GetUserPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user permissions params
func (o *GetUserPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user permissions params
func (o *GetUserPermissionsParams) WithHTTPClient(client *http.Client) *GetUserPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user permissions params
func (o *GetUserPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the get user permissions params
func (o *GetUserPermissionsParams) WithRepository(repository string) *GetUserPermissionsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get user permissions params
func (o *GetUserPermissionsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithUsername adds the username to the get user permissions params
func (o *GetUserPermissionsParams) WithUsername(username string) *GetUserPermissionsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get user permissions params
func (o *GetUserPermissionsParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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
	"github.com/go-openapi/swag"
)

// NewDeleteUserGroupParams creates a new DeleteUserGroupParams object
// with the default values initialized.
func NewDeleteUserGroupParams() *DeleteUserGroupParams {
	var ()
	return &DeleteUserGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserGroupParamsWithTimeout creates a new DeleteUserGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserGroupParamsWithTimeout(timeout time.Duration) *DeleteUserGroupParams {
	var ()
	return &DeleteUserGroupParams{

		timeout: timeout,
	}
}

// NewDeleteUserGroupParamsWithContext creates a new DeleteUserGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserGroupParamsWithContext(ctx context.Context) *DeleteUserGroupParams {
	var ()
	return &DeleteUserGroupParams{

		Context: ctx,
	}
}

// NewDeleteUserGroupParamsWithHTTPClient creates a new DeleteUserGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserGroupParamsWithHTTPClient(client *http.Client) *DeleteUserGroupParams {
	var ()
	return &DeleteUserGroupParams{
		HTTPClient: client,
	}
}

/*DeleteUserGroupParams contains all the parameters to send to the API endpoint
for the delete user group operation typically these are written to a http.Request
*/
type DeleteUserGroupParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*GroupID*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user group params
func (o *DeleteUserGroupParams) WithTimeout(timeout time.Duration) *DeleteUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user group params
func (o *DeleteUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user group params
func (o *DeleteUserGroupParams) WithContext(ctx context.Context) *DeleteUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user group params
func (o *DeleteUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user group params
func (o *DeleteUserGroupParams) WithHTTPClient(client *http.Client) *DeleteUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user group params
func (o *DeleteUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete user group params
func (o *DeleteUserGroupParams) WithXRequestID(xRequestID *string) *DeleteUserGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete user group params
func (o *DeleteUserGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGroupID adds the groupID to the delete user group params
func (o *DeleteUserGroupParams) WithGroupID(groupID int64) *DeleteUserGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete user group params
func (o *DeleteUserGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
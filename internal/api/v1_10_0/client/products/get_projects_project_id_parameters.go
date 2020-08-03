// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewGetProjectsProjectIDParams creates a new GetProjectsProjectIDParams object
// with the default values initialized.
func NewGetProjectsProjectIDParams() *GetProjectsProjectIDParams {
	var ()
	return &GetProjectsProjectIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsProjectIDParamsWithTimeout creates a new GetProjectsProjectIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsProjectIDParamsWithTimeout(timeout time.Duration) *GetProjectsProjectIDParams {
	var ()
	return &GetProjectsProjectIDParams{

		timeout: timeout,
	}
}

// NewGetProjectsProjectIDParamsWithContext creates a new GetProjectsProjectIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsProjectIDParamsWithContext(ctx context.Context) *GetProjectsProjectIDParams {
	var ()
	return &GetProjectsProjectIDParams{

		Context: ctx,
	}
}

// NewGetProjectsProjectIDParamsWithHTTPClient creates a new GetProjectsProjectIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectsProjectIDParamsWithHTTPClient(client *http.Client) *GetProjectsProjectIDParams {
	var ()
	return &GetProjectsProjectIDParams{
		HTTPClient: client,
	}
}

/*GetProjectsProjectIDParams contains all the parameters to send to the API endpoint
for the get projects project ID operation typically these are written to a http.Request
*/
type GetProjectsProjectIDParams struct {

	/*ProjectID
	  Project ID for filtering results.

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get projects project ID params
func (o *GetProjectsProjectIDParams) WithTimeout(timeout time.Duration) *GetProjectsProjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects project ID params
func (o *GetProjectsProjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects project ID params
func (o *GetProjectsProjectIDParams) WithContext(ctx context.Context) *GetProjectsProjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects project ID params
func (o *GetProjectsProjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects project ID params
func (o *GetProjectsProjectIDParams) WithHTTPClient(client *http.Client) *GetProjectsProjectIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects project ID params
func (o *GetProjectsProjectIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get projects project ID params
func (o *GetProjectsProjectIDParams) WithProjectID(projectID int64) *GetProjectsProjectIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get projects project ID params
func (o *GetProjectsProjectIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsProjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

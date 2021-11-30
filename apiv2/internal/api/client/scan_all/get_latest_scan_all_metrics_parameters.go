// Code generated by go-swagger; DO NOT EDIT.

package scan_all

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

// NewGetLatestScanAllMetricsParams creates a new GetLatestScanAllMetricsParams object
// with the default values initialized.
func NewGetLatestScanAllMetricsParams() *GetLatestScanAllMetricsParams {
	var ()
	return &GetLatestScanAllMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestScanAllMetricsParamsWithTimeout creates a new GetLatestScanAllMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLatestScanAllMetricsParamsWithTimeout(timeout time.Duration) *GetLatestScanAllMetricsParams {
	var ()
	return &GetLatestScanAllMetricsParams{

		timeout: timeout,
	}
}

// NewGetLatestScanAllMetricsParamsWithContext creates a new GetLatestScanAllMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLatestScanAllMetricsParamsWithContext(ctx context.Context) *GetLatestScanAllMetricsParams {
	var ()
	return &GetLatestScanAllMetricsParams{

		Context: ctx,
	}
}

// NewGetLatestScanAllMetricsParamsWithHTTPClient creates a new GetLatestScanAllMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLatestScanAllMetricsParamsWithHTTPClient(client *http.Client) *GetLatestScanAllMetricsParams {
	var ()
	return &GetLatestScanAllMetricsParams{
		HTTPClient: client,
	}
}

/*GetLatestScanAllMetricsParams contains all the parameters to send to the API endpoint
for the get latest scan all metrics operation typically these are written to a http.Request
*/
type GetLatestScanAllMetricsParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) WithTimeout(timeout time.Duration) *GetLatestScanAllMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) WithContext(ctx context.Context) *GetLatestScanAllMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) WithHTTPClient(client *http.Client) *GetLatestScanAllMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) WithXRequestID(xRequestID *string) *GetLatestScanAllMetricsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get latest scan all metrics params
func (o *GetLatestScanAllMetricsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestScanAllMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
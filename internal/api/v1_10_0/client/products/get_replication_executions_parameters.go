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

// NewGetReplicationExecutionsParams creates a new GetReplicationExecutionsParams object
// with the default values initialized.
func NewGetReplicationExecutionsParams() *GetReplicationExecutionsParams {
	var ()
	return &GetReplicationExecutionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationExecutionsParamsWithTimeout creates a new GetReplicationExecutionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReplicationExecutionsParamsWithTimeout(timeout time.Duration) *GetReplicationExecutionsParams {
	var ()
	return &GetReplicationExecutionsParams{

		timeout: timeout,
	}
}

// NewGetReplicationExecutionsParamsWithContext creates a new GetReplicationExecutionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReplicationExecutionsParamsWithContext(ctx context.Context) *GetReplicationExecutionsParams {
	var ()
	return &GetReplicationExecutionsParams{

		Context: ctx,
	}
}

// NewGetReplicationExecutionsParamsWithHTTPClient creates a new GetReplicationExecutionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReplicationExecutionsParamsWithHTTPClient(client *http.Client) *GetReplicationExecutionsParams {
	var ()
	return &GetReplicationExecutionsParams{
		HTTPClient: client,
	}
}

/*GetReplicationExecutionsParams contains all the parameters to send to the API endpoint
for the get replication executions operation typically these are written to a http.Request
*/
type GetReplicationExecutionsParams struct {

	/*Page
	  The page.

	*/
	Page *int64
	/*PageSize
	  The page size.

	*/
	PageSize *int64
	/*PolicyID
	  The policy ID.

	*/
	PolicyID *int64
	/*Status
	  The execution status.

	*/
	Status *string
	/*Trigger
	  The trigger mode.

	*/
	Trigger *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get replication executions params
func (o *GetReplicationExecutionsParams) WithTimeout(timeout time.Duration) *GetReplicationExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication executions params
func (o *GetReplicationExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication executions params
func (o *GetReplicationExecutionsParams) WithContext(ctx context.Context) *GetReplicationExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication executions params
func (o *GetReplicationExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication executions params
func (o *GetReplicationExecutionsParams) WithHTTPClient(client *http.Client) *GetReplicationExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication executions params
func (o *GetReplicationExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get replication executions params
func (o *GetReplicationExecutionsParams) WithPage(page *int64) *GetReplicationExecutionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get replication executions params
func (o *GetReplicationExecutionsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get replication executions params
func (o *GetReplicationExecutionsParams) WithPageSize(pageSize *int64) *GetReplicationExecutionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get replication executions params
func (o *GetReplicationExecutionsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPolicyID adds the policyID to the get replication executions params
func (o *GetReplicationExecutionsParams) WithPolicyID(policyID *int64) *GetReplicationExecutionsParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the get replication executions params
func (o *GetReplicationExecutionsParams) SetPolicyID(policyID *int64) {
	o.PolicyID = policyID
}

// WithStatus adds the status to the get replication executions params
func (o *GetReplicationExecutionsParams) WithStatus(status *string) *GetReplicationExecutionsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get replication executions params
func (o *GetReplicationExecutionsParams) SetStatus(status *string) {
	o.Status = status
}

// WithTrigger adds the trigger to the get replication executions params
func (o *GetReplicationExecutionsParams) WithTrigger(trigger *string) *GetReplicationExecutionsParams {
	o.SetTrigger(trigger)
	return o
}

// SetTrigger adds the trigger to the get replication executions params
func (o *GetReplicationExecutionsParams) SetTrigger(trigger *string) {
	o.Trigger = trigger
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.PolicyID != nil {

		// query param policy_id
		var qrPolicyID int64
		if o.PolicyID != nil {
			qrPolicyID = *o.PolicyID
		}
		qPolicyID := swag.FormatInt64(qrPolicyID)
		if qPolicyID != "" {
			if err := r.SetQueryParam("policy_id", qPolicyID); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Trigger != nil {

		// query param trigger
		var qrTrigger string
		if o.Trigger != nil {
			qrTrigger = *o.Trigger
		}
		qTrigger := qrTrigger
		if qTrigger != "" {
			if err := r.SetQueryParam("trigger", qTrigger); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

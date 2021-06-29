// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
)

// GetProjectsProjectIDImmutabletagrulesReader is a Reader for the GetProjectsProjectIDImmutabletagrules structure.
type GetProjectsProjectIDImmutabletagrulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDImmutabletagrulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDImmutabletagrulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDImmutabletagrulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDImmutabletagrulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDImmutabletagrulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDImmutabletagrulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDImmutabletagrulesOK creates a GetProjectsProjectIDImmutabletagrulesOK with default headers values
func NewGetProjectsProjectIDImmutabletagrulesOK() *GetProjectsProjectIDImmutabletagrulesOK {
	return &GetProjectsProjectIDImmutabletagrulesOK{}
}

/* GetProjectsProjectIDImmutabletagrulesOK describes a response with status code 200, with default header values.

List project immutable tag rules successfully.
*/
type GetProjectsProjectIDImmutabletagrulesOK struct {
	Payload []*legacy.ImmutableRule
}

func (o *GetProjectsProjectIDImmutabletagrulesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/immutabletagrules][%d] getProjectsProjectIdImmutabletagrulesOK  %+v", 200, o.Payload)
}
func (o *GetProjectsProjectIDImmutabletagrulesOK) GetPayload() []*legacy.ImmutableRule {
	return o.Payload
}

func (o *GetProjectsProjectIDImmutabletagrulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDImmutabletagrulesBadRequest creates a GetProjectsProjectIDImmutabletagrulesBadRequest with default headers values
func NewGetProjectsProjectIDImmutabletagrulesBadRequest() *GetProjectsProjectIDImmutabletagrulesBadRequest {
	return &GetProjectsProjectIDImmutabletagrulesBadRequest{}
}

/* GetProjectsProjectIDImmutabletagrulesBadRequest describes a response with status code 400, with default header values.

Illegal format of provided ID value.
*/
type GetProjectsProjectIDImmutabletagrulesBadRequest struct {
}

func (o *GetProjectsProjectIDImmutabletagrulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/immutabletagrules][%d] getProjectsProjectIdImmutabletagrulesBadRequest ", 400)
}

func (o *GetProjectsProjectIDImmutabletagrulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDImmutabletagrulesUnauthorized creates a GetProjectsProjectIDImmutabletagrulesUnauthorized with default headers values
func NewGetProjectsProjectIDImmutabletagrulesUnauthorized() *GetProjectsProjectIDImmutabletagrulesUnauthorized {
	return &GetProjectsProjectIDImmutabletagrulesUnauthorized{}
}

/* GetProjectsProjectIDImmutabletagrulesUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetProjectsProjectIDImmutabletagrulesUnauthorized struct {
}

func (o *GetProjectsProjectIDImmutabletagrulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/immutabletagrules][%d] getProjectsProjectIdImmutabletagrulesUnauthorized ", 401)
}

func (o *GetProjectsProjectIDImmutabletagrulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDImmutabletagrulesForbidden creates a GetProjectsProjectIDImmutabletagrulesForbidden with default headers values
func NewGetProjectsProjectIDImmutabletagrulesForbidden() *GetProjectsProjectIDImmutabletagrulesForbidden {
	return &GetProjectsProjectIDImmutabletagrulesForbidden{}
}

/* GetProjectsProjectIDImmutabletagrulesForbidden describes a response with status code 403, with default header values.

User have no permission to list immutable tag rules of the project.
*/
type GetProjectsProjectIDImmutabletagrulesForbidden struct {
}

func (o *GetProjectsProjectIDImmutabletagrulesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/immutabletagrules][%d] getProjectsProjectIdImmutabletagrulesForbidden ", 403)
}

func (o *GetProjectsProjectIDImmutabletagrulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDImmutabletagrulesInternalServerError creates a GetProjectsProjectIDImmutabletagrulesInternalServerError with default headers values
func NewGetProjectsProjectIDImmutabletagrulesInternalServerError() *GetProjectsProjectIDImmutabletagrulesInternalServerError {
	return &GetProjectsProjectIDImmutabletagrulesInternalServerError{}
}

/* GetProjectsProjectIDImmutabletagrulesInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDImmutabletagrulesInternalServerError struct {
}

func (o *GetProjectsProjectIDImmutabletagrulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/immutabletagrules][%d] getProjectsProjectIdImmutabletagrulesInternalServerError ", 500)
}

func (o *GetProjectsProjectIDImmutabletagrulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

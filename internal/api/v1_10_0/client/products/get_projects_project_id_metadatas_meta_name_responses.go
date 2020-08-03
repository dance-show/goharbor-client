// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/model/v1_10_0"
)

// GetProjectsProjectIDMetadatasMetaNameReader is a Reader for the GetProjectsProjectIDMetadatasMetaName structure.
type GetProjectsProjectIDMetadatasMetaNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDMetadatasMetaNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDMetadatasMetaNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectsProjectIDMetadatasMetaNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDMetadatasMetaNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsProjectIDMetadatasMetaNameOK creates a GetProjectsProjectIDMetadatasMetaNameOK with default headers values
func NewGetProjectsProjectIDMetadatasMetaNameOK() *GetProjectsProjectIDMetadatasMetaNameOK {
	return &GetProjectsProjectIDMetadatasMetaNameOK{}
}

/*GetProjectsProjectIDMetadatasMetaNameOK handles this case with default header values.

Get metadata successfully.
*/
type GetProjectsProjectIDMetadatasMetaNameOK struct {
	Payload *v1_10_0.ProjectMetadata
}

func (o *GetProjectsProjectIDMetadatasMetaNameOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/metadatas/{meta_name}][%d] getProjectsProjectIdMetadatasMetaNameOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDMetadatasMetaNameOK) GetPayload() *v1_10_0.ProjectMetadata {
	return o.Payload
}

func (o *GetProjectsProjectIDMetadatasMetaNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(v1_10_0.ProjectMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDMetadatasMetaNameUnauthorized creates a GetProjectsProjectIDMetadatasMetaNameUnauthorized with default headers values
func NewGetProjectsProjectIDMetadatasMetaNameUnauthorized() *GetProjectsProjectIDMetadatasMetaNameUnauthorized {
	return &GetProjectsProjectIDMetadatasMetaNameUnauthorized{}
}

/*GetProjectsProjectIDMetadatasMetaNameUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDMetadatasMetaNameUnauthorized struct {
}

func (o *GetProjectsProjectIDMetadatasMetaNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/metadatas/{meta_name}][%d] getProjectsProjectIdMetadatasMetaNameUnauthorized ", 401)
}

func (o *GetProjectsProjectIDMetadatasMetaNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewGetProjectsProjectIDMetadatasMetaNameInternalServerError creates a GetProjectsProjectIDMetadatasMetaNameInternalServerError with default headers values
func NewGetProjectsProjectIDMetadatasMetaNameInternalServerError() *GetProjectsProjectIDMetadatasMetaNameInternalServerError {
	return &GetProjectsProjectIDMetadatasMetaNameInternalServerError{}
}

/*GetProjectsProjectIDMetadatasMetaNameInternalServerError handles this case with default header values.

Internal server errors.
*/
type GetProjectsProjectIDMetadatasMetaNameInternalServerError struct {
}

func (o *GetProjectsProjectIDMetadatasMetaNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/metadatas/{meta_name}][%d] getProjectsProjectIdMetadatasMetaNameInternalServerError ", 500)
}

func (o *GetProjectsProjectIDMetadatasMetaNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

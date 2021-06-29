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

// GetRegistriesIDInfoReader is a Reader for the GetRegistriesIDInfo structure.
type GetRegistriesIDInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistriesIDInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistriesIDInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistriesIDInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistriesIDInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistriesIDInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistriesIDInfoOK creates a GetRegistriesIDInfoOK with default headers values
func NewGetRegistriesIDInfoOK() *GetRegistriesIDInfoOK {
	return &GetRegistriesIDInfoOK{}
}

/* GetRegistriesIDInfoOK describes a response with status code 200, with default header values.

Get registry successfully.
*/
type GetRegistriesIDInfoOK struct {
	Payload *legacy.RegistryInfo
}

func (o *GetRegistriesIDInfoOK) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistriesIdInfoOK  %+v", 200, o.Payload)
}
func (o *GetRegistriesIDInfoOK) GetPayload() *legacy.RegistryInfo {
	return o.Payload
}

func (o *GetRegistriesIDInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.RegistryInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistriesIDInfoUnauthorized creates a GetRegistriesIDInfoUnauthorized with default headers values
func NewGetRegistriesIDInfoUnauthorized() *GetRegistriesIDInfoUnauthorized {
	return &GetRegistriesIDInfoUnauthorized{}
}

/* GetRegistriesIDInfoUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetRegistriesIDInfoUnauthorized struct {
}

func (o *GetRegistriesIDInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistriesIdInfoUnauthorized ", 401)
}

func (o *GetRegistriesIDInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistriesIDInfoNotFound creates a GetRegistriesIDInfoNotFound with default headers values
func NewGetRegistriesIDInfoNotFound() *GetRegistriesIDInfoNotFound {
	return &GetRegistriesIDInfoNotFound{}
}

/* GetRegistriesIDInfoNotFound describes a response with status code 404, with default header values.

Registry not found
*/
type GetRegistriesIDInfoNotFound struct {
}

func (o *GetRegistriesIDInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistriesIdInfoNotFound ", 404)
}

func (o *GetRegistriesIDInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistriesIDInfoInternalServerError creates a GetRegistriesIDInfoInternalServerError with default headers values
func NewGetRegistriesIDInfoInternalServerError() *GetRegistriesIDInfoInternalServerError {
	return &GetRegistriesIDInfoInternalServerError{}
}

/* GetRegistriesIDInfoInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetRegistriesIDInfoInternalServerError struct {
}

func (o *GetRegistriesIDInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistriesIdInfoInternalServerError ", 500)
}

func (o *GetRegistriesIDInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2DownloadHostIgnitionReader is a Reader for the V2DownloadHostIgnition structure.
type V2DownloadHostIgnitionReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V2DownloadHostIgnitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2DownloadHostIgnitionOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2DownloadHostIgnitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2DownloadHostIgnitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2DownloadHostIgnitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2DownloadHostIgnitionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2DownloadHostIgnitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2DownloadHostIgnitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2DownloadHostIgnitionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2DownloadHostIgnitionOK creates a V2DownloadHostIgnitionOK with default headers values
func NewV2DownloadHostIgnitionOK(writer io.Writer) *V2DownloadHostIgnitionOK {
	return &V2DownloadHostIgnitionOK{
		Payload: writer,
	}
}

/*V2DownloadHostIgnitionOK handles this case with default header values.

Success.
*/
type V2DownloadHostIgnitionOK struct {
	Payload io.Writer
}

func (o *V2DownloadHostIgnitionOK) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionOK  %+v", 200, o.Payload)
}

func (o *V2DownloadHostIgnitionOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V2DownloadHostIgnitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionUnauthorized creates a V2DownloadHostIgnitionUnauthorized with default headers values
func NewV2DownloadHostIgnitionUnauthorized() *V2DownloadHostIgnitionUnauthorized {
	return &V2DownloadHostIgnitionUnauthorized{}
}

/*V2DownloadHostIgnitionUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2DownloadHostIgnitionUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2DownloadHostIgnitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionUnauthorized  %+v", 401, o.Payload)
}

func (o *V2DownloadHostIgnitionUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2DownloadHostIgnitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionForbidden creates a V2DownloadHostIgnitionForbidden with default headers values
func NewV2DownloadHostIgnitionForbidden() *V2DownloadHostIgnitionForbidden {
	return &V2DownloadHostIgnitionForbidden{}
}

/*V2DownloadHostIgnitionForbidden handles this case with default header values.

Forbidden.
*/
type V2DownloadHostIgnitionForbidden struct {
	Payload *models.InfraError
}

func (o *V2DownloadHostIgnitionForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionForbidden  %+v", 403, o.Payload)
}

func (o *V2DownloadHostIgnitionForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2DownloadHostIgnitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionNotFound creates a V2DownloadHostIgnitionNotFound with default headers values
func NewV2DownloadHostIgnitionNotFound() *V2DownloadHostIgnitionNotFound {
	return &V2DownloadHostIgnitionNotFound{}
}

/*V2DownloadHostIgnitionNotFound handles this case with default header values.

Error.
*/
type V2DownloadHostIgnitionNotFound struct {
	Payload *models.Error
}

func (o *V2DownloadHostIgnitionNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionNotFound  %+v", 404, o.Payload)
}

func (o *V2DownloadHostIgnitionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadHostIgnitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionMethodNotAllowed creates a V2DownloadHostIgnitionMethodNotAllowed with default headers values
func NewV2DownloadHostIgnitionMethodNotAllowed() *V2DownloadHostIgnitionMethodNotAllowed {
	return &V2DownloadHostIgnitionMethodNotAllowed{}
}

/*V2DownloadHostIgnitionMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2DownloadHostIgnitionMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2DownloadHostIgnitionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2DownloadHostIgnitionMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadHostIgnitionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionConflict creates a V2DownloadHostIgnitionConflict with default headers values
func NewV2DownloadHostIgnitionConflict() *V2DownloadHostIgnitionConflict {
	return &V2DownloadHostIgnitionConflict{}
}

/*V2DownloadHostIgnitionConflict handles this case with default header values.

Error.
*/
type V2DownloadHostIgnitionConflict struct {
	Payload *models.Error
}

func (o *V2DownloadHostIgnitionConflict) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionConflict  %+v", 409, o.Payload)
}

func (o *V2DownloadHostIgnitionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadHostIgnitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionInternalServerError creates a V2DownloadHostIgnitionInternalServerError with default headers values
func NewV2DownloadHostIgnitionInternalServerError() *V2DownloadHostIgnitionInternalServerError {
	return &V2DownloadHostIgnitionInternalServerError{}
}

/*V2DownloadHostIgnitionInternalServerError handles this case with default header values.

Error.
*/
type V2DownloadHostIgnitionInternalServerError struct {
	Payload *models.Error
}

func (o *V2DownloadHostIgnitionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionInternalServerError  %+v", 500, o.Payload)
}

func (o *V2DownloadHostIgnitionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadHostIgnitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadHostIgnitionServiceUnavailable creates a V2DownloadHostIgnitionServiceUnavailable with default headers values
func NewV2DownloadHostIgnitionServiceUnavailable() *V2DownloadHostIgnitionServiceUnavailable {
	return &V2DownloadHostIgnitionServiceUnavailable{}
}

/*V2DownloadHostIgnitionServiceUnavailable handles this case with default header values.

Unavailable.
*/
type V2DownloadHostIgnitionServiceUnavailable struct {
	Payload *models.Error
}

func (o *V2DownloadHostIgnitionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/infra-env/{infra_env_id}/hosts/{host_id}/downloads/ignition][%d] v2DownloadHostIgnitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2DownloadHostIgnitionServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadHostIgnitionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

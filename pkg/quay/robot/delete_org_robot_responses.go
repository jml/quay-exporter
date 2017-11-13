// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// DeleteOrgRobotReader is a Reader for the DeleteOrgRobot structure.
type DeleteOrgRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOrgRobotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOrgRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteOrgRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteOrgRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrgRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrgRobotNoContent creates a DeleteOrgRobotNoContent with default headers values
func NewDeleteOrgRobotNoContent() *DeleteOrgRobotNoContent {
	return &DeleteOrgRobotNoContent{}
}

/*DeleteOrgRobotNoContent handles this case with default header values.

Deleted
*/
type DeleteOrgRobotNoContent struct {
}

func (o *DeleteOrgRobotNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] deleteOrgRobotNoContent ", 204)
}

func (o *DeleteOrgRobotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgRobotBadRequest creates a DeleteOrgRobotBadRequest with default headers values
func NewDeleteOrgRobotBadRequest() *DeleteOrgRobotBadRequest {
	return &DeleteOrgRobotBadRequest{}
}

/*DeleteOrgRobotBadRequest handles this case with default header values.

Bad Request
*/
type DeleteOrgRobotBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteOrgRobotBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] deleteOrgRobotBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrgRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgRobotUnauthorized creates a DeleteOrgRobotUnauthorized with default headers values
func NewDeleteOrgRobotUnauthorized() *DeleteOrgRobotUnauthorized {
	return &DeleteOrgRobotUnauthorized{}
}

/*DeleteOrgRobotUnauthorized handles this case with default header values.

Session required
*/
type DeleteOrgRobotUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteOrgRobotUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] deleteOrgRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrgRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgRobotForbidden creates a DeleteOrgRobotForbidden with default headers values
func NewDeleteOrgRobotForbidden() *DeleteOrgRobotForbidden {
	return &DeleteOrgRobotForbidden{}
}

/*DeleteOrgRobotForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteOrgRobotForbidden struct {
	Payload *models.APIError
}

func (o *DeleteOrgRobotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] deleteOrgRobotForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrgRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgRobotNotFound creates a DeleteOrgRobotNotFound with default headers values
func NewDeleteOrgRobotNotFound() *DeleteOrgRobotNotFound {
	return &DeleteOrgRobotNotFound{}
}

/*DeleteOrgRobotNotFound handles this case with default header values.

Not found
*/
type DeleteOrgRobotNotFound struct {
	Payload *models.APIError
}

func (o *DeleteOrgRobotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] deleteOrgRobotNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrgRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
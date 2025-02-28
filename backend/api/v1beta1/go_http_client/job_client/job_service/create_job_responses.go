// Code generated by go-swagger; DO NOT EDIT.

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	job_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/job_model"
)

// CreateJobReader is a Reader for the CreateJob structure.
type CreateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateJobOK creates a CreateJobOK with default headers values
func NewCreateJobOK() *CreateJobOK {
	return &CreateJobOK{}
}

/*CreateJobOK handles this case with default header values.

A successful response.
*/
type CreateJobOK struct {
	Payload *job_model.V1beta1Job
}

func (o *CreateJobOK) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/jobs][%d] createJobOK  %+v", 200, o.Payload)
}

func (o *CreateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.V1beta1Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobDefault creates a CreateJobDefault with default headers values
func NewCreateJobDefault(code int) *CreateJobDefault {
	return &CreateJobDefault{
		_statusCode: code,
	}
}

/*CreateJobDefault handles this case with default header values.

CreateJobDefault create job default
*/
type CreateJobDefault struct {
	_statusCode int

	Payload *job_model.V1beta1Status
}

// Code gets the status code for the create job default response
func (o *CreateJobDefault) Code() int {
	return o._statusCode
}

func (o *CreateJobDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/jobs][%d] CreateJob default  %+v", o._statusCode, o.Payload)
}

func (o *CreateJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.V1beta1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

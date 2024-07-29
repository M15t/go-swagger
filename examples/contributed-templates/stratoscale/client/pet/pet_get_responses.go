// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/M15t/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetGetReader is a Reader for the PetGet structure.
type PetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPetGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pet/{petId}] PetGet", response, response.Code())
	}
}

// NewPetGetOK creates a PetGetOK with default headers values
func NewPetGetOK() *PetGetOK {
	return &PetGetOK{}
}

/*
PetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type PetGetOK struct {
	Payload *models.Pet
}

// IsSuccess returns true when this pet get o k response has a 2xx status code
func (o *PetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pet get o k response has a 3xx status code
func (o *PetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet get o k response has a 4xx status code
func (o *PetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pet get o k response has a 5xx status code
func (o *PetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pet get o k response a status code equal to that given
func (o *PetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pet get o k response
func (o *PetGetOK) Code() int {
	return 200
}

func (o *PetGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pet/{petId}][%d] petGetOK %s", 200, payload)
}

func (o *PetGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pet/{petId}][%d] petGetOK %s", 200, payload)
}

func (o *PetGetOK) GetPayload() *models.Pet {
	return o.Payload
}

func (o *PetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPetGetBadRequest creates a PetGetBadRequest with default headers values
func NewPetGetBadRequest() *PetGetBadRequest {
	return &PetGetBadRequest{}
}

/*
PetGetBadRequest describes a response with status code 400, with default header values.

Invalid ID supplied
*/
type PetGetBadRequest struct {
}

// IsSuccess returns true when this pet get bad request response has a 2xx status code
func (o *PetGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pet get bad request response has a 3xx status code
func (o *PetGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet get bad request response has a 4xx status code
func (o *PetGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pet get bad request response has a 5xx status code
func (o *PetGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pet get bad request response a status code equal to that given
func (o *PetGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pet get bad request response
func (o *PetGetBadRequest) Code() int {
	return 400
}

func (o *PetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] petGetBadRequest", 400)
}

func (o *PetGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] petGetBadRequest", 400)
}

func (o *PetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetGetNotFound creates a PetGetNotFound with default headers values
func NewPetGetNotFound() *PetGetNotFound {
	return &PetGetNotFound{}
}

/*
PetGetNotFound describes a response with status code 404, with default header values.

Pet not found
*/
type PetGetNotFound struct {
}

// IsSuccess returns true when this pet get not found response has a 2xx status code
func (o *PetGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pet get not found response has a 3xx status code
func (o *PetGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet get not found response has a 4xx status code
func (o *PetGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pet get not found response has a 5xx status code
func (o *PetGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pet get not found response a status code equal to that given
func (o *PetGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pet get not found response
func (o *PetGetNotFound) Code() int {
	return 404
}

func (o *PetGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] petGetNotFound", 404)
}

func (o *PetGetNotFound) String() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] petGetNotFound", 404)
}

func (o *PetGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

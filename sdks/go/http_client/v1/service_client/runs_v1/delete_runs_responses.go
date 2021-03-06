// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteRunsReader is a Reader for the DeleteRuns structure.
type DeleteRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRunsOK creates a DeleteRunsOK with default headers values
func NewDeleteRunsOK() *DeleteRunsOK {
	return &DeleteRunsOK{}
}

/*DeleteRunsOK handles this case with default header values.

A successful response.
*/
type DeleteRunsOK struct {
}

func (o *DeleteRunsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsOK ", 200)
}

func (o *DeleteRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunsNoContent creates a DeleteRunsNoContent with default headers values
func NewDeleteRunsNoContent() *DeleteRunsNoContent {
	return &DeleteRunsNoContent{}
}

/*DeleteRunsNoContent handles this case with default header values.

No content.
*/
type DeleteRunsNoContent struct {
	Payload interface{}
}

func (o *DeleteRunsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunsForbidden creates a DeleteRunsForbidden with default headers values
func NewDeleteRunsForbidden() *DeleteRunsForbidden {
	return &DeleteRunsForbidden{}
}

/*DeleteRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteRunsForbidden struct {
	Payload interface{}
}

func (o *DeleteRunsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunsNotFound creates a DeleteRunsNotFound with default headers values
func NewDeleteRunsNotFound() *DeleteRunsNotFound {
	return &DeleteRunsNotFound{}
}

/*DeleteRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteRunsNotFound struct {
	Payload interface{}
}

func (o *DeleteRunsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunsDefault creates a DeleteRunsDefault with default headers values
func NewDeleteRunsDefault(code int) *DeleteRunsDefault {
	return &DeleteRunsDefault{
		_statusCode: code,
	}
}

/*DeleteRunsDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete runs default response
func (o *DeleteRunsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRunsDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] DeleteRuns default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

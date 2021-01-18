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

// GetRunClonesLineageReader is a Reader for the GetRunClonesLineage structure.
type GetRunClonesLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunClonesLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunClonesLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunClonesLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunClonesLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunClonesLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunClonesLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunClonesLineageOK creates a GetRunClonesLineageOK with default headers values
func NewGetRunClonesLineageOK() *GetRunClonesLineageOK {
	return &GetRunClonesLineageOK{}
}

/*GetRunClonesLineageOK handles this case with default header values.

A successful response.
*/
type GetRunClonesLineageOK struct {
	Payload *service_model.V1ListRunsResponse
}

func (o *GetRunClonesLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/clones][%d] getRunClonesLineageOK  %+v", 200, o.Payload)
}

func (o *GetRunClonesLineageOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *GetRunClonesLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunClonesLineageNoContent creates a GetRunClonesLineageNoContent with default headers values
func NewGetRunClonesLineageNoContent() *GetRunClonesLineageNoContent {
	return &GetRunClonesLineageNoContent{}
}

/*GetRunClonesLineageNoContent handles this case with default header values.

No content.
*/
type GetRunClonesLineageNoContent struct {
	Payload interface{}
}

func (o *GetRunClonesLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/clones][%d] getRunClonesLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetRunClonesLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunClonesLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunClonesLineageForbidden creates a GetRunClonesLineageForbidden with default headers values
func NewGetRunClonesLineageForbidden() *GetRunClonesLineageForbidden {
	return &GetRunClonesLineageForbidden{}
}

/*GetRunClonesLineageForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetRunClonesLineageForbidden struct {
	Payload interface{}
}

func (o *GetRunClonesLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/clones][%d] getRunClonesLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetRunClonesLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunClonesLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunClonesLineageNotFound creates a GetRunClonesLineageNotFound with default headers values
func NewGetRunClonesLineageNotFound() *GetRunClonesLineageNotFound {
	return &GetRunClonesLineageNotFound{}
}

/*GetRunClonesLineageNotFound handles this case with default header values.

Resource does not exist.
*/
type GetRunClonesLineageNotFound struct {
	Payload interface{}
}

func (o *GetRunClonesLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/clones][%d] getRunClonesLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetRunClonesLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunClonesLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunClonesLineageDefault creates a GetRunClonesLineageDefault with default headers values
func NewGetRunClonesLineageDefault(code int) *GetRunClonesLineageDefault {
	return &GetRunClonesLineageDefault{
		_statusCode: code,
	}
}

/*GetRunClonesLineageDefault handles this case with default header values.

An unexpected error response.
*/
type GetRunClonesLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get run clones lineage default response
func (o *GetRunClonesLineageDefault) Code() int {
	return o._statusCode
}

func (o *GetRunClonesLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/clones][%d] GetRunClonesLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunClonesLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunClonesLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
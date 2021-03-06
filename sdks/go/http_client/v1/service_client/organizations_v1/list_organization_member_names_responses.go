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

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListOrganizationMemberNamesReader is a Reader for the ListOrganizationMemberNames structure.
type ListOrganizationMemberNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationMemberNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationMemberNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListOrganizationMemberNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListOrganizationMemberNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOrganizationMemberNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListOrganizationMemberNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrganizationMemberNamesOK creates a ListOrganizationMemberNamesOK with default headers values
func NewListOrganizationMemberNamesOK() *ListOrganizationMemberNamesOK {
	return &ListOrganizationMemberNamesOK{}
}

/*ListOrganizationMemberNamesOK handles this case with default header values.

A successful response.
*/
type ListOrganizationMemberNamesOK struct {
	Payload *service_model.V1ListOrganizationMembersResponse
}

func (o *ListOrganizationMemberNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members/names][%d] listOrganizationMemberNamesOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationMemberNamesOK) GetPayload() *service_model.V1ListOrganizationMembersResponse {
	return o.Payload
}

func (o *ListOrganizationMemberNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListOrganizationMembersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMemberNamesNoContent creates a ListOrganizationMemberNamesNoContent with default headers values
func NewListOrganizationMemberNamesNoContent() *ListOrganizationMemberNamesNoContent {
	return &ListOrganizationMemberNamesNoContent{}
}

/*ListOrganizationMemberNamesNoContent handles this case with default header values.

No content.
*/
type ListOrganizationMemberNamesNoContent struct {
	Payload interface{}
}

func (o *ListOrganizationMemberNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members/names][%d] listOrganizationMemberNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationMemberNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationMemberNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMemberNamesForbidden creates a ListOrganizationMemberNamesForbidden with default headers values
func NewListOrganizationMemberNamesForbidden() *ListOrganizationMemberNamesForbidden {
	return &ListOrganizationMemberNamesForbidden{}
}

/*ListOrganizationMemberNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListOrganizationMemberNamesForbidden struct {
	Payload interface{}
}

func (o *ListOrganizationMemberNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members/names][%d] listOrganizationMemberNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationMemberNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationMemberNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMemberNamesNotFound creates a ListOrganizationMemberNamesNotFound with default headers values
func NewListOrganizationMemberNamesNotFound() *ListOrganizationMemberNamesNotFound {
	return &ListOrganizationMemberNamesNotFound{}
}

/*ListOrganizationMemberNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListOrganizationMemberNamesNotFound struct {
	Payload interface{}
}

func (o *ListOrganizationMemberNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members/names][%d] listOrganizationMemberNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationMemberNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationMemberNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMemberNamesDefault creates a ListOrganizationMemberNamesDefault with default headers values
func NewListOrganizationMemberNamesDefault(code int) *ListOrganizationMemberNamesDefault {
	return &ListOrganizationMemberNamesDefault{
		_statusCode: code,
	}
}

/*ListOrganizationMemberNamesDefault handles this case with default header values.

An unexpected error response.
*/
type ListOrganizationMemberNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list organization member names default response
func (o *ListOrganizationMemberNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListOrganizationMemberNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members/names][%d] ListOrganizationMemberNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationMemberNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListOrganizationMemberNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

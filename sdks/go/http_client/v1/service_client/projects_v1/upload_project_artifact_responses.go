// Copyright 2019 Polyaxon, Inc.
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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UploadProjectArtifactReader is a Reader for the UploadProjectArtifact structure.
type UploadProjectArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadProjectArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadProjectArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUploadProjectArtifactNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUploadProjectArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadProjectArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadProjectArtifactOK creates a UploadProjectArtifactOK with default headers values
func NewUploadProjectArtifactOK() *UploadProjectArtifactOK {
	return &UploadProjectArtifactOK{}
}

/*UploadProjectArtifactOK handles this case with default header values.

A successful response.
*/
type UploadProjectArtifactOK struct {
}

func (o *UploadProjectArtifactOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/artifacts/{uuid}/upload][%d] uploadProjectArtifactOK ", 200)
}

func (o *UploadProjectArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProjectArtifactNoContent creates a UploadProjectArtifactNoContent with default headers values
func NewUploadProjectArtifactNoContent() *UploadProjectArtifactNoContent {
	return &UploadProjectArtifactNoContent{}
}

/*UploadProjectArtifactNoContent handles this case with default header values.

No content.
*/
type UploadProjectArtifactNoContent struct {
	Payload interface{}
}

func (o *UploadProjectArtifactNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/artifacts/{uuid}/upload][%d] uploadProjectArtifactNoContent  %+v", 204, o.Payload)
}

func (o *UploadProjectArtifactNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadProjectArtifactNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadProjectArtifactForbidden creates a UploadProjectArtifactForbidden with default headers values
func NewUploadProjectArtifactForbidden() *UploadProjectArtifactForbidden {
	return &UploadProjectArtifactForbidden{}
}

/*UploadProjectArtifactForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UploadProjectArtifactForbidden struct {
	Payload interface{}
}

func (o *UploadProjectArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/artifacts/{uuid}/upload][%d] uploadProjectArtifactForbidden  %+v", 403, o.Payload)
}

func (o *UploadProjectArtifactForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadProjectArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadProjectArtifactNotFound creates a UploadProjectArtifactNotFound with default headers values
func NewUploadProjectArtifactNotFound() *UploadProjectArtifactNotFound {
	return &UploadProjectArtifactNotFound{}
}

/*UploadProjectArtifactNotFound handles this case with default header values.

Resource does not exist.
*/
type UploadProjectArtifactNotFound struct {
	Payload interface{}
}

func (o *UploadProjectArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/artifacts/{uuid}/upload][%d] uploadProjectArtifactNotFound  %+v", 404, o.Payload)
}

func (o *UploadProjectArtifactNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadProjectArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

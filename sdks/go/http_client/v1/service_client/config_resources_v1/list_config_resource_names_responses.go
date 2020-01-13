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

package config_resources_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListConfigResourceNamesReader is a Reader for the ListConfigResourceNames structure.
type ListConfigResourceNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConfigResourceNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConfigResourceNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListConfigResourceNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListConfigResourceNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListConfigResourceNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListConfigResourceNamesOK creates a ListConfigResourceNamesOK with default headers values
func NewListConfigResourceNamesOK() *ListConfigResourceNamesOK {
	return &ListConfigResourceNamesOK{}
}

/*ListConfigResourceNamesOK handles this case with default header values.

A successful response.
*/
type ListConfigResourceNamesOK struct {
	Payload *service_model.V1ListConfigResourcesResponse
}

func (o *ListConfigResourceNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/config_resources/names][%d] listConfigResourceNamesOK  %+v", 200, o.Payload)
}

func (o *ListConfigResourceNamesOK) GetPayload() *service_model.V1ListConfigResourcesResponse {
	return o.Payload
}

func (o *ListConfigResourceNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListConfigResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConfigResourceNamesNoContent creates a ListConfigResourceNamesNoContent with default headers values
func NewListConfigResourceNamesNoContent() *ListConfigResourceNamesNoContent {
	return &ListConfigResourceNamesNoContent{}
}

/*ListConfigResourceNamesNoContent handles this case with default header values.

No content.
*/
type ListConfigResourceNamesNoContent struct {
	Payload interface{}
}

func (o *ListConfigResourceNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/config_resources/names][%d] listConfigResourceNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListConfigResourceNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConfigResourceNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConfigResourceNamesForbidden creates a ListConfigResourceNamesForbidden with default headers values
func NewListConfigResourceNamesForbidden() *ListConfigResourceNamesForbidden {
	return &ListConfigResourceNamesForbidden{}
}

/*ListConfigResourceNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListConfigResourceNamesForbidden struct {
	Payload interface{}
}

func (o *ListConfigResourceNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/config_resources/names][%d] listConfigResourceNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListConfigResourceNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConfigResourceNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConfigResourceNamesNotFound creates a ListConfigResourceNamesNotFound with default headers values
func NewListConfigResourceNamesNotFound() *ListConfigResourceNamesNotFound {
	return &ListConfigResourceNamesNotFound{}
}

/*ListConfigResourceNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListConfigResourceNamesNotFound struct {
	Payload interface{}
}

func (o *ListConfigResourceNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/config_resources/names][%d] listConfigResourceNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListConfigResourceNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConfigResourceNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
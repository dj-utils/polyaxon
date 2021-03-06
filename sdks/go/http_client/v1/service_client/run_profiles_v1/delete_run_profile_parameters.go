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

package run_profiles_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRunProfileParams creates a new DeleteRunProfileParams object
// with the default values initialized.
func NewDeleteRunProfileParams() *DeleteRunProfileParams {
	var ()
	return &DeleteRunProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunProfileParamsWithTimeout creates a new DeleteRunProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRunProfileParamsWithTimeout(timeout time.Duration) *DeleteRunProfileParams {
	var ()
	return &DeleteRunProfileParams{

		timeout: timeout,
	}
}

// NewDeleteRunProfileParamsWithContext creates a new DeleteRunProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRunProfileParamsWithContext(ctx context.Context) *DeleteRunProfileParams {
	var ()
	return &DeleteRunProfileParams{

		Context: ctx,
	}
}

// NewDeleteRunProfileParamsWithHTTPClient creates a new DeleteRunProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRunProfileParamsWithHTTPClient(client *http.Client) *DeleteRunProfileParams {
	var ()
	return &DeleteRunProfileParams{
		HTTPClient: client,
	}
}

/*DeleteRunProfileParams contains all the parameters to send to the API endpoint
for the delete run profile operation typically these are written to a http.Request
*/
type DeleteRunProfileParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Unique integer identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete run profile params
func (o *DeleteRunProfileParams) WithTimeout(timeout time.Duration) *DeleteRunProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete run profile params
func (o *DeleteRunProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete run profile params
func (o *DeleteRunProfileParams) WithContext(ctx context.Context) *DeleteRunProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete run profile params
func (o *DeleteRunProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete run profile params
func (o *DeleteRunProfileParams) WithHTTPClient(client *http.Client) *DeleteRunProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete run profile params
func (o *DeleteRunProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the delete run profile params
func (o *DeleteRunProfileParams) WithOwner(owner string) *DeleteRunProfileParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete run profile params
func (o *DeleteRunProfileParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the delete run profile params
func (o *DeleteRunProfileParams) WithUUID(uuid string) *DeleteRunProfileParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete run profile params
func (o *DeleteRunProfileParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

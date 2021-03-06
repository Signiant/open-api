package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListHookTypesParams creates a new ListHookTypesParams object
// with the default values initialized.
func NewListHookTypesParams() *ListHookTypesParams {

	return &ListHookTypesParams{}
}

/*ListHookTypesParams contains all the parameters to send to the API endpoint
for the list hook types operation typically these are written to a http.Request
*/
type ListHookTypesParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *ListHookTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

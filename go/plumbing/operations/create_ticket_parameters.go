package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateTicketParams creates a new CreateTicketParams object
// with the default values initialized.
func NewCreateTicketParams() *CreateTicketParams {
	var ()
	return &CreateTicketParams{}
}

/*CreateTicketParams contains all the parameters to send to the API endpoint
for the create ticket operation typically these are written to a http.Request
*/
type CreateTicketParams struct {

	/*ClientID*/
	ClientID string
}

// WithClientID adds the clientId to the create ticket params
func (o *CreateTicketParams) WithClientID(ClientID string) *CreateTicketParams {
	o.ClientID = ClientID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// query param client_id
	qrClientID := o.ClientID
	qClientID := qrClientID
	if qClientID != "" {
		if err := r.SetQueryParam("client_id", qClientID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

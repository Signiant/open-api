package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSiteSnippetParams creates a new DeleteSiteSnippetParams object
// with the default values initialized.
func NewDeleteSiteSnippetParams() *DeleteSiteSnippetParams {
	var ()
	return &DeleteSiteSnippetParams{}
}

/*DeleteSiteSnippetParams contains all the parameters to send to the API endpoint
for the delete site snippet operation typically these are written to a http.Request
*/
type DeleteSiteSnippetParams struct {

	/*SiteID*/
	SiteID string
	/*SnippetID*/
	SnippetID string
}

// WithSiteID adds the siteId to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithSiteID(SiteID string) *DeleteSiteSnippetParams {
	o.SiteID = SiteID
	return o
}

// WithSnippetID adds the snippetId to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithSnippetID(SnippetID string) *DeleteSiteSnippetParams {
	o.SnippetID = SnippetID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSiteSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param snippet_id
	if err := r.SetPathParam("snippet_id", o.SnippetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

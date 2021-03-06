package plumbing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/plumbing/operations"
)

// Default netlify HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new netlify HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Netlify {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("api.netlify.com", "/api/v1", []string{"https"})
	return New(transport, formats)
}

// New creates a new netlify client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Netlify {
	cli := new(Netlify)
	cli.Transport = transport

	cli.Operations = operations.New(transport, formats)

	return cli
}

// Netlify is a client for netlify
type Netlify struct {
	Operations *operations.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Netlify) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Operations.SetTransport(transport)

}

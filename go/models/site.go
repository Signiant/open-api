package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Site site

swagger:model site
*/
type Site struct {

	/* admin url
	 */
	AdminURL string `json:"admin_url,omitempty"`

	/* created at
	 */
	CreatedAt string `json:"created_at,omitempty"`

	/* custom domain
	 */
	CustomDomain string `json:"custom_domain,omitempty"`

	/* deploy url
	 */
	DeployURL string `json:"deploy_url,omitempty"`

	/* domain aliases
	 */
	DomainAliases []string `json:"domain_aliases,omitempty"`

	/* force ssl
	 */
	ForceSsl bool `json:"force_ssl,omitempty"`

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* managed dns
	 */
	ManagedDNS bool `json:"managed_dns,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* notification email
	 */
	NotificationEmail string `json:"notification_email,omitempty"`

	/* password
	 */
	Password string `json:"password,omitempty"`

	/* plan
	 */
	Plan string `json:"plan,omitempty"`

	/* published deploy
	 */
	PublishedDeploy *Deploy `json:"published_deploy,omitempty"`

	/* screenshot url
	 */
	ScreenshotURL string `json:"screenshot_url,omitempty"`

	/* ssl
	 */
	Ssl bool `json:"ssl,omitempty"`

	/* state
	 */
	State string `json:"state,omitempty"`

	/* updated at
	 */
	UpdatedAt string `json:"updated_at,omitempty"`

	/* url
	 */
	URL string `json:"url,omitempty"`

	/* user id
	 */
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainAliases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateDomainAliases(formats strfmt.Registry) error {

	if swag.IsZero(m.DomainAliases) { // not required
		return nil
	}

	return nil
}

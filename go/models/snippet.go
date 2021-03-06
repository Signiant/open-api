package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Snippet snippet

swagger:model snippet
*/
type Snippet struct {

	/* general
	 */
	General string `json:"general,omitempty"`

	/* general position
	 */
	GeneralPosition string `json:"general_position,omitempty"`

	/* goal
	 */
	Goal string `json:"goal,omitempty"`

	/* goal position
	 */
	GoalPosition string `json:"goal_position,omitempty"`

	/* id
	 */
	ID int32 `json:"id,omitempty"`

	/* site id
	 */
	SiteID string `json:"site_id,omitempty"`

	/* title
	 */
	Title string `json:"title,omitempty"`
}

// Validate validates this snippet
func (m *Snippet) Validate(formats strfmt.Registry) error {
	return nil
}

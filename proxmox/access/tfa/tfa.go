// File generated by proxmox json schema, DO NOT EDIT

package tfa

import (
	"context"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
)

type HTTPClient interface {
	Do(context.Context, string, string, interface{}, interface{}) error
}

type Client struct {
	httpClient HTTPClient
}

func New(c HTTPClient) *Client {
	return &Client{
		httpClient: c,
	}
}

type IndexResponse []*struct {
	Entries []*struct {
		Created     int    `url:"created",json:"created"`         // Creation time of this entry as unix epoch.
		Description string `url:"description",json:"description"` // User chosen description for this entry.
		Id          string `url:"id",json:"id"`                   // The id used to reference this entry.
		Type        string `url:"type",json:"type"`               // TFA Entry Type.

		// The following parameters are optional
		Enable *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"` // Whether this TFA entry is currently enabled.
	} `url:"entries",json:"entries"`
	Userid string `url:"userid",json:"userid"` // User this entry belongs to.

}

// Index List TFA configurations of users.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/access/tfa", "GET", &resp, nil)
	return resp, err
}

type CreateRequest struct {
	Response string `url:"response",json:"response"` // The response to the current authentication challenge.

}

type CreateResponse struct {
	Ticket string `url:"ticket",json:"ticket"`
}

// Create Finish a u2f challenge.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/access/tfa", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Userid string `url:"userid",json:"userid"` // User ID

}

type FindResponse []*struct {
	Created     int    `url:"created",json:"created"`         // Creation time of this entry as unix epoch.
	Description string `url:"description",json:"description"` // User chosen description for this entry.
	Id          string `url:"id",json:"id"`                   // The id used to reference this entry.
	Type        string `url:"type",json:"type"`               // TFA Entry Type.

	// The following parameters are optional
	Enable *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"` // Whether this TFA entry is currently enabled.
}

// Find List TFA configurations of users.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/access/tfa/{userid}", "GET", &resp, req)
	return resp, err
}

type ChildCreateRequest struct {
	Type   string `url:"type",json:"type"`     // TFA Entry Type.
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Challenge   *string `url:"challenge,omitempty",json:"challenge,omitempty"`     // When responding to a u2f challenge: the original challenge string
	Description *string `url:"description,omitempty",json:"description,omitempty"` // A description to distinguish multiple entries from one another
	Password    *string `url:"password,omitempty",json:"password,omitempty"`       // The current password.
	Totp        *string `url:"totp,omitempty",json:"totp,omitempty"`               // A totp URI.
	Value       *string `url:"value,omitempty",json:"value,omitempty"`             // The current value for the provided totp URI, or a Webauthn/U2F challenge response
}

type ChildCreateResponse struct {
	Id string `url:"id",json:"id"` // The id of a newly added TFA entry.

	// The following parameters are optional
	Challenge *string  `url:"challenge,omitempty",json:"challenge,omitempty"` // When adding u2f entries, this contains a challenge the user must respond to in order to finish the registration.
	Recovery  []string `url:"recovery,omitempty",json:"recovery,omitempty"`   // When adding recovery codes, this contains the list of codes to be displayed to the user
}

// ChildCreate Add a TFA entry for a user.
func (c *Client) ChildCreate(ctx context.Context, req *ChildCreateRequest) (*ChildCreateResponse, error) {
	var resp *ChildCreateResponse

	err := c.httpClient.Do(ctx, "/access/tfa/{userid}", "POST", &resp, req)
	return resp, err
}

type GetTfaEntryIdRequest struct {
	Id     string `url:"id",json:"id"`         // A TFA entry id.
	Userid string `url:"userid",json:"userid"` // User ID

}

type GetTfaEntryIdResponse struct {
	Created     int    `url:"created",json:"created"`         // Creation time of this entry as unix epoch.
	Description string `url:"description",json:"description"` // User chosen description for this entry.
	Id          string `url:"id",json:"id"`                   // The id used to reference this entry.
	Type        string `url:"type",json:"type"`               // TFA Entry Type.

	// The following parameters are optional
	Enable *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"` // Whether this TFA entry is currently enabled.
}

// GetTfaEntryId Fetch a requested TFA entry if present.
func (c *Client) GetTfaEntryId(ctx context.Context, req *GetTfaEntryIdRequest) (*GetTfaEntryIdResponse, error) {
	var resp *GetTfaEntryIdResponse

	err := c.httpClient.Do(ctx, "/access/tfa/{userid}/{id}", "GET", &resp, req)
	return resp, err
}

type UpdateTfaEntryIdRequest struct {
	Id     string `url:"id",json:"id"`         // A TFA entry id.
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Description *string           `url:"description,omitempty",json:"description,omitempty"` // A description to distinguish multiple entries from one another
	Enable      *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"`           // Whether the entry should be enabled for login.
	Password    *string           `url:"password,omitempty",json:"password,omitempty"`       // The current password.
}

type UpdateTfaEntryIdResponse map[string]interface{}

// UpdateTfaEntryId Add a TFA entry for a user.
func (c *Client) UpdateTfaEntryId(ctx context.Context, req *UpdateTfaEntryIdRequest) (*UpdateTfaEntryIdResponse, error) {
	var resp *UpdateTfaEntryIdResponse

	err := c.httpClient.Do(ctx, "/access/tfa/{userid}/{id}", "PUT", &resp, req)
	return resp, err
}

type DeleteTfaIdRequest struct {
	Id     string `url:"id",json:"id"`         // A TFA entry id.
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Password *string `url:"password,omitempty",json:"password,omitempty"` // The current password.
}

type DeleteTfaIdResponse map[string]interface{}

// DeleteTfaId Delete a TFA entry by ID.
func (c *Client) DeleteTfaId(ctx context.Context, req *DeleteTfaIdRequest) (*DeleteTfaIdResponse, error) {
	var resp *DeleteTfaIdResponse

	err := c.httpClient.Do(ctx, "/access/tfa/{userid}/{id}", "DELETE", &resp, req)
	return resp, err
}

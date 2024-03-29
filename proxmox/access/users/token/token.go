// File generated by proxmox json schema, DO NOT EDIT

package token

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

type IndexRequest struct {
	Userid string `url:"userid" json:"userid"` // Full User ID, in the `name@realm` format.

}
type _IndexRequest IndexRequest

type IndexResponse struct {
	Tokenid string `url:"tokenid" json:"tokenid"` // User-specific token identifier.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}
type _IndexResponse IndexResponse

type FindRequest struct {
	Tokenid string `url:"tokenid" json:"tokenid"` // User-specific token identifier.
	Userid  string `url:"userid" json:"userid"`   // Full User ID, in the `name@realm` format.

}
type _FindRequest FindRequest

type FindResponse struct {

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}
type _FindResponse FindResponse

type ChildCreateRequest struct {
	Tokenid string `url:"tokenid" json:"tokenid"` // User-specific token identifier.
	Userid  string `url:"userid" json:"userid"`   // Full User ID, in the `name@realm` format.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}
type _ChildCreateRequest ChildCreateRequest

type Info struct {

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}
type _Info Info

type ChildCreateResponse struct {
	FullTokenid string `url:"full-tokenid" json:"full-tokenid"` // The full token id.
	Info        Info   `url:"info" json:"info"`
	Value       string `url:"value" json:"value"` // API token value used for authentication.

}
type _ChildCreateResponse ChildCreateResponse

type UpdateRequest struct {
	Tokenid string `url:"tokenid" json:"tokenid"` // User-specific token identifier.
	Userid  string `url:"userid" json:"userid"`   // Full User ID, in the `name@realm` format.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}
type _UpdateRequest UpdateRequest

// Updated token information.
type UpdateResponse struct {

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}
type _UpdateResponse UpdateResponse

type DeleteRequest struct {
	Tokenid string `url:"tokenid" json:"tokenid"` // User-specific token identifier.
	Userid  string `url:"userid" json:"userid"`   // Full User ID, in the `name@realm` format.

}
type _DeleteRequest DeleteRequest

// Index Get user API tokens.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}/token", "GET", &resp, req)
	return resp, err
}

// Find Get specific API token information.
func (c *Client) Find(ctx context.Context, req FindRequest) (FindResponse, error) {
	var resp FindResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}/token/{tokenid}", "GET", &resp, req)
	return resp, err
}

// ChildCreate Generate a new API token for a specific user. NOTE: returns API token value, which needs to be stored as it cannot be retrieved afterwards!
func (c *Client) ChildCreate(ctx context.Context, req ChildCreateRequest) (ChildCreateResponse, error) {
	var resp ChildCreateResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}/token/{tokenid}", "POST", &resp, req)
	return resp, err
}

// Update Update API token for a specific user.
func (c *Client) Update(ctx context.Context, req UpdateRequest) (UpdateResponse, error) {
	var resp UpdateResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}/token/{tokenid}", "PUT", &resp, req)
	return resp, err
}

// Delete Remove API token for a specific user.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/access/users/{userid}/token/{tokenid}", "DELETE", nil, req)
	return err
}

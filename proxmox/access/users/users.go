// File generated by proxmox json schema, DO NOT EDIT

package users

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

	// The following parameters are optional
	Enabled *util.SpecialBool `url:"enabled,omitempty",json:"enabled,omitempty"` // Optional filter for enable property.
	Full    *util.SpecialBool `url:"full,omitempty",json:"full,omitempty"`       // Include group and token information.
}

type IndexResponse []*struct {
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Comment   *string           `url:"comment,omitempty",json:"comment,omitempty"`
	Email     *string           `url:"email,omitempty",json:"email,omitempty"`
	Enable    *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int              `url:"expire,omitempty",json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string           `url:"firstname,omitempty",json:"firstname,omitempty"`
	Groups    *string           `url:"groups,omitempty",json:"groups,omitempty"`
	Keys      *string           `url:"keys,omitempty",json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string           `url:"lastname,omitempty",json:"lastname,omitempty"`
	RealmType *string           `url:"realm-type,omitempty",json:"realm-type,omitempty"` // The type of the users realm
	Tokens    []*struct {
		Tokenid string `url:"tokenid",json:"tokenid"` // User-specific token identifier.

		// The following parameters are optional
		Comment *string           `url:"comment,omitempty",json:"comment,omitempty"`
		Expire  *int              `url:"expire,omitempty",json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
		Privsep *util.SpecialBool `url:"privsep,omitempty",json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	} `url:"tokens,omitempty",json:"tokens,omitempty"`
}

// Index User index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/access/users", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Comment   *string           `url:"comment,omitempty",json:"comment,omitempty"`
	Email     *string           `url:"email,omitempty",json:"email,omitempty"`
	Enable    *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int              `url:"expire,omitempty",json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string           `url:"firstname,omitempty",json:"firstname,omitempty"`
	Groups    *string           `url:"groups,omitempty",json:"groups,omitempty"`
	Keys      *string           `url:"keys,omitempty",json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string           `url:"lastname,omitempty",json:"lastname,omitempty"`
	Password  *string           `url:"password,omitempty",json:"password,omitempty"` // Initial password.
}

type CreateResponse map[string]interface{}

// Create Create new user.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/access/users", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Userid string `url:"userid",json:"userid"` // User ID

}

type FindResponse struct {

	// The following parameters are optional
	Comment   *string                `url:"comment,omitempty",json:"comment,omitempty"`
	Email     *string                `url:"email,omitempty",json:"email,omitempty"`
	Enable    *util.SpecialBool      `url:"enable,omitempty",json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int                   `url:"expire,omitempty",json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string                `url:"firstname,omitempty",json:"firstname,omitempty"`
	Groups    []string               `url:"groups,omitempty",json:"groups,omitempty"`
	Keys      *string                `url:"keys,omitempty",json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string                `url:"lastname,omitempty",json:"lastname,omitempty"`
	Tokens    map[string]interface{} `url:"tokens,omitempty",json:"tokens,omitempty"`
}

// Find Get user configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Append    *util.SpecialBool `url:"append,omitempty",json:"append,omitempty"`
	Comment   *string           `url:"comment,omitempty",json:"comment,omitempty"`
	Email     *string           `url:"email,omitempty",json:"email,omitempty"`
	Enable    *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int              `url:"expire,omitempty",json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string           `url:"firstname,omitempty",json:"firstname,omitempty"`
	Groups    *string           `url:"groups,omitempty",json:"groups,omitempty"`
	Keys      *string           `url:"keys,omitempty",json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string           `url:"lastname,omitempty",json:"lastname,omitempty"`
}

type UpdateResponse map[string]interface{}

// Update Update user configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Userid string `url:"userid",json:"userid"` // User ID

}

type DeleteResponse map[string]interface{}

// Delete Delete user.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}", "DELETE", &resp, req)
	return resp, err
}

type ReadUserTfaTypeTfaRequest struct {
	Userid string `url:"userid",json:"userid"` // User ID

	// The following parameters are optional
	Multiple *util.SpecialBool `url:"multiple,omitempty",json:"multiple,omitempty"` // Request all entries as an array.
}

type ReadUserTfaTypeTfaResponse struct {

	// The following parameters are optional
	Realm *string  `url:"realm,omitempty",json:"realm,omitempty"` // The type of TFA the users realm has set, if any.
	Types []string `url:"types,omitempty",json:"types,omitempty"` // Array of the user configured TFA types, if any. Only available if 'multiple' was not passed.
	User  *string  `url:"user,omitempty",json:"user,omitempty"`   // The type of TFA the user has set, if any. Only set if 'multiple' was not passed.
}

// ReadUserTfaTypeTfa Get user TFA types (Personal and Realm).
func (c *Client) ReadUserTfaTypeTfa(ctx context.Context, req *ReadUserTfaTypeTfaRequest) (*ReadUserTfaTypeTfaResponse, error) {
	var resp *ReadUserTfaTypeTfaResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}/tfa", "GET", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package users

import (
	"context"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
)

const (
	Realm_OATH   Realm = "oath"
	Realm_YUBICO Realm = "yubico"

	User_OATH User = "oath"
	User_U2F  User = "u2f"
)

type Realm string
type User string

func PtrRealm(i Realm) *Realm {
	return &i
}
func PtrUser(i User) *User {
	return &i
}

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
	Enabled *util.PVEBool `url:"enabled,omitempty" json:"enabled,omitempty"` // Optional filter for enable property.
	Full    *util.PVEBool `url:"full,omitempty" json:"full,omitempty"`       // Include group and token information.
}

type Tokens struct {
	Tokenid string `url:"tokenid" json:"tokenid"` // User-specific token identifier.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Expire  *int          `url:"expire,omitempty" json:"expire,omitempty"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *util.PVEBool `url:"privsep,omitempty" json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}

type IndexResponse struct {
	Userid string `url:"userid" json:"userid"` // User ID

	// The following parameters are optional
	Comment   *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Email     *string       `url:"email,omitempty" json:"email,omitempty"`
	Enable    *util.PVEBool `url:"enable,omitempty" json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int          `url:"expire,omitempty" json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string       `url:"firstname,omitempty" json:"firstname,omitempty"`
	Groups    *string       `url:"groups,omitempty" json:"groups,omitempty"`
	Keys      *string       `url:"keys,omitempty" json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string       `url:"lastname,omitempty" json:"lastname,omitempty"`
	RealmType *string       `url:"realm-type,omitempty" json:"realm-type,omitempty"` // The type of the users realm
	Tokens    *[]Tokens     `url:"tokens,omitempty" json:"tokens,omitempty"`
}

type CreateRequest struct {
	Userid string `url:"userid" json:"userid"` // User ID

	// The following parameters are optional
	Comment   *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Email     *string       `url:"email,omitempty" json:"email,omitempty"`
	Enable    *util.PVEBool `url:"enable,omitempty" json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int          `url:"expire,omitempty" json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string       `url:"firstname,omitempty" json:"firstname,omitempty"`
	Groups    *string       `url:"groups,omitempty" json:"groups,omitempty"`
	Keys      *string       `url:"keys,omitempty" json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string       `url:"lastname,omitempty" json:"lastname,omitempty"`
	Password  *string       `url:"password,omitempty" json:"password,omitempty"` // Initial password.
}

type FindRequest struct {
	Userid string `url:"userid" json:"userid"` // User ID

}

type FindResponse struct {

	// The following parameters are optional
	Comment   *string                 `url:"comment,omitempty" json:"comment,omitempty"`
	Email     *string                 `url:"email,omitempty" json:"email,omitempty"`
	Enable    *util.PVEBool           `url:"enable,omitempty" json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int                    `url:"expire,omitempty" json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string                 `url:"firstname,omitempty" json:"firstname,omitempty"`
	Groups    *[]string               `url:"groups,omitempty" json:"groups,omitempty"`
	Keys      *string                 `url:"keys,omitempty" json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string                 `url:"lastname,omitempty" json:"lastname,omitempty"`
	Tokens    *map[string]interface{} `url:"tokens,omitempty" json:"tokens,omitempty"`
}

type UpdateRequest struct {
	Userid string `url:"userid" json:"userid"` // User ID

	// The following parameters are optional
	Append    *util.PVEBool `url:"append,omitempty" json:"append,omitempty"`
	Comment   *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Email     *string       `url:"email,omitempty" json:"email,omitempty"`
	Enable    *util.PVEBool `url:"enable,omitempty" json:"enable,omitempty"` // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int          `url:"expire,omitempty" json:"expire,omitempty"` // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string       `url:"firstname,omitempty" json:"firstname,omitempty"`
	Groups    *string       `url:"groups,omitempty" json:"groups,omitempty"`
	Keys      *string       `url:"keys,omitempty" json:"keys,omitempty"` // Keys for two factor auth (yubico).
	Lastname  *string       `url:"lastname,omitempty" json:"lastname,omitempty"`
}

type DeleteRequest struct {
	Userid string `url:"userid" json:"userid"` // User ID

}

type ReadUserTfaTypeTfaRequest struct {
	Userid string `url:"userid" json:"userid"` // User ID

	// The following parameters are optional
	Multiple *util.PVEBool `url:"multiple,omitempty" json:"multiple,omitempty"` // Request all entries as an array.
}

type ReadUserTfaTypeTfaResponse struct {

	// The following parameters are optional
	Realm *Realm    `url:"realm,omitempty" json:"realm,omitempty"` // The type of TFA the users realm has set, if any.
	Types *[]string `url:"types,omitempty" json:"types,omitempty"` // Array of the user configured TFA types, if any. Only available if 'multiple' was not passed.
	User  *User     `url:"user,omitempty" json:"user,omitempty"`   // The type of TFA the user has set, if any. Only set if 'multiple' was not passed.
}

// Index User index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/access/users", "GET", &resp, req)
	return resp, err
}

// Create Create new user.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/access/users", "POST", nil, req)
	return err
}

// Find Get user configuration.
func (c *Client) Find(ctx context.Context, req FindRequest) (FindResponse, error) {
	var resp FindResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}", "GET", &resp, req)
	return resp, err
}

// Update Update user configuration.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/access/users/{userid}", "PUT", nil, req)
	return err
}

// Delete Delete user.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/access/users/{userid}", "DELETE", nil, req)
	return err
}

// ReadUserTfaTypeTfa Get user TFA types (Personal and Realm).
func (c *Client) ReadUserTfaTypeTfa(ctx context.Context, req ReadUserTfaTypeTfaRequest) (ReadUserTfaTypeTfaResponse, error) {
	var resp ReadUserTfaTypeTfaResponse

	err := c.httpClient.Do(ctx, "/access/users/{userid}/tfa", "GET", &resp, req)
	return resp, err
}

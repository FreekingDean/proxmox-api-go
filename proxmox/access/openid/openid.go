// File generated by proxmox json schema, DO NOT EDIT

package openid

import (
	"context"
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

type IndexResponse struct {
	Subdir string `url:"subdir" json:"subdir"`
}

type AuthUrlRequest struct {
	Realm       string `url:"realm" json:"realm"`               // Authentication domain ID
	RedirectUrl string `url:"redirect-url" json:"redirect-url"` // Redirection Url. The client should set this to the used server url (location.origin).

}

type LoginRequest struct {
	Code        string `url:"code" json:"code"`                 // OpenId authorization code.
	RedirectUrl string `url:"redirect-url" json:"redirect-url"` // Redirection Url. The client should set this to the used server url (location.origin).
	State       string `url:"state" json:"state"`               // OpenId state.

}

type LoginResponse struct {
	Cap                 map[string]interface{} `url:"cap" json:"cap"`
	Csrfpreventiontoken string                 `url:"CSRFPreventionToken" json:"CSRFPreventionToken"`
	Ticket              string                 `url:"ticket" json:"ticket"`
	Username            string                 `url:"username" json:"username"`

	// The following parameters are optional
	Clustername *string `url:"clustername,omitempty" json:"clustername,omitempty"`
}

// Index Directory index.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/access/openid", "GET", &resp, nil)
	return resp, err
}

// AuthUrl Get the OpenId Authorization Url for the specified realm.
func (c *Client) AuthUrl(ctx context.Context, req AuthUrlRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/access/openid/auth-url", "POST", &resp, req)
	return resp, err
}

// Login  Verify OpenID authorization code and create a ticket.
func (c *Client) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	var resp LoginResponse

	err := c.httpClient.Do(ctx, "/access/openid/login", "POST", &resp, req)
	return resp, err
}

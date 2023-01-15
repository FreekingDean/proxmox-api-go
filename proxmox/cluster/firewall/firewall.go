// File generated by proxmox json schema, DO NOT EDIT

package firewall

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

type IndexResponse []*map[string]interface{}

// Index Directory index.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall", "GET", &resp, nil)
	return resp, err
}

type GetOptionsResponse struct {

	// The following parameters are optional
	Ebtables     *util.SpecialBool `url:"ebtables,omitempty",json:"ebtables,omitempty"`           // Enable ebtables rules cluster wide.
	Enable       *int              `url:"enable,omitempty",json:"enable,omitempty"`               // Enable or disable the firewall cluster wide.
	LogRatelimit *string           `url:"log_ratelimit,omitempty",json:"log_ratelimit,omitempty"` // Log ratelimiting settings
	PolicyIn     *string           `url:"policy_in,omitempty",json:"policy_in,omitempty"`         // Input policy.
	PolicyOut    *string           `url:"policy_out,omitempty",json:"policy_out,omitempty"`       // Output policy.
}

// GetOptions Get Firewall options.
func (c *Client) GetOptions(ctx context.Context) (*GetOptionsResponse, error) {
	var resp *GetOptionsResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/options", "GET", &resp, nil)
	return resp, err
}

type SetOptionsRequest struct {

	// The following parameters are optional
	Delete       *string           `url:"delete,omitempty",json:"delete,omitempty"`               // A list of settings you want to delete.
	Digest       *string           `url:"digest,omitempty",json:"digest,omitempty"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Ebtables     *util.SpecialBool `url:"ebtables,omitempty",json:"ebtables,omitempty"`           // Enable ebtables rules cluster wide.
	Enable       *int              `url:"enable,omitempty",json:"enable,omitempty"`               // Enable or disable the firewall cluster wide.
	LogRatelimit *string           `url:"log_ratelimit,omitempty",json:"log_ratelimit,omitempty"` // Log ratelimiting settings
	PolicyIn     *string           `url:"policy_in,omitempty",json:"policy_in,omitempty"`         // Input policy.
	PolicyOut    *string           `url:"policy_out,omitempty",json:"policy_out,omitempty"`       // Output policy.
}

type SetOptionsResponse map[string]interface{}

// SetOptions Set Firewall options.
func (c *Client) SetOptions(ctx context.Context, req *SetOptionsRequest) (*SetOptionsResponse, error) {
	var resp *SetOptionsResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/options", "PUT", &resp, req)
	return resp, err
}

type GetMacrosResponse []*struct {
	Descr string `url:"descr",json:"descr"` // More verbose description (if available).
	Macro string `url:"macro",json:"macro"` // Macro name.

}

// GetMacros List available macros
func (c *Client) GetMacros(ctx context.Context) (*GetMacrosResponse, error) {
	var resp *GetMacrosResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/macros", "GET", &resp, nil)
	return resp, err
}

type RefsRequest struct {

	// The following parameters are optional
	Type *string `url:"type,omitempty",json:"type,omitempty"` // Only list references of specified type.
}

type RefsResponse []*struct {
	Name string `url:"name",json:"name"`
	Ref  string `url:"ref",json:"ref"`
	Type string `url:"type",json:"type"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
}

// Refs Lists possible IPSet/Alias reference which are allowed in source/dest properties.
func (c *Client) Refs(ctx context.Context, req *RefsRequest) (*RefsResponse, error) {
	var resp *RefsResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/refs", "GET", &resp, req)
	return resp, err
}

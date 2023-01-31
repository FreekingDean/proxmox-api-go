// File generated by proxmox json schema, DO NOT EDIT

package aliases

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

type IndexRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

}
type _IndexRequest IndexRequest

type IndexResponse struct {
	Cidr   string `url:"cidr" json:"cidr"`
	Digest string `url:"digest" json:"digest"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `url:"name" json:"name"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
}
type _IndexResponse IndexResponse

type CreateRequest struct {
	Cidr string `url:"cidr" json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name" json:"name"` // Alias name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
}
type _CreateRequest CreateRequest

type FindRequest struct {
	Name string `url:"name" json:"name"` // Alias name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

}
type _FindRequest FindRequest

type UpdateRequest struct {
	Cidr string `url:"cidr" json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name" json:"name"` // Alias name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
	Digest  *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Rename  *string `url:"rename,omitempty" json:"rename,omitempty"` // Rename an existing alias.
}
type _UpdateRequest UpdateRequest

type DeleteRequest struct {
	Name string `url:"name" json:"name"` // Alias name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}
type _DeleteRequest DeleteRequest

// Index List aliases
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases", "GET", &resp, req)
	return resp, err
}

// Create Create IP or Network Alias.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases", "POST", nil, req)
	return err
}

// Find Read alias.
func (c *Client) Find(ctx context.Context, req FindRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}", "GET", &resp, req)
	return resp, err
}

// Update Update IP or Network alias.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}", "PUT", nil, req)
	return err
}

// Delete Remove IP or Network alias.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}", "DELETE", nil, req)
	return err
}

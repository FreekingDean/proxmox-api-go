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
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type IndexResponse []*struct {
	Cidr   string `url:"cidr",json:"cidr"`
	Digest string `url:"digest",json:"digest"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `url:"name",json:"name"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
}

// Index List aliases
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Cidr string `url:"cidr",json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name",json:"name"` // Alias name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
}

type CreateResponse map[string]interface{}

// Create Create IP or Network Alias.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Name string `url:"name",json:"name"` // Alias name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type FindResponse map[string]interface{}

// Find Read alias.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Cidr string `url:"cidr",json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name",json:"name"` // Alias name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
	Digest  *string `url:"digest,omitempty",json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Rename  *string `url:"rename,omitempty",json:"rename,omitempty"` // Rename an existing alias.
}

type UpdateResponse map[string]interface{}

// Update Update IP or Network alias.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Name string `url:"name",json:"name"` // Alias name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty",json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

type DeleteResponse map[string]interface{}

// Delete Remove IP or Network alias.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}", "DELETE", &resp, req)
	return resp, err
}

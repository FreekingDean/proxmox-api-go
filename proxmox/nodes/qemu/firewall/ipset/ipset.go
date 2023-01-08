// File generated by proxmox json schema, DO NOT EDIT

package ipset

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
	Digest string `url:"digest",json:"digest"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `url:"name",json:"name"`     // IP set name.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
}

// Index List IPSets
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
	Digest  *string `url:"digest,omitempty",json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Rename  *string `url:"rename,omitempty",json:"rename,omitempty"` // Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
}

type CreateResponse map[string]interface{}

// Create Create new IPSet
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type FindResponse []*struct {
	Cidr   string `url:"cidr",json:"cidr"`
	Digest string `url:"digest",json:"digest"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
	Nomatch *bool   `url:"nomatch,omitempty",json:"nomatch,omitempty"`
}

// Find List IPSet content
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}", "GET", &resp, req)
	return resp, err
}

type ChildCreateRequest struct {
	Cidr string `url:"cidr",json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
	Nomatch *bool   `url:"nomatch,omitempty",json:"nomatch,omitempty"`
}

type ChildCreateResponse map[string]interface{}

// ChildCreate Add IP or Network to IPSet.
func (c *Client) ChildCreate(ctx context.Context, req *ChildCreateRequest) (*ChildCreateResponse, error) {
	var resp *ChildCreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}", "POST", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Force *bool `url:"force,omitempty",json:"force,omitempty"` // Delete all members of the IPSet, if there are any.
}

type DeleteResponse map[string]interface{}

// Delete Delete IPSet
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}", "DELETE", &resp, req)
	return resp, err
}

type ReadIpCidrRequest struct {
	Cidr string `url:"cidr",json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type ReadIpCidrResponse map[string]interface{}

// ReadIpCidr Read IP or Network settings from IPSet.
func (c *Client) ReadIpCidr(ctx context.Context, req *ReadIpCidrRequest) (*ReadIpCidrResponse, error) {
	var resp *ReadIpCidrResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}", "GET", &resp, req)
	return resp, err
}

type UpdateIpCidrRequest struct {
	Cidr string `url:"cidr",json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
	Digest  *string `url:"digest,omitempty",json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Nomatch *bool   `url:"nomatch,omitempty",json:"nomatch,omitempty"`
}

type UpdateIpCidrResponse map[string]interface{}

// UpdateIpCidr Update IP or Network settings
func (c *Client) UpdateIpCidr(ctx context.Context, req *UpdateIpCidrRequest) (*UpdateIpCidrResponse, error) {
	var resp *UpdateIpCidrResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}", "PUT", &resp, req)
	return resp, err
}

type RemoveIpCidrRequest struct {
	Cidr string `url:"cidr",json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name",json:"name"` // IP set name.
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty",json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

type RemoveIpCidrResponse map[string]interface{}

// RemoveIpCidr Remove IP or Network from IPSet.
func (c *Client) RemoveIpCidr(ctx context.Context, req *RemoveIpCidrRequest) (*RemoveIpCidrResponse, error) {
	var resp *RemoveIpCidrResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}", "DELETE", &resp, req)
	return resp, err
}

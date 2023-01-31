// File generated by proxmox json schema, DO NOT EDIT

package ipset

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
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

}
type _IndexRequest IndexRequest

type IndexResponse struct {
	Digest string `url:"digest" json:"digest"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `url:"name" json:"name"`     // IP set name.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
}
type _IndexResponse IndexResponse

type CreateRequest struct {
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
	Digest  *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Rename  *string `url:"rename,omitempty" json:"rename,omitempty"` // Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
}
type _CreateRequest CreateRequest

type FindRequest struct {
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

}
type _FindRequest FindRequest

type FindResponse struct {
	Cidr   string `url:"cidr" json:"cidr"`
	Digest string `url:"digest" json:"digest"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Nomatch *util.PVEBool `url:"nomatch,omitempty" json:"nomatch,omitempty"`
}
type _FindResponse FindResponse

type ChildCreateRequest struct {
	Cidr string `url:"cidr" json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Nomatch *util.PVEBool `url:"nomatch,omitempty" json:"nomatch,omitempty"`
}
type _ChildCreateRequest ChildCreateRequest

type DeleteRequest struct {
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Force *util.PVEBool `url:"force,omitempty" json:"force,omitempty"` // Delete all members of the IPSet, if there are any.
}
type _DeleteRequest DeleteRequest

type ReadIpCidrRequest struct {
	Cidr string `url:"cidr" json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

}
type _ReadIpCidrRequest ReadIpCidrRequest

type UpdateIpCidrRequest struct {
	Cidr string `url:"cidr" json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Digest  *string       `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Nomatch *util.PVEBool `url:"nomatch,omitempty" json:"nomatch,omitempty"`
}
type _UpdateIpCidrRequest UpdateIpCidrRequest

type RemoveIpCidrRequest struct {
	Cidr string `url:"cidr" json:"cidr"` // Network/IP specification in CIDR format.
	Name string `url:"name" json:"name"` // IP set name.
	Node string `url:"node" json:"node"` // The cluster node name.
	Vmid int    `url:"vmid" json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}
type _RemoveIpCidrRequest RemoveIpCidrRequest

// Index List IPSets
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset", "GET", &resp, req)
	return resp, err
}

// Create Create new IPSet
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset", "POST", nil, req)
	return err
}

// Find List IPSet content
func (c *Client) Find(ctx context.Context, req FindRequest) ([]FindResponse, error) {
	var resp []FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}", "GET", &resp, req)
	return resp, err
}

// ChildCreate Add IP or Network to IPSet.
func (c *Client) ChildCreate(ctx context.Context, req ChildCreateRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}", "POST", nil, req)
	return err
}

// Delete Delete IPSet
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}", "DELETE", nil, req)
	return err
}

// ReadIpCidr Read IP or Network settings from IPSet.
func (c *Client) ReadIpCidr(ctx context.Context, req ReadIpCidrRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}", "GET", &resp, req)
	return resp, err
}

// UpdateIpCidr Update IP or Network settings
func (c *Client) UpdateIpCidr(ctx context.Context, req UpdateIpCidrRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}", "PUT", nil, req)
	return err
}

// RemoveIpCidr Remove IP or Network from IPSet.
func (c *Client) RemoveIpCidr(ctx context.Context, req RemoveIpCidrRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}", "DELETE", nil, req)
	return err
}

// File generated by proxmox json schema, DO NOT EDIT

package ipams

import (
	"context"
)

const (
	Type_NETBOX  Type = "netbox"
	Type_PHPIPAM Type = "phpipam"
	Type_PVE     Type = "pve"
)

type Type string

func PtrType(i Type) *Type {
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
	Type *Type `url:"type,omitempty" json:"type,omitempty"` // Only list sdn ipams of specific type
}

type IndexResponse struct {
	Ipam string `url:"ipam" json:"ipam"`
	Type string `url:"type" json:"type"`
}

type CreateRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.
	Type Type   `url:"type" json:"type"` // Plugin type.

	// The following parameters are optional
	Section *int    `url:"section,omitempty" json:"section,omitempty"`
	Token   *string `url:"token,omitempty" json:"token,omitempty"`
	Url     *string `url:"url,omitempty" json:"url,omitempty"`
}

type FindRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.

}

type UpdateRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.

	// The following parameters are optional
	Delete  *string `url:"delete,omitempty" json:"delete,omitempty"` // A list of settings you want to delete.
	Digest  *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Section *int    `url:"section,omitempty" json:"section,omitempty"`
	Token   *string `url:"token,omitempty" json:"token,omitempty"`
	Url     *string `url:"url,omitempty" json:"url,omitempty"`
}

type DeleteRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.

}

// Index SDN ipams index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams", "GET", &resp, req)
	return resp, err
}

// Create Create a new sdn ipam object.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams", "POST", nil, req)
	return err
}

// Find Read sdn ipam configuration.
func (c *Client) Find(ctx context.Context, req FindRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams/{ipam}", "GET", &resp, req)
	return resp, err
}

// Update Update sdn ipam object configuration.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams/{ipam}", "PUT", nil, req)
	return err
}

// Delete Delete sdn ipam object configuration.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams/{ipam}", "DELETE", nil, req)
	return err
}

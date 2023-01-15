// File generated by proxmox json schema, DO NOT EDIT

package ipams

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

	// The following parameters are optional
	Type *string `url:"type,omitempty" json:"type,omitempty"` // Only list sdn ipams of specific type
}

type IndexResponse []*struct {
	Ipam string `url:"ipam" json:"ipam"`
	Type string `url:"type" json:"type"`
}

// Index SDN ipams index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.
	Type string `url:"type" json:"type"` // Plugin type.

	// The following parameters are optional
	Section *int    `url:"section,omitempty" json:"section,omitempty"`
	Token   *string `url:"token,omitempty" json:"token,omitempty"`
	Url     *string `url:"url,omitempty" json:"url,omitempty"`
}

type CreateResponse map[string]interface{}

// Create Create a new sdn ipam object.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.

}

type FindResponse map[string]interface{}

// Find Read sdn ipam configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams/{ipam}", "GET", &resp, req)
	return resp, err
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

type UpdateResponse map[string]interface{}

// Update Update sdn ipam object configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams/{ipam}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Ipam string `url:"ipam" json:"ipam"` // The SDN ipam object identifier.

}

type DeleteResponse map[string]interface{}

// Delete Delete sdn ipam object configuration.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/ipams/{ipam}", "DELETE", &resp, req)
	return resp, err
}

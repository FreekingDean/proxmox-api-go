// File generated by proxmox json schema, DO NOT EDIT

package pools

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

type IndexResponse struct {
	Poolid string `url:"poolid" json:"poolid"`
}

type CreateRequest struct {
	Poolid string `url:"poolid" json:"poolid"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
}

type FindRequest struct {
	Poolid string `url:"poolid" json:"poolid"`

	// The following parameters are optional
	Type *string `url:"type,omitempty" json:"type,omitempty"`
}

type Members struct {
	Id   string `url:"id" json:"id"`
	Node string `url:"node" json:"node"`
	Type string `url:"type" json:"type"`

	// The following parameters are optional
	Storage *string `url:"storage,omitempty" json:"storage,omitempty"`
	Vmid    *int    `url:"vmid,omitempty" json:"vmid,omitempty"`
}

type FindResponse struct {
	Members []Members `url:"members" json:"members"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
}

type UpdateRequest struct {
	Poolid string `url:"poolid" json:"poolid"`

	// The following parameters are optional
	Comment *string       `url:"comment,omitempty" json:"comment,omitempty"`
	Delete  *util.PVEBool `url:"delete,omitempty" json:"delete,omitempty"`   // Remove vms/storage (instead of adding it).
	Storage *string       `url:"storage,omitempty" json:"storage,omitempty"` // List of storage IDs.
	Vms     *string       `url:"vms,omitempty" json:"vms,omitempty"`         // List of virtual machines.
}

type DeleteRequest struct {
	Poolid string `url:"poolid" json:"poolid"`
}

// Index Pool index.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/pools", "GET", &resp, nil)
	return resp, err
}

// Create Create new pool.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/pools", "POST", nil, req)
	return err
}

// Find Get pool configuration.
func (c *Client) Find(ctx context.Context, req FindRequest) (FindResponse, error) {
	var resp FindResponse

	err := c.httpClient.Do(ctx, "/pools/{poolid}", "GET", &resp, req)
	return resp, err
}

// Update Update pool data.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/pools/{poolid}", "PUT", nil, req)
	return err
}

// Delete Delete pool.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/pools/{poolid}", "DELETE", nil, req)
	return err
}

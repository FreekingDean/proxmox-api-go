// File generated by proxmox json schema, DO NOT EDIT

package status

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

type IndexResponse []*map[string]interface{}

// Index Directory index.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/status", "GET", &resp, nil)
	return resp, err
}

type StatusCurrentResponse []*map[string]interface{}

// StatusCurrent Get HA manger status.
func (c *Client) StatusCurrent(ctx context.Context) (*StatusCurrentResponse, error) {
	var resp *StatusCurrentResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/status/current", "GET", &resp, nil)
	return resp, err
}

type ManagerStatusResponse map[string]interface{}

// ManagerStatus Get full HA manger status, including LRM status.
func (c *Client) ManagerStatus(ctx context.Context) (*ManagerStatusResponse, error) {
	var resp *ManagerStatusResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/status/manager_status", "GET", &resp, nil)
	return resp, err
}
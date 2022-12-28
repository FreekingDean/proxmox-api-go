// File generated by proxmox json schema, DO NOT EDIT

package cloudinit

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
	Node string `url:"node",json:"node"`
	Vmid int    `url:"vmid",json:"vmid"`
}

type IndexResponse []*struct {
	New *string `url:"new,omitempty",json:"new,omitempty"`
	Old *string `url:"old,omitempty",json:"old,omitempty"`
	Key string  `url:"key",json:"key"`
}

// Index Get the cloudinit configuration with both current and pending values.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/cloudinit", "GET", &resp, req)
	return resp, err
}

type MassUpdateRequest struct {
	Node string `url:"node",json:"node"`
	Vmid int    `url:"vmid",json:"vmid"`
}

type MassUpdateResponse map[string]interface{}

// MassUpdate Regenerate and change cloudinit config drive.
func (c *Client) MassUpdate(ctx context.Context, req *MassUpdateRequest) (*MassUpdateResponse, error) {
	var resp *MassUpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/cloudinit", "PUT", &resp, req)
	return resp, err
}

type CloudinitGeneratedConfigDumpRequest struct {
	Node string `url:"node",json:"node"`
	Type string `url:"type",json:"type"`
	Vmid int    `url:"vmid",json:"vmid"`
}

type CloudinitGeneratedConfigDumpResponse string

// CloudinitGeneratedConfigDump Get automatically generated cloudinit config.
func (c *Client) CloudinitGeneratedConfigDump(ctx context.Context, req *CloudinitGeneratedConfigDumpRequest) (*CloudinitGeneratedConfigDumpResponse, error) {
	var resp *CloudinitGeneratedConfigDumpResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/cloudinit/dump", "GET", &resp, req)
	return resp, err
}

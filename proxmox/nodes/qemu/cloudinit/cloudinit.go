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
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type IndexResponse []*struct {
	Key string `url:"key",json:"key"` // Configuration option name.

	// The following parameters are optional
	New *string `url:"new,omitempty",json:"new,omitempty"` // The new pending value.
	Old *string `url:"old,omitempty",json:"old,omitempty"` // Value as it was used to generate the current cloudinit image.
}

// Index Get the cloudinit configuration with both current and pending values.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/cloudinit", "GET", &resp, req)
	return resp, err
}

type MassUpdateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type MassUpdateResponse map[string]interface{}

// MassUpdate Regenerate and change cloudinit config drive.
func (c *Client) MassUpdate(ctx context.Context, req *MassUpdateRequest) (*MassUpdateResponse, error) {
	var resp *MassUpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/cloudinit", "PUT", &resp, req)
	return resp, err
}

type CloudinitGeneratedConfigDumpRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Type string `url:"type",json:"type"` // Config type.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type CloudinitGeneratedConfigDumpResponse string

// CloudinitGeneratedConfigDump Get automatically generated cloudinit config.
func (c *Client) CloudinitGeneratedConfigDump(ctx context.Context, req *CloudinitGeneratedConfigDumpRequest) (*CloudinitGeneratedConfigDumpResponse, error) {
	var resp *CloudinitGeneratedConfigDumpResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/cloudinit/dump", "GET", &resp, req)
	return resp, err
}

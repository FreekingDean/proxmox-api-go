// File generated by proxmox json schema, DO NOT EDIT

package disks

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

}

type ListRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	IncludePartitions *util.PVEBool `url:"include-partitions,omitempty" json:"include-partitions,omitempty"` // Also include partitions.
	Skipsmart         *util.PVEBool `url:"skipsmart,omitempty" json:"skipsmart,omitempty"`                   // Skip smart checks.
	Type              *string       `url:"type,omitempty" json:"type,omitempty"`                             // Only list specific types of disks.
}

type ListResponse struct {
	Devpath string       `url:"devpath" json:"devpath"` // The device path
	Gpt     util.PVEBool `url:"gpt" json:"gpt"`
	Mounted util.PVEBool `url:"mounted" json:"mounted"`
	Osdid   int          `url:"osdid" json:"osdid"`
	Size    int          `url:"size" json:"size"`

	// The following parameters are optional
	Health *string `url:"health,omitempty" json:"health,omitempty"`
	Model  *string `url:"model,omitempty" json:"model,omitempty"`
	Parent *string `url:"parent,omitempty" json:"parent,omitempty"` // For partitions only. The device path of the disk the partition resides on.
	Serial *string `url:"serial,omitempty" json:"serial,omitempty"`
	Used   *string `url:"used,omitempty" json:"used,omitempty"`
	Vendor *string `url:"vendor,omitempty" json:"vendor,omitempty"`
	Wwn    *string `url:"wwn,omitempty" json:"wwn,omitempty"`
}

type SmartRequest struct {
	Disk string `url:"disk" json:"disk"` // Block device name
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Healthonly *util.PVEBool `url:"healthonly,omitempty" json:"healthonly,omitempty"` // If true returns only the health status
}

type SmartResponse struct {
	Health string `url:"health" json:"health"`

	// The following parameters are optional
	Attributes *[]map[string]interface{} `url:"attributes,omitempty" json:"attributes,omitempty"`
	Text       *string                   `url:"text,omitempty" json:"text,omitempty"`
	Type       *string                   `url:"type,omitempty" json:"type,omitempty"`
}

type InitgptRequest struct {
	Disk string `url:"disk" json:"disk"` // Block device name
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Uuid *string `url:"uuid,omitempty" json:"uuid,omitempty"` // UUID for the GPT table
}

type WipeDiskWipediskRequest struct {
	Disk string `url:"disk" json:"disk"` // Block device name
	Node string `url:"node" json:"node"` // The cluster node name.

}

// Index Node index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks", "GET", &resp, req)
	return resp, err
}

// List List local disks.
func (c *Client) List(ctx context.Context, req ListRequest) ([]ListResponse, error) {
	var resp []ListResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/list", "GET", &resp, req)
	return resp, err
}

// Smart Get SMART Health of a disk.
func (c *Client) Smart(ctx context.Context, req SmartRequest) (SmartResponse, error) {
	var resp SmartResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/smart", "GET", &resp, req)
	return resp, err
}

// Initgpt Initialize Disk with GPT
func (c *Client) Initgpt(ctx context.Context, req InitgptRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/initgpt", "POST", &resp, req)
	return resp, err
}

// WipeDiskWipedisk Wipe a disk or partition.
func (c *Client) WipeDiskWipedisk(ctx context.Context, req WipeDiskWipediskRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/wipedisk", "PUT", &resp, req)
	return resp, err
}

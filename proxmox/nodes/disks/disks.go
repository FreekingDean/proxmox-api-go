// File generated by proxmox json schema, DO NOT EDIT

package disks

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
}

type IndexResponse []*map[string]interface{}

// Index Node index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks", "GET", &resp, req)
	return resp, err
}

type ListRequest struct {
	Type              *string `url:"type,omitempty",json:"type,omitempty"`
	IncludePartitions *bool   `url:"include-partitions,omitempty",json:"include-partitions,omitempty"`
	Node              string  `url:"node",json:"node"`
	Skipsmart         *bool   `url:"skipsmart,omitempty",json:"skipsmart,omitempty"`
}

type ListResponse []*struct {
	Vendor  *string `url:"vendor,omitempty",json:"vendor,omitempty"`
	Gpt     bool    `url:"gpt",json:"gpt"`
	Health  *string `url:"health,omitempty",json:"health,omitempty"`
	Mounted bool    `url:"mounted",json:"mounted"`
	Size    int     `url:"size",json:"size"`
	Serial  *string `url:"serial,omitempty",json:"serial,omitempty"`
	Used    *string `url:"used,omitempty",json:"used,omitempty"`
	Wwn     *string `url:"wwn,omitempty",json:"wwn,omitempty"`
	Devpath string  `url:"devpath",json:"devpath"`
	Model   *string `url:"model,omitempty",json:"model,omitempty"`
	Osdid   int     `url:"osdid",json:"osdid"`
	Parent  *string `url:"parent,omitempty",json:"parent,omitempty"`
}

// List List local disks.
func (c *Client) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	var resp *ListResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/list", "GET", &resp, req)
	return resp, err
}

type SmartRequest struct {
	Healthonly *bool  `url:"healthonly,omitempty",json:"healthonly,omitempty"`
	Node       string `url:"node",json:"node"`
	Disk       string `url:"disk",json:"disk"`
}

type SmartResponse struct {
	Attributes []*map[string]interface{} `url:"attributes,omitempty",json:"attributes,omitempty"`
	Health     string                    `url:"health",json:"health"`
	Text       *string                   `url:"text,omitempty",json:"text,omitempty"`
	Type       *string                   `url:"type,omitempty",json:"type,omitempty"`
}

// Smart Get SMART Health of a disk.
func (c *Client) Smart(ctx context.Context, req *SmartRequest) (*SmartResponse, error) {
	var resp *SmartResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/smart", "GET", &resp, req)
	return resp, err
}

type InitgptRequest struct {
	Disk string  `url:"disk",json:"disk"`
	Node string  `url:"node",json:"node"`
	Uuid *string `url:"uuid,omitempty",json:"uuid,omitempty"`
}

type InitgptResponse string

// Initgpt Initialize Disk with GPT
func (c *Client) Initgpt(ctx context.Context, req *InitgptRequest) (*InitgptResponse, error) {
	var resp *InitgptResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/initgpt", "POST", &resp, req)
	return resp, err
}

type WipeDiskWipediskRequest struct {
	Disk string `url:"disk",json:"disk"`
	Node string `url:"node",json:"node"`
}

type WipeDiskWipediskResponse string

// WipeDiskWipedisk Wipe a disk or partition.
func (c *Client) WipeDiskWipedisk(ctx context.Context, req *WipeDiskWipediskRequest) (*WipeDiskWipediskResponse, error) {
	var resp *WipeDiskWipediskResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/wipedisk", "PUT", &resp, req)
	return resp, err
}

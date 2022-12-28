// File generated by proxmox json schema, DO NOT EDIT

package qemu

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

// Index QEMU capabilities index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/capabilities/qemu", "GET", &resp, req)
	return resp, err
}

type IndexCpuRequest struct {
	Node string `url:"node",json:"node"`
}

type IndexCpuResponse []*struct {
	Custom bool   `url:"custom",json:"custom"`
	Name   string `url:"name",json:"name"`
	Vendor string `url:"vendor",json:"vendor"`
}

// IndexCpu List all custom and default CPU models.
func (c *Client) IndexCpu(ctx context.Context, req *IndexCpuRequest) (*IndexCpuResponse, error) {
	var resp *IndexCpuResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/capabilities/qemu/cpu", "GET", &resp, req)
	return resp, err
}

type TypesMachinesRequest struct {
	Node string `url:"node",json:"node"`
}

type TypesMachinesResponse []*struct {
	Id      string `url:"id",json:"id"`
	Type    string `url:"type",json:"type"`
	Version string `url:"version",json:"version"`
}

// TypesMachines Get available QEMU/KVM machine types.
func (c *Client) TypesMachines(ctx context.Context, req *TypesMachinesRequest) (*TypesMachinesResponse, error) {
	var resp *TypesMachinesResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/capabilities/qemu/machines", "GET", &resp, req)
	return resp, err
}

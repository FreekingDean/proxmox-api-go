// File generated by proxmox json schema, DO NOT EDIT

package qemu

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

type IndexCpuRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}

type IndexCpuResponse struct {
	Custom util.SpecialBool `url:"custom" json:"custom"` // True if this is a custom CPU model.
	Name   string           `url:"name" json:"name"`     // Name of the CPU model. Identifies it for subsequent API calls. Prefixed with 'custom-' for custom models.
	Vendor string           `url:"vendor" json:"vendor"` // CPU vendor visible to the guest when this model is selected. Vendor of 'reported-model' in case of custom models.

}

type TypesMachinesRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}

type TypesMachinesResponse struct {
	Id      string `url:"id" json:"id"`           // Full name of machine type and version.
	Type    string `url:"type" json:"type"`       // The machine type.
	Version string `url:"version" json:"version"` // The machine version.

}

// Index QEMU capabilities index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/capabilities/qemu", "GET", &resp, req)
	return resp, err
}

// IndexCpu List all custom and default CPU models.
func (c *Client) IndexCpu(ctx context.Context, req IndexCpuRequest) ([]IndexCpuResponse, error) {
	var resp []IndexCpuResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/capabilities/qemu/cpu", "GET", &resp, req)
	return resp, err
}

// TypesMachines Get available QEMU/KVM machine types.
func (c *Client) TypesMachines(ctx context.Context, req TypesMachinesRequest) ([]TypesMachinesResponse, error) {
	var resp []TypesMachinesResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/capabilities/qemu/machines", "GET", &resp, req)
	return resp, err
}

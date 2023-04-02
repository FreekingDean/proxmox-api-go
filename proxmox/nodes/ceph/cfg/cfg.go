// File generated by proxmox json schema, DO NOT EDIT

package cfg

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
type _IndexRequest IndexRequest

type RawRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _RawRequest RawRequest

type DbRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _DbRequest DbRequest

type DbResponse struct {
	CanUpdateAtRuntime util.PVEBool `url:"can_update_at_runtime" json:"can_update_at_runtime"`
	Level              string       `url:"level" json:"level"`
	Mask               string       `url:"mask" json:"mask"`
	Name               string       `url:"name" json:"name"`
	Section            string       `url:"section" json:"section"`
	Value              string       `url:"value" json:"value"`
}
type _DbResponse DbResponse

// Index Directory index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/cfg", "GET", &resp, req)
	return resp, err
}

// Raw Get the Ceph configuration file.
func (c *Client) Raw(ctx context.Context, req RawRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/cfg/raw", "GET", &resp, req)
	return resp, err
}

// Db Get the Ceph configuration database.
func (c *Client) Db(ctx context.Context, req DbRequest) ([]DbResponse, error) {
	var resp []DbResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/cfg/db", "GET", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package mgr

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
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _IndexRequest IndexRequest

type IndexResponse struct {
	State string `url:"state" json:"state"` // State of the MGR

	// The following parameters are optional
	Addr *string `url:"addr,omitempty" json:"addr,omitempty"`
	Host *string `url:"host,omitempty" json:"host,omitempty"`
}
type _IndexResponse IndexResponse

type ChildCreateRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Id *string `url:"id,omitempty" json:"id,omitempty"` // The ID for the manager, when omitted the same as the nodename
}
type _ChildCreateRequest ChildCreateRequest

type DeleteRequest struct {
	Id   string `url:"id" json:"id"`     // The ID of the manager
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _DeleteRequest DeleteRequest

// Index MGR directory index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mgr", "GET", &resp, req)
	return resp, err
}

// ChildCreate Create Ceph Manager
func (c *Client) ChildCreate(ctx context.Context, req ChildCreateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mgr/{id}", "POST", &resp, req)
	return resp, err
}

// Delete Destroy Ceph Manager.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mgr/{id}", "DELETE", &resp, req)
	return resp, err
}

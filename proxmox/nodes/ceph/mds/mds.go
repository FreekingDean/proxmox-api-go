// File generated by proxmox json schema, DO NOT EDIT

package mds

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

type IndexResponse struct {
	State string `url:"state" json:"state"` // State of the MDS

	// The following parameters are optional
	Addr          *string       `url:"addr,omitempty" json:"addr,omitempty"`
	Host          *string       `url:"host,omitempty" json:"host,omitempty"`
	Rank          *int          `url:"rank,omitempty" json:"rank,omitempty"`
	StandbyReplay *util.PVEBool `url:"standby_replay,omitempty" json:"standby_replay,omitempty"` // If true, the standby MDS is polling the active MDS for faster recovery (hot standby).
}
type _IndexResponse IndexResponse

type ChildCreateRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Hotstandby *util.PVEBool `url:"hotstandby,omitempty" json:"hotstandby,omitempty"` // Determines whether a ceph-mds daemon should poll and replay the log of an active MDS. Faster switch on MDS failure, but needs more idle resources.
	Name       *string       `url:"name,omitempty" json:"name,omitempty"`             // The ID for the mds, when omitted the same as the nodename
}
type _ChildCreateRequest ChildCreateRequest

type DeleteRequest struct {
	Name string `url:"name" json:"name"` // The name (ID) of the mds
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _DeleteRequest DeleteRequest

// Index MDS directory index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mds", "GET", &resp, req)
	return resp, err
}

// ChildCreate Create Ceph Metadata Server (MDS)
func (c *Client) ChildCreate(ctx context.Context, req ChildCreateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mds/{name}", "POST", &resp, req)
	return resp, err
}

// Delete Destroy Ceph Metadata Server
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mds/{name}", "DELETE", &resp, req)
	return resp, err
}

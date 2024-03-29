// File generated by proxmox json schema, DO NOT EDIT

package mon

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
	Name string `url:"name" json:"name"`

	// The following parameters are optional
	Addr             *string       `url:"addr,omitempty" json:"addr,omitempty"`
	CephVersion      *string       `url:"ceph_version,omitempty" json:"ceph_version,omitempty"`
	CephVersionShort *string       `url:"ceph_version_short,omitempty" json:"ceph_version_short,omitempty"`
	Direxists        *string       `url:"direxists,omitempty" json:"direxists,omitempty"`
	Host             *util.PVEBool `url:"host,omitempty" json:"host,omitempty"`
	Quorum           *util.PVEBool `url:"quorum,omitempty" json:"quorum,omitempty"`
	Rank             *int          `url:"rank,omitempty" json:"rank,omitempty"`
	Service          *int          `url:"service,omitempty" json:"service,omitempty"`
	State            *string       `url:"state,omitempty" json:"state,omitempty"`
}
type _IndexResponse IndexResponse

type ChildCreateRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	MonAddress *string `url:"mon-address,omitempty" json:"mon-address,omitempty"` // Overwrites autodetected monitor IP address(es). Must be in the public network(s) of Ceph.
	Monid      *string `url:"monid,omitempty" json:"monid,omitempty"`             // The ID for the monitor, when omitted the same as the nodename
}
type _ChildCreateRequest ChildCreateRequest

type DeleteRequest struct {
	Monid string `url:"monid" json:"monid"` // Monitor ID
	Node  string `url:"node" json:"node"`   // The cluster node name.

}
type _DeleteRequest DeleteRequest

// Index Get Ceph monitor list.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mon", "GET", &resp, req)
	return resp, err
}

// ChildCreate Create Ceph Monitor and Manager
func (c *Client) ChildCreate(ctx context.Context, req ChildCreateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mon/{monid}", "POST", &resp, req)
	return resp, err
}

// Delete Destroy Ceph Monitor and Manager.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/mon/{monid}", "DELETE", &resp, req)
	return resp, err
}

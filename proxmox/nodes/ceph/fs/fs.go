// File generated by proxmox json schema, DO NOT EDIT

package fs

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
	Node string `url:"node",json:"node"` // The cluster node name.

}

type IndexResponse []*struct {
	DataPool     string `url:"data_pool",json:"data_pool"`         // The name of the data pool.
	MetadataPool string `url:"metadata_pool",json:"metadata_pool"` // The name of the metadata pool.
	Name         string `url:"name",json:"name"`                   // The ceph filesystem name.

}

// Index Directory index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/fs", "GET", &resp, req)
	return resp, err
}

type ChildCreateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	AddStorage *util.SpecialBool `url:"add-storage,omitempty",json:"add-storage,omitempty"` // Configure the created CephFS as storage for this cluster.
	Name       *string           `url:"name,omitempty",json:"name,omitempty"`               // The ceph filesystem name.
	PgNum      *int              `url:"pg_num,omitempty",json:"pg_num,omitempty"`           // Number of placement groups for the backing data pool. The metadata pool will use a quarter of this.
}

type ChildCreateResponse string

// ChildCreate Create a Ceph filesystem
func (c *Client) ChildCreate(ctx context.Context, req *ChildCreateRequest) (*ChildCreateResponse, error) {
	var resp *ChildCreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/fs/{name}", "POST", &resp, req)
	return resp, err
}

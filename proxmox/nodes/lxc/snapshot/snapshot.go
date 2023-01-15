// File generated by proxmox json schema, DO NOT EDIT

package snapshot

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
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type IndexResponse []*struct {
	Description string `url:"description",json:"description"` // Snapshot description.
	Name        string `url:"name",json:"name"`               // Snapshot identifier. Value 'current' identifies the current VM.

	// The following parameters are optional
	Parent   *string `url:"parent,omitempty",json:"parent,omitempty"`     // Parent snapshot identifier.
	Snaptime *int    `url:"snaptime,omitempty",json:"snaptime,omitempty"` // Snapshot creation time
}

// Index List all snapshots.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Node     string `url:"node",json:"node"`         // The cluster node name.
	Snapname string `url:"snapname",json:"snapname"` // The name of the snapshot.
	Vmid     int    `url:"vmid",json:"vmid"`         // The (unique) ID of the VM.

	// The following parameters are optional
	Description *string `url:"description,omitempty",json:"description,omitempty"` // A textual description or comment.
}

type CreateResponse string

// Create Snapshot a container.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Node     string `url:"node",json:"node"`         // The cluster node name.
	Snapname string `url:"snapname",json:"snapname"` // The name of the snapshot.
	Vmid     int    `url:"vmid",json:"vmid"`         // The (unique) ID of the VM.

}

type FindResponse []*map[string]interface{}

// Find
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot/{snapname}", "GET", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Node     string `url:"node",json:"node"`         // The cluster node name.
	Snapname string `url:"snapname",json:"snapname"` // The name of the snapshot.
	Vmid     int    `url:"vmid",json:"vmid"`         // The (unique) ID of the VM.

	// The following parameters are optional
	Force *util.SpecialBool `url:"force,omitempty",json:"force,omitempty"` // For removal from config file, even if removing disk snapshots fails.
}

type DeleteResponse string

// Delete Delete a LXC snapshot.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot/{snapname}", "DELETE", &resp, req)
	return resp, err
}

type RollbackRequest struct {
	Node     string `url:"node",json:"node"`         // The cluster node name.
	Snapname string `url:"snapname",json:"snapname"` // The name of the snapshot.
	Vmid     int    `url:"vmid",json:"vmid"`         // The (unique) ID of the VM.

	// The following parameters are optional
	Start *util.SpecialBool `url:"start,omitempty",json:"start,omitempty"` // Whether the container should get started after rolling back successfully
}

type RollbackResponse string

// Rollback Rollback LXC state to specified snapshot.
func (c *Client) Rollback(ctx context.Context, req *RollbackRequest) (*RollbackResponse, error) {
	var resp *RollbackResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot/{snapname}/rollback", "POST", &resp, req)
	return resp, err
}

type GetSnapshotConfigRequest struct {
	Node     string `url:"node",json:"node"`         // The cluster node name.
	Snapname string `url:"snapname",json:"snapname"` // The name of the snapshot.
	Vmid     int    `url:"vmid",json:"vmid"`         // The (unique) ID of the VM.

}

type GetSnapshotConfigResponse map[string]interface{}

// GetSnapshotConfig Get snapshot configuration
func (c *Client) GetSnapshotConfig(ctx context.Context, req *GetSnapshotConfigRequest) (*GetSnapshotConfigResponse, error) {
	var resp *GetSnapshotConfigResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config", "GET", &resp, req)
	return resp, err
}

type UpdateSnapshotConfigRequest struct {
	Node     string `url:"node",json:"node"`         // The cluster node name.
	Snapname string `url:"snapname",json:"snapname"` // The name of the snapshot.
	Vmid     int    `url:"vmid",json:"vmid"`         // The (unique) ID of the VM.

	// The following parameters are optional
	Description *string `url:"description,omitempty",json:"description,omitempty"` // A textual description or comment.
}

type UpdateSnapshotConfigResponse map[string]interface{}

// UpdateSnapshotConfig Update snapshot metadata.
func (c *Client) UpdateSnapshotConfig(ctx context.Context, req *UpdateSnapshotConfigRequest) (*UpdateSnapshotConfigResponse, error) {
	var resp *UpdateSnapshotConfigResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config", "PUT", &resp, req)
	return resp, err
}

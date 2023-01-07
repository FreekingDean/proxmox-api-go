// File generated by proxmox json schema, DO NOT EDIT

package resources

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

type IndexRequest map[string]interface{}

type IndexResponse []*struct {
	Sid string `url:"sid",json:"sid"`
}

// Index List HA resources.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Sid string `url:"sid",json:"sid"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).

	// The following parameters are optional
	Comment     *string `url:"comment,omitempty",json:"comment,omitempty"`           // Description.
	Group       *string `url:"group,omitempty",json:"group,omitempty"`               // The HA group identifier.
	MaxRelocate *int    `url:"max_relocate,omitempty",json:"max_relocate,omitempty"` // Maximal number of service relocate tries when a service failes to start.
	MaxRestart  *int    `url:"max_restart,omitempty",json:"max_restart,omitempty"`   // Maximal number of tries to restart the service on a node after its start failed.
	State       *string `url:"state,omitempty",json:"state,omitempty"`               // Requested resource state.
	Type        *string `url:"type,omitempty",json:"type,omitempty"`                 // Resource type.
}

type CreateResponse map[string]interface{}

// Create Create a new HA resource.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Sid string `url:"sid",json:"sid"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).

}

type FindResponse struct {
	Digest string `url:"digest",json:"digest"` // Can be used to prevent concurrent modifications.
	Sid    string `url:"sid",json:"sid"`       // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
	Type   string `url:"type",json:"type"`     // The type of the resources.

	// The following parameters are optional
	Comment     *string `url:"comment,omitempty",json:"comment,omitempty"`           // Description.
	Group       *string `url:"group,omitempty",json:"group,omitempty"`               // The HA group identifier.
	MaxRelocate *int    `url:"max_relocate,omitempty",json:"max_relocate,omitempty"` // Maximal number of service relocate tries when a service failes to start.
	MaxRestart  *int    `url:"max_restart,omitempty",json:"max_restart,omitempty"`   // Maximal number of tries to restart the service on a node after its start failed.
	State       *string `url:"state,omitempty",json:"state,omitempty"`               // Requested resource state.
}

// Find Read resource configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources/{sid}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Sid string `url:"sid",json:"sid"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).

	// The following parameters are optional
	Comment     *string `url:"comment,omitempty",json:"comment,omitempty"`           // Description.
	Delete      *string `url:"delete,omitempty",json:"delete,omitempty"`             // A list of settings you want to delete.
	Digest      *string `url:"digest,omitempty",json:"digest,omitempty"`             // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group       *string `url:"group,omitempty",json:"group,omitempty"`               // The HA group identifier.
	MaxRelocate *int    `url:"max_relocate,omitempty",json:"max_relocate,omitempty"` // Maximal number of service relocate tries when a service failes to start.
	MaxRestart  *int    `url:"max_restart,omitempty",json:"max_restart,omitempty"`   // Maximal number of tries to restart the service on a node after its start failed.
	State       *string `url:"state,omitempty",json:"state,omitempty"`               // Requested resource state.
}

type UpdateResponse map[string]interface{}

// Update Update resource configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources/{sid}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Sid string `url:"sid",json:"sid"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).

}

type DeleteResponse map[string]interface{}

// Delete Delete resource configuration.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources/{sid}", "DELETE", &resp, req)
	return resp, err
}

type MigrateRequest struct {
	Node string `url:"node",json:"node"` // Target node.
	Sid  string `url:"sid",json:"sid"`   // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).

}

type MigrateResponse map[string]interface{}

// Migrate Request resource migration (online) to another node.
func (c *Client) Migrate(ctx context.Context, req *MigrateRequest) (*MigrateResponse, error) {
	var resp *MigrateResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources/{sid}/migrate", "POST", &resp, req)
	return resp, err
}

type RelocateRequest struct {
	Node string `url:"node",json:"node"` // Target node.
	Sid  string `url:"sid",json:"sid"`   // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).

}

type RelocateResponse map[string]interface{}

// Relocate Request resource relocatzion to another node. This stops the service on the old node, and restarts it on the target node.
func (c *Client) Relocate(ctx context.Context, req *RelocateRequest) (*RelocateResponse, error) {
	var resp *RelocateResponse

	err := c.httpClient.Do(ctx, "/cluster/ha/resources/{sid}/relocate", "POST", &resp, req)
	return resp, err
}
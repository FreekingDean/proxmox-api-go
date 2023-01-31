// File generated by proxmox json schema, DO NOT EDIT

package directory

import (
	"context"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
)

const (
	Filesystem_EXT4 Filesystem = "ext4"
	Filesystem_XFS  Filesystem = "xfs"
)

type Filesystem string

func PtrFilesystem(i Filesystem) *Filesystem {
	return &i
}

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
	Device   string `url:"device" json:"device"`     // The mounted device.
	Options  string `url:"options" json:"options"`   // The mount options.
	Path     string `url:"path" json:"path"`         // The mount path.
	Type     string `url:"type" json:"type"`         // The filesystem type.
	Unitfile string `url:"unitfile" json:"unitfile"` // The path of the mount unit.

}
type _IndexResponse IndexResponse

type CreateRequest struct {
	Device string `url:"device" json:"device"` // The block device you want to create the filesystem on.
	Name   string `url:"name" json:"name"`     // The storage identifier.
	Node   string `url:"node" json:"node"`     // The cluster node name.

	// The following parameters are optional
	AddStorage *util.PVEBool `url:"add_storage,omitempty" json:"add_storage,omitempty"` // Configure storage using the directory.
	Filesystem *Filesystem   `url:"filesystem,omitempty" json:"filesystem,omitempty"`   // The desired filesystem.
}
type _CreateRequest CreateRequest

type DeleteRequest struct {
	Name string `url:"name" json:"name"` // The storage identifier.
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	CleanupConfig *util.PVEBool `url:"cleanup-config,omitempty" json:"cleanup-config,omitempty"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  *util.PVEBool `url:"cleanup-disks,omitempty" json:"cleanup-disks,omitempty"`   // Also wipe disk so it can be repurposed afterwards.
}
type _DeleteRequest DeleteRequest

// Index PVE Managed Directory storages.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/directory", "GET", &resp, req)
	return resp, err
}

// Create Create a Filesystem on an unused disk. Will be mounted under '/mnt/pve/NAME'.
func (c *Client) Create(ctx context.Context, req CreateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/directory", "POST", &resp, req)
	return resp, err
}

// Delete Unmounts the storage and removes the mount unit.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/directory/{name}", "DELETE", &resp, req)
	return resp, err
}

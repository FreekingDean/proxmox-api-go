// File generated by proxmox json schema, DO NOT EDIT

package zfs

import (
	"context"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
	"net/url"
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

type IndexResponse struct {
	Alloc  int     `url:"alloc" json:"alloc"`
	Dedup  float64 `url:"dedup" json:"dedup"`
	Frag   int     `url:"frag" json:"frag"`
	Free   int     `url:"free" json:"free"`
	Health string  `url:"health" json:"health"`
	Name   string  `url:"name" json:"name"`
	Size   int     `url:"size" json:"size"`
}

// Array of DraidConfig
type DraidConfigArr []DraidConfig

func (t DraidConfigArr) EncodeValues(key string, v *url.Values) error {
	return util.EncodeArray(key, v, t)
}

type DraidConfig struct {
	Data   int `url:"data" json:"data"`     // The number of data devices per redundancy group. (dRAID)
	Spares int `url:"spares" json:"spares"` // Number of dRAID spares.

}

func (t DraidConfig) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `data=<integer> ,spares=<integer>`)
}

type CreateRequest struct {
	Devices   string `url:"devices" json:"devices"`     // The block devices you want to create the zpool on.
	Name      string `url:"name" json:"name"`           // The storage identifier.
	Node      string `url:"node" json:"node"`           // The cluster node name.
	Raidlevel string `url:"raidlevel" json:"raidlevel"` // The RAID level to use.

	// The following parameters are optional
	AddStorage  *util.SpecialBool `url:"add_storage,omitempty" json:"add_storage,omitempty"` // Configure storage using the zpool.
	Ashift      *int              `url:"ashift,omitempty" json:"ashift,omitempty"`           // Pool sector size exponent.
	Compression *string           `url:"compression,omitempty" json:"compression,omitempty"` // The compression algorithm to use.
	DraidConfig *DraidConfig      `url:"draid-config,omitempty" json:"draid-config,omitempty"`
}

type FindRequest struct {
	Name string `url:"name" json:"name"` // The storage identifier.
	Node string `url:"node" json:"node"` // The cluster node name.

}

type Children struct {
	Msg  string `url:"msg" json:"msg"`   // An optional message about the vdev.
	Name string `url:"name" json:"name"` // The name of the vdev or section.

	// The following parameters are optional
	Cksum *float64 `url:"cksum,omitempty" json:"cksum,omitempty"`
	Read  *float64 `url:"read,omitempty" json:"read,omitempty"`
	State *string  `url:"state,omitempty" json:"state,omitempty"` // The state of the vdev.
	Write *float64 `url:"write,omitempty" json:"write,omitempty"`
}

type FindResponse struct {
	Children []Children `url:"children" json:"children"` // The pool configuration information, including the vdevs for each section (e.g. spares, cache), may be nested.
	Errors   string     `url:"errors" json:"errors"`     // Information about the errors on the zpool.
	Name     string     `url:"name" json:"name"`         // The name of the zpool.
	State    string     `url:"state" json:"state"`       // The state of the zpool.

	// The following parameters are optional
	Action *string `url:"action,omitempty" json:"action,omitempty"` // Information about the recommended action to fix the state.
	Scan   *string `url:"scan,omitempty" json:"scan,omitempty"`     // Information about the last/current scrub.
	Status *string `url:"status,omitempty" json:"status,omitempty"` // Information about the state of the zpool.
}

type DeleteRequest struct {
	Name string `url:"name" json:"name"` // The storage identifier.
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	CleanupConfig *util.SpecialBool `url:"cleanup-config,omitempty" json:"cleanup-config,omitempty"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  *util.SpecialBool `url:"cleanup-disks,omitempty" json:"cleanup-disks,omitempty"`   // Also wipe disks so they can be repurposed afterwards.
}

// Index List Zpools.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/zfs", "GET", &resp, req)
	return resp, err
}

// Create Create a ZFS pool.
func (c *Client) Create(ctx context.Context, req CreateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/zfs", "POST", &resp, req)
	return resp, err
}

// Find Get details about a zpool.
func (c *Client) Find(ctx context.Context, req FindRequest) (FindResponse, error) {
	var resp FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/zfs/{name}", "GET", &resp, req)
	return resp, err
}

// Delete Destroy a ZFS pool.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/disks/zfs/{name}", "DELETE", &resp, req)
	return resp, err
}

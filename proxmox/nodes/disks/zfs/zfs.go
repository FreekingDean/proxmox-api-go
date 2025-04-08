// File generated by proxmox json schema, DO NOT EDIT

package zfs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
	"net/url"
	"strings"
)

const (
	Compression_ON   Compression = "on"
	Compression_OFF  Compression = "off"
	Compression_GZIP Compression = "gzip"
	Compression_LZ4  Compression = "lz4"
	Compression_LZJB Compression = "lzjb"
	Compression_ZLE  Compression = "zle"
	Compression_ZSTD Compression = "zstd"

	Raidlevel_SINGLE Raidlevel = "single"
	Raidlevel_MIRROR Raidlevel = "mirror"
	Raidlevel_RAID10 Raidlevel = "raid10"
	Raidlevel_RAIDZ  Raidlevel = "raidz"
	Raidlevel_RAIDZ2 Raidlevel = "raidz2"
	Raidlevel_RAIDZ3 Raidlevel = "raidz3"
	Raidlevel_DRAID  Raidlevel = "draid"
	Raidlevel_DRAID2 Raidlevel = "draid2"
	Raidlevel_DRAID3 Raidlevel = "draid3"
)

type Compression string
type Raidlevel string

func PtrCompression(i Compression) *Compression {
	return &i
}
func PtrRaidlevel(i Raidlevel) *Raidlevel {
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
	Alloc  int     `url:"alloc" json:"alloc"`
	Dedup  float64 `url:"dedup" json:"dedup"`
	Frag   int     `url:"frag" json:"frag"`
	Free   int     `url:"free" json:"free"`
	Health string  `url:"health" json:"health"`
	Name   string  `url:"name" json:"name"`
	Size   int     `url:"size" json:"size"`
}
type _IndexResponse IndexResponse

type DraidConfig struct {
	Data   int `url:"data" json:"data"`     // The number of data devices per redundancy group. (dRAID)
	Spares int `url:"spares" json:"spares"` // Number of dRAID spares.

}
type _DraidConfig DraidConfig

func (t DraidConfig) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `data=<integer> ,spares=<integer>`)
}

func (t *DraidConfig) UnmarshalJSON(d []byte) error {
	if len(d) == 0 || string(d) == `""` {
		return nil
	}
	cleaned := string(d)[1 : len(d)-1]
	parts := strings.Split(cleaned, ",")
	values := map[string]string{}
	for _, p := range parts {
		kv := strings.Split(p, "=")
		if len(kv) > 2 {
			return fmt.Errorf("Wrong number of parts for kv pair '%s'", p)
		}
		if len(kv) == 1 {
			values[""] = kv[0]
			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["data"]; ok {

		err := json.Unmarshal([]byte(v), &t.Data)
		if err != nil {
			return err
		}
	}

	if v, ok := values["spares"]; ok {

		err := json.Unmarshal([]byte(v), &t.Spares)
		if err != nil {
			return err
		}
	}

	return nil
}

type CreateRequest struct {
	Devices   string    `url:"devices" json:"devices"`     // The block devices you want to create the zpool on.
	Name      string    `url:"name" json:"name"`           // The storage identifier.
	Node      string    `url:"node" json:"node"`           // The cluster node name.
	Raidlevel Raidlevel `url:"raidlevel" json:"raidlevel"` // The RAID level to use.

	// The following parameters are optional
	AddStorage  *util.PVEBool `url:"add_storage,omitempty" json:"add_storage,omitempty"` // Configure storage using the zpool.
	Ashift      *int          `url:"ashift,omitempty" json:"ashift,omitempty"`           // Pool sector size exponent.
	Compression *Compression  `url:"compression,omitempty" json:"compression,omitempty"` // The compression algorithm to use.
	DraidConfig *DraidConfig  `url:"draid-config,omitempty" json:"draid-config,omitempty"`
}
type _CreateRequest CreateRequest

type FindRequest struct {
	Name string `url:"name" json:"name"` // The storage identifier.
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _FindRequest FindRequest

type Children struct {
	Msg  string `url:"msg" json:"msg"`   // An optional message about the vdev.
	Name string `url:"name" json:"name"` // The name of the vdev or section.

	// The following parameters are optional
	Cksum *float64 `url:"cksum,omitempty" json:"cksum,omitempty"`
	Read  *float64 `url:"read,omitempty" json:"read,omitempty"`
	State *string  `url:"state,omitempty" json:"state,omitempty"` // The state of the vdev.
	Write *float64 `url:"write,omitempty" json:"write,omitempty"`
}
type _Children Children

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
type _FindResponse FindResponse

type DeleteRequest struct {
	Name string `url:"name" json:"name"` // The storage identifier.
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	CleanupConfig *util.PVEBool `url:"cleanup-config,omitempty" json:"cleanup-config,omitempty"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  *util.PVEBool `url:"cleanup-disks,omitempty" json:"cleanup-disks,omitempty"`   // Also wipe disks so they can be repurposed afterwards.
}
type _DeleteRequest DeleteRequest

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

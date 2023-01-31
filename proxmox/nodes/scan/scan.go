// File generated by proxmox json schema, DO NOT EDIT

package scan

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
	Method string `url:"method" json:"method"`
}
type _IndexResponse IndexResponse

type NfsscanRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Server string `url:"server" json:"server"` // The server address (name or IP).

}
type _NfsscanRequest NfsscanRequest

type NfsscanResponse struct {
	Options string `url:"options" json:"options"` // NFS export options.
	Path    string `url:"path" json:"path"`       // The exported path.

}
type _NfsscanResponse NfsscanResponse

type CifsscanRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Server string `url:"server" json:"server"` // The server address (name or IP).

	// The following parameters are optional
	Domain   *string `url:"domain,omitempty" json:"domain,omitempty"`     // SMB domain (Workgroup).
	Password *string `url:"password,omitempty" json:"password,omitempty"` // User password.
	Username *string `url:"username,omitempty" json:"username,omitempty"` // User name.
}
type _CifsscanRequest CifsscanRequest

type CifsscanResponse struct {
	Description string `url:"description" json:"description"` // Descriptive text from server.
	Share       string `url:"share" json:"share"`             // The cifs share name.

}
type _CifsscanResponse CifsscanResponse

type PbsscanRequest struct {
	Node     string `url:"node" json:"node"`         // The cluster node name.
	Password string `url:"password" json:"password"` // User password or API token secret.
	Server   string `url:"server" json:"server"`     // The server address (name or IP).
	Username string `url:"username" json:"username"` // User-name or API token-ID.

	// The following parameters are optional
	Fingerprint *string `url:"fingerprint,omitempty" json:"fingerprint,omitempty"` // Certificate SHA 256 fingerprint.
	Port        *int    `url:"port,omitempty" json:"port,omitempty"`               // Optional port.
}
type _PbsscanRequest PbsscanRequest

type PbsscanResponse struct {
	Store string `url:"store" json:"store"` // The datastore name.

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"` // Comment from server.
}
type _PbsscanResponse PbsscanResponse

type GlusterfsscanRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Server string `url:"server" json:"server"` // The server address (name or IP).

}
type _GlusterfsscanRequest GlusterfsscanRequest

type GlusterfsscanResponse struct {
	Volname string `url:"volname" json:"volname"` // The volume name.

}
type _GlusterfsscanResponse GlusterfsscanResponse

type IscsiscanRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Portal string `url:"portal" json:"portal"` // The iSCSI portal (IP or DNS name with optional port).

}
type _IscsiscanRequest IscsiscanRequest

type IscsiscanResponse struct {
	Portal string `url:"portal" json:"portal"` // The iSCSI portal name.
	Target string `url:"target" json:"target"` // The iSCSI target name.

}
type _IscsiscanResponse IscsiscanResponse

type LvmscanRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _LvmscanRequest LvmscanRequest

type LvmscanResponse struct {
	Vg string `url:"vg" json:"vg"` // The LVM logical volume group name.

}
type _LvmscanResponse LvmscanResponse

type LvmthinscanRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.
	Vg   string `url:"vg" json:"vg"`
}
type _LvmthinscanRequest LvmthinscanRequest

type LvmthinscanResponse struct {
	Lv string `url:"lv" json:"lv"` // The LVM Thin Pool name (LVM logical volume).

}
type _LvmthinscanResponse LvmthinscanResponse

type ZfsscanRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _ZfsscanRequest ZfsscanRequest

type ZfsscanResponse struct {
	Pool string `url:"pool" json:"pool"` // ZFS pool name.

}
type _ZfsscanResponse ZfsscanResponse

// Index Index of available scan methods
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan", "GET", &resp, req)
	return resp, err
}

// Nfsscan Scan remote NFS server.
func (c *Client) Nfsscan(ctx context.Context, req NfsscanRequest) ([]NfsscanResponse, error) {
	var resp []NfsscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/nfs", "GET", &resp, req)
	return resp, err
}

// Cifsscan Scan remote CIFS server.
func (c *Client) Cifsscan(ctx context.Context, req CifsscanRequest) ([]CifsscanResponse, error) {
	var resp []CifsscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/cifs", "GET", &resp, req)
	return resp, err
}

// Pbsscan Scan remote Proxmox Backup Server.
func (c *Client) Pbsscan(ctx context.Context, req PbsscanRequest) ([]PbsscanResponse, error) {
	var resp []PbsscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/pbs", "GET", &resp, req)
	return resp, err
}

// Glusterfsscan Scan remote GlusterFS server.
func (c *Client) Glusterfsscan(ctx context.Context, req GlusterfsscanRequest) ([]GlusterfsscanResponse, error) {
	var resp []GlusterfsscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/glusterfs", "GET", &resp, req)
	return resp, err
}

// Iscsiscan Scan remote iSCSI server.
func (c *Client) Iscsiscan(ctx context.Context, req IscsiscanRequest) ([]IscsiscanResponse, error) {
	var resp []IscsiscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/iscsi", "GET", &resp, req)
	return resp, err
}

// Lvmscan List local LVM volume groups.
func (c *Client) Lvmscan(ctx context.Context, req LvmscanRequest) ([]LvmscanResponse, error) {
	var resp []LvmscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/lvm", "GET", &resp, req)
	return resp, err
}

// Lvmthinscan List local LVM Thin Pools.
func (c *Client) Lvmthinscan(ctx context.Context, req LvmthinscanRequest) ([]LvmthinscanResponse, error) {
	var resp []LvmthinscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/lvmthin", "GET", &resp, req)
	return resp, err
}

// Zfsscan Scan zfs pool list on local node.
func (c *Client) Zfsscan(ctx context.Context, req ZfsscanRequest) ([]ZfsscanResponse, error) {
	var resp []ZfsscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/scan/zfs", "GET", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package vzdump

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

type CreateRequest struct {

	// The following parameters are optional
	All              *util.SpecialBool `url:"all,omitempty" json:"all,omitempty"`                           // Backup all known guest systems on this host.
	Bwlimit          *int              `url:"bwlimit,omitempty" json:"bwlimit,omitempty"`                   // Limit I/O bandwidth (KBytes per second).
	Compress         *string           `url:"compress,omitempty" json:"compress,omitempty"`                 // Compress dump file.
	Dumpdir          *string           `url:"dumpdir,omitempty" json:"dumpdir,omitempty"`                   // Store resulting files to specified directory.
	Exclude          *string           `url:"exclude,omitempty" json:"exclude,omitempty"`                   // Exclude specified guest systems (assumes --all)
	ExcludePath      *string           `url:"exclude-path,omitempty" json:"exclude-path,omitempty"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	Ionice           *int              `url:"ionice,omitempty" json:"ionice,omitempty"`                     // Set CFQ ionice priority.
	Lockwait         *int              `url:"lockwait,omitempty" json:"lockwait,omitempty"`                 // Maximal time to wait for the global lock (minutes).
	Mailnotification *string           `url:"mailnotification,omitempty" json:"mailnotification,omitempty"` // Specify when to send an email
	Mailto           *string           `url:"mailto,omitempty" json:"mailto,omitempty"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Maxfiles         *int              `url:"maxfiles,omitempty" json:"maxfiles,omitempty"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Mode             *string           `url:"mode,omitempty" json:"mode,omitempty"`                         // Backup mode.
	Node             *string           `url:"node,omitempty" json:"node,omitempty"`                         // Only run if executed on this node.
	NotesTemplate    *string           `url:"notes-template,omitempty" json:"notes-template,omitempty"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	Performance      *string           `url:"performance,omitempty" json:"performance,omitempty"`           // Other performance-related settings.
	Pigz             *int              `url:"pigz,omitempty" json:"pigz,omitempty"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pool             *string           `url:"pool,omitempty" json:"pool,omitempty"`                         // Backup all known guest systems included in the specified pool.
	Protected        *util.SpecialBool `url:"protected,omitempty" json:"protected,omitempty"`               // If true, mark backup(s) as protected.
	PruneBackups     *string           `url:"prune-backups,omitempty" json:"prune-backups,omitempty"`       // Use these retention options instead of those from the storage configuration.
	Quiet            *util.SpecialBool `url:"quiet,omitempty" json:"quiet,omitempty"`                       // Be quiet.
	Remove           *util.SpecialBool `url:"remove,omitempty" json:"remove,omitempty"`                     // Prune older backups according to 'prune-backups'.
	Script           *string           `url:"script,omitempty" json:"script,omitempty"`                     // Use specified hook script.
	Stdexcludes      *util.SpecialBool `url:"stdexcludes,omitempty" json:"stdexcludes,omitempty"`           // Exclude temporary files and logs.
	Stdout           *util.SpecialBool `url:"stdout,omitempty" json:"stdout,omitempty"`                     // Write tar to stdout, not to a file.
	Stop             *util.SpecialBool `url:"stop,omitempty" json:"stop,omitempty"`                         // Stop running backup jobs on this host.
	Stopwait         *int              `url:"stopwait,omitempty" json:"stopwait,omitempty"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Storage          *string           `url:"storage,omitempty" json:"storage,omitempty"`                   // Store resulting file to this storage.
	Tmpdir           *string           `url:"tmpdir,omitempty" json:"tmpdir,omitempty"`                     // Store temporary files to specified directory.
	Vmid             *string           `url:"vmid,omitempty" json:"vmid,omitempty"`                         // The ID of the guest system you want to backup.
	Zstd             *int              `url:"zstd,omitempty" json:"zstd,omitempty"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
}

type CreateResponse string

// Create Create backup.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/vzdump", "POST", &resp, req)
	return resp, err
}

type DefaultsRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Storage *string `url:"storage,omitempty" json:"storage,omitempty"` // The storage identifier.
}

type DefaultsResponse struct {

	// The following parameters are optional
	All              *util.SpecialBool `url:"all,omitempty" json:"all,omitempty"`                           // Backup all known guest systems on this host.
	Bwlimit          *int              `url:"bwlimit,omitempty" json:"bwlimit,omitempty"`                   // Limit I/O bandwidth (KBytes per second).
	Compress         *string           `url:"compress,omitempty" json:"compress,omitempty"`                 // Compress dump file.
	Dumpdir          *string           `url:"dumpdir,omitempty" json:"dumpdir,omitempty"`                   // Store resulting files to specified directory.
	Exclude          *string           `url:"exclude,omitempty" json:"exclude,omitempty"`                   // Exclude specified guest systems (assumes --all)
	ExcludePath      *string           `url:"exclude-path,omitempty" json:"exclude-path,omitempty"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	Ionice           *int              `url:"ionice,omitempty" json:"ionice,omitempty"`                     // Set CFQ ionice priority.
	Lockwait         *int              `url:"lockwait,omitempty" json:"lockwait,omitempty"`                 // Maximal time to wait for the global lock (minutes).
	Mailnotification *string           `url:"mailnotification,omitempty" json:"mailnotification,omitempty"` // Specify when to send an email
	Mailto           *string           `url:"mailto,omitempty" json:"mailto,omitempty"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Maxfiles         *int              `url:"maxfiles,omitempty" json:"maxfiles,omitempty"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Mode             *string           `url:"mode,omitempty" json:"mode,omitempty"`                         // Backup mode.
	Node             *string           `url:"node,omitempty" json:"node,omitempty"`                         // Only run if executed on this node.
	NotesTemplate    *string           `url:"notes-template,omitempty" json:"notes-template,omitempty"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	Performance      *string           `url:"performance,omitempty" json:"performance,omitempty"`           // Other performance-related settings.
	Pigz             *int              `url:"pigz,omitempty" json:"pigz,omitempty"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pool             *string           `url:"pool,omitempty" json:"pool,omitempty"`                         // Backup all known guest systems included in the specified pool.
	Protected        *util.SpecialBool `url:"protected,omitempty" json:"protected,omitempty"`               // If true, mark backup(s) as protected.
	PruneBackups     *string           `url:"prune-backups,omitempty" json:"prune-backups,omitempty"`       // Use these retention options instead of those from the storage configuration.
	Quiet            *util.SpecialBool `url:"quiet,omitempty" json:"quiet,omitempty"`                       // Be quiet.
	Remove           *util.SpecialBool `url:"remove,omitempty" json:"remove,omitempty"`                     // Prune older backups according to 'prune-backups'.
	Script           *string           `url:"script,omitempty" json:"script,omitempty"`                     // Use specified hook script.
	Stdexcludes      *util.SpecialBool `url:"stdexcludes,omitempty" json:"stdexcludes,omitempty"`           // Exclude temporary files and logs.
	Stop             *util.SpecialBool `url:"stop,omitempty" json:"stop,omitempty"`                         // Stop running backup jobs on this host.
	Stopwait         *int              `url:"stopwait,omitempty" json:"stopwait,omitempty"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Storage          *string           `url:"storage,omitempty" json:"storage,omitempty"`                   // Store resulting file to this storage.
	Tmpdir           *string           `url:"tmpdir,omitempty" json:"tmpdir,omitempty"`                     // Store temporary files to specified directory.
	Vmid             *string           `url:"vmid,omitempty" json:"vmid,omitempty"`                         // The ID of the guest system you want to backup.
	Zstd             *int              `url:"zstd,omitempty" json:"zstd,omitempty"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
}

// Defaults Get the currently configured vzdump defaults.
func (c *Client) Defaults(ctx context.Context, req *DefaultsRequest) (*DefaultsResponse, error) {
	var resp *DefaultsResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/vzdump/defaults", "GET", &resp, req)
	return resp, err
}

type ExtractconfigRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Volume string `url:"volume" json:"volume"` // Volume identifier

}

type ExtractconfigResponse string

// Extractconfig Extract configuration from vzdump backup archive.
func (c *Client) Extractconfig(ctx context.Context, req *ExtractconfigRequest) (*ExtractconfigResponse, error) {
	var resp *ExtractconfigResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/vzdump/extractconfig", "GET", &resp, req)
	return resp, err
}

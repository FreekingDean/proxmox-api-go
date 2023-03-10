// File generated by proxmox json schema, DO NOT EDIT

package backup

import (
	"context"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
)

const (
	ChildrenType_QEMU    ChildrenType = "qemu"
	ChildrenType_LXC     ChildrenType = "lxc"
	ChildrenType_UNKNOWN ChildrenType = "unknown"

	Compress_0    Compress = "0"
	Compress_1    Compress = "1"
	Compress_GZIP Compress = "gzip"
	Compress_LZO  Compress = "lzo"
	Compress_ZSTD Compress = "zstd"

	Mailnotification_ALWAYS  Mailnotification = "always"
	Mailnotification_FAILURE Mailnotification = "failure"

	Mode_SNAPSHOT Mode = "snapshot"
	Mode_SUSPEND  Mode = "suspend"
	Mode_STOP     Mode = "stop"
)

type ChildrenType string
type Compress string
type Mailnotification string
type Mode string

func PtrChildrenType(i ChildrenType) *ChildrenType {
	return &i
}
func PtrCompress(i Compress) *Compress {
	return &i
}
func PtrMailnotification(i Mailnotification) *Mailnotification {
	return &i
}
func PtrMode(i Mode) *Mode {
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

type IndexResponse struct {
	Id string `url:"id" json:"id"` // The job ID.

}
type _IndexResponse IndexResponse

type CreateRequest struct {

	// The following parameters are optional
	All              *util.PVEBool     `url:"all,omitempty" json:"all,omitempty"`                           // Backup all known guest systems on this host.
	Bwlimit          *int              `url:"bwlimit,omitempty" json:"bwlimit,omitempty"`                   // Limit I/O bandwidth (KBytes per second).
	Comment          *string           `url:"comment,omitempty" json:"comment,omitempty"`                   // Description for the Job.
	Compress         *Compress         `url:"compress,omitempty" json:"compress,omitempty"`                 // Compress dump file.
	Dow              *string           `url:"dow,omitempty" json:"dow,omitempty"`                           // Day of week selection.
	Dumpdir          *string           `url:"dumpdir,omitempty" json:"dumpdir,omitempty"`                   // Store resulting files to specified directory.
	Enabled          *util.PVEBool     `url:"enabled,omitempty" json:"enabled,omitempty"`                   // Enable or disable the job.
	Exclude          *string           `url:"exclude,omitempty" json:"exclude,omitempty"`                   // Exclude specified guest systems (assumes --all)
	ExcludePath      *string           `url:"exclude-path,omitempty" json:"exclude-path,omitempty"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root, other paths match relative to each subdirectory.
	Id               *string           `url:"id,omitempty" json:"id,omitempty"`                             // Job ID (will be autogenerated).
	Ionice           *int              `url:"ionice,omitempty" json:"ionice,omitempty"`                     // Set CFQ ionice priority.
	Lockwait         *int              `url:"lockwait,omitempty" json:"lockwait,omitempty"`                 // Maximal time to wait for the global lock (minutes).
	Mailnotification *Mailnotification `url:"mailnotification,omitempty" json:"mailnotification,omitempty"` // Specify when to send an email
	Mailto           *string           `url:"mailto,omitempty" json:"mailto,omitempty"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Maxfiles         *int              `url:"maxfiles,omitempty" json:"maxfiles,omitempty"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Mode             *Mode             `url:"mode,omitempty" json:"mode,omitempty"`                         // Backup mode.
	Node             *string           `url:"node,omitempty" json:"node,omitempty"`                         // Only run if executed on this node.
	NotesTemplate    *string           `url:"notes-template,omitempty" json:"notes-template,omitempty"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	Performance      *string           `url:"performance,omitempty" json:"performance,omitempty"`           // Other performance-related settings.
	Pigz             *int              `url:"pigz,omitempty" json:"pigz,omitempty"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pool             *string           `url:"pool,omitempty" json:"pool,omitempty"`                         // Backup all known guest systems included in the specified pool.
	Protected        *util.PVEBool     `url:"protected,omitempty" json:"protected,omitempty"`               // If true, mark backup(s) as protected.
	PruneBackups     *string           `url:"prune-backups,omitempty" json:"prune-backups,omitempty"`       // Use these retention options instead of those from the storage configuration.
	Quiet            *util.PVEBool     `url:"quiet,omitempty" json:"quiet,omitempty"`                       // Be quiet.
	Remove           *util.PVEBool     `url:"remove,omitempty" json:"remove,omitempty"`                     // Prune older backups according to 'prune-backups'.
	RepeatMissed     *util.PVEBool     `url:"repeat-missed,omitempty" json:"repeat-missed,omitempty"`       // If true, the job will be run as soon as possible if it was missed while the scheduler was not running.
	Schedule         *string           `url:"schedule,omitempty" json:"schedule,omitempty"`                 // Backup schedule. The format is a subset of `systemd` calendar events.
	Script           *string           `url:"script,omitempty" json:"script,omitempty"`                     // Use specified hook script.
	Starttime        *string           `url:"starttime,omitempty" json:"starttime,omitempty"`               // Job Start time.
	Stdexcludes      *util.PVEBool     `url:"stdexcludes,omitempty" json:"stdexcludes,omitempty"`           // Exclude temporary files and logs.
	Stop             *util.PVEBool     `url:"stop,omitempty" json:"stop,omitempty"`                         // Stop running backup jobs on this host.
	Stopwait         *int              `url:"stopwait,omitempty" json:"stopwait,omitempty"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Storage          *string           `url:"storage,omitempty" json:"storage,omitempty"`                   // Store resulting file to this storage.
	Tmpdir           *string           `url:"tmpdir,omitempty" json:"tmpdir,omitempty"`                     // Store temporary files to specified directory.
	Vmid             *string           `url:"vmid,omitempty" json:"vmid,omitempty"`                         // The ID of the guest system you want to backup.
	Zstd             *int              `url:"zstd,omitempty" json:"zstd,omitempty"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
}
type _CreateRequest CreateRequest

type FindRequest struct {
	Id string `url:"id" json:"id"` // The job ID.

}
type _FindRequest FindRequest

type UpdateRequest struct {
	Id string `url:"id" json:"id"` // The job ID.

	// The following parameters are optional
	All              *util.PVEBool     `url:"all,omitempty" json:"all,omitempty"`                           // Backup all known guest systems on this host.
	Bwlimit          *int              `url:"bwlimit,omitempty" json:"bwlimit,omitempty"`                   // Limit I/O bandwidth (KBytes per second).
	Comment          *string           `url:"comment,omitempty" json:"comment,omitempty"`                   // Description for the Job.
	Compress         *Compress         `url:"compress,omitempty" json:"compress,omitempty"`                 // Compress dump file.
	Delete           *string           `url:"delete,omitempty" json:"delete,omitempty"`                     // A list of settings you want to delete.
	Dow              *string           `url:"dow,omitempty" json:"dow,omitempty"`                           // Day of week selection.
	Dumpdir          *string           `url:"dumpdir,omitempty" json:"dumpdir,omitempty"`                   // Store resulting files to specified directory.
	Enabled          *util.PVEBool     `url:"enabled,omitempty" json:"enabled,omitempty"`                   // Enable or disable the job.
	Exclude          *string           `url:"exclude,omitempty" json:"exclude,omitempty"`                   // Exclude specified guest systems (assumes --all)
	ExcludePath      *string           `url:"exclude-path,omitempty" json:"exclude-path,omitempty"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root, other paths match relative to each subdirectory.
	Ionice           *int              `url:"ionice,omitempty" json:"ionice,omitempty"`                     // Set CFQ ionice priority.
	Lockwait         *int              `url:"lockwait,omitempty" json:"lockwait,omitempty"`                 // Maximal time to wait for the global lock (minutes).
	Mailnotification *Mailnotification `url:"mailnotification,omitempty" json:"mailnotification,omitempty"` // Specify when to send an email
	Mailto           *string           `url:"mailto,omitempty" json:"mailto,omitempty"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Maxfiles         *int              `url:"maxfiles,omitempty" json:"maxfiles,omitempty"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Mode             *Mode             `url:"mode,omitempty" json:"mode,omitempty"`                         // Backup mode.
	Node             *string           `url:"node,omitempty" json:"node,omitempty"`                         // Only run if executed on this node.
	NotesTemplate    *string           `url:"notes-template,omitempty" json:"notes-template,omitempty"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	Performance      *string           `url:"performance,omitempty" json:"performance,omitempty"`           // Other performance-related settings.
	Pigz             *int              `url:"pigz,omitempty" json:"pigz,omitempty"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pool             *string           `url:"pool,omitempty" json:"pool,omitempty"`                         // Backup all known guest systems included in the specified pool.
	Protected        *util.PVEBool     `url:"protected,omitempty" json:"protected,omitempty"`               // If true, mark backup(s) as protected.
	PruneBackups     *string           `url:"prune-backups,omitempty" json:"prune-backups,omitempty"`       // Use these retention options instead of those from the storage configuration.
	Quiet            *util.PVEBool     `url:"quiet,omitempty" json:"quiet,omitempty"`                       // Be quiet.
	Remove           *util.PVEBool     `url:"remove,omitempty" json:"remove,omitempty"`                     // Prune older backups according to 'prune-backups'.
	RepeatMissed     *util.PVEBool     `url:"repeat-missed,omitempty" json:"repeat-missed,omitempty"`       // If true, the job will be run as soon as possible if it was missed while the scheduler was not running.
	Schedule         *string           `url:"schedule,omitempty" json:"schedule,omitempty"`                 // Backup schedule. The format is a subset of `systemd` calendar events.
	Script           *string           `url:"script,omitempty" json:"script,omitempty"`                     // Use specified hook script.
	Starttime        *string           `url:"starttime,omitempty" json:"starttime,omitempty"`               // Job Start time.
	Stdexcludes      *util.PVEBool     `url:"stdexcludes,omitempty" json:"stdexcludes,omitempty"`           // Exclude temporary files and logs.
	Stop             *util.PVEBool     `url:"stop,omitempty" json:"stop,omitempty"`                         // Stop running backup jobs on this host.
	Stopwait         *int              `url:"stopwait,omitempty" json:"stopwait,omitempty"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Storage          *string           `url:"storage,omitempty" json:"storage,omitempty"`                   // Store resulting file to this storage.
	Tmpdir           *string           `url:"tmpdir,omitempty" json:"tmpdir,omitempty"`                     // Store temporary files to specified directory.
	Vmid             *string           `url:"vmid,omitempty" json:"vmid,omitempty"`                         // The ID of the guest system you want to backup.
	Zstd             *int              `url:"zstd,omitempty" json:"zstd,omitempty"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
}
type _UpdateRequest UpdateRequest

type DeleteRequest struct {
	Id string `url:"id" json:"id"` // The job ID.

}
type _DeleteRequest DeleteRequest

type GetVolumeBackupIncludedIncludedVolumesRequest struct {
	Id string `url:"id" json:"id"` // The job ID.

}
type _GetVolumeBackupIncludedIncludedVolumesRequest GetVolumeBackupIncludedIncludedVolumesRequest

type Children struct {
	Id       string       `url:"id" json:"id"`             // Configuration key of the volume.
	Included util.PVEBool `url:"included" json:"included"` // Whether the volume is included in the backup or not.
	Name     string       `url:"name" json:"name"`         // Name of the volume.
	Reason   string       `url:"reason" json:"reason"`     // The reason why the volume is included (or excluded).

}
type _Children Children

type SubChildren struct {
	Id   int          `url:"id" json:"id"`     // VMID of the guest.
	Type ChildrenType `url:"type" json:"type"` // Type of the guest, VM, CT or unknown for removed but not purged guests.

	// The following parameters are optional
	Children *[]Children `url:"children,omitempty" json:"children,omitempty"` // The volumes of the guest with the information if they will be included in backups.
	Name     *string     `url:"name,omitempty" json:"name,omitempty"`         // Name of the guest
}
type _SubChildren SubChildren

// Root node of the tree object. Children represent guests, grandchildren represent volumes of that guest.
type GetVolumeBackupIncludedIncludedVolumesResponse struct {
	Children []Children `url:"children" json:"children"`
}
type _GetVolumeBackupIncludedIncludedVolumesResponse GetVolumeBackupIncludedIncludedVolumesResponse

// Index List vzdump backup schedule.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/backup", "GET", &resp, nil)
	return resp, err
}

// Create Create new vzdump backup job.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/backup", "POST", nil, req)
	return err
}

// Find Read vzdump backup job definition.
func (c *Client) Find(ctx context.Context, req FindRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}", "GET", &resp, req)
	return resp, err
}

// Update Update vzdump backup job definition.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}", "PUT", nil, req)
	return err
}

// Delete Delete vzdump backup job definition.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}", "DELETE", nil, req)
	return err
}

// GetVolumeBackupIncludedIncludedVolumes Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
func (c *Client) GetVolumeBackupIncludedIncludedVolumes(ctx context.Context, req GetVolumeBackupIncludedIncludedVolumesRequest) (GetVolumeBackupIncludedIncludedVolumesResponse, error) {
	var resp GetVolumeBackupIncludedIncludedVolumesResponse

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}/included_volumes", "GET", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package backup

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

type IndexResponse []*struct {
	Id string `url:"id",json:"id"`
}

// Index List vzdump backup schedule.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/backup", "GET", &resp, nil)
	return resp, err
}

type CreateRequest struct {
	All              *bool   `url:"all,omitempty",json:"all,omitempty"`
	Enabled          *bool   `url:"enabled,omitempty",json:"enabled,omitempty"`
	Maxfiles         *int    `url:"maxfiles,omitempty",json:"maxfiles,omitempty"`
	RepeatMissed     *bool   `url:"repeat-missed,omitempty",json:"repeat-missed,omitempty"`
	Stopwait         *int    `url:"stopwait,omitempty",json:"stopwait,omitempty"`
	Storage          *string `url:"storage,omitempty",json:"storage,omitempty"`
	Bwlimit          *int    `url:"bwlimit,omitempty",json:"bwlimit,omitempty"`
	Exclude          *string `url:"exclude,omitempty",json:"exclude,omitempty"`
	Ionice           *int    `url:"ionice,omitempty",json:"ionice,omitempty"`
	NotesTemplate    *string `url:"notes-template,omitempty",json:"notes-template,omitempty"`
	Zstd             *int    `url:"zstd,omitempty",json:"zstd,omitempty"`
	Performance      *string `url:"performance,omitempty",json:"performance,omitempty"`
	Remove           *bool   `url:"remove,omitempty",json:"remove,omitempty"`
	Script           *string `url:"script,omitempty",json:"script,omitempty"`
	Tmpdir           *string `url:"tmpdir,omitempty",json:"tmpdir,omitempty"`
	Stop             *bool   `url:"stop,omitempty",json:"stop,omitempty"`
	Vmid             *string `url:"vmid,omitempty",json:"vmid,omitempty"`
	Comment          *string `url:"comment,omitempty",json:"comment,omitempty"`
	ExcludePath      *string `url:"exclude-path,omitempty",json:"exclude-path,omitempty"`
	Mode             *string `url:"mode,omitempty",json:"mode,omitempty"`
	Protected        *bool   `url:"protected,omitempty",json:"protected,omitempty"`
	Starttime        *string `url:"starttime,omitempty",json:"starttime,omitempty"`
	Compress         *string `url:"compress,omitempty",json:"compress,omitempty"`
	Dumpdir          *string `url:"dumpdir,omitempty",json:"dumpdir,omitempty"`
	Id               *string `url:"id,omitempty",json:"id,omitempty"`
	Quiet            *bool   `url:"quiet,omitempty",json:"quiet,omitempty"`
	Pool             *string `url:"pool,omitempty",json:"pool,omitempty"`
	Schedule         *string `url:"schedule,omitempty",json:"schedule,omitempty"`
	Dow              *string `url:"dow,omitempty",json:"dow,omitempty"`
	Mailto           *string `url:"mailto,omitempty",json:"mailto,omitempty"`
	PruneBackups     *string `url:"prune-backups,omitempty",json:"prune-backups,omitempty"`
	Stdexcludes      *bool   `url:"stdexcludes,omitempty",json:"stdexcludes,omitempty"`
	Lockwait         *int    `url:"lockwait,omitempty",json:"lockwait,omitempty"`
	Mailnotification *string `url:"mailnotification,omitempty",json:"mailnotification,omitempty"`
	Node             *string `url:"node,omitempty",json:"node,omitempty"`
	Pigz             *int    `url:"pigz,omitempty",json:"pigz,omitempty"`
}

type CreateResponse map[string]interface{}

// Create Create new vzdump backup job.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/cluster/backup", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Id string `url:"id",json:"id"`
}

type FindResponse map[string]interface{}

// Find Read vzdump backup job definition.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Exclude          *string `url:"exclude,omitempty",json:"exclude,omitempty"`
	Stop             *bool   `url:"stop,omitempty",json:"stop,omitempty"`
	Storage          *string `url:"storage,omitempty",json:"storage,omitempty"`
	Compress         *string `url:"compress,omitempty",json:"compress,omitempty"`
	Mailto           *string `url:"mailto,omitempty",json:"mailto,omitempty"`
	PruneBackups     *string `url:"prune-backups,omitempty",json:"prune-backups,omitempty"`
	Remove           *bool   `url:"remove,omitempty",json:"remove,omitempty"`
	RepeatMissed     *bool   `url:"repeat-missed,omitempty",json:"repeat-missed,omitempty"`
	Schedule         *string `url:"schedule,omitempty",json:"schedule,omitempty"`
	Script           *string `url:"script,omitempty",json:"script,omitempty"`
	Dow              *string `url:"dow,omitempty",json:"dow,omitempty"`
	ExcludePath      *string `url:"exclude-path,omitempty",json:"exclude-path,omitempty"`
	Maxfiles         *int    `url:"maxfiles,omitempty",json:"maxfiles,omitempty"`
	Performance      *string `url:"performance,omitempty",json:"performance,omitempty"`
	Pigz             *int    `url:"pigz,omitempty",json:"pigz,omitempty"`
	Stopwait         *int    `url:"stopwait,omitempty",json:"stopwait,omitempty"`
	All              *bool   `url:"all,omitempty",json:"all,omitempty"`
	Comment          *string `url:"comment,omitempty",json:"comment,omitempty"`
	Lockwait         *int    `url:"lockwait,omitempty",json:"lockwait,omitempty"`
	Mailnotification *string `url:"mailnotification,omitempty",json:"mailnotification,omitempty"`
	Vmid             *string `url:"vmid,omitempty",json:"vmid,omitempty"`
	Bwlimit          *int    `url:"bwlimit,omitempty",json:"bwlimit,omitempty"`
	Mode             *string `url:"mode,omitempty",json:"mode,omitempty"`
	Protected        *bool   `url:"protected,omitempty",json:"protected,omitempty"`
	Tmpdir           *string `url:"tmpdir,omitempty",json:"tmpdir,omitempty"`
	Node             *string `url:"node,omitempty",json:"node,omitempty"`
	Stdexcludes      *bool   `url:"stdexcludes,omitempty",json:"stdexcludes,omitempty"`
	Delete           *string `url:"delete,omitempty",json:"delete,omitempty"`
	Dumpdir          *string `url:"dumpdir,omitempty",json:"dumpdir,omitempty"`
	Id               string  `url:"id",json:"id"`
	Ionice           *int    `url:"ionice,omitempty",json:"ionice,omitempty"`
	NotesTemplate    *string `url:"notes-template,omitempty",json:"notes-template,omitempty"`
	Quiet            *bool   `url:"quiet,omitempty",json:"quiet,omitempty"`
	Enabled          *bool   `url:"enabled,omitempty",json:"enabled,omitempty"`
	Pool             *string `url:"pool,omitempty",json:"pool,omitempty"`
	Starttime        *string `url:"starttime,omitempty",json:"starttime,omitempty"`
	Zstd             *int    `url:"zstd,omitempty",json:"zstd,omitempty"`
}

type UpdateResponse map[string]interface{}

// Update Update vzdump backup job definition.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Id string `url:"id",json:"id"`
}

type DeleteResponse map[string]interface{}

// Delete Delete vzdump backup job definition.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}", "DELETE", &resp, req)
	return resp, err
}

type GetVolumeBackupIncludedIncludedVolumesRequest struct {
	Id string `url:"id",json:"id"`
}

type GetVolumeBackupIncludedIncludedVolumesResponse struct {
	Children []*struct {
		Id       int     `url:"id",json:"id"`
		Name     *string `url:"name,omitempty",json:"name,omitempty"`
		Type     string  `url:"type",json:"type"`
		Children []*struct {
			Id       string `url:"id",json:"id"`
			Included bool   `url:"included",json:"included"`
			Name     string `url:"name",json:"name"`
			Reason   string `url:"reason",json:"reason"`
		} `url:"children,omitempty",json:"children,omitempty"`
	} `url:"children",json:"children"`
}

// GetVolumeBackupIncludedIncludedVolumes Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
func (c *Client) GetVolumeBackupIncludedIncludedVolumes(ctx context.Context, req *GetVolumeBackupIncludedIncludedVolumesRequest) (*GetVolumeBackupIncludedIncludedVolumesResponse, error) {
	var resp *GetVolumeBackupIncludedIncludedVolumesResponse

	err := c.httpClient.Do(ctx, "/cluster/backup/{id}/included_volumes", "GET", &resp, req)
	return resp, err
}

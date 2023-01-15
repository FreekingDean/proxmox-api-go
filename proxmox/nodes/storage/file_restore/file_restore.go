// File generated by proxmox json schema, DO NOT EDIT

package file_restore

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

type ListRequest struct {
	Filepath string `url:"filepath" json:"filepath"` // base64-path to the directory or file being listed, or "/".
	Node     string `url:"node" json:"node"`         // The cluster node name.
	Storage  string `url:"storage" json:"storage"`   // The storage identifier.
	Volume   string `url:"volume" json:"volume"`     // Backup volume ID or name. Currently only PBS snapshots are supported.

}

type ListResponse []*struct {
	Filepath string           `url:"filepath" json:"filepath"` // base64 path of the current entry
	Leaf     util.SpecialBool `url:"leaf" json:"leaf"`         // If this entry is a leaf in the directory graph.
	Text     string           `url:"text" json:"text"`         // Entry display text.
	Type     string           `url:"type" json:"type"`         // Entry type.

	// The following parameters are optional
	Mtime *int `url:"mtime,omitempty" json:"mtime,omitempty"` // Entry last-modified time (unix timestamp).
	Size  *int `url:"size,omitempty" json:"size,omitempty"`   // Entry file size.
}

// List List files and directories for single file restore under the given path.
func (c *Client) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	var resp *ListResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/file-restore/list", "GET", &resp, req)
	return resp, err
}

type DownloadRequest struct {
	Filepath string `url:"filepath" json:"filepath"` // base64-path to the directory or file to download.
	Node     string `url:"node" json:"node"`         // The cluster node name.
	Storage  string `url:"storage" json:"storage"`   // The storage identifier.
	Volume   string `url:"volume" json:"volume"`     // Backup volume ID or name. Currently only PBS snapshots are supported.

}

type DownloadResponse interface{}

// Download Extract a file or directory (as zip archive) from a PBS backup.
func (c *Client) Download(ctx context.Context, req *DownloadRequest) (*DownloadResponse, error) {
	var resp *DownloadResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/file-restore/download", "GET", &resp, req)
	return resp, err
}

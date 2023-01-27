// File generated by proxmox json schema, DO NOT EDIT

package apt

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
	Node string `url:"node" json:"node"` // The cluster node name.

}

type IndexResponse struct {
	Id string `url:"id" json:"id"`
}

type ListUpdatesUpdateRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}

type UpdateDatabaseUpdateRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Notify *util.SpecialBool `url:"notify,omitempty" json:"notify,omitempty"` // Send notification mail about new packages (to email address specified for user 'root@pam').
	Quiet  *util.SpecialBool `url:"quiet,omitempty" json:"quiet,omitempty"`   // Only produces output suitable for logging, omitting progress indicators.
}

type ChangelogRequest struct {
	Name string `url:"name" json:"name"` // Package name.
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Version *string `url:"version,omitempty" json:"version,omitempty"` // Package version.
}

type RepositoriesRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}

type Errors struct {
	Error string `url:"error" json:"error"` // The error message
	Path  string `url:"path" json:"path"`   // Path to the problematic file.

}

type Options struct {
	Key    string   `url:"Key" json:"Key"`
	Values []string `url:"Values" json:"Values"`
}

type Repositories struct {
	Enabled  util.SpecialBool `url:"Enabled" json:"Enabled"`   // Whether the repository is enabled or not
	Filetype string           `url:"FileType" json:"FileType"` // Format of the defining file.
	Suites   []string         `url:"Suites" json:"Suites"`     // List of package distribuitions
	Types    []string         `url:"Types" json:"Types"`       // List of package types.
	Uris     []string         `url:"URIs" json:"URIs"`         // List of repository URIs.

	// The following parameters are optional
	Comment    *string   `url:"Comment,omitempty" json:"Comment,omitempty"`       // Associated comment
	Components []string  `url:"Components,omitempty" json:"Components,omitempty"` // List of repository components
	Options    []Options `url:"Options,omitempty" json:"Options,omitempty"`       // Additional options
}

type Files struct {
	Digest       []int          `url:"digest" json:"digest"`             // Digest of the file as bytes.
	FileType     string         `url:"file-type" json:"file-type"`       // Format of the file.
	Path         string         `url:"path" json:"path"`                 // Path to the problematic file.
	Repositories []Repositories `url:"repositories" json:"repositories"` // The parsed repositories.

}

type Infos struct {
	Index   string `url:"index" json:"index"`     // Index of the associated repository within the file.
	Kind    string `url:"kind" json:"kind"`       // Kind of the information (e.g. warning).
	Message string `url:"message" json:"message"` // Information message.
	Path    string `url:"path" json:"path"`       // Path to the associated file.

	// The following parameters are optional
	Property *string `url:"property,omitempty" json:"property,omitempty"` // Property from which the info originates.
}

type StandardRepos struct {
	Handle string `url:"handle" json:"handle"` // Handle to identify the repository.
	Name   string `url:"name" json:"name"`     // Full name of the repository.

	// The following parameters are optional
	Status *util.SpecialBool `url:"status,omitempty" json:"status,omitempty"` // Indicating enabled/disabled status, if the repository is configured.
}

// Result from parsing the APT repository files in /etc/apt/.
type RepositoriesResponse struct {
	Digest        string          `url:"digest" json:"digest"`                 // Common digest of all files.
	Errors        []Errors        `url:"errors" json:"errors"`                 // List of problematic repository files.
	Files         []Files         `url:"files" json:"files"`                   // List of parsed repository files.
	Infos         []Infos         `url:"infos" json:"infos"`                   // Additional information/warnings for APT repositories.
	StandardRepos []StandardRepos `url:"standard-repos" json:"standard-repos"` // List of standard repositories and their configuration status

}

type ChangeRepositoryRepositoriesRequest struct {
	Index int    `url:"index" json:"index"` // Index within the file (starting from 0).
	Node  string `url:"node" json:"node"`   // The cluster node name.
	Path  string `url:"path" json:"path"`   // Path to the containing file.

	// The following parameters are optional
	Digest  *string           `url:"digest,omitempty" json:"digest,omitempty"`   // Digest to detect modifications.
	Enabled *util.SpecialBool `url:"enabled,omitempty" json:"enabled,omitempty"` // Whether the repository should be enabled or not.
}

type AddRepositoryRepositoriesRequest struct {
	Handle string `url:"handle" json:"handle"` // Handle that identifies a repository.
	Node   string `url:"node" json:"node"`     // The cluster node name.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty" json:"digest,omitempty"` // Digest to detect modifications.
}

type VersionsRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}

// Index Directory index for apt (Advanced Package Tool).
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt", "GET", &resp, req)
	return resp, err
}

// ListUpdatesUpdate List available updates.
func (c *Client) ListUpdatesUpdate(ctx context.Context, req ListUpdatesUpdateRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/update", "GET", &resp, req)
	return resp, err
}

// UpdateDatabaseUpdate This is used to resynchronize the package index files from their sources (apt-get update).
func (c *Client) UpdateDatabaseUpdate(ctx context.Context, req UpdateDatabaseUpdateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/update", "POST", &resp, req)
	return resp, err
}

// Changelog Get package changelogs.
func (c *Client) Changelog(ctx context.Context, req ChangelogRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/changelog", "GET", &resp, req)
	return resp, err
}

// Repositories Get APT repository information.
func (c *Client) Repositories(ctx context.Context, req RepositoriesRequest) (RepositoriesResponse, error) {
	var resp RepositoriesResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/repositories", "GET", &resp, req)
	return resp, err
}

// ChangeRepositoryRepositories Change the properties of a repository. Currently only allows enabling/disabling.
func (c *Client) ChangeRepositoryRepositories(ctx context.Context, req ChangeRepositoryRepositoriesRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/repositories", "POST", nil, req)
	return err
}

// AddRepositoryRepositories Add a standard repository to the configuration
func (c *Client) AddRepositoryRepositories(ctx context.Context, req AddRepositoryRepositoriesRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/repositories", "PUT", nil, req)
	return err
}

// Versions Get package information for important Proxmox packages.
func (c *Client) Versions(ctx context.Context, req VersionsRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/apt/versions", "GET", &resp, req)
	return resp, err
}

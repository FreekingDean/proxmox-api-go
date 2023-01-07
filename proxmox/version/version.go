// File generated by proxmox json schema, DO NOT EDIT

package version

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

type IndexResponse struct {
	Release string `url:"release",json:"release"` // The current Proxmox VE point release in `x.y` format.
	Repoid  string `url:"repoid",json:"repoid"`   // The short git revision from which this version was build.
	Version string `url:"version",json:"version"` // The full pve-manager package version of this node.

	// The following parameters are optional
	Console *string `url:"console,omitempty",json:"console,omitempty"` // The default console viewer to use.
}

// Index API version details, including some parts of the global datacenter config.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/version", "GET", &resp, nil)
	return resp, err
}
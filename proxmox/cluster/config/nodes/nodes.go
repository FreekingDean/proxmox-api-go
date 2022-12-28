// File generated by proxmox json schema, DO NOT EDIT

package nodes

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
	Node string `url:"node",json:"node"`
}

// Index Corosync node list.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/config/nodes", "GET", &resp, nil)
	return resp, err
}

type ChildCreateRequest struct {
	Force      *bool   `url:"force,omitempty",json:"force,omitempty"`
	Linkn      *string `url:"link[n],omitempty",json:"link[n],omitempty"`
	NewNodeIp  *string `url:"new_node_ip,omitempty",json:"new_node_ip,omitempty"`
	Node       string  `url:"node",json:"node"`
	Nodeid     *int    `url:"nodeid,omitempty",json:"nodeid,omitempty"`
	Votes      *int    `url:"votes,omitempty",json:"votes,omitempty"`
	Apiversion *int    `url:"apiversion,omitempty",json:"apiversion,omitempty"`
}

type ChildCreateResponse struct {
	CorosyncAuthkey string   `url:"corosync_authkey",json:"corosync_authkey"`
	CorosyncConf    string   `url:"corosync_conf",json:"corosync_conf"`
	Warnings        []string `url:"warnings",json:"warnings"`
}

// ChildCreate Adds a node to the cluster configuration. This call is for internal use.
func (c *Client) ChildCreate(ctx context.Context, req *ChildCreateRequest) (*ChildCreateResponse, error) {
	var resp *ChildCreateResponse

	err := c.httpClient.Do(ctx, "/cluster/config/nodes/{node}", "POST", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Node string `url:"node",json:"node"`
}

type DeleteResponse map[string]interface{}

// Delete Removes a node from the cluster configuration.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/config/nodes/{node}", "DELETE", &resp, req)
	return resp, err
}

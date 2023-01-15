// File generated by proxmox json schema, DO NOT EDIT

package config

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

type IndexResponse []*map[string]interface{}

// Index Directory index.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/config", "GET", &resp, nil)
	return resp, err
}

type CreateRequest struct {
	Clustername string `url:"clustername",json:"clustername"` // The name of the cluster.

	// The following parameters are optional
	Linkn  *string `url:"link[n],omitempty",json:"link[n],omitempty"` // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid *int    `url:"nodeid,omitempty",json:"nodeid,omitempty"`   // Node id for this node.
	Votes  *int    `url:"votes,omitempty",json:"votes,omitempty"`     // Number of votes for this node.
}

type CreateResponse string

// Create Generate new cluster configuration. If no links given, default to local IP address as link0.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/cluster/config", "POST", &resp, req)
	return resp, err
}

type JoinApiVersionApiversionResponse int

// JoinApiVersionApiversion Return the version of the cluster join API available on this node.
func (c *Client) JoinApiVersionApiversion(ctx context.Context) (*JoinApiVersionApiversionResponse, error) {
	var resp *JoinApiVersionApiversionResponse

	err := c.httpClient.Do(ctx, "/cluster/config/apiversion", "GET", &resp, nil)
	return resp, err
}

type JoinInfoJoinRequest struct {

	// The following parameters are optional
	Node *string `url:"node,omitempty",json:"node,omitempty"` // The node for which the joinee gets the nodeinfo.
}

type JoinInfoJoinResponse struct {
	ConfigDigest string `url:"config_digest",json:"config_digest"`
	Nodelist     []*struct {
		Name        string `url:"name",json:"name"` // The cluster node name.
		PveAddr     string `url:"pve_addr",json:"pve_addr"`
		PveFp       string `url:"pve_fp",json:"pve_fp"` // Certificate SHA 256 fingerprint.
		QuorumVotes int    `url:"quorum_votes",json:"quorum_votes"`

		// The following parameters are optional
		Nodeid    *int    `url:"nodeid,omitempty",json:"nodeid,omitempty"`         // Node id for this node.
		Ring0Addr *string `url:"ring0_addr,omitempty",json:"ring0_addr,omitempty"` // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	} `url:"nodelist",json:"nodelist"`
	PreferredNode string                 `url:"preferred_node",json:"preferred_node"` // The cluster node name.
	Totem         map[string]interface{} `url:"totem",json:"totem"`
}

// JoinInfoJoin Get information needed to join this cluster over the connected node.
func (c *Client) JoinInfoJoin(ctx context.Context, req *JoinInfoJoinRequest) (*JoinInfoJoinResponse, error) {
	var resp *JoinInfoJoinResponse

	err := c.httpClient.Do(ctx, "/cluster/config/join", "GET", &resp, req)
	return resp, err
}

type JoinRequest struct {
	Fingerprint string `url:"fingerprint",json:"fingerprint"` // Certificate SHA 256 fingerprint.
	Hostname    string `url:"hostname",json:"hostname"`       // Hostname (or IP) of an existing cluster member.
	Password    string `url:"password",json:"password"`       // Superuser (root) password of peer node.

	// The following parameters are optional
	Force  *util.SpecialBool `url:"force,omitempty",json:"force,omitempty"`     // Do not throw error if node already exists.
	Linkn  *string           `url:"link[n],omitempty",json:"link[n],omitempty"` // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid *int              `url:"nodeid,omitempty",json:"nodeid,omitempty"`   // Node id for this node.
	Votes  *int              `url:"votes,omitempty",json:"votes,omitempty"`     // Number of votes for this node
}

type JoinResponse string

// Join Joins this node into an existing cluster. If no links are given, default to IP resolved by node's hostname on single link (fallback fails for clusters with multiple links).
func (c *Client) Join(ctx context.Context, req *JoinRequest) (*JoinResponse, error) {
	var resp *JoinResponse

	err := c.httpClient.Do(ctx, "/cluster/config/join", "POST", &resp, req)
	return resp, err
}

type TotemResponse map[string]interface{}

// Totem Get corosync totem protocol settings.
func (c *Client) Totem(ctx context.Context) (*TotemResponse, error) {
	var resp *TotemResponse

	err := c.httpClient.Do(ctx, "/cluster/config/totem", "GET", &resp, nil)
	return resp, err
}

type StatusQdeviceResponse map[string]interface{}

// StatusQdevice Get QDevice status
func (c *Client) StatusQdevice(ctx context.Context) (*StatusQdeviceResponse, error) {
	var resp *StatusQdeviceResponse

	err := c.httpClient.Do(ctx, "/cluster/config/qdevice", "GET", &resp, nil)
	return resp, err
}

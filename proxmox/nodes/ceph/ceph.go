// File generated by proxmox json schema, DO NOT EDIT

package ceph

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
	Node string `url:"node",json:"node"` // The cluster node name.

}

type IndexResponse []*map[string]interface{}

// Index Directory index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph", "GET", &resp, req)
	return resp, err
}

type ConfigRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type ConfigResponse string

// Config Get Ceph configuration.
func (c *Client) Config(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	var resp *ConfigResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/config", "GET", &resp, req)
	return resp, err
}

type ConfigdbRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type ConfigdbResponse []*struct {
	CanUpdateAtRuntime util.SpecialBool `url:"can_update_at_runtime",json:"can_update_at_runtime"`
	Level              string           `url:"level",json:"level"`
	Mask               string           `url:"mask",json:"mask"`
	Name               string           `url:"name",json:"name"`
	Section            string           `url:"section",json:"section"`
	Value              string           `url:"value",json:"value"`
}

// Configdb Get Ceph configuration database.
func (c *Client) Configdb(ctx context.Context, req *ConfigdbRequest) (*ConfigdbResponse, error) {
	var resp *ConfigdbResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/configdb", "GET", &resp, req)
	return resp, err
}

type InitRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	ClusterNetwork *string           `url:"cluster-network,omitempty",json:"cluster-network,omitempty"` // Declare a separate cluster network, OSDs will routeheartbeat, object replication and recovery traffic over it
	DisableCephx   *util.SpecialBool `url:"disable_cephx,omitempty",json:"disable_cephx,omitempty"`     // Disable cephx authentication.WARNING: cephx is a security feature protecting against man-in-the-middle attacks. Only consider disabling cephx if your network is private!
	MinSize        *int              `url:"min_size,omitempty",json:"min_size,omitempty"`               // Minimum number of available replicas per object to allow I/O
	Network        *string           `url:"network,omitempty",json:"network,omitempty"`                 // Use specific network for all ceph related traffic
	PgBits         *int              `url:"pg_bits,omitempty",json:"pg_bits,omitempty"`                 // Placement group bits, used to specify the default number of placement groups.NOTE: 'osd pool default pg num' does not work for default pools.
	Size           *int              `url:"size,omitempty",json:"size,omitempty"`                       // Targeted number of replicas per object
}

type InitResponse map[string]interface{}

// Init Create initial ceph default configuration and setup symlinks.
func (c *Client) Init(ctx context.Context, req *InitRequest) (*InitResponse, error) {
	var resp *InitResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/init", "POST", &resp, req)
	return resp, err
}

type StopRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Service *string `url:"service,omitempty",json:"service,omitempty"` // Ceph service name.
}

type StopResponse string

// Stop Stop ceph services.
func (c *Client) Stop(ctx context.Context, req *StopRequest) (*StopResponse, error) {
	var resp *StopResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/stop", "POST", &resp, req)
	return resp, err
}

type StartRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Service *string `url:"service,omitempty",json:"service,omitempty"` // Ceph service name.
}

type StartResponse string

// Start Start ceph services.
func (c *Client) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	var resp *StartResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/start", "POST", &resp, req)
	return resp, err
}

type RestartRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Service *string `url:"service,omitempty",json:"service,omitempty"` // Ceph service name.
}

type RestartResponse string

// Restart Restart ceph services.
func (c *Client) Restart(ctx context.Context, req *RestartRequest) (*RestartResponse, error) {
	var resp *RestartResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/restart", "POST", &resp, req)
	return resp, err
}

type StatusRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type StatusResponse map[string]interface{}

// Status Get ceph status.
func (c *Client) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	var resp *StatusResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/status", "GET", &resp, req)
	return resp, err
}

type CrushRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type CrushResponse string

// Crush Get OSD crush map
func (c *Client) Crush(ctx context.Context, req *CrushRequest) (*CrushResponse, error) {
	var resp *CrushResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/crush", "GET", &resp, req)
	return resp, err
}

type LogRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Limit *int `url:"limit,omitempty",json:"limit,omitempty"`
	Start *int `url:"start,omitempty",json:"start,omitempty"`
}

type LogResponse []*struct {
	N int    `url:"n",json:"n"` // Line number
	T string `url:"t",json:"t"` // Line text

}

// Log Read ceph log
func (c *Client) Log(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	var resp *LogResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/log", "GET", &resp, req)
	return resp, err
}

type RulesRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type RulesResponse []*map[string]interface{}

// Rules List ceph rules.
func (c *Client) Rules(ctx context.Context, req *RulesRequest) (*RulesResponse, error) {
	var resp *RulesResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/rules", "GET", &resp, req)
	return resp, err
}

type CmdSafetyRequest struct {
	Action  string `url:"action",json:"action"`   // Action to check
	Id      string `url:"id",json:"id"`           // ID of the service
	Node    string `url:"node",json:"node"`       // The cluster node name.
	Service string `url:"service",json:"service"` // Service type

}

type CmdSafetyResponse struct {
	Safe util.SpecialBool `url:"safe",json:"safe"` // If it is safe to run the command.

	// The following parameters are optional
	Status *string `url:"status,omitempty",json:"status,omitempty"` // Status message given by Ceph.
}

// CmdSafety Heuristical check if it is safe to perform an action.
func (c *Client) CmdSafety(ctx context.Context, req *CmdSafetyRequest) (*CmdSafetyResponse, error) {
	var resp *CmdSafetyResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/cmd-safety", "GET", &resp, req)
	return resp, err
}

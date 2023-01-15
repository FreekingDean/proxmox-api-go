// File generated by proxmox json schema, DO NOT EDIT

package firewall

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
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type IndexResponse []*map[string]interface{}

// Index Directory index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall", "GET", &resp, req)
	return resp, err
}

type GetOptionsRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type GetOptionsResponse struct {

	// The following parameters are optional
	Dhcp        *util.SpecialBool `url:"dhcp,omitempty",json:"dhcp,omitempty"`                   // Enable DHCP.
	Enable      *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"`               // Enable/disable firewall rules.
	Ipfilter    *util.SpecialBool `url:"ipfilter,omitempty",json:"ipfilter,omitempty"`           // Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	LogLevelIn  *string           `url:"log_level_in,omitempty",json:"log_level_in,omitempty"`   // Log level for incoming traffic.
	LogLevelOut *string           `url:"log_level_out,omitempty",json:"log_level_out,omitempty"` // Log level for outgoing traffic.
	Macfilter   *util.SpecialBool `url:"macfilter,omitempty",json:"macfilter,omitempty"`         // Enable/disable MAC address filter.
	Ndp         *util.SpecialBool `url:"ndp,omitempty",json:"ndp,omitempty"`                     // Enable NDP (Neighbor Discovery Protocol).
	PolicyIn    *string           `url:"policy_in,omitempty",json:"policy_in,omitempty"`         // Input policy.
	PolicyOut   *string           `url:"policy_out,omitempty",json:"policy_out,omitempty"`       // Output policy.
	Radv        *util.SpecialBool `url:"radv,omitempty",json:"radv,omitempty"`                   // Allow sending Router Advertisement.
}

// GetOptions Get VM firewall options.
func (c *Client) GetOptions(ctx context.Context, req *GetOptionsRequest) (*GetOptionsResponse, error) {
	var resp *GetOptionsResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/options", "GET", &resp, req)
	return resp, err
}

type SetOptionsRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Delete      *string           `url:"delete,omitempty",json:"delete,omitempty"`               // A list of settings you want to delete.
	Dhcp        *util.SpecialBool `url:"dhcp,omitempty",json:"dhcp,omitempty"`                   // Enable DHCP.
	Digest      *string           `url:"digest,omitempty",json:"digest,omitempty"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Enable      *util.SpecialBool `url:"enable,omitempty",json:"enable,omitempty"`               // Enable/disable firewall rules.
	Ipfilter    *util.SpecialBool `url:"ipfilter,omitempty",json:"ipfilter,omitempty"`           // Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	LogLevelIn  *string           `url:"log_level_in,omitempty",json:"log_level_in,omitempty"`   // Log level for incoming traffic.
	LogLevelOut *string           `url:"log_level_out,omitempty",json:"log_level_out,omitempty"` // Log level for outgoing traffic.
	Macfilter   *util.SpecialBool `url:"macfilter,omitempty",json:"macfilter,omitempty"`         // Enable/disable MAC address filter.
	Ndp         *util.SpecialBool `url:"ndp,omitempty",json:"ndp,omitempty"`                     // Enable NDP (Neighbor Discovery Protocol).
	PolicyIn    *string           `url:"policy_in,omitempty",json:"policy_in,omitempty"`         // Input policy.
	PolicyOut   *string           `url:"policy_out,omitempty",json:"policy_out,omitempty"`       // Output policy.
	Radv        *util.SpecialBool `url:"radv,omitempty",json:"radv,omitempty"`                   // Allow sending Router Advertisement.
}

type SetOptionsResponse map[string]interface{}

// SetOptions Set Firewall options.
func (c *Client) SetOptions(ctx context.Context, req *SetOptionsRequest) (*SetOptionsResponse, error) {
	var resp *SetOptionsResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/options", "PUT", &resp, req)
	return resp, err
}

type LogRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Limit *int `url:"limit,omitempty",json:"limit,omitempty"`
	Start *int `url:"start,omitempty",json:"start,omitempty"`
}

type LogResponse []*struct {
	N int    `url:"n",json:"n"` // Line number
	T string `url:"t",json:"t"` // Line text

}

// Log Read firewall log
func (c *Client) Log(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	var resp *LogResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/log", "GET", &resp, req)
	return resp, err
}

type RefsRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Type *string `url:"type,omitempty",json:"type,omitempty"` // Only list references of specified type.
}

type RefsResponse []*struct {
	Name string `url:"name",json:"name"`
	Type string `url:"type",json:"type"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
}

// Refs Lists possible IPSet/Alias reference which are allowed in source/dest properties.
func (c *Client) Refs(ctx context.Context, req *RefsRequest) (*RefsResponse, error) {
	var resp *RefsResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/refs", "GET", &resp, req)
	return resp, err
}

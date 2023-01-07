// File generated by proxmox json schema, DO NOT EDIT

package rules

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

type IndexRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type IndexResponse []*struct {
	Pos int `url:"pos",json:"pos"`
}

// Index List rules.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/rules", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Action string `url:"action",json:"action"` // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Node   string `url:"node",json:"node"`     // The cluster node name.
	Type   string `url:"type",json:"type"`     // Rule type.
	Vmid   int    `url:"vmid",json:"vmid"`     // The (unique) ID of the VM.

	// The following parameters are optional
	Comment  *string `url:"comment,omitempty",json:"comment,omitempty"`     // Descriptive comment.
	Dest     *string `url:"dest,omitempty",json:"dest,omitempty"`           // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Digest   *string `url:"digest,omitempty",json:"digest,omitempty"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    *string `url:"dport,omitempty",json:"dport,omitempty"`         // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Enable   *int    `url:"enable,omitempty",json:"enable,omitempty"`       // Flag to enable/disable a rule.
	IcmpType *string `url:"icmp-type,omitempty",json:"icmp-type,omitempty"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Iface    *string `url:"iface,omitempty",json:"iface,omitempty"`         // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Log      *string `url:"log,omitempty",json:"log,omitempty"`             // Log level for firewall rule.
	Macro    *string `url:"macro,omitempty",json:"macro,omitempty"`         // Use predefined standard macro.
	Pos      *int    `url:"pos,omitempty",json:"pos,omitempty"`             // Update rule at position <pos>.
	Proto    *string `url:"proto,omitempty",json:"proto,omitempty"`         // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source   *string `url:"source,omitempty",json:"source,omitempty"`       // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport    *string `url:"sport,omitempty",json:"sport,omitempty"`         // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
}

type CreateResponse map[string]interface{}

// Create Create new rule.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/rules", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Pos *int `url:"pos,omitempty",json:"pos,omitempty"` // Update rule at position <pos>.
}

type FindResponse struct {
	Action string `url:"action",json:"action"`
	Pos    int    `url:"pos",json:"pos"`
	Type   string `url:"type",json:"type"`

	// The following parameters are optional
	Comment   *string `url:"comment,omitempty",json:"comment,omitempty"`
	Dest      *string `url:"dest,omitempty",json:"dest,omitempty"`
	Dport     *string `url:"dport,omitempty",json:"dport,omitempty"`
	Enable    *int    `url:"enable,omitempty",json:"enable,omitempty"`
	IcmpType  *string `url:"icmp-type,omitempty",json:"icmp-type,omitempty"`
	Iface     *string `url:"iface,omitempty",json:"iface,omitempty"`
	Ipversion *int    `url:"ipversion,omitempty",json:"ipversion,omitempty"`
	Log       *string `url:"log,omitempty",json:"log,omitempty"` // Log level for firewall rule
	Macro     *string `url:"macro,omitempty",json:"macro,omitempty"`
	Proto     *string `url:"proto,omitempty",json:"proto,omitempty"`
	Source    *string `url:"source,omitempty",json:"source,omitempty"`
	Sport     *string `url:"sport,omitempty",json:"sport,omitempty"`
}

// Find Get single rule data.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Action   *string `url:"action,omitempty",json:"action,omitempty"`       // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Comment  *string `url:"comment,omitempty",json:"comment,omitempty"`     // Descriptive comment.
	Delete   *string `url:"delete,omitempty",json:"delete,omitempty"`       // A list of settings you want to delete.
	Dest     *string `url:"dest,omitempty",json:"dest,omitempty"`           // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Digest   *string `url:"digest,omitempty",json:"digest,omitempty"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    *string `url:"dport,omitempty",json:"dport,omitempty"`         // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Enable   *int    `url:"enable,omitempty",json:"enable,omitempty"`       // Flag to enable/disable a rule.
	IcmpType *string `url:"icmp-type,omitempty",json:"icmp-type,omitempty"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Iface    *string `url:"iface,omitempty",json:"iface,omitempty"`         // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Log      *string `url:"log,omitempty",json:"log,omitempty"`             // Log level for firewall rule.
	Macro    *string `url:"macro,omitempty",json:"macro,omitempty"`         // Use predefined standard macro.
	Moveto   *int    `url:"moveto,omitempty",json:"moveto,omitempty"`       // Move rule to new position <moveto>. Other arguments are ignored.
	Pos      *int    `url:"pos,omitempty",json:"pos,omitempty"`             // Update rule at position <pos>.
	Proto    *string `url:"proto,omitempty",json:"proto,omitempty"`         // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source   *string `url:"source,omitempty",json:"source,omitempty"`       // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport    *string `url:"sport,omitempty",json:"sport,omitempty"`         // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Type     *string `url:"type,omitempty",json:"type,omitempty"`           // Rule type.
}

type UpdateResponse map[string]interface{}

// Update Modify rule data.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty",json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Pos    *int    `url:"pos,omitempty",json:"pos,omitempty"`       // Update rule at position <pos>.
}

type DeleteResponse map[string]interface{}

// Delete Delete rule.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}", "DELETE", &resp, req)
	return resp, err
}

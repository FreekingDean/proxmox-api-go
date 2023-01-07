// File generated by proxmox json schema, DO NOT EDIT

package zones

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

type IndexRequest map[string]interface{}

type IndexResponse []*struct {
	Type string `url:"type",json:"type"`
	Zone string `url:"zone",json:"zone"`

	// The following parameters are optional
	Dns        *string `url:"dns,omitempty",json:"dns,omitempty"`
	Dnszone    *string `url:"dnszone,omitempty",json:"dnszone,omitempty"`
	Ipam       *string `url:"ipam,omitempty",json:"ipam,omitempty"`
	Mtu        *int    `url:"mtu,omitempty",json:"mtu,omitempty"`
	Nodes      *string `url:"nodes,omitempty",json:"nodes,omitempty"`
	Reversedns *string `url:"reversedns,omitempty",json:"reversedns,omitempty"`
	State      *string `url:"state,omitempty",json:"state,omitempty"`
}

// Index SDN zones index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/zones", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Type string `url:"type",json:"type"` // Plugin type.
	Zone string `url:"zone",json:"zone"` // The SDN zone object identifier.

	// The following parameters are optional
	AdvertiseSubnets         *bool   `url:"advertise-subnets,omitempty",json:"advertise-subnets,omitempty"` // Advertise evpn subnets if you have silent hosts
	Bridge                   *string `url:"bridge,omitempty",json:"bridge,omitempty"`
	BridgeDisableMacLearning *bool   `url:"bridge-disable-mac-learning,omitempty",json:"bridge-disable-mac-learning,omitempty"` // Disable auto mac learning.
	Controller               *string `url:"controller,omitempty",json:"controller,omitempty"`                                   // Frr router name
	DisableArpNdSuppression  *bool   `url:"disable-arp-nd-suppression,omitempty",json:"disable-arp-nd-suppression,omitempty"`   // Disable ipv4 arp && ipv6 neighbour discovery suppression
	Dns                      *string `url:"dns,omitempty",json:"dns,omitempty"`                                                 // dns api server
	Dnszone                  *string `url:"dnszone,omitempty",json:"dnszone,omitempty"`                                         // dns domain zone  ex: mydomain.com
	DpId                     *int    `url:"dp-id,omitempty",json:"dp-id,omitempty"`                                             // Faucet dataplane id
	Exitnodes                *string `url:"exitnodes,omitempty",json:"exitnodes,omitempty"`                                     // List of cluster node names.
	ExitnodesLocalRouting    *bool   `url:"exitnodes-local-routing,omitempty",json:"exitnodes-local-routing,omitempty"`         // Allow exitnodes to connect to evpn guests
	ExitnodesPrimary         *string `url:"exitnodes-primary,omitempty",json:"exitnodes-primary,omitempty"`                     // Force traffic to this exitnode first.
	Ipam                     *string `url:"ipam,omitempty",json:"ipam,omitempty"`                                               // use a specific ipam
	Mac                      *string `url:"mac,omitempty",json:"mac,omitempty"`                                                 // Anycast logical router mac address
	Mtu                      *int    `url:"mtu,omitempty",json:"mtu,omitempty"`                                                 // MTU
	Nodes                    *string `url:"nodes,omitempty",json:"nodes,omitempty"`                                             // List of cluster node names.
	Peers                    *string `url:"peers,omitempty",json:"peers,omitempty"`                                             // peers address list.
	Reversedns               *string `url:"reversedns,omitempty",json:"reversedns,omitempty"`                                   // reverse dns api server
	RtImport                 *string `url:"rt-import,omitempty",json:"rt-import,omitempty"`                                     // Route-Target import
	Tag                      *int    `url:"tag,omitempty",json:"tag,omitempty"`                                                 // Service-VLAN Tag
	VlanProtocol             *string `url:"vlan-protocol,omitempty",json:"vlan-protocol,omitempty"`
	VrfVxlan                 *int    `url:"vrf-vxlan,omitempty",json:"vrf-vxlan,omitempty"` // l3vni.
}

type CreateResponse map[string]interface{}

// Create Create a new sdn zone object.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/zones", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Zone string `url:"zone",json:"zone"` // The SDN zone object identifier.

	// The following parameters are optional
	Pending *bool `url:"pending,omitempty",json:"pending,omitempty"` // Display pending config.
	Running *bool `url:"running,omitempty",json:"running,omitempty"` // Display running config.
}

type FindResponse map[string]interface{}

// Find Read sdn zone configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/zones/{zone}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Zone string `url:"zone",json:"zone"` // The SDN zone object identifier.

	// The following parameters are optional
	AdvertiseSubnets         *bool   `url:"advertise-subnets,omitempty",json:"advertise-subnets,omitempty"` // Advertise evpn subnets if you have silent hosts
	Bridge                   *string `url:"bridge,omitempty",json:"bridge,omitempty"`
	BridgeDisableMacLearning *bool   `url:"bridge-disable-mac-learning,omitempty",json:"bridge-disable-mac-learning,omitempty"` // Disable auto mac learning.
	Controller               *string `url:"controller,omitempty",json:"controller,omitempty"`                                   // Frr router name
	Delete                   *string `url:"delete,omitempty",json:"delete,omitempty"`                                           // A list of settings you want to delete.
	Digest                   *string `url:"digest,omitempty",json:"digest,omitempty"`                                           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	DisableArpNdSuppression  *bool   `url:"disable-arp-nd-suppression,omitempty",json:"disable-arp-nd-suppression,omitempty"`   // Disable ipv4 arp && ipv6 neighbour discovery suppression
	Dns                      *string `url:"dns,omitempty",json:"dns,omitempty"`                                                 // dns api server
	Dnszone                  *string `url:"dnszone,omitempty",json:"dnszone,omitempty"`                                         // dns domain zone  ex: mydomain.com
	DpId                     *int    `url:"dp-id,omitempty",json:"dp-id,omitempty"`                                             // Faucet dataplane id
	Exitnodes                *string `url:"exitnodes,omitempty",json:"exitnodes,omitempty"`                                     // List of cluster node names.
	ExitnodesLocalRouting    *bool   `url:"exitnodes-local-routing,omitempty",json:"exitnodes-local-routing,omitempty"`         // Allow exitnodes to connect to evpn guests
	ExitnodesPrimary         *string `url:"exitnodes-primary,omitempty",json:"exitnodes-primary,omitempty"`                     // Force traffic to this exitnode first.
	Ipam                     *string `url:"ipam,omitempty",json:"ipam,omitempty"`                                               // use a specific ipam
	Mac                      *string `url:"mac,omitempty",json:"mac,omitempty"`                                                 // Anycast logical router mac address
	Mtu                      *int    `url:"mtu,omitempty",json:"mtu,omitempty"`                                                 // MTU
	Nodes                    *string `url:"nodes,omitempty",json:"nodes,omitempty"`                                             // List of cluster node names.
	Peers                    *string `url:"peers,omitempty",json:"peers,omitempty"`                                             // peers address list.
	Reversedns               *string `url:"reversedns,omitempty",json:"reversedns,omitempty"`                                   // reverse dns api server
	RtImport                 *string `url:"rt-import,omitempty",json:"rt-import,omitempty"`                                     // Route-Target import
	Tag                      *int    `url:"tag,omitempty",json:"tag,omitempty"`                                                 // Service-VLAN Tag
	VlanProtocol             *string `url:"vlan-protocol,omitempty",json:"vlan-protocol,omitempty"`
	VrfVxlan                 *int    `url:"vrf-vxlan,omitempty",json:"vrf-vxlan,omitempty"` // l3vni.
}

type UpdateResponse map[string]interface{}

// Update Update sdn zone object configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/zones/{zone}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Zone string `url:"zone",json:"zone"` // The SDN zone object identifier.

}

type DeleteResponse map[string]interface{}

// Delete Delete sdn zone object configuration.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/zones/{zone}", "DELETE", &resp, req)
	return resp, err
}
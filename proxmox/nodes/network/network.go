// File generated by proxmox json schema, DO NOT EDIT

package network

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

	// The following parameters are optional
	Type *string `url:"type,omitempty",json:"type,omitempty"` // Only list specific interface types.
}

type IndexResponse []*map[string]interface{}

// Index List available networks
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Iface string `url:"iface",json:"iface"` // Network interface name.
	Node  string `url:"node",json:"node"`   // The cluster node name.
	Type  string `url:"type",json:"type"`   // Network interface type

	// The following parameters are optional
	Address            *string `url:"address,omitempty",json:"address,omitempty"`                             // IP address.
	Address6           *string `url:"address6,omitempty",json:"address6,omitempty"`                           // IP address.
	Autostart          *bool   `url:"autostart,omitempty",json:"autostart,omitempty"`                         // Automatically start interface on boot.
	BondMode           *string `url:"bond_mode,omitempty",json:"bond_mode,omitempty"`                         // Bonding mode.
	BondPrimary        *string `url:"bond-primary,omitempty",json:"bond-primary,omitempty"`                   // Specify the primary interface for active-backup bond.
	BondXmitHashPolicy *string `url:"bond_xmit_hash_policy,omitempty",json:"bond_xmit_hash_policy,omitempty"` // Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	BridgePorts        *string `url:"bridge_ports,omitempty",json:"bridge_ports,omitempty"`                   // Specify the interfaces you want to add to your bridge.
	BridgeVlanAware    *bool   `url:"bridge_vlan_aware,omitempty",json:"bridge_vlan_aware,omitempty"`         // Enable bridge vlan support.
	Cidr               *string `url:"cidr,omitempty",json:"cidr,omitempty"`                                   // IPv4 CIDR.
	Cidr6              *string `url:"cidr6,omitempty",json:"cidr6,omitempty"`                                 // IPv6 CIDR.
	Comments           *string `url:"comments,omitempty",json:"comments,omitempty"`                           // Comments
	Comments6          *string `url:"comments6,omitempty",json:"comments6,omitempty"`                         // Comments
	Gateway            *string `url:"gateway,omitempty",json:"gateway,omitempty"`                             // Default gateway address.
	Gateway6           *string `url:"gateway6,omitempty",json:"gateway6,omitempty"`                           // Default ipv6 gateway address.
	Mtu                *int    `url:"mtu,omitempty",json:"mtu,omitempty"`                                     // MTU.
	Netmask            *string `url:"netmask,omitempty",json:"netmask,omitempty"`                             // Network mask.
	Netmask6           *int    `url:"netmask6,omitempty",json:"netmask6,omitempty"`                           // Network mask.
	OvsBonds           *string `url:"ovs_bonds,omitempty",json:"ovs_bonds,omitempty"`                         // Specify the interfaces used by the bonding device.
	OvsBridge          *string `url:"ovs_bridge,omitempty",json:"ovs_bridge,omitempty"`                       // The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	OvsOptions         *string `url:"ovs_options,omitempty",json:"ovs_options,omitempty"`                     // OVS interface options.
	OvsPorts           *string `url:"ovs_ports,omitempty",json:"ovs_ports,omitempty"`                         // Specify the interfaces you want to add to your bridge.
	OvsTag             *int    `url:"ovs_tag,omitempty",json:"ovs_tag,omitempty"`                             // Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	Slaves             *string `url:"slaves,omitempty",json:"slaves,omitempty"`                               // Specify the interfaces used by the bonding device.
	VlanId             *int    `url:"vlan-id,omitempty",json:"vlan-id,omitempty"`                             // vlan-id for a custom named vlan interface (ifupdown2 only).
	VlanRawDevice      *string `url:"vlan-raw-device,omitempty",json:"vlan-raw-device,omitempty"`             // Specify the raw interface for the vlan interface.
}

type CreateResponse map[string]interface{}

// Create Create network device configuration
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network", "POST", &resp, req)
	return resp, err
}

type MassUpdateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type MassUpdateResponse string

// MassUpdate Reload network configuration
func (c *Client) MassUpdate(ctx context.Context, req *MassUpdateRequest) (*MassUpdateResponse, error) {
	var resp *MassUpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network", "PUT", &resp, req)
	return resp, err
}

type MassDeleteRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type MassDeleteResponse map[string]interface{}

// MassDelete Revert network configuration changes.
func (c *Client) MassDelete(ctx context.Context, req *MassDeleteRequest) (*MassDeleteResponse, error) {
	var resp *MassDeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network", "DELETE", &resp, req)
	return resp, err
}

type FindRequest struct {
	Iface string `url:"iface",json:"iface"` // Network interface name.
	Node  string `url:"node",json:"node"`   // The cluster node name.

}

type FindResponse struct {
	Method string `url:"method",json:"method"`
	Type   string `url:"type",json:"type"`
}

// Find Read network device configuration
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network/{iface}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Iface string `url:"iface",json:"iface"` // Network interface name.
	Node  string `url:"node",json:"node"`   // The cluster node name.
	Type  string `url:"type",json:"type"`   // Network interface type

	// The following parameters are optional
	Address            *string `url:"address,omitempty",json:"address,omitempty"`                             // IP address.
	Address6           *string `url:"address6,omitempty",json:"address6,omitempty"`                           // IP address.
	Autostart          *bool   `url:"autostart,omitempty",json:"autostart,omitempty"`                         // Automatically start interface on boot.
	BondMode           *string `url:"bond_mode,omitempty",json:"bond_mode,omitempty"`                         // Bonding mode.
	BondPrimary        *string `url:"bond-primary,omitempty",json:"bond-primary,omitempty"`                   // Specify the primary interface for active-backup bond.
	BondXmitHashPolicy *string `url:"bond_xmit_hash_policy,omitempty",json:"bond_xmit_hash_policy,omitempty"` // Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	BridgePorts        *string `url:"bridge_ports,omitempty",json:"bridge_ports,omitempty"`                   // Specify the interfaces you want to add to your bridge.
	BridgeVlanAware    *bool   `url:"bridge_vlan_aware,omitempty",json:"bridge_vlan_aware,omitempty"`         // Enable bridge vlan support.
	Cidr               *string `url:"cidr,omitempty",json:"cidr,omitempty"`                                   // IPv4 CIDR.
	Cidr6              *string `url:"cidr6,omitempty",json:"cidr6,omitempty"`                                 // IPv6 CIDR.
	Comments           *string `url:"comments,omitempty",json:"comments,omitempty"`                           // Comments
	Comments6          *string `url:"comments6,omitempty",json:"comments6,omitempty"`                         // Comments
	Delete             *string `url:"delete,omitempty",json:"delete,omitempty"`                               // A list of settings you want to delete.
	Gateway            *string `url:"gateway,omitempty",json:"gateway,omitempty"`                             // Default gateway address.
	Gateway6           *string `url:"gateway6,omitempty",json:"gateway6,omitempty"`                           // Default ipv6 gateway address.
	Mtu                *int    `url:"mtu,omitempty",json:"mtu,omitempty"`                                     // MTU.
	Netmask            *string `url:"netmask,omitempty",json:"netmask,omitempty"`                             // Network mask.
	Netmask6           *int    `url:"netmask6,omitempty",json:"netmask6,omitempty"`                           // Network mask.
	OvsBonds           *string `url:"ovs_bonds,omitempty",json:"ovs_bonds,omitempty"`                         // Specify the interfaces used by the bonding device.
	OvsBridge          *string `url:"ovs_bridge,omitempty",json:"ovs_bridge,omitempty"`                       // The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	OvsOptions         *string `url:"ovs_options,omitempty",json:"ovs_options,omitempty"`                     // OVS interface options.
	OvsPorts           *string `url:"ovs_ports,omitempty",json:"ovs_ports,omitempty"`                         // Specify the interfaces you want to add to your bridge.
	OvsTag             *int    `url:"ovs_tag,omitempty",json:"ovs_tag,omitempty"`                             // Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	Slaves             *string `url:"slaves,omitempty",json:"slaves,omitempty"`                               // Specify the interfaces used by the bonding device.
	VlanId             *int    `url:"vlan-id,omitempty",json:"vlan-id,omitempty"`                             // vlan-id for a custom named vlan interface (ifupdown2 only).
	VlanRawDevice      *string `url:"vlan-raw-device,omitempty",json:"vlan-raw-device,omitempty"`             // Specify the raw interface for the vlan interface.
}

type UpdateResponse map[string]interface{}

// Update Update network device configuration
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network/{iface}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Iface string `url:"iface",json:"iface"` // Network interface name.
	Node  string `url:"node",json:"node"`   // The cluster node name.

}

type DeleteResponse map[string]interface{}

// Delete Delete network device configuration
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/network/{iface}", "DELETE", &resp, req)
	return resp, err
}

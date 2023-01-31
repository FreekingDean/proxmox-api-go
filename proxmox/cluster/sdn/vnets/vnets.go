// File generated by proxmox json schema, DO NOT EDIT

package vnets

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

	// The following parameters are optional
	Pending *util.PVEBool `url:"pending,omitempty" json:"pending,omitempty"` // Display pending config.
	Running *util.PVEBool `url:"running,omitempty" json:"running,omitempty"` // Display running config.
}

type CreateRequest struct {
	Vnet string `url:"vnet" json:"vnet"` // The SDN vnet object identifier.
	Zone string `url:"zone" json:"zone"` // zone id

	// The following parameters are optional
	Alias     *string       `url:"alias,omitempty" json:"alias,omitempty"`         // alias name of the vnet
	Tag       *int          `url:"tag,omitempty" json:"tag,omitempty"`             // vlan or vxlan id
	Type      *string       `url:"type,omitempty" json:"type,omitempty"`           // Type
	Vlanaware *util.PVEBool `url:"vlanaware,omitempty" json:"vlanaware,omitempty"` // Allow vm VLANs to pass through this vnet.
}

type FindRequest struct {
	Vnet string `url:"vnet" json:"vnet"` // The SDN vnet object identifier.

	// The following parameters are optional
	Pending *util.PVEBool `url:"pending,omitempty" json:"pending,omitempty"` // Display pending config.
	Running *util.PVEBool `url:"running,omitempty" json:"running,omitempty"` // Display running config.
}

type UpdateRequest struct {
	Vnet string `url:"vnet" json:"vnet"` // The SDN vnet object identifier.

	// The following parameters are optional
	Alias     *string       `url:"alias,omitempty" json:"alias,omitempty"`         // alias name of the vnet
	Delete    *string       `url:"delete,omitempty" json:"delete,omitempty"`       // A list of settings you want to delete.
	Digest    *string       `url:"digest,omitempty" json:"digest,omitempty"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Tag       *int          `url:"tag,omitempty" json:"tag,omitempty"`             // vlan or vxlan id
	Vlanaware *util.PVEBool `url:"vlanaware,omitempty" json:"vlanaware,omitempty"` // Allow vm VLANs to pass through this vnet.
	Zone      *string       `url:"zone,omitempty" json:"zone,omitempty"`           // zone id
}

type DeleteRequest struct {
	Vnet string `url:"vnet" json:"vnet"` // The SDN vnet object identifier.

}

// Index SDN vnets index.
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets", "GET", &resp, req)
	return resp, err
}

// Create Create a new sdn vnet object.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets", "POST", nil, req)
	return err
}

// Find Read sdn vnet configuration.
func (c *Client) Find(ctx context.Context, req FindRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}", "GET", &resp, req)
	return resp, err
}

// Update Update sdn vnet object configuration.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}", "PUT", nil, req)
	return err
}

// Delete Delete sdn vnet object configuration.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}", "DELETE", nil, req)
	return err
}

// File generated by proxmox json schema, DO NOT EDIT

package subnets

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
	Vnet string `url:"vnet" json:"vnet"` // The SDN vnet object identifier.

	// The following parameters are optional
	Pending *util.SpecialBool `url:"pending,omitempty" json:"pending,omitempty"` // Display pending config.
	Running *util.SpecialBool `url:"running,omitempty" json:"running,omitempty"` // Display running config.
}

type IndexResponse []*map[string]interface{}

// Index SDN subnets index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}/subnets", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Subnet string `url:"subnet" json:"subnet"` // The SDN subnet object identifier.
	Type   string `url:"type" json:"type"`
	Vnet   string `url:"vnet" json:"vnet"` // associated vnet

	// The following parameters are optional
	Dnszoneprefix *string           `url:"dnszoneprefix,omitempty" json:"dnszoneprefix,omitempty"` // dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Gateway       *string           `url:"gateway,omitempty" json:"gateway,omitempty"`             // Subnet Gateway: Will be assign on vnet for layer3 zones
	Snat          *util.SpecialBool `url:"snat,omitempty" json:"snat,omitempty"`                   // enable masquerade for this subnet if pve-firewall
}

type CreateResponse map[string]interface{}

// Create Create a new sdn subnet object.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}/subnets", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Subnet string `url:"subnet" json:"subnet"` // The SDN subnet object identifier.
	Vnet   string `url:"vnet" json:"vnet"`     // The SDN vnet object identifier.

	// The following parameters are optional
	Pending *util.SpecialBool `url:"pending,omitempty" json:"pending,omitempty"` // Display pending config.
	Running *util.SpecialBool `url:"running,omitempty" json:"running,omitempty"` // Display running config.
}

type FindResponse map[string]interface{}

// Find Read sdn subnet configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}/subnets/{subnet}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Subnet string `url:"subnet" json:"subnet"` // The SDN subnet object identifier.

	// The following parameters are optional
	Delete        *string           `url:"delete,omitempty" json:"delete,omitempty"`               // A list of settings you want to delete.
	Digest        *string           `url:"digest,omitempty" json:"digest,omitempty"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dnszoneprefix *string           `url:"dnszoneprefix,omitempty" json:"dnszoneprefix,omitempty"` // dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Gateway       *string           `url:"gateway,omitempty" json:"gateway,omitempty"`             // Subnet Gateway: Will be assign on vnet for layer3 zones
	Snat          *util.SpecialBool `url:"snat,omitempty" json:"snat,omitempty"`                   // enable masquerade for this subnet if pve-firewall
	Vnet          *string           `url:"vnet,omitempty" json:"vnet,omitempty"`                   // associated vnet
}

type UpdateResponse map[string]interface{}

// Update Update sdn subnet object configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}/subnets/{subnet}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Subnet string `url:"subnet" json:"subnet"` // The SDN subnet object identifier.
	Vnet   string `url:"vnet" json:"vnet"`     // The SDN vnet object identifier.

}

type DeleteResponse map[string]interface{}

// Delete Delete sdn subnet object configuration.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/sdn/vnets/{vnet}/subnets/{subnet}", "DELETE", &resp, req)
	return resp, err
}

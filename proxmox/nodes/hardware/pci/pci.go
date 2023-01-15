// File generated by proxmox json schema, DO NOT EDIT

package pci

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

	// The following parameters are optional
	PciClassBlacklist *string           `url:"pci-class-blacklist,omitempty",json:"pci-class-blacklist,omitempty"` // A list of blacklisted PCI classes, which will not be returned. Following are filtered by default: Memory Controller (05), Bridge (06) and Processor (0b).
	Verbose           *util.SpecialBool `url:"verbose,omitempty",json:"verbose,omitempty"`                         // If disabled, does only print the PCI IDs. Otherwise, additional information like vendor and device will be returned.
}

type IndexResponse []*struct {
	Class      string `url:"class",json:"class"`           // The PCI Class of the device.
	Device     string `url:"device",json:"device"`         // The Device ID.
	Id         string `url:"id",json:"id"`                 // The PCI ID.
	Iommugroup int    `url:"iommugroup",json:"iommugroup"` // The IOMMU group in which the device is in. If no IOMMU group is detected, it is set to -1.
	Vendor     string `url:"vendor",json:"vendor"`         // The Vendor ID.

	// The following parameters are optional
	DeviceName          *string           `url:"device_name,omitempty",json:"device_name,omitempty"`
	Mdev                *util.SpecialBool `url:"mdev,omitempty",json:"mdev,omitempty"`                         // If set, marks that the device is capable of creating mediated devices.
	SubsystemDevice     *string           `url:"subsystem_device,omitempty",json:"subsystem_device,omitempty"` // The Subsystem Device ID.
	SubsystemDeviceName *string           `url:"subsystem_device_name,omitempty",json:"subsystem_device_name,omitempty"`
	SubsystemVendor     *string           `url:"subsystem_vendor,omitempty",json:"subsystem_vendor,omitempty"` // The Subsystem Vendor ID.
	SubsystemVendorName *string           `url:"subsystem_vendor_name,omitempty",json:"subsystem_vendor_name,omitempty"`
	VendorName          *string           `url:"vendor_name,omitempty",json:"vendor_name,omitempty"`
}

// Index List local PCI devices.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/hardware/pci", "GET", &resp, req)
	return resp, err
}

type FindRequest struct {
	Node  string `url:"node",json:"node"` // The cluster node name.
	Pciid string `url:"pciid",json:"pciid"`
}

type FindResponse []*struct {
	Method string `url:"method",json:"method"`
}

// Find Index of available pci methods
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/hardware/pci/{pciid}", "GET", &resp, req)
	return resp, err
}

type MdevscanMdevRequest struct {
	Node  string `url:"node",json:"node"`   // The cluster node name.
	Pciid string `url:"pciid",json:"pciid"` // The PCI ID to list the mdev types for.

}

type MdevscanMdevResponse []*struct {
	Available   int    `url:"available",json:"available"` // The number of still available instances of this type.
	Description string `url:"description",json:"description"`
	Type        string `url:"type",json:"type"` // The name of the mdev type.

}

// MdevscanMdev List mediated device types for given PCI device.
func (c *Client) MdevscanMdev(ctx context.Context, req *MdevscanMdevRequest) (*MdevscanMdevResponse, error) {
	var resp *MdevscanMdevResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/hardware/pci/{pciid}/mdev", "GET", &resp, req)
	return resp, err
}

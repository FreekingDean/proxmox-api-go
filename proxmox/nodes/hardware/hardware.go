// File generated by proxmox json schema, DO NOT EDIT

package hardware

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
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _IndexRequest IndexRequest

type IndexResponse struct {
	Type string `url:"type" json:"type"`
}
type _IndexResponse IndexResponse

type UsbscanRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _UsbscanRequest UsbscanRequest

type UsbscanResponse struct {
	Busnum int    `url:"busnum" json:"busnum"`
	Class  int    `url:"class" json:"class"`
	Devnum int    `url:"devnum" json:"devnum"`
	Level  int    `url:"level" json:"level"`
	Port   int    `url:"port" json:"port"`
	Prodid string `url:"prodid" json:"prodid"`
	Speed  string `url:"speed" json:"speed"`
	Vendid string `url:"vendid" json:"vendid"`

	// The following parameters are optional
	Manufacturer *string `url:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Product      *string `url:"product,omitempty" json:"product,omitempty"`
	Serial       *string `url:"serial,omitempty" json:"serial,omitempty"`
	Usbpath      *string `url:"usbpath,omitempty" json:"usbpath,omitempty"`
}
type _UsbscanResponse UsbscanResponse

// Index Index of hardware types
func (c *Client) Index(ctx context.Context, req IndexRequest) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/hardware", "GET", &resp, req)
	return resp, err
}

// Usbscan List local USB devices.
func (c *Client) Usbscan(ctx context.Context, req UsbscanRequest) ([]UsbscanResponse, error) {
	var resp []UsbscanResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/hardware/usb", "GET", &resp, req)
	return resp, err
}

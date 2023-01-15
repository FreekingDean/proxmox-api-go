// File generated by proxmox json schema, DO NOT EDIT

package acme

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

// Index ACME index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/certificates/acme", "GET", &resp, req)
	return resp, err
}

type NewCertificateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Force *util.SpecialBool `url:"force,omitempty",json:"force,omitempty"` // Overwrite existing custom certificate.
}

type NewCertificateResponse string

// NewCertificate Order a new certificate from ACME-compatible CA.
func (c *Client) NewCertificate(ctx context.Context, req *NewCertificateRequest) (*NewCertificateResponse, error) {
	var resp *NewCertificateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/certificates/acme/certificate", "POST", &resp, req)
	return resp, err
}

type RenewCertificateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Force *util.SpecialBool `url:"force,omitempty",json:"force,omitempty"` // Force renewal even if expiry is more than 30 days away.
}

type RenewCertificateResponse string

// RenewCertificate Renew existing certificate from CA.
func (c *Client) RenewCertificate(ctx context.Context, req *RenewCertificateRequest) (*RenewCertificateResponse, error) {
	var resp *RenewCertificateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/certificates/acme/certificate", "PUT", &resp, req)
	return resp, err
}

type RevokeCertificateRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

}

type RevokeCertificateResponse string

// RevokeCertificate Revoke existing certificate from CA.
func (c *Client) RevokeCertificate(ctx context.Context, req *RevokeCertificateRequest) (*RevokeCertificateResponse, error) {
	var resp *RevokeCertificateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/certificates/acme/certificate", "DELETE", &resp, req)
	return resp, err
}

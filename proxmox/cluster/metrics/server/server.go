// File generated by proxmox json schema, DO NOT EDIT

package server

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

type IndexResponse []*struct {
	Type    string `url:"type",json:"type"`
	Disable bool   `url:"disable",json:"disable"`
	Id      string `url:"id",json:"id"`
	Port    int    `url:"port",json:"port"`
	Server  string `url:"server",json:"server"`
}

// Index List configured metric servers.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/metrics/server", "GET", &resp, nil)
	return resp, err
}

type FindRequest struct {
	Id string `url:"id",json:"id"`
}

type FindResponse map[string]interface{}

// Find Read metric server configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/cluster/metrics/server/{id}", "GET", &resp, req)
	return resp, err
}

type ChildCreateRequest struct {
	Id                string  `url:"id",json:"id"`
	Influxdbproto     *string `url:"influxdbproto,omitempty",json:"influxdbproto,omitempty"`
	Type              string  `url:"type",json:"type"`
	Server            string  `url:"server",json:"server"`
	Token             *string `url:"token,omitempty",json:"token,omitempty"`
	ApiPathPrefix     *string `url:"api-path-prefix,omitempty",json:"api-path-prefix,omitempty"`
	Bucket            *string `url:"bucket,omitempty",json:"bucket,omitempty"`
	Proto             *string `url:"proto,omitempty",json:"proto,omitempty"`
	Port              int     `url:"port",json:"port"`
	Timeout           *int    `url:"timeout,omitempty",json:"timeout,omitempty"`
	VerifyCertificate *bool   `url:"verify-certificate,omitempty",json:"verify-certificate,omitempty"`
	MaxBodySize       *int    `url:"max-body-size,omitempty",json:"max-body-size,omitempty"`
	Organization      *string `url:"organization,omitempty",json:"organization,omitempty"`
	Path              *string `url:"path,omitempty",json:"path,omitempty"`
	Disable           *bool   `url:"disable,omitempty",json:"disable,omitempty"`
	Mtu               *int    `url:"mtu,omitempty",json:"mtu,omitempty"`
}

type ChildCreateResponse map[string]interface{}

// ChildCreate Create a new external metric server config
func (c *Client) ChildCreate(ctx context.Context, req *ChildCreateRequest) (*ChildCreateResponse, error) {
	var resp *ChildCreateResponse

	err := c.httpClient.Do(ctx, "/cluster/metrics/server/{id}", "POST", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Organization      *string `url:"organization,omitempty",json:"organization,omitempty"`
	Port              int     `url:"port",json:"port"`
	VerifyCertificate *bool   `url:"verify-certificate,omitempty",json:"verify-certificate,omitempty"`
	ApiPathPrefix     *string `url:"api-path-prefix,omitempty",json:"api-path-prefix,omitempty"`
	Delete            *string `url:"delete,omitempty",json:"delete,omitempty"`
	Digest            *string `url:"digest,omitempty",json:"digest,omitempty"`
	Id                string  `url:"id",json:"id"`
	Timeout           *int    `url:"timeout,omitempty",json:"timeout,omitempty"`
	Bucket            *string `url:"bucket,omitempty",json:"bucket,omitempty"`
	MaxBodySize       *int    `url:"max-body-size,omitempty",json:"max-body-size,omitempty"`
	Mtu               *int    `url:"mtu,omitempty",json:"mtu,omitempty"`
	Path              *string `url:"path,omitempty",json:"path,omitempty"`
	Proto             *string `url:"proto,omitempty",json:"proto,omitempty"`
	Influxdbproto     *string `url:"influxdbproto,omitempty",json:"influxdbproto,omitempty"`
	Server            string  `url:"server",json:"server"`
	Token             *string `url:"token,omitempty",json:"token,omitempty"`
	Disable           *bool   `url:"disable,omitempty",json:"disable,omitempty"`
}

type UpdateResponse map[string]interface{}

// Update Update metric server configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/cluster/metrics/server/{id}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Id string `url:"id",json:"id"`
}

type DeleteResponse map[string]interface{}

// Delete Remove Metric server.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/cluster/metrics/server/{id}", "DELETE", &resp, req)
	return resp, err
}

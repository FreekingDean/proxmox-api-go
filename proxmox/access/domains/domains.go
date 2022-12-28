// File generated by proxmox json schema, DO NOT EDIT

package domains

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
	Comment *string `url:"comment,omitempty",json:"comment,omitempty"`
	Realm   string  `url:"realm",json:"realm"`
	Tfa     *string `url:"tfa,omitempty",json:"tfa,omitempty"`
	Type    string  `url:"type",json:"type"`
}

// Index Authentication domain index.
func (c *Client) Index(ctx context.Context) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/access/domains", "GET", &resp, nil)
	return resp, err
}

type CreateRequest struct {
	Mode                *string `url:"mode,omitempty",json:"mode,omitempty"`
	Password            *string `url:"password,omitempty",json:"password,omitempty"`
	Port                *int    `url:"port,omitempty",json:"port,omitempty"`
	Tfa                 *string `url:"tfa,omitempty",json:"tfa,omitempty"`
	Realm               string  `url:"realm",json:"realm"`
	Scopes              *string `url:"scopes,omitempty",json:"scopes,omitempty"`
	Secure              *bool   `url:"secure,omitempty",json:"secure,omitempty"`
	UserClasses         *string `url:"user_classes,omitempty",json:"user_classes,omitempty"`
	BaseDn              *string `url:"base_dn,omitempty",json:"base_dn,omitempty"`
	Certkey             *string `url:"certkey,omitempty",json:"certkey,omitempty"`
	ClientKey           *string `url:"client-key,omitempty",json:"client-key,omitempty"`
	Filter              *string `url:"filter,omitempty",json:"filter,omitempty"`
	UsernameClaim       *string `url:"username-claim,omitempty",json:"username-claim,omitempty"`
	Cert                *string `url:"cert,omitempty",json:"cert,omitempty"`
	UserAttr            *string `url:"user_attr,omitempty",json:"user_attr,omitempty"`
	Comment             *string `url:"comment,omitempty",json:"comment,omitempty"`
	Default             *bool   `url:"default,omitempty",json:"default,omitempty"`
	GroupDn             *string `url:"group_dn,omitempty",json:"group_dn,omitempty"`
	GroupNameAttr       *string `url:"group_name_attr,omitempty",json:"group_name_attr,omitempty"`
	Autocreate          *bool   `url:"autocreate,omitempty",json:"autocreate,omitempty"`
	BindDn              *string `url:"bind_dn,omitempty",json:"bind_dn,omitempty"`
	Capath              *string `url:"capath,omitempty",json:"capath,omitempty"`
	ClientId            *string `url:"client-id,omitempty",json:"client-id,omitempty"`
	AcrValues           *string `url:"acr-values,omitempty",json:"acr-values,omitempty"`
	CaseSensitive       *bool   `url:"case-sensitive,omitempty",json:"case-sensitive,omitempty"`
	Domain              *string `url:"domain,omitempty",json:"domain,omitempty"`
	Type                string  `url:"type",json:"type"`
	IssuerUrl           *string `url:"issuer-url,omitempty",json:"issuer-url,omitempty"`
	Server1             *string `url:"server1,omitempty",json:"server1,omitempty"`
	SyncAttributes      *string `url:"sync_attributes,omitempty",json:"sync_attributes,omitempty"`
	Verify              *bool   `url:"verify,omitempty",json:"verify,omitempty"`
	GroupClasses        *string `url:"group_classes,omitempty",json:"group_classes,omitempty"`
	Prompt              *string `url:"prompt,omitempty",json:"prompt,omitempty"`
	Server2             *string `url:"server2,omitempty",json:"server2,omitempty"`
	GroupFilter         *string `url:"group_filter,omitempty",json:"group_filter,omitempty"`
	Sslversion          *string `url:"sslversion,omitempty",json:"sslversion,omitempty"`
	SyncDefaultsOptions *string `url:"sync-defaults-options,omitempty",json:"sync-defaults-options,omitempty"`
}

type CreateResponse map[string]interface{}

// Create Add an authentication server.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/access/domains", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Realm string `url:"realm",json:"realm"`
}

// Find Get auth server configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) error {

	err := c.httpClient.Do(ctx, "/access/domains/{realm}", "GET", nil, req)
	return err
}

type UpdateRequest struct {
	Default             *bool   `url:"default,omitempty",json:"default,omitempty"`
	Mode                *string `url:"mode,omitempty",json:"mode,omitempty"`
	Server1             *string `url:"server1,omitempty",json:"server1,omitempty"`
	ClientKey           *string `url:"client-key,omitempty",json:"client-key,omitempty"`
	Prompt              *string `url:"prompt,omitempty",json:"prompt,omitempty"`
	Tfa                 *string `url:"tfa,omitempty",json:"tfa,omitempty"`
	ClientId            *string `url:"client-id,omitempty",json:"client-id,omitempty"`
	Comment             *string `url:"comment,omitempty",json:"comment,omitempty"`
	GroupClasses        *string `url:"group_classes,omitempty",json:"group_classes,omitempty"`
	GroupDn             *string `url:"group_dn,omitempty",json:"group_dn,omitempty"`
	Server2             *string `url:"server2,omitempty",json:"server2,omitempty"`
	Digest              *string `url:"digest,omitempty",json:"digest,omitempty"`
	UserClasses         *string `url:"user_classes,omitempty",json:"user_classes,omitempty"`
	SyncDefaultsOptions *string `url:"sync-defaults-options,omitempty",json:"sync-defaults-options,omitempty"`
	SyncAttributes      *string `url:"sync_attributes,omitempty",json:"sync_attributes,omitempty"`
	CaseSensitive       *bool   `url:"case-sensitive,omitempty",json:"case-sensitive,omitempty"`
	Cert                *string `url:"cert,omitempty",json:"cert,omitempty"`
	Domain              *string `url:"domain,omitempty",json:"domain,omitempty"`
	GroupNameAttr       *string `url:"group_name_attr,omitempty",json:"group_name_attr,omitempty"`
	IssuerUrl           *string `url:"issuer-url,omitempty",json:"issuer-url,omitempty"`
	Sslversion          *string `url:"sslversion,omitempty",json:"sslversion,omitempty"`
	GroupFilter         *string `url:"group_filter,omitempty",json:"group_filter,omitempty"`
	Port                *int    `url:"port,omitempty",json:"port,omitempty"`
	Realm               string  `url:"realm",json:"realm"`
	Scopes              *string `url:"scopes,omitempty",json:"scopes,omitempty"`
	UserAttr            *string `url:"user_attr,omitempty",json:"user_attr,omitempty"`
	AcrValues           *string `url:"acr-values,omitempty",json:"acr-values,omitempty"`
	Autocreate          *bool   `url:"autocreate,omitempty",json:"autocreate,omitempty"`
	BindDn              *string `url:"bind_dn,omitempty",json:"bind_dn,omitempty"`
	Capath              *string `url:"capath,omitempty",json:"capath,omitempty"`
	Certkey             *string `url:"certkey,omitempty",json:"certkey,omitempty"`
	Filter              *string `url:"filter,omitempty",json:"filter,omitempty"`
	BaseDn              *string `url:"base_dn,omitempty",json:"base_dn,omitempty"`
	Delete              *string `url:"delete,omitempty",json:"delete,omitempty"`
	Password            *string `url:"password,omitempty",json:"password,omitempty"`
	Secure              *bool   `url:"secure,omitempty",json:"secure,omitempty"`
	Verify              *bool   `url:"verify,omitempty",json:"verify,omitempty"`
}

type UpdateResponse map[string]interface{}

// Update Update authentication server settings.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/access/domains/{realm}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Realm string `url:"realm",json:"realm"`
}

type DeleteResponse map[string]interface{}

// Delete Delete an authentication server.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/access/domains/{realm}", "DELETE", &resp, req)
	return resp, err
}

type SyncRequest struct {
	Full           *bool   `url:"full,omitempty",json:"full,omitempty"`
	Purge          *bool   `url:"purge,omitempty",json:"purge,omitempty"`
	Realm          string  `url:"realm",json:"realm"`
	RemoveVanished *string `url:"remove-vanished,omitempty",json:"remove-vanished,omitempty"`
	Scope          *string `url:"scope,omitempty",json:"scope,omitempty"`
	DryRun         *bool   `url:"dry-run,omitempty",json:"dry-run,omitempty"`
	EnableNew      *bool   `url:"enable-new,omitempty",json:"enable-new,omitempty"`
}

type SyncResponse string

// Sync Syncs users and/or groups from the configured LDAP to user.cfg. NOTE: Synced groups will have the name 'name-$realm', so make sure those groups do not exist to prevent overwriting.
func (c *Client) Sync(ctx context.Context, req *SyncRequest) (*SyncResponse, error) {
	var resp *SyncResponse

	err := c.httpClient.Do(ctx, "/access/domains/{realm}/sync", "POST", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package domains

import (
	"context"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
)

const (
	Mode_LDAP              Mode = "ldap"
	Mode_LDAPS             Mode = "ldaps"
	Mode_LDAP_AND_STARTTLS Mode = "ldap+starttls"

	Scope_USERS  Scope = "users"
	Scope_GROUPS Scope = "groups"
	Scope_BOTH   Scope = "both"

	Sslversion_TLSV1   Sslversion = "tlsv1"
	Sslversion_TLSV1_1 Sslversion = "tlsv1_1"
	Sslversion_TLSV1_2 Sslversion = "tlsv1_2"
	Sslversion_TLSV1_3 Sslversion = "tlsv1_3"

	Tfa_YUBICO Tfa = "yubico"
	Tfa_OATH   Tfa = "oath"

	Type_AD     Type = "ad"
	Type_LDAP   Type = "ldap"
	Type_OPENID Type = "openid"
	Type_PAM    Type = "pam"
	Type_PVE    Type = "pve"
)

type Mode string
type Scope string
type Sslversion string
type Tfa string
type Type string

func PtrMode(i Mode) *Mode {
	return &i
}
func PtrScope(i Scope) *Scope {
	return &i
}
func PtrSslversion(i Sslversion) *Sslversion {
	return &i
}
func PtrTfa(i Tfa) *Tfa {
	return &i
}
func PtrType(i Type) *Type {
	return &i
}

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

type IndexResponse struct {
	Realm string `url:"realm" json:"realm"`
	Type  string `url:"type" json:"type"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"` // A comment. The GUI use this text when you select a domain (Realm) on the login window.
	Tfa     *Tfa    `url:"tfa,omitempty" json:"tfa,omitempty"`         // Two-factor authentication provider.
}
type _IndexResponse IndexResponse

type CreateRequest struct {
	Realm string `url:"realm" json:"realm"` // Authentication domain ID
	Type  Type   `url:"type" json:"type"`   // Realm type.

	// The following parameters are optional
	AcrValues           *string       `url:"acr-values,omitempty" json:"acr-values,omitempty"`                       // Specifies the Authentication Context Class Reference values that theAuthorization Server is being requested to use for the Auth Request.
	Autocreate          *util.PVEBool `url:"autocreate,omitempty" json:"autocreate,omitempty"`                       // Automatically create users if they do not exist.
	BaseDn              *string       `url:"base_dn,omitempty" json:"base_dn,omitempty"`                             // LDAP base domain name
	BindDn              *string       `url:"bind_dn,omitempty" json:"bind_dn,omitempty"`                             // LDAP bind domain name
	Capath              *string       `url:"capath,omitempty" json:"capath,omitempty"`                               // Path to the CA certificate store
	CaseSensitive       *util.PVEBool `url:"case-sensitive,omitempty" json:"case-sensitive,omitempty"`               // username is case-sensitive
	Cert                *string       `url:"cert,omitempty" json:"cert,omitempty"`                                   // Path to the client certificate
	Certkey             *string       `url:"certkey,omitempty" json:"certkey,omitempty"`                             // Path to the client certificate key
	ClientId            *string       `url:"client-id,omitempty" json:"client-id,omitempty"`                         // OpenID Client ID
	ClientKey           *string       `url:"client-key,omitempty" json:"client-key,omitempty"`                       // OpenID Client Key
	Comment             *string       `url:"comment,omitempty" json:"comment,omitempty"`                             // Description.
	Default             *util.PVEBool `url:"default,omitempty" json:"default,omitempty"`                             // Use this as default realm
	Domain              *string       `url:"domain,omitempty" json:"domain,omitempty"`                               // AD domain name
	Filter              *string       `url:"filter,omitempty" json:"filter,omitempty"`                               // LDAP filter for user sync.
	GroupClasses        *string       `url:"group_classes,omitempty" json:"group_classes,omitempty"`                 // The objectclasses for groups.
	GroupDn             *string       `url:"group_dn,omitempty" json:"group_dn,omitempty"`                           // LDAP base domain name for group sync. If not set, the base_dn will be used.
	GroupFilter         *string       `url:"group_filter,omitempty" json:"group_filter,omitempty"`                   // LDAP filter for group sync.
	GroupNameAttr       *string       `url:"group_name_attr,omitempty" json:"group_name_attr,omitempty"`             // LDAP attribute representing a groups name. If not set or found, the first value of the DN will be used as name.
	IssuerUrl           *string       `url:"issuer-url,omitempty" json:"issuer-url,omitempty"`                       // OpenID Issuer Url
	Mode                *Mode         `url:"mode,omitempty" json:"mode,omitempty"`                                   // LDAP protocol mode.
	Password            *string       `url:"password,omitempty" json:"password,omitempty"`                           // LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.
	Port                *int          `url:"port,omitempty" json:"port,omitempty"`                                   // Server port.
	Prompt              *string       `url:"prompt,omitempty" json:"prompt,omitempty"`                               // Specifies whether the Authorization Server prompts the End-User for reauthentication and consent.
	Scopes              *string       `url:"scopes,omitempty" json:"scopes,omitempty"`                               // Specifies the scopes (user details) that should be authorized and returned, for example 'email' or 'profile'.
	Secure              *util.PVEBool `url:"secure,omitempty" json:"secure,omitempty"`                               // Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.
	Server1             *string       `url:"server1,omitempty" json:"server1,omitempty"`                             // Server IP address (or DNS name)
	Server2             *string       `url:"server2,omitempty" json:"server2,omitempty"`                             // Fallback Server IP address (or DNS name)
	Sslversion          *Sslversion   `url:"sslversion,omitempty" json:"sslversion,omitempty"`                       // LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!
	SyncAttributes      *string       `url:"sync_attributes,omitempty" json:"sync_attributes,omitempty"`             // Comma separated list of key=value pairs for specifying which LDAP attributes map to which PVE user field. For example, to map the LDAP attribute 'mail' to PVEs 'email', write 'email=mail'. By default, each PVE user field is represented by an LDAP attribute of the same name.
	SyncDefaultsOptions *string       `url:"sync-defaults-options,omitempty" json:"sync-defaults-options,omitempty"` // The default options for behavior of synchronizations.
	Tfa                 *string       `url:"tfa,omitempty" json:"tfa,omitempty"`                                     // Use Two-factor authentication.
	UserAttr            *string       `url:"user_attr,omitempty" json:"user_attr,omitempty"`                         // LDAP user attribute name
	UserClasses         *string       `url:"user_classes,omitempty" json:"user_classes,omitempty"`                   // The objectclasses for users.
	UsernameClaim       *string       `url:"username-claim,omitempty" json:"username-claim,omitempty"`               // OpenID claim used to generate the unique username.
	Verify              *util.PVEBool `url:"verify,omitempty" json:"verify,omitempty"`                               // Verify the server's SSL certificate
}
type _CreateRequest CreateRequest

type FindRequest struct {
	Realm string `url:"realm" json:"realm"` // Authentication domain ID

}
type _FindRequest FindRequest

type UpdateRequest struct {
	Realm string `url:"realm" json:"realm"` // Authentication domain ID

	// The following parameters are optional
	AcrValues           *string       `url:"acr-values,omitempty" json:"acr-values,omitempty"`                       // Specifies the Authentication Context Class Reference values that theAuthorization Server is being requested to use for the Auth Request.
	Autocreate          *util.PVEBool `url:"autocreate,omitempty" json:"autocreate,omitempty"`                       // Automatically create users if they do not exist.
	BaseDn              *string       `url:"base_dn,omitempty" json:"base_dn,omitempty"`                             // LDAP base domain name
	BindDn              *string       `url:"bind_dn,omitempty" json:"bind_dn,omitempty"`                             // LDAP bind domain name
	Capath              *string       `url:"capath,omitempty" json:"capath,omitempty"`                               // Path to the CA certificate store
	CaseSensitive       *util.PVEBool `url:"case-sensitive,omitempty" json:"case-sensitive,omitempty"`               // username is case-sensitive
	Cert                *string       `url:"cert,omitempty" json:"cert,omitempty"`                                   // Path to the client certificate
	Certkey             *string       `url:"certkey,omitempty" json:"certkey,omitempty"`                             // Path to the client certificate key
	ClientId            *string       `url:"client-id,omitempty" json:"client-id,omitempty"`                         // OpenID Client ID
	ClientKey           *string       `url:"client-key,omitempty" json:"client-key,omitempty"`                       // OpenID Client Key
	Comment             *string       `url:"comment,omitempty" json:"comment,omitempty"`                             // Description.
	Default             *util.PVEBool `url:"default,omitempty" json:"default,omitempty"`                             // Use this as default realm
	Delete              *string       `url:"delete,omitempty" json:"delete,omitempty"`                               // A list of settings you want to delete.
	Digest              *string       `url:"digest,omitempty" json:"digest,omitempty"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Domain              *string       `url:"domain,omitempty" json:"domain,omitempty"`                               // AD domain name
	Filter              *string       `url:"filter,omitempty" json:"filter,omitempty"`                               // LDAP filter for user sync.
	GroupClasses        *string       `url:"group_classes,omitempty" json:"group_classes,omitempty"`                 // The objectclasses for groups.
	GroupDn             *string       `url:"group_dn,omitempty" json:"group_dn,omitempty"`                           // LDAP base domain name for group sync. If not set, the base_dn will be used.
	GroupFilter         *string       `url:"group_filter,omitempty" json:"group_filter,omitempty"`                   // LDAP filter for group sync.
	GroupNameAttr       *string       `url:"group_name_attr,omitempty" json:"group_name_attr,omitempty"`             // LDAP attribute representing a groups name. If not set or found, the first value of the DN will be used as name.
	IssuerUrl           *string       `url:"issuer-url,omitempty" json:"issuer-url,omitempty"`                       // OpenID Issuer Url
	Mode                *Mode         `url:"mode,omitempty" json:"mode,omitempty"`                                   // LDAP protocol mode.
	Password            *string       `url:"password,omitempty" json:"password,omitempty"`                           // LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.
	Port                *int          `url:"port,omitempty" json:"port,omitempty"`                                   // Server port.
	Prompt              *string       `url:"prompt,omitempty" json:"prompt,omitempty"`                               // Specifies whether the Authorization Server prompts the End-User for reauthentication and consent.
	Scopes              *string       `url:"scopes,omitempty" json:"scopes,omitempty"`                               // Specifies the scopes (user details) that should be authorized and returned, for example 'email' or 'profile'.
	Secure              *util.PVEBool `url:"secure,omitempty" json:"secure,omitempty"`                               // Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.
	Server1             *string       `url:"server1,omitempty" json:"server1,omitempty"`                             // Server IP address (or DNS name)
	Server2             *string       `url:"server2,omitempty" json:"server2,omitempty"`                             // Fallback Server IP address (or DNS name)
	Sslversion          *Sslversion   `url:"sslversion,omitempty" json:"sslversion,omitempty"`                       // LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!
	SyncAttributes      *string       `url:"sync_attributes,omitempty" json:"sync_attributes,omitempty"`             // Comma separated list of key=value pairs for specifying which LDAP attributes map to which PVE user field. For example, to map the LDAP attribute 'mail' to PVEs 'email', write 'email=mail'. By default, each PVE user field is represented by an LDAP attribute of the same name.
	SyncDefaultsOptions *string       `url:"sync-defaults-options,omitempty" json:"sync-defaults-options,omitempty"` // The default options for behavior of synchronizations.
	Tfa                 *string       `url:"tfa,omitempty" json:"tfa,omitempty"`                                     // Use Two-factor authentication.
	UserAttr            *string       `url:"user_attr,omitempty" json:"user_attr,omitempty"`                         // LDAP user attribute name
	UserClasses         *string       `url:"user_classes,omitempty" json:"user_classes,omitempty"`                   // The objectclasses for users.
	Verify              *util.PVEBool `url:"verify,omitempty" json:"verify,omitempty"`                               // Verify the server's SSL certificate
}
type _UpdateRequest UpdateRequest

type DeleteRequest struct {
	Realm string `url:"realm" json:"realm"` // Authentication domain ID

}
type _DeleteRequest DeleteRequest

type SyncRequest struct {
	Realm string `url:"realm" json:"realm"` // Authentication domain ID

	// The following parameters are optional
	DryRun         *util.PVEBool `url:"dry-run,omitempty" json:"dry-run,omitempty"`                 // If set, does not write anything.
	EnableNew      *util.PVEBool `url:"enable-new,omitempty" json:"enable-new,omitempty"`           // Enable newly synced users immediately.
	Full           *util.PVEBool `url:"full,omitempty" json:"full,omitempty"`                       // DEPRECATED: use 'remove-vanished' instead. If set, uses the LDAP Directory as source of truth, deleting users or groups not returned from the sync and removing all locally modified properties of synced users. If not set, only syncs information which is present in the synced data, and does not delete or modify anything else.
	Purge          *util.PVEBool `url:"purge,omitempty" json:"purge,omitempty"`                     // DEPRECATED: use 'remove-vanished' instead. Remove ACLs for users or groups which were removed from the config during a sync.
	RemoveVanished *string       `url:"remove-vanished,omitempty" json:"remove-vanished,omitempty"` // A semicolon-seperated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	Scope          *Scope        `url:"scope,omitempty" json:"scope,omitempty"`                     // Select what to sync.
}
type _SyncRequest SyncRequest

// Index Authentication domain index.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/access/domains", "GET", &resp, nil)
	return resp, err
}

// Create Add an authentication server.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/access/domains", "POST", nil, req)
	return err
}

// Find Get auth server configuration.
func (c *Client) Find(ctx context.Context, req FindRequest) error {

	err := c.httpClient.Do(ctx, "/access/domains/{realm}", "GET", nil, req)
	return err
}

// Update Update authentication server settings.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/access/domains/{realm}", "PUT", nil, req)
	return err
}

// Delete Delete an authentication server.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/access/domains/{realm}", "DELETE", nil, req)
	return err
}

// Sync Syncs users and/or groups from the configured LDAP to user.cfg. NOTE: Synced groups will have the name 'name-$realm', so make sure those groups do not exist to prevent overwriting.
func (c *Client) Sync(ctx context.Context, req SyncRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/access/domains/{realm}/sync", "POST", &resp, req)
	return resp, err
}

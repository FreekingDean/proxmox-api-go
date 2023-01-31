// File generated by proxmox json schema, DO NOT EDIT

package roles

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

type IndexResponse struct {
	Roleid string `url:"roleid" json:"roleid"`

	// The following parameters are optional
	Privs   *string       `url:"privs,omitempty" json:"privs,omitempty"`
	Special *util.PVEBool `url:"special,omitempty" json:"special,omitempty"`
}
type _IndexResponse IndexResponse

type CreateRequest struct {
	Roleid string `url:"roleid" json:"roleid"`

	// The following parameters are optional
	Privs *string `url:"privs,omitempty" json:"privs,omitempty"`
}
type _CreateRequest CreateRequest

type FindRequest struct {
	Roleid string `url:"roleid" json:"roleid"`
}
type _FindRequest FindRequest

type FindResponse struct {

	// The following parameters are optional
	DatastoreAllocate         *util.PVEBool `url:"Datastore.Allocate,omitempty" json:"Datastore.Allocate,omitempty"`
	DatastoreAllocatespace    *util.PVEBool `url:"Datastore.AllocateSpace,omitempty" json:"Datastore.AllocateSpace,omitempty"`
	DatastoreAllocatetemplate *util.PVEBool `url:"Datastore.AllocateTemplate,omitempty" json:"Datastore.AllocateTemplate,omitempty"`
	DatastoreAudit            *util.PVEBool `url:"Datastore.Audit,omitempty" json:"Datastore.Audit,omitempty"`
	GroupAllocate             *util.PVEBool `url:"Group.Allocate,omitempty" json:"Group.Allocate,omitempty"`
	PermissionsModify         *util.PVEBool `url:"Permissions.Modify,omitempty" json:"Permissions.Modify,omitempty"`
	PoolAllocate              *util.PVEBool `url:"Pool.Allocate,omitempty" json:"Pool.Allocate,omitempty"`
	PoolAudit                 *util.PVEBool `url:"Pool.Audit,omitempty" json:"Pool.Audit,omitempty"`
	RealmAllocate             *util.PVEBool `url:"Realm.Allocate,omitempty" json:"Realm.Allocate,omitempty"`
	RealmAllocateuser         *util.PVEBool `url:"Realm.AllocateUser,omitempty" json:"Realm.AllocateUser,omitempty"`
	SdnAllocate               *util.PVEBool `url:"SDN.Allocate,omitempty" json:"SDN.Allocate,omitempty"`
	SdnAudit                  *util.PVEBool `url:"SDN.Audit,omitempty" json:"SDN.Audit,omitempty"`
	SysAudit                  *util.PVEBool `url:"Sys.Audit,omitempty" json:"Sys.Audit,omitempty"`
	SysConsole                *util.PVEBool `url:"Sys.Console,omitempty" json:"Sys.Console,omitempty"`
	SysIncoming               *util.PVEBool `url:"Sys.Incoming,omitempty" json:"Sys.Incoming,omitempty"`
	SysModify                 *util.PVEBool `url:"Sys.Modify,omitempty" json:"Sys.Modify,omitempty"`
	SysPowermgmt              *util.PVEBool `url:"Sys.PowerMgmt,omitempty" json:"Sys.PowerMgmt,omitempty"`
	SysSyslog                 *util.PVEBool `url:"Sys.Syslog,omitempty" json:"Sys.Syslog,omitempty"`
	UserModify                *util.PVEBool `url:"User.Modify,omitempty" json:"User.Modify,omitempty"`
	VmAllocate                *util.PVEBool `url:"VM.Allocate,omitempty" json:"VM.Allocate,omitempty"`
	VmAudit                   *util.PVEBool `url:"VM.Audit,omitempty" json:"VM.Audit,omitempty"`
	VmBackup                  *util.PVEBool `url:"VM.Backup,omitempty" json:"VM.Backup,omitempty"`
	VmClone                   *util.PVEBool `url:"VM.Clone,omitempty" json:"VM.Clone,omitempty"`
	VmConfigCdrom             *util.PVEBool `url:"VM.Config.CDROM,omitempty" json:"VM.Config.CDROM,omitempty"`
	VmConfigCloudinit         *util.PVEBool `url:"VM.Config.Cloudinit,omitempty" json:"VM.Config.Cloudinit,omitempty"`
	VmConfigCpu               *util.PVEBool `url:"VM.Config.CPU,omitempty" json:"VM.Config.CPU,omitempty"`
	VmConfigDisk              *util.PVEBool `url:"VM.Config.Disk,omitempty" json:"VM.Config.Disk,omitempty"`
	VmConfigHwtype            *util.PVEBool `url:"VM.Config.HWType,omitempty" json:"VM.Config.HWType,omitempty"`
	VmConfigMemory            *util.PVEBool `url:"VM.Config.Memory,omitempty" json:"VM.Config.Memory,omitempty"`
	VmConfigNetwork           *util.PVEBool `url:"VM.Config.Network,omitempty" json:"VM.Config.Network,omitempty"`
	VmConfigOptions           *util.PVEBool `url:"VM.Config.Options,omitempty" json:"VM.Config.Options,omitempty"`
	VmConsole                 *util.PVEBool `url:"VM.Console,omitempty" json:"VM.Console,omitempty"`
	VmMigrate                 *util.PVEBool `url:"VM.Migrate,omitempty" json:"VM.Migrate,omitempty"`
	VmMonitor                 *util.PVEBool `url:"VM.Monitor,omitempty" json:"VM.Monitor,omitempty"`
	VmPowermgmt               *util.PVEBool `url:"VM.PowerMgmt,omitempty" json:"VM.PowerMgmt,omitempty"`
	VmSnapshot                *util.PVEBool `url:"VM.Snapshot,omitempty" json:"VM.Snapshot,omitempty"`
	VmSnapshotRollback        *util.PVEBool `url:"VM.Snapshot.Rollback,omitempty" json:"VM.Snapshot.Rollback,omitempty"`
}
type _FindResponse FindResponse

type UpdateRequest struct {
	Roleid string `url:"roleid" json:"roleid"`

	// The following parameters are optional
	Append *util.PVEBool `url:"append,omitempty" json:"append,omitempty"`
	Privs  *string       `url:"privs,omitempty" json:"privs,omitempty"`
}
type _UpdateRequest UpdateRequest

type DeleteRequest struct {
	Roleid string `url:"roleid" json:"roleid"`
}
type _DeleteRequest DeleteRequest

// Index Role index.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/access/roles", "GET", &resp, nil)
	return resp, err
}

// Create Create new role.
func (c *Client) Create(ctx context.Context, req CreateRequest) error {

	err := c.httpClient.Do(ctx, "/access/roles", "POST", nil, req)
	return err
}

// Find Get role configuration.
func (c *Client) Find(ctx context.Context, req FindRequest) (FindResponse, error) {
	var resp FindResponse

	err := c.httpClient.Do(ctx, "/access/roles/{roleid}", "GET", &resp, req)
	return resp, err
}

// Update Update an existing role.
func (c *Client) Update(ctx context.Context, req UpdateRequest) error {

	err := c.httpClient.Do(ctx, "/access/roles/{roleid}", "PUT", nil, req)
	return err
}

// Delete Delete role.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/access/roles/{roleid}", "DELETE", nil, req)
	return err
}

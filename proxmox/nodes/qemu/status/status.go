// File generated by proxmox json schema, DO NOT EDIT

package status

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
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type IndexResponse []*struct {
	Subdir string `url:"subdir",json:"subdir"`
}

// Index Directory index
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status", "GET", &resp, req)
	return resp, err
}

type VmStatusCurrentRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

}

type VmStatusCurrentResponse struct {
	Ha     map[string]interface{} `url:"ha",json:"ha"`         // HA manager service status.
	Status string                 `url:"status",json:"status"` // Qemu process status.
	Vmid   int                    `url:"vmid",json:"vmid"`     // The (unique) ID of the VM.

	// The following parameters are optional
	Agent          *util.SpecialBool `url:"agent,omitempty",json:"agent,omitempty"`                     // Qemu GuestAgent enabled in config.
	Cpus           *float64          `url:"cpus,omitempty",json:"cpus,omitempty"`                       // Maximum usable CPUs.
	Lock           *string           `url:"lock,omitempty",json:"lock,omitempty"`                       // The current config lock, if any.
	Maxdisk        *int              `url:"maxdisk,omitempty",json:"maxdisk,omitempty"`                 // Root disk size in bytes.
	Maxmem         *int              `url:"maxmem,omitempty",json:"maxmem,omitempty"`                   // Maximum memory in bytes.
	Name           *string           `url:"name,omitempty",json:"name,omitempty"`                       // VM name.
	Pid            *int              `url:"pid,omitempty",json:"pid,omitempty"`                         // PID of running qemu process.
	Qmpstatus      *string           `url:"qmpstatus,omitempty",json:"qmpstatus,omitempty"`             // Qemu QMP agent status.
	RunningMachine *string           `url:"running-machine,omitempty",json:"running-machine,omitempty"` // The currently running machine type (if running).
	RunningQemu    *string           `url:"running-qemu,omitempty",json:"running-qemu,omitempty"`       // The currently running QEMU version (if running).
	Spice          *util.SpecialBool `url:"spice,omitempty",json:"spice,omitempty"`                     // Qemu VGA configuration supports spice.
	Tags           *string           `url:"tags,omitempty",json:"tags,omitempty"`                       // The current configured tags, if any
	Uptime         *int              `url:"uptime,omitempty",json:"uptime,omitempty"`                   // Uptime.
}

// VmStatusCurrent Get virtual machine status.
func (c *Client) VmStatusCurrent(ctx context.Context, req *VmStatusCurrentRequest) (*VmStatusCurrentResponse, error) {
	var resp *VmStatusCurrentResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/current", "GET", &resp, req)
	return resp, err
}

type VmStartRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	ForceCpu         *string           `url:"force-cpu,omitempty",json:"force-cpu,omitempty"`                 // Override QEMU's -cpu argument with the given string.
	Machine          *string           `url:"machine,omitempty",json:"machine,omitempty"`                     // Specifies the Qemu machine type.
	Migratedfrom     *string           `url:"migratedfrom,omitempty",json:"migratedfrom,omitempty"`           // The cluster node name.
	MigrationNetwork *string           `url:"migration_network,omitempty",json:"migration_network,omitempty"` // CIDR of the (sub) network that is used for migration.
	MigrationType    *string           `url:"migration_type,omitempty",json:"migration_type,omitempty"`       // Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.
	Skiplock         *util.SpecialBool `url:"skiplock,omitempty",json:"skiplock,omitempty"`                   // Ignore locks - only root is allowed to use this option.
	Stateuri         *string           `url:"stateuri,omitempty",json:"stateuri,omitempty"`                   // Some command save/restore state from this location.
	Targetstorage    *string           `url:"targetstorage,omitempty",json:"targetstorage,omitempty"`         // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Timeout          *int              `url:"timeout,omitempty",json:"timeout,omitempty"`                     // Wait maximal timeout seconds.
}

type VmStartResponse string

// VmStart Start virtual machine.
func (c *Client) VmStart(ctx context.Context, req *VmStartRequest) (*VmStartResponse, error) {
	var resp *VmStartResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/start", "POST", &resp, req)
	return resp, err
}

type VmStopRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Keepactive   *util.SpecialBool `url:"keepActive,omitempty",json:"keepActive,omitempty"`     // Do not deactivate storage volumes.
	Migratedfrom *string           `url:"migratedfrom,omitempty",json:"migratedfrom,omitempty"` // The cluster node name.
	Skiplock     *util.SpecialBool `url:"skiplock,omitempty",json:"skiplock,omitempty"`         // Ignore locks - only root is allowed to use this option.
	Timeout      *int              `url:"timeout,omitempty",json:"timeout,omitempty"`           // Wait maximal timeout seconds.
}

type VmStopResponse string

// VmStop Stop virtual machine. The qemu process will exit immediately. Thisis akin to pulling the power plug of a running computer and may damage the VM data
func (c *Client) VmStop(ctx context.Context, req *VmStopRequest) (*VmStopResponse, error) {
	var resp *VmStopResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/stop", "POST", &resp, req)
	return resp, err
}

type VmResetRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Skiplock *util.SpecialBool `url:"skiplock,omitempty",json:"skiplock,omitempty"` // Ignore locks - only root is allowed to use this option.
}

type VmResetResponse string

// VmReset Reset virtual machine.
func (c *Client) VmReset(ctx context.Context, req *VmResetRequest) (*VmResetResponse, error) {
	var resp *VmResetResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/reset", "POST", &resp, req)
	return resp, err
}

type VmShutdownRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Forcestop  *util.SpecialBool `url:"forceStop,omitempty",json:"forceStop,omitempty"`   // Make sure the VM stops.
	Keepactive *util.SpecialBool `url:"keepActive,omitempty",json:"keepActive,omitempty"` // Do not deactivate storage volumes.
	Skiplock   *util.SpecialBool `url:"skiplock,omitempty",json:"skiplock,omitempty"`     // Ignore locks - only root is allowed to use this option.
	Timeout    *int              `url:"timeout,omitempty",json:"timeout,omitempty"`       // Wait maximal timeout seconds.
}

type VmShutdownResponse string

// VmShutdown Shutdown virtual machine. This is similar to pressing the power button on a physical machine.This will send an ACPI event for the guest OS, which should then proceed to a clean shutdown.
func (c *Client) VmShutdown(ctx context.Context, req *VmShutdownRequest) (*VmShutdownResponse, error) {
	var resp *VmShutdownResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/shutdown", "POST", &resp, req)
	return resp, err
}

type VmRebootRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Timeout *int `url:"timeout,omitempty",json:"timeout,omitempty"` // Wait maximal timeout seconds for the shutdown.
}

type VmRebootResponse string

// VmReboot Reboot the VM by shutting it down, and starting it again. Applies pending changes.
func (c *Client) VmReboot(ctx context.Context, req *VmRebootRequest) (*VmRebootResponse, error) {
	var resp *VmRebootResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/reboot", "POST", &resp, req)
	return resp, err
}

type VmSuspendRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Skiplock     *util.SpecialBool `url:"skiplock,omitempty",json:"skiplock,omitempty"`         // Ignore locks - only root is allowed to use this option.
	Statestorage *string           `url:"statestorage,omitempty",json:"statestorage,omitempty"` // The storage for the VM state
	Todisk       *util.SpecialBool `url:"todisk,omitempty",json:"todisk,omitempty"`             // If set, suspends the VM to disk. Will be resumed on next VM start.
}

type VmSuspendResponse string

// VmSuspend Suspend virtual machine.
func (c *Client) VmSuspend(ctx context.Context, req *VmSuspendRequest) (*VmSuspendResponse, error) {
	var resp *VmSuspendResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/suspend", "POST", &resp, req)
	return resp, err
}

type VmResumeRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Vmid int    `url:"vmid",json:"vmid"` // The (unique) ID of the VM.

	// The following parameters are optional
	Nocheck  *util.SpecialBool `url:"nocheck,omitempty",json:"nocheck,omitempty"`
	Skiplock *util.SpecialBool `url:"skiplock,omitempty",json:"skiplock,omitempty"` // Ignore locks - only root is allowed to use this option.
}

type VmResumeResponse string

// VmResume Resume virtual machine.
func (c *Client) VmResume(ctx context.Context, req *VmResumeRequest) (*VmResumeResponse, error) {
	var resp *VmResumeResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/qemu/{vmid}/status/resume", "POST", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package nodes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/FreekingDean/proxmox-api-go/internal/util"
)

const (
	Cf_AVERAGE Cf = "AVERAGE"
	Cf_MAX     Cf = "MAX"

	Cmd_CEPH_INSTALL Cmd = "ceph_install"
	Cmd_LOGIN        Cmd = "login"
	Cmd_UPGRADE      Cmd = "upgrade"

	Command_REBOOT   Command = "reboot"
	Command_SHUTDOWN Command = "shutdown"

	Property_ACME                  Property = "acme"
	Property_ACMEDOMAIN0           Property = "acmedomain0"
	Property_ACMEDOMAIN1           Property = "acmedomain1"
	Property_ACMEDOMAIN2           Property = "acmedomain2"
	Property_ACMEDOMAIN3           Property = "acmedomain3"
	Property_ACMEDOMAIN4           Property = "acmedomain4"
	Property_ACMEDOMAIN5           Property = "acmedomain5"
	Property_DESCRIPTION           Property = "description"
	Property_STARTALL_ONBOOT_DELAY Property = "startall-onboot-delay"
	Property_WAKEONLAN             Property = "wakeonlan"

	Status_UNKNOWN Status = "unknown"
	Status_ONLINE  Status = "online"
	Status_OFFLINE Status = "offline"

	Timeframe_HOUR  Timeframe = "hour"
	Timeframe_DAY   Timeframe = "day"
	Timeframe_WEEK  Timeframe = "week"
	Timeframe_MONTH Timeframe = "month"
	Timeframe_YEAR  Timeframe = "year"
)

type Cf string
type Cmd string
type Command string
type Property string
type Status string
type Timeframe string

func PtrCf(i Cf) *Cf {
	return &i
}
func PtrCmd(i Cmd) *Cmd {
	return &i
}
func PtrCommand(i Command) *Command {
	return &i
}
func PtrProperty(i Property) *Property {
	return &i
}
func PtrStatus(i Status) *Status {
	return &i
}
func PtrTimeframe(i Timeframe) *Timeframe {
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
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Status Status `url:"status" json:"status"` // Node status.

	// The following parameters are optional
	Cpu            *float64 `url:"cpu,omitempty" json:"cpu,omitempty"`                         // CPU utilization.
	Level          *string  `url:"level,omitempty" json:"level,omitempty"`                     // Support level.
	Maxcpu         *int     `url:"maxcpu,omitempty" json:"maxcpu,omitempty"`                   // Number of available CPUs.
	Maxmem         *int     `url:"maxmem,omitempty" json:"maxmem,omitempty"`                   // Number of available memory in bytes.
	Mem            *int     `url:"mem,omitempty" json:"mem,omitempty"`                         // Used memory in bytes.
	SslFingerprint *string  `url:"ssl_fingerprint,omitempty" json:"ssl_fingerprint,omitempty"` // The SSL fingerprint for the node certificate.
	Uptime         *int     `url:"uptime,omitempty" json:"uptime,omitempty"`                   // Node uptime in seconds.
}
type _IndexResponse IndexResponse

type FindRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _FindRequest FindRequest

type GetSubscriptionRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _GetSubscriptionRequest GetSubscriptionRequest

type UpdateSubscriptionRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Force *util.PVEBool `url:"force,omitempty" json:"force,omitempty"` // Always connect to server, even if we have up to date info inside local cache.
}
type _UpdateSubscriptionRequest UpdateSubscriptionRequest

type SetSubscriptionRequest struct {
	Key  string `url:"key" json:"key"`   // Proxmox VE subscription key
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _SetSubscriptionRequest SetSubscriptionRequest

type DeleteSubscriptionRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _DeleteSubscriptionRequest DeleteSubscriptionRequest

type GetConfigRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Property *Property `url:"property,omitempty" json:"property,omitempty"` // Return only a specific property from the node configuration.
}
type _GetConfigRequest GetConfigRequest

// Node specific ACME settings.
type Acme struct {

	// The following parameters are optional
	Account *string `url:"account,omitempty" json:"account,omitempty"` // ACME account config file name.
	Domains *string `url:"domains,omitempty" json:"domains,omitempty"` // List of domains for this node's ACME certificate
}
type _Acme Acme

func (t Acme) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[account=<name>] [,domains=<domain[;domain;...]>]`)
}

func (t *Acme) UnmarshalJSON(d []byte) error {
	if len(d) == 0 || string(d) == `""` {
		return nil
	}
	cleaned := string(d)[1 : len(d)-1]
	parts := strings.Split(cleaned, ",")
	values := map[string]string{}
	for _, p := range parts {
		kv := strings.Split(p, "=")
		if len(kv) > 2 {
			return fmt.Errorf("Wrong number of parts for kv pair '%s'", p)
		}
		if len(kv) == 1 {

			values["account"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["account"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Account)
		if err != nil {
			return err
		}
	}

	if v, ok := values["domains"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Domains)
		if err != nil {
			return err
		}
	}

	return nil
}

// ACME domain and validation plugin
type Acmedomain struct {
	Domain string `url:"domain" json:"domain"` // domain for this node's ACME certificate

	// The following parameters are optional
	Alias  *string `url:"alias,omitempty" json:"alias,omitempty"`   // Alias for the Domain to verify ACME Challenge over DNS
	Plugin *string `url:"plugin,omitempty" json:"plugin,omitempty"` // The ACME plugin ID
}
type _Acmedomain Acmedomain

func (t Acmedomain) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[domain=]<domain> [,alias=<domain>] [,plugin=<name of the plugin configuration>]`)
}

func (t *Acmedomain) UnmarshalJSON(d []byte) error {
	if len(d) == 0 || string(d) == `""` {
		return nil
	}
	cleaned := string(d)[1 : len(d)-1]
	parts := strings.Split(cleaned, ",")
	values := map[string]string{}
	for _, p := range parts {
		kv := strings.Split(p, "=")
		if len(kv) > 2 {
			return fmt.Errorf("Wrong number of parts for kv pair '%s'", p)
		}
		if len(kv) == 1 {

			values["domain"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["domain"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Domain)
		if err != nil {
			return err
		}
	}

	if v, ok := values["alias"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Alias)
		if err != nil {
			return err
		}
	}

	if v, ok := values["plugin"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Plugin)
		if err != nil {
			return err
		}
	}

	return nil
}

// Array of Acmedomain
type Acmedomains []*Acmedomain
type _Acmedomains Acmedomains

func (t Acmedomains) EncodeValues(key string, v *url.Values) error {
	return util.EncodeArray(key, v, t)
}

type GetConfigResponse struct {

	// The following parameters are optional
	Acme                *Acme        `url:"acme,omitempty" json:"acme,omitempty"`                                   // Node specific ACME settings.
	Acmedomains         *Acmedomains `url:"acmedomain[n],omitempty" json:"acmedomain[n],omitempty"`                 // ACME domain and validation plugin
	Description         *string      `url:"description,omitempty" json:"description,omitempty"`                     // Description for the Node. Shown in the web-interface node notes panel. This is saved as comment inside the configuration file.
	Digest              *string      `url:"digest,omitempty" json:"digest,omitempty"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	StartallOnbootDelay *int         `url:"startall-onboot-delay,omitempty" json:"startall-onboot-delay,omitempty"` // Initial delay in seconds, before starting all the Virtual Guests with on-boot enabled.
	Wakeonlan           *string      `url:"wakeonlan,omitempty" json:"wakeonlan,omitempty"`                         // MAC address for wake on LAN
}
type _GetConfigResponse GetConfigResponse

func (t *GetConfigResponse) UnmarshalJSON(d []byte) error {
	tmp := _GetConfigResponse{}
	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}
	*t = GetConfigResponse(tmp)
	rest := map[string]json.RawMessage{}
	err = json.Unmarshal(d, &rest)
	if err != nil {
		return err
	}
	for k, v := range rest {

		if strings.HasPrefix(k, "acmedomain") {
			idxStr := strings.TrimPrefix(k, "acmedomain")
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return err
			}
			if t.Acmedomains == nil {
				arr := make(Acmedomains, 0)
				t.Acmedomains = &arr
			}
			for len(*t.Acmedomains) < idx+1 {
				*t.Acmedomains = append(*t.Acmedomains, nil)
			}
			var newVal Acmedomain
			err = json.Unmarshal(v, &newVal)
			if err != nil {
				return err
			}
			(*t.Acmedomains)[idx] = &newVal
		}

	}
	return nil
}

type SetOptionsConfigRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Acme                *Acme        `url:"acme,omitempty" json:"acme,omitempty"`                                   // Node specific ACME settings.
	Acmedomains         *Acmedomains `url:"acmedomain[n],omitempty" json:"acmedomain[n],omitempty"`                 // ACME domain and validation plugin
	Delete              *string      `url:"delete,omitempty" json:"delete,omitempty"`                               // A list of settings you want to delete.
	Description         *string      `url:"description,omitempty" json:"description,omitempty"`                     // Description for the Node. Shown in the web-interface node notes panel. This is saved as comment inside the configuration file.
	Digest              *string      `url:"digest,omitempty" json:"digest,omitempty"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	StartallOnbootDelay *int         `url:"startall-onboot-delay,omitempty" json:"startall-onboot-delay,omitempty"` // Initial delay in seconds, before starting all the Virtual Guests with on-boot enabled.
	Wakeonlan           *string      `url:"wakeonlan,omitempty" json:"wakeonlan,omitempty"`                         // MAC address for wake on LAN
}
type _SetOptionsConfigRequest SetOptionsConfigRequest

func (t *SetOptionsConfigRequest) UnmarshalJSON(d []byte) error {
	tmp := _SetOptionsConfigRequest{}
	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}
	*t = SetOptionsConfigRequest(tmp)
	rest := map[string]json.RawMessage{}
	err = json.Unmarshal(d, &rest)
	if err != nil {
		return err
	}
	for k, v := range rest {

		if strings.HasPrefix(k, "acmedomain") {
			idxStr := strings.TrimPrefix(k, "acmedomain")
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return err
			}
			if t.Acmedomains == nil {
				arr := make(Acmedomains, 0)
				t.Acmedomains = &arr
			}
			for len(*t.Acmedomains) < idx+1 {
				*t.Acmedomains = append(*t.Acmedomains, nil)
			}
			var newVal Acmedomain
			err = json.Unmarshal(v, &newVal)
			if err != nil {
				return err
			}
			(*t.Acmedomains)[idx] = &newVal
		}

	}
	return nil
}

type VersionRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _VersionRequest VersionRequest

type VersionResponse struct {
	Release string `url:"release" json:"release"` // The current installed Proxmox VE Release
	Repoid  string `url:"repoid" json:"repoid"`   // The short git commit hash ID from which this version was build
	Version string `url:"version" json:"version"` // The current installed pve-manager package version

}
type _VersionResponse VersionResponse

type StatusRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _StatusRequest StatusRequest

type NodeCmdStatusRequest struct {
	Command Command `url:"command" json:"command"` // Specify the command.
	Node    string  `url:"node" json:"node"`       // The cluster node name.

}
type _NodeCmdStatusRequest NodeCmdStatusRequest

type NetstatRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _NetstatRequest NetstatRequest

type ExecuteRequest struct {
	Commands string `url:"commands" json:"commands"` // JSON encoded array of commands.
	Node     string `url:"node" json:"node"`         // The cluster node name.

}
type _ExecuteRequest ExecuteRequest

type WakeonlanRequest struct {
	Node string `url:"node" json:"node"` // target node for wake on LAN packet

}
type _WakeonlanRequest WakeonlanRequest

type RrdRequest struct {
	Ds        string    `url:"ds" json:"ds"`               // The list of datasources you want to display.
	Node      string    `url:"node" json:"node"`           // The cluster node name.
	Timeframe Timeframe `url:"timeframe" json:"timeframe"` // Specify the time frame you are interested in.

	// The following parameters are optional
	Cf *Cf `url:"cf,omitempty" json:"cf,omitempty"` // The RRD consolidation function
}
type _RrdRequest RrdRequest

type RrdResponse struct {
	Filename string `url:"filename" json:"filename"`
}
type _RrdResponse RrdResponse

type RrddataRequest struct {
	Node      string    `url:"node" json:"node"`           // The cluster node name.
	Timeframe Timeframe `url:"timeframe" json:"timeframe"` // Specify the time frame you are interested in.

	// The following parameters are optional
	Cf *Cf `url:"cf,omitempty" json:"cf,omitempty"` // The RRD consolidation function
}
type _RrddataRequest RrddataRequest

type SyslogRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Limit   *int    `url:"limit,omitempty" json:"limit,omitempty"`
	Service *string `url:"service,omitempty" json:"service,omitempty"` // Service ID
	Since   *string `url:"since,omitempty" json:"since,omitempty"`     // Display all log since this date-time string.
	Start   *int    `url:"start,omitempty" json:"start,omitempty"`
	Until   *string `url:"until,omitempty" json:"until,omitempty"` // Display all log until this date-time string.
}
type _SyslogRequest SyslogRequest

type SyslogResponse struct {
	N int    `url:"n" json:"n"` // Line number
	T string `url:"t" json:"t"` // Line text

}
type _SyslogResponse SyslogResponse

type JournalRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Endcursor   *string `url:"endcursor,omitempty" json:"endcursor,omitempty"`     // End before the given Cursor. Conflicts with 'until'
	Lastentries *int    `url:"lastentries,omitempty" json:"lastentries,omitempty"` // Limit to the last X lines. Conflicts with a range.
	Since       *int    `url:"since,omitempty" json:"since,omitempty"`             // Display all log since this UNIX epoch. Conflicts with 'startcursor'.
	Startcursor *string `url:"startcursor,omitempty" json:"startcursor,omitempty"` // Start after the given Cursor. Conflicts with 'since'
	Until       *int    `url:"until,omitempty" json:"until,omitempty"`             // Display all log until this UNIX epoch. Conflicts with 'endcursor'.
}
type _JournalRequest JournalRequest

type VncshellRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Cmd       *Cmd          `url:"cmd,omitempty" json:"cmd,omitempty"`             // Run specific command or default to login.
	CmdOpts   *string       `url:"cmd-opts,omitempty" json:"cmd-opts,omitempty"`   // Add parameters to a command. Encoded as null terminated strings.
	Height    *int          `url:"height,omitempty" json:"height,omitempty"`       // sets the height of the console in pixels.
	Websocket *util.PVEBool `url:"websocket,omitempty" json:"websocket,omitempty"` // use websocket instead of standard vnc.
	Width     *int          `url:"width,omitempty" json:"width,omitempty"`         // sets the width of the console in pixels.
}
type _VncshellRequest VncshellRequest

type VncshellResponse struct {
	Cert   string `url:"cert" json:"cert"`
	Port   int    `url:"port" json:"port"`
	Ticket string `url:"ticket" json:"ticket"`
	Upid   string `url:"upid" json:"upid"`
	User   string `url:"user" json:"user"`
}
type _VncshellResponse VncshellResponse

type TermproxyRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Cmd     *Cmd    `url:"cmd,omitempty" json:"cmd,omitempty"`           // Run specific command or default to login.
	CmdOpts *string `url:"cmd-opts,omitempty" json:"cmd-opts,omitempty"` // Add parameters to a command. Encoded as null terminated strings.
}
type _TermproxyRequest TermproxyRequest

type TermproxyResponse struct {
	Port   int    `url:"port" json:"port"`
	Ticket string `url:"ticket" json:"ticket"`
	Upid   string `url:"upid" json:"upid"`
	User   string `url:"user" json:"user"`
}
type _TermproxyResponse TermproxyResponse

type VncwebsocketRequest struct {
	Node      string `url:"node" json:"node"`           // The cluster node name.
	Port      int    `url:"port" json:"port"`           // Port number returned by previous vncproxy call.
	Vncticket string `url:"vncticket" json:"vncticket"` // Ticket from previous call to vncproxy.

}
type _VncwebsocketRequest VncwebsocketRequest

type VncwebsocketResponse struct {
	Port string `url:"port" json:"port"`
}
type _VncwebsocketResponse VncwebsocketResponse

type SpiceshellRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Cmd     *Cmd    `url:"cmd,omitempty" json:"cmd,omitempty"`           // Run specific command or default to login.
	CmdOpts *string `url:"cmd-opts,omitempty" json:"cmd-opts,omitempty"` // Add parameters to a command. Encoded as null terminated strings.
	Proxy   *string `url:"proxy,omitempty" json:"proxy,omitempty"`       // SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
}
type _SpiceshellRequest SpiceshellRequest

// Returned values can be directly passed to the 'remote-viewer' application.
type SpiceshellResponse struct {
	Host     string `url:"host" json:"host"`
	Password string `url:"password" json:"password"`
	Proxy    string `url:"proxy" json:"proxy"`
	TlsPort  int    `url:"tls-port" json:"tls-port"`
	Type     string `url:"type" json:"type"`
}
type _SpiceshellResponse SpiceshellResponse

type DnsRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _DnsRequest DnsRequest

type DnsResponse struct {

	// The following parameters are optional
	Dns1   *string `url:"dns1,omitempty" json:"dns1,omitempty"`     // First name server IP address.
	Dns2   *string `url:"dns2,omitempty" json:"dns2,omitempty"`     // Second name server IP address.
	Dns3   *string `url:"dns3,omitempty" json:"dns3,omitempty"`     // Third name server IP address.
	Search *string `url:"search,omitempty" json:"search,omitempty"` // Search domain for host-name lookup.
}
type _DnsResponse DnsResponse

type UpdateDnsRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Search string `url:"search" json:"search"` // Search domain for host-name lookup.

	// The following parameters are optional
	Dns1 *string `url:"dns1,omitempty" json:"dns1,omitempty"` // First name server IP address.
	Dns2 *string `url:"dns2,omitempty" json:"dns2,omitempty"` // Second name server IP address.
	Dns3 *string `url:"dns3,omitempty" json:"dns3,omitempty"` // Third name server IP address.
}
type _UpdateDnsRequest UpdateDnsRequest

type TimeRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _TimeRequest TimeRequest

type TimeResponse struct {
	Localtime int    `url:"localtime" json:"localtime"` // Seconds since 1970-01-01 00:00:00 (local time)
	Time      int    `url:"time" json:"time"`           // Seconds since 1970-01-01 00:00:00 UTC.
	Timezone  string `url:"timezone" json:"timezone"`   // Time zone

}
type _TimeResponse TimeResponse

type SetTimezoneTimeRequest struct {
	Node     string `url:"node" json:"node"`         // The cluster node name.
	Timezone string `url:"timezone" json:"timezone"` // Time zone. The file '/usr/share/zoneinfo/zone.tab' contains the list of valid names.

}
type _SetTimezoneTimeRequest SetTimezoneTimeRequest

type AplinfoRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _AplinfoRequest AplinfoRequest

type AplDownloadAplinfoRequest struct {
	Node     string `url:"node" json:"node"`         // The cluster node name.
	Storage  string `url:"storage" json:"storage"`   // The storage where the template will be stored
	Template string `url:"template" json:"template"` // The template which will downloaded

}
type _AplDownloadAplinfoRequest AplDownloadAplinfoRequest

type QueryUrlMetadataRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.
	Url  string `url:"url" json:"url"`   // The URL to query the metadata from.

	// The following parameters are optional
	VerifyCertificates *util.PVEBool `url:"verify-certificates,omitempty" json:"verify-certificates,omitempty"` // If false, no SSL/TLS certificates will be verified.
}
type _QueryUrlMetadataRequest QueryUrlMetadataRequest

type QueryUrlMetadataResponse struct {

	// The following parameters are optional
	Filename *string `url:"filename,omitempty" json:"filename,omitempty"`
	Mimetype *string `url:"mimetype,omitempty" json:"mimetype,omitempty"`
	Size     *int    `url:"size,omitempty" json:"size,omitempty"`
}
type _QueryUrlMetadataResponse QueryUrlMetadataResponse

type ReportRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _ReportRequest ReportRequest

type StartallRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Force *util.PVEBool `url:"force,omitempty" json:"force,omitempty"` // Issue start command even if virtual guest have 'onboot' not set or set to off.
	Vms   *string       `url:"vms,omitempty" json:"vms,omitempty"`     // Only consider guests from this comma separated list of VMIDs.
}
type _StartallRequest StartallRequest

type StopallRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Vms *string `url:"vms,omitempty" json:"vms,omitempty"` // Only consider Guests with these IDs.
}
type _StopallRequest StopallRequest

type MigrateallRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Target string `url:"target" json:"target"` // Target node.

	// The following parameters are optional
	Maxworkers     *int          `url:"maxworkers,omitempty" json:"maxworkers,omitempty"`             // Maximal number of parallel migration job. If not set use 'max_workers' from datacenter.cfg, one of both must be set!
	Vms            *string       `url:"vms,omitempty" json:"vms,omitempty"`                           // Only consider Guests with these IDs.
	WithLocalDisks *util.PVEBool `url:"with-local-disks,omitempty" json:"with-local-disks,omitempty"` // Enable live storage migration for local disk
}
type _MigrateallRequest MigrateallRequest

type GetEtcHostsRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _GetEtcHostsRequest GetEtcHostsRequest

type GetEtcHostsResponse struct {
	Data string `url:"data" json:"data"` // The content of /etc/hosts.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}
type _GetEtcHostsResponse GetEtcHostsResponse

type WriteEtcHostsRequest struct {
	Data string `url:"data" json:"data"` // The target content of /etc/hosts.
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Digest *string `url:"digest,omitempty" json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}
type _WriteEtcHostsRequest WriteEtcHostsRequest

// Index Cluster node index.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/nodes", "GET", &resp, nil)
	return resp, err
}

// Find Node index.
func (c *Client) Find(ctx context.Context, req FindRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}", "GET", &resp, req)
	return resp, err
}

// GetSubscription Read subscription info.
func (c *Client) GetSubscription(ctx context.Context, req GetSubscriptionRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/subscription", "GET", &resp, req)
	return resp, err
}

// UpdateSubscription Update subscription info.
func (c *Client) UpdateSubscription(ctx context.Context, req UpdateSubscriptionRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/subscription", "POST", nil, req)
	return err
}

// SetSubscription Set subscription key.
func (c *Client) SetSubscription(ctx context.Context, req SetSubscriptionRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/subscription", "PUT", nil, req)
	return err
}

// DeleteSubscription Delete subscription key of this node.
func (c *Client) DeleteSubscription(ctx context.Context, req DeleteSubscriptionRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/subscription", "DELETE", nil, req)
	return err
}

// GetConfig Get node configuration options.
func (c *Client) GetConfig(ctx context.Context, req GetConfigRequest) (GetConfigResponse, error) {
	var resp GetConfigResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/config", "GET", &resp, req)
	return resp, err
}

// SetOptionsConfig Set node configuration options.
func (c *Client) SetOptionsConfig(ctx context.Context, req SetOptionsConfigRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/config", "PUT", nil, req)
	return err
}

// Version API version details
func (c *Client) Version(ctx context.Context, req VersionRequest) (VersionResponse, error) {
	var resp VersionResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/version", "GET", &resp, req)
	return resp, err
}

// Status Read node status
func (c *Client) Status(ctx context.Context, req StatusRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/status", "GET", &resp, req)
	return resp, err
}

// NodeCmdStatus Reboot or shutdown a node.
func (c *Client) NodeCmdStatus(ctx context.Context, req NodeCmdStatusRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/status", "POST", nil, req)
	return err
}

// Netstat Read tap/vm network device interface counters
func (c *Client) Netstat(ctx context.Context, req NetstatRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/netstat", "GET", &resp, req)
	return resp, err
}

// Execute Execute multiple commands in order.
func (c *Client) Execute(ctx context.Context, req ExecuteRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/execute", "POST", &resp, req)
	return resp, err
}

// Wakeonlan Try to wake a node via 'wake on LAN' network packet.
func (c *Client) Wakeonlan(ctx context.Context, req WakeonlanRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/wakeonlan", "POST", &resp, req)
	return resp, err
}

// Rrd Read node RRD statistics (returns PNG)
func (c *Client) Rrd(ctx context.Context, req RrdRequest) (RrdResponse, error) {
	var resp RrdResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/rrd", "GET", &resp, req)
	return resp, err
}

// Rrddata Read node RRD statistics
func (c *Client) Rrddata(ctx context.Context, req RrddataRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/rrddata", "GET", &resp, req)
	return resp, err
}

// Syslog Read system log
func (c *Client) Syslog(ctx context.Context, req SyslogRequest) ([]SyslogResponse, error) {
	var resp []SyslogResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/syslog", "GET", &resp, req)
	return resp, err
}

// Journal Read Journal
func (c *Client) Journal(ctx context.Context, req JournalRequest) ([]string, error) {
	var resp []string

	err := c.httpClient.Do(ctx, "/nodes/{node}/journal", "GET", &resp, req)
	return resp, err
}

// Vncshell Creates a VNC Shell proxy.
func (c *Client) Vncshell(ctx context.Context, req VncshellRequest) (VncshellResponse, error) {
	var resp VncshellResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/vncshell", "POST", &resp, req)
	return resp, err
}

// Termproxy Creates a VNC Shell proxy.
func (c *Client) Termproxy(ctx context.Context, req TermproxyRequest) (TermproxyResponse, error) {
	var resp TermproxyResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/termproxy", "POST", &resp, req)
	return resp, err
}

// Vncwebsocket Opens a websocket for VNC traffic.
func (c *Client) Vncwebsocket(ctx context.Context, req VncwebsocketRequest) (VncwebsocketResponse, error) {
	var resp VncwebsocketResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/vncwebsocket", "GET", &resp, req)
	return resp, err
}

// Spiceshell Creates a SPICE shell.
func (c *Client) Spiceshell(ctx context.Context, req SpiceshellRequest) (SpiceshellResponse, error) {
	var resp SpiceshellResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/spiceshell", "POST", &resp, req)
	return resp, err
}

// Dns Read DNS settings.
func (c *Client) Dns(ctx context.Context, req DnsRequest) (DnsResponse, error) {
	var resp DnsResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/dns", "GET", &resp, req)
	return resp, err
}

// UpdateDns Write DNS settings.
func (c *Client) UpdateDns(ctx context.Context, req UpdateDnsRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/dns", "PUT", nil, req)
	return err
}

// Time Read server time and time zone settings.
func (c *Client) Time(ctx context.Context, req TimeRequest) (TimeResponse, error) {
	var resp TimeResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/time", "GET", &resp, req)
	return resp, err
}

// SetTimezoneTime Set time zone.
func (c *Client) SetTimezoneTime(ctx context.Context, req SetTimezoneTimeRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/time", "PUT", nil, req)
	return err
}

// Aplinfo Get list of appliances.
func (c *Client) Aplinfo(ctx context.Context, req AplinfoRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/nodes/{node}/aplinfo", "GET", &resp, req)
	return resp, err
}

// AplDownloadAplinfo Download appliance templates.
func (c *Client) AplDownloadAplinfo(ctx context.Context, req AplDownloadAplinfoRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/aplinfo", "POST", &resp, req)
	return resp, err
}

// QueryUrlMetadata Query metadata of an URL: file size, file name and mime type.
func (c *Client) QueryUrlMetadata(ctx context.Context, req QueryUrlMetadataRequest) (QueryUrlMetadataResponse, error) {
	var resp QueryUrlMetadataResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/query-url-metadata", "GET", &resp, req)
	return resp, err
}

// Report Gather various systems information about a node
func (c *Client) Report(ctx context.Context, req ReportRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/report", "GET", &resp, req)
	return resp, err
}

// Startall Start all VMs and containers located on this node (by default only those with onboot=1).
func (c *Client) Startall(ctx context.Context, req StartallRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/startall", "POST", &resp, req)
	return resp, err
}

// Stopall Stop all VMs and Containers.
func (c *Client) Stopall(ctx context.Context, req StopallRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/stopall", "POST", &resp, req)
	return resp, err
}

// Migrateall Migrate all VMs and Containers.
func (c *Client) Migrateall(ctx context.Context, req MigrateallRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/nodes/{node}/migrateall", "POST", &resp, req)
	return resp, err
}

// GetEtcHosts Get the content of /etc/hosts.
func (c *Client) GetEtcHosts(ctx context.Context, req GetEtcHostsRequest) (GetEtcHostsResponse, error) {
	var resp GetEtcHostsResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/hosts", "GET", &resp, req)
	return resp, err
}

// WriteEtcHosts Write /etc/hosts.
func (c *Client) WriteEtcHosts(ctx context.Context, req WriteEtcHostsRequest) error {

	err := c.httpClient.Do(ctx, "/nodes/{node}/hosts", "POST", nil, req)
	return err
}

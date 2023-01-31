// File generated by proxmox json schema, DO NOT EDIT

package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
	"net/url"
	"strings"
)

const (
	Console_APPLET  Console = "applet"
	Console_VV      Console = "vv"
	Console_HTML5   Console = "html5"
	Console_XTERMJS Console = "xtermjs"

	CrsHa_BASIC  CrsHa = "basic"
	CrsHa_STATIC CrsHa = "static"

	Fencing_WATCHDOG Fencing = "watchdog"
	Fencing_HARDWARE Fencing = "hardware"
	Fencing_BOTH     Fencing = "both"

	HaShutdownPolicy_FREEZE      HaShutdownPolicy = "freeze"
	HaShutdownPolicy_FAILOVER    HaShutdownPolicy = "failover"
	HaShutdownPolicy_CONDITIONAL HaShutdownPolicy = "conditional"
	HaShutdownPolicy_MIGRATE     HaShutdownPolicy = "migrate"

	Keyboard_DE    Keyboard = "de"
	Keyboard_DE_CH Keyboard = "de-ch"
	Keyboard_DA    Keyboard = "da"
	Keyboard_EN_GB Keyboard = "en-gb"
	Keyboard_EN_US Keyboard = "en-us"
	Keyboard_ES    Keyboard = "es"
	Keyboard_FI    Keyboard = "fi"
	Keyboard_FR    Keyboard = "fr"
	Keyboard_FR_BE Keyboard = "fr-be"
	Keyboard_FR_CA Keyboard = "fr-ca"
	Keyboard_FR_CH Keyboard = "fr-ch"
	Keyboard_HU    Keyboard = "hu"
	Keyboard_IS    Keyboard = "is"
	Keyboard_IT    Keyboard = "it"
	Keyboard_JA    Keyboard = "ja"
	Keyboard_LT    Keyboard = "lt"
	Keyboard_MK    Keyboard = "mk"
	Keyboard_NL    Keyboard = "nl"
	Keyboard_NO    Keyboard = "no"
	Keyboard_PL    Keyboard = "pl"
	Keyboard_PT    Keyboard = "pt"
	Keyboard_PT_BR Keyboard = "pt-br"
	Keyboard_SV    Keyboard = "sv"
	Keyboard_SL    Keyboard = "sl"
	Keyboard_TR    Keyboard = "tr"

	Language_CA    Language = "ca"
	Language_DA    Language = "da"
	Language_DE    Language = "de"
	Language_EN    Language = "en"
	Language_ES    Language = "es"
	Language_EU    Language = "eu"
	Language_FA    Language = "fa"
	Language_FR    Language = "fr"
	Language_HE    Language = "he"
	Language_IT    Language = "it"
	Language_JA    Language = "ja"
	Language_NB    Language = "nb"
	Language_NN    Language = "nn"
	Language_PL    Language = "pl"
	Language_PT_BR Language = "pt_BR"
	Language_RU    Language = "ru"
	Language_SL    Language = "sl"
	Language_SV    Language = "sv"
	Language_TR    Language = "tr"
	Language_ZH_CN Language = "zh_CN"
	Language_ZH_TW Language = "zh_TW"

	MigrationType_SECURE   MigrationType = "secure"
	MigrationType_INSECURE MigrationType = "insecure"

	TagStyleOrdering_CONFIG       TagStyleOrdering = "config"
	TagStyleOrdering_ALPHABETICAL TagStyleOrdering = "alphabetical"

	TagStyleShape_FULL   TagStyleShape = "full"
	TagStyleShape_CIRCLE TagStyleShape = "circle"
	TagStyleShape_DENSE  TagStyleShape = "dense"
	TagStyleShape_NONE   TagStyleShape = "none"

	Type_VM      Type = "vm"
	Type_STORAGE Type = "storage"
	Type_NODE    Type = "node"
	Type_SDN     Type = "sdn"
	Type_POOL    Type = "pool"
	Type_QEMU    Type = "qemu"
	Type_LXC     Type = "lxc"
	Type_OPENVZ  Type = "openvz"
	Type_CLUSTER Type = "cluster"

	UserTagAccessUserAllow_NONE     UserTagAccessUserAllow = "none"
	UserTagAccessUserAllow_LIST     UserTagAccessUserAllow = "list"
	UserTagAccessUserAllow_EXISTING UserTagAccessUserAllow = "existing"
	UserTagAccessUserAllow_FREE     UserTagAccessUserAllow = "free"
)

type Console string
type CrsHa string
type Fencing string
type HaShutdownPolicy string
type Keyboard string
type Language string
type MigrationType string
type TagStyleOrdering string
type TagStyleShape string
type Type string
type UserTagAccessUserAllow string

func PtrConsole(i Console) *Console {
	return &i
}
func PtrCrsHa(i CrsHa) *CrsHa {
	return &i
}
func PtrFencing(i Fencing) *Fencing {
	return &i
}
func PtrHaShutdownPolicy(i HaShutdownPolicy) *HaShutdownPolicy {
	return &i
}
func PtrKeyboard(i Keyboard) *Keyboard {
	return &i
}
func PtrLanguage(i Language) *Language {
	return &i
}
func PtrMigrationType(i MigrationType) *MigrationType {
	return &i
}
func PtrTagStyleOrdering(i TagStyleOrdering) *TagStyleOrdering {
	return &i
}
func PtrTagStyleShape(i TagStyleShape) *TagStyleShape {
	return &i
}
func PtrType(i Type) *Type {
	return &i
}
func PtrUserTagAccessUserAllow(i UserTagAccessUserAllow) *UserTagAccessUserAllow {
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

type LogRequest struct {

	// The following parameters are optional
	Max *int `url:"max,omitempty" json:"max,omitempty"` // Maximum number of entries.
}
type _LogRequest LogRequest

type ResourcesRequest struct {

	// The following parameters are optional
	Type *Type `url:"type,omitempty" json:"type,omitempty"`
}
type _ResourcesRequest ResourcesRequest

type ResourcesResponse struct {
	Id   string `url:"id" json:"id"`
	Type Type   `url:"type" json:"type"` // Resource type.

	// The following parameters are optional
	CgroupMode *int     `url:"cgroup-mode,omitempty" json:"cgroup-mode,omitempty"` // The cgroup mode the node operates under (when type == node).
	Content    *string  `url:"content,omitempty" json:"content,omitempty"`         // Allowed storage content types (when type == storage).
	Cpu        *float64 `url:"cpu,omitempty" json:"cpu,omitempty"`                 // CPU utilization (when type in node,qemu,lxc).
	Disk       *int     `url:"disk,omitempty" json:"disk,omitempty"`               // Used disk space in bytes (when type in storage), used root image spave for VMs (type in qemu,lxc).
	Hastate    *string  `url:"hastate,omitempty" json:"hastate,omitempty"`         // HA service status (for HA managed VMs).
	Level      *string  `url:"level,omitempty" json:"level,omitempty"`             // Support level (when type == node).
	Maxcpu     *float64 `url:"maxcpu,omitempty" json:"maxcpu,omitempty"`           // Number of available CPUs (when type in node,qemu,lxc).
	Maxdisk    *int     `url:"maxdisk,omitempty" json:"maxdisk,omitempty"`         // Storage size in bytes (when type in storage), root image size for VMs (type in qemu,lxc).
	Maxmem     *int     `url:"maxmem,omitempty" json:"maxmem,omitempty"`           // Number of available memory in bytes (when type in node,qemu,lxc).
	Mem        *int     `url:"mem,omitempty" json:"mem,omitempty"`                 // Used memory in bytes (when type in node,qemu,lxc).
	Name       *string  `url:"name,omitempty" json:"name,omitempty"`               // Name of the resource.
	Node       *string  `url:"node,omitempty" json:"node,omitempty"`               // The cluster node name (when type in node,storage,qemu,lxc).
	Plugintype *string  `url:"plugintype,omitempty" json:"plugintype,omitempty"`   // More specific type, if available.
	Pool       *string  `url:"pool,omitempty" json:"pool,omitempty"`               // The pool name (when type in pool,qemu,lxc).
	Status     *string  `url:"status,omitempty" json:"status,omitempty"`           // Resource type dependent status.
	Storage    *string  `url:"storage,omitempty" json:"storage,omitempty"`         // The storage identifier (when type == storage).
	Uptime     *int     `url:"uptime,omitempty" json:"uptime,omitempty"`           // Node uptime in seconds (when type in node,qemu,lxc).
	Vmid       *int     `url:"vmid,omitempty" json:"vmid,omitempty"`               // The numerical vmid (when type in qemu,lxc).
}
type _ResourcesResponse ResourcesResponse

type TasksResponse struct {
	Upid string `url:"upid" json:"upid"`
}
type _TasksResponse TasksResponse

// Cluster wide HA settings.
type Ha struct {
	ShutdownPolicy HaShutdownPolicy `url:"shutdown_policy" json:"shutdown_policy"` // The policy for HA services on node shutdown. 'freeze' disables auto-recovery, 'failover' ensures recovery, 'conditional' recovers on poweroff and freezes on reboot. 'migrate' will migrate running services to other nodes, if possible. With 'freeze' or 'failover', HA Services will always get stopped first on shutdown.

}
type _Ha Ha

func (t Ha) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `shutdown_policy=<enum>`)
}

func (t *Ha) UnmarshalJSON(d []byte) error {
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

			values["shutdown_policy"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["shutdown_policy"]; ok {

		err := json.Unmarshal([]byte(v), &t.ShutdownPolicy)
		if err != nil {
			return err
		}
	}

	return nil
}

// Tag style options.
type TagStyle struct {

	// The following parameters are optional
	CaseSensitive *util.PVEBool     `url:"case-sensitive,omitempty" json:"case-sensitive,omitempty"` // Controls if filtering for unique tags on update should check case-sensitive.
	ColorMap      *string           `url:"color-map,omitempty" json:"color-map,omitempty"`           // Manual color mapping for tags (semicolon separated).
	Ordering      *TagStyleOrdering `url:"ordering,omitempty" json:"ordering,omitempty"`             // Controls the sorting of the tags in the web-interface and the API update.
	Shape         *TagStyleShape    `url:"shape,omitempty" json:"shape,omitempty"`                   // Tag shape for the web ui tree. 'full' draws the full tag. 'circle' draws only a circle with the background color. 'dense' only draws a small rectancle (useful when many tags are assigned to each guest).'none' disables showing the tags.
}
type _TagStyle TagStyle

func (t TagStyle) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[case-sensitive=<1|0>] [,color-map=<tag>:<hex-color>[:<hex-color-for-text>][;<tag>=...]] [,ordering=<config|alphabetical>] [,shape=<enum>]`)
}

func (t *TagStyle) UnmarshalJSON(d []byte) error {
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

			values["case-sensitive"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["case-sensitive"]; ok {

		err := json.Unmarshal([]byte(v), &t.CaseSensitive)
		if err != nil {
			return err
		}
	}

	if v, ok := values["color-map"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.ColorMap)
		if err != nil {
			return err
		}
	}

	if v, ok := values["ordering"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Ordering)
		if err != nil {
			return err
		}
	}

	if v, ok := values["shape"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Shape)
		if err != nil {
			return err
		}
	}

	return nil
}

// webauthn configuration
type Webauthn struct {

	// The following parameters are optional
	AllowSubdomains *util.PVEBool `url:"allow-subdomains,omitempty" json:"allow-subdomains,omitempty"` // Whether to allow the origin to be a subdomain, rather than the exact URL.
	Id              *string       `url:"id,omitempty" json:"id,omitempty"`                             // Relying party ID. Must be the domain name without protocol, port or location. Changing this *will* break existing credentials.
	Origin          *string       `url:"origin,omitempty" json:"origin,omitempty"`                     // Site origin. Must be a `https://` URL (or `http://localhost`). Should contain the address users type in their browsers to access the web interface. Changing this *may* break existing credentials.
	Rp              *string       `url:"rp,omitempty" json:"rp,omitempty"`                             // Relying party name. Any text identifier. Changing this *may* break existing credentials.
}
type _Webauthn Webauthn

func (t Webauthn) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[allow-subdomains=<1|0>] [,id=<DOMAINNAME>] [,origin=<URL>] [,rp=<RELYING_PARTY>]`)
}

func (t *Webauthn) UnmarshalJSON(d []byte) error {
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

			values["allow-subdomains"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["allow-subdomains"]; ok {

		err := json.Unmarshal([]byte(v), &t.AllowSubdomains)
		if err != nil {
			return err
		}
	}

	if v, ok := values["id"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Id)
		if err != nil {
			return err
		}
	}

	if v, ok := values["origin"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Origin)
		if err != nil {
			return err
		}
	}

	if v, ok := values["rp"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Rp)
		if err != nil {
			return err
		}
	}

	return nil
}

// Cluster resource scheduling settings.
type Crs struct {
	Ha CrsHa `url:"ha" json:"ha"` // Use this resource scheduler mode for HA.

}
type _Crs Crs

func (t Crs) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `ha=<basic|static>`)
}

func (t *Crs) UnmarshalJSON(d []byte) error {
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

			values["ha"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["ha"]; ok {

		err := json.Unmarshal([]byte(v), &t.Ha)
		if err != nil {
			return err
		}
	}

	return nil
}

// Control the range for the free VMID auto-selection pool.
type NextId struct {

	// The following parameters are optional
	Lower *int `url:"lower,omitempty" json:"lower,omitempty"` // Lower, inclusive boundary for free next-id API range.
	Upper *int `url:"upper,omitempty" json:"upper,omitempty"` // Upper, exclusive boundary for free next-id API range.
}
type _NextId NextId

func (t NextId) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[lower=<integer>] [,upper=<integer>]`)
}

func (t *NextId) UnmarshalJSON(d []byte) error {
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

			values["lower"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["lower"]; ok {

		err := json.Unmarshal([]byte(v), &t.Lower)
		if err != nil {
			return err
		}
	}

	if v, ok := values["upper"]; ok {

		err := json.Unmarshal([]byte(v), &t.Upper)
		if err != nil {
			return err
		}
	}

	return nil
}

// Privilege options for user-settable tags
type UserTagAccess struct {

	// The following parameters are optional
	UserAllow     *UserTagAccessUserAllow `url:"user-allow,omitempty" json:"user-allow,omitempty"`           // Controls tag usage for users without `Sys.Modify` on `/` by either allowing `none`, a `list`, already `existing` or anything (`free`).
	UserAllowList *string                 `url:"user-allow-list,omitempty" json:"user-allow-list,omitempty"` // List of tags users are allowed to set and delete (semicolon separated) for 'user-allow' values 'list' and 'existing'.
}
type _UserTagAccess UserTagAccess

func (t UserTagAccess) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[user-allow=<enum>] [,user-allow-list=<tag>[;<tag>...]]`)
}

func (t *UserTagAccess) UnmarshalJSON(d []byte) error {
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

			values["user-allow"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["user-allow"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.UserAllow)
		if err != nil {
			return err
		}
	}

	if v, ok := values["user-allow-list"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.UserAllowList)
		if err != nil {
			return err
		}
	}

	return nil
}

// Set bandwidth/io limits various operations.
type Bwlimit struct {

	// The following parameters are optional
	Clone     *float64 `url:"clone,omitempty" json:"clone,omitempty"`         // bandwidth limit in KiB/s for cloning disks
	Default   *float64 `url:"default,omitempty" json:"default,omitempty"`     // default bandwidth limit in KiB/s
	Migration *float64 `url:"migration,omitempty" json:"migration,omitempty"` // bandwidth limit in KiB/s for migrating guests (including moving local disks)
	Move      *float64 `url:"move,omitempty" json:"move,omitempty"`           // bandwidth limit in KiB/s for moving disks
	Restore   *float64 `url:"restore,omitempty" json:"restore,omitempty"`     // bandwidth limit in KiB/s for restoring guests from backups
}
type _Bwlimit Bwlimit

func (t Bwlimit) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[clone=<LIMIT>] [,default=<LIMIT>] [,migration=<LIMIT>] [,move=<LIMIT>] [,restore=<LIMIT>]`)
}

func (t *Bwlimit) UnmarshalJSON(d []byte) error {
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

			values["clone"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["clone"]; ok {

		err := json.Unmarshal([]byte(v), &t.Clone)
		if err != nil {
			return err
		}
	}

	if v, ok := values["default"]; ok {

		err := json.Unmarshal([]byte(v), &t.Default)
		if err != nil {
			return err
		}
	}

	if v, ok := values["migration"]; ok {

		err := json.Unmarshal([]byte(v), &t.Migration)
		if err != nil {
			return err
		}
	}

	if v, ok := values["move"]; ok {

		err := json.Unmarshal([]byte(v), &t.Move)
		if err != nil {
			return err
		}
	}

	if v, ok := values["restore"]; ok {

		err := json.Unmarshal([]byte(v), &t.Restore)
		if err != nil {
			return err
		}
	}

	return nil
}

// For cluster wide migration settings.
type Migration struct {
	Type MigrationType `url:"type" json:"type"` // Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.

	// The following parameters are optional
	Network *string `url:"network,omitempty" json:"network,omitempty"` // CIDR of the (sub) network that is used for migration.
}
type _Migration Migration

func (t Migration) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[type=]<secure|insecure> [,network=<CIDR>]`)
}

func (t *Migration) UnmarshalJSON(d []byte) error {
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

			values["type"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["type"]; ok {

		err := json.Unmarshal([]byte(v), &t.Type)
		if err != nil {
			return err
		}
	}

	if v, ok := values["network"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Network)
		if err != nil {
			return err
		}
	}

	return nil
}

// u2f
type U2f struct {

	// The following parameters are optional
	Appid  *string `url:"appid,omitempty" json:"appid,omitempty"`   // U2F AppId URL override. Defaults to the origin.
	Origin *string `url:"origin,omitempty" json:"origin,omitempty"` // U2F Origin override. Mostly useful for single nodes with a single URL.
}
type _U2f U2f

func (t U2f) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[appid=<APPID>] [,origin=<URL>]`)
}

func (t *U2f) UnmarshalJSON(d []byte) error {
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

			values["appid"] = kv[0]

			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["appid"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Appid)
		if err != nil {
			return err
		}
	}

	if v, ok := values["origin"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Origin)
		if err != nil {
			return err
		}
	}

	return nil
}

type SetOptionsRequest struct {

	// The following parameters are optional
	Bwlimit           *Bwlimit       `url:"bwlimit,omitempty" json:"bwlimit,omitempty"`                       // Set bandwidth/io limits various operations.
	Console           *Console       `url:"console,omitempty" json:"console,omitempty"`                       // Select the default Console viewer. You can either use the builtin java applet (VNC; deprecated and maps to html5), an external virt-viewer comtatible application (SPICE), an HTML5 based vnc viewer (noVNC), or an HTML5 based console client (xtermjs). If the selected viewer is not available (e.g. SPICE not activated for the VM), the fallback is noVNC.
	Crs               *Crs           `url:"crs,omitempty" json:"crs,omitempty"`                               // Cluster resource scheduling settings.
	Delete            *string        `url:"delete,omitempty" json:"delete,omitempty"`                         // A list of settings you want to delete.
	Description       *string        `url:"description,omitempty" json:"description,omitempty"`               // Datacenter description. Shown in the web-interface datacenter notes panel. This is saved as comment inside the configuration file.
	EmailFrom         *string        `url:"email_from,omitempty" json:"email_from,omitempty"`                 // Specify email address to send notification from (default is root@$hostname)
	Fencing           *Fencing       `url:"fencing,omitempty" json:"fencing,omitempty"`                       // Set the fencing mode of the HA cluster. Hardware mode needs a valid configuration of fence devices in /etc/pve/ha/fence.cfg. With both all two modes are used. WARNING: 'hardware' and 'both' are EXPERIMENTAL & WIP
	Ha                *Ha            `url:"ha,omitempty" json:"ha,omitempty"`                                 // Cluster wide HA settings.
	HttpProxy         *string        `url:"http_proxy,omitempty" json:"http_proxy,omitempty"`                 // Specify external http proxy which is used for downloads (example: 'http://username:password@host:port/')
	Keyboard          *Keyboard      `url:"keyboard,omitempty" json:"keyboard,omitempty"`                     // Default keybord layout for vnc server.
	Language          *Language      `url:"language,omitempty" json:"language,omitempty"`                     // Default GUI language.
	MacPrefix         *string        `url:"mac_prefix,omitempty" json:"mac_prefix,omitempty"`                 // Prefix for autogenerated MAC addresses.
	MaxWorkers        *int           `url:"max_workers,omitempty" json:"max_workers,omitempty"`               // Defines how many workers (per node) are maximal started on actions like 'stopall VMs' or task from the ha-manager.
	Migration         *Migration     `url:"migration,omitempty" json:"migration,omitempty"`                   // For cluster wide migration settings.
	MigrationUnsecure *util.PVEBool  `url:"migration_unsecure,omitempty" json:"migration_unsecure,omitempty"` // Migration is secure using SSH tunnel by default. For secure private networks you can disable it to speed up migration. Deprecated, use the 'migration' property instead!
	NextId            *NextId        `url:"next-id,omitempty" json:"next-id,omitempty"`                       // Control the range for the free VMID auto-selection pool.
	RegisteredTags    *string        `url:"registered-tags,omitempty" json:"registered-tags,omitempty"`       // A list of tags that require a `Sys.Modify` on '/' to set and delete. Tags set here that are also in 'user-tag-access' also require `Sys.Modify`.
	TagStyle          *TagStyle      `url:"tag-style,omitempty" json:"tag-style,omitempty"`                   // Tag style options.
	U2f               *U2f           `url:"u2f,omitempty" json:"u2f,omitempty"`                               // u2f
	UserTagAccess     *UserTagAccess `url:"user-tag-access,omitempty" json:"user-tag-access,omitempty"`       // Privilege options for user-settable tags
	Webauthn          *Webauthn      `url:"webauthn,omitempty" json:"webauthn,omitempty"`                     // webauthn configuration
}
type _SetOptionsRequest SetOptionsRequest

type GetStatusResponse struct {
	Id   string `url:"id" json:"id"`
	Name string `url:"name" json:"name"`
	Type Type   `url:"type" json:"type"` // Indicates the type, either cluster or node. The type defines the object properties e.g. quorate available for type cluster.

	// The following parameters are optional
	Ip      *string       `url:"ip,omitempty" json:"ip,omitempty"`           // [node] IP of the resolved nodename.
	Level   *string       `url:"level,omitempty" json:"level,omitempty"`     // [node] Proxmox VE Subscription level, indicates if eligible for enterprise support as well as access to the stable Proxmox VE Enterprise Repository.
	Local   *util.PVEBool `url:"local,omitempty" json:"local,omitempty"`     // [node] Indicates if this is the responding node.
	Nodeid  *int          `url:"nodeid,omitempty" json:"nodeid,omitempty"`   // [node] ID of the node from the corosync configuration.
	Nodes   *int          `url:"nodes,omitempty" json:"nodes,omitempty"`     // [cluster] Nodes count, including offline nodes.
	Online  *util.PVEBool `url:"online,omitempty" json:"online,omitempty"`   // [node] Indicates if the node is online or offline.
	Quorate *util.PVEBool `url:"quorate,omitempty" json:"quorate,omitempty"` // [cluster] Indicates if there is a majority of nodes online to make decisions
	Version *int          `url:"version,omitempty" json:"version,omitempty"` // [cluster] Current version of the corosync configuration file.
}
type _GetStatusResponse GetStatusResponse

type NextidRequest struct {

	// The following parameters are optional
	Vmid *int `url:"vmid,omitempty" json:"vmid,omitempty"` // The (unique) ID of the VM.
}
type _NextidRequest NextidRequest

// Index Cluster index.
func (c *Client) Index(ctx context.Context) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster", "GET", &resp, nil)
	return resp, err
}

// Log Read cluster log
func (c *Client) Log(ctx context.Context, req LogRequest) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/log", "GET", &resp, req)
	return resp, err
}

// Resources Resources index (cluster wide).
func (c *Client) Resources(ctx context.Context, req ResourcesRequest) ([]ResourcesResponse, error) {
	var resp []ResourcesResponse

	err := c.httpClient.Do(ctx, "/cluster/resources", "GET", &resp, req)
	return resp, err
}

// Tasks List recent tasks (cluster wide).
func (c *Client) Tasks(ctx context.Context) ([]TasksResponse, error) {
	var resp []TasksResponse

	err := c.httpClient.Do(ctx, "/cluster/tasks", "GET", &resp, nil)
	return resp, err
}

// GetOptions Get datacenter options. Without 'Sys.Audit' on '/' not all options are returned.
func (c *Client) GetOptions(ctx context.Context) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/options", "GET", &resp, nil)
	return resp, err
}

// SetOptions Set datacenter options.
func (c *Client) SetOptions(ctx context.Context, req SetOptionsRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/options", "PUT", nil, req)
	return err
}

// GetStatus Get cluster status information.
func (c *Client) GetStatus(ctx context.Context) ([]GetStatusResponse, error) {
	var resp []GetStatusResponse

	err := c.httpClient.Do(ctx, "/cluster/status", "GET", &resp, nil)
	return resp, err
}

// Nextid Get next free VMID. Pass a VMID to assert that its free (at time of check).
func (c *Client) Nextid(ctx context.Context, req NextidRequest) (int, error) {
	var resp int

	err := c.httpClient.Do(ctx, "/cluster/nextid", "GET", &resp, req)
	return resp, err
}

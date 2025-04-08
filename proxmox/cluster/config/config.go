// File generated by proxmox json schema, DO NOT EDIT

package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

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

// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
type Link struct {
	Address string `url:"address" json:"address"` // Hostname (or IP) of this corosync link address.

	// The following parameters are optional
	Priority *int `url:"priority,omitempty" json:"priority,omitempty"` // The priority for the link when knet is used in 'passive' mode (default). Lower value means higher priority. Only valid for cluster create, ignored on node add.
}
type _Link Link

func (t Link) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[address=]<IP> [,priority=<integer>]`)
}

func (t *Link) UnmarshalJSON(d []byte) error {
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
			values["address"] = kv[0]
			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["address"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Address)
		if err != nil {
			return err
		}
	}

	if v, ok := values["priority"]; ok {

		err := json.Unmarshal([]byte(v), &t.Priority)
		if err != nil {
			return err
		}
	}

	return nil
}

// Set of Link
type Links map[string]*Link
type _Links Links

func (t Links) EncodeValues(key string, v *url.Values) error {
	return util.EncodeArray(key, v, t)
}

type CreateRequest struct {
	Clustername string `url:"clustername" json:"clustername"` // The name of the cluster.

	// The following parameters are optional
	Links  *Links `url:"link[n],omitempty" json:"link[n],omitempty"` // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid *int   `url:"nodeid,omitempty" json:"nodeid,omitempty"`   // Node id for this node.
	Votes  *int   `url:"votes,omitempty" json:"votes,omitempty"`     // Number of votes for this node.
}
type _CreateRequest CreateRequest

func (t *CreateRequest) UnmarshalJSON(d []byte) error {
	tmp := _CreateRequest{}
	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}
	*t = CreateRequest(tmp)
	rest := map[string]json.RawMessage{}
	err = json.Unmarshal(d, &rest)
	if err != nil {
		return err
	}
	for k, v := range rest {

		if ok, err := regexp.MatchString("^link[0-9]+$", k); ok {
			if t.Links == nil {
				set := make(Links)
				t.Links = &set
			}
			var newVal Link
			err = json.Unmarshal(v, &newVal)
			if err != nil {
				return err
			}
			(*t.Links)[k] = &newVal
		} else if err != nil {
			return err
		}

	}
	return nil
}

type JoinInfoRequest struct {

	// The following parameters are optional
	Node *string `url:"node,omitempty" json:"node,omitempty"` // The node for which the joinee gets the nodeinfo.
}
type _JoinInfoRequest JoinInfoRequest

// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
type Ring0Addr struct {
	Address string `url:"address" json:"address"` // Hostname (or IP) of this corosync link address.

	// The following parameters are optional
	Priority *int `url:"priority,omitempty" json:"priority,omitempty"` // The priority for the link when knet is used in 'passive' mode (default). Lower value means higher priority. Only valid for cluster create, ignored on node add.
}
type _Ring0Addr Ring0Addr

type Nodelist struct {
	Name        string `url:"name" json:"name"` // The cluster node name.
	PveAddr     string `url:"pve_addr" json:"pve_addr"`
	PveFp       string `url:"pve_fp" json:"pve_fp"` // Certificate SHA 256 fingerprint.
	QuorumVotes int    `url:"quorum_votes" json:"quorum_votes"`

	// The following parameters are optional
	Nodeid    *int       `url:"nodeid,omitempty" json:"nodeid,omitempty"`         // Node id for this node.
	Ring0Addr *Ring0Addr `url:"ring0_addr,omitempty" json:"ring0_addr,omitempty"` // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
}
type _Nodelist Nodelist

type JoinInfoResponse struct {
	ConfigDigest  string                 `url:"config_digest" json:"config_digest"`
	Nodelist      []Nodelist             `url:"nodelist" json:"nodelist"`
	PreferredNode string                 `url:"preferred_node" json:"preferred_node"` // The cluster node name.
	Totem         map[string]interface{} `url:"totem" json:"totem"`
}
type _JoinInfoResponse JoinInfoResponse

type JoinRequest struct {
	Fingerprint string `url:"fingerprint" json:"fingerprint"` // Certificate SHA 256 fingerprint.
	Hostname    string `url:"hostname" json:"hostname"`       // Hostname (or IP) of an existing cluster member.
	Password    string `url:"password" json:"password"`       // Superuser (root) password of peer node.

	// The following parameters are optional
	Force  *util.PVEBool `url:"force,omitempty" json:"force,omitempty"`     // Do not throw error if node already exists.
	Links  *Links        `url:"link[n],omitempty" json:"link[n],omitempty"` // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid *int          `url:"nodeid,omitempty" json:"nodeid,omitempty"`   // Node id for this node.
	Votes  *int          `url:"votes,omitempty" json:"votes,omitempty"`     // Number of votes for this node
}
type _JoinRequest JoinRequest

func (t *JoinRequest) UnmarshalJSON(d []byte) error {
	tmp := _JoinRequest{}
	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}
	*t = JoinRequest(tmp)
	rest := map[string]json.RawMessage{}
	err = json.Unmarshal(d, &rest)
	if err != nil {
		return err
	}
	for k, v := range rest {

		if ok, err := regexp.MatchString("^link[0-9]+$", k); ok {
			if t.Links == nil {
				set := make(Links)
				t.Links = &set
			}
			var newVal Link
			err = json.Unmarshal(v, &newVal)
			if err != nil {
				return err
			}
			(*t.Links)[k] = &newVal
		} else if err != nil {
			return err
		}

	}
	return nil
}

// Index Directory index.
func (c *Client) Index(ctx context.Context) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/config", "GET", &resp, nil)
	return resp, err
}

// Create Generate new cluster configuration. If no links given, default to local IP address as link0.
func (c *Client) Create(ctx context.Context, req CreateRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/cluster/config", "POST", &resp, req)
	return resp, err
}

// JoinApiVersionApiversion Return the version of the cluster join API available on this node.
func (c *Client) JoinApiVersionApiversion(ctx context.Context) (int, error) {
	var resp int

	err := c.httpClient.Do(ctx, "/cluster/config/apiversion", "GET", &resp, nil)
	return resp, err
}

// JoinInfo Get information needed to join this cluster over the connected node.
func (c *Client) JoinInfo(ctx context.Context, req JoinInfoRequest) (JoinInfoResponse, error) {
	var resp JoinInfoResponse

	err := c.httpClient.Do(ctx, "/cluster/config/join", "GET", &resp, req)
	return resp, err
}

// Join Joins this node into an existing cluster. If no links are given, default to IP resolved by node's hostname on single link (fallback fails for clusters with multiple links).
func (c *Client) Join(ctx context.Context, req JoinRequest) (string, error) {
	var resp string

	err := c.httpClient.Do(ctx, "/cluster/config/join", "POST", &resp, req)
	return resp, err
}

// Totem Get corosync totem protocol settings.
func (c *Client) Totem(ctx context.Context) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/config/totem", "GET", &resp, nil)
	return resp, err
}

// StatusQdevice Get QDevice status
func (c *Client) StatusQdevice(ctx context.Context) (map[string]interface{}, error) {
	var resp map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/config/qdevice", "GET", &resp, nil)
	return resp, err
}

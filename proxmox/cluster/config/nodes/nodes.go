// File generated by proxmox json schema, DO NOT EDIT

package nodes

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

type IndexResponse struct {
	Node string `url:"node" json:"node"`
}
type _IndexResponse IndexResponse

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

type ChildCreateRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

	// The following parameters are optional
	Apiversion *int          `url:"apiversion,omitempty" json:"apiversion,omitempty"`   // The JOIN_API_VERSION of the new node.
	Force      *util.PVEBool `url:"force,omitempty" json:"force,omitempty"`             // Do not throw error if node already exists.
	Links      *Links        `url:"link[n],omitempty" json:"link[n],omitempty"`         // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	NewNodeIp  *string       `url:"new_node_ip,omitempty" json:"new_node_ip,omitempty"` // IP Address of node to add. Used as fallback if no links are given.
	Nodeid     *int          `url:"nodeid,omitempty" json:"nodeid,omitempty"`           // Node id for this node.
	Votes      *int          `url:"votes,omitempty" json:"votes,omitempty"`             // Number of votes for this node
}
type _ChildCreateRequest ChildCreateRequest

func (t *ChildCreateRequest) UnmarshalJSON(d []byte) error {
	tmp := _ChildCreateRequest{}
	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}
	*t = ChildCreateRequest(tmp)
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

type ChildCreateResponse struct {
	CorosyncAuthkey string   `url:"corosync_authkey" json:"corosync_authkey"`
	CorosyncConf    string   `url:"corosync_conf" json:"corosync_conf"`
	Warnings        []string `url:"warnings" json:"warnings"`
}
type _ChildCreateResponse ChildCreateResponse

type DeleteRequest struct {
	Node string `url:"node" json:"node"` // The cluster node name.

}
type _DeleteRequest DeleteRequest

// Index Corosync node list.
func (c *Client) Index(ctx context.Context) ([]IndexResponse, error) {
	var resp []IndexResponse

	err := c.httpClient.Do(ctx, "/cluster/config/nodes", "GET", &resp, nil)
	return resp, err
}

// ChildCreate Adds a node to the cluster configuration. This call is for internal use.
func (c *Client) ChildCreate(ctx context.Context, req ChildCreateRequest) (ChildCreateResponse, error) {
	var resp ChildCreateResponse

	err := c.httpClient.Do(ctx, "/cluster/config/nodes/{node}", "POST", &resp, req)
	return resp, err
}

// Delete Removes a node from the cluster configuration.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/config/nodes/{node}", "DELETE", nil, req)
	return err
}

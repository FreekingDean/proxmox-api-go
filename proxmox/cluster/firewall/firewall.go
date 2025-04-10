// File generated by proxmox json schema, DO NOT EDIT

package firewall

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FreekingDean/proxmox-api-go/internal/util"
	"net/url"
	"strings"
)

const (
	PolicyIn_ACCEPT PolicyIn = "ACCEPT"
	PolicyIn_REJECT PolicyIn = "REJECT"
	PolicyIn_DROP   PolicyIn = "DROP"

	PolicyOut_ACCEPT PolicyOut = "ACCEPT"
	PolicyOut_REJECT PolicyOut = "REJECT"
	PolicyOut_DROP   PolicyOut = "DROP"

	Type_ALIAS Type = "alias"
	Type_IPSET Type = "ipset"
)

type PolicyIn string
type PolicyOut string
type Type string

func PtrPolicyIn(i PolicyIn) *PolicyIn {
	return &i
}
func PtrPolicyOut(i PolicyOut) *PolicyOut {
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

// Log ratelimiting settings
type LogRatelimit struct {
	Enable util.PVEBool `url:"enable" json:"enable"` // Enable or disable log rate limiting

	// The following parameters are optional
	Burst *int    `url:"burst,omitempty" json:"burst,omitempty"` // Initial burst of packages which will always get logged before the rate is applied
	Rate  *string `url:"rate,omitempty" json:"rate,omitempty"`   // Frequency with which the burst bucket gets refilled
}
type _LogRatelimit LogRatelimit

func (t LogRatelimit) EncodeValues(key string, v *url.Values) error {
	return util.EncodeString(key, v, t, `[enable=]<1|0> [,burst=<integer>] [,rate=<rate>]`)
}

func (t *LogRatelimit) UnmarshalJSON(d []byte) error {
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
			values["enable"] = kv[0]
			continue
		}
		values[kv[0]] = kv[1]
	}

	if v, ok := values["enable"]; ok {

		err := json.Unmarshal([]byte(v), &t.Enable)
		if err != nil {
			return err
		}
	}

	if v, ok := values["burst"]; ok {

		err := json.Unmarshal([]byte(v), &t.Burst)
		if err != nil {
			return err
		}
	}

	if v, ok := values["rate"]; ok {

		v = fmt.Sprintf("\"%s\"", v)

		err := json.Unmarshal([]byte(v), &t.Rate)
		if err != nil {
			return err
		}
	}

	return nil
}

type GetOptionsResponse struct {

	// The following parameters are optional
	Ebtables     *util.PVEBool `url:"ebtables,omitempty" json:"ebtables,omitempty"`           // Enable ebtables rules cluster wide.
	Enable       *int          `url:"enable,omitempty" json:"enable,omitempty"`               // Enable or disable the firewall cluster wide.
	LogRatelimit *LogRatelimit `url:"log_ratelimit,omitempty" json:"log_ratelimit,omitempty"` // Log ratelimiting settings
	PolicyIn     *PolicyIn     `url:"policy_in,omitempty" json:"policy_in,omitempty"`         // Input policy.
	PolicyOut    *PolicyOut    `url:"policy_out,omitempty" json:"policy_out,omitempty"`       // Output policy.
}
type _GetOptionsResponse GetOptionsResponse

type SetOptionsRequest struct {

	// The following parameters are optional
	Delete       *string       `url:"delete,omitempty" json:"delete,omitempty"`               // A list of settings you want to delete.
	Digest       *string       `url:"digest,omitempty" json:"digest,omitempty"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Ebtables     *util.PVEBool `url:"ebtables,omitempty" json:"ebtables,omitempty"`           // Enable ebtables rules cluster wide.
	Enable       *int          `url:"enable,omitempty" json:"enable,omitempty"`               // Enable or disable the firewall cluster wide.
	LogRatelimit *LogRatelimit `url:"log_ratelimit,omitempty" json:"log_ratelimit,omitempty"` // Log ratelimiting settings
	PolicyIn     *PolicyIn     `url:"policy_in,omitempty" json:"policy_in,omitempty"`         // Input policy.
	PolicyOut    *PolicyOut    `url:"policy_out,omitempty" json:"policy_out,omitempty"`       // Output policy.
}
type _SetOptionsRequest SetOptionsRequest

type GetMacrosResponse struct {
	Descr string `url:"descr" json:"descr"` // More verbose description (if available).
	Macro string `url:"macro" json:"macro"` // Macro name.

}
type _GetMacrosResponse GetMacrosResponse

type RefsRequest struct {

	// The following parameters are optional
	Type *Type `url:"type,omitempty" json:"type,omitempty"` // Only list references of specified type.
}
type _RefsRequest RefsRequest

type RefsResponse struct {
	Name string `url:"name" json:"name"`
	Ref  string `url:"ref" json:"ref"`
	Type Type   `url:"type" json:"type"`

	// The following parameters are optional
	Comment *string `url:"comment,omitempty" json:"comment,omitempty"`
}
type _RefsResponse RefsResponse

// Index Directory index.
func (c *Client) Index(ctx context.Context) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}

	err := c.httpClient.Do(ctx, "/cluster/firewall", "GET", &resp, nil)
	return resp, err
}

// GetOptions Get Firewall options.
func (c *Client) GetOptions(ctx context.Context) (GetOptionsResponse, error) {
	var resp GetOptionsResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/options", "GET", &resp, nil)
	return resp, err
}

// SetOptions Set Firewall options.
func (c *Client) SetOptions(ctx context.Context, req SetOptionsRequest) error {

	err := c.httpClient.Do(ctx, "/cluster/firewall/options", "PUT", nil, req)
	return err
}

// GetMacros List available macros
func (c *Client) GetMacros(ctx context.Context) ([]GetMacrosResponse, error) {
	var resp []GetMacrosResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/macros", "GET", &resp, nil)
	return resp, err
}

// Refs Lists possible IPSet/Alias reference which are allowed in source/dest properties.
func (c *Client) Refs(ctx context.Context, req RefsRequest) ([]RefsResponse, error) {
	var resp []RefsResponse

	err := c.httpClient.Do(ctx, "/cluster/firewall/refs", "GET", &resp, req)
	return resp, err
}

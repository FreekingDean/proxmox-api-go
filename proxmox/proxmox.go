package proxmox

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/FreekingDean/proxmox-api-go/internal/util"
	"github.com/google/go-querystring/query"
)

var (
	routeRegex = regexp.MustCompile(`\{(.*?)\}`)
)

type Client struct {
	client   *http.Client
	baseAddr string
	cookie   string
	csrf     string
}

type ClientOption func(*Client)

func NewClient(baseAddr string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &Client{
		client:   &http.Client{},
		baseAddr: baseAddr,
	}
}

type Response struct {
	Data interface{} `json:"data"`
}

func (c *Client) SetCookie(cookie string) {
	c.cookie = cookie
}

func (c *Client) SetCsrf(csrf string) {
	c.csrf = csrf
}

func (c *Client) Do(ctx context.Context, route string, method string, response interface{}, request interface{}) error {
	v, err := query.Values(request)
	if err != nil {
		return err
	}
	params := []interface{}{}
	paramRoute := routeRegex.ReplaceAllStringFunc(route, func(s string) string {
		key := s[1 : len(s)-1]
		params = append(params, v.Get(key))
		v.Del(key)
		return "%s"
	})
	route = fmt.Sprintf(paramRoute, params...)
	query, buf := queryBuf(method, v)
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s%s", c.baseAddr, route, query), buf)
	if err != nil {
		return err
	}
	if c.cookie != "" {
		req.Header.Add("Authorization", fmt.Sprintf("PVEAuthCookie=%s", c.cookie))
	}
	if c.csrf != "" && method != http.MethodGet {
		req.Header.Add("CSRFPreventionToken", c.csrf)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 400 {
		return fmt.Errorf("parameter error: %s", v.Encode())
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non 200: %s", resp.Status)
	}
	dec := json.NewDecoder(resp.Body)
	respObj := Response{
		Data: response,
	}
	return dec.Decode(&respObj)
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

func Float64(f float64) *float64 {
	return &f
}

func PVEBool(b bool) *util.PVEBool {
	return (*util.PVEBool)(&b)
}

func queryBuf(method string, v url.Values) (string, io.Reader) {
	if method == http.MethodGet || method == http.MethodDelete {
		return fmt.Sprintf("?%s", v.Encode()), nil
	} else {
		return "", bytes.NewBufferString(v.Encode())
	}
}

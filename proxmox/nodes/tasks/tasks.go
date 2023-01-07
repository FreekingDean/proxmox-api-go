// File generated by proxmox json schema, DO NOT EDIT

package tasks

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

type IndexRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Errors       *bool   `url:"errors,omitempty",json:"errors,omitempty"`             // Only list tasks with a status of ERROR.
	Limit        *int    `url:"limit,omitempty",json:"limit,omitempty"`               // Only list this amount of tasks.
	Since        *int    `url:"since,omitempty",json:"since,omitempty"`               // Only list tasks since this UNIX epoch.
	Source       *string `url:"source,omitempty",json:"source,omitempty"`             // List archived, active or all tasks.
	Start        *int    `url:"start,omitempty",json:"start,omitempty"`               // List tasks beginning from this offset.
	Statusfilter *string `url:"statusfilter,omitempty",json:"statusfilter,omitempty"` // List of Task States that should be returned.
	Typefilter   *string `url:"typefilter,omitempty",json:"typefilter,omitempty"`     // Only list tasks of this type (e.g., vzstart, vzdump).
	Until        *int    `url:"until,omitempty",json:"until,omitempty"`               // Only list tasks until this UNIX epoch.
	Userfilter   *string `url:"userfilter,omitempty",json:"userfilter,omitempty"`     // Only list tasks from this user.
	Vmid         *int    `url:"vmid,omitempty",json:"vmid,omitempty"`                 // Only list tasks for this VM.
}

type IndexResponse []*struct {
	Id        string `url:"id",json:"id"`
	Node      string `url:"node",json:"node"`
	Pid       int    `url:"pid",json:"pid"`
	Pstart    int    `url:"pstart",json:"pstart"`
	Starttime int    `url:"starttime",json:"starttime"`
	Type      string `url:"type",json:"type"`
	Upid      string `url:"upid",json:"upid"`
	User      string `url:"user",json:"user"`

	// The following parameters are optional
	Endtime *int    `url:"endtime,omitempty",json:"endtime,omitempty"`
	Status  *string `url:"status,omitempty",json:"status,omitempty"`
}

// Index Read task list for one node (finished tasks).
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/tasks", "GET", &resp, req)
	return resp, err
}

type FindRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Upid string `url:"upid",json:"upid"`
}

type FindResponse []*map[string]interface{}

// Find
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/tasks/{upid}", "GET", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Upid string `url:"upid",json:"upid"`
}

type DeleteResponse map[string]interface{}

// Delete Stop a task.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/tasks/{upid}", "DELETE", &resp, req)
	return resp, err
}

type ReadTaskLogRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Upid string `url:"upid",json:"upid"` // The task's unique ID.

	// The following parameters are optional
	Limit *int `url:"limit,omitempty",json:"limit,omitempty"` // The maximum amount of lines that should be printed.
	Start *int `url:"start,omitempty",json:"start,omitempty"` // The line number to start printing at.
}

type ReadTaskLogResponse []*struct {
	N int    `url:"n",json:"n"` // Line number
	T string `url:"t",json:"t"` // Line text

}

// ReadTaskLog Read task log.
func (c *Client) ReadTaskLog(ctx context.Context, req *ReadTaskLogRequest) (*ReadTaskLogResponse, error) {
	var resp *ReadTaskLogResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/tasks/{upid}/log", "GET", &resp, req)
	return resp, err
}

type ReadTaskStatusRequest struct {
	Node string `url:"node",json:"node"` // The cluster node name.
	Upid string `url:"upid",json:"upid"` // The task's unique ID.

}

type ReadTaskStatusResponse struct {
	Id        string  `url:"id",json:"id"`
	Node      string  `url:"node",json:"node"`
	Pid       int     `url:"pid",json:"pid"`
	Starttime float64 `url:"starttime",json:"starttime"`
	Status    string  `url:"status",json:"status"`
	Type      string  `url:"type",json:"type"`
	Upid      string  `url:"upid",json:"upid"`
	User      string  `url:"user",json:"user"`

	// The following parameters are optional
	Exitstatus *string `url:"exitstatus,omitempty",json:"exitstatus,omitempty"`
}

// ReadTaskStatus Read task status.
func (c *Client) ReadTaskStatus(ctx context.Context, req *ReadTaskStatusRequest) (*ReadTaskStatusResponse, error) {
	var resp *ReadTaskStatusResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/tasks/{upid}/status", "GET", &resp, req)
	return resp, err
}

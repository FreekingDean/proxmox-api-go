// File generated by proxmox json schema, DO NOT EDIT

package pools

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

}

type IndexResponse []*struct {
	BytesUsed     int     `url:"bytes_used",json:"bytes_used"`
	CrushRule     int     `url:"crush_rule",json:"crush_rule"`
	CrushRuleName string  `url:"crush_rule_name",json:"crush_rule_name"`
	MinSize       int     `url:"min_size",json:"min_size"`
	PercentUsed   float64 `url:"percent_used",json:"percent_used"`
	PgNum         int     `url:"pg_num",json:"pg_num"`
	Pool          int     `url:"pool",json:"pool"`
	PoolName      string  `url:"pool_name",json:"pool_name"`
	Size          int     `url:"size",json:"size"`
	Type          string  `url:"type",json:"type"`

	// The following parameters are optional
	ApplicationMetadata map[string]interface{} `url:"application_metadata,omitempty",json:"application_metadata,omitempty"`
	AutoscaleStatus     map[string]interface{} `url:"autoscale_status,omitempty",json:"autoscale_status,omitempty"`
	PgAutoscaleMode     *string                `url:"pg_autoscale_mode,omitempty",json:"pg_autoscale_mode,omitempty"`
	PgNumFinal          *int                   `url:"pg_num_final,omitempty",json:"pg_num_final,omitempty"`
	PgNumMin            *int                   `url:"pg_num_min,omitempty",json:"pg_num_min,omitempty"`
	TargetSize          *int                   `url:"target_size,omitempty",json:"target_size,omitempty"`
	TargetSizeRatio     *float64               `url:"target_size_ratio,omitempty",json:"target_size_ratio,omitempty"`
}

// Index List all pools.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/pools", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Name string `url:"name",json:"name"` // The name of the pool. It must be unique.
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	AddStorages     *util.SpecialBool `url:"add_storages,omitempty",json:"add_storages,omitempty"`           // Configure VM and CT storage using the new pool.
	Application     *string           `url:"application,omitempty",json:"application,omitempty"`             // The application of the pool.
	CrushRule       *string           `url:"crush_rule,omitempty",json:"crush_rule,omitempty"`               // The rule to use for mapping object placement in the cluster.
	ErasureCoding   *string           `url:"erasure-coding,omitempty",json:"erasure-coding,omitempty"`       // Create an erasure coded pool for RBD with an accompaning replicated pool for metadata storage. With EC, the common ceph options 'size', 'min_size' and 'crush_rule' parameters will be applied to the metadata pool.
	MinSize         *int              `url:"min_size,omitempty",json:"min_size,omitempty"`                   // Minimum number of replicas per object
	PgAutoscaleMode *string           `url:"pg_autoscale_mode,omitempty",json:"pg_autoscale_mode,omitempty"` // The automatic PG scaling mode of the pool.
	PgNum           *int              `url:"pg_num,omitempty",json:"pg_num,omitempty"`                       // Number of placement groups.
	PgNumMin        *int              `url:"pg_num_min,omitempty",json:"pg_num_min,omitempty"`               // Minimal number of placement groups.
	Size            *int              `url:"size,omitempty",json:"size,omitempty"`                           // Number of replicas per object
	TargetSize      *string           `url:"target_size,omitempty",json:"target_size,omitempty"`             // The estimated target size of the pool for the PG autoscaler.
	TargetSizeRatio *float64          `url:"target_size_ratio,omitempty",json:"target_size_ratio,omitempty"` // The estimated target ratio of the pool for the PG autoscaler.
}

type CreateResponse string

// Create Create Ceph pool
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/pools", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Name string `url:"name",json:"name"` // The name of the pool. It must be unique.
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Verbose *util.SpecialBool `url:"verbose,omitempty",json:"verbose,omitempty"` // If enabled, will display additional data(eg. statistics).
}

type FindResponse struct {
	FastRead             util.SpecialBool `url:"fast_read",json:"fast_read"`
	Hashpspool           util.SpecialBool `url:"hashpspool",json:"hashpspool"`
	Id                   int              `url:"id",json:"id"`
	Name                 string           `url:"name",json:"name"` // The name of the pool. It must be unique.
	NodeepScrub          util.SpecialBool `url:"nodeep-scrub",json:"nodeep-scrub"`
	Nodelete             util.SpecialBool `url:"nodelete",json:"nodelete"`
	Nopgchange           util.SpecialBool `url:"nopgchange",json:"nopgchange"`
	Noscrub              util.SpecialBool `url:"noscrub",json:"noscrub"`
	Nosizechange         util.SpecialBool `url:"nosizechange",json:"nosizechange"`
	PgpNum               int              `url:"pgp_num",json:"pgp_num"`
	UseGmtHitset         util.SpecialBool `url:"use_gmt_hitset",json:"use_gmt_hitset"`
	WriteFadviseDontneed util.SpecialBool `url:"write_fadvise_dontneed",json:"write_fadvise_dontneed"`

	// The following parameters are optional
	Application     *string                   `url:"application,omitempty",json:"application,omitempty"` // The application of the pool.
	ApplicationList []*map[string]interface{} `url:"application_list,omitempty",json:"application_list,omitempty"`
	AutoscaleStatus map[string]interface{}    `url:"autoscale_status,omitempty",json:"autoscale_status,omitempty"`
	CrushRule       *string                   `url:"crush_rule,omitempty",json:"crush_rule,omitempty"`               // The rule to use for mapping object placement in the cluster.
	MinSize         *int                      `url:"min_size,omitempty",json:"min_size,omitempty"`                   // Minimum number of replicas per object
	PgAutoscaleMode *string                   `url:"pg_autoscale_mode,omitempty",json:"pg_autoscale_mode,omitempty"` // The automatic PG scaling mode of the pool.
	PgNum           *int                      `url:"pg_num,omitempty",json:"pg_num,omitempty"`                       // Number of placement groups.
	PgNumMin        *int                      `url:"pg_num_min,omitempty",json:"pg_num_min,omitempty"`               // Minimal number of placement groups.
	Size            *int                      `url:"size,omitempty",json:"size,omitempty"`                           // Number of replicas per object
	Statistics      map[string]interface{}    `url:"statistics,omitempty",json:"statistics,omitempty"`
	TargetSize      *string                   `url:"target_size,omitempty",json:"target_size,omitempty"`             // The estimated target size of the pool for the PG autoscaler.
	TargetSizeRatio *float64                  `url:"target_size_ratio,omitempty",json:"target_size_ratio,omitempty"` // The estimated target ratio of the pool for the PG autoscaler.
}

// Find List pool settings.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/pools/{name}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Name string `url:"name",json:"name"` // The name of the pool. It must be unique.
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Application     *string  `url:"application,omitempty",json:"application,omitempty"`             // The application of the pool.
	CrushRule       *string  `url:"crush_rule,omitempty",json:"crush_rule,omitempty"`               // The rule to use for mapping object placement in the cluster.
	MinSize         *int     `url:"min_size,omitempty",json:"min_size,omitempty"`                   // Minimum number of replicas per object
	PgAutoscaleMode *string  `url:"pg_autoscale_mode,omitempty",json:"pg_autoscale_mode,omitempty"` // The automatic PG scaling mode of the pool.
	PgNum           *int     `url:"pg_num,omitempty",json:"pg_num,omitempty"`                       // Number of placement groups.
	PgNumMin        *int     `url:"pg_num_min,omitempty",json:"pg_num_min,omitempty"`               // Minimal number of placement groups.
	Size            *int     `url:"size,omitempty",json:"size,omitempty"`                           // Number of replicas per object
	TargetSize      *string  `url:"target_size,omitempty",json:"target_size,omitempty"`             // The estimated target size of the pool for the PG autoscaler.
	TargetSizeRatio *float64 `url:"target_size_ratio,omitempty",json:"target_size_ratio,omitempty"` // The estimated target ratio of the pool for the PG autoscaler.
}

type UpdateResponse string

// Update Change POOL settings
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/pools/{name}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Name string `url:"name",json:"name"` // The name of the pool. It must be unique.
	Node string `url:"node",json:"node"` // The cluster node name.

	// The following parameters are optional
	Force           *util.SpecialBool `url:"force,omitempty",json:"force,omitempty"`                       // If true, destroys pool even if in use
	RemoveEcprofile *util.SpecialBool `url:"remove_ecprofile,omitempty",json:"remove_ecprofile,omitempty"` // Remove the erasure code profile. Defaults to true, if applicable.
	RemoveStorages  *util.SpecialBool `url:"remove_storages,omitempty",json:"remove_storages,omitempty"`   // Remove all pveceph-managed storages configured for this pool
}

type DeleteResponse string

// Delete Destroy pool
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/ceph/pools/{name}", "DELETE", &resp, req)
	return resp, err
}

// File generated by proxmox json schema, DO NOT EDIT

package content

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
	Node    string `url:"node" json:"node"`       // The cluster node name.
	Storage string `url:"storage" json:"storage"` // The storage identifier.

	// The following parameters are optional
	Content *string `url:"content,omitempty" json:"content,omitempty"` // Only list content of this type.
	Vmid    *int    `url:"vmid,omitempty" json:"vmid,omitempty"`       // Only list images for this VM
}

type IndexResponse []*struct {
	Format string `url:"format" json:"format"` // Format identifier ('raw', 'qcow2', 'subvol', 'iso', 'tgz' ...)
	Size   int    `url:"size" json:"size"`     // Volume size in bytes.
	Volid  string `url:"volid" json:"volid"`   // Volume identifier.

	// The following parameters are optional
	Ctime        *int              `url:"ctime,omitempty" json:"ctime,omitempty"`         // Creation time (seconds since the UNIX Epoch).
	Encrypted    *string           `url:"encrypted,omitempty" json:"encrypted,omitempty"` // If whole backup is encrypted, value is the fingerprint or '1'  if encrypted. Only useful for the Proxmox Backup Server storage type.
	Notes        *string           `url:"notes,omitempty" json:"notes,omitempty"`         // Optional notes. If they contain multiple lines, only the first one is returned here.
	Parent       *string           `url:"parent,omitempty" json:"parent,omitempty"`       // Volume identifier of parent (for linked cloned).
	Protected    *util.SpecialBool `url:"protected,omitempty" json:"protected,omitempty"` // Protection status. Currently only supported for backups.
	Used         *int              `url:"used,omitempty" json:"used,omitempty"`           // Used space. Please note that most storage plugins do not report anything useful here.
	Verification struct {
		State string `url:"state" json:"state"` // Last backup verification state.
		Upid  string `url:"upid" json:"upid"`   // Last backup verification UPID.

	} `url:"verification,omitempty" json:"verification,omitempty"` // Last backup verification result, only useful for PBS storages.
	Vmid *int `url:"vmid,omitempty" json:"vmid,omitempty"` // Associated Owner VMID.
}

// Index List storage content.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/content", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Filename string `url:"filename" json:"filename"` // The name of the file to create.
	Node     string `url:"node" json:"node"`         // The cluster node name.
	Size     string `url:"size" json:"size"`         // Size in kilobyte (1024 bytes). Optional suffixes 'M' (megabyte, 1024K) and 'G' (gigabyte, 1024M)
	Storage  string `url:"storage" json:"storage"`   // The storage identifier.
	Vmid     int    `url:"vmid" json:"vmid"`         // Specify owner VM

	// The following parameters are optional
	Format *string `url:"format,omitempty" json:"format,omitempty"`
}

type CreateResponse string

// Create Allocate disk images.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/content", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Volume string `url:"volume" json:"volume"` // Volume identifier

	// The following parameters are optional
	Storage *string `url:"storage,omitempty" json:"storage,omitempty"` // The storage identifier.
}

type FindResponse struct {
	Format string `url:"format" json:"format"` // Format identifier ('raw', 'qcow2', 'subvol', 'iso', 'tgz' ...)
	Path   string `url:"path" json:"path"`     // The Path
	Size   int    `url:"size" json:"size"`     // Volume size in bytes.
	Used   int    `url:"used" json:"used"`     // Used space. Please note that most storage plugins do not report anything useful here.

	// The following parameters are optional
	Notes     *string           `url:"notes,omitempty" json:"notes,omitempty"`         // Optional notes.
	Protected *util.SpecialBool `url:"protected,omitempty" json:"protected,omitempty"` // Protection status. Currently only supported for backups.
}

// Find Get volume attributes
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/content/{volume}", "GET", &resp, req)
	return resp, err
}

type ChildCreateRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Target string `url:"target" json:"target"` // Target volume identifier
	Volume string `url:"volume" json:"volume"` // Source volume identifier

	// The following parameters are optional
	Storage    *string `url:"storage,omitempty" json:"storage,omitempty"`         // The storage identifier.
	TargetNode *string `url:"target_node,omitempty" json:"target_node,omitempty"` // Target node. Default is local node.
}

type ChildCreateResponse string

// ChildCreate Copy a volume. This is experimental code - do not use.
func (c *Client) ChildCreate(ctx context.Context, req *ChildCreateRequest) (*ChildCreateResponse, error) {
	var resp *ChildCreateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/content/{volume}", "POST", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Volume string `url:"volume" json:"volume"` // Volume identifier

	// The following parameters are optional
	Notes     *string           `url:"notes,omitempty" json:"notes,omitempty"`         // The new notes.
	Protected *util.SpecialBool `url:"protected,omitempty" json:"protected,omitempty"` // Protection status. Currently only supported for backups.
	Storage   *string           `url:"storage,omitempty" json:"storage,omitempty"`     // The storage identifier.
}

type UpdateResponse map[string]interface{}

// Update Update volume attributes
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/content/{volume}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Node   string `url:"node" json:"node"`     // The cluster node name.
	Volume string `url:"volume" json:"volume"` // Volume identifier

	// The following parameters are optional
	Delay   *int    `url:"delay,omitempty" json:"delay,omitempty"`     // Time to wait for the task to finish. We return 'null' if the task finish within that time.
	Storage *string `url:"storage,omitempty" json:"storage,omitempty"` // The storage identifier.
}

type DeleteResponse *string

// Delete Delete volume
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/nodes/{node}/storage/{storage}/content/{volume}", "DELETE", &resp, req)
	return resp, err
}

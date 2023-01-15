// File generated by proxmox json schema, DO NOT EDIT

package storage

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

	// The following parameters are optional
	Type *string `url:"type,omitempty",json:"type,omitempty"` // Only list storage of specific type
}

type IndexResponse []*struct {
	Storage string `url:"storage",json:"storage"`
}

// Index Storage index.
func (c *Client) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	var resp *IndexResponse

	err := c.httpClient.Do(ctx, "/storage", "GET", &resp, req)
	return resp, err
}

type CreateRequest struct {
	Storage string `url:"storage",json:"storage"` // The storage identifier.
	Type    string `url:"type",json:"type"`       // Storage type.

	// The following parameters are optional
	Authsupported        *string `url:"authsupported,omitempty",json:"authsupported,omitempty"`                 // Authsupported.
	Base                 *string `url:"base,omitempty",json:"base,omitempty"`                                   // Base volume. This volume is automatically activated.
	Blocksize            *string `url:"blocksize,omitempty",json:"blocksize,omitempty"`                         // block size
	Bwlimit              *string `url:"bwlimit,omitempty",json:"bwlimit,omitempty"`                             // Set bandwidth/io limits various operations.
	ComstarHg            *string `url:"comstar_hg,omitempty",json:"comstar_hg,omitempty"`                       // host group for comstar views
	ComstarTg            *string `url:"comstar_tg,omitempty",json:"comstar_tg,omitempty"`                       // target group for comstar views
	Content              *string `url:"content,omitempty",json:"content,omitempty"`                             // Allowed content types.NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	DataPool             *string `url:"data-pool,omitempty",json:"data-pool,omitempty"`                         // Data Pool (for erasure coding only)
	Datastore            *string `url:"datastore,omitempty",json:"datastore,omitempty"`                         // Proxmox Backup Server datastore name.
	Disable              *bool   `url:"disable,omitempty",json:"disable,omitempty"`                             // Flag to disable the storage.
	Domain               *string `url:"domain,omitempty",json:"domain,omitempty"`                               // CIFS domain.
	EncryptionKey        *string `url:"encryption-key,omitempty",json:"encryption-key,omitempty"`               // Encryption key. Use 'autogen' to generate one automatically without passphrase.
	Export               *string `url:"export,omitempty",json:"export,omitempty"`                               // NFS export path.
	Fingerprint          *string `url:"fingerprint,omitempty",json:"fingerprint,omitempty"`                     // Certificate SHA 256 fingerprint.
	Format               *string `url:"format,omitempty",json:"format,omitempty"`                               // Default image format.
	FsName               *string `url:"fs-name,omitempty",json:"fs-name,omitempty"`                             // The Ceph filesystem name.
	Fuse                 *bool   `url:"fuse,omitempty",json:"fuse,omitempty"`                                   // Mount CephFS through FUSE.
	IsMountpoint         *string `url:"is_mountpoint,omitempty",json:"is_mountpoint,omitempty"`                 // Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	Iscsiprovider        *string `url:"iscsiprovider,omitempty",json:"iscsiprovider,omitempty"`                 // iscsi provider
	Keyring              *string `url:"keyring,omitempty",json:"keyring,omitempty"`                             // Client keyring contents (for external clusters).
	Krbd                 *bool   `url:"krbd,omitempty",json:"krbd,omitempty"`                                   // Always access rbd through krbd kernel module.
	LioTpg               *string `url:"lio_tpg,omitempty",json:"lio_tpg,omitempty"`                             // target portal group for Linux LIO targets
	MasterPubkey         *string `url:"master-pubkey,omitempty",json:"master-pubkey,omitempty"`                 // Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	MaxProtectedBackups  *int    `url:"max-protected-backups,omitempty",json:"max-protected-backups,omitempty"` // Maximal number of protected backups per guest. Use '-1' for unlimited.
	Maxfiles             *int    `url:"maxfiles,omitempty",json:"maxfiles,omitempty"`                           // Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Mkdir                *bool   `url:"mkdir,omitempty",json:"mkdir,omitempty"`                                 // Create the directory if it doesn't exist.
	Monhost              *string `url:"monhost,omitempty",json:"monhost,omitempty"`                             // IP addresses of monitors (for external clusters).
	Mountpoint           *string `url:"mountpoint,omitempty",json:"mountpoint,omitempty"`                       // mount point
	Namespace            *string `url:"namespace,omitempty",json:"namespace,omitempty"`                         // Namespace.
	Nocow                *bool   `url:"nocow,omitempty",json:"nocow,omitempty"`                                 // Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	Nodes                *string `url:"nodes,omitempty",json:"nodes,omitempty"`                                 // List of cluster node names.
	Nowritecache         *bool   `url:"nowritecache,omitempty",json:"nowritecache,omitempty"`                   // disable write caching on the target
	Options              *string `url:"options,omitempty",json:"options,omitempty"`                             // NFS mount options (see 'man nfs')
	Password             *string `url:"password,omitempty",json:"password,omitempty"`                           // Password for accessing the share/datastore.
	Path                 *string `url:"path,omitempty",json:"path,omitempty"`                                   // File system path.
	Pool                 *string `url:"pool,omitempty",json:"pool,omitempty"`                                   // Pool.
	Port                 *int    `url:"port,omitempty",json:"port,omitempty"`                                   // For non default port.
	Portal               *string `url:"portal,omitempty",json:"portal,omitempty"`                               // iSCSI portal (IP or DNS name with optional port).
	Preallocation        *string `url:"preallocation,omitempty",json:"preallocation,omitempty"`                 // Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	PruneBackups         *string `url:"prune-backups,omitempty",json:"prune-backups,omitempty"`                 // The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	Saferemove           *bool   `url:"saferemove,omitempty",json:"saferemove,omitempty"`                       // Zero-out data when removing LVs.
	SaferemoveThroughput *string `url:"saferemove_throughput,omitempty",json:"saferemove_throughput,omitempty"` // Wipe throughput (cstream -t parameter value).
	Server               *string `url:"server,omitempty",json:"server,omitempty"`                               // Server IP or DNS name.
	Server2              *string `url:"server2,omitempty",json:"server2,omitempty"`                             // Backup volfile server IP or DNS name.
	Share                *string `url:"share,omitempty",json:"share,omitempty"`                                 // CIFS share.
	Shared               *bool   `url:"shared,omitempty",json:"shared,omitempty"`                               // Mark storage as shared.
	Smbversion           *string `url:"smbversion,omitempty",json:"smbversion,omitempty"`                       // SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Sparse               *bool   `url:"sparse,omitempty",json:"sparse,omitempty"`                               // use sparse volumes
	Subdir               *string `url:"subdir,omitempty",json:"subdir,omitempty"`                               // Subdir to mount.
	TaggedOnly           *bool   `url:"tagged_only,omitempty",json:"tagged_only,omitempty"`                     // Only use logical volumes tagged with 'pve-vm-ID'.
	Target               *string `url:"target,omitempty",json:"target,omitempty"`                               // iSCSI target.
	Thinpool             *string `url:"thinpool,omitempty",json:"thinpool,omitempty"`                           // LVM thin pool LV name.
	Transport            *string `url:"transport,omitempty",json:"transport,omitempty"`                         // Gluster transport: tcp or rdma
	Username             *string `url:"username,omitempty",json:"username,omitempty"`                           // RBD Id.
	Vgname               *string `url:"vgname,omitempty",json:"vgname,omitempty"`                               // Volume group name.
	Volume               *string `url:"volume,omitempty",json:"volume,omitempty"`                               // Glusterfs Volume.
}

type CreateResponse struct {
	Storage string `url:"storage",json:"storage"` // The ID of the created storage.
	Type    string `url:"type",json:"type"`       // The type of the created storage.

	// The following parameters are optional
	Config struct {

		// The following parameters are optional
		EncryptionKey *string `url:"encryption-key,omitempty",json:"encryption-key,omitempty"` // The, possible auto-generated, encryption-key.
	} `url:"config,omitempty",json:"config,omitempty"` // Partial, possible server generated, configuration properties.
}

// Create Create a new storage.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var resp *CreateResponse

	err := c.httpClient.Do(ctx, "/storage", "POST", &resp, req)
	return resp, err
}

type FindRequest struct {
	Storage string `url:"storage",json:"storage"` // The storage identifier.

}

type FindResponse map[string]interface{}

// Find Read storage configuration.
func (c *Client) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	var resp *FindResponse

	err := c.httpClient.Do(ctx, "/storage/{storage}", "GET", &resp, req)
	return resp, err
}

type UpdateRequest struct {
	Storage string `url:"storage",json:"storage"` // The storage identifier.

	// The following parameters are optional
	Blocksize            *string `url:"blocksize,omitempty",json:"blocksize,omitempty"`                         // block size
	Bwlimit              *string `url:"bwlimit,omitempty",json:"bwlimit,omitempty"`                             // Set bandwidth/io limits various operations.
	ComstarHg            *string `url:"comstar_hg,omitempty",json:"comstar_hg,omitempty"`                       // host group for comstar views
	ComstarTg            *string `url:"comstar_tg,omitempty",json:"comstar_tg,omitempty"`                       // target group for comstar views
	Content              *string `url:"content,omitempty",json:"content,omitempty"`                             // Allowed content types.NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	DataPool             *string `url:"data-pool,omitempty",json:"data-pool,omitempty"`                         // Data Pool (for erasure coding only)
	Delete               *string `url:"delete,omitempty",json:"delete,omitempty"`                               // A list of settings you want to delete.
	Digest               *string `url:"digest,omitempty",json:"digest,omitempty"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Disable              *bool   `url:"disable,omitempty",json:"disable,omitempty"`                             // Flag to disable the storage.
	Domain               *string `url:"domain,omitempty",json:"domain,omitempty"`                               // CIFS domain.
	EncryptionKey        *string `url:"encryption-key,omitempty",json:"encryption-key,omitempty"`               // Encryption key. Use 'autogen' to generate one automatically without passphrase.
	Fingerprint          *string `url:"fingerprint,omitempty",json:"fingerprint,omitempty"`                     // Certificate SHA 256 fingerprint.
	Format               *string `url:"format,omitempty",json:"format,omitempty"`                               // Default image format.
	FsName               *string `url:"fs-name,omitempty",json:"fs-name,omitempty"`                             // The Ceph filesystem name.
	Fuse                 *bool   `url:"fuse,omitempty",json:"fuse,omitempty"`                                   // Mount CephFS through FUSE.
	IsMountpoint         *string `url:"is_mountpoint,omitempty",json:"is_mountpoint,omitempty"`                 // Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	Keyring              *string `url:"keyring,omitempty",json:"keyring,omitempty"`                             // Client keyring contents (for external clusters).
	Krbd                 *bool   `url:"krbd,omitempty",json:"krbd,omitempty"`                                   // Always access rbd through krbd kernel module.
	LioTpg               *string `url:"lio_tpg,omitempty",json:"lio_tpg,omitempty"`                             // target portal group for Linux LIO targets
	MasterPubkey         *string `url:"master-pubkey,omitempty",json:"master-pubkey,omitempty"`                 // Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	MaxProtectedBackups  *int    `url:"max-protected-backups,omitempty",json:"max-protected-backups,omitempty"` // Maximal number of protected backups per guest. Use '-1' for unlimited.
	Maxfiles             *int    `url:"maxfiles,omitempty",json:"maxfiles,omitempty"`                           // Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Mkdir                *bool   `url:"mkdir,omitempty",json:"mkdir,omitempty"`                                 // Create the directory if it doesn't exist.
	Monhost              *string `url:"monhost,omitempty",json:"monhost,omitempty"`                             // IP addresses of monitors (for external clusters).
	Mountpoint           *string `url:"mountpoint,omitempty",json:"mountpoint,omitempty"`                       // mount point
	Namespace            *string `url:"namespace,omitempty",json:"namespace,omitempty"`                         // Namespace.
	Nocow                *bool   `url:"nocow,omitempty",json:"nocow,omitempty"`                                 // Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	Nodes                *string `url:"nodes,omitempty",json:"nodes,omitempty"`                                 // List of cluster node names.
	Nowritecache         *bool   `url:"nowritecache,omitempty",json:"nowritecache,omitempty"`                   // disable write caching on the target
	Options              *string `url:"options,omitempty",json:"options,omitempty"`                             // NFS mount options (see 'man nfs')
	Password             *string `url:"password,omitempty",json:"password,omitempty"`                           // Password for accessing the share/datastore.
	Pool                 *string `url:"pool,omitempty",json:"pool,omitempty"`                                   // Pool.
	Port                 *int    `url:"port,omitempty",json:"port,omitempty"`                                   // For non default port.
	Preallocation        *string `url:"preallocation,omitempty",json:"preallocation,omitempty"`                 // Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	PruneBackups         *string `url:"prune-backups,omitempty",json:"prune-backups,omitempty"`                 // The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	Saferemove           *bool   `url:"saferemove,omitempty",json:"saferemove,omitempty"`                       // Zero-out data when removing LVs.
	SaferemoveThroughput *string `url:"saferemove_throughput,omitempty",json:"saferemove_throughput,omitempty"` // Wipe throughput (cstream -t parameter value).
	Server               *string `url:"server,omitempty",json:"server,omitempty"`                               // Server IP or DNS name.
	Server2              *string `url:"server2,omitempty",json:"server2,omitempty"`                             // Backup volfile server IP or DNS name.
	Shared               *bool   `url:"shared,omitempty",json:"shared,omitempty"`                               // Mark storage as shared.
	Smbversion           *string `url:"smbversion,omitempty",json:"smbversion,omitempty"`                       // SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Sparse               *bool   `url:"sparse,omitempty",json:"sparse,omitempty"`                               // use sparse volumes
	Subdir               *string `url:"subdir,omitempty",json:"subdir,omitempty"`                               // Subdir to mount.
	TaggedOnly           *bool   `url:"tagged_only,omitempty",json:"tagged_only,omitempty"`                     // Only use logical volumes tagged with 'pve-vm-ID'.
	Transport            *string `url:"transport,omitempty",json:"transport,omitempty"`                         // Gluster transport: tcp or rdma
	Username             *string `url:"username,omitempty",json:"username,omitempty"`                           // RBD Id.
}

type UpdateResponse struct {
	Storage string `url:"storage",json:"storage"` // The ID of the created storage.
	Type    string `url:"type",json:"type"`       // The type of the created storage.

	// The following parameters are optional
	Config struct {

		// The following parameters are optional
		EncryptionKey *string `url:"encryption-key,omitempty",json:"encryption-key,omitempty"` // The, possible auto-generated, encryption-key.
	} `url:"config,omitempty",json:"config,omitempty"` // Partial, possible server generated, configuration properties.
}

// Update Update storage configuration.
func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp *UpdateResponse

	err := c.httpClient.Do(ctx, "/storage/{storage}", "PUT", &resp, req)
	return resp, err
}

type DeleteRequest struct {
	Storage string `url:"storage",json:"storage"` // The storage identifier.

}

type DeleteResponse map[string]interface{}

// Delete Delete storage configuration.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp *DeleteResponse

	err := c.httpClient.Do(ctx, "/storage/{storage}", "DELETE", &resp, req)
	return resp, err
}

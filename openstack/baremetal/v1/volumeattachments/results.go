package volumeattachments

import (
	"github.com/gophercloud/gophercloud"
)

type volumeattachmentResult struct {
	gophercloud.Result
}

// Extract interprets any volumeattachmentResult as a VolumeTarget, if possible.
func (r volumeattachmentResult) Extract() (*VolumeTarget, error) {
	var s VolumeTarget
	err := r.ExtractInto(&s)
	return &s, err
}

func (r volumeattachmentResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// VolumeTarget represents a volume target in the OpenStack Bare Metal API.
type VolumeTarget struct {
	// UUID for the resource.
	UUID string `json:"uuid"`

	// The type of volume target such as ‘iscsi’ and ‘fibre_channel’.
	VolumeType string `json:"volume_type"`

	// A set of physical information of the volume such as the identifier
	// (eg. IQN) and LUN number of the volume. This information is used to
	// connect the node to the volume by the storage interface.
	// The contents depend on the volume type. 
	Properties map[string]interface{} `json:"properties"`

	// The boot index of the Volume target. “0” indicates that this volume
	// is used as a boot volume.
	BootIndex string `json:"boot_index"`

	// The identifier of the volume. This ID is used by storage interface
	// to distinguish volumes.
	VolumeID string `json:"volume_id"`

	// A set of one or more arbitrary metadata key and value pairs.
	Extra map[string]interface{} `json:"extra"`

	// UUID of the Node this resource belongs to.
	NodeUUID string `json:"node_uuid"`
	
	// The UTC date and time when the resource was created, ISO 8601 format.
	CreatedAt time.Time `json:"created_at"`

	// The UTC date and time when the resource was updated, ISO 8601 format.
	// May be “null”.
	UpdatedAt time.Time `json:"updated_at"`

	// A list of relative links. Includes the self and bookmark links.
	Links []interface{} `json:"links"`

	// A URL to request a next collection of the resource. This parameter is
	// returned when limit is specified in a request and there remain items.
	Next string `json:"next"`
}

// CreateResult is the response from a Create operation.
type CreateResult struct {
	volumeattachmentResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

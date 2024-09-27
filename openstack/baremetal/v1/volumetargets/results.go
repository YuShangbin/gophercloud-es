package volumetargets

import (
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type volumetargetResult struct {
	gophercloud.Result
}

func (r volumetargetResult) Extract() (*VolumeTarget, error) {
	var s VolumeTarget
	err := r.ExtractInto(&s)
	return &s, err
}

func (r volumetargetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

func ExtractVolumeTargetsInto(r pagination.Page, v interface{}) error {
	return r.(VolumeTargetPage).Result.ExtractIntoSlicePtr(v, "targets")
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

// VolumeTargetPage abstracts the raw results of making a List() request against
// the API.
type VolumeTargetPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no VolumeTarget results.
func (r VolumeTargetPage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	s, err := ExtractVolumeTargets(r)
	return len(s) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r VolumeTargetPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"targets_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractVolumeTargets interprets the results of a single page from a List() call,
// producing a slice of VolumeTarget entities.
func ExtractVolumeTargets(r pagination.Page) ([]VolumeTarget, error) {
	var s []VolumeTarget
	err := ExtractVolumeTargetsInto(r, &s)
	return s, err
}

package volumeattachments

import (
	"github.com/gophercloud/gophercloud"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeAttachmentCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies volume attach parameters.
type CreateOpts struct {
	// The type for a volume target, e.g. “iscsi”.
	VolumeType string `json:"volume_type,omitempty"`

    // The boot index for a volume target, e.g. “0”.
	BootIndex string `json:"boot_index,omitempty"`

	// The UUID for a volume.
	VolumeID string `json:"volume_id,omitempty"`
}

// ToVolumeAttachmentCreateMap assembles a request body based on the contents of a CreateOpts.
func (opts CreateOpts) ToVolumeAttachmentCreateMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Create requests a volume to be attached to a node.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder, nodeID string) (r CreateResult) {
	reqBody, err := opts.ToVolumeAttachmentCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := client.Post(createURL(client, nodeID), reqBody, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete requests a volume to be detached from a node.
func Delete(client *gophercloud.ServiceClient, nodeID string, volumeID string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, nodeID, volumeID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

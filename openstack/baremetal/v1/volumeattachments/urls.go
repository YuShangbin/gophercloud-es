package volumeattachments

import "github.com/gophercloud/gophercloud"

func createURL(client *gophercloud.ServiceClient, nodeID string) string {
	return client.ServiceURL("nodes/" + nodeID + "/volume_attachments")
}

func deleteURL(client *gophercloud.ServiceClient, nodeID string, volumeID string) string {
	return client.ServiceURL("nodes/" + nodeID + "/volume_attachments/" + volumeID)
}

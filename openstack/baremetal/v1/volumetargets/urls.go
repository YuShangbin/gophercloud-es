package volumetargets

import "github.com/gophercloud/gophercloud"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("volume/targets")
}

func listURL(client *gophercloud.ServiceClient) string {
	return createURL(client)
}

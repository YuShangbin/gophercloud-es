package volumetargets

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToVolumeTargetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the node attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListOpts struct {
	// Filter the list by the name or uuid of the Node
	Node string `q:"node"`

	// One or more fields to be returned in the response.
	Fields []string `q:"fields"`

	// Requests a page size of items.
	Limit int `q:"limit"`

	// The ID of the last-seen item
	Marker string `q:"marker"`

	// Sorts the response by the requested sort direction.
	// Valid value is asc (ascending) or desc (descending). Default is asc.
	SortDir string `q:"sort_dir"`

	// Sorts the response by the this attribute value. Default is id.
	SortKey string `q:"sort_key"`
}

// ToVolumeTargetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeTargetListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List makes a request against the API to list volume targets accessible to you.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeTargetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumeTargetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

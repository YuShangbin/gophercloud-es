package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/baremetal/v1/volumetargets"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListVolumeTargets(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleVolumeTargetListSuccessfully(t)

	pages := 0
	err := volumetargets.List(client.ServiceClient(), volumetargets.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := volumetargets.ExtractVolumeTargets(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 1 {
			t.Fatalf("Expected 1 volumetargets, got %d", len(actual))
		}
		th.AssertEquals(t, "bd4d008c-7d31-463d-abf9-6c23d9d55f7f", actual[0].UUID)

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListOpts(t *testing.T) {
	opts := volumetargets.ListOpts{
		Fields: []string{"uuid", "volume_id", "properties"},
	}

	// Regular ListOpts can
	query, err := opts.ToVolumeTargetListQuery()
	th.AssertEquals(t, query, "?fields=uuid&fields=volume_id&fields=properties")
	th.AssertNoErr(t, err)
}
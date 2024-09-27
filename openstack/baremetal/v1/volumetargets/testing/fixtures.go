package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/baremetal/v1/volumetargets"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// VolumeTargetListBody contains the canned body of a volumetargets.List response.
const VolumeTargetListBody = `
{
    "targets": [
        {
            "boot_index": "0",
            "created_at": "2016-08-18T22:28:48.643434+11:11",
            "extra": {},
            "links": [
                {
                    "href": "http://127.0.0.1:6385/v1/volume/targets/bd4d008c-7d31-463d-abf9-6c23d9d55f7f",
                    "rel": "self"
                },
                {
                    "href": "http://127.0.0.1:6385/volume/targets/bd4d008c-7d31-463d-abf9-6c23d9d55f7f",
                    "rel": "bookmark"
                }
            ],
            "node_uuid": "6d85703a-565d-469a-96ce-30b6de53079d",
            "properties": {},
            "updated_at": null,
            "uuid": "bd4d008c-7d31-463d-abf9-6c23d9d55f7f",
            "volume_id": "04452bed-5367-4202-8bf5-de4335ac56d2",
            "volume_type": "iscsi"
        }
    ]
}
`

var (
	fooCreated, _ = time.Parse(time.RFC3339, "2016-08-18T22:28:48.643434+11:11")
	VolumeTargetFoo       = volumetargets.VolumeTarget{
		UUID:                "bd4d008c-7d31-463d-abf9-6c23d9d55f7f",
		VolumeType:          "iscsi",
		Properties:          map[string]interface{}{},
		BootIndex:           "0",
        VolumeID:            "04452bed-5367-4202-8bf5-de4335ac56d2",
        Extra:               map[string]interface{}{},
		NodeUUID:            "6d85703a-565d-469a-96ce-30b6de53079d",
		CreatedAt:           fooCreated,
		UpdatedAt:           nil,
		Links:               []interface{}{map[string]interface{}{"href": "http://192.168.0.8/baremetal/v1/volume/targets/bd4d008c-7d31-463d-abf9-6c23d9d55f7f", "rel": "self"}, map[string]interface{}{"href": "http://192.168.0.8/baremetal/volume/targets/bd4d008c-7d31-463d-abf9-6c23d9d55f7f", "rel": "bookmark"}},
        Next:                nil,
	}
)

// HandleVolumeTargetListSuccessfully sets up the test server to respond to a volume target List request.
func HandleVolumeTargetListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/volume/targets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()

		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, VolumeTargetListBody)

		case "bd4d008c-7d31-463d-abf9-6c23d9d55f7f":
			fmt.Fprintf(w, `{ "targets": [] }`)
		default:
			t.Fatalf("/volume/targets invoked with unexpected marker=[%s]", marker)
		}
	})
}

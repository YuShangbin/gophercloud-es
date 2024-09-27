/*
	Package volumetargets contains the functionality to Listing of bare metal
	VolumeTarget resources

	API reference: https://developer.openstack.org/api-ref/baremetal/#list-volume-targets

Example to List VolumeTargets

		listOpts := volumetargets.ListOpts{
	 		Limit: 10,
		}

	 	volumetargets.List(client, listOpts).EachPage(func(page pagination.Page) (bool, error) {
	 		volumetargetList, err := volumetargets.ExtractVolumeTargets(page)
	 		if err != nil {
	 			return false, err
	 		}

	 		for _, n := range volumetargetList {
	 			// Do something
	 		}

	 		return true, nil
	 	})
*/
package volumetargets
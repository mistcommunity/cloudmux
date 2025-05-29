
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type DiskListOptions struct {
		StorageId  string
		InstanceId string
	}
	shellutils.R(&DiskListOptions{}, "disk-list", "list disks", func(cli *nutanix.SRegion, args *DiskListOptions) error {
		disks, err := cli.GetDisks(args.StorageId, args.InstanceId)
		if err != nil {
			return err
		}
		printList(disks, 0, 0, 0, []string{})
		return nil
	})

	type DiskIdOptions struct {
		ID string
	}

	shellutils.R(&DiskIdOptions{}, "disk-show", "show disk", func(cli *nutanix.SRegion, args *DiskIdOptions) error {
		disk, err := cli.GetDisk(args.ID)
		if err != nil {
			return err
		}
		printObject(disk)
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type DiskListOptions struct {
		Instance         string `help:"Instance ID"`
		Zone             string `help:"Zone ID"`
		Category         string `help:"Disk category"`
		SnapshotpolicyId string
	}
	shellutils.R(&DiskListOptions{}, "disk-list", "List disks", func(cli *aliyun.SRegion, args *DiskListOptions) error {
		disks, err := cli.GetDisks(args.Instance, args.Zone, args.Category, nil, args.SnapshotpolicyId)
		if err != nil {
			return err
		}
		printList(disks, 0, 0, 0, []string{})
		return nil
	})

	type DiskDeleteOptions struct {
		ID string `help:"Instance ID"`
	}
	shellutils.R(&DiskDeleteOptions{}, "disk-delete", "List disks", func(cli *aliyun.SRegion, args *DiskDeleteOptions) error {
		e := cli.DeleteDisk(args.ID)
		if e != nil {
			return e
		}
		return nil
	})
}

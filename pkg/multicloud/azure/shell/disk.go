
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type DiskListOptions struct {
	}
	shellutils.R(&DiskListOptions{}, "disk-list", "List disks", func(cli *azure.SRegion, args *DiskListOptions) error {
		disks, err := cli.GetDisks()
		if err != nil {
			return err
		}
		printList(disks, len(disks), 0, 0, []string{})
		return nil
	})

	type DiskCreateOptions struct {
		NAME          string `help:"Disk name"`
		STORAGETYPE   string `help:"Storage type" choices:"Standard_LRS|Premium_LRS|StandardSSD_LRS"`
		SizeGb        int32  `help:"Disk size"`
		Image         string `help:"Image id"`
		SnapshotId    string `help:"Create disk by snapshot"`
		ResourceGroup string `help:"ResourceGroup Name"`
	}

	shellutils.R(&DiskCreateOptions{}, "disk-create", "Create disk", func(cli *azure.SRegion, args *DiskCreateOptions) error {
		disk, err := cli.CreateDisk(args.STORAGETYPE, args.NAME, args.SizeGb, args.Image, args.SnapshotId, args.ResourceGroup)
		if err != nil {
			return err
		}
		printObject(disk)
		return nil
	})

	type DiskOptions struct {
		ID string `help:"Disk ID"`
	}

	shellutils.R(&DiskOptions{}, "disk-show", "Show disk", func(cli *azure.SRegion, args *DiskOptions) error {
		disk, err := cli.GetDisk(args.ID)
		if err != nil {
			return err
		}
		printObject(disk)
		return nil
	})

	shellutils.R(&DiskOptions{}, "disk-delete", "Delete disks", func(cli *azure.SRegion, args *DiskOptions) error {
		return cli.DeleteDisk(args.ID)
	})

	type DiskResizeOptions struct {
		ID   string `help:"Disk ID"`
		SIZE int32  `help:"Disk SizeGb"`
	}

	shellutils.R(&DiskResizeOptions{}, "disk-resize", "Delete disks", func(cli *azure.SRegion, args *DiskResizeOptions) error {
		return cli.ResizeDisk(args.ID, args.SIZE)
	})
}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type DiskListOptions struct {
		ZONE        string
		StorageType string
		MaxResults  int
		PageToken   string
	}
	shellutils.R(&DiskListOptions{}, "disk-list", "List disks", func(cli *google.SRegion, args *DiskListOptions) error {
		disks, err := cli.GetDisks(args.ZONE, args.StorageType, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(disks, 0, 0, 0, nil)
		return nil
	})

	type DiskIdOptions struct {
		ID string
	}
	shellutils.R(&DiskIdOptions{}, "disk-show", "Show disk", func(cli *google.SRegion, args *DiskIdOptions) error {
		disk, err := cli.GetDisk(args.ID)
		if err != nil {
			return err
		}
		printObject(disk)
		return nil
	})

	shellutils.R(&DiskIdOptions{}, "disk-delete", "Delete disk", func(cli *google.SRegion, args *DiskIdOptions) error {
		return cli.Delete(args.ID)
	})

	type DiskCreateOptions struct {
		NAME         string
		Desc         string
		ZONE         string
		SIZE_GB      int
		Image        string
		STORAGE_TYPE string `choices:"pd-standard|pd-ssd"`
	}

	shellutils.R(&DiskCreateOptions{}, "disk-create", "Create disks", func(cli *google.SRegion, args *DiskCreateOptions) error {
		disk, err := cli.CreateDisk(args.NAME, args.SIZE_GB, args.ZONE, args.STORAGE_TYPE, args.Image, args.Desc)
		if err != nil {
			return err
		}
		printObject(disk)
		return nil
	})

	type DiskResizeOptions struct {
		ID      string
		SIZE_GB int
	}

	shellutils.R(&DiskResizeOptions{}, "disk-resize", "Resize disk", func(cli *google.SRegion, args *DiskResizeOptions) error {
		return cli.ResizeDisk(args.ID, args.SIZE_GB)
	})

	type RegionDiskListOptions struct {
		StorageType string
		MaxResults  int
		PageToken   string
	}
	shellutils.R(&RegionDiskListOptions{}, "region-disk-list", "List region disks", func(cli *google.SRegion, args *RegionDiskListOptions) error {
		disks, err := cli.GetRegionDisks(args.StorageType, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(disks, 0, 0, 0, nil)
		return nil
	})

	type RegionDiskShowOptions struct {
		ID string
	}
	shellutils.R(&RegionDiskShowOptions{}, "region-disk-show", "Show region disk", func(cli *google.SRegion, args *RegionDiskShowOptions) error {
		disk, err := cli.GetRegionDisk(args.ID)
		if err != nil {
			return err
		}
		printObject(disk)
		return nil
	})

}

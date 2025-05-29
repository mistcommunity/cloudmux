package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type DiskListOptions struct {
		Instance   string `help:"Instance ID"`
		Zone       string `help:"Zone ID"`
		VolumeType string `help:"Disk category" choices:"gp2|gp3|io1|io2|st1|sc1|standard"`
		DiskIds    []string
	}
	shellutils.R(&DiskListOptions{}, "disk-list", "List disks", func(cli *aws.SRegion, args *DiskListOptions) error {
		disks, err := cli.GetDisks(args.Instance, args.Zone, args.VolumeType, args.DiskIds)
		if err != nil {
			return err
		}
		printList(disks, 0, 0, 0, []string{})
		return nil
	})

	type DiskDeleteOptions struct {
		ID string `help:"Disk ID"`
	}
	shellutils.R(&DiskDeleteOptions{}, "disk-delete", "List disks", func(cli *aws.SRegion, args *DiskDeleteOptions) error {
		e := cli.DeleteDisk(args.ID)
		if e != nil {
			return e
		}
		return nil
	})

	type DiskResizeOptions struct {
		ID      string `help:"Disk ID"`
		SIZE_GB int64
	}
	shellutils.R(&DiskResizeOptions{}, "disk-resize", "List disks", func(cli *aws.SRegion, args *DiskResizeOptions) error {
		e := cli.ResizeDisk(args.ID, args.SIZE_GB)
		if e != nil {
			return e
		}
		return nil
	})

	type VolumeCreateOptions struct {
		Name       string
		Desc       string
		VolumeType string `choices:"gp2|gp3|io1|io2|st1|sc1|standard" default:"gp2"`
		ZoneId     string
		SizeGb     int `default:"10"`
		Iops       int
		Throughput int
		SnapshotId string
	}

	shellutils.R(&VolumeCreateOptions{}, "disk-create", "create a volume", func(cli *aws.SRegion, args *VolumeCreateOptions) error {
		volume, err := cli.CreateDisk(args.ZoneId, args.VolumeType, args.Name, args.SizeGb, args.Iops, args.Throughput, args.SnapshotId, args.Desc)
		if err != nil {
			return err
		}
		printObject(volume)
		return nil
	})

}

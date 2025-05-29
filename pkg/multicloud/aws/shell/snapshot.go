package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type SnapshotListOptions struct {
		DiskId string   `help:"Disk ID"`
		Ids    []string `helo:"Snapshot ids"`
		Name   string   `help:"Snapshot Name"`
	}
	shellutils.R(&SnapshotListOptions{}, "snapshot-list", "List snapshot", func(cli *aws.SRegion, args *SnapshotListOptions) error {
		snapshots, err := cli.GetSnapshots(args.DiskId, args.Name, args.Ids)
		if err != nil {
			return err
		}
		printList(snapshots, 0, 0, 0, []string{})
		return nil
	})

	type SnapshotDeleteOptions struct {
		ID string `help:"Snapshot ID"`
	}

	shellutils.R(&SnapshotDeleteOptions{}, "snapshot-delete", "Delete snapshot", func(cli *aws.SRegion, args *SnapshotDeleteOptions) error {
		return cli.DeleteSnapshot(args.ID)
	})

	type SnapshotCreateOptions struct {
		DiskId string `help:"Disk ID"`
		Name   string `help:"Snapeshot Name"`
		Desc   string `help:"Snapshot Desc"`
	}

	shellutils.R(&SnapshotCreateOptions{}, "snapshot-create", "Create snapshot", func(cli *aws.SRegion, args *SnapshotCreateOptions) error {
		_, err := cli.CreateSnapshot(args.DiskId, args.Name, args.Desc)
		return err
	})

}

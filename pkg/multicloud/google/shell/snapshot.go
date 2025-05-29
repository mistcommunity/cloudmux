
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type SnapshotListOptions struct {
		Disk       string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&SnapshotListOptions{}, "snapshot-list", "List snapshots", func(cli *google.SRegion, args *SnapshotListOptions) error {
		snapshots, err := cli.GetSnapshots(args.Disk, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(snapshots, 0, 0, 0, nil)
		return nil
	})

	type SnapshotIdOptions struct {
		ID string
	}
	shellutils.R(&SnapshotIdOptions{}, "snapshot-show", "Show snapshot", func(cli *google.SRegion, args *SnapshotIdOptions) error {
		snapshot, err := cli.GetSnapshot(args.ID)
		if err != nil {
			return err
		}
		printObject(snapshot)
		return nil
	})

	shellutils.R(&SnapshotIdOptions{}, "snapshot-delete", "Delete snapshot", func(cli *google.SRegion, args *SnapshotIdOptions) error {
		return cli.Delete(args.ID)
	})

	type SnapshotCreateOptions struct {
		NAME string
		Desc string
		DISK string
	}

	shellutils.R(&SnapshotCreateOptions{}, "snapshot-create", "Create snapshot", func(cli *google.SRegion, args *SnapshotCreateOptions) error {
		snapshot, err := cli.CreateSnapshot(args.DISK, args.NAME, args.Desc)
		if err != nil {
			return err
		}
		printObject(snapshot)
		return nil
	})

}

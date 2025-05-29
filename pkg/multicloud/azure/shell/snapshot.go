
package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type SnapshotListOptions struct {
	}
	shellutils.R(&SnapshotListOptions{}, "snapshot-list", "List snapshot", func(cli *azure.SRegion, args *SnapshotListOptions) error {
		snapshots, err := cli.ListSnapshots()
		if err != nil {
			return err
		}
		printList(snapshots, len(snapshots), 0, 0, []string{})
		return nil
	})

	type SnapshotCreateOptions struct {
		DISK string `help:"SourceID"`
		NAME string `help:"Snapshot name"`
		Desc string `help:"Snapshot description"`
	}

	shellutils.R(&SnapshotCreateOptions{}, "snapshot-create", "Create snapshot", func(cli *azure.SRegion, args *SnapshotCreateOptions) error {
		snapshot, err := cli.CreateSnapshot(args.DISK, args.NAME, args.Desc)
		if err != nil {
			return err
		}
		printObject(snapshot)
		return nil
	})

	type SnapshotOptions struct {
		ID string `help:"Snapshot ID"`
	}

	shellutils.R(&SnapshotOptions{}, "snapshot-delete", "Delete snapshot", func(cli *azure.SRegion, args *SnapshotOptions) error {
		return cli.DeleteSnapshot(args.ID)
	})

	shellutils.R(&SnapshotOptions{}, "snapshot-show", "List snapshot", func(cli *azure.SRegion, args *SnapshotOptions) error {
		snapshot, err := cli.GetSnapshot(args.ID)
		if err != nil {
			return err
		}
		printObject(snapshot)
		return nil
	})

	shellutils.R(&SnapshotOptions{}, "snapshot-grant-access", "Grant access for snapshot", func(cli *azure.SRegion, args *SnapshotOptions) error {
		if uri, err := cli.GrantAccessSnapshot(args.ID); err != nil {
			return err
		} else {
			fmt.Printf("download link %s", uri)
			return nil
		}
	})

}

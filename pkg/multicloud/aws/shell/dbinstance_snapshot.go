package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type DBInstanceSnapshotListOptions struct {
		DBInstanceId string
		Id           string
	}

	shellutils.R(&DBInstanceSnapshotListOptions{}, "dbinstance-snapshot-list", "List rds intance snapshots", func(cli *aws.SRegion, args *DBInstanceSnapshotListOptions) error {
		snapshots, err := cli.GetDBInstanceSnapshots(args.DBInstanceId, args.Id)
		if err != nil {
			return err
		}
		printList(snapshots, 0, 0, 0, []string{})
		return nil
	})

}

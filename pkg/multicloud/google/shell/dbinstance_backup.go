
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type DBInstanceBackupListOptions struct {
		RDS string
	}
	shellutils.R(&DBInstanceBackupListOptions{}, "dbinstance-backup-list", "List dbinstance backup", func(cli *google.SRegion, args *DBInstanceBackupListOptions) error {
		backups, err := cli.GetDBInstanceBackups(args.RDS)
		if err != nil {
			return err
		}
		printList(backups, 0, 0, 0, nil)
		return nil
	})

	type DBInstanceBackupCreateOptions struct {
		RDS  string
		NAME string
		Desc string
	}

	shellutils.R(&DBInstanceBackupCreateOptions{}, "dbinstance-backup-create", "Create dbinstance backup", func(cli *google.SRegion, args *DBInstanceBackupCreateOptions) error {
		return cli.CreateDBInstanceBackup(args.RDS, args.NAME, args.Desc)
	})
}


package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type DBInstanceBackupJobListOptions struct {
		INSTANCE string
		JobId    string
	}
	shellutils.R(&DBInstanceBackupJobListOptions{}, "dbinstance-backup-job-list", "Get dbinstance backup jobs", func(cli *aliyun.SRegion, args *DBInstanceBackupJobListOptions) error {
		jobs, err := cli.GetDBInstanceBackupJobs(args.INSTANCE, args.JobId)
		if err != nil {
			return err
		}
		printObject(jobs)
		return nil
	})

	type DBInstanceBackupOptions struct {
		INSTANCE string
		BACKUP   string
	}

	shellutils.R(&DBInstanceBackupOptions{}, "dbinstance-backup-delete", "Delete dbinstance backup", func(cli *aliyun.SRegion, args *DBInstanceBackupOptions) error {
		return cli.DeleteDBInstanceBackup(args.INSTANCE, args.BACKUP)
	})

	shellutils.R(&DBInstanceBackupOptions{}, "dbinstance-backup-job-list", "Get dbinstance backup jobs", func(cli *aliyun.SRegion, args *DBInstanceBackupOptions) error {
		return cli.DeleteDBInstanceBackup(args.INSTANCE, args.BACKUP)
	})

	type DBInstanceIdExtraOptions struct {
		ID     string `help:"ID of instances to show"`
		Limit  int    `help:"page size"`
		Offset int    `help:"page offset"`
	}

	shellutils.R(&DBInstanceIdExtraOptions{}, "dbinstance-backup-list", "List dbintance backups", func(cli *aliyun.SRegion, args *DBInstanceIdExtraOptions) error {
		backups, _, err := cli.GetDBInstanceBackups(args.ID, "", args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(backups, 0, 0, 0, []string{})
		return nil
	})

	type DBInstanceBackupCreateOptions struct {
		INSTANCE_ID string
		Database    []string
	}

	shellutils.R(&DBInstanceBackupCreateOptions{}, "dbinstance-backup-create", "Create dbintance backup", func(cli *aliyun.SRegion, args *DBInstanceBackupCreateOptions) error {
		backupId, err := cli.CreateDBInstanceBackup(args.INSTANCE_ID, args.Database)
		if err != nil {
			return err
		}
		fmt.Println("backup id: ", backupId)
		return nil
	})

}

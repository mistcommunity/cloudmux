
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type DBInstanceDatabasesOptions struct {
		ID string
	}

	shellutils.R(&DBInstanceDatabasesOptions{}, "dbinstance-database-show", "Show rds intance database", func(cli *azure.SRegion, args *DBInstanceDatabasesOptions) error {
		databases, err := cli.ListSDBInstanceDatabase(args.ID)
		if err != nil {
			return err
		}
		printList(databases, 0, 0, 0, []string{})
		return nil
	})

}

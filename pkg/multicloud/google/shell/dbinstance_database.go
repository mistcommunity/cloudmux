
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type DBInstanceDatabaseListOptions struct {
		RDS string
	}
	shellutils.R(&DBInstanceDatabaseListOptions{}, "dbinstance-database-list", "List dbinstance database", func(cli *google.SRegion, args *DBInstanceDatabaseListOptions) error {
		databases, err := cli.GetDBInstanceDatabases(args.RDS)
		if err != nil {
			return err
		}
		printList(databases, 0, 0, 0, nil)
		return nil
	})

	type DBInstanceDatabaseCreateOptions struct {
		RDS     string
		NAME    string
		CHARSET string
	}

	shellutils.R(&DBInstanceDatabaseCreateOptions{}, "dbinstance-database-create", "Create dbinstance database", func(cli *google.SRegion, args *DBInstanceDatabaseCreateOptions) error {
		return cli.CreateDatabase(args.RDS, args.NAME, args.CHARSET)
	})

}

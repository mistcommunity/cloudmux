
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type DBInstanceListOptions struct {
		TYPE string `help:"dbinstance type" choices:"Microsoft.DBForMariaDB/servers|Microsoft.DBforMySQL/servers|Microsoft.DBforMySQL/flexibleServers|Microsoft.DBforPostgreSQL/servers|Microsoft.DBforPostgreSQL/flexibleServers|Microsoft.Sql/servers|Microsoft.Sql/managedInstances"`
	}
	shellutils.R(&DBInstanceListOptions{}, "dbinstance-list", "List rds intances", func(cli *azure.SRegion, args *DBInstanceListOptions) error {
		instances, err := cli.ListDBInstance(args.TYPE)
		if err != nil {
			return err
		}
		printList(instances, 0, 0, len(instances), []string{})
		return nil
	})

	type DBInstanceIdOptions struct {
		ID string
	}

	shellutils.R(&DBInstanceIdOptions{}, "dbinstance-show", "Show rds intance", func(cli *azure.SRegion, args *DBInstanceIdOptions) error {
		instance, err := cli.GetDBInstanceById(args.ID)
		if err != nil {
			return err
		}
		printObject(instance)
		return nil
	})

}

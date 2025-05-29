
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type DBInstanceAccountListOptions struct {
		RDS string
	}
	shellutils.R(&DBInstanceAccountListOptions{}, "dbinstance-account-list", "List dbinstance account", func(cli *google.SRegion, args *DBInstanceAccountListOptions) error {
		accounts, err := cli.GetDBInstanceAccounts(args.RDS)
		if err != nil {
			return err
		}
		printList(accounts, 0, 0, 0, nil)
		return nil
	})

	type DBInstanceAccountCreateOptions struct {
		RDS      string
		NAME     string
		PASSWORD string
		Host     string
	}
	shellutils.R(&DBInstanceAccountCreateOptions{}, "dbinstance-account-create", "Create dbinstance account", func(cli *google.SRegion, args *DBInstanceAccountCreateOptions) error {
		return cli.CreateDBInstanceAccount(args.RDS, args.NAME, args.PASSWORD, args.Host)
	})
}

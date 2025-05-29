
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type DBInstanceParameterOptions struct {
		ID string
	}

	shellutils.R(&DBInstanceParameterOptions{}, "dbinstance-parameter-show", "Show rds intance parameter", func(cli *azure.SRegion, args *DBInstanceParameterOptions) error {
		parameters, err := cli.ListDBInstanceConfiguration(args.ID)
		if err != nil {
			return err
		}
		printList(parameters, 0, 0, 0, []string{})
		return nil
	})

}

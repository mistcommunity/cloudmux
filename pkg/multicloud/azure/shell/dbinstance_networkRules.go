
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type DBInstanceNetworkOptions struct {
		ID string
	}

	shellutils.R(&DBInstanceNetworkOptions{}, "dbinstance-network-show", "Show rds intance networks", func(cli *azure.SRegion, args *DBInstanceNetworkOptions) error {
		networks, err := cli.ListDBInstanceVirtualNetworkRule(args.ID)
		if err != nil {
			return err
		}
		printList(networks, 0, 0, 0, []string{})
		return nil
	})

}

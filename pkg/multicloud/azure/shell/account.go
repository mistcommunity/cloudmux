
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type AccountListOptions struct {
	}
	shellutils.R(&AccountListOptions{}, "account-list", "List sub account", func(cli *azure.SRegion, args *AccountListOptions) error {
		if accounts, err := cli.GetClient().GetSubAccounts(); err != nil {
			return err
		} else {
			printObject(accounts)
			return nil
		}
	})
}

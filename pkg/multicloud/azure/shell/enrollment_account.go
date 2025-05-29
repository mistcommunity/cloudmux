
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type EnrollmentAccountListOptions struct {
	}
	shellutils.R(&EnrollmentAccountListOptions{}, "enrollment-account-list", "List enrollment accounts", func(cli *azure.SRegion, args *EnrollmentAccountListOptions) error {
		accounts, err := cli.GetClient().GetEnrollmentAccounts()
		if err != nil {
			return err
		}
		printList(accounts, 0, 0, 0, nil)
		return nil
	})
}

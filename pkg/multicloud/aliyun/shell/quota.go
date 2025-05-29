
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type QuotaListOptions struct {
	}
	shellutils.R(&QuotaListOptions{}, "quota-list", "List quota", func(cli *aliyun.SRegion, args *QuotaListOptions) error {
		quotas, err := cli.GetQuotas()
		if err != nil {
			return err
		}
		printList(quotas, 0, 0, 0, nil)
		return nil
	})
}

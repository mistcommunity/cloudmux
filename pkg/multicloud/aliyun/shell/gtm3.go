
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type GtmListOptions struct {
	}
	shellutils.R(&GtmListOptions{}, "gtm3-instance-list", "List Gtm", func(cli *aliyun.SRegion, args *GtmListOptions) error {
		ret, err := cli.GetClient().ListCloudGtmInstanceConfigs()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type GtmShowOptions struct {
		ID string
	}
	shellutils.R(&GtmShowOptions{}, "gtm3-pool-show", "Show Gtm pool", func(cli *aliyun.SRegion, args *GtmShowOptions) error {
		ret, err := cli.GetClient().DescribeCloudGtmAddressPool(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

}

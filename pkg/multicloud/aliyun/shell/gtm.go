
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type GtmListOptions struct {
	}
	shellutils.R(&GtmListOptions{}, "gtm-instance-list", "List Gtm", func(cli *aliyun.SRegion, args *GtmListOptions) error {
		ret, err := cli.GetClient().DescribeDnsGtmInstances()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type GtmPoolListOptions struct {
		ID string
	}
	shellutils.R(&GtmPoolListOptions{}, "gtm-instance-address-pool-list", "List Gtm address pool", func(cli *aliyun.SRegion, args *GtmPoolListOptions) error {
		ret, err := cli.GetClient().DescribeDnsGtmInstanceAddressPools(args.ID)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type GtmPoolShowOptions struct {
		ID string
	}
	shellutils.R(&GtmPoolShowOptions{}, "gtm-instance-address-pool-show", "Show Gtm address pool", func(cli *aliyun.SRegion, args *GtmPoolShowOptions) error {
		ret, err := cli.GetClient().DescribeDnsGtmInstanceAddressPool(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type AccountBalanceOptions struct {
	}
	shellutils.R(&AccountBalanceOptions{}, "balance", "Get account balance", func(cli *aliyun.SRegion, args *AccountBalanceOptions) error {
		result1, err := cli.GetClient().QueryAccountBalance()
		if err != nil {
			return err
		}
		printObject(result1)

		result2, err := cli.GetClient().QueryCashCoupons()
		if err != nil {
			return err
		}
		if len(result2) > 0 {
			printList(result2, len(result2), 0, 0, nil)
		}

		result3, err := cli.GetClient().QueryPrepaidCards()
		if err != nil {
			return err
		}
		if len(result3) > 0 {
			printList(result3, len(result3), 0, 0, nil)
		}
		return nil
	})

	type AliyunSubscribeBillOptions struct {
		BUCKET string `help:"bucket name to store billing records"`
	}
	shellutils.R(&AliyunSubscribeBillOptions{}, "subscribe-bill", "Subscribe bill to OSS storage", func(cli *aliyun.SRegion, args *AliyunSubscribeBillOptions) error {
		err := cli.GetClient().SubscribeBillToOSS(args.BUCKET)
		if err != nil {
			return err
		}
		return nil
	})

}

package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type LambdaListOptions struct {
	}
	shellutils.R(&LambdaListOptions{}, "lambda-function-list", "List lambda functions", func(cli *aws.SRegion, args *LambdaListOptions) error {
		functions, err := cli.ListFunctions()
		if err != nil {
			return err
		}
		printList(functions, 0, 0, 0, []string{})
		return nil
	})

	type LambdaProvisionedOptions struct {
		NAME    string
		VERSION string
	}
	shellutils.R(&LambdaProvisionedOptions{}, "lambda-function-provisioned-show", "Show lambda function provisioned", func(cli *aws.SRegion, args *LambdaProvisionedOptions) error {
		ret, err := cli.GetProvisionedConcurrencyConfig(args.NAME, args.VERSION)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

}

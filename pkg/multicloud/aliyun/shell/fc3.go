
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type FunctionListOptions struct {
	}
	shellutils.R(&FunctionListOptions{}, "fc3-function-list", "List Function", func(cli *aliyun.SRegion, args *FunctionListOptions) error {
		ret, err := cli.GetFunctions()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type InstanceListOptions struct {
		FUNCTION string
	}
	shellutils.R(&InstanceListOptions{}, "fc3-instance-list", "List Instance", func(cli *aliyun.SRegion, args *InstanceListOptions) error {
		ret, err := cli.GetFunctionInstances(args.FUNCTION)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

}

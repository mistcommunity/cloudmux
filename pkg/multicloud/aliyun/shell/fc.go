
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ServiceListOptions struct {
	}
	shellutils.R(&ServiceListOptions{}, "fc-service-list", "List Service", func(cli *aliyun.SRegion, args *ServiceListOptions) error {
		ret, err := cli.GetFcServices()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type FunctionListOptions struct {
		SERVICE string
	}

	shellutils.R(&FunctionListOptions{}, "fc-function-list", "List Function", func(cli *aliyun.SRegion, args *FunctionListOptions) error {
		ret, err := cli.GetFcFunctions(args.SERVICE)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type InstanceListOptions struct {
		SERVICE  string
		FUNCTION string
	}

	shellutils.R(&InstanceListOptions{}, "fc-instance-list", "List Instance", func(cli *aliyun.SRegion, args *InstanceListOptions) error {
		ret, err := cli.GetFcInstances(args.SERVICE, args.FUNCTION)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

}

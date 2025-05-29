
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type NetworkInterfaceListOptions struct {
		InstanceId string `help:"Id or instance"`
		Status     string
		PageSize   int
		NextToken  string
	}
	shellutils.R(&NetworkInterfaceListOptions{}, "network-interface-list", "List networkinterfaces", func(cli *aliyun.SRegion, args *NetworkInterfaceListOptions) error {
		interfaces, _, err := cli.GetNetworkInterfaces(args.InstanceId, args.Status, args.NextToken, args.PageSize)
		if err != nil {
			return err
		}
		printList(interfaces, 0, 0, 0, nil)
		return nil
	})
}

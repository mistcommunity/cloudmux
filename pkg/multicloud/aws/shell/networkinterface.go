package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type NetworkInterfaceListOptions struct {
		Id string
	}
	shellutils.R(&NetworkInterfaceListOptions{}, "network-interface-list", "List network interfaces", func(cli *aws.SRegion, args *NetworkInterfaceListOptions) error {
		interfaces, err := cli.GetNetworkInterfaces(args.Id)
		if err != nil {
			return err
		}
		printList(interfaces, 0, 0, 0, []string{})
		return nil
	})
}

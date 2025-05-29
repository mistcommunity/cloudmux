
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type NetworkListOptions struct {
	}
	shellutils.R(&NetworkListOptions{}, "network-list", "list networks", func(cli *proxmox.SRegion, args *NetworkListOptions) error {
		networks, err := cli.GetNetworks()
		if err != nil {
			return err
		}
		printList(networks, 0, 0, 0, []string{})
		return nil
	})

	type NetworkIdOptions struct {
		ID string
	}

	shellutils.R(&NetworkIdOptions{}, "network-show", "show network", func(cli *proxmox.SRegion, args *NetworkIdOptions) error {
		ret, err := cli.GetNetwork(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

}

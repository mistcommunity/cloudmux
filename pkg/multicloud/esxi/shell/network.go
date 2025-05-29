
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/esxi"
)

func init() {
	type NetworkListOptions struct {
		// DATACENTER string `help:"List datastores in datacenter"`
	}
	shellutils.R(&NetworkListOptions{}, "network-list", "List networks in datacenter", func(cli *esxi.SESXiClient, args *NetworkListOptions) error {
		nets, err := cli.GetNetworks()
		if err != nil {
			return err
		}
		printList(nets, nil)
		return nil
	})

	shellutils.R(&NetworkListOptions{}, "wire-list", "List wires in datacenter", func(cli *esxi.SESXiClient, args *NetworkListOptions) error {
		vpcs, err := cli.GetIVpcs()
		if err != nil {
			return err
		}
		wires, err := vpcs[0].GetIWires()
		if err != nil {
			return err
		}
		printList(wires, nil)
		return nil
	})
}

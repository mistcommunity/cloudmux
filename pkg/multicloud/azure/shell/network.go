
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type NetworkListOptions struct {
		VPC    string `help:"Vpc Id"`
		Limit  int    `help:"page size"`
		Offset int    `help:"page offset"`
	}
	shellutils.R(&NetworkListOptions{}, "network-list", "List networks", func(cli *azure.SRegion, args *NetworkListOptions) error {
		vpc, err := cli.GetVpc(args.VPC)
		if err != nil {
			return err
		}
		networks := vpc.GetNetworks()
		printList(networks, len(networks), args.Offset, args.Limit, []string{})
		return nil
	})

	type NetworkCreateOptions struct {
		VPC  string
		NAME string
		CIDR string
	}

	shellutils.R(&NetworkCreateOptions{}, "network-create", "Create network", func(cli *azure.SRegion, args *NetworkCreateOptions) error {
		network, err := cli.CreateNetwork(args.VPC, args.NAME, args.CIDR, "")
		if err != nil {
			return err
		}
		printObject(network)
		return nil
	})

	type NetworkInterfaceListOptions struct {
	}

	shellutils.R(&NetworkInterfaceListOptions{}, "network-interface-list", "List network interface", func(cli *azure.SRegion, args *NetworkInterfaceListOptions) error {
		nics, err := cli.GetNetworkInterfaces()
		if err != nil {
			return err
		}
		printList(nics, len(nics), 0, 0, []string{})
		return nil
	})

	type NetworkInterfaceOptions struct {
		ID string `help:"Network ineterface ID"`
	}

	shellutils.R(&NetworkInterfaceOptions{}, "network-interface-show", "Show network interface", func(cli *azure.SRegion, args *NetworkInterfaceOptions) error {
		nic, err := cli.GetNetworkInterface(args.ID)
		if err != nil {
			return err
		}
		printObject(nic)
		return nil
	})

	type NetworkInterfaceCreateOptions struct {
		ResourceGroup string `help:"ResourceGroup Name"`
		NAME          string `help:"Nic interface name"`
		IP            string `help:"Nic private ip address"`
		NETWORK       string `help:"Netowrk ID"`
		SecurityGroup string `help:"SecurityGroup ID"`
	}

	shellutils.R(&NetworkInterfaceCreateOptions{}, "network-interface-create", "Create network interface", func(cli *azure.SRegion, args *NetworkInterfaceCreateOptions) error {
		nic, err := cli.CreateNetworkInterface(args.ResourceGroup, args.NAME, args.IP, args.NETWORK, args.SecurityGroup)
		if err != nil {
			return err
		}
		printObject(nic)
		return nil
	})

}

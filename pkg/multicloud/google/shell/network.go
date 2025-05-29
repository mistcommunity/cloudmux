
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type NetworkListOptions struct {
	}
	shellutils.R(&NetworkListOptions{}, "vpc-list", "List networks", func(cli *google.SRegion, args *NetworkListOptions) error {
		networks, err := cli.GetVpcs()
		if err != nil {
			return err
		}
		printList(networks, 0, 0, 0, nil)
		return nil
	})

	type NetworkIdOptions struct {
		ID string
	}

	shellutils.R(&NetworkIdOptions{}, "vpc-show", "Show network", func(cli *google.SRegion, args *NetworkIdOptions) error {
		vpc, err := cli.GetVpc(args.ID)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

	shellutils.R(&NetworkIdOptions{}, "vpc-delete", "Delete network", func(cli *google.SRegion, args *NetworkIdOptions) error {
		return cli.Delete(args.ID)
	})

	type NetworkCreateOptions struct {
		NAME string
		VPC  string
		CIDR string
		Desc string
	}

	shellutils.R(&NetworkCreateOptions{}, "vpc-create", "Create network", func(cli *google.SRegion, args *NetworkCreateOptions) error {
		network, err := cli.CreateVpc(args.NAME, args.VPC, args.CIDR, args.Desc)
		if err != nil {
			return err
		}
		printObject(network)
		return nil
	})

}

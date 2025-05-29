
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type VpcListOptions struct {
		Limit  int `help:"page size"`
		Offset int `help:"page offset"`
	}
	shellutils.R(&VpcListOptions{}, "vpc-list", "List vpcs", func(cli *azure.SRegion, args *VpcListOptions) error {
		vpcs, err := cli.ListVpcs()
		if err != nil {
			return err
		}
		printList(vpcs, len(vpcs), args.Offset, args.Limit, []string{})
		return nil
	})

	type VpcOptions struct {
		ID string `help:"vpc ID"`
	}

	shellutils.R(&VpcOptions{}, "vpc-show", "Show vpc details", func(cli *azure.SRegion, args *VpcOptions) error {
		vpc, err := cli.GetVpc(args.ID)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

	shellutils.R(&VpcOptions{}, "vpc-delete", "Delete vpc", func(cli *azure.SRegion, args *VpcOptions) error {
		return cli.DeleteVpc(args.ID)
	})

	type VpcCreateOptions struct {
		NAME string `help:"vpc Name"`
		CIDR string `help:"vpc cidr"`
		Desc string `help:"vpc description"`
	}

	shellutils.R(&VpcCreateOptions{}, "vpc-create", "Create vpc", func(cli *azure.SRegion, args *VpcCreateOptions) error {
		opts := &cloudprovider.VpcCreateOptions{
			NAME: args.NAME,
			CIDR: args.CIDR,
			Desc: args.Desc,
		}
		vpc, err := cli.CreateIVpc(opts)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})
}

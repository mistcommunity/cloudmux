
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type EipListOptions struct {
		Offset int `help:"List offset"`
		Limit  int `help:"List limit"`
	}
	shellutils.R(&EipListOptions{}, "eip-list", "List eips", func(cli *azure.SRegion, args *EipListOptions) error {
		eips, err := cli.GetEips()
		if err != nil {
			return err
		}
		printList(eips, len(eips), args.Offset, args.Limit, []string{})
		return nil
	})

	type EipAllocateOptions struct {
		NAME          string `help:"Eip Name"`
		ResourceGroup string `help:"ResourceGroup Name"`
	}
	shellutils.R(&EipAllocateOptions{}, "eip-create", "Allocate an EIP", func(cli *azure.SRegion, args *EipAllocateOptions) error {
		if eip, err := cli.AllocateEIP(args.NAME, args.ResourceGroup); err != nil {
			return err
		} else {
			printObject(eip)
			return nil
		}
	})

	type EipReleaseOptions struct {
		ID string `help:"EIP allocation ID"`
	}
	shellutils.R(&EipReleaseOptions{}, "eip-delete", "Release an EIP", func(cli *azure.SRegion, args *EipReleaseOptions) error {
		return cli.DeallocateEIP(args.ID)
	})

	type EipShowOptions struct {
		ID string `help:"EIP ID"`
	}
	shellutils.R(&EipShowOptions{}, "eip-show", "Show an EIP", func(cli *azure.SRegion, args *EipShowOptions) error {
		if eip, err := cli.GetEip(args.ID); err != nil {
			return err
		} else {
			printObject(eip)
			return nil
		}
	})

	type EipAssociateOptions struct {
		ID       string `help:"EIP allocation ID"`
		INSTANCE string `help:"Instance ID"`
	}
	shellutils.R(&EipAssociateOptions{}, "eip-associate", "Associate an EIP", func(cli *azure.SRegion, args *EipAssociateOptions) error {
		err := cli.AssociateEip(args.ID, args.INSTANCE)
		return err
	})

	type EipDissociateOptions struct {
		ID string `help:"EIP allocation ID"`
	}

	shellutils.R(&EipDissociateOptions{}, "eip-dissociate", "Dissociate an EIP", func(cli *azure.SRegion, args *EipDissociateOptions) error {
		err := cli.DissociateEip(args.ID)
		return err
	})
}

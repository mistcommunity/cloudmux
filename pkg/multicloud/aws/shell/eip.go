package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type EipListOptions struct {
		Id          string
		Addr        string
		AssociateId string
	}
	shellutils.R(&EipListOptions{}, "eip-list", "List eips", func(cli *aws.SRegion, args *EipListOptions) error {
		eips, err := cli.GetEips(args.Id, args.Addr, args.AssociateId)
		if err != nil {
			return err
		}
		printList(eips, 0, 0, 0, []string{})
		return nil
	})

	type EipAllocateOptions struct {
		Name string
	}
	shellutils.R(&EipAllocateOptions{}, "eip-create", "Allocate an EIP", func(cli *aws.SRegion, args *EipAllocateOptions) error {
		opts := cloudprovider.SEip{Name: args.Name}
		eip, err := cli.AllocateEIP(&opts)
		if err != nil {
			return err
		}
		printObject(eip)
		return nil
	})

	type EipReleaseOptions struct {
		ID string `help:"EIP allocation ID"`
	}
	shellutils.R(&EipReleaseOptions{}, "eip-delete", "Release an EIP", func(cli *aws.SRegion, args *EipReleaseOptions) error {
		err := cli.DeallocateEIP(args.ID)
		return err
	})

	type EipAssociateOptions struct {
		ID       string `help:"EIP allocation ID"`
		INSTANCE string `help:"Instance ID"`
	}
	shellutils.R(&EipAssociateOptions{}, "eip-associate", "Associate an EIP", func(cli *aws.SRegion, args *EipAssociateOptions) error {
		err := cli.AssociateEip(args.ID, args.INSTANCE)
		return err
	})

	type EipDissociateOptions struct {
		INSTANCE string `help:"Instance ID"`
	}

	shellutils.R(&EipDissociateOptions{}, "eip-dissociate", "Dissociate an EIP", func(cli *aws.SRegion, args *EipDissociateOptions) error {
		err := cli.DissociateEip(args.INSTANCE)
		return err
	})
}

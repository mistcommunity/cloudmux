
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type EipListOptions struct {
		Id          string `help:"Eip id"`
		AssociateId string `help:"Id of associate resource"`
		Addr        string `help:"Eip "`
		Offset      int    `help:"List offset"`
		Limit       int    `help:"List limit"`
	}
	shellutils.R(&EipListOptions{}, "eip-list", "List eips", func(cli *aliyun.SRegion, args *EipListOptions) error {
		eips, total, e := cli.GetEips(args.Id, args.AssociateId, args.Addr, args.Offset, args.Limit)
		if e != nil {
			return e
		}
		printList(eips, total, args.Offset, args.Limit, []string{})
		return nil
	})

	shellutils.R(&cloudprovider.SEip{}, "eip-create", "Allocate an EIP", func(cli *aliyun.SRegion, args *cloudprovider.SEip) error {
		eip, err := cli.AllocateEIP(args)
		if err != nil {
			return err
		}
		printObject(eip)
		return nil
	})

	type EipReleaseOptions struct {
		ID string `help:"EIP allocation ID"`
	}
	shellutils.R(&EipReleaseOptions{}, "eip-delete", "Release an EIP", func(cli *aliyun.SRegion, args *EipReleaseOptions) error {
		err := cli.DeallocateEIP(args.ID)
		return err
	})

	type EipAssociateOptions struct {
		ID       string `help:"EIP allocation ID"`
		INSTANCE string `help:"Instance ID"`
	}
	shellutils.R(&EipAssociateOptions{}, "eip-associate", "Associate an EIP", func(cli *aliyun.SRegion, args *EipAssociateOptions) error {
		err := cli.AssociateEip(args.ID, args.INSTANCE)
		return err
	})
	shellutils.R(&EipAssociateOptions{}, "eip-dissociate", "Dissociate an EIP", func(cli *aliyun.SRegion, args *EipAssociateOptions) error {
		err := cli.DissociateEip(args.ID, args.INSTANCE)
		return err
	})
}

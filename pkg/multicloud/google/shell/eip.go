
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type EipListOptions struct {
		Address    string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&EipListOptions{}, "eip-list", "List eips", func(cli *google.SRegion, args *EipListOptions) error {
		eips, err := cli.GetEips(args.Address, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(eips, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&EipListOptions{}, "global-eip-list", "List global eips", func(cli *google.SRegion, args *EipListOptions) error {
		eips, err := cli.GetClient().GetGlobalRegion().GetEips(args.Address)
		if err != nil {
			return err
		}
		printList(eips, 0, 0, 0, nil)
		return nil
	})

	type EipIdOptions struct {
		ID string
	}
	shellutils.R(&EipIdOptions{}, "eip-show", "Show eip", func(cli *google.SRegion, args *EipIdOptions) error {
		eip, err := cli.GetEip(args.ID)
		if err != nil {
			return err
		}
		printObject(eip)
		return nil
	})

	shellutils.R(&EipIdOptions{}, "eip-delete", "Delete eip", func(cli *google.SRegion, args *EipIdOptions) error {
		return cli.Delete(args.ID)
	})

	type EipCreateOptions struct {
		NAME string
		Desc string
	}

	shellutils.R(&EipCreateOptions{}, "eip-create", "Create eip", func(cli *google.SRegion, args *EipCreateOptions) error {
		eip, err := cli.CreateEip(args.NAME, args.Desc)
		if err != nil {
			return err
		}
		printObject(eip)
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type VpcListOptions struct {
	}
	shellutils.R(&VpcListOptions{}, "vpc-list", "list vpc", func(cli *nutanix.SRegion, args *VpcListOptions) error {
		vpcs, err := cli.GetVpcs()
		if err != nil {
			return err
		}
		printList(vpcs, 0, 0, 0, []string{})
		return nil
	})

	type VpcIdOptions struct {
		ID string
	}

	shellutils.R(&VpcIdOptions{}, "vpc-show", "show vpc", func(cli *nutanix.SRegion, args *VpcIdOptions) error {
		vpc, err := cli.GetVpc(args.ID)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

	shellutils.R(&VpcIdOptions{}, "vpc-delete", "delete vpc", func(cli *nutanix.SRegion, args *VpcIdOptions) error {
		return cli.DeleteVpc(args.ID)
	})

	type VpcCreateOptions struct {
		Name string
		Desc string
		CIDR string
	}

	shellutils.R(&VpcCreateOptions{}, "vpc-create", "Create vpc", func(cli *nutanix.SRegion, args *VpcCreateOptions) error {
		opts := cloudprovider.VpcCreateOptions{
			NAME: args.Name,
			CIDR: args.CIDR,
			Desc: args.Desc,
		}
		vpc, err := cli.CreateVpc(&opts)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

}

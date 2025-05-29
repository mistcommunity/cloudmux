
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type VpcListOptions struct {
		Limit  int `help:"page size"`
		Offset int `help:"page offset"`
	}
	shellutils.R(&VpcListOptions{}, "vpc-list", "List vpcs", func(cli *aliyun.SRegion, args *VpcListOptions) error {
		vpcs, total, e := cli.GetVpcs(nil, args.Offset, args.Limit)
		if e != nil {
			return e
		}
		printList(vpcs, total, args.Offset, args.Limit, []string{})
		return nil
	})

	type VpcCreateOptions struct {
		Name string
		Desc string
		CIDR string
	}

	shellutils.R(&VpcCreateOptions{}, "vpc-create", "Create vpc", func(cli *aliyun.SRegion, args *VpcCreateOptions) error {
		opts := cloudprovider.VpcCreateOptions{
			NAME: args.Name,
			CIDR: args.CIDR,
			Desc: args.Desc,
		}
		vpc, err := cli.CreateIVpc(&opts)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

	type VpcOptions struct {
		ID string `help:"VPC id"`
	}

	shellutils.R(&VpcOptions{}, "vpc-delete", "Delete vpc", func(cli *aliyun.SRegion, args *VpcOptions) error {
		return cli.DeleteVpc(args.ID)
	})

	type VpcMoveResourceGroup struct {
		ResourceType    string `choices:"vpc|eip|bandwidthpackage" default:"vpc"`
		ResourceGroupId string
		ResourceId      string
	}

	shellutils.R(&VpcMoveResourceGroup{}, "vpc-mv-resource-group", "Delete vpc", func(cli *aliyun.SRegion, args *VpcMoveResourceGroup) error {
		return cli.VpcMoveResourceGroup(args.ResourceType, args.ResourceGroupId, args.ResourceId)
	})

}

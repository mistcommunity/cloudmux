package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type InstanceTypeListOptions struct {
	}
	shellutils.R(&InstanceTypeListOptions{}, "instance-type-list", "List intance types", func(cli *aws.SRegion, args *InstanceTypeListOptions) error {
		skus, err := cli.GetInstanceTypes()
		if err != nil {
			return err
		}
		printList(skus, 0, 0, 0, []string{})
		return nil
	})

	type SkuListOptions struct {
		Arch      string
		NextToken string
	}
	shellutils.R(&SkuListOptions{}, "sku-list", "List intance types", func(cli *aws.SRegion, args *SkuListOptions) error {
		skus, _, err := cli.DescribeInstanceTypes(args.Arch, args.NextToken)
		if err != nil {
			return err
		}
		printList(skus, 0, 0, 0, []string{})
		return nil
	})

}

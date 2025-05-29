
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ResourceGroupListOptions struct {
		PageSize   int
		PageNumber int
	}
	shellutils.R(&ResourceGroupListOptions{}, "resource-group-list", "List resource group", func(cli *aliyun.SRegion, args *ResourceGroupListOptions) error {
		groups, _, err := cli.GetClient().GetResourceGroups(args.PageNumber, args.PageSize)
		if err != nil {
			return err
		}
		printList(groups, 0, 0, 0, nil)
		return nil
	})

	type ResourceGroupShowOptions struct {
		ID string
	}

	shellutils.R(&ResourceGroupShowOptions{}, "resource-group-show", "Show resource group", func(cli *aliyun.SRegion, args *ResourceGroupShowOptions) error {
		group, err := cli.GetClient().GetResourceGroup(args.ID)
		if err != nil {
			return err
		}
		printObject(group)
		return nil
	})

	type ResourceGroupCreateOptions struct {
		NAME string
	}

	shellutils.R(&ResourceGroupCreateOptions{}, "resource-group-create", "Create resource group", func(cli *aliyun.SRegion, args *ResourceGroupCreateOptions) error {
		group, err := cli.GetClient().CreateResourceGroup(args.NAME)
		if err != nil {
			return err
		}
		printObject(group)
		return nil
	})

	type ResourceListOptions struct {
		Service      string
		ResourceType string
	}

	shellutils.R(&ResourceListOptions{}, "resource-list", "List resources", func(cli *aliyun.SRegion, args *ResourceListOptions) error {
		ret, err := cli.GetClient().ListResources(args.Service, args.ResourceType)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, nil)
		return nil
	})

}

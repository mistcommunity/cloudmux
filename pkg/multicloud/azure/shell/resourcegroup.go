
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type ResourceGroupListOptions struct {
		Limit  int `help:"page size"`
		Offset int `help:"page offset"`
	}
	shellutils.R(&ResourceGroupListOptions{}, "resource-group-list", "List group", func(cli *azure.SRegion, args *ResourceGroupListOptions) error {
		groups, err := cli.GetClient().ListResourceGroups()
		if err != nil {
			return err
		}
		printList(groups, len(groups), 0, 0, []string{})
		return nil
	})

	type ResourceGroupOptions struct {
		GROUP string `help:"ResourceGrop Name"`
	}

	shellutils.R(&ResourceGroupOptions{}, "resource-group-show", "Show group detail", func(cli *azure.SRegion, args *ResourceGroupOptions) error {
		if group, err := cli.GetResourceGroupDetail(args.GROUP); err != nil {
			return err
		} else {
			printObject(group)
			return nil
		}
	})

	shellutils.R(&ResourceGroupOptions{}, "resource-group-create", "Create resource group", func(cli *azure.SRegion, args *ResourceGroupOptions) error {
		resp, err := cli.CreateResourceGroup(args.GROUP)
		if err != nil {
			return err
		}
		printObject(resp)
		return nil
	})

	shellutils.R(&ResourceGroupOptions{}, "resource-group-delete", "Delete resource group", func(cli *azure.SRegion, args *ResourceGroupOptions) error {
		err := cli.DeleteResourceGroup(args.GROUP)
		if err != nil {
			return err
		}
		return nil
	})

	type ResourceGroupUpdateOptions struct {
		GROUP string `help:"Name of resource group to update"`
		NAME  string `help:"New name of resource group"`
	}
	shellutils.R(&ResourceGroupUpdateOptions{}, "resource-group-update", "Update resource group detail", func(cli *azure.SRegion, args *ResourceGroupUpdateOptions) error {
		err := cli.UpdateResourceGroup(args.GROUP, args.NAME)
		if err != nil {
			return err
		}
		return nil
	})

}

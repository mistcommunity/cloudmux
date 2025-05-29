
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type RegionListOptions struct {
	}
	shellutils.R(&RegionListOptions{}, "region-list", "List regions", func(cli *azure.SRegion, args *RegionListOptions) error {
		regions := cli.GetClient().GetRegions()
		printList(regions, 0, 0, 0, nil)
		return nil
	})

	type ResourceIdOptions struct {
		ID string `help:"resource id"`
	}
	shellutils.R(&ResourceIdOptions{}, "delete", "delete resource", func(cli *azure.SRegion, args *ResourceIdOptions) error {
		return cli.Delete(args.ID)
	})

	shellutils.R(&ResourceIdOptions{}, "show", "Show resource", func(cli *azure.SRegion, args *ResourceIdOptions) error {
		ret, err := cli.Show(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

}

package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type RegionListOptions struct {
	}
	shellutils.R(&RegionListOptions{}, "region-list", "List regions", func(cli *aws.SRegion, args *RegionListOptions) error {
		regions, err := cli.GetClient().GetRegions()
		if err != nil {
			return err
		}
		printList(regions, 0, 0, 0, nil)
		return nil
	})
}

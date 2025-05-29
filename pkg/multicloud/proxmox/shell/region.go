
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type RegionListOptions struct {
	}
	shellutils.R(&RegionListOptions{}, "region-list", "list regions", func(cli *proxmox.SRegion, args *RegionListOptions) error {
		regions, err := cli.GetClient().GetRegions()
		if err != nil {
			return err
		}
		printList(regions, 0, 0, 0, []string{})
		return nil
	})

}

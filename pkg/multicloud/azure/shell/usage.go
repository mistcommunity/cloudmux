
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type UsageListOptions struct {
		TYPE string `choices:"Microsoft.Network|Microsoft.Compute|Microsoft.Storage"`
	}
	shellutils.R(&UsageListOptions{}, "usage-list", "List usage", func(cli *azure.SRegion, args *UsageListOptions) error {
		usage, err := cli.GetUsage(args.TYPE)
		if err != nil {
			return err
		}
		printList(usage, 0, 0, 0, []string{})
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type VRouterListOptions struct {
		Limit  int `help:"page size"`
		Offset int `help:"page offset"`
	}
	shellutils.R(&VRouterListOptions{}, "vrouter-list", "List vrouters", func(cli *aliyun.SRegion, args *VRouterListOptions) error {
		vrouters, total, e := cli.GetVRouters(args.Offset, args.Limit)
		if e != nil {
			return e
		}
		printList(vrouters, total, args.Offset, args.Limit, []string{})
		return nil
	})
}

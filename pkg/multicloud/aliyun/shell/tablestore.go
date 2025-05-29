
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type TableStoreListOptions struct {
		PageSize   int `help:"page size" default:"10"`
		PageNumber int `help:"page number" default:"1"`
	}
	shellutils.R(&TableStoreListOptions{}, "tablestore-list", "List tablestores", func(cli *aliyun.SRegion, args *TableStoreListOptions) error {
		ret, total, err := cli.GetTablestoreInstances(args.PageSize, args.PageNumber)
		if err != nil {
			return err
		}
		printList(ret, total, 0, 0, []string{})
		return nil
	})
}

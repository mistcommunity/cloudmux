
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type HbaseListOptions struct {
	}
	shellutils.R(&HbaseListOptions{}, "hbase-instance-list", "List Hbase", func(cli *aliyun.SRegion, args *HbaseListOptions) error {
		ret, err := cli.GetHbaseInstances()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

}

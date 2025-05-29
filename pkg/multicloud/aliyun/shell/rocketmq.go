
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type RocketmqListOptions struct {
	}
	shellutils.R(&RocketmqListOptions{}, "rocketmq-list", "List rocketmq", func(cli *aliyun.SRegion, args *RocketmqListOptions) error {
		ret, err := cli.GetRocketmqInstances()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})
}

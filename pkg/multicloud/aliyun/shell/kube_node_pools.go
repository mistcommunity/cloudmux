
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type KubeNodePoolListOptions struct {
		CLUSTER string
	}
	shellutils.R(&KubeNodePoolListOptions{}, "kube-node-pool-list", "List kube node pools", func(cli *aliyun.SRegion, args *KubeNodePoolListOptions) error {
		pools, err := cli.GetKubeNodePools(args.CLUSTER)
		if err != nil {
			return err
		}
		printList(pools, 0, 0, 0, []string{})
		return nil
	})

}

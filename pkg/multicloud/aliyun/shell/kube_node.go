
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type KubeNodeListOptions struct {
		CLUSTER    string
		PageSize   int
		PageNumber int
	}
	shellutils.R(&KubeNodeListOptions{}, "kube-node-list", "List kube nodes", func(cli *aliyun.SRegion, args *KubeNodeListOptions) error {
		nodes, _, err := cli.GetKubeNodes(args.CLUSTER, args.PageSize, args.PageNumber)
		if err != nil {
			return err
		}
		printList(nodes, 0, 0, 0, []string{})
		return nil
	})

}

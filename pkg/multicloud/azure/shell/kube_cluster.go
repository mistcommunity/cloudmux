
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type KubeClusterListOptions struct {
	}
	shellutils.R(&KubeClusterListOptions{}, "kube-cluster-list", "List kube cluster", func(cli *azure.SRegion, args *KubeClusterListOptions) error {
		clusters, err := cli.GetKubeClusters()
		if err != nil {
			return err
		}
		printList(clusters, len(clusters), 0, 0, []string{})
		return nil
	})

	type KubeClusterIdOptions struct {
		ID string
	}

	shellutils.R(&KubeClusterIdOptions{}, "kube-cluster-kubeconfig", "Show kubeconfg", func(cli *azure.SRegion, args *KubeClusterIdOptions) error {
		config, err := cli.GetKubeConfig(args.ID)
		if err != nil {
			return err
		}
		printObject(config)
		return nil
	})

}

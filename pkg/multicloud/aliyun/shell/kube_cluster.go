
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type KubeClusterListOptions struct {
		PageSize   int
		PageNumber int
	}
	shellutils.R(&KubeClusterListOptions{}, "kube-cluster-list", "List kube clusters", func(cli *aliyun.SRegion, args *KubeClusterListOptions) error {
		clusters, _, err := cli.GetKubeClusters(args.PageSize, args.PageNumber)
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

	type KubeClusterIdOptions struct {
		ID string
	}

	shellutils.R(&KubeClusterIdOptions{}, "kube-cluster-show", "Show kube cluster", func(cli *aliyun.SRegion, args *KubeClusterIdOptions) error {
		cluster, err := cli.GetKubeCluster(args.ID)
		if err != nil {
			return err
		}
		printObject(cluster)
		return nil
	})

	shellutils.R(&KubeClusterIdOptions{}, "kube-cluster-delete", "Delete kube cluster", func(cli *aliyun.SRegion, args *KubeClusterIdOptions) error {
		return cli.DeleteKubeCluster(args.ID, false)
	})

	shellutils.R(&cloudprovider.KubeClusterCreateOptions{}, "kube-cluster-create", "Create kube cluster", func(cli *aliyun.SRegion, args *cloudprovider.KubeClusterCreateOptions) error {
		cluster, err := cli.CreateKubeCluster(args)
		if err != nil {
			return err
		}
		printObject(cluster)
		return nil
	})

	type KubeClusterKubeconfigOptions struct {
		ID            string
		Private       bool
		ExpireMinutes int
	}

	shellutils.R(&KubeClusterKubeconfigOptions{}, "kube-cluster-kubeconfig", "Get kube cluster kubeconfig", func(cli *aliyun.SRegion, args *KubeClusterKubeconfigOptions) error {
		config, err := cli.GetKubeConfig(args.ID, args.Private, args.ExpireMinutes)
		if err != nil {
			return err
		}
		printObject(config)
		return nil
	})

}

package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type KubeClusterListOptions struct {
		NextToken string
	}
	shellutils.R(&KubeClusterListOptions{}, "kube-cluster-list", "List kube cluster", func(cli *aws.SRegion, args *KubeClusterListOptions) error {
		ret, _, err := cli.GetKubeClusters(args.NextToken)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type KubeClusterNameOptions struct {
		NAME string
	}

	shellutils.R(&KubeClusterNameOptions{}, "kube-cluster-show", "Show kube cluster", func(cli *aws.SRegion, args *KubeClusterNameOptions) error {
		ret, err := cli.GetKubeCluster(args.NAME)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&KubeClusterNameOptions{}, "kube-cluster-delete", "Delete kube cluster", func(cli *aws.SRegion, args *KubeClusterNameOptions) error {
		return cli.DeleteKubeCluster(args.NAME)
	})

	shellutils.R(&cloudprovider.KubeClusterCreateOptions{}, "kube-cluster-create", "Create kube cluster", func(cli *aws.SRegion, args *cloudprovider.KubeClusterCreateOptions) error {
		cluster, err := cli.CreateKubeCluster(args)
		if err != nil {
			return err
		}
		printObject(cluster)
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type ClusterListOptions struct {
	}
	shellutils.R(&ClusterListOptions{}, "cluster-list", "list clusters", func(cli *nutanix.SRegion, args *ClusterListOptions) error {
		clusters, err := cli.GetClusters()
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

	type ClusterIdOptions struct {
		ID string
	}

	shellutils.R(&ClusterIdOptions{}, "cluster-show", "show clusters", func(cli *nutanix.SRegion, args *ClusterIdOptions) error {
		cluster, err := cli.GetCluster(args.ID)
		if err != nil {
			return err
		}
		printObject(cluster)
		return nil
	})

}

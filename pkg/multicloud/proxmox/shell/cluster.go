
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type ClusterListOptions struct {
	}
	shellutils.R(&ClusterListOptions{}, "cluster-list", "list clusters", func(cli *proxmox.SRegion, args *ClusterListOptions) error {
		clusters, err := cli.GetClusterAllResources()
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&ClusterListOptions{}, "cluster-node-list", "list nodes", func(cli *proxmox.SRegion, args *ClusterListOptions) error {
		clusters, err := cli.GetClusterNodeResources()
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

}

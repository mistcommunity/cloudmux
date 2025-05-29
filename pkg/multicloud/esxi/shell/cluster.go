
package shell

import (
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/esxi"
)

func init() {
	type ClusterListOptions struct {
		DATACENTER string `help:"List clusters in datacenter"`
	}
	shellutils.R(&ClusterListOptions{}, "cluster-list", "List all clusters", func(cli *esxi.SESXiClient, args *ClusterListOptions) error {
		dc, err := cli.FindDatacenterByMoId(args.DATACENTER)
		if err != nil {
			return err
		}
		clusters, err := dc.ListClusters()
		if err != nil {
			return err
		}
		printList(clusters, nil)
		return nil
	})

	type ClusterPoolListOptions struct {
		DATACENTER string `help:"List clusters in datacenter"`
		CLUSTER    string `help:"List cluster resource pool"`
	}

	shellutils.R(&ClusterPoolListOptions{}, "cluster-pool-list", "List all cluster resource pool", func(cli *esxi.SESXiClient, args *ClusterPoolListOptions) error {
		dc, err := cli.FindDatacenterByMoId(args.DATACENTER)
		if err != nil {
			return err
		}
		cluster, err := dc.GetCluster(args.CLUSTER)
		if err != nil {
			return err
		}
		pools, err := cluster.ListResourcePools()
		if err != nil {
			return errors.Wrap(err, "ListResourcePools")
		}
		printList(pools, nil)
		return nil
	})

	type ClusterPoolSyncOptions struct {
		DATACENTER string `help:"List clusters in datacenter"`
		CLUSTER    string `help:"List cluster resource pool"`
		Name       string `help:"Resource pool name"`
	}

	shellutils.R(&ClusterPoolSyncOptions{}, "cluster-pool-sync", "Sync cluster resource pool", func(cli *esxi.SESXiClient, args *ClusterPoolSyncOptions) error {
		dc, err := cli.FindDatacenterByMoId(args.DATACENTER)
		if err != nil {
			return err
		}
		cluster, err := dc.GetCluster(args.CLUSTER)
		if err != nil {
			return err
		}
		pool, err := cluster.SyncResourcePool(args.Name)
		if err != nil {
			return errors.Wrap(err, "SyncResourcePool")
		}
		printObject(pool)
		return nil
	})

}

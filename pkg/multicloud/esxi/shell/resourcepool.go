
package shell

import (
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/esxi"
)

func init() {
	type ResourcePoolListOptions struct {
		DATACENTER string `help:"List resource pool in datacenter"`
	}
	shellutils.R(&ResourcePoolListOptions{}, "resource-pool-list", "List all resource pools", func(cli *esxi.SESXiClient, args *ResourcePoolListOptions) error {
		dc, err := cli.FindDatacenterByMoId(args.DATACENTER)
		if err != nil {
			return err
		}
		pools, err := dc.GetResourcePools()
		if err != nil {
			return err
		}
		printList(pools, nil)
		return nil
	})

	type ResourcePoolCreateOptions struct {
		DATACENTER string `help:"Create resource pool in datacenter"`
		CLUSTER    string `help:"Cluster name"`
		NAME       string `help:"Resource pool name"`
	}

	shellutils.R(&ResourcePoolCreateOptions{}, "resource-pool-create", "Create resource pool", func(cli *esxi.SESXiClient, args *ResourcePoolCreateOptions) error {
		dc, err := cli.FindDatacenterByMoId(args.DATACENTER)
		if err != nil {
			return err
		}
		cluster, err := dc.GetCluster(args.CLUSTER)
		if err != nil {
			return errors.Wrap(err, "GetCluster")
		}
		pool, err := cluster.CreateResourcePool(args.NAME)
		if err != nil {
			return err
		}
		printObject(pool)
		return nil
	})
}

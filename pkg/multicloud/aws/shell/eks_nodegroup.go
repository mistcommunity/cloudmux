package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type NodegroupListOptions struct {
		CLUSTER_NAME string
		NextToken    string
	}
	shellutils.R(&NodegroupListOptions{}, "node-group-list", "List node group", func(cli *aws.SRegion, args *NodegroupListOptions) error {
		ret, _, err := cli.GetNodegroups(args.CLUSTER_NAME, args.NextToken)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type NodegroupNameOptions struct {
		CLUSTER_NAME string
		NAME         string
	}

	shellutils.R(&NodegroupNameOptions{}, "node-group-show", "Show node group", func(cli *aws.SRegion, args *NodegroupNameOptions) error {
		ret, err := cli.GetNodegroup(args.CLUSTER_NAME, args.NAME)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&NodegroupNameOptions{}, "node-group-delete", "Delete node group", func(cli *aws.SRegion, args *NodegroupNameOptions) error {
		return cli.DeleteNodegroup(args.CLUSTER_NAME, args.NAME)
	})

}

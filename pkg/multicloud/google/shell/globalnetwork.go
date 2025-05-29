
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type GlobalNetworkListOptions struct {
		MaxResults int
		PageToken  string
	}
	shellutils.R(&GlobalNetworkListOptions{}, "global-network-list", "List globalnetworks", func(cli *google.SRegion, args *GlobalNetworkListOptions) error {
		globalnetworks, err := cli.GetClient().GetGlobalNetworks(args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(globalnetworks, 0, 0, 0, nil)
		return nil
	})

	type GlobalNetworkShowOptions struct {
		ID string
	}
	shellutils.R(&GlobalNetworkShowOptions{}, "global-network-show", "Show globalnetwork", func(cli *google.SRegion, args *GlobalNetworkShowOptions) error {
		globalnetwork, err := cli.GetClient().GetGlobalNetwork(args.ID)
		if err != nil {
			return err
		}
		printObject(globalnetwork)
		return nil
	})

	type GlobalNetworkCreateOptions struct {
		NAME string
		Desc string
	}

	shellutils.R(&GlobalNetworkCreateOptions{}, "global-network-create", "Create globalnetwork", func(cli *google.SRegion, args *GlobalNetworkCreateOptions) error {
		globalnetwork, err := cli.GetClient().CreateGlobalNetwork(args.NAME, args.Desc)
		if err != nil {
			return err
		}
		printObject(globalnetwork)
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type FirewallListOptions struct {
		Network    string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&FirewallListOptions{}, "firewall-list", "List firewalls", func(cli *google.SRegion, args *FirewallListOptions) error {
		firewalls, err := cli.GetClient().GetFirewalls(args.Network, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(firewalls, 0, 0, 0, nil)
		return nil
	})

	type FirewallShowOptions struct {
		ID string
	}
	shellutils.R(&FirewallShowOptions{}, "firewall-show", "Show firewall", func(cli *google.SRegion, args *FirewallShowOptions) error {
		firewall, err := cli.GetClient().GetFirewall(args.ID)
		if err != nil {
			return err
		}
		printObject(firewall)
		return nil
	})

}

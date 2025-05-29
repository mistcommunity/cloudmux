package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type TrifficPolicyInstanceShowOptions struct {
		ID string
	}
	shellutils.R(&TrifficPolicyInstanceShowOptions{}, "dns-traffic-policy-instance-show", "Show traffic policy instance", func(cli *aws.SRegion, args *TrifficPolicyInstanceShowOptions) error {
		ret, err := cli.GetClient().GetDnsTrafficPolicyInstance(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	type TrifficPolicyShowOptions struct {
		ID      string
		VERSION string
	}
	shellutils.R(&TrifficPolicyShowOptions{}, "dns-traffic-policy-show", "Show traffic policy", func(cli *aws.SRegion, args *TrifficPolicyShowOptions) error {
		ret, err := cli.GetClient().GetTrafficPolicy(args.ID, args.VERSION)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&TrifficPolicyInstanceShowOptions{}, "dns-extra-address-list", "List traffic policy address", func(cli *aws.SRegion, args *TrifficPolicyInstanceShowOptions) error {
		ret, err := cli.GetClient().GetDnsExtraAddresses(args.ID)
		if err != nil {
			return err
		}
		fmt.Println(ret)
		return nil
	})

}

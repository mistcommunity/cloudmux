
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type RuleGroupListOptions struct {
	}
	shellutils.R(&RuleGroupListOptions{}, "waf-rule-group-list", "List waf rule groups", func(cli *azure.SRegion, args *RuleGroupListOptions) error {
		groups, err := cli.ListAppWafManagedRuleGroup()
		if err != nil {
			return err
		}
		printList(groups, len(groups), 0, 0, []string{})
		return nil
	})

	type FrontDoorPolicyListOptions struct {
		RESOURCE_GROUP string
	}
	shellutils.R(&FrontDoorPolicyListOptions{}, "front-door-policy-list", "List front door policies", func(cli *azure.SRegion, args *FrontDoorPolicyListOptions) error {
		policies, err := cli.ListFrontDoorWafs(args.RESOURCE_GROUP)
		if err != nil {
			return err
		}
		printList(policies, 0, 0, 0, []string{})
		return nil
	})

	type AppGatewayWafListOptions struct {
	}

	shellutils.R(&AppGatewayWafListOptions{}, "app-gateway-waf-list", "List app gateway wafs", func(cli *azure.SRegion, args *AppGatewayWafListOptions) error {
		wafs, err := cli.ListAppWafs()
		if err != nil {
			return err
		}
		printList(wafs, 0, 0, 0, []string{})
		return nil
	})

	type AppGatewayWafRuleGroupListOptions struct {
	}

	shellutils.R(&AppGatewayWafRuleGroupListOptions{}, "app-gateway-waf-rule-group-list", "List app gateway wafs", func(cli *azure.SRegion, args *AppGatewayWafRuleGroupListOptions) error {
		group, err := cli.ListAppWafManagedRuleGroup()
		if err != nil {
			return err
		}
		printList(group, 0, 0, 0, []string{})
		return nil
	})

}

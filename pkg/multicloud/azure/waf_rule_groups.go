
package azure

import "net/url"

type SWafRule struct {
	RuleId        string
	Description   string
	DefaultAction string
	DefaultState  string
}

type SRuleGroup struct {
	ruleGroupName string
	description   string
	Rules         []SWafRule
}

type SManagedRuleGroupProperties struct {
	ProvisioningState string
	RuleSetId         string
	RuleSetType       string
	RuleSetVersion    string
	RuleGroups        []SRuleGroup
}

type SManagedRuleGroup struct {
	Name       string
	Id         string
	Type       string
	Properties SManagedRuleGroupProperties
}

func (self *SRegion) ListManagedRuleGroups() ([]SManagedRuleGroup, error) {
	groups := []SManagedRuleGroup{}
	err := self.list("Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets", url.Values{}, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

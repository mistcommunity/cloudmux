
package azure

import "net/url"

type SFrontDoorProperties struct {
	ResourceState     string
	ProvisioningState string
	PolicySettings    struct {
		EnabledState                  string
		Mode                          string
		RedirectUrl                   string
		CustomBlockResponseStatusCode int
		CustomBlockResponseBody       string
		RequestBodyCheck              string
	}
	CustomRules struct {
		Rules []struct{}
	}
	ManagedRules struct {
		ManagedRuleSets []struct {
			RuleSetType        string
			RuleSetVersion     string
			RuleSetAction      string
			RuleGroupOverrides []struct {
			}
			Exclusions []struct{}
		}
	}
	FrontendEndpointLinks []struct{}
	RoutingRuleLinks      []struct{}
	SecurityPolicyLinks   []struct{}
}

type SFrontDoorWaf struct {
	Id       string
	Name     string
	Type     string
	Tags     map[string]string
	Location string
	Sku      struct {
		Name string
	}
	Properties SFrontDoorProperties
}

func (self *SRegion) ListFrontDoorWafs(resGroup string) ([]SFrontDoorWaf, error) {
	params := url.Values{}
	params.Set("resourceGroups", resGroup)
	ret := []SFrontDoorWaf{}
	err := self.list("Microsoft.Network/frontdoorWebApplicationFirewallPolicies", params, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

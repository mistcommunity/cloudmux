package cloudprovider

import (
	"fmt"
	"strings"

	"yunion.io/x/pkg/util/secrules"
)

type SecurityGroupReference struct {
	Id   string
	Name string
}

type SecurityGroupCreateInput struct {
	Name      string
	Desc      string
	VpcId     string
	ProjectId string

	Tags map[string]string
}

type SecurityGroupRuleCreateOptions struct {
	Desc      string
	Priority  int
	Protocol  string
	Ports     string
	Direction secrules.TSecurityRuleDirection
	CIDR      string
	Action    secrules.TSecurityRuleAction
}

type SecurityGroupRuleUpdateOptions struct {
	CIDR     string
	Action   secrules.TSecurityRuleAction
	Desc     string
	Ports    string
	Protocol string
	Priority int
}

func (self *SecurityGroupRuleCreateOptions) String() string {
	ret := fmt.Sprintf("%s_%s_%s", self.Direction, self.Action, self.Protocol)
	if len(self.CIDR) > 0 {
		ret += "_" + self.CIDR
	}
	if len(self.Ports) > 0 {
		ret += "_" + self.Ports
	}
	ret = strings.ReplaceAll(ret, ".", "_")
	ret = strings.ReplaceAll(ret, ",", "_")
	return ret
}

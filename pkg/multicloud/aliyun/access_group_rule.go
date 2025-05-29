
package aliyun

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SAccessGroupRule struct {
	region *SRegion

	AccessGroupName string
	RWAccess        string
	UserAccess      string
	Priority        int
	SourceCidrIp    string
	AccessRuleId    string
}

func (self *SAccessGroupRule) GetGlobalId() string {
	return self.AccessRuleId
}

func (self *SAccessGroupRule) Delete() error {
	return self.region.DeleteAccessGroupRule(self.AccessGroupName, self.AccessRuleId)
}

func (self *SAccessGroupRule) GetPriority() int {
	return self.Priority
}

func (self *SAccessGroupRule) GetSource() string {
	return self.SourceCidrIp
}

func (self *SAccessGroupRule) GetRWAccessType() cloudprovider.TRWAccessType {
	switch self.RWAccess {
	case "RDWR":
		return cloudprovider.RWAccessTypeRW
	case "RDONLY":
		return cloudprovider.RWAccessTypeR
	default:
		return cloudprovider.TRWAccessType(self.RWAccess)
	}
}

func (self *SAccessGroupRule) GetUserAccessType() cloudprovider.TUserAccessType {
	switch self.UserAccess {
	case "no_squash":
		return cloudprovider.UserAccessTypeNoRootSquash
	case "root_squash":
		return cloudprovider.UserAccessTypeRootSquash
	case "all_squash":
		return cloudprovider.UserAccessTypeAllSquash
	default:
		return cloudprovider.TUserAccessType(self.UserAccess)
	}
}

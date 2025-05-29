
package aliyun

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SElasticcacheAcl struct {
	multicloud.SElasticcacheAclBase
	AliyunTags

	cacheDB *SElasticcache

	SecurityIPList           string `json:"SecurityIpList"`
	SecurityIPGroupAttribute string `json:"SecurityIpGroupAttribute"`
	SecurityIPGroupName      string `json:"SecurityIpGroupName"`
}

func (self *SElasticcacheAcl) GetId() string {
	return fmt.Sprintf("%s/%s", self.cacheDB.GetId(), self.SecurityIPGroupName)
}

func (self *SElasticcacheAcl) GetName() string {
	return self.SecurityIPGroupName
}

func (self *SElasticcacheAcl) GetGlobalId() string {
	return self.GetId()
}

func (self *SElasticcacheAcl) GetStatus() string {
	return api.ELASTIC_CACHE_ACL_STATUS_AVAILABLE
}

func (self *SElasticcacheAcl) Refresh() error {
	iacl, err := self.cacheDB.GetICloudElasticcacheAcl(self.GetId())
	if err != nil {
		return err
	}

	err = jsonutils.Update(self, iacl.(*SElasticcacheAcl))
	if err != nil {
		return err
	}

	return nil
}

func (self *SElasticcacheAcl) GetIpList() string {
	return self.SecurityIPList
}

// https://help.aliyun.com/document_detail/61002.html?spm=a2c4g.11186623.6.764.3752782fJpbjxH
func (self *SElasticcacheAcl) Delete() error {
	params := make(map[string]string)
	params["InstanceId"] = self.cacheDB.GetId()
	params["ModifyMode"] = "Delete"
	params["SecurityIpGroupName"] = self.GetName()
	params["SecurityIps"] = self.SecurityIPList

	err := DoAction(self.cacheDB.region.kvsRequest, "ModifySecurityIps", params, nil, nil)
	if err != nil {
		return errors.Wrap(err, "elasticcacheAcl.Delete")
	}

	return nil
}

// https://help.aliyun.com/document_detail/61002.html?spm=a2c4g.11186623.6.764.3752782fJpbjxH
func (self *SElasticcacheAcl) UpdateAcl(securityIps string) error {
	return self.cacheDB.region.createAcl(self.cacheDB.GetId(), self.GetName(), securityIps)
}

func (self *SRegion) createAcl(instanceId, aclName, securityIps string) error {
	params := make(map[string]string)
	params["InstanceId"] = instanceId
	params["SecurityIpGroupName"] = aclName
	params["ModifyMode"] = "Cover"
	params["SecurityIps"] = securityIps

	err := DoAction(self.kvsRequest, "ModifySecurityIps", params, nil, nil)
	if err != nil {
		return errors.Wrap(err, "region.UpdateAcl")
	}

	return nil
}

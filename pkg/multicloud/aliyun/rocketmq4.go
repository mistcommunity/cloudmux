
package aliyun

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
)

const (
	ROCKYMQ4_API_VERSION = "2019-02-14"
)

type SRocketmq4Instance struct {
	InstanceId   string
	InstanceName string
}

func (self *SRegion) onsRequest(apiName string, params map[string]string, body interface{}) (jsonutils.JSONObject, error) {
	client, err := self.getSdkClient()
	if err != nil {
		return nil, err
	}
	params = self.client.SetResourceGropuId(params)
	return doRequest(client, fmt.Sprintf("ons.%s.aliyuncs.com", self.RegionId), ROCKYMQ4_API_VERSION, apiName, params, body, self.client.debug)
}

func (region *SRegion) GetRocketmq4Instances() ([]SRocketmq4Instance, error) {
	params := map[string]string{}
	resp, err := region.onsRequest("OnsInstanceInServiceList", params, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "OnsInstanceInServiceList")
	}
	ret := []SRocketmq4Instance{}
	return ret, resp.Unmarshal(&ret, "Data", "InstanceVO")
}

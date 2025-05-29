
package aliyun

import (
	"fmt"

	"yunion.io/x/jsonutils"
)

type SApigateway struct {
	InstanceId   string
	InstanceName string
}

func (self *SRegion) apiRequest(apiName string, params map[string]string) (jsonutils.JSONObject, error) {
	client, err := self.getSdkClient()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("apigateway.%s.aliyuncs.com", self.RegionId)
	params = self.client.SetResourceGropuId(params)
	return jsonRequest(client, endpoint, "2016-07-14", apiName, params, self.client.debug)
}

func (self *SRegion) GetApigateways() ([]SApigateway, error) {
	params := map[string]string{}
	resp, err := self.apiRequest("DescribeInstances", params)
	if err != nil {
		return nil, err
	}
	part := struct {
		Instances struct {
			InstanceAttribute []SApigateway
		}
		TotalCount int
	}{}
	err = resp.Unmarshal(&part)
	if err != nil {
		return nil, err
	}
	return part.Instances.InstanceAttribute, nil
}

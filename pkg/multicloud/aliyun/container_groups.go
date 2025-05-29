
package aliyun

import (
	"fmt"

	"yunion.io/x/jsonutils"
)

type SContainerGroup struct {
	ContainerGroupId   string
	ContainerGroupName string
}

func (region *SRegion) eciRequest(apiName string, params map[string]string) (jsonutils.JSONObject, error) {
	client, err := region.getSdkClient()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("eci.%s.aliyuncs.com", region.RegionId)
	params = region.client.SetResourceGropuId(params)
	return jsonRequest(client, endpoint, "2018-08-08", apiName, params, region.client.debug)
}

func (self *SRegion) GetContainerGroups() ([]SContainerGroup, error) {
	ret := []SContainerGroup{}
	params := map[string]string{
		"RegionId": self.RegionId,
	}
	for {
		resp, err := self.eciRequest("DescribeContainerGroups", params)
		if err != nil {
			return nil, err
		}
		part := struct {
			NextToken       string
			ContainerGroups []SContainerGroup
		}{}
		err = resp.Unmarshal(&part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.ContainerGroups...)
		if len(part.NextToken) == 0 || len(part.ContainerGroups) == 0 {
			break
		}
		params["NextToken"] = part.NextToken
	}
	return ret, nil
}

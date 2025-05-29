
package aliyun

import "fmt"

type SAliyunResource struct {
	ResourceGroupId string
	ResourceId      string
	Service         string
	ResourceType    string
	RegionId        string
	CreateDate      string
}

func (self *SAliyunClient) ListResources(service, resourceType string) ([]SAliyunResource, error) {
	params := map[string]string{
		"PageSize": "100",
	}
	if len(service) > 0 {
		params["Service"] = service
	}
	if len(resourceType) > 0 {
		params["ResourceType"] = resourceType
	}
	pageNumber := 1
	ret := []SAliyunResource{}
	for {
		params["PageNumber"] = fmt.Sprintf("%d", pageNumber)
		resp, err := self.rmRequest("ListResources", params)
		if err != nil {
			return nil, err
		}
		part := struct {
			Resources struct {
				Resource []SAliyunResource
			}
			TotalCount int
		}{}
		err = resp.Unmarshal(&part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.Resources.Resource...)
		if len(ret) >= part.TotalCount || len(part.Resources.Resource) == 0 {
			break
		}
		pageNumber++
	}
	return ret, nil
}

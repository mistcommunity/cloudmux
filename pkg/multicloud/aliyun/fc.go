
package aliyun

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
)

const (
	FC_API_VERSION = "2021-04-06"
)

type SFcService struct {
	ServiceId   string
	ServiceName string
}

func (self *SRegion) fcRequest(apiName string, params map[string]string, body interface{}) (jsonutils.JSONObject, error) {
	client, err := self.getSdkClient()
	if err != nil {
		return nil, err
	}
	params = self.client.SetResourceGropuId(params)
	return doRequest(client, fmt.Sprintf("fc.%s.aliyuncs.com", self.RegionId), FC_API_VERSION, apiName, params, body, self.client.debug)
}

func (self *SRegion) GetFcServices() ([]SFcService, error) {
	params := map[string]string{
		"PathPattern": "/services",
		"limit":       "100",
	}
	ret := []SFcService{}
	for {
		resp, err := self.fc3Request("ListServices", params, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "ListServices")
		}
		part := struct {
			Services  []SFcService
			NextToken string
		}{}
		err = resp.Unmarshal(&part)
		if err != nil {
			return nil, errors.Wrapf(err, "resp.Unmarshal")
		}
		ret = append(ret, part.Services...)
		if len(part.NextToken) == 0 || len(part.Services) == 0 {
			break
		}
		params["nextToken"] = part.NextToken
	}
	return ret, nil
}

type SFcFunction struct {
	FunctionId   string
	FunctionName string
}

func (self *SRegion) GetFcFunctions(service string) ([]SFcFunction, error) {
	params := map[string]string{
		"PathPattern": fmt.Sprintf("/services/%s/functions", service),
		"limit":       "100",
	}
	ret := []SFcFunction{}
	for {
		resp, err := self.fc3Request("ListFunctions", params, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "ListFunctions")
		}
		part := struct {
			Functions []SFcFunction
			NextToken string
		}{}
		err = resp.Unmarshal(&part)
		if err != nil {
			return nil, errors.Wrapf(err, "resp.Unmarshal")
		}
		ret = append(ret, part.Functions...)
		if len(part.NextToken) == 0 || len(part.Functions) == 0 {
			break
		}
		params["nextToken"] = part.NextToken
	}
	return ret, nil
}

type SFcInstance struct {
	InstanceId string
	VersionId  string
}

func (self *SRegion) GetFcInstances(service, funcName string) ([]SFcInstance, error) {
	params := map[string]string{
		"PathPattern": fmt.Sprintf("/services/%s/functions/%s/instances", service, funcName),
		"limit":       "100",
	}
	resp, err := self.fc3Request("ListFunctions", params, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "ListFunctions")
	}
	ret := []SFcInstance{}
	return ret, resp.Unmarshal(&ret, "instances")
}

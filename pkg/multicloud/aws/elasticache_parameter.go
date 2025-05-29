
package aws

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SElasticacheParameter struct {
	multicloud.SElasticcacheParameterBase
	AwsTags

	AllowedValues        string `xml:"AllowedValues"`
	ChangeType           string `xml:"ChangeType"`
	DataType             string `xml:"DataType"`
	Description          string `xml:"Description"`
	IsModifiable         bool   `xml:"IsModifiable"`
	MinimumEngineVersion string `xml:"MinimumEngineVersion"`
	ParameterName        string `xml:"ParameterName"`
	ParameterValue       string `xml:"ParameterValue"`
	Source               string `xml:"Source"`
}

func (region *SRegion) GetCacheParameters(id string) ([]SElasticacheParameter, error) {
	params := map[string]string{}
	if len(id) > 0 {
		params["CacheParameterGroupName"] = id
	}
	ret := []SElasticacheParameter{}
	for {
		part := struct {
			CacheNodeTypeSpecificParameters []SElasticacheParameter `xml:"CacheNodeTypeSpecificParameters>CacheNodeTypeSpecificParameter"`
			Marker                          string
		}{}
		err := region.ecRequest("DescribeCacheParameters", params, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.CacheNodeTypeSpecificParameters...)
		if len(part.CacheNodeTypeSpecificParameters) == 0 || len(part.Marker) == 0 {
			break
		}
		params["Marker"] = part.Marker
	}
	return ret, nil
}

func (self *SElasticacheParameter) GetId() string {
	return self.ParameterName
}

func (self *SElasticacheParameter) GetName() string {
	return self.ParameterName
}

func (self *SElasticacheParameter) GetGlobalId() string {
	return self.GetId()
}

func (self *SElasticacheParameter) GetStatus() string {
	return api.ELASTIC_CACHE_PARAMETER_STATUS_AVAILABLE
}

func (self *SElasticacheParameter) GetParameterKey() string {
	return self.ParameterName
}

func (self *SElasticacheParameter) GetParameterValue() string {
	return self.ParameterValue
}

func (self *SElasticacheParameter) GetParameterValueRange() string {
	return self.AllowedValues
}

func (self *SElasticacheParameter) GetDescription() string {
	return self.Description
}

func (self *SElasticacheParameter) GetModifiable() bool {
	return self.IsModifiable
}

func (self *SElasticacheParameter) GetForceRestart() bool {
	return self.ChangeType == "requires-reboot"
}

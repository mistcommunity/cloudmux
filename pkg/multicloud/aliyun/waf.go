
package aliyun

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SInstanceSpecs struct {
	Code  int
	Value bool
}

type SWafInstance struct {
	Version           string
	InstanceSpecInfos []SInstanceSpecs
	InstanceId        string
	ExpireTime        uint64
}

func (self *SRegion) DescribeInstanceSpecInfo() (*SWafInstance, error) {
	params := map[string]string{
		"RegionId": self.RegionId,
	}
	resp, err := self.wafRequest("DescribeInstanceSpecInfo", params)
	if err != nil {
		return nil, errors.Wrapf(err, "DescribeInstanceSpecInfo")
	}
	ret := &SWafInstance{}
	err = resp.Unmarshal(&ret)
	if err != nil {
		return nil, errors.Wrapf(err, "resp.Unmarshal")
	}
	if len(ret.InstanceId) == 0 {
		return nil, cloudprovider.ErrNotFound
	}
	return ret, nil
}

func (self *SRegion) DeleteInstance(id string) error {
	params := map[string]string{
		"RegionId":   self.RegionId,
		"InstanceId": id,
	}
	_, err := self.wafRequest("DeleteInstance", params)
	return errors.Wrapf(err, "DeleteInstance")
}

func (self *SRegion) GetICloudWafInstancesV1() ([]cloudprovider.ICloudWafInstance, error) {
	ins, err := self.DescribeInstanceSpecInfo()
	if err != nil {
		if errors.Cause(err) == cloudprovider.ErrNotFound {
			return []cloudprovider.ICloudWafInstance{}, nil
		}
		return nil, errors.Wrapf(err, "DescribeInstanceSpecInfo")
	}
	domains, err := self.DescribeDomainNames(ins.InstanceId)
	if err != nil {
		return nil, errors.Wrapf(err, "DescribeDomainNames")
	}
	ret := []cloudprovider.ICloudWafInstance{}
	for i := range domains {
		domain, err := self.DescribeDomain(ins.InstanceId, domains[i])
		if err != nil {
			return nil, errors.Wrapf(err, "DescribeDomain %s", domains[i])
		}
		domain.region = self
		domain.insId = ins.InstanceId
		domain.name = domains[i]
		ret = append(ret, domain)
	}
	return ret, nil
}

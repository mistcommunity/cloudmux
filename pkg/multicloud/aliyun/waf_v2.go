
package aliyun

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/errors"
)

func (self *SRegion) DescribeWafInstance() (*SWafInstance, error) {
	params := map[string]string{
		"RegionId": self.RegionId,
	}
	resp, err := self.wafv2Request("DescribeInstance", params)
	if err != nil {
		return nil, errors.Wrapf(err, "DescribeInstance")
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

func (self *SRegion) GetICloudWafInstancesV2() ([]cloudprovider.ICloudWafInstance, error) {
	ins, err := self.DescribeWafInstance()
	if err != nil {
		if errors.Cause(err) == cloudprovider.ErrNotFound {
			return []cloudprovider.ICloudWafInstance{}, nil
		}
		return nil, errors.Wrapf(err, "DescribeInstanceSpecInfo")
	}
	domains, err := self.DescribeWafDomains(ins.InstanceId)
	if err != nil {
		return nil, errors.Wrapf(err, "DescribeDomainNames")
	}
	ret := []cloudprovider.ICloudWafInstance{}
	for i := range domains {
		domain, err := self.DescribeDomainV2(ins.InstanceId, domains[i].Domain)
		if err != nil {
			return nil, errors.Wrapf(err, "DescribeDomain %s", domains[i].Domain)
		}
		domain.region = self
		domain.insId = ins.InstanceId
		ret = append(ret, domain)
	}
	return ret, nil
}

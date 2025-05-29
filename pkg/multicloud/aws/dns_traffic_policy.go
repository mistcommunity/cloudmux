package aws

import (
	"fmt"
)

type SDnsTrafficPolicyInstance struct {
	Id                   string `xml:"Id"`
	Name                 string `xml:"Name"`
	TrafficPolicyId      string `xml:"TrafficPolicyId"`
	TrafficPolicyType    string `xml:"TrafficPolicyType"`
	TrafficPolicyVersion string `xml:"TrafficPolicyVersion"`
}

func (self *SAwsClient) GetDnsTrafficPolicyInstance(id string) (*SDnsTrafficPolicyInstance, error) {
	params := map[string]string{"Id": fmt.Sprintf("trafficpolicyinstance/%s", id)}
	ret := &struct {
		TrafficPolicyInstance SDnsTrafficPolicyInstance `xml:"TrafficPolicyInstance"`
	}{}
	err := self.dnsRequest("GetTrafficPolicyInstance", params, ret)
	if err != nil {
		return nil, err
	}
	return &ret.TrafficPolicyInstance, nil
}

type SDnsTrafficPolicy struct {
	Comment  string `xml:"Comment"`
	Document string `xml:"Document"`
	Name     string `xml:"Name"`
	Type     string `xml:"Type"`
}

func (self *SAwsClient) GetTrafficPolicy(id string, version string) (*SDnsTrafficPolicy, error) {
	params := map[string]string{"Id": fmt.Sprintf("trafficpolicy/%s/%s", id, version)}
	ret := &struct {
		TrafficPolicy SDnsTrafficPolicy `xml:"TrafficPolicy"`
	}{}
	err := self.dnsRequest("GetTrafficPolicy", params, ret)
	if err != nil {
		return nil, err
	}
	return &ret.TrafficPolicy, nil
}

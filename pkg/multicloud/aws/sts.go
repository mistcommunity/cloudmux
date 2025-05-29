
package aws

import "yunion.io/x/pkg/errors"

type SCallerIdentity struct {
	Arn     string `xml:"Arn"`
	UserId  string `xml:"UserId"`
	Account string `xml:"Account"`
}

func (self *SAwsClient) GetCallerIdentity() (*SCallerIdentity, error) {
	ret := &SCallerIdentity{}
	err := self.stsRequest("GetCallerIdentity", nil, ret)
	if err != nil {
		return nil, errors.Wrap(err, "GetCallerIdentity")
	}
	return ret, nil
}

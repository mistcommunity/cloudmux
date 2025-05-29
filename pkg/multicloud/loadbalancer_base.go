
package multicloud

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/errors"
)

type SLoadbalancerBase struct {
	SVirtualResourceBase
}

func (lb *SLoadbalancerBase) GetIEIP() (cloudprovider.ICloudEIP, error) {
	return nil, nil
}

func (lb *SLoadbalancerBase) GetSecurityGroupIds() ([]string, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetSecurityGroupIds")
}

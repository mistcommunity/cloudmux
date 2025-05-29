
package multicloud

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SNatGatewayBase struct {
	SResourceBase
	SBillingBase
}

func (nat *SNatGatewayBase) GetIEips() ([]cloudprovider.ICloudEIP, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIEips")
}

func (nat *SNatGatewayBase) GetIDNatEntries() ([]cloudprovider.ICloudNatDEntry, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIDNatEntries")
}

func (nat *SNatGatewayBase) GetISNatEntries() ([]cloudprovider.ICloudNatSEntry, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISNatEntries")
}

func (nat *SNatGatewayBase) Delete() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Delete")
}

func (nat *SNatGatewayBase) GetIpAddr() string {
	return ""
}

func (nat *SNatGatewayBase) GetBandwidthMb() int {
	return 0
}

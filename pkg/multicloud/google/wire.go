
package google

import (
	"fmt"
	"time"

	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SWire struct {
	multicloud.SResourceBase
	GoogleTags
	vpc *SVpc
}

func (wire *SWire) GetId() string {
	return wire.vpc.GetGlobalId()
}

func (wire *SWire) GetGlobalId() string {
	return fmt.Sprintf("%s-%s", wire.GetId(), wire.vpc.region.Name)
}

func (wire *SWire) GetName() string {
	return wire.vpc.GetName()
}

func (wire *SWire) GetCreatedAt() time.Time {
	return time.Time{}
}

func (wire *SWire) CreateINetwork(opts *cloudprovider.SNetworkCreateOptions) (cloudprovider.ICloudNetwork, error) {
	return nil, cloudprovider.ErrNotSupported
}

func (wire *SWire) GetIVpc() cloudprovider.ICloudVpc {
	return wire.vpc
}

func (wire *SWire) GetIZone() cloudprovider.ICloudZone {
	return nil
}

func (self *SWire) GetINetworks() ([]cloudprovider.ICloudNetwork, error) {
	network := SNetwork{wire: self}
	return []cloudprovider.ICloudNetwork{&network}, nil
}

func (self *SWire) GetINetworkById(id string) (cloudprovider.ICloudNetwork, error) {
	networks, err := self.GetINetworks()
	if err != nil {
		return nil, errors.Wrapf(err, "GetINetwork")
	}
	for i := range networks {
		if networks[i].GetGlobalId() == id {
			return networks[i], nil
		}
	}
	return nil, errors.Wrapf(cloudprovider.ErrNotFound, id)
}

func (wire *SWire) GetBandwidth() int {
	return 0
}

func (wire *SWire) GetStatus() string {
	return api.WIRE_STATUS_AVAILABLE
}

func (wire *SWire) IsEmulated() bool {
	return true
}

func (wire *SWire) Refresh() error {
	return nil
}


package proxmox

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SWire struct {
	multicloud.SResourceBase
	ProxmoxTags

	region *SRegion
}

func (self *SWire) GetId() string {
	return self.region.GetId()
}

func (self *SWire) GetName() string {
	return "Default"
}

func (self *SWire) GetGlobalId() string {
	return self.GetId()
}

func (self *SWire) CreateINetwork(opts *cloudprovider.SNetworkCreateOptions) (cloudprovider.ICloudNetwork, error) {
	return nil, cloudprovider.ErrNotSupported
}

func (self *SWire) GetBandwidth() int {
	return 10000
}

func (self *SWire) GetINetworks() ([]cloudprovider.ICloudNetwork, error) {
	nets, err := self.region.GetNetworks()
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.ICloudNetwork{}
	for i := range nets {
		nets[i].wire = self
		ret = append(ret, &nets[i])
	}
	return ret, nil
}

func (self *SWire) GetINetworkById(id string) (cloudprovider.ICloudNetwork, error) {
	net, err := self.region.GetNetwork(id)
	if err != nil {
		return nil, err
	}
	net.wire = self
	return net, nil
}

func (self *SWire) GetIVpc() cloudprovider.ICloudVpc {
	return self.region.getVpc()
}

func (self *SWire) GetIZone() cloudprovider.ICloudZone {
	zone, _ := self.region.GetZone()
	return zone
}

func (self *SWire) GetStatus() string {
	return api.WIRE_STATUS_AVAILABLE
}

func (self *SRegion) GetWires() ([]SWire, error) {
	ret := []SWire{}
	wire := &SWire{region: self}
	ret = append(ret, *wire)
	return ret, nil
}

func (self *SRegion) GetWire() (*SWire, error) {
	ret := &SWire{region: self}
	return ret, nil
}

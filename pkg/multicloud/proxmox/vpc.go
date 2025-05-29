
package proxmox

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SVpc struct {
	multicloud.SVpc
	ProxmoxTags

	region *SRegion
}

func (self *SVpc) GetName() string {
	return "Default"
}

func (self *SVpc) GetId() string {
	return self.region.GetId()
}

func (self *SVpc) GetGlobalId() string {
	return self.GetId()
}

func (self *SVpc) Delete() error {
	return cloudprovider.ErrNotSupported
}

func (self *SVpc) IsEmulated() bool {
	return true
}

func (self *SVpc) GetCidrBlock() string {
	return "0.0.0.0/0"
}

func (self *SVpc) GetIRouteTables() ([]cloudprovider.ICloudRouteTable, error) {
	return []cloudprovider.ICloudRouteTable{}, nil
}

func (self *SVpc) GetIRouteTableById(routeTableId string) (cloudprovider.ICloudRouteTable, error) {
	return nil, cloudprovider.ErrNotFound
}

func (self *SVpc) GetISecurityGroups() ([]cloudprovider.ICloudSecurityGroup, error) {
	return []cloudprovider.ICloudSecurityGroup{}, nil
}

func (self *SVpc) GetIWires() ([]cloudprovider.ICloudWire, error) {
	wires, err := self.region.GetWires()
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.ICloudWire{}
	for i := range wires {
		wires[i].region = self.region
		ret = append(ret, &wires[i])
	}
	return ret, nil
}

func (self *SVpc) GetIWireById(id string) (cloudprovider.ICloudWire, error) {
	wire, err := self.region.GetWire()
	if err != nil {
		return nil, err
	}
	wire.region = self.region
	return wire, nil
}

func (self *SVpc) GetIsDefault() bool {
	return true
}

func (self *SVpc) GetRegion() cloudprovider.ICloudRegion {
	return self.region
}

func (self *SVpc) GetStatus() string {
	return api.VPC_STATUS_AVAILABLE
}


package nutanix

import (
	"fmt"
	"net/url"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SInstanceNic struct {
	cloudprovider.DummyICloudNic
	ins *SInstance

	MacAddress  string   `json:"mac_address"`
	NetworkUUID string   `json:"network_uuid"`
	NicUUID     string   `json:"nic_uuid"`
	Model       string   `json:"model"`
	IPAddress   string   `json:"ip_address"`
	IPAddresses []string `json:"ip_addresses"`
	VlanMode    string   `json:"vlan_mode"`
	IsConnected bool     `json:"is_connected"`
}

func (self *SInstanceNic) GetId() string {
	return self.NicUUID
}

func (self *SInstanceNic) GetIP() string {
	return self.IPAddress
}

func (self *SInstanceNic) GetMAC() string {
	return self.MacAddress
}

func (self *SInstanceNic) GetDriver() string {
	return "virtio"
}

func (self *SInstanceNic) InClassicNetwork() bool {
	return false
}

func (self *SInstanceNic) GetSubAddress() ([]string, error) {
	ret := []string{}
	for _, addr := range self.IPAddresses {
		if addr != self.IPAddress {
			ret = append(ret, addr)
		}
	}
	return ret, nil
}

func (self *SInstanceNic) GetINetworkId() string {
	if len(self.IPAddress) == 0 {
		return self.NetworkUUID
	}
	vpc, err := self.ins.host.zone.region.GetVpc(self.NetworkUUID)
	if err != nil {
		return self.NetworkUUID
	}
	wires, err := vpc.GetIWires()
	if err != nil {
		return self.NetworkUUID
	}
	for i := range wires {
		networks, err := wires[i].GetINetworks()
		if err != nil {
			continue
		}
		for j := range networks {
			network := networks[j].(*SNetwork)
			if network.Contains(self.IPAddress) {
				return network.GetGlobalId()
			}
		}
	}
	return self.NetworkUUID
}

func (self *SInstanceNic) AssignAddress(ipAddrs []string) error {
	return cloudprovider.ErrNotImplemented
}

func (self *SRegion) GetInstanceNics(id string) ([]SInstanceNic, error) {
	nics := []SInstanceNic{}
	res := fmt.Sprintf("vms/%s/nics", id)
	params := url.Values{}
	params.Set("include_address_assignments", "true")
	_, err := self.list(res, params, &nics)
	return nics, err
}

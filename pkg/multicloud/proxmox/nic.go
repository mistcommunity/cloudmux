
package proxmox

import (
	"fmt"
	"strings"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SInstanceNic struct {
	cloudprovider.DummyICloudNic
	ins *SInstance

	MacAddr   string
	IpAddr    string
	NicId     string
	Model     string
	NetworkId string
}

func (self *SInstanceNic) GetId() string {
	return fmt.Sprintf("%d/%s", self.ins.VmID, self.NicId)
}

func (self *SInstanceNic) GetIP() string {
	return self.IpAddr
}

func (self *SInstanceNic) GetMAC() string {
	return self.MacAddr
}

func (self *SInstanceNic) GetDriver() string {
	return strings.ToLower(self.Model)
}

func (self *SInstanceNic) InClassicNetwork() bool {
	return true
}

func (self *SInstanceNic) GetSubAddress() ([]string, error) {
	return []string{}, nil
}

func (self *SInstanceNic) GetINetworkId() string {
	return self.NetworkId
}

func (self *SInstanceNic) AssignAddress(ipAddrs []string) error {
	return cloudprovider.ErrNotSupported
}

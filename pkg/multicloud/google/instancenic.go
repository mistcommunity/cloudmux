
package google

import (
	"yunion.io/x/pkg/util/netutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SNetworkInterface struct {
	instance *SInstance

	Network       string
	Subnetwork    string
	NetworkIP     string
	Name          string
	AccessConfigs []AccessConfig
	Fingerprint   string
	Kind          string

	cloudprovider.DummyICloudNic
}

func (nic *SNetworkInterface) GetId() string {
	return ""
}

func (nic *SNetworkInterface) GetIP() string {
	return nic.NetworkIP
}

func (nic *SNetworkInterface) GetMAC() string {
	ip, _ := netutils.NewIPV4Addr(nic.NetworkIP)
	return ip.ToMac("42:01:")
}

func (nic *SNetworkInterface) GetDriver() string {
	return "virtio"
}

func (nic *SNetworkInterface) InClassicNetwork() bool {
	return false
}

func (nic *SNetworkInterface) GetINetworkId() string {
	vpc := &SVpc{region: nic.instance.host.zone.region}
	err := nic.instance.host.zone.region.GetBySelfId(nic.Subnetwork, vpc)
	if err != nil {
		return ""
	}
	return vpc.Id
}

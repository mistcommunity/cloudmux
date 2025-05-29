
package esxi

import (
	"github.com/vmware/govmomi/vim25/types"

	"yunion.io/x/log"
	"yunion.io/x/pkg/util/netutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SVirtualNIC struct {
	SVirtualDevice

	cloudprovider.DummyICloudNic
}

func NewVirtualNIC(vm *SVirtualMachine, dev types.BaseVirtualDevice, index int) SVirtualNIC {
	return SVirtualNIC{
		SVirtualDevice: NewVirtualDevice(vm, dev, index),
	}
}

func (nic *SVirtualNIC) getVirtualEthernetCard() *types.VirtualEthernetCard {
	card := types.VirtualEthernetCard{}
	if FetchAnonymousFieldValue(nic.dev, &card) {
		return &card
	}
	return nil
}

func (nic *SVirtualNIC) GetId() string {
	return ""
}

func (nic *SVirtualNIC) GetIP() string {
	guestIps := nic.vm.getGuestIps()
	if nicConf, ok := guestIps[nic.GetMAC()]; ok {
		if len(nicConf.IPs) > 0 {
			return nicConf.IPs[0]
		}
	}
	log.Warningf("cannot find ip for mac %s", nic.GetMAC())
	return ""
}

func (nic *SVirtualNIC) GetDriver() string {
	return nic.SVirtualDevice.GetDriver()
}

func (nic *SVirtualNIC) GetMAC() string {
	return netutils.FormatMacAddr(nic.getVirtualEthernetCard().MacAddress)
}

func (nic *SVirtualNIC) InClassicNetwork() bool {
	return false
}

func (nic *SVirtualNIC) GetINetworkId() string {
	return ""
}

func (nic *SVirtualNIC) GetSubAddress() ([]string, error) {
	guestIps := nic.vm.getGuestIps()
	if nicConf, ok := guestIps[nic.GetMAC()]; ok {
		if len(nicConf.IPs) > 1 {
			return nicConf.IPs[1:], nil
		}
	}
	return nil, nil
}

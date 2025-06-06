
package azure

import (
	"strings"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/util/netutils"
	"yunion.io/x/pkg/util/rbacscope"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SubnetPropertiesFormat struct {
	AddressPrefix   string   `json:"addressPrefix,omitempty"`
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

func (self *SubnetPropertiesFormat) getPrefix() string {
	if len(self.AddressPrefix) > 0 {
		return self.AddressPrefix
	}
	for _, prefix := range self.AddressPrefixes {
		return prefix
	}
	return ""
}

type SNetwork struct {
	multicloud.SNetworkBase
	AzureTags
	wire *SWire

	//AvailableIpAddressCount int `json:"availableIpAddressCount,omitempty"`
	ID            string
	Name          string
	Properties    SubnetPropertiesFormat
	AddressPrefix string `json:"addressPrefix,omitempty"`
}

func (self *SNetwork) GetTags() (map[string]string, error) {
	return self.Tags, nil
}

func (self *SNetwork) GetId() string {
	return self.ID
}

func (self *SNetwork) GetName() string {
	return self.Name
}

func (self *SNetwork) GetGlobalId() string {
	return strings.ToLower(self.ID)
}

func (self *SNetwork) GetStatus() string {
	return api.NETWORK_STATUS_AVAILABLE
}

func (self *SNetwork) Delete() error {
	return self.wire.vpc.region.del(self.ID)
}

func (self *SNetwork) getPrefix() string {
	if len(self.AddressPrefix) > 0 {
		return self.AddressPrefix
	}
	return self.Properties.getPrefix()
}

func (self *SNetwork) GetGateway() string {
	pref, _ := netutils.NewIPV4Prefix(self.getPrefix())
	endIp := pref.Address.BroadcastAddr(pref.MaskLen) // 255
	endIp = endIp.StepDown()                          // 254
	return endIp.String()
}

func (self *SNetwork) GetIWire() cloudprovider.ICloudWire {
	return self.wire
}

func (self *SNetwork) GetIpEnd() string {
	pref, _ := netutils.NewIPV4Prefix(self.getPrefix())
	endIp := pref.Address.BroadcastAddr(pref.MaskLen) // 255
	endIp = endIp.StepDown()                          // 254
	return endIp.String()
}

func (self *SNetwork) GetIpMask() int8 {
	pref, _ := netutils.NewIPV4Prefix(self.getPrefix())
	return pref.MaskLen
}

// https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-faq
func (self *SNetwork) GetIpStart() string {
	pref, _ := netutils.NewIPV4Prefix(self.getPrefix())
	startIp := pref.Address.NetAddr(pref.MaskLen) // 0
	startIp = startIp.StepUp()                    // 1
	startIp = startIp.StepUp()                    // 2
	startIp = startIp.StepUp()                    // 3
	startIp = startIp.StepUp()                    // 4
	return startIp.String()
}

func (self *SNetwork) GetIsPublic() bool {
	return true
}

func (self *SNetwork) GetPublicScope() rbacscope.TRbacScope {
	return rbacscope.ScopeDomain
}

func (self *SNetwork) GetServerType() string {
	return api.NETWORK_TYPE_GUEST
}

func (self *SNetwork) Refresh() error {
	network, err := self.wire.zone.region.GetNetwork(self.ID)
	if err != nil {
		return err
	}
	return jsonutils.Update(self, network)
}

func (self *SNetwork) GetAllocTimeoutSeconds() int {
	return 120 // 2 minutes
}

func (self *SNetwork) GetProjectId() string {
	return getResourceGroup(self.ID)
}

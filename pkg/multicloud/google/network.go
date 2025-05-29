
package google

import (
	"time"

	"yunion.io/x/pkg/util/netutils"
	"yunion.io/x/pkg/util/rbacscope"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SNetwork struct {
	multicloud.SNetworkBase
	GoogleTags
	wire *SWire
}

func (network *SNetwork) GetProjectId() string {
	return network.wire.vpc.region.GetProjectId()
}

func (self *SNetwork) GetName() string {
	return self.wire.vpc.GetName()
}

func (self *SNetwork) GetId() string {
	return self.wire.vpc.GetId()
}

func (self *SNetwork) GetDescription() string {
	return ""
}

func (self *SNetwork) GetGlobalId() string {
	return self.wire.vpc.GetGlobalId()
}

func (self *SNetwork) Refresh() error {
	return self.wire.vpc.Refresh()
}

func (network *SNetwork) IsEmulated() bool {
	return false
}

func (network *SNetwork) GetStatus() string {
	return api.NETWORK_INTERFACE_STATUS_AVAILABLE
}

func (network *SNetwork) GetCreatedAt() time.Time {
	return time.Time{}
}

func (network *SNetwork) Delete() error {
	return network.wire.vpc.Delete()
}

func (network *SNetwork) GetAllocTimeoutSeconds() int {
	return 300
}

func (network *SNetwork) GetIWire() cloudprovider.ICloudWire {
	return network.wire
}

func (self *SNetwork) GetIpStart() string {
	pref, _ := netutils.NewIPV4Prefix(self.wire.vpc.IpCidrRange)
	startIp := pref.Address.NetAddr(pref.MaskLen) // 0
	startIp = startIp.StepUp()                    // 1
	return startIp.String()
}

func (self *SNetwork) GetIpEnd() string {
	pref, _ := netutils.NewIPV4Prefix(self.wire.vpc.IpCidrRange)
	endIp := pref.Address.BroadcastAddr(pref.MaskLen) // 255
	endIp = endIp.StepDown()                          // 254
	return endIp.String()
}

func (self *SNetwork) GetIpMask() int8 {
	pref, _ := netutils.NewIPV4Prefix(self.wire.vpc.IpCidrRange)
	return pref.MaskLen
}

func (self *SNetwork) GetGateway() string {
	return self.wire.vpc.GatewayAddress
}

func (network *SNetwork) GetServerType() string {
	return api.NETWORK_TYPE_GUEST
}

func (network *SNetwork) GetIsPublic() bool {
	return true
}

func (network *SNetwork) GetPublicScope() rbacscope.TRbacScope {
	return rbacscope.ScopeDomain
}

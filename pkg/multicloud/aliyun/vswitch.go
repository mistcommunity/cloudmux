
package aliyun

import (
	"fmt"
	"strings"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/util/netutils"
	"yunion.io/x/pkg/util/rbacscope"
	"yunion.io/x/pkg/utils"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

// {"AvailableIpAddressCount":4091,"CidrBlock":"172.31.32.0/20","CreationTime":"2017-03-19T13:37:44Z","Description":"System created default virtual switch.","IsDefault":true,"Status":"Available","VSwitchId":"vsw-j6c3gig5ub4fmi2veyrus","VSwitchName":"","VpcId":"vpc-j6c86z3sh8ufhgsxwme0q","ZoneId":"cn-hongkong-b"}

const (
	VSwitchPending   = "Pending"
	VSwitchAvailable = "Available"
)

type SCloudResources struct {
	CloudResourceSetType []string
}

type SVSwitch struct {
	multicloud.SNetworkBase
	AliyunTags
	wire *SWire

	AvailableIpAddressCount int

	CidrBlock     string
	Ipv6CidrBlock string
	EnabledIpv6   bool
	CreationTime  time.Time
	Description   string
	IsDefault     bool
	Status        string
	VSwitchId     string
	VSwitchName   string
	VpcId         string
	ZoneId        string

	CloudResources  SCloudResources
	ResourceGroupId string
	RouteTable      SRouteTable
}

func (self *SVSwitch) GetId() string {
	return self.VSwitchId
}

func (self *SVSwitch) GetName() string {
	if len(self.VSwitchName) > 0 {
		return self.VSwitchName
	}
	return self.VSwitchId
}

func (self *SVSwitch) GetGlobalId() string {
	return self.VSwitchId
}

func (self *SVSwitch) GetStatus() string {
	return strings.ToLower(self.Status)
}

func (self *SVSwitch) Refresh() error {
	net, err := self.wire.zone.region.GetVSwitchAttributes(self.VSwitchId)
	if err != nil {
		return err
	}
	return jsonutils.Update(self, net)
}

func (self *SVSwitch) GetIWire() cloudprovider.ICloudWire {
	return self.wire
}

func (self *SVSwitch) SetTags(tags map[string]string, replace bool) error {
	return self.wire.zone.region.SetResourceTags(ALIYUN_SERVICE_VPC, "VSWITCH", self.VSwitchId, tags, replace)
}

func (net *SVSwitch) GetIp6Start() string {
	if len(net.Ipv6CidrBlock) > 0 {
		prefix, err := netutils.NewIPV6Prefix(net.Ipv6CidrBlock)
		if err != nil {
			return ""
		}
		return prefix.Address.NetAddr(prefix.MaskLen).StepUp().String()
	}
	return ""
}

func (net *SVSwitch) GetIp6End() string {
	if len(net.Ipv6CidrBlock) > 0 {
		prefix, err := netutils.NewIPV6Prefix(net.Ipv6CidrBlock)
		if err != nil {
			return ""
		}
		end := prefix.Address.NetAddr(prefix.MaskLen).BroadcastAddr(prefix.MaskLen)
		for i := 0; i < 15; i++ {
			end = end.StepDown()
		}
		return end.String()
	}
	return ""
}

func (net *SVSwitch) GetIp6Mask() uint8 {
	if len(net.Ipv6CidrBlock) > 0 {
		prefix, err := netutils.NewIPV6Prefix(net.Ipv6CidrBlock)
		if err != nil {
			return 0
		}
		return prefix.MaskLen
	}
	return 0
}

func (net *SVSwitch) GetGateway6() string {
	if len(net.Ipv6CidrBlock) > 0 {
		prefix, err := netutils.NewIPV6Prefix(net.Ipv6CidrBlock)
		if err != nil {
			return ""
		}
		return prefix.Address.NetAddr(prefix.MaskLen).StepUp().String()
	}
	return ""
}

func (self *SVSwitch) GetIpStart() string {
	pref, _ := netutils.NewIPV4Prefix(self.CidrBlock)
	startIp := pref.Address.NetAddr(pref.MaskLen) // 0
	startIp = startIp.StepUp()                    // 1
	return startIp.String()
}

func (self *SVSwitch) GetIpEnd() string {
	pref, _ := netutils.NewIPV4Prefix(self.CidrBlock)
	endIp := pref.Address.BroadcastAddr(pref.MaskLen) // 255
	endIp = endIp.StepDown()                          // 254
	endIp = endIp.StepDown()                          // 253
	endIp = endIp.StepDown()                          // 252
	return endIp.String()
}

func (self *SVSwitch) GetIpMask() int8 {
	pref, _ := netutils.NewIPV4Prefix(self.CidrBlock)
	return pref.MaskLen
}

func (self *SVSwitch) GetGateway() string {
	pref, _ := netutils.NewIPV4Prefix(self.CidrBlock)
	endIp := pref.Address.BroadcastAddr(pref.MaskLen) // 255
	endIp = endIp.StepDown()                          // 254
	return endIp.String()
}

func (self *SVSwitch) GetServerType() string {
	return api.NETWORK_TYPE_GUEST
}

func (self *SVSwitch) GetIsPublic() bool {
	// return self.IsDefault
	return true
}

func (self *SVSwitch) GetPublicScope() rbacscope.TRbacScope {
	return rbacscope.ScopeDomain
}

func (self *SRegion) createVSwitch(zoneId string, vpcId string, name string, cidr string, desc string) (string, error) {
	params := make(map[string]string)
	params["ZoneId"] = zoneId
	params["VpcId"] = vpcId
	params["CidrBlock"] = cidr
	params["VSwitchName"] = name
	if len(desc) > 0 {
		params["Description"] = desc
	}
	params["ClientToken"] = utils.GenRequestId(20)

	body, err := self.vpcRequest("CreateVSwitch", params)
	if err != nil {
		return "", err
	}
	return body.GetString("VSwitchId")
}

func (self *SRegion) DeleteVSwitch(vswitchId string) error {
	params := make(map[string]string)
	params["VSwitchId"] = vswitchId

	_, err := self.vpcRequest("DeleteVSwitch", params)
	return err
}

func (self *SVSwitch) Delete() error {
	err := self.Refresh()
	if err != nil {
		log.Errorf("refresh vswitch fail %s", err)
		return err
	}
	if len(self.RouteTable.RouteTableId) > 0 && !self.RouteTable.IsSystem() {
		err = self.wire.zone.region.UnassociateRouteTable(self.RouteTable.RouteTableId, self.VSwitchId)
		if err != nil {
			log.Errorf("unassociate routetable fail %s", err)
			return err
		}
	}
	err = self.dissociateWithSNAT()
	if err != nil {
		log.Errorf("fail to dissociateWithSNAT")
		return err
	}
	err = cloudprovider.Wait(10*time.Second, 60*time.Second, func() (bool, error) {
		err := self.wire.zone.region.DeleteVSwitch(self.VSwitchId)
		if err != nil {
			// delete network immediately after deleting vm on it
			// \"Code\":\"DependencyViolation\",\"Message\":\"Specified object has dependent resources.\"}
			if isError(err, "DependencyViolation") {
				return false, nil
			}
			return false, err
		} else {
			return true, nil
		}
	})
	return err
}

func (self *SVSwitch) GetAllocTimeoutSeconds() int {
	return 120 // 2 minutes
}

func (self *SRegion) GetVSwitches(ids []string, vpcId string, offset int, limit int) ([]SVSwitch, int, error) {
	if limit > 50 || limit <= 0 {
		limit = 50
	}
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["PageSize"] = fmt.Sprintf("%d", limit)
	params["PageNumber"] = fmt.Sprintf("%d", (offset/limit)+1)
	if ids != nil && len(ids) > 0 {
		params["VSwitchId"] = strings.Join(ids, ",")
	}
	if len(vpcId) > 0 {
		params["VpcId"] = vpcId
	}

	body, err := self.vpcRequest("DescribeVSwitches", params)
	if err != nil {
		log.Errorf("GetVSwitches fail %s", err)
		return nil, 0, err
	}

	switches := make([]SVSwitch, 0)
	err = body.Unmarshal(&switches, "VSwitches", "VSwitch")
	if err != nil {
		log.Errorf("Unmarshal vswitches fail %s", err)
		return nil, 0, err
	}
	total, _ := body.Int("TotalCount")
	return switches, int(total), nil
}

func (self *SRegion) GetVSwitchAttributes(idstr string) (*SVSwitch, error) {
	params := make(map[string]string)
	params["VSwitchId"] = idstr

	body, err := self.vpcRequest("DescribeVSwitchAttributes", params)
	if err != nil {
		log.Errorf("DescribeVSwitchAttributes fail %s", err)
		return nil, err
	}
	if self.client.debug {
		log.Debugf("%s", body.PrettyString())
	}
	switches := SVSwitch{}
	err = body.Unmarshal(&switches)
	if err != nil {
		log.Errorf("Unmarshal vswitches fail %s", err)
		return nil, err
	}
	return &switches, nil
}

func (vsw *SVSwitch) dissociateWithSNAT() error {
	natgatways, err := vsw.wire.vpc.getNatGateways()
	if err != nil {
		return err
	}
	for i := range natgatways {
		err = natgatways[i].dissociateWithVswitch(vsw.VSwitchId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *SVSwitch) GetProjectId() string {
	return self.ResourceGroupId
}


package aliyun

import (
	"strings"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

const (
	VpcAvailable = "Available"
	VpcPending   = "Pending"
)

// "CidrBlock":"172.31.0.0/16","CreationTime":"2017-03-19T13:37:40Z","Description":"System created default VPC.","IsDefault":true,"RegionId":"cn-hongkong","Status":"Available","UserCidrs":{"UserCidr":[]},"VRouterId":"vrt-j6c00qrol733dg36iq4qj","VSwitchIds":{"VSwitchId":["vsw-j6c3gig5ub4fmi2veyrus"]},"VpcId":"vpc-j6c86z3sh8ufhgsxwme0q","VpcName":""

type SUserCIDRs struct {
	UserCidr []string
}

type SVSwitchIds struct {
	VSwitchId []string
}

type SVpc struct {
	multicloud.SVpc
	AliyunTags

	region *SRegion

	iwires []cloudprovider.ICloudWire

	routeTables []cloudprovider.ICloudRouteTable

	CidrBlock     string
	Ipv6CidrBlock string
	CreationTime  time.Time
	Description   string
	IsDefault     bool
	RegionId      string
	Status        string
	UserCidrs     SUserCIDRs
	VRouterId     string
	VSwitchIds    SVSwitchIds
	VpcId         string
	VpcName       string
}

func (self *SVpc) GetId() string {
	return self.VpcId
}

func (self *SVpc) GetName() string {
	if len(self.VpcName) > 0 {
		return self.VpcName
	}
	return self.VpcId
}

func (self *SVpc) GetGlobalId() string {
	return self.VpcId
}

func (self *SVpc) IsEmulated() bool {
	return false
}

func (self *SVpc) GetIsDefault() bool {
	return self.IsDefault
}

func (self *SVpc) GetCidrBlock() string {
	return self.CidrBlock
}

func (self *SVpc) GetCidrBlock6() string {
	return self.Ipv6CidrBlock
}

func (self *SVpc) GetStatus() string {
	return strings.ToLower(self.Status)
}

func (self *SVpc) Refresh() error {
	new, err := self.region.getVpc(self.VpcId)
	if err != nil {
		return err
	}
	return jsonutils.Update(self, new)
}

func (self *SVpc) GetRegion() cloudprovider.ICloudRegion {
	return self.region
}

func (self *SVpc) addWire(wire *SWire) {
	if self.iwires == nil {
		self.iwires = make([]cloudprovider.ICloudWire, 0)
	}
	self.iwires = append(self.iwires, wire)
}

func (self *SVpc) getWireByZoneId(zoneId string) *SWire {
	for i := 0; i < len(self.iwires); i += 1 {
		wire := self.iwires[i].(*SWire)
		if wire.zone.ZoneId == zoneId {
			return wire
		}
	}
	return nil
}

func (self *SVpc) fetchVSwitches() error {
	switches, total, err := self.region.GetVSwitches(nil, self.VpcId, 0, 50)
	if err != nil {
		return err
	}
	if total > len(switches) {
		switches, _, err = self.region.GetVSwitches(nil, self.VpcId, 0, total)
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(switches); i += 1 {
		wire := self.getWireByZoneId(switches[i].ZoneId)
		if wire != nil {
			switches[i].wire = wire
			wire.addNetwork(&switches[i])
		}
	}
	return nil
}

func (self *SVpc) GetIWires() ([]cloudprovider.ICloudWire, error) {
	if self.iwires == nil {
		err := self.fetchVSwitches()
		if err != nil {
			return nil, err
		}
	}
	return self.iwires, nil
}

func (self *SVpc) GetIWireById(wireId string) (cloudprovider.ICloudWire, error) {
	if self.iwires == nil {
		err := self.fetchVSwitches()
		if err != nil {
			return nil, err
		}
	}
	for i := 0; i < len(self.iwires); i += 1 {
		if self.iwires[i].GetGlobalId() == wireId {
			return self.iwires[i], nil
		}
	}
	return nil, cloudprovider.ErrNotFound
}

func (self *SVpc) GetISecurityGroups() ([]cloudprovider.ICloudSecurityGroup, error) {
	groups, err := self.region.GetSecurityGroups(self.VpcId, "", nil)
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.ICloudSecurityGroup{}
	for i := range groups {
		groups[i].region = self.region
		ret = append(ret, &groups[i])
	}
	return ret, nil
}

func (self *SVpc) fetchRouteTables() error {
	routeTables := make([]*SRouteTable, 0)
	for {
		parts, total, err := self.RemoteGetRouteTableList(len(routeTables), 50)
		if err != nil {
			return err
		}
		routeTables = append(routeTables, parts...)
		if len(routeTables) >= total {
			break
		}
	}
	self.routeTables = make([]cloudprovider.ICloudRouteTable, len(routeTables))
	for i := 0; i < len(routeTables); i++ {
		routeTables[i].vpc = self
		self.routeTables[i] = routeTables[i]
	}
	return nil
}

func (self *SVpc) GetIRouteTables() ([]cloudprovider.ICloudRouteTable, error) {
	if self.routeTables == nil {
		err := self.fetchRouteTables()
		if err != nil {
			return nil, err
		}
	}
	return self.routeTables, nil
}

func (self *SVpc) GetIRouteTableById(routeTableId string) (cloudprovider.ICloudRouteTable, error) {
	tables, err := self.GetIRouteTables()
	if err != nil {
		return nil, errors.Wrapf(err, "GetIRouteTables")
	}
	for i := range tables {
		if tables[i].GetGlobalId() == routeTableId {
			return tables[i], nil
		}
	}
	return nil, errors.Wrapf(cloudprovider.ErrNotFound, routeTableId)
}

func (self *SVpc) Delete() error {
	return self.region.DeleteVpc(self.VpcId)
}

func (self *SVpc) getNatGateways() ([]SNatGateway, error) {
	natgatways := make([]SNatGateway, 0)
	gwTotal := -1
	for gwTotal < 0 || len(natgatways) < gwTotal {
		parts, total, err := self.region.GetNatGateways(self.VpcId, "", len(natgatways), 50)
		if err != nil {
			return nil, err
		}
		if len(parts) > 0 {
			natgatways = append(natgatways, parts...)
		}
		gwTotal = total
	}
	for i := 0; i < len(natgatways); i += 1 {
		natgatways[i].vpc = self
	}
	return natgatways, nil
}

func (self *SVpc) GetINatGateways() ([]cloudprovider.ICloudNatGateway, error) {
	nats := []SNatGateway{}
	for {
		parts, total, err := self.region.GetNatGateways(self.VpcId, "", len(nats), 50)
		if err != nil {
			return nil, err
		}
		nats = append(nats, parts...)
		if len(nats) >= total {
			break
		}
	}
	inats := []cloudprovider.ICloudNatGateway{}
	for i := 0; i < len(nats); i++ {
		nats[i].vpc = self
		inats = append(inats, &nats[i])
	}
	return inats, nil
}

func (self *SVpc) GetAuthorityOwnerId() string {
	return self.region.client.ownerId
}

func (self *SRegion) GrantInstanceToCen(opts *cloudprovider.SVpcJointInterVpcNetworkOption, instance SCenAttachInstanceInput) error {
	params := make(map[string]string)
	params["CenId"] = opts.InterVpcNetworkId
	params["CenOwnerId"] = opts.NetworkAuthorityOwnerId

	params["InstanceId"] = instance.InstanceId
	params["InstanceType"] = instance.InstanceType
	params["RegionId"] = instance.InstanceRegion
	_, err := self.vpcRequest("GrantInstanceToCen", params)
	if err != nil {
		return errors.Wrapf(err, `self.vpcRequest("GrantInstanceToCen", %s)`, jsonutils.Marshal(params).String())
	}
	return nil
}

func (self *SVpc) ProposeJoinICloudInterVpcNetwork(opts *cloudprovider.SVpcJointInterVpcNetworkOption) error {
	instance := SCenAttachInstanceInput{
		InstanceType:   "VPC",
		InstanceId:     self.GetId(),
		InstanceRegion: self.region.GetId(),
	}
	err := self.region.GrantInstanceToCen(opts, instance)
	if err != nil {
		return errors.Wrapf(err, "self.region.GrantInstanceToCen(%s,%s)", self.GetId(), jsonutils.Marshal(opts).String())
	}
	return nil
}

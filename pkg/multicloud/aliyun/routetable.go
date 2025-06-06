
package aliyun

import (
	"fmt"
	"strings"
	"time"

	"yunion.io/x/log"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

// {"CreationTime":"2017-03-19T13:37:40Z","RouteEntrys":{"RouteEntry":[{"DestinationCidrBlock":"172.31.32.0/20","InstanceId":"","NextHopType":"local","NextHops":{"NextHop":[]},"RouteTableId":"vtb-j6c60lectdi80rk5xz43g","Status":"Available","Type":"System"},{"DestinationCidrBlock":"100.64.0.0/10","InstanceId":"","NextHopType":"service","NextHops":{"NextHop":[]},"RouteTableId":"vtb-j6c60lectdi80rk5xz43g","Status":"Available","Type":"System"}]},"RouteTableId":"vtb-j6c60lectdi80rk5xz43g","RouteTableType":"System","VRouterId":"vrt-j6c00qrol733dg36iq4qj"}

type SNextHops struct {
	NextHop []string
}

type SRouteEntry struct {
	multicloud.SResourceBase
	AliyunTags
	routeTable *SRouteTable

	RouteTableId         string
	Type                 string
	DestinationCidrBlock string
	NextHopType          string
	InstanceId           string
	RouteEntryId         string
	RouteEntryName       string
	NextHops             SNextHops
}

func (route *SRouteEntry) GetId() string {
	return fmt.Sprintf("%s-%s-%s", route.RouteTableId, route.DestinationCidrBlock, route.NextHopType)
}

func (route *SRouteEntry) GetName() string {
	return route.RouteEntryName
}

func (route *SRouteEntry) GetGlobalId() string {
	return route.GetId()
}

func (route *SRouteEntry) GetStatus() string {
	return api.ROUTE_ENTRY_STATUS_AVAILIABLE
}

func (route *SRouteEntry) Refresh() error {
	return nil
}

func (route *SRouteEntry) IsEmulated() bool {
	return false
}

// Custom：自定义路由。 System：系统路由。
func (route *SRouteEntry) GetType() string {
	return route.Type
}

func (route *SRouteEntry) GetCidr() string {
	return route.DestinationCidrBlock
}

func (route *SRouteEntry) GetNextHopType() string {
	switch route.NextHopType {
	case "Instance":
		return api.NEXT_HOP_TYPE_INSTANCE
	case "HaVip":
		return api.NEXT_HOP_TYPE_HAVIP
	case "VpnGateway":
		return api.NEXT_HOP_TYPE_VPN
	case "NatGateway":
		return api.NEXT_HOP_TYPE_NAT
	case "NetworkInterface":
		return api.NEXT_HOP_TYPE_NETWORK
	case "RouterInterface":
		return api.NEXT_HOP_TYPE_ROUTER
	case "IPv6Gateway":
		return api.NEXT_HOP_TYPE_IPV6
	case "InternetGateway":
		return api.NEXT_HOP_TYPE_INTERNET
	case "Next_HOP_TYPE_EGRESS_INTERNET":
		return api.NEXT_HOP_TYPE_EGRESS_INTERNET
	default:
		return ""
	}
}

func (route *SRouteEntry) GetNextHop() string {
	return route.InstanceId
}

type SRouteEntrys struct {
	RouteEntry []*SRouteEntry
}

type SRouteTable struct {
	multicloud.SResourceBase
	AliyunTags
	region *SRegion
	vpc    *SVpc
	routes []cloudprovider.ICloudRoute

	VpcId        string
	CreationTime time.Time
	RouteEntrys  SRouteEntrys
	VRouterId    string
	Description  string

	RouteTableId   string
	RouteTableName string
	RouteTableType string
	RouterId       string
	RouterType     string
	VSwitchIds     SRouteTableVSwitchIds
}

type SRouteTableVSwitchIds struct {
	VSwitchId []string
}

type sDescribeRouteTablesResponseRouteTables struct {
	RouteTable []SRouteTable
}

type sDescribeRouteTablesResponse struct {
	RouteTables sDescribeRouteTablesResponseRouteTables
	TotalCount  int
}

func (self *SRouteTable) GetDescription() string {
	return self.Description
}

func (self *SRouteTable) GetId() string {
	return self.RouteTableId
}

func (self *SRouteTable) GetGlobalId() string {
	return self.RouteTableId
}

func (self *SRouteTable) GetName() string {
	return self.RouteTableName
}

func (self *SRouteTable) GetRegionId() string {
	return self.region.RegionId
}

// VRouter：VPC路由器。 VBR：边界路由器。
func (self *SRouteTable) GetType() cloudprovider.RouteTableType {
	switch self.RouteTableType {
	case "System":
		return cloudprovider.RouteTableTypeSystem
	case "Custom":
		return cloudprovider.RouteTableTypeCustom
	default:
		return cloudprovider.RouteTableTypeSystem
	}
}

func (self *SRouteTable) GetVpcId() string {
	return self.VpcId
}

func (self *SRouteTable) GetIRoutes() ([]cloudprovider.ICloudRoute, error) {
	if self.routes == nil {
		err := self.fetchRoutes()
		if err != nil {
			return nil, err
		}
	}
	return self.routes, nil
}

func (self *SRouteTable) GetStatus() string {
	return api.ROUTE_TABLE_AVAILABLE
}

func (self *SRouteTable) IsEmulated() bool {
	return false
}

func (self *SRouteTable) Refresh() error {
	return nil
}

func (self *SRouteTable) fetchRoutes() error {
	routes := []*SRouteEntry{}
	for {
		parts, total, err := self.RemoteGetRoutes(len(routes), 50)
		if err != nil {
			return err
		}
		routes = append(routes, parts...)
		if len(routes) >= total {
			break
		}
	}
	self.routes = make([]cloudprovider.ICloudRoute, len(routes))
	for i := 0; i < len(routes); i++ {
		routes[i].routeTable = self
		self.routes[i] = routes[i]
	}
	return nil
}

func (self *SRouteTable) RemoteGetRoutes(offset int, limit int) ([]*SRouteEntry, int, error) {
	if limit > 50 || limit <= 0 {
		limit = 50
	}
	params := make(map[string]string)
	params["RouteTableId"] = self.RouteTableId
	params["PageSize"] = fmt.Sprintf("%d", limit)
	params["PageNumber"] = fmt.Sprintf("%d", (offset/limit)+1)

	body, err := self.region.ecsRequest("DescribeRouteTables", params)
	if err != nil {
		log.Errorf("RemoteGetRoutes fail %s", err)
		return nil, 0, err
	}

	resp := sDescribeRouteTablesResponse{}
	err = body.Unmarshal(&resp)
	if err != nil {
		log.Errorf("Unmarshal routeEntrys fail %s", err)
		return nil, 0, err
	}
	routeTables := resp.RouteTables.RouteTable
	if len(routeTables) != 1 {
		return nil, 0, fmt.Errorf("expecting 1 route table, got %d", len(routeTables))
	}
	routeTable := routeTables[0]
	return routeTable.RouteEntrys.RouteEntry, resp.TotalCount, nil
}

func (self *SRouteTable) GetAssociations() []cloudprovider.RouteTableAssociation {
	result := []cloudprovider.RouteTableAssociation{}
	switch self.RouterType {
	case "VRouter":
		for i := range self.VSwitchIds.VSwitchId {
			association := cloudprovider.RouteTableAssociation{
				AssociationId:        self.RouteTableId + ":" + self.VSwitchIds.VSwitchId[i],
				AssociationType:      cloudprovider.RouteTableAssociaToSubnet,
				AssociatedResourceId: self.VSwitchIds.VSwitchId[i],
			}
			result = append(result, association)
		}
	case "VBR":
		association := cloudprovider.RouteTableAssociation{
			AssociationId:        self.RouteTableId + ":" + self.RouterId,
			AssociationType:      cloudprovider.RouteTableAssociaToRouter,
			AssociatedResourceId: self.RouterId,
		}
		result = append(result, association)
	}

	return result
}

func (self *SRouteTable) CreateRoute(route cloudprovider.RouteSet) error {
	return cloudprovider.ErrNotSupported
}

func (self *SRouteTable) UpdateRoute(route cloudprovider.RouteSet) error {
	return cloudprovider.ErrNotSupported
}

func (self *SRouteTable) RemoveRoute(route cloudprovider.RouteSet) error {
	return cloudprovider.ErrNotSupported
}

func (self *SVpc) RemoteGetRouteTableList(offset int, limit int) ([]*SRouteTable, int, error) {
	if limit > 50 || limit <= 0 {
		limit = 50
	}
	params := make(map[string]string)
	params["VpcId"] = self.VpcId
	params["PageSize"] = fmt.Sprintf("%d", limit)
	params["PageNumber"] = fmt.Sprintf("%d", (offset/limit)+1)

	body, err := self.region.vpcRequest("DescribeRouteTableList", params)
	if err != nil {
		log.Errorf("RemoteGetRouteTableList fail %s", err)
		return nil, 0, err
	}

	routeTables := make([]*SRouteTable, 0)
	err = body.Unmarshal(&routeTables, "RouterTableList", "RouterTableListType")
	if err != nil {
		log.Errorf("Unmarshal routeTables fail %s", err)
		return nil, 0, err
	}
	for _, routeTable := range routeTables {
		routeTable.region = self.region
	}
	total, _ := body.Int("TotalCount")
	return routeTables, int(total), nil
}

func (region *SRegion) AssociateRouteTable(rtableId string, vswitchId string) error {
	params := make(map[string]string)
	params["RegionId"] = region.RegionId
	params["RouteTableId"] = rtableId
	params["VSwitchId"] = vswitchId
	_, err := region.vpcRequest("AssociateRouteTable", params)
	return err
}

func (region *SRegion) UnassociateRouteTable(rtableId string, vswitchId string) error {
	params := make(map[string]string)
	params["RegionId"] = region.RegionId
	params["RouteTableId"] = rtableId
	params["VSwitchId"] = vswitchId
	_, err := region.vpcRequest("UnassociateRouteTable", params)
	return err
}

func (routeTable *SRouteTable) IsSystem() bool {
	return strings.ToLower(routeTable.RouteTableType) == "system"
}

package cloudprovider

type RouteTableAssociationType string

type RouteTableType string

const (
	RouteTableAssociaToRouter = RouteTableAssociationType("Router")
	RouteTableAssociaToSubnet = RouteTableAssociationType("Subnet")
)

const (
	RouteTableTypeSystem = RouteTableType("System")
	RouteTableTypeCustom = RouteTableType("Custom")
)

type RouteTableAssociation struct {
	AssociationId        string
	AssociationType      RouteTableAssociationType
	AssociatedResourceId string
}

func (self RouteTableAssociation) GetGlobalId() string {
	return self.AssociationId
}

type RouteSet struct {
	RouteId     string
	Destination string // route destination
	NextHopType string // route next hop type
	NextHop     string // route next hop (ip or id)
}

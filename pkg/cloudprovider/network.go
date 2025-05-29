package cloudprovider

type SNetworkCreateOptions struct {
	Name           string
	Desc           string
	ProjectId      string
	Cidr           string
	AssignPublicIp bool
}

type SWireCreateOptions struct {
	Name           string
	ZoneId         string
	Bandwidth      int
	Mtu            int
	AssignPublicIp bool
}

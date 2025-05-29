package cloudprovider

type VpcPeeringConnectionCreateOptions struct {
	Name          string
	Desc          string
	PeerVpcId     string
	PeerAccountId string
	PeerRegionId  string
	Bandwidth     int //qcloud cross region,Mbps
}

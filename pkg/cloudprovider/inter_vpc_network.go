package cloudprovider

type SInterVpcNetworkAttachVpcOption struct {
	VpcId               string
	VpcRegionId         string
	VpcAuthorityOwnerId string
}

type SInterVpcNetworkDetachVpcOption struct {
	VpcId               string
	VpcRegionId         string
	VpcAuthorityOwnerId string
}

type SVpcJointInterVpcNetworkOption struct {
	InterVpcNetworkId       string
	NetworkAuthorityOwnerId string
}

type SInterVpcNetworkCreateOptions struct {
	Name string
	Desc string
}

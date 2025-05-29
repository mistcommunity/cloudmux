package cloudprovider

import "yunion.io/x/pkg/util/billing"

// These two structures are designed for modifying snat table and dnat table.
// There is a so strange point that they have both field of ExternalIP and ExternalIPID.
// The reason is that you must pass ExternalIPID to modify in Huawei Cloud for now.
// So please construct a valid parameter ExternalIPID instead of ExternalIP in Huawei Cloud.
// A more general approach is to pass both valid parameters.

type SNatSRule struct {
	SourceCIDR string
	NetworkID  string

	ExternalIP   string
	ExternalIPID string
}

type SNatDRule struct {
	Protocol string

	InternalIP   string
	InternalPort int

	ExternalIP   string
	ExternalIPID string
	ExternalPort int
}

type NatGatewayCreateOptions struct {
	Name      string
	VpcId     string
	NetworkId string
	Desc      string
	NatSpec   string

	BillingCycle *billing.SBillingCycle
}

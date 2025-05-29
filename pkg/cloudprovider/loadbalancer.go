package cloudprovider

import "yunion.io/x/pkg/util/billing"

type SLoadbalancerCreateOptions struct {
	Name             string
	Desc             string
	ZoneId           string
	SlaveZoneId      string
	VpcId            string
	NetworkIds       []string
	EipId            string // eip id
	Address          string
	AddressType      string
	LoadbalancerSpec string
	ChargeType       string
	EgressMbps       int
	BillingCycle     *billing.SBillingCycle
	ProjectId        string
	Tags             map[string]string
}

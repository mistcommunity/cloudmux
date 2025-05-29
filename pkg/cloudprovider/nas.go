package cloudprovider

import "yunion.io/x/pkg/util/billing"

type FileSystemCraeteOptions struct {
	Name           string
	Desc           string
	VpcId          string
	NetworkId      string
	Capacity       int64
	StorageType    string
	Protocol       string
	FileSystemType string
	ZoneId         string

	BillingCycle *billing.SBillingCycle
}

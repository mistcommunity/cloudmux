package cloudprovider

type TElasticipChargeType string

var (
	ElasticipChargeTypeByTraffic   = TElasticipChargeType("traffic")
	ElasticipChargeTypeByBandwidth = TElasticipChargeType("bandwidth")
)

type SEip struct {
	Name              string
	BandwidthMbps     int
	ChargeType        string
	BGPType           string
	NetworkExternalId string
	Ip                string
	ProjectId         string
	Tags              map[string]string
}

type AssociateConfig struct {
	InstanceId    string
	AssociateType string
	Bandwidth     int
	ChargeType    string
}

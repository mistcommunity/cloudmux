
package multicloud

type SElasticcacheBase struct {
	SVirtualResourceBase
	SBillingBase
}

func (self *SElasticcacheBase) GetBandwidth() int {
	return 0
}

func (self *SElasticcacheBase) GetConnections() int {
	return 0
}

type SElasticcacheBackupBase struct {
	SResourceBase
}

type SElasticcacheAccountBase struct {
	SResourceBase
}

type SElasticcacheAclBase struct {
	SResourceBase
}

type SElasticcacheParameterBase struct {
	SResourceBase
}

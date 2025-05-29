package cloudprovider

import "yunion.io/x/pkg/util/billing"

type TBackupMethod string

const (
	BackupMethodLogical  = TBackupMethod("Logical")
	BackupMethodPhysical = TBackupMethod("Physical")
	BackupMethodUnknown  = TBackupMethod("")
)

type SDBInstanceNetwork struct {
	IP        string
	NetworkId string
}

type SExtraIp struct {
	IP  string
	URL string
}

type SZoneInfo struct {
	Zone1  string
	Zone2  string
	Zone3  string
	ZoneId string
}

type SInstanceType struct {
	InstanceType string
	SZoneInfo
}

type SManagedDBInstanceCreateConfig struct {
	Name        string
	Description string
	StorageType string
	DiskSizeGB  int
	SInstanceType
	VcpuCount        int
	VmemSizeMb       int
	VpcId            string
	SecgroupIds      []string
	NetworkId        string
	Address          string
	Engine           string
	EngineVersion    string
	Category         string
	Port             int
	MasterInstanceId string
	Password         string
	Username         string
	ProjectId        string

	MultiAz bool

	BillingCycle *billing.SBillingCycle
	Tags         map[string]string

	// 仅从备份恢复到新实例用到
	RdsId    string
	BackupId string
}

type SManagedDBInstanceChangeConfig struct {
	DiskSizeGB   int
	InstanceType string
}

type SDBInstanceDatabaseCreateConfig struct {
	Name         string
	CharacterSet string
	Description  string
}

type SDBInstancePrivilege struct {
	Account   string
	Database  string
	Privilege string
}

type SDBInstanceAccountCreateConfig struct {
	Name        string
	Host        string
	Description string
	Password    string
}

type SDBInstanceBackupCreateConfig struct {
	Name        string
	Description string
	Databases   []string
}

type SDBInstanceRecoveryConfig struct {
	BackupId                   string
	Databases                  map[string]string
	OriginDBInstanceExternalId string
}

type SDBInstanceUpdateOptions struct {
	NAME        string
	Description string
}

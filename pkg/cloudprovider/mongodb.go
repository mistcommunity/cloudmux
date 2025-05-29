package cloudprovider

import "time"

// 备份状态
type TMongoDBBackupStatus string

// 备份方法
type TMongoDBBackupMethod string

// 备份方式
type TMongoDBBackupType string

const (
	MongoDBBackupStatusCreating  = TMongoDBBackupStatus("creating")
	MongoDBBackupStatusAvailable = TMongoDBBackupStatus("available")
	MongoDBBackupStatusFailed    = TMongoDBBackupStatus("failed")
	MongoDBBackupStatusUnknown   = TMongoDBBackupStatus("unknown")

	MongoDBBackupMethodPhysical = TMongoDBBackupMethod("physical")
	MongoDBBackupMethodLogical  = TMongoDBBackupMethod("logical")

	MongoDBBackupTypeAuto   = TMongoDBBackupType("auto")
	MongoDBBackupTypeManual = TMongoDBBackupType("manual")
)

type SMongoDBBackup struct {
	Name         string
	Description  string
	StartTime    time.Time
	EndTime      time.Time
	Status       TMongoDBBackupStatus
	BackupMethod TMongoDBBackupMethod
	BackupType   TMongoDBBackupType
	BackupSizeKb int
}

type SMongoDBBackups struct {
	Data  []SMongoDBBackup
	Total int
}

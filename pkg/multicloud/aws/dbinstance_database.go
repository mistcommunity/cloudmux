package aws

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SDBInstanceDatabase struct {
	multicloud.SDBInstanceDatabaseBase
	AwsTags

	DBName string
}

func (database *SDBInstanceDatabase) GetId() string {
	return database.DBName
}

func (database *SDBInstanceDatabase) GetGlobalId() string {
	return database.DBName
}

func (database *SDBInstanceDatabase) GetName() string {
	return database.DBName
}

func (database *SDBInstanceDatabase) GetStatus() string {
	return api.DBINSTANCE_DATABASE_RUNNING
}

func (database *SDBInstanceDatabase) GetCharacterSet() string {
	return ""
}

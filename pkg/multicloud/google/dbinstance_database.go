
package google

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SSqlserverDatabaseDetails struct {
	CompatibilityLevel int
	RecoveryModel      string
}

type SDBInstanceDatabase struct {
	multicloud.SResourceBase
	GoogleTags
	rds                      *SDBInstance
	Kind                     string
	Collation                string
	Etag                     string
	Name                     string
	Instance                 string
	SelfLink                 string
	Charset                  string
	Project                  string
	SqlserverDatabaseDetails SSqlserverDatabaseDetails
}

func (region *SRegion) GetDBInstanceDatabases(instance string) ([]SDBInstanceDatabase, error) {
	databases := []SDBInstanceDatabase{}
	params := map[string]string{}
	resource := fmt.Sprintf("instances/%s/databases", instance)
	err := region.RdsListAll(resource, params, &databases)
	if err != nil {
		return nil, errors.Wrap(err, "RdsListAll")
	}
	return databases, nil
}

func (region *SRegion) DeleteDBInstanceDatabase(id string) error {
	return region.rdsDelete(id)
}

func (database *SDBInstanceDatabase) Delete() error {
	return database.rds.region.DeleteDBInstanceDatabase(database.SelfLink)
}

func (database *SDBInstanceDatabase) GetCharacterSet() string {
	return database.Charset
}

func (database *SDBInstanceDatabase) GetGlobalId() string {
	return database.Name
}

func (database *SDBInstanceDatabase) GetId() string {
	return database.SelfLink
}

func (database *SDBInstanceDatabase) GetName() string {
	return database.Name
}

func (database *SDBInstanceDatabase) GetStatus() string {
	return api.DBINSTANCE_DATABASE_RUNNING
}

func (database *SDBInstanceDatabase) IsEmulated() bool {
	return false
}

func (database *SDBInstanceDatabase) Refresh() error {
	_database := SDBInstanceDatabase{}
	err := database.rds.region.rdsGet(database.SelfLink, &_database)
	if err != nil {
		return errors.Wrap(err, "rdsGet")
	}
	return jsonutils.Update(database, _database)
}

func (region *SRegion) CreateDatabase(instanceId string, name, charset string) error {
	body := map[string]interface{}{
		"charset": charset,
		"name":    name,
	}
	return region.rdsDo(instanceId, "databases", nil, jsonutils.Marshal(body))
}

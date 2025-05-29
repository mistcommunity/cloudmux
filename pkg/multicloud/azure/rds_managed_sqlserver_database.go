
package azure

import (
	"net/url"
	"strings"
	"time"

	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SManagedSQLServerDatabase struct {
	multicloud.SDBInstanceDatabaseBase
	AzureTags
	rds *SManagedSQLServer

	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		Collation                string    `json:"collation"`
		Creationdate             time.Time `json:"creationDate"`
		Defaultsecondarylocation string    `json:"defaultSecondaryLocation"`
		Status                   string    `json:"status"`
	} `json:"properties"`
	Type string `json:"type"`
}

func (self *SManagedSQLServerDatabase) GetName() string {
	return self.Name
}

func (self *SManagedSQLServerDatabase) GetId() string {
	return self.ID
}

func (self *SManagedSQLServerDatabase) GetStatus() string {
	switch self.Properties.Status {
	case "Online":
		return api.DBINSTANCE_DATABASE_RUNNING
	case "Creating":
		return api.DBINSTANCE_DATABASE_CREATING
	default:
		return strings.ToLower(self.Properties.Status)
	}
}

func (self *SManagedSQLServerDatabase) GetGlobalId() string {
	return strings.ToLower(self.Name)
}

func (self *SManagedSQLServerDatabase) GetCharacterSet() string {
	return self.Properties.Collation
}

func (self *SRegion) GetManagedSQLServerDatabases(id string) ([]SManagedSQLServerDatabase, error) {
	result := struct {
		Value []SManagedSQLServerDatabase
	}{}
	return result.Value, self.get(id+"/databases", url.Values{}, &result)
}

func (self *SManagedSQLServer) GetIDBInstanceDatabases() ([]cloudprovider.ICloudDBInstanceDatabase, error) {
	dbs, err := self.region.GetManagedSQLServerDatabases(self.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "GetSQLServerDatabases")
	}
	ret := []cloudprovider.ICloudDBInstanceDatabase{}
	for i := range dbs {
		dbs[i].rds = self
		ret = append(ret, &dbs[i])
	}
	return ret, nil
}

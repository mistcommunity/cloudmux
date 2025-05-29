
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

type SSQLServerDatabase struct {
	multicloud.SDBInstanceDatabaseBase
	AzureTags
	rds *SSQLServer
	Sku struct {
		Name     string `json:"name"`
		Tier     string `json:"tier"`
		Capacity int    `json:"capacity"`
	} `json:"sku"`
	Kind       string `json:"kind"`
	Properties struct {
		Collation                     string    `json:"collation"`
		Maxsizebytes                  int64     `json:"maxSizeBytes"`
		Status                        string    `json:"status"`
		Databaseid                    string    `json:"databaseId"`
		Creationdate                  string    `json:"creationDate"`
		Currentserviceobjectivename   string    `json:"currentServiceObjectiveName"`
		Requestedserviceobjectivename string    `json:"requestedServiceObjectiveName"`
		Defaultsecondarylocation      string    `json:"defaultSecondaryLocation"`
		Catalogcollation              string    `json:"catalogCollation"`
		Licensetype                   string    `json:"licenseType"`
		Maxlogsizebytes               int       `json:"maxLogSizeBytes"`
		Storageaccounttype            string    `json:"storageAccountType"`
		Zoneredundant                 bool      `json:"zoneRedundant"`
		Readscale                     string    `json:"readScale"`
		Earliestrestoredate           time.Time `json:"earliestRestoreDate"`
		Currentsku                    struct {
			Name     string `json:"name"`
			Tier     string `json:"tier"`
			Capacity int    `json:"capacity"`
			Family   string `json:"family"`
		} `json:"currentSku"`
	} `json:"properties"`
	Location string `json:"location"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

func (self *SSQLServerDatabase) GetName() string {
	return self.Name
}

func (self *SSQLServerDatabase) GetId() string {
	return self.ID
}

func (self *SSQLServerDatabase) GetStatus() string {
	switch self.Properties.Status {
	case "Online":
		return api.DBINSTANCE_DATABASE_RUNNING
	case "Creating":
		return api.DBINSTANCE_DATABASE_CREATING
	default:
		return strings.ToLower(self.Properties.Status)
	}
}

func (self *SSQLServerDatabase) GetDiskSizeMb() int {
	return int(self.Properties.Maxsizebytes / 1024 / 1024)
}

func (self *SSQLServerDatabase) GetVcpuCount() int {
	if len(self.Properties.Currentsku.Family) > 0 {
		return self.Properties.Currentsku.Capacity
	}
	return 0
}

func (self *SSQLServerDatabase) GetVmemSizeMb() int {
	if len(self.Properties.Currentsku.Family) > 0 {
		return int(5.2 * 1024 * float32(self.Properties.Currentsku.Capacity))
	}
	return 0
}

func (self *SSQLServerDatabase) GetDTU() int {
	if len(self.Properties.Currentsku.Family) == 0 {
		return self.Properties.Currentsku.Capacity
	}
	return 0
}

func (self *SSQLServerDatabase) GetGlobalId() string {
	return strings.ToLower(self.Name)
}

func (self *SSQLServerDatabase) GetCharacterSet() string {
	return self.Properties.Collation
}

func (self *SRegion) GetSQLServerDatabases(id string) ([]SSQLServerDatabase, error) {
	result := struct {
		Value []SSQLServerDatabase
	}{}
	return result.Value, self.get(id+"/databases", url.Values{}, &result)
}

func (self *SSQLServer) GetIDBInstanceDatabases() ([]cloudprovider.ICloudDBInstanceDatabase, error) {
	dbs, err := self.fetchDatabase()
	if err != nil {
		return nil, errors.Wrapf(err, "fetchDatabase")
	}
	ret := []cloudprovider.ICloudDBInstanceDatabase{}
	for i := range dbs {
		dbs[i].rds = self
		ret = append(ret, &dbs[i])
	}
	return ret, nil
}

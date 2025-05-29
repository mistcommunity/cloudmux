
package azure

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SManagedSQLServerAccount struct {
	multicloud.SDBInstanceAccountBase
	rds *SManagedSQLServer

	Name string
}

func (self *SManagedSQLServerAccount) GetName() string {
	return self.Name
}

func (self *SManagedSQLServer) GetIDBInstanceAccounts() ([]cloudprovider.ICloudDBInstanceAccount, error) {
	accounts := []cloudprovider.ICloudDBInstanceAccount{}
	if len(self.Properties.Administratorlogin) > 0 {
		account := &SManagedSQLServerAccount{rds: self, Name: self.Properties.Administratorlogin}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

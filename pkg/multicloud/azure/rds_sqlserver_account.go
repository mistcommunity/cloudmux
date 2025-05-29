
package azure

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SSQLServerAccount struct {
	multicloud.SDBInstanceAccountBase
	rds *SSQLServer

	Name string
}

func (self *SSQLServerAccount) GetName() string {
	return self.Name
}

func (self *SSQLServer) GetIDBInstanceAccounts() ([]cloudprovider.ICloudDBInstanceAccount, error) {
	accounts := []cloudprovider.ICloudDBInstanceAccount{}
	if len(self.Properties.Administratorlogin) > 0 {
		account := &SSQLServerAccount{rds: self, Name: self.Properties.Administratorlogin}
		accounts = append(accounts, account)
	}
	return accounts, nil
}


package azure

import "yunion.io/x/cloudmux/pkg/multicloud"

type SDBInstanceAccount struct {
	multicloud.SDBInstanceAccountBase
	AzureTags
	instance *SDBInstance

	AccountName string
}

func (account *SDBInstanceAccount) GetName() string {
	return account.AccountName + "@" + account.instance.Name
}

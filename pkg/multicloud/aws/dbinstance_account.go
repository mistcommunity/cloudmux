package aws

import (
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SDBInstanceAccount struct {
	multicloud.SDBInstanceAccountBase
	instance *SDBInstance

	AccountName string
}

func (account *SDBInstanceAccount) GetName() string {
	return account.AccountName
}

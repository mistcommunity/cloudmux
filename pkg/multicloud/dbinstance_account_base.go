

package multicloud

import (
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SDBInstanceAccountBase struct {
}

func (account *SDBInstanceAccountBase) GetIDBInstanceAccountPrivileges() ([]cloudprovider.ICloudDBInstanceAccountPrivilege, error) {
	return []cloudprovider.ICloudDBInstanceAccountPrivilege{}, nil
}

func (account *SDBInstanceAccountBase) Delete() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Delete")
}

func (account *SDBInstanceAccountBase) GetHost() string {
	return "%"
}

func (account *SDBInstanceAccountBase) GetStatus() string {
	return api.DBINSTANCE_USER_AVAILABLE
}

func (account *SDBInstanceAccountBase) ResetPassword(password string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "ResetPassword")
}

func (backup *SDBInstanceAccountBase) GrantPrivilege(database, privilege string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "GrantPrivilege")
}

func (backup *SDBInstanceAccountBase) RevokePrivilege(database string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "RevokePrivilege")
}


package multicloud

import (
	"context"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/billing"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SDBInstanceBase struct {
	SVirtualResourceBase
	SBillingBase
}

func (instance *SDBInstanceBase) GetConnectionStr() string {
	return ""
}

func (instance *SDBInstanceBase) GetInternalConnectionStr() string {
	return ""
}

func (instance *SDBInstanceBase) GetDBNetworks() ([]cloudprovider.SDBInstanceNetwork, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetDBNetworks")
}

func (instance *SDBInstanceBase) GetIDBInstanceParameters() ([]cloudprovider.ICloudDBInstanceParameter, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIDBInstanceParameters")
}

func (instance *SDBInstanceBase) GetIDBInstanceDatabases() ([]cloudprovider.ICloudDBInstanceDatabase, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIDBInstanceDatabases")
}

func (instance *SDBInstanceBase) GetIDBInstanceAccounts() ([]cloudprovider.ICloudDBInstanceAccount, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIDBInstanceAccounts")
}

func (instance *SDBInstanceBase) GetIDBInstanceBackups() ([]cloudprovider.ICloudDBInstanceBackup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIDBInstanceBackups")
}

func (instance *SDBInstanceBase) GetCategory() string {
	return ""
}

func (instance *SDBInstanceBase) Reboot() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Reboot")
}

func (instance *SDBInstanceBase) Delete() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Delete")
}

func (instance *SDBInstanceBase) GetMasterInstanceId() string {
	return ""
}

func (instance *SDBInstanceBase) GetSecurityGroupIds() ([]string, error) {
	return []string{}, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetSecurityGroupIds")
}

func (self *SDBInstanceBase) SetSecurityGroups(ids []string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "SetSecurityGroups")
}

func (instance *SDBInstanceBase) Renew(bc billing.SBillingCycle) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Renew")
}

func (instance *SDBInstanceBase) ChangeConfig(ctx context.Context, config *cloudprovider.SManagedDBInstanceChangeConfig) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "ChangeConfig")
}

func (instance *SDBInstanceBase) OpenPublicConnection() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "OpenPublicConnection")
}

func (instance *SDBInstanceBase) ClosePublicConnection() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "ClosePublicConnection")
}

func (instance *SDBInstanceBase) CreateDatabase(conf *cloudprovider.SDBInstanceDatabaseCreateConfig) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateDatabase")
}

func (instance *SDBInstanceBase) CreateAccount(conf *cloudprovider.SDBInstanceAccountCreateConfig) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateAccount")
}

func (instance *SDBInstanceBase) CreateIBackup(conf *cloudprovider.SDBInstanceBackupCreateConfig) (string, error) {
	return "", errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateIBackup")
}

func (instance *SDBInstanceBase) RecoveryFromBackup(conf *cloudprovider.SDBInstanceRecoveryConfig) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "RecoveryFromBackup")
}

func (instance *SDBInstanceBase) GetIops() int {
	return 0
}

func (instance *SDBInstanceBase) GetDiskSizeUsedMB() int {
	return 0
}

func (instance *SDBInstanceBase) Update(ctx context.Context, input cloudprovider.SDBInstanceUpdateOptions) error {
	return errors.Wrap(errors.ErrNotImplemented, "DBInsatnce Update")
}


package multicloud

import (
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SDBInstanceBackupBase struct {
	SResourceBase
}

func (backup *SDBInstanceBackupBase) GetBackMode() string {
	return api.BACKUP_MODE_AUTOMATED
}

func (backup *SDBInstanceBackupBase) Delete() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Delete")
}

func (backup *SDBInstanceBackupBase) GetProjectId() string {
	return ""
}

func (backup *SDBInstanceBackupBase) CreateICloudDBInstance(opts *cloudprovider.SManagedDBInstanceCreateConfig) (cloudprovider.ICloudDBInstance, error) {
	return nil, errors.Wrap(cloudprovider.ErrNotImplemented, "CreateICloudDBInstance")
}

func (backup *SDBInstanceBackupBase) GetBackupMethod() cloudprovider.TBackupMethod {
	return cloudprovider.BackupMethodUnknown
}

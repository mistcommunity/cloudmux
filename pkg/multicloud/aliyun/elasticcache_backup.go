
package aliyun

import (
	"fmt"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

// https://help.aliyun.com/document_detail/61081.html?spm=a2c4g.11186623.6.752.3d7630beuL57kI
type SElasticcacheBackup struct {
	multicloud.SElasticcacheBackupBase
	AliyunTags

	cacheDB *SElasticcache

	BackupIntranetDownloadURL string    `json:"BackupIntranetDownloadURL"`
	BackupType                string    `json:"BackupType"`
	BackupEndTime             time.Time `json:"BackupEndTime"`
	BackupMethod              string    `json:"BackupMethod"`
	BackupID                  int64     `json:"BackupId"`
	BackupStartTime           time.Time `json:"BackupStartTime"`
	BackupDownloadURL         string    `json:"BackupDownloadURL"`
	BackupDBNames             string    `json:"BackupDBNames"`
	NodeInstanceID            string    `json:"NodeInstanceId"`
	BackupMode                string    `json:"BackupMode"`
	BackupStatus              string    `json:"BackupStatus"`
	BackupSizeByte            int64     `json:"BackupSize"`
	EngineVersion             string    `json:"EngineVersion"`
}

func (self *SElasticcacheBackup) GetId() string {
	return fmt.Sprintf("%d", self.BackupID)
}

func (self *SElasticcacheBackup) GetName() string {
	return self.GetId()
}

func (self *SElasticcacheBackup) GetGlobalId() string {
	return self.GetId()
}

func (self *SElasticcacheBackup) GetStatus() string {
	switch self.BackupStatus {
	case "Success", "running":
		return api.ELASTIC_CACHE_BACKUP_STATUS_SUCCESS
	case "Failed":
		return api.ELASTIC_CACHE_BACKUP_STATUS_FAILED
	default:
		return self.BackupStatus
	}
}

func (self *SElasticcacheBackup) Refresh() error {
	ibackup, err := self.cacheDB.GetICloudElasticcacheBackup(self.GetId())
	if err != nil {
		return err
	}

	err = jsonutils.Update(self, ibackup.(*SElasticcacheBackup))
	if err != nil {
		return err
	}

	return nil
}

func (self *SElasticcacheBackup) GetBackupSizeMb() int {
	return int(self.BackupSizeByte / 1024 / 1024)
}

func (self *SElasticcacheBackup) GetBackupType() string {
	switch self.BackupType {
	case "FullBackup":
		return api.ELASTIC_CACHE_BACKUP_TYPE_FULL
	case "IncrementalBackup":
		return api.ELASTIC_CACHE_BACKUP_TYPE_INCREMENTAL
	default:
		return self.BackupType
	}
}

func (self *SElasticcacheBackup) GetBackupMode() string {
	switch self.BackupMode {
	case "Automated":
		return api.ELASTIC_CACHE_BACKUP_MODE_AUTOMATED
	case "Manual":
		return api.ELASTIC_CACHE_BACKUP_MODE_MANUAL
	default:
		return self.BackupMode
	}
}

func (self *SElasticcacheBackup) GetDownloadURL() string {
	return self.BackupDownloadURL
}

func (self *SElasticcacheBackup) GetStartTime() time.Time {
	return self.BackupStartTime
}

func (self *SElasticcacheBackup) GetEndTime() time.Time {
	return self.BackupEndTime
}

func (self *SElasticcacheBackup) Delete() error {
	return cloudprovider.ErrNotSupported
}

// https://help.aliyun.com/document_detail/61083.html?spm=a2c4g.11186623.6.753.216f67ddzpyvTL
func (self *SElasticcacheBackup) RestoreInstance(instanceId string) error {
	params := make(map[string]string)
	params["InstanceId"] = instanceId
	params["BackupId"] = self.GetId()

	// 目前没有查询备份ID的接口，因此，备份ID没什么用
	err := DoAction(self.cacheDB.region.kvsRequest, "RestoreInstance", params, nil, nil)
	if err != nil {
		return errors.Wrap(err, "elasticcache.RestoreInstance")
	}

	return nil
}

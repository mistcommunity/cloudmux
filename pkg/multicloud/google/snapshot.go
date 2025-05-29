
package google

import (
	"fmt"
	"time"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
)

type SSnapshot struct {
	region *SRegion
	SResourceBase
	GoogleTags

	CreationTimestamp  time.Time
	Status             string
	SourceDisk         string
	SourceDiskId       string
	DiskSizeGb         int32
	StorageBytes       int
	StorageBytesStatus string
	Licenses           []string
	LabelFingerprint   string
	LicenseCodes       []string
	StorageLocations   []string
	Kind               string
}

func (region *SRegion) GetSnapshots(disk string, maxResults int, pageToken string) ([]SSnapshot, error) {
	snapshots := []SSnapshot{}
	params := map[string]string{}
	if len(disk) > 0 {
		params["filter"] = fmt.Sprintf(`sourceDisk="%s"`, disk)
	}
	resource := "global/snapshots"
	return snapshots, region.List(resource, params, maxResults, pageToken, &snapshots)
}

func (region *SRegion) GetSnapshot(id string) (*SSnapshot, error) {
	snapshot := &SSnapshot{region: region}
	return snapshot, region.Get("global/snapshots", id, snapshot)
}

// CREATING, DELETING, FAILED, READY, or UPLOADING
func (snapshot *SSnapshot) GetStatus() string {
	switch snapshot.Status {
	case "CREATING":
		return api.SNAPSHOT_CREATING
	case "DELETING":
		return api.SNAPSHOT_DELETING
	case "FAILED":
		return api.SNAPSHOT_UNKNOWN
	case "READY", "UPLOADING":
		return api.SNAPSHOT_READY
	default:
		return api.SNAPSHOT_UNKNOWN
	}
}

func (snapshot *SSnapshot) IsEmulated() bool {
	return false
}

func (self *SSnapshot) GetCreatedAt() time.Time {
	return self.CreationTimestamp
}

func (snapshot *SSnapshot) Refresh() error {
	_snapshot, err := snapshot.region.GetSnapshot(snapshot.Id)
	if err != nil {
		return err
	}
	return jsonutils.Update(snapshot, _snapshot)
}

func (snapshot *SSnapshot) GetSizeMb() int32 {
	return snapshot.DiskSizeGb * 1024
}

func (snapshot *SSnapshot) GetDiskId() string {
	return snapshot.SourceDisk
}

func (snapshot *SSnapshot) GetDiskType() string {
	if len(snapshot.Licenses) > 0 {
		return api.DISK_TYPE_SYS
	}
	return api.DISK_TYPE_DATA
}

func (snapshot *SSnapshot) Delete() error {
	return snapshot.region.Delete(snapshot.SelfLink)
}

func (snapshot *SSnapshot) GetProjectId() string {
	return snapshot.region.GetProjectId()
}

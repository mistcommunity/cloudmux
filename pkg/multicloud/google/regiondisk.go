
package google

import (
	"fmt"
	"time"
)

type SRegionDisk struct {
	storage *SStorage

	Id                     string
	CreationTimestamp      time.Time
	Name                   string
	SizeGB                 int
	Zone                   string
	Status                 string
	SelfLink               string
	Type                   string
	LastAttachTimestamp    time.Time
	LastDetachTimestamp    time.Time
	LabelFingerprint       string
	PhysicalBlockSizeBytes string
	Kind                   string
}

func (region *SRegion) GetRegionDisks(storageType string, maxResults int, pageToken string) ([]SRegionDisk, error) {
	disks := []SRegionDisk{}
	params := map[string]string{}
	if len(storageType) > 0 {
		params["filter"] = fmt.Sprintf(`type="%s/%s/projects/%s/regions/%s/diskTypes/%s"`, GOOGLE_COMPUTE_DOMAIN, GOOGLE_API_VERSION, region.GetProjectId(), region.Name, storageType)
	}
	resource := fmt.Sprintf("regions/%s/disks", region.Name)
	return disks, region.List(resource, params, maxResults, pageToken, &disks)
}

func (region *SRegion) GetRegionDisk(id string) (*SRegionDisk, error) {
	disk := &SRegionDisk{}
	return disk, region.GetBySelfId(id, disk)
}


package google

import (
	"fmt"
	"time"
)

type SRegionStorage struct {
	region *SRegion

	CreationTimestamp time.Time
	Name              string
	Description       string
	ValidDiskSize     string
	Zone              string
	SelfLink          string
	DefaultDiskSizeGb string
	Kind              string
}

func (region *SRegion) GetRegionStorages(maxResults int, pageToken string) ([]SRegionStorage, error) {
	storages := []SRegionStorage{}
	resource := fmt.Sprintf("regions/%s/diskTypes", region.Name)
	params := map[string]string{}
	return storages, region.List(resource, params, maxResults, pageToken, &storages)
}

func (region *SRegion) GetRegionStorage(id string) (*SRegionStorage, error) {
	storage := &SRegionStorage{region: region}
	return storage, region.GetBySelfId(id, storage)
}

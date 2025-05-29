
package google

import (
	"fmt"
	"time"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SMachineType struct {
	Id                           string
	CreationTimestamp            time.Time
	Name                         string
	Description                  string
	GuestCpus                    int
	MemoryMb                     int
	ImageSpaceGb                 int
	MaximumPersistentDisks       int
	MaximumPersistentDisksSizeGb int
	Zone                         string
	SelfLink                     string
	IsSharedCpu                  bool
	Kind                         string
}

func (region *SRegion) GetMachineTypes(zone string, maxResults int, pageToken string) ([]SMachineType, error) {
	machines := []SMachineType{}
	params := map[string]string{}
	if len(zone) == 0 {
		return nil, cloudprovider.ErrNotFound
	}
	resource := fmt.Sprintf("zones/%s/machineTypes", zone)
	return machines, region.List(resource, params, maxResults, pageToken, &machines)
}

func (region *SRegion) GetMachineType(id string) (*SMachineType, error) {
	machine := &SMachineType{}
	err := region.client.ecsGet("machineTypes", id, machine)
	if err != nil {
		return nil, err
	}
	return machine, nil
}


package aliyun

import (
	"context"
	"fmt"
	"strings"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SLoadbalancerMasterSlaveBackend struct {
	multicloud.SResourceBase
	AliyunTags
	lbbg *SLoadbalancerMasterSlaveBackendGroup

	ServerId   string
	Weight     int
	Port       int
	ServerType string
}

func (backend *SLoadbalancerMasterSlaveBackend) GetName() string {
	return backend.ServerId
}

func (backend *SLoadbalancerMasterSlaveBackend) GetId() string {
	return fmt.Sprintf("%s/%s", backend.lbbg.MasterSlaveServerGroupId, backend.ServerId)
}

func (backend *SLoadbalancerMasterSlaveBackend) GetGlobalId() string {
	return backend.GetId()
}

func (backend *SLoadbalancerMasterSlaveBackend) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (backend *SLoadbalancerMasterSlaveBackend) IsEmulated() bool {
	return false
}

func (backend *SLoadbalancerMasterSlaveBackend) Refresh() error {
	return nil
}

func (backend *SLoadbalancerMasterSlaveBackend) GetWeight() int {
	return backend.Weight
}

func (backend *SLoadbalancerMasterSlaveBackend) GetPort() int {
	return backend.Port
}

func (backend *SLoadbalancerMasterSlaveBackend) GetBackendType() string {
	return api.LB_BACKEND_GUEST
}

func (backend *SLoadbalancerMasterSlaveBackend) GetBackendRole() string {
	return strings.ToLower(backend.ServerType)
}

func (backend *SLoadbalancerMasterSlaveBackend) GetBackendId() string {
	return backend.ServerId
}

func (backend *SLoadbalancerMasterSlaveBackend) GetIpAddress() string {
	return ""
}

func (backend *SLoadbalancerMasterSlaveBackend) GetProjectId() string {
	return backend.lbbg.GetProjectId()
}

func (backend *SLoadbalancerMasterSlaveBackend) SyncConf(ctx context.Context, port, weight int) error {
	return cloudprovider.ErrNotSupported
}

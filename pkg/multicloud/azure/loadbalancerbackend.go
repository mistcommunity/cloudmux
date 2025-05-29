
package azure

import (
	"context"
	"strings"

	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SLoadbalancerBackend struct {
	multicloud.SResourceBase
	AzureTags

	lbbg                         *SLoadbalancerBackendGroup
	PrivateIPAddress             string
	LoadBalancerBackendAddresses []string
}

func (self *SLoadbalancerBackend) GetId() string {
	return self.PrivateIPAddress + strings.Join(self.LoadBalancerBackendAddresses, ",")
}

func (self *SLoadbalancerBackend) GetName() string {
	return self.PrivateIPAddress + strings.Join(self.LoadBalancerBackendAddresses, ",")
}

func (self *SLoadbalancerBackend) GetGlobalId() string {
	return strings.ToLower(self.GetId())
}

func (self *SLoadbalancerBackend) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SLoadbalancerBackend) GetProjectId() string {
	return ""
}

func (self *SLoadbalancerBackend) GetWeight() int {
	return 0
}

func (self *SLoadbalancerBackend) GetPort() int {
	return 0
}

func (self *SLoadbalancerBackend) GetBackendType() string {
	return api.LB_BACKEND_IP
}

func (self *SLoadbalancerBackend) GetBackendRole() string {
	return api.LB_BACKEND_ROLE_DEFAULT
}

func (self *SLoadbalancerBackend) GetBackendId() string {
	return ""
}

func (self *SLoadbalancerBackend) GetIpAddress() string {
	return self.PrivateIPAddress + strings.Join(self.LoadBalancerBackendAddresses, ",")
}

func (self *SLoadbalancerBackend) SyncConf(ctx context.Context, port, weight int) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SyncConf")
}

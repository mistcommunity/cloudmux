
package google

import (
	"context"
	"fmt"
	"strings"
	"time"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SGlobalLoadbalancerListenerRule struct {
	lbl            *SGlobalLoadbalancerListener
	pathMatcher    PathMatcher
	pathRule       PathRule
	backendService SBackendServices

	ListenerName       string `json:"listener_name"`
	BackendServiceName string `json:"backend_service_name"`
	Domain             string `json:"domain"`
	Path               string `json:"path"`
	Port               string `json:"Port"`
}

func (self SGlobalLoadbalancerListenerRule) GetId() string {
	return fmt.Sprintf("%s::%s::%s", self.lbl.GetGlobalId(), self.Domain, strings.Join(self.pathRule.Paths, ","))
}

func (self SGlobalLoadbalancerListenerRule) GetName() string {
	return fmt.Sprintf("%s::%s::%s", self.lbl.GetName(), self.Domain, strings.Join(self.pathRule.Paths, ","))
}

func (self SGlobalLoadbalancerListenerRule) GetGlobalId() string {
	return self.GetId()
}

func (self SGlobalLoadbalancerListenerRule) GetCreatedAt() time.Time {
	return time.Time{}
}

func (self SGlobalLoadbalancerListenerRule) GetDescription() string {
	return ""
}

func (self SGlobalLoadbalancerListenerRule) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self SGlobalLoadbalancerListenerRule) Refresh() error {
	return nil
}

func (self SGlobalLoadbalancerListenerRule) IsEmulated() bool {
	return true
}

func (self SGlobalLoadbalancerListenerRule) GetSysTags() map[string]string {
	return nil
}

func (self SGlobalLoadbalancerListenerRule) GetTags() (map[string]string, error) {
	return nil, nil
}

func (self SGlobalLoadbalancerListenerRule) SetTags(tags map[string]string, replace bool) error {
	return cloudprovider.ErrNotSupported
}

func (self SGlobalLoadbalancerListenerRule) GetRedirect() string {
	return ""
}

func (self SGlobalLoadbalancerListenerRule) GetRedirectCode() int64 {
	return 0
}

func (self SGlobalLoadbalancerListenerRule) GetRedirectScheme() string {
	return ""
}

func (self SGlobalLoadbalancerListenerRule) GetRedirectHost() string {
	return ""
}

func (self SGlobalLoadbalancerListenerRule) GetRedirectPath() string {
	return ""
}

func (self SGlobalLoadbalancerListenerRule) IsDefault() bool {
	return false
}

func (self SGlobalLoadbalancerListenerRule) GetDomain() string {
	return self.Domain
}

func (self SGlobalLoadbalancerListenerRule) GetPath() string {
	return self.Path
}

func (self SGlobalLoadbalancerListenerRule) GetCondition() string {
	return ""
}

func (self SGlobalLoadbalancerListenerRule) GetBackendGroupId() string {
	return self.backendService.GetGlobalId()
}

func (self SGlobalLoadbalancerListenerRule) Delete(ctx context.Context) error {
	return cloudprovider.ErrNotSupported
}

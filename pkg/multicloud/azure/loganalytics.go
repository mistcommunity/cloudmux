
package azure

import (
	"fmt"
	"net/url"
	"time"

	"yunion.io/x/pkg/util/httputils"
)

type SLoganalyticsWorkspace struct {
	Id                               string                           `json:"id"`
	Location                         string                           `json:"location"`
	Name                             string                           `json:"name"`
	SLoganalyticsWorkspaceProperties SLoganalyticsWorkspaceProperties `json:"properties"`
	Type                             string                           `json:"type"`
}

type SLoganalyticsWorkspaceProperties struct {
	CreatedDate                                      string                                           `json:"createdDate"`
	CustomerId                                       string                                           `json:"customerId"`
	SLoganalyticsWorkspacePropertiesFeatures         SLoganalyticsWorkspacePropertiesFeatures         `json:"features"`
	ModifiedDate                                     string                                           `json:"modifiedDate"`
	ProvisioningState                                string                                           `json:"provisioningState"`
	PublicNetworkAccessForIngestion                  string                                           `json:"publicNetworkAccessForIngestion"`
	PublicNetworkAccessForQuery                      string                                           `json:"publicNetworkAccessForQuery"`
	RetentionInDays                                  int                                              `json:"retentionInDays"`
	SLoganalyticsWorkspacePropertiesSku              SLoganalyticsWorkspacePropertiesSku              `json:"sku"`
	Source                                           string                                           `json:"source"`
	SLoganalyticsWorkspacePropertiesWorkspaceCapping SLoganalyticsWorkspacePropertiesWorkspaceCapping `json:"workspaceCapping"`
}

type SLoganalyticsWorkspacePropertiesWorkspaceCapping struct {
	DailyQuotaGb        int    `json:"dailyQuotaGb"`
	DataIngestionStatus string `json:"dataIngestionStatus"`
	QuotaNextResetTime  string `json:"quotaNextResetTime"`
}

type SLoganalyticsWorkspacePropertiesSku struct {
	LastSkuUpdate string `json:"lastSkuUpdate"`
	Name          string `json:"name"`
}

type SLoganalyticsWorkspacePropertiesFeatures struct {
	EnableLogAccessUsingOnlyResourcePermissions bool `json:"enableLogAccessUsingOnlyResourcePermissions"`
	Legacy                                      int  `json:"legacy"`
	SearchVersion                               int  `json:"searchVersion"`
}

func (self *SAzureClient) GetLoganalyticsWorkspaces() ([]SLoganalyticsWorkspace, error) {
	if len(self.workspaces) > 0 {
		return self.workspaces, nil
	}
	self.workspaces = []SLoganalyticsWorkspace{}
	return self.workspaces, self.list("Microsoft.OperationalInsights/workspaces", url.Values{}, &self.workspaces)
}

type WorkspaceData struct {
	Name    string
	Columns []struct {
		Name string
		Type string
	}
	Rows [][]string
}

func (self *SAzureClient) GetInstanceDiskUsage(workspace string, instanceId string, start, end time.Time) ([]WorkspaceData, error) {
	params := url.Values{}
	params.Set("timespan", "P1D")
	query := fmt.Sprintf(`Perf | where ObjectName == "LogicalDisk" or ObjectName == "Logical Disk" | where _ResourceId == "%s" | where InstanceName != "_Total" | where CounterName == "%% Free Space" | where TimeGenerated between(datetime("%s") .. datetime("%s")) |  project TimeGenerated, InstanceName, CounterValue, Computer, _ResourceId`, instanceId, start.Format(time.RFC3339), end.Format(time.RFC3339))
	params.Set("query", query)
	resp, err := self.ljsonRequest(string(httputils.GET), fmt.Sprintf("v1/workspaces/%s/query", workspace), nil, params)
	if err != nil {
		return nil, err
	}
	ret := []WorkspaceData{}
	return ret, resp.Unmarshal(&ret, "tables")
}

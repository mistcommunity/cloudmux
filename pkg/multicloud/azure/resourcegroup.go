
package azure

import (
	"fmt"
	"net/url"
	"strings"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type GroupProperties struct {
	ProvisioningState string
}

type SResourceGroup struct {
	multicloud.SProjectBase
	AzureTags
	client *SAzureClient

	ID         string
	Name       string
	Location   string
	Properties GroupProperties
	ManagedBy  string
	subId      string
}

func (self *SRegion) GetResourceGroupDetail(groupName string) (*SResourceGroup, error) {
	resourceGroup := SResourceGroup{client: self.client, subId: self.client._subscriptionId()}
	idStr := fmt.Sprintf("subscriptions/%s/resourcegroups/%s", self.client._subscriptionId(), groupName)
	return &resourceGroup, self.get(idStr, url.Values{}, &resourceGroup)
}

// not support update, resource group name is immutable???
func (self *SRegion) UpdateResourceGroup(groupName string, newName string) error {
	resourceGroup := SResourceGroup{Name: newName, client: self.client, subId: self.client._subscriptionId()}
	resource := fmt.Sprintf("subscriptions/%s/resourcegroups/%s", self.client.subscriptionId, groupName)
	_, err := self.client.patch(resource, jsonutils.Marshal(&resourceGroup))
	return err
}

func (self *SRegion) CreateResourceGroup(groupName string) (jsonutils.JSONObject, error) {
	resourceGroup := SResourceGroup{Location: self.Name, client: self.client, subId: self.client._subscriptionId()}
	idStr := fmt.Sprintf("subscriptions/%s/resourcegroups/%s", self.client._subscriptionId(), groupName)
	return self.client.put(idStr, jsonutils.Marshal(resourceGroup))
}

func (self *SRegion) DeleteResourceGroup(groupName string) error {
	idStr := fmt.Sprintf("subscriptions/%s/resourcegroups/%s", self.client._subscriptionId(), groupName)
	return self.del(idStr)
}

func (r *SResourceGroup) GetName() string {
	return r.Name
}

func (r *SResourceGroup) GetId() string {
	return r.ID
}

func (self *SResourceGroup) GetAccountId() string {
	return fmt.Sprintf("%s/%s", self.client.tenantId, self.subId)
}

func (r *SResourceGroup) GetGlobalId() string {
	return strings.ToLower(fmt.Sprintf("%s/%s", r.subId, r.Name))
}

func (r *SResourceGroup) GetStatus() string {
	return api.EXTERNAL_PROJECT_STATUS_AVAILABLE
}


package azure

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
)

type SSubscription struct {
	SubscriptionId string
	State          string
	DisplayName    string
}

func (self *SSubscription) GetHealthStatus() string {
	if self.State == "Enabled" {
		return api.CLOUD_PROVIDER_HEALTH_NORMAL
	}
	return api.CLOUD_PROVIDER_HEALTH_SUSPENDED
}

func (self *SAzureClient) ListSubscriptions() ([]SSubscription, error) {
	resp, err := self.list_v2("subscriptions", "2014-02-26", nil)
	if err != nil {
		return nil, err
	}
	result := []SSubscription{}
	err = resp.Unmarshal(&result, "value")
	if err != nil {
		return nil, err
	}
	return result, nil
}

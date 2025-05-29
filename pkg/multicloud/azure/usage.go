
package azure

import (
	"fmt"
	"net/url"
	"strings"

	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type UsageName struct {
	Value          string
	LocalizedValue string
}

// {"value":[{"unit":"Count","currentValue":0,"limit":250,"name":{"value":"StorageAccounts","localizedValue":"Storage Accounts"}}]}
type SUsage struct {
	Unit         string
	CurrentValue int
	Limit        int
	Name         UsageName
}

func (u *SUsage) GetGlobalId() string {
	return strings.ToLower(u.Name.Value)
}

func (u *SUsage) GetQuotaType() string {
	return u.Name.Value
}

func (u *SUsage) GetDesc() string {
	return u.Name.LocalizedValue
}

func (u *SUsage) GetMaxQuotaCount() int {
	return u.Limit
}

func (u *SUsage) GetCurrentQuotaUsedCount() int {
	return u.CurrentValue
}

func (region *SRegion) GetUsage(resourceType string) ([]SUsage, error) {
	usage := []SUsage{}
	resource := fmt.Sprintf("%s/locations/%s/usages", resourceType, region.Name)
	err := region.client.list(resource, url.Values{}, &usage)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAll(%s)", resource)
	}
	return usage, nil
}

func (region *SRegion) GetICloudQuotas() ([]cloudprovider.ICloudQuota, error) {
	ret := []cloudprovider.ICloudQuota{}
	for _, resourceType := range []string{"Microsoft.Network", "Microsoft.Storage", "Microsoft.Compute"} {
		usages, err := region.GetUsage(resourceType)
		if err != nil {
			return nil, errors.Wrap(err, "GetUsage")
		}
		for i := range usages {
			ret = append(ret, &usages[i])
		}
	}
	return ret, nil
}

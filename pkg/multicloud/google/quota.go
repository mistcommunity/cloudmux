
package google

import (
	"strings"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SQuota struct {
	Metric string
	Limit  int
	Usage  int
	Owner  string
}

func (q *SQuota) GetGlobalId() string {
	return strings.ToLower(q.Metric)
}

func (q *SQuota) GetName() string {
	return q.Metric
}

func (q *SQuota) GetDesc() string {
	return q.Metric
}

func (q *SQuota) GetQuotaType() string {
	return q.Metric
}

func (q *SQuota) GetMaxQuotaCount() int {
	return q.Limit
}

func (q *SQuota) GetCurrentQuotaUsedCount() int {
	return q.Usage
}

func (region *SRegion) GetICloudQuotas() ([]cloudprovider.ICloudQuota, error) {
	ret := []cloudprovider.ICloudQuota{}
	for i := range region.Quotas {
		ret = append(ret, &region.Quotas[i])
	}
	return ret, nil
}

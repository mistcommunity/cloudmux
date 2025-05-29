
package aliyun

import (
	"time"

	api "yunion.io/x/cloudmux/pkg/apis/billing"
)

func convertChargeType(ct TChargeType) string {
	switch ct {
	case PrePaidInstanceChargeType, PrePaidDBInstanceChargeType:
		return api.BILLING_TYPE_PREPAID
	case PostPaidInstanceChargeType, PostPaidDBInstanceChargeType:
		return api.BILLING_TYPE_POSTPAID
	default:
		return ""
	}
}

func convertExpiredAt(expired time.Time) time.Time {
	if !expired.IsZero() {
		now := time.Now()
		if expired.Sub(now) < time.Hour*24*365*6 {
			return expired
		}
	}
	return time.Time{}
}

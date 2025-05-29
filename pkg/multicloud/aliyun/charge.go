package aliyun

import (
	"time"
)

func convertChargeType(ct TChargeType) string {
	return ""
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


package multicloud

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
)

type SDisk struct {
	SVirtualResourceBase
	SBillingBase
}

func (self *SDisk) GetIStorageId() string {
	return ""
}

func (self *SDisk) GetIops() int {
	return 0
}

func (self *SDisk) GetPreallocation() string {
	return api.DISK_PREALLOCATION_OFF
}

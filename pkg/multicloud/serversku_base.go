
package multicloud

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
)

type SServerSku struct {
	SResourceBase
}

func (self *SServerSku) GetStatus() string {
	return api.SkuStatusReady
}

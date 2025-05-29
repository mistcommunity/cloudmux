
package multicloud

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/errors"
)

type SNasBase struct {
	SVirtualResourceBase
	SBillingBase
}

func (self *SNasBase) SetQuota(input *cloudprovider.SFileSystemSetQuotaInput) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "SetQuota")
}

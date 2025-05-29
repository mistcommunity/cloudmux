
package multicloud

import (
	"time"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/billing"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SBillingBase struct{}

func (self *SBillingBase) GetBillingType() string {
	return ""
}

func (self *SBillingBase) GetExpiredAt() time.Time {
	return time.Time{}
}

func (self *SBillingBase) SetAutoRenew(bc billing.SBillingCycle) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "SetAutoRenew")
}

func (self *SBillingBase) IsAutoRenew() bool {
	return false
}

func (self *SBillingBase) Renew(bc billing.SBillingCycle) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "Renew")
}

func (self *SBillingBase) ChangeBillingType(billType string) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "ChangeBillingType")
}

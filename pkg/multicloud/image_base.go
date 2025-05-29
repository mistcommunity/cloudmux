
package multicloud

import (
	"github.com/pkg/errors"
	"yunion.io/x/pkg/util/rbacscope"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SImageBase struct {
	SVirtualResourceBase
}

func (self *SImageBase) GetPublicScope() rbacscope.TRbacScope {
	return rbacscope.ScopeSystem
}

func (self *SImageBase) GetSubImages() []cloudprovider.SSubImage {
	return nil
}

func (self *SImageBase) Export(opts *cloudprovider.SImageExportOptions) ([]cloudprovider.SImageExportInfo, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "Export")
}

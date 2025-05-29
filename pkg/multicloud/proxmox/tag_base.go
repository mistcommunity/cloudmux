
package proxmox

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type ProxmoxTags struct {
}

func (self *ProxmoxTags) GetTags() (map[string]string, error) {
	tags := map[string]string{}
	return tags, nil
}

func (self *ProxmoxTags) GetSysTags() map[string]string {
	return nil
}

func (self *ProxmoxTags) SetTags(tags map[string]string, replace bool) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetTags")
}

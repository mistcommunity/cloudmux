
package azure

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type AzureTags struct {
	Tags map[string]string
}

func (self *AzureTags) GetTags() (map[string]string, error) {
	return self.Tags, nil
}

func (self *AzureTags) GetSysTags() map[string]string {
	return nil
}

func (self *AzureTags) SetTags(tags map[string]string, replace bool) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetTags")
}

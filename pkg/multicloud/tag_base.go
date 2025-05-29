
package multicloud

import (
	"strings"

	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type STagBase struct {
}

func (self STagBase) GetSysTags() map[string]string {
	return nil
}

func (self STagBase) GetTags() (map[string]string, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetTags")
}

func (self STagBase) SetTags(tags map[string]string, replace bool) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetTags")
}

type STag struct {
	TagKey   string
	TagValue string

	Key   string
	Value string
}

func (self STag) IsSysTagPrefix(keys []string) bool {
	for _, prefix := range keys {
		if strings.HasPrefix(self.TagKey, prefix) {
			return true
		}
		if strings.HasPrefix(self.Key, prefix) {
			return true
		}
	}
	return false
}

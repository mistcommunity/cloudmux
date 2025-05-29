
package google

import (
	"strings"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/encode"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type GoogleTags struct {
	Labels map[string]string
}

func (self *GoogleTags) GetTags() (map[string]string, error) {
	ret := map[string]string{}
	for k, v := range self.Labels {
		if strings.HasPrefix(k, "goog-") {
			continue
		}
		ret[encode.DecodeGoogleLable(k)] = encode.DecodeGoogleLable(v)
	}
	return ret, nil
}

func (self *GoogleTags) GetSysTags() map[string]string {
	ret := map[string]string{}
	for k, v := range self.Labels {
		if strings.HasPrefix(k, "goog-") {
			ret[k] = v
		}
	}
	return ret
}

func (self *GoogleTags) SetTags(tags map[string]string, replace bool) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetTags")
}

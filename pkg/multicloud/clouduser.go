
package multicloud

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SBaseClouduser struct {
}

func (self *SBaseClouduser) CreateAccessKey(name string) (*cloudprovider.SAccessKey, error) {
	return nil, errors.Wrap(cloudprovider.ErrNotImplemented, "base CreateAccessKey")
}

func (self *SBaseClouduser) GetAccessKeys() ([]cloudprovider.SAccessKey, error) {
	return nil, errors.Wrap(cloudprovider.ErrNotImplemented, "base GetAccessKeys")
}

func (self *SBaseClouduser) DeleteAccessKey(accesskey string) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "base DeleteAccessKey")
}

func (self *SBaseClouduser) SetDisable() error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetDisable")
}

func (self *SBaseClouduser) SetEnable(opts *cloudprovider.SClouduserEnableOptions) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetDisable")
}

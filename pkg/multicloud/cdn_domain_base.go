
package multicloud

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/errors"
)

type SCDNDomainBase struct {
	SResourceBase
}

func (self *SCDNDomainBase) GetCacheKeys() (*cloudprovider.SCDNCacheKeys, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetCacheKeys")
}

func (self *SCDNDomainBase) GetRangeOriginPull() (*cloudprovider.SCDNRangeOriginPull, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetRangeOriginPull")
}

func (self *SCDNDomainBase) GetCache() (*cloudprovider.SCDNCache, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetCache")
}

func (self *SCDNDomainBase) GetHTTPS() (*cloudprovider.SCDNHttps, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetHTTPS")
}

func (self *SCDNDomainBase) GetForceRedirect() (*cloudprovider.SCDNForceRedirect, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetForceRedirect")
}

func (self *SCDNDomainBase) GetReferer() (*cloudprovider.SCDNReferer, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetReferer")
}

func (self *SCDNDomainBase) GetMaxAge() (*cloudprovider.SCDNMaxAge, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetMaxAge")
}

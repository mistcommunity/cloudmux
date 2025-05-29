
package multicloud

import (
	"yunion.io/x/cloudmux/pkg/apis"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
)

type SHostBase struct {
	SResourceBase
	STagBase
}

func (host *SHostBase) GetCpuCmtbound() float32 {
	return 0.0
}

func (host *SHostBase) GetMemCmtbound() float32 {
	return 0.0
}

func (host *SHostBase) GetReservedMemoryMb() int {
	return 0
}

func (host *SHostBase) GetSchedtags() ([]string, error) {
	return nil, nil
}

func (host *SHostBase) GetOvnVersion() string {
	return ""
}

func (host *SHostBase) GetCpuArchitecture() string {
	return apis.OS_ARCH_X86_64
}

func (host *SHostBase) GetStorageDriver() string {
	return ""
}

func (host *SHostBase) GetStorageInfo() jsonutils.JSONObject {
	return nil
}

func (host *SHostBase) GetIsolateDevices() ([]cloudprovider.IsolateDevice, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIsolateDevices")
}

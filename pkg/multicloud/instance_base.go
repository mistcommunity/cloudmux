
package multicloud

import (
	"context"

	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SInstanceBase struct {
	SResourceBase
	SBillingBase
}

func (instance *SInstanceBase) GetIHostId() string {
	return ""
}

func (instance *SInstanceBase) GetCpuSockets() int {
	return 1
}

func (instance *SInstanceBase) GetSerialOutput(port int) (string, error) {
	return "", cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) ConvertPublicIpToEip() error {
	return cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) MigrateVM(hostId string) error {
	return cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) LiveMigrateVM(hostId string) error {
	return cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) GetInstanceSnapshot(idStr string) (cloudprovider.ICloudInstanceSnapshot, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) GetInstanceSnapshots() ([]cloudprovider.ICloudInstanceSnapshot, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) CreateInstanceSnapshot(ctx context.Context, name string, desc string) (cloudprovider.ICloudInstanceSnapshot, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (instance *SInstanceBase) ResetToInstanceSnapshot(ctx context.Context, idStr string) error {
	return cloudprovider.ErrNotImplemented
}

func (self *SInstanceBase) SaveImage(opts *cloudprovider.SaveImageOptions) (cloudprovider.ICloudImage, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "SaveImage")
}

func (self *SInstanceBase) AllocatePublicIpAddress() (string, error) {
	return "", errors.Wrapf(cloudprovider.ErrNotImplemented, "AllocatePublicIpAddress")
}

func (self *SInstanceBase) CreateDisk(ctx context.Context, opts *cloudprovider.GuestDiskCreateOptions) (string, error) {
	return "", errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateDisk")
}

func (self *SInstanceBase) GetThroughput() int {
	return 0
}

func (self *SInstanceBase) GetInternetMaxBandwidthOut() int {
	return 0
}

func (ins *SInstanceBase) GetPowerStates() string {
	return ""
}

func (instance *SInstanceBase) GetError() error {
	return nil
}

func (instance *SInstanceBase) GetIsolateDeviceIds() ([]string, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIsolateDeviceIds")
}

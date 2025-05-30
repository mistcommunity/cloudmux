package aliyun

import (
	"fmt"
	"strings"

	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/utils"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type TChargeType string

var LOCAL_STORAGES = []string{
	api.STORAGE_LOCAL_HDD_PRO,
	api.STORAGE_LOCAL_SSD_PRO,
}

const (
	PrePaidInstanceChargeType  TChargeType = "PrePaid"
	PostPaidInstanceChargeType TChargeType = "PostPaid"
	DefaultInstanceChargeType              = PostPaidInstanceChargeType
)

type SpotStrategyType string

const (
	NoSpotStrategy             SpotStrategyType = "NoSpot"
	SpotWithPriceLimitStrategy SpotStrategyType = "SpotWithPriceLimit"
	SpotAsPriceGoStrategy      SpotStrategyType = "SpotAsPriceGo"
	DefaultSpotStrategy                         = NoSpotStrategy
)

type SDedicatedHostGenerations struct {
	DedicatedHostGeneration []string
}

type SVolumeCategories struct {
	VolumeCategories []string
}

type SSupportedDataDiskCategories struct {
	SupportedDataDiskCategory []string
}

type SSupportedInstanceGenerations struct {
	SupportedInstanceGeneration []string
}

type SSupportedInstanceTypeFamilies struct {
	SupportedInstanceTypeFamily []string
}

type SSupportedInstanceTypes struct {
	SupportedInstanceType []string
}

type SSupportedNetworkTypes struct {
	SupportedNetworkCategory []string
}

type SSupportedSystemDiskCategories struct {
	SupportedSystemDiskCategory []string
}

type SResourcesInfo struct {
	DataDiskCategories   SSupportedDataDiskCategories
	InstanceGenerations  SSupportedInstanceGenerations
	InstanceTypeFamilies SSupportedInstanceTypeFamilies
	InstanceTypes        SSupportedInstanceTypes
	IoOptimized          bool
	NetworkTypes         SSupportedNetworkTypes
	SystemDiskCategories SSupportedSystemDiskCategories
}

type SResources struct {
	ResourcesInfo []SResourcesInfo
}

type SResourceCreation struct {
	ResourceTypes []string
}

type SInstanceTypes struct {
	InstanceTypes []string
}

type SDiskCategories struct {
	DiskCategories []string
}

type SDedicatedHostTypes struct {
	DedicatedHostType []string
}

type SZone struct {
	multicloud.SResourceBase
	AliyunTags
	region *SRegion

	iwires []cloudprovider.ICloudWire

	host *SHost

	istorages []cloudprovider.ICloudStorage

	ZoneId                    string
	LocalName                 string
	DedicatedHostGenerations  SDedicatedHostGenerations
	AvailableVolumeCategories SVolumeCategories
	/* 可供创建的具体资源，AvailableResourcesType 组成的数组 */
	AvailableResources SResources
	/* 允许创建的资源类型集合 */
	AvailableResourceCreation SResourceCreation
	/* 允许创建的实例规格类型 */
	AvailableInstanceTypes SInstanceTypes
	/* 支持的磁盘种类集合 */
	AvailableDiskCategories     SDiskCategories
	AvailableDedicatedHostTypes SDedicatedHostTypes
}

func (self *SZone) GetId() string {
	return self.ZoneId
}

func (self *SZone) GetName() string {
	if self.region.GetCloudEnv() == ALIYUN_FINANCE_CLOUDENV && !strings.Contains(self.LocalName, "金融") {
		i := strings.Index(self.LocalName, "可用")
		var localname string
		if i >= 0 {
			localname = self.LocalName[0:i] + "金融云 " + self.LocalName[i:]
		} else {
			localname = self.LocalName + " 金融云"
		}

		return fmt.Sprintf("%s %s", CLOUD_PROVIDER_ALIYUN_CN, localname)
	} else {
		return fmt.Sprintf("%s %s", CLOUD_PROVIDER_ALIYUN_CN, self.LocalName)
	}
}

func (self *SZone) GetI18n() cloudprovider.SModelI18nTable {
	table := cloudprovider.SModelI18nTable{}
	table["name"] = cloudprovider.NewSModelI18nEntry(self.GetName()).CN(self.GetName())
	return table
}

func (self *SZone) GetGlobalId() string {
	return fmt.Sprintf("%s/%s", self.region.GetGlobalId(), self.ZoneId)
}

func (self *SZone) IsEmulated() bool {
	return false
}

func (self *SZone) GetStatus() string {
	if len(self.AvailableResourceCreation.ResourceTypes) == 0 || !utils.IsInStringArray("Instance", self.AvailableResourceCreation.ResourceTypes) {
		return api.ZONE_SOLDOUT
	} else {
		return api.ZONE_ENABLE
	}
}

func (self *SZone) Refresh() error {
	// do nothing
	return nil
}

func (self *SZone) GetIRegion() cloudprovider.ICloudRegion {
	return self.region
}

func (self *SZone) fetchStorages() error {
	categories := self.AvailableDiskCategories.DiskCategories
	self.istorages = []cloudprovider.ICloudStorage{}
	for _, sc := range categories {
		storage := SStorage{zone: self, storageType: sc}
		self.istorages = append(self.istorages, &storage)
		if sc == api.STORAGE_CLOUD_ESSD {
			storage_l0 := SStorage{zone: self, storageType: api.STORAGE_CLOUD_ESSD_PL0}
			self.istorages = append(self.istorages, &storage_l0)
			storage_l2 := SStorage{zone: self, storageType: api.STORAGE_CLOUD_ESSD_PL2}
			self.istorages = append(self.istorages, &storage_l2)
			storage_l3 := SStorage{zone: self, storageType: api.STORAGE_CLOUD_ESSD_PL3}
			self.istorages = append(self.istorages, &storage_l3)
		}
	}
	for _, localStorage := range LOCAL_STORAGES {
		storage := SStorage{zone: self, storageType: localStorage}
		self.istorages = append(self.istorages, &storage)
	}
	return nil
}

func (self *SZone) getStorageByCategory(category string) (*SStorage, error) {
	storages, err := self.GetIStorages()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(storages); i += 1 {
		storage := storages[i].(*SStorage)
		if storage.storageType == category {
			return storage, nil
		}
	}
	return nil, fmt.Errorf("No such storage %s", category)
}

func (self *SZone) GetIStorages() ([]cloudprovider.ICloudStorage, error) {
	if self.istorages == nil {
		err := self.fetchStorages()
		if err != nil {
			return nil, errors.Wrapf(err, "fetchStorages")
		}
	}
	return self.istorages, nil
}

func (self *SZone) GetIStorageById(id string) (cloudprovider.ICloudStorage, error) {
	if self.istorages == nil {
		err := self.fetchStorages()
		if err != nil {
			return nil, errors.Wrapf(err, "fetchStorages")
		}
	}
	for i := 0; i < len(self.istorages); i += 1 {
		if self.istorages[i].GetGlobalId() == id {
			return self.istorages[i], nil
		}
	}
	return nil, cloudprovider.ErrNotFound
}

func (self *SZone) getHost() *SHost {
	if self.host == nil {
		self.host = &SHost{zone: self}
	}
	return self.host
}

func (self *SZone) GetIHosts() ([]cloudprovider.ICloudHost, error) {
	return []cloudprovider.ICloudHost{self.getHost()}, nil
}

func (self *SZone) GetIHostById(id string) (cloudprovider.ICloudHost, error) {
	host := self.getHost()
	if host.GetGlobalId() == id {
		return host, nil
	}
	return nil, cloudprovider.ErrNotFound
}

func (self *SZone) addWire(wire *SWire) {
	if self.iwires == nil {
		self.iwires = make([]cloudprovider.ICloudWire, 0)
	}
	self.iwires = append(self.iwires, wire)
}

func (self *SZone) GetIWires() ([]cloudprovider.ICloudWire, error) {
	return self.iwires, nil
}

func (self *SZone) getNetworkById(vswitchId string) *SVSwitch {
	log.Debugf("Search in wires %d", len(self.iwires))
	for i := 0; i < len(self.iwires); i += 1 {
		log.Debugf("Search in wire %s", self.iwires[i].GetName())
		wire := self.iwires[i].(*SWire)
		net := wire.getNetworkById(vswitchId)
		if net != nil {
			return net
		}
	}
	return nil
}

func (self *SZone) getSysDiskCategories() []string {
	ret := []string{}
	for _, res := range self.AvailableResources.ResourcesInfo {
		for _, category := range res.SystemDiskCategories.SupportedSystemDiskCategory {
			if !utils.IsInStringArray(category, ret) {
				ret = append(ret, category)
			}
		}
	}
	if utils.IsInStringArray(api.STORAGE_CLOUD_ESSD, ret) {
		ret = append(ret, []string{api.STORAGE_CLOUD_ESSD_PL2, api.STORAGE_CLOUD_ESSD_PL3, api.STORAGE_CLOUD_ESSD_PL0}...)
	}
	return ret
}

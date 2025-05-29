
package nutanix

import (
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SZone struct {
	multicloud.STagBase
	multicloud.SResourceBase

	SCluster
	region *SRegion
}

func (self *SZone) GetName() string {
	return self.Name
}

func (self *SZone) GetId() string {
	return self.UUID
}

func (self *SZone) GetGlobalId() string {
	return self.UUID
}

func (self *SZone) GetI18n() cloudprovider.SModelI18nTable {
	table := cloudprovider.SModelI18nTable{}
	table["name"] = cloudprovider.NewSModelI18nEntry(self.GetName()).CN(self.GetName())
	return table
}

func (self *SZone) GetIHosts() ([]cloudprovider.ICloudHost, error) {
	hosts, err := self.region.GetHosts()
	if err != nil {
		return nil, errors.Wrapf(err, "GetIHosts")
	}
	firstHost := true
	ret := []cloudprovider.ICloudHost{}
	for i := range hosts {
		if hosts[i].ClusterUUID != self.UUID {
			continue
		}
		hosts[i].zone = self
		hosts[i].firstHost = firstHost
		ret = append(ret, &hosts[i])
		if firstHost {
			firstHost = false
		}
	}
	return ret, nil
}

func (self *SZone) GetIHostById(id string) (cloudprovider.ICloudHost, error) {
	host, err := self.region.GetHost(id)
	if err != nil {
		return nil, errors.Wrapf(err, "GetIHostById(%s)", id)
	}
	if host.ClusterUUID != self.UUID {
		return nil, errors.Wrapf(cloudprovider.ErrNotFound, id)
	}
	host.zone = self
	return host, nil
}

func (self *SZone) GetIRegion() cloudprovider.ICloudRegion {
	return self.region
}

func (self *SZone) GetStatus() string {
	return api.ZONE_ENABLE
}

func (self *SZone) GetIStorages() ([]cloudprovider.ICloudStorage, error) {
	storages, err := self.region.GetStorages()
	if err != nil {
		return nil, errors.Wrapf(err, "GetStorages")
	}
	ret := []cloudprovider.ICloudStorage{}
	for i := range storages {
		if storages[i].ClusterUUID != self.UUID {
			continue
		}
		storages[i].zone = self
		ret = append(ret, &storages[i])
	}
	return ret, nil
}

func (self *SZone) GetIStorageById(id string) (cloudprovider.ICloudStorage, error) {
	storage, err := self.region.GetStorage(id)
	if err != nil {
		return nil, errors.Wrapf(err, "GetStorage %s", id)
	}
	if storage.ClusterUUID != self.UUID {
		return nil, errors.Wrapf(cloudprovider.ErrNotFound, id)
	}
	storage.zone = self
	return storage, nil
}


package esxi

import (
	"context"
	"path"
	"strings"
	"time"

	"github.com/vmware/govmomi/object"

	"yunion.io/x/pkg/util/imagetools"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SImage struct {
	multicloud.SImageBase
	multicloud.STagBase
	cache    *SDatastoreImageCache
	filename string
	size     int64
	createAt time.Time

	imageInfo *imagetools.ImageInfo
}

func (self *SImage) GetMinRamSizeMb() int {
	return 0
}

func (self *SImage) getDatacenter() *object.Datacenter {
	return self.cache.datastore.datacenter.getDcObj()
}

func (self *SImage) getFullFilename() string {
	return self.cache.datastore.getPathString(self.filename)
}

func (self *SImage) GetId() string {
	idstr := path.Base(self.filename)
	if strings.HasSuffix(idstr, ".vmdk") {
		idstr = idstr[:len(idstr)-5]
	}
	return strings.ToLower(idstr)
}

func (self *SImage) GetName() string {
	return self.GetId()
}

func (self *SImage) GetGlobalId() string {
	return self.GetId()
}

func (self *SImage) GetStatus() string {
	dm := object.NewVirtualDiskManager(self.cache.datastore.manager.client.Client)
	ctx := context.Background()
	_, err := dm.QueryVirtualDiskInfo(ctx, self.getFullFilename(), self.getDatacenter(), true)
	if err != nil {
		return api.CACHED_IMAGE_STATUS_CACHE_FAILED
	}
	return api.CACHED_IMAGE_STATUS_ACTIVE
}

func (self *SImage) GetImageStatus() string {
	status := self.GetStatus()
	if status == api.CACHED_IMAGE_STATUS_ACTIVE {
		return cloudprovider.IMAGE_STATUS_ACTIVE
	}
	return cloudprovider.IMAGE_STATUS_DELETED
}

func (self *SImage) Refresh() error {
	return nil
}

func (self *SImage) IsEmulated() bool {
	return false
}

func (self *SImage) Delete(ctx context.Context) error {
	return self.cache.datastore.DeleteVmdk(ctx, self.filename)
}

func (self *SImage) GetIStoragecache() cloudprovider.ICloudStoragecache {
	return self.cache
}

func (self *SImage) GetImageType() cloudprovider.TImageType {
	return cloudprovider.ImageTypeCustomized
}

func (self *SImage) GetSizeByte() int64 {
	return self.size
}

func (img *SImage) getNormalizedImageInfo() *imagetools.ImageInfo {
	if img.imageInfo == nil {
		imgInfo := imagetools.NormalizeImageInfo(img.filename, "", "", "", "")
		img.imageInfo = &imgInfo
	}
	return img.imageInfo
}

func (img *SImage) GetFullOsName() string {
	return img.filename
}

func (img *SImage) GetOsType() cloudprovider.TOsType {
	return cloudprovider.TOsType(img.getNormalizedImageInfo().OsType)
}

func (img *SImage) GetOsArch() string {
	return img.getNormalizedImageInfo().OsArch
}

func (img *SImage) GetOsDist() string {
	return img.getNormalizedImageInfo().OsDistro
}

func (img *SImage) GetOsVersion() string {
	return img.getNormalizedImageInfo().OsVersion
}

func (img *SImage) GetOsLang() string {
	return img.getNormalizedImageInfo().OsLang
}

func (img *SImage) GetBios() cloudprovider.TBiosType {
	return cloudprovider.ToBiosType(img.getNormalizedImageInfo().OsBios)
}

func (self *SImage) GetMinOsDiskSizeGb() int {
	return int(self.GetSizeByte() / 1024 / 1024 / 1024)
}

func (self *SImage) GetImageFormat() string {
	return "vmdk"
}

func (self *SImage) GetCreatedAt() time.Time {
	return self.createAt
}


package proxmox

import (
	"context"
	"fmt"
	"net/url"

	"yunion.io/x/pkg/util/imagetools"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SImage struct {
	multicloud.SImageBase
	ProxmoxTags
	cache *SStoragecache

	imageInfo *imagetools.ImageInfo

	Volid   string
	Size    int64
	Ctime   int64
	Content string
	Format  string
}

func (self *SImage) GetMinRamSizeMb() int {
	return 0
}

func (self *SImage) GetId() string {
	return self.Volid
}

func (self *SImage) GetName() string {
	return self.Volid
}

func (self *SImage) Delete(ctx context.Context) error {
	return cloudprovider.ErrNotImplemented
}

func (self *SImage) GetGlobalId() string {
	return self.GetId()
}

func (self *SImage) GetIStoragecache() cloudprovider.ICloudStoragecache {
	return self.cache
}

func (self *SImage) GetStatus() string {
	return api.CACHED_IMAGE_STATUS_ACTIVE
}

func (self *SImage) GetImageStatus() string {
	return cloudprovider.IMAGE_STATUS_ACTIVE
}

func (self *SImage) GetImageType() cloudprovider.TImageType {
	return cloudprovider.ImageTypeSystem
}

func (self *SImage) GetSizeByte() int64 {
	return self.Size
}

func (img *SImage) getNormalizedImageInfo() *imagetools.ImageInfo {
	if img.imageInfo == nil {
		imgInfo := imagetools.NormalizeImageInfo(img.Volid, "", "", "", "")
		img.imageInfo = &imgInfo
	}
	return img.imageInfo
}

func (img *SImage) GetOsType() cloudprovider.TOsType {
	return cloudprovider.TOsType(img.getNormalizedImageInfo().OsType)
}

func (img *SImage) GetOsDist() string {
	return img.getNormalizedImageInfo().OsDistro
}

func (img *SImage) GetOsVersion() string {
	return img.getNormalizedImageInfo().OsVersion
}

func (img *SImage) GetOsArch() string {
	return img.getNormalizedImageInfo().OsArch
}

func (img *SImage) GetOsLang() string {
	return img.getNormalizedImageInfo().OsLang
}

func (img *SImage) GetFullOsName() string {
	return ""
}

func (img *SImage) GetBios() cloudprovider.TBiosType {
	return cloudprovider.BIOS
}

func (self *SImage) GetMinOsDiskSizeGb() int {
	if self.GetOsType() == "windows" {
		return 40
	}
	return 30
}

func (self *SImage) GetImageFormat() string {
	return self.Format
}

func (self *SProxmoxClient) GetImages(node, storageName string) ([]SImage, error) {
	images := []SImage{}
	params := url.Values{}
	params.Set("content", "iso")
	path := fmt.Sprintf("/nodes/%s/storage/%s/content", node, storageName)
	err := self.get(path, params, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (self *SRegion) GetImages(node, storageName string) ([]SImage, error) {
	return self.client.GetImages(node, storageName)
}


package aliyun

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/imagetools"

	"yunion.io/x/cloudmux/pkg/apis"
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type ImageStatusType string

const (
	ImageStatusCreating     ImageStatusType = "Creating"
	ImageStatusAvailable    ImageStatusType = "Available"
	ImageStatusUnAvailable  ImageStatusType = "UnAvailable"
	ImageStatusCreateFailed ImageStatusType = "CreateFailed"
)

type ImageOwnerType string

const (
	ImageOwnerSystem      ImageOwnerType = "system"
	ImageOwnerSelf        ImageOwnerType = "self"
	ImageOwnerOthers      ImageOwnerType = "others"
	ImageOwnerMarketplace ImageOwnerType = "marketplace"
)

type ImageUsageType string

const (
	ImageUsageInstance ImageUsageType = "instance"
	ImageUsageNone     ImageUsageType = "none"
)

type SImage struct {
	multicloud.SImageBase
	AliyunTags
	storageCache *SStoragecache

	// normalized image info
	imgInfo *imagetools.ImageInfo

	Architecture         string
	CreationTime         time.Time
	Description          string
	ImageId              string
	ImageName            string
	OSName               string
	OSType               string
	ImageOwnerAlias      ImageOwnerType
	IsSupportCloudinit   bool
	IsSupportIoOptimized bool
	Platform             string
	Size                 int
	Status               ImageStatusType
	Usage                string
}

func (self *SImage) GetMinRamSizeMb() int {
	return 0
}

func (self *SImage) GetId() string {
	return self.ImageId
}

func (self *SImage) GetName() string {
	if self.ImageOwnerAlias == ImageOwnerSystem {
		return self.OSName
	} else {
		return self.ImageName
	}
}

func (self *SImage) IsEmulated() bool {
	return false
}

func (self *SImage) Delete(ctx context.Context) error {
	return self.storageCache.region.DeleteImage(self.ImageId)
}

func (self *SImage) GetGlobalId() string {
	return self.ImageId
}

func (self *SImage) GetIStoragecache() cloudprovider.ICloudStoragecache {
	return self.storageCache
}

func (self *SImage) GetStatus() string {
	switch self.Status {
	case ImageStatusCreating:
		return api.CACHED_IMAGE_STATUS_SAVING
	case ImageStatusAvailable:
		return api.CACHED_IMAGE_STATUS_ACTIVE
	case ImageStatusUnAvailable:
		return api.CACHED_IMAGE_STATUS_CACHE_FAILED
	case ImageStatusCreateFailed:
		return api.CACHED_IMAGE_STATUS_CACHE_FAILED
	default:
		return api.CACHED_IMAGE_STATUS_CACHE_FAILED
	}
}

func (self *SImage) GetImageStatus() string {
	switch self.Status {
	case ImageStatusCreating:
		return cloudprovider.IMAGE_STATUS_QUEUED
	case ImageStatusAvailable:
		return cloudprovider.IMAGE_STATUS_ACTIVE
	case ImageStatusUnAvailable:
		return cloudprovider.IMAGE_STATUS_DELETED
	case ImageStatusCreateFailed:
		return cloudprovider.IMAGE_STATUS_KILLED
	default:
		return cloudprovider.IMAGE_STATUS_KILLED
	}
}

func (self *SImage) Refresh() error {
	new, err := self.storageCache.region.GetImage(self.ImageId)
	if err != nil {
		return err
	}
	return jsonutils.Update(self, new)
}

func (self *SImage) GetImageType() cloudprovider.TImageType {
	switch self.ImageOwnerAlias {
	case ImageOwnerSystem:
		return cloudprovider.ImageTypeSystem
	case ImageOwnerSelf, ImageOwnerOthers:
		return cloudprovider.ImageTypeCustomized
	case ImageOwnerMarketplace:
		return cloudprovider.ImageTypeMarket
	default:
		return cloudprovider.ImageTypeCustomized
	}
}

func (self *SImage) GetSizeByte() int64 {
	return int64(self.Size) * 1024 * 1024 * 1024
}

func (self *SImage) GetOsType() cloudprovider.TOsType {
	return cloudprovider.TOsType(self.getNormalizedImageInfo().OsType)
}

func (self *SImage) GetOsDist() string {
	return self.getNormalizedImageInfo().OsDistro
}

func (self *SImage) getNormalizedImageInfo() *imagetools.ImageInfo {
	if self.imgInfo == nil {
		imgInfo := imagetools.NormalizeImageInfo(self.OSName, self.Architecture, self.OSType, self.Platform, "")
		self.imgInfo = &imgInfo
	}

	return self.imgInfo
}

func (self *SImage) GetFullOsName() string {
	return self.OSName
}

func (self *SImage) GetOsVersion() string {
	return self.getNormalizedImageInfo().OsVersion
}

func (self *SImage) GetOsLang() string {
	return self.getNormalizedImageInfo().OsLang
}

func (self *SImage) GetOsArch() string {
	return self.getNormalizedImageInfo().OsArch
}

func (self *SImage) GetBios() cloudprovider.TBiosType {
	return cloudprovider.ToBiosType(self.getNormalizedImageInfo().OsBios)
}

func (self *SImage) GetMinOsDiskSizeGb() int {
	return 40
}

func (self *SImage) GetImageFormat() string {
	return "vhd"
}

func (self *SImage) GetCreatedAt() time.Time {
	return self.CreationTime
}

type ImageExportTask struct {
	ImageId  string
	RegionId string
	// RequestId string
	TaskId string
}

func (self *SImage) Export(opts *cloudprovider.SImageExportOptions) ([]cloudprovider.SImageExportInfo, error) {
	err := self.storageCache.region.GetClient().EnableImageExport()
	if err != nil {
		return nil, err
	}

	if len(opts.BucketName) == 0 {
		opts.BucketName = fmt.Sprintf("image-export-%s", self.ImageId)
		err := self.storageCache.region.CreateIBucket(opts.BucketName, "", "")
		if err != nil {
			return nil, errors.Wrapf(err, "CreateIBucket")
		}
	}
	bucket, err := self.storageCache.region.checkBucket(opts.BucketName)
	if err != nil {
		return nil, errors.Wrapf(err, "GetIBucketByName(%s)", opts.BucketName)
	}
	task, err := self.storageCache.region.ExportImage(self.ImageId, opts.BucketName)
	if err != nil {
		return nil, errors.Wrapf(err, "ExportImage")
	}

	err = self.storageCache.region.WaitTaskStatus(ExportImageTask, task.TaskId, TaskStatusFinished, time.Second*10, time.Minute*30, 0, 100, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "WaitTaskStatus")
	}

	images, err := bucket.ListObjects(oss.Prefix(fmt.Sprintf("%sexport", strings.Replace(self.ImageId, "-", "", -1))))
	if err != nil {
		return nil, errors.Wrap(err, "bucket.ListObjects")
	}
	ret := []cloudprovider.SImageExportInfo{}
	for _, image := range images.Objects {
		url, err := bucket.SignURL(image.Key, oss.HTTPMethod("GET"), 32400)
		if err != nil {
			return nil, errors.Wrapf(err, "SignURL(%s)", image.Key)
		}
		ret = append(ret, cloudprovider.SImageExportInfo{
			DownloadUrl:    url,
			Name:           strings.TrimSuffix(image.Key, ".tar.gz"),
			CompressFormat: "tar.gz",
		})
	}
	return ret, nil
}

func (self *SRegion) ExportImage(imageId, bucketName string) (*ImageExportTask, error) {
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["ImageId"] = imageId
	params["OssBucket"] = bucketName
	params["OssPrefix"] = fmt.Sprintf("%sexport", strings.Replace(imageId, "-", "", -1))

	body, err := self.ecsRequest("ExportImage", params)
	if err != nil {
		return nil, errors.Wrapf(err, "ExportImage")
	}
	result := &ImageExportTask{}
	if err := body.Unmarshal(result); err != nil {
		return nil, errors.Wrapf(err, "Unmarshal")
	}
	return result, nil
}

// {"ImageId":"m-j6c1qlpa7oebbg1n2k60","RegionId":"cn-hongkong","RequestId":"F8B2F6A1-F6AA-4C92-A54C-C4A309CF811F","TaskId":"t-j6c1qlpa7oebbg1rcl9t"}

type ImageImportTask struct {
	ImageId  string
	RegionId string
	// RequestId string
	TaskId string
}

func (self *SRegion) ImportImage(name string, osArch string, osType string, osDist string, bucket string, key string) (*ImageImportTask, error) {
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["ImageName"] = name
	if osDist == "RHEL" {
		osDist = "CentOS"
	}
	params["Platform"] = osDist // "Others Linux"
	params["OSType"] = osType   // "linux"
	switch osArch {
	case apis.OS_ARCH_I386, apis.OS_ARCH_X86_32:
		params["Architecture"] = "i386"
	case apis.OS_ARCH_X86, apis.OS_ARCH_X86_64:
		params["Architecture"] = "x86_64"
	case apis.OS_ARCH_ARM, apis.OS_ARCH_AARCH32, apis.OS_ARCH_AARCH64:
		params["Architecture"] = "arm64"
	default:
		params["Architecture"] = osArch // "x86_64"
	}
	params["DiskDeviceMapping.1.OSSBucket"] = bucket
	params["DiskDeviceMapping.1.OSSObject"] = key

	log.Debugf("Upload image with params %#v", params)

	body, err := self.ecsRequest("ImportImage", params)
	if err != nil {
		log.Errorf("ImportImage fail %s", err)
		return nil, err
	}

	log.Infof("%s", body)
	result := ImageImportTask{}
	err = body.Unmarshal(&result)
	if err != nil {
		log.Errorf("unmarshal result error %s", err)
		return nil, err
	}

	return &result, nil
}

func (self *SRegion) GetImage(imageId string) (*SImage, error) {
	images, _, err := self.GetImages("", "", []string{imageId}, "", 0, 1)
	if err != nil {
		return nil, err
	}
	if len(images) == 0 {
		return nil, cloudprovider.ErrNotFound
	}
	return &images[0], nil
}

func (self *SRegion) GetImageByName(name string) (*SImage, error) {
	images, _, err := self.GetImages("", "", nil, name, 0, 1)
	if err != nil {
		return nil, err
	}
	if len(images) == 0 {
		return nil, cloudprovider.ErrNotFound
	}
	return &images[0], nil
}

func (self *SRegion) GetImagesBySnapshot(snapshotId string, offset int, limit int) ([]SImage, int, error) {
	if limit > 50 || limit <= 0 {
		limit = 50
	}
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["PageSize"] = fmt.Sprintf("%d", limit)
	params["PageNumber"] = fmt.Sprintf("%d", (offset/limit)+1)
	params["SnapshotId"] = snapshotId

	return self.getImages(params)
}

func (self *SRegion) GetImageStatus(imageId string) (ImageStatusType, error) {
	image, err := self.GetImage(imageId)
	if err != nil {
		return "", err
	}
	return image.Status, nil
}

func (self *SRegion) GetImages(status ImageStatusType, owner ImageOwnerType, imageId []string, name string, offset int, limit int) ([]SImage, int, error) {
	if limit > 50 || limit <= 0 {
		limit = 50
	}
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["PageSize"] = fmt.Sprintf("%d", limit)
	params["PageNumber"] = fmt.Sprintf("%d", (offset/limit)+1)

	if len(status) > 0 {
		params["Status"] = string(status)
	} else {
		params["Status"] = "Creating,Available,UnAvailable,CreateFailed"
	}
	if imageId != nil && len(imageId) > 0 {
		params["ImageId"] = strings.Join(imageId, ",")
	}
	if len(owner) > 0 {
		params["ImageOwnerAlias"] = string(owner)
	}

	if len(name) > 0 {
		params["ImageName"] = name
	}

	return self.getImages(params)
}

func (self *SRegion) getImages(params map[string]string) ([]SImage, int, error) {
	body, err := self.ecsRequest("DescribeImages", params)
	if err != nil {
		log.Errorf("DescribeImages fail %s", err)
		return nil, 0, err
	}

	images := make([]SImage, 0)
	err = body.Unmarshal(&images, "Images", "Image")
	if err != nil {
		log.Errorf("unmarshal images fail %s", err)
		return nil, 0, nil
	}
	total, _ := body.Int("TotalCount")
	return images, int(total), nil
}

func (self *SRegion) DeleteImage(imageId string) error {
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["ImageId"] = imageId
	params["Force"] = "true"

	_, err := self.ecsRequest("DeleteImage", params)
	if err != nil {
		log.Errorf("DeleteImage fail %s", err)
		return err
	}
	return nil
}


package aliyun

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/qemuimgfmt"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SStoragecache struct {
	multicloud.SResourceBase
	AliyunTags
	region *SRegion
}

func (self *SStoragecache) GetId() string {
	return fmt.Sprintf("%s-%s", self.region.client.cpcfg.Id, self.region.GetId())
}

func (self *SStoragecache) GetName() string {
	return fmt.Sprintf("%s-%s", self.region.client.cpcfg.Name, self.region.GetId())
}

func (self *SStoragecache) GetStatus() string {
	return "available"
}

func (self *SStoragecache) Refresh() error {
	return nil
}

func (self *SStoragecache) GetGlobalId() string {
	return fmt.Sprintf("%s-%s", self.region.client.cpcfg.Id, self.region.GetGlobalId())
}

func (self *SStoragecache) IsEmulated() bool {
	return false
}

func (self *SStoragecache) GetICloudImages() ([]cloudprovider.ICloudImage, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (self *SStoragecache) GetICustomizedCloudImages() ([]cloudprovider.ICloudImage, error) {
	images := make([]SImage, 0)
	for {
		parts, total, err := self.region.GetImages(ImageStatusType(""), ImageOwnerSelf, nil, "", len(images), 50)
		if err != nil {
			return nil, errors.Wrapf(err, "GetImages")
		}
		images = append(images, parts...)
		if len(images) >= total {
			break
		}
	}
	ret := []cloudprovider.ICloudImage{}
	for i := range images {
		images[i].storageCache = self
		ret = append(ret, &images[i])
	}
	// 共享镜像
	images = []SImage{}
	for {
		parts, total, err := self.region.GetImages(ImageStatusType(""), ImageOwnerOthers, nil, "", len(images), 50)
		if err != nil {
			return nil, errors.Wrapf(err, "GetImages")
		}
		images = append(images, parts...)
		if len(images) >= total {
			break
		}
	}
	for i := range images {
		images[i].storageCache = self
		ret = append(ret, &images[i])
	}

	return ret, nil
}

func (self *SStoragecache) GetIImageById(extId string) (cloudprovider.ICloudImage, error) {
	img, err := self.region.GetImage(extId)
	if err != nil {
		return nil, err
	}
	img.storageCache = self
	return img, nil
}

func (self *SStoragecache) GetPath() string {
	return ""
}

func (self *SStoragecache) UploadImage(ctx context.Context, image *cloudprovider.SImageCreateOption, callback func(float32)) (string, error) {
	if len(image.ExternalId) > 0 {
		status, err := self.region.GetImageStatus(image.ExternalId)
		if err != nil {
			log.Errorf("GetImageStatus error %s", err)
		}
		// 不能直接删除 ImageStatusCreating 状态的image ,需要先取消importImage Task
		if status == ImageStatusCreating {
			err := self.region.CancelImageImportTasks()
			if err != nil {
				log.Errorln(err)
			}
		}
	}
	return self.uploadImage(ctx, image, callback)
}

func (self *SStoragecache) uploadImage(ctx context.Context, image *cloudprovider.SImageCreateOption, callback func(float32)) (string, error) {
	reader, sizeByte, err := image.GetReader(image.ImageId, string(qemuimgfmt.QCOW2))
	if err != nil {
		return "", errors.Wrapf(err, "GetReader")
	}

	bucketName := strings.ToLower(fmt.Sprintf("imgcache-%s-%s", self.region.GetId(), image.ImageId))
	exist, err := self.region.IBucketExist(bucketName)
	if err != nil {
		return "", errors.Wrapf(err, "IBucketExist(%s)", bucketName)
	}
	if !exist {
		log.Debugf("Bucket %s not exists, to create ...", bucketName)
		err = self.region.CreateIBucket(bucketName, "", "")
		if err != nil {
			return "", errors.Wrapf(err, "CreateIBucket %s", bucketName)
		}
	} else {
		log.Debugf("Bucket %s exists", bucketName)
	}

	defer self.region.DeleteIBucket(bucketName) // remove bucket

	bucket, err := self.region.GetIBucketByName(bucketName)
	if err != nil {
		return "", errors.Wrapf(err, "GetIBucketByName %s", bucketName)
	}
	log.Debugf("To upload image to bucket %s ...", bucketName)
	body := multicloud.NewProgress(sizeByte, 80, reader, callback)
	err = cloudprovider.UploadObject(context.Background(), bucket, image.ImageId, 0, body, sizeByte, "", "", nil, false)
	if err != nil {
		return "", errors.Wrapf(err, "UploadObject %s", image.ImageId)
	}

	defer bucket.DeleteObject(context.Background(), image.ImageId) // remove object

	imageBaseName := image.ImageId
	if imageBaseName[0] >= '0' && imageBaseName[0] <= '9' {
		imageBaseName = fmt.Sprintf("img%s", image.ImageId)
	}
	imageName := imageBaseName
	nameIdx := 1

	// check image name, avoid name conflict
	for {
		_, err = self.region.GetImageByName(imageName)
		if err != nil {
			if errors.Cause(err) == cloudprovider.ErrNotFound {
				break
			} else {
				return "", err
			}
		}
		imageName = fmt.Sprintf("%s-%d", imageBaseName, nameIdx)
		nameIdx += 1
	}

	log.Debugf("Import image %s", imageName)

	// ensure privileges
	err = self.region.GetClient().EnableImageImport()
	if err != nil {
		return "", errors.Wrapf(err, "EnableImageImport")
	}

	task, err := self.region.ImportImage(imageName, image.OsArch, image.OsType, image.OsDistribution, bucketName, image.ImageId)

	if err != nil {
		return "", errors.Wrapf(err, "ImportImage %s %s", image.ImageId, bucketName)
	}

	// timeout: 1hour = 3600 seconds
	err = self.region.WaitTaskStatus(ImportImageTask, task.TaskId, TaskStatusFinished, 15*time.Second, 3600*time.Second, 80, 100, callback)
	if err != nil {
		return task.ImageId, errors.Wrapf(err, "waitTaskStatus")
	}

	return task.ImageId, nil
}

func (self *SRegion) CheckBucket(bucketName string) (*oss.Bucket, error) {
	return self.checkBucket(bucketName)
}

func (self *SRegion) checkBucket(bucketName string) (*oss.Bucket, error) {
	oss, err := self.GetOssClient()
	if err != nil {
		log.Errorf("GetOssClient err %s", err)
		return nil, err
	}
	if exist, err := oss.IsBucketExist(bucketName); err != nil {
		log.Errorf("IsBucketExist err %s", err)
		return nil, err
	} else if !exist {
		log.Debugf("Bucket %s not exists, to create ...", bucketName)
		if err := oss.CreateBucket(bucketName); err != nil {
			log.Errorf("Create bucket error %s", err)
			return nil, err
		}
	}
	log.Debugf("Bucket %s exists", bucketName)
	if bucket, err := oss.Bucket(bucketName); err != nil {
		log.Errorf("Bucket error %s %s", bucketName, err)
		return nil, err
	} else {
		return bucket, nil
	}
}

func (self *SRegion) CreateImage(snapshoutId, imageName, imageDesc string) (string, error) {
	return self.createIImage(snapshoutId, imageName, imageDesc)
}

func (self *SRegion) createIImage(snapshoutId, imageName, imageDesc string) (string, error) {
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["OssBucket"] = strings.ToLower(fmt.Sprintf("imgcache-%s", self.GetId()))
	params["SnapshotId"] = snapshoutId
	params["ImageName"] = imageName
	params["Description"] = imageDesc

	if _, err := self.checkBucket(params["OssBucket"]); err != nil {
		return "", err
	}

	if body, err := self.ecsRequest("CreateImage", params); err != nil {
		log.Errorf("CreateImage fail %s", err)
		return "", err
	} else {
		log.Infof("%s", body)
		return body.GetString("ImageId")
	}
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		log.Debugf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		log.Debugf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		log.Debugf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		log.Debugf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

func (region *SRegion) GetIStoragecaches() ([]cloudprovider.ICloudStoragecache, error) {
	storageCache := region.getStoragecache()
	return []cloudprovider.ICloudStoragecache{storageCache}, nil
}

func (region *SRegion) GetIStoragecacheById(id string) (cloudprovider.ICloudStoragecache, error) {
	storageCache := region.getStoragecache()
	if id == storageCache.GetGlobalId() {
		return storageCache, nil
	}
	return nil, cloudprovider.ErrNotFound
}


package aliyun

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type AliyunTags struct {
	Tags struct {
		Tag []multicloud.STag

		// Kafka
		TagVO []multicloud.STag `json:"TagVO" yunion-deprecated-by:"Tag"`
	}
}

var sysTags = []string{
	"aliyun", "creator",
	"acs:", "serverless/", "alloc_id", "virtual-kubelet",
	"diskId", "diskNum", "serverId", "restoreId", "cnfs-id", "from", "shadowId",
	"ack.aliyun.com", "cluster-id.ack.aliyun.com", "ack.alibabacloud.com",
	"k8s.io", "k8s.aliyun.com", "kubernetes.do.not.delete", "kubernetes.reused.by.user",
	"HBR InstanceId", "HBR Retention Days", "HBR Retention Type", "HBR JobId",
	"createdBy", "recoveryPointTime", "recoveryPointId",
	"eas_resource_group_name", "eas_tenant_name", "managedby",
}

func (self *AliyunTags) GetTags() (map[string]string, error) {
	ret := map[string]string{}
	for _, tag := range self.Tags.Tag {
		if tag.IsSysTagPrefix(sysTags) {
			continue
		}
		if len(tag.TagKey) > 0 {
			ret[tag.TagKey] = tag.TagValue
		} else if len(tag.Key) > 0 {
			ret[tag.Key] = tag.Value
		}
	}

	return ret, nil
}

func (self *AliyunTags) GetSysTags() map[string]string {
	ret := map[string]string{}
	for _, tag := range self.Tags.Tag {
		if tag.IsSysTagPrefix(sysTags) {
			if len(tag.TagKey) > 0 {
				ret[tag.TagKey] = tag.TagValue
			} else if len(tag.Key) > 0 {
				ret[tag.Key] = tag.Value
			}
		}
	}
	return ret
}

func (self *AliyunTags) SetTags(tags map[string]string, replace bool) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetTags")
}

type SAliyunTag struct {
	ResourceId   string
	ResourceType string
	TagKey       string
	TagValue     string
}

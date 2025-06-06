
package aliyun

import (
	"fmt"
	"strings"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/utils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func (self *SRegion) tagRequest(serviceType, action string, params map[string]string, body interface{}) (jsonutils.JSONObject, error) {
	switch serviceType {
	case ALIYUN_SERVICE_ECS:
		return self.ecsRequest(action, params)
	case ALIYUN_SERVICE_VPC:
		return self.vpcRequest(action, params)
	case ALIYUN_SERVICE_RDS:
		return self.rdsRequest(action, params)
	case ALIYUN_SERVICE_ES:
		return self.esRequest(action, params, body)
	case ALIYUN_SERVICE_KAFKA:
		return self.kafkaRequest(action, params)
	case ALIYUN_SERVICE_SLB:
		return self.lbRequest(action, params)
	case ALIYUN_SERVICE_KVS:
		return self.kvsRequest(action, params)
	case ALIYUN_SERVICE_NAS:
		return self.nasRequest(action, params)
	case ALIYUN_SERVICE_MONGO_DB:
		return self.mongodbRequest(action, params)
	case ALIYUN_SERVICE_CDN:
		return self.client.cdnRequest(action, params)
	default:
		return nil, fmt.Errorf("invalid service type")
	}
}

// 资源类型。取值范围：
// disk, instance, image, securitygroup, snapshot
func (self *SRegion) ListTags(serviceType string, resourceType string, resourceId string) ([]SAliyunTag, error) {
	tags := []SAliyunTag{}
	params := make(map[string]string)
	params["RegionId"] = self.RegionId
	params["ResourceType"] = resourceType

	if serviceType == ALIYUN_SERVICE_ES {
		params["PathPattern"] = fmt.Sprintf("/openapi/tags")
		params["ResourceIds"] = jsonutils.Marshal([]string{resourceId}).String()
	} else {
		params["ResourceId.1"] = resourceId
	}

	nextToken := ""
	for {
		if len(nextToken) > 0 {
			params["NextToken"] = nextToken
		}
		resp, err := self.tagRequest(serviceType, "ListTagResources", params, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "%s ListTagResources %s", serviceType, params)
		}
		part := []SAliyunTag{}
		err = resp.Unmarshal(&part, "TagResources", "TagResource")
		if err != nil {
			return nil, errors.Wrapf(err, "resp.Unmarshal")
		}
		tags = append(tags, part...)
		nextToken, _ = resp.GetString("NextToken")
		if len(nextToken) == 0 {
			break
		}
	}
	return tags, nil
}

func (self *SRegion) UntagResource(serviceType string, resourceType string, resId string, keys []string, removeAll bool) error {
	if len(resId) == 0 || (len(keys) == 0 && !removeAll) {
		return nil
	}

	params := map[string]string{
		"RegionId":     self.RegionId,
		"ResourceType": resourceType,
	}

	if removeAll {
		params["All"] = "true"
	}

	if serviceType == ALIYUN_SERVICE_ES {
		params["PathPattern"] = fmt.Sprintf("/openapi/tags")
		params["ResourceIds"] = jsonutils.Marshal([]string{resId}).String()
		if keys != nil && len(keys) > 0 {
			params["TagKeys"] = jsonutils.Marshal(keys).String()
		}
	} else {
		params["ResourceId.1"] = resId
		if keys != nil && len(keys) > 0 {
			for i, key := range keys {
				params[fmt.Sprintf("TagKey.%d", i+1)] = key
			}
		}
	}

	_, err := self.tagRequest(serviceType, "UntagResources", params, nil)
	return errors.Wrapf(err, "UntagResources %s", params)
}

func (self *SRegion) SetResourceTags(serviceType string, resourceType string, resId string, tags map[string]string, replace bool) error {
	_, _tags, err := self.ListSysAndUserTags(serviceType, resourceType, resId)
	if err != nil {
		return errors.Wrapf(err, "ListTags")
	}
	keys, upperKeys := []string{}, []string{}
	for k := range tags {
		keys = append(keys, k)
		upperKeys = append(upperKeys, strings.ToUpper(k))
	}
	if replace {
		if len(tags) > 0 {
			removeKeys := []string{}
			for k := range _tags {
				if !utils.IsInStringArray(k, keys) {
					removeKeys = append(removeKeys, k)
				}
			}
			if len(removeKeys) > 0 {
				err := self.UntagResource(serviceType, resourceType, resId, removeKeys, false)
				if err != nil {
					return errors.Wrapf(err, "UntagResource")
				}
			}
		}
	} else {
		removeKeys := []string{}
		for k := range _tags {
			if !utils.IsInStringArray(k, keys) && utils.IsInStringArray(strings.ToUpper(k), upperKeys) {
				removeKeys = append(removeKeys, k)
			}
		}
		if len(removeKeys) > 0 {
			err := self.UntagResource(serviceType, resourceType, resId, removeKeys, false)
			if err != nil {
				return errors.Wrapf(err, "UntagResource")
			}
		}
	}
	return self.TagResource(serviceType, resourceType, resId, tags)
}

func (self *SRegion) TagResource(serviceType string, resourceType string, resourceId string, tags map[string]string) error {
	if len(tags) > 20 {
		return errors.Wrap(cloudprovider.ErrNotSupported, "tags count exceed 20 for one request")
	}

	if len(tags) == 0 {
		return self.UntagResource(serviceType, resourceType, resourceId, nil, true)
	}

	params := make(map[string]string)

	body := map[string]interface{}{}
	if serviceType == ALIYUN_SERVICE_ES {
		params["PathPattern"] = fmt.Sprintf("/openapi/tags")
		body["ResourceIds"] = []string{resourceId}
		body["ResourceType"] = resourceType
	} else {
		params["RegionId"] = self.RegionId
		params["ResourceType"] = resourceType
		params["ResourceId.1"] = resourceId
	}

	if serviceType == ALIYUN_SERVICE_ES {
		var bodyTags []map[string]string
		for k, v := range tags {
			if strings.HasPrefix(k, "aliyun") ||
				strings.HasPrefix(k, "acs:") ||
				strings.HasPrefix(k, "http://") ||
				strings.HasPrefix(k, "https://") ||
				strings.HasPrefix(v, "http://") ||
				strings.HasPrefix(v, "https://") ||
				strings.HasPrefix(v, "acs:") {
				continue
			}
			bodyTags = append(bodyTags, map[string]string{"key": k, "value": v})
		}
		if len(bodyTags) > 0 {
			body["Tags"] = bodyTags
		}
	} else {
		i := 0
		for k, v := range tags {
			if strings.HasPrefix(k, "aliyun") ||
				strings.HasPrefix(k, "acs:") ||
				strings.HasPrefix(k, "http://") ||
				strings.HasPrefix(k, "https://") ||
				strings.HasPrefix(v, "http://") ||
				strings.HasPrefix(v, "https://") ||
				strings.HasPrefix(v, "acs:") {
				continue
			}
			params[fmt.Sprintf("Tag.%d.Key", i+1)] = k
			params[fmt.Sprintf("Tag.%d.Value", i+1)] = v
			i++
		}
	}

	_, err := self.tagRequest(serviceType, "TagResources", params, body)
	if err != nil {
		return errors.Wrapf(err, "%s %s %s", "TagResources", resourceId, params)
	}
	return nil
}

func (self *SRegion) ListSysAndUserTags(serviceType string, resourceType string, resourceId string) (map[string]string, map[string]string, error) {
	tags, err := self.ListTags(serviceType, resourceType, resourceId)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "ListTags(%s, %s)", resourceType, resourceId)
	}
	sys, user := map[string]string{}, map[string]string{}
	for _, tag := range tags {
		if strings.HasPrefix(tag.TagKey, "aliyun") || strings.HasPrefix(tag.TagKey, "acs:") {
			sys[tag.TagKey] = tag.TagValue
			continue
		}
		user[tag.TagKey] = tag.TagValue
	}
	return sys, user, nil
}

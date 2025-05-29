
package aws

import (
	"strings"

	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SAwsTag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type SAwsRdsTag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type SAwsLbTag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type AwsTags struct {
	TagSet []SAwsTag `xml:"tagSet>item"`
	// rds
	TagList []SAwsRdsTag `xml:"TagList>Tag"`
	// elb
	Tags []SAwsLbTag `xml:"Tags>member"`
}

func (self AwsTags) GetName() string {
	for _, tag := range self.TagSet {
		if strings.ToLower(tag.Key) == "name" {
			return tag.Value
		}
	}
	for _, tag := range self.TagList {
		if strings.ToLower(tag.Key) == "name" {
			return tag.Value
		}
	}
	for _, tag := range self.Tags {
		if strings.ToLower(tag.Key) == "name" {
			return tag.Value
		}
	}
	return ""
}

func (self AwsTags) GetDescription() string {
	for _, tag := range self.TagSet {
		if strings.ToLower(tag.Key) == "description" {
			return tag.Value
		}
	}
	for _, tag := range self.TagList {
		if strings.ToLower(tag.Key) == "description" {
			return tag.Value
		}
	}
	for _, tag := range self.Tags {
		if strings.ToLower(tag.Key) == "description" {
			return tag.Value
		}
	}
	return ""
}

func (self *AwsTags) GetTags() (map[string]string, error) {
	ret := map[string]string{}
	for _, tag := range self.TagSet {
		if tag.Key == "Name" || tag.Key == "Description" {
			continue
		}
		if strings.HasPrefix(tag.Key, "aws:") {
			continue
		}
		ret[tag.Key] = tag.Value
	}
	for _, tag := range self.TagList {
		if strings.ToLower(tag.Key) == "name" || strings.ToLower(tag.Key) == "description" {
			continue
		}
		if strings.HasPrefix(tag.Key, "aws:") {
			continue
		}
		ret[tag.Key] = tag.Value
	}
	for _, tag := range self.Tags {
		if strings.ToLower(tag.Key) == "name" || strings.ToLower(tag.Key) == "description" {
			continue
		}
		if strings.HasPrefix(tag.Key, "aws:") {
			continue
		}
		ret[tag.Key] = tag.Value
	}
	return ret, nil
}

func (self *AwsTags) GetSysTags() map[string]string {
	ret := map[string]string{}
	for _, tag := range self.TagSet {
		if strings.HasPrefix(tag.Key, "aws:") {
			ret[tag.Key] = tag.Value
		}
	}
	for _, tag := range self.TagList {
		if strings.HasPrefix(tag.Key, "aws:") {
			ret[tag.Key] = tag.Value
		}
	}
	for _, tag := range self.Tags {
		if strings.HasPrefix(tag.Key, "aws:") {
			ret[tag.Key] = tag.Value
		}
	}
	return ret
}

func (self *AwsTags) SetTags(tags map[string]string, replace bool) error {
	return errors.Wrap(cloudprovider.ErrNotImplemented, "SetTags")
}

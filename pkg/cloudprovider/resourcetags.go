package cloudprovider

import (
	"context"
	"reflect"
	"strings"
	"time"

	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
)

const (
	SET_TAGS = "set-tags"
)

type TagsUpdateInfo struct {
	OldTags map[string]string
	NewTags map[string]string
}

func (t TagsUpdateInfo) IsChanged() bool {
	return !reflect.DeepEqual(t.OldTags, t.NewTags)
}

func SetTags(ctx context.Context, res ICloudResource, managerId string, tags map[string]string, replace bool) error {
	err := res.SetTags(tags, replace)
	if err != nil {
		return errors.Wrapf(err, "SetTags")
	}

	// 避免设置标签后未及时生效，导致本地同步和云上不一致
	Wait(time.Second*5, time.Minute, func() (bool, error) {
		res.Refresh()
		newTags, err := res.GetTags()
		if err != nil {
			return false, errors.Wrapf(err, "GetTags")
		}
		for k, v := range tags {
			_, ok := newTags[k]
			_, ok2 := newTags[strings.ToLower(k)]
			if !ok && !ok2 {
				log.Warningf("tag %s:%s not found waitting....", k, v)
				return false, nil
			}
		}
		return true, nil
	})
	return nil
}

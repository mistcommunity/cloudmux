
package shell

import (
	"strings"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type TagGetOptions struct {
		SERVICE string `help:"service, eg. ecs" choices:"ecs|kvs|rds|vpc|slb"`
		RESTYPE string `help:"resource type, eg. instance"`
		ID      string `help:"resource Id, eg. ins-123xxx"`
	}
	shellutils.R(&TagGetOptions{}, "tag-list", "List tag of a specific resource", func(cli *aliyun.SRegion, args *TagGetOptions) error {
		tags, err := cli.ListTags(args.SERVICE, args.RESTYPE, args.ID)
		if err != nil {
			return err
		}
		printObject(tags)
		return nil
	})

	type TagOptions struct {
		TagGetOptions
		KEY   string
		VALUE string
	}
	shellutils.R(&TagOptions{}, "tag-resource", "set tags of a specific resource", func(cli *aliyun.SRegion, args *TagOptions) error {
		return cli.TagResource(args.SERVICE, args.RESTYPE, args.ID, map[string]string{args.KEY: args.VALUE})
	})

	type UnTagOptions struct {
		TagGetOptions
		KEY []string
	}

	shellutils.R(&UnTagOptions{}, "untag-resource", "un tags of a specific resource", func(cli *aliyun.SRegion, args *UnTagOptions) error {
		return cli.UntagResource(args.SERVICE, args.RESTYPE, args.ID, args.KEY, false)
	})

	type TagSetOptions struct {
		TagGetOptions
		VALUES  []string
		Replace bool
	}

	shellutils.R(&TagSetOptions{}, "tag-set", "set tags of a specific resource", func(cli *aliyun.SRegion, args *TagSetOptions) error {
		tags := map[string]string{}
		for _, value := range args.VALUES {
			v := strings.Split(value, ":")
			if len(v) != 2 {
				return errors.Errorf("invalid tag %s", value)
			}
			tags[v[0]] = v[1]
		}
		return cli.SetResourceTags(args.SERVICE, args.RESTYPE, args.ID, tags, args.Replace)
	})

}

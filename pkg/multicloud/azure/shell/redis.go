
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type RedisListOptions struct {
	}
	shellutils.R(&RedisListOptions{}, "redis-list", "List redis", func(cli *azure.SRegion, args *RedisListOptions) error {
		redis, err := cli.GetRedisCaches()
		if err != nil {
			return err
		}
		printList(redis, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&RedisListOptions{}, "enterprise-redis-list", "List enterprise redis", func(cli *azure.SRegion, args *RedisListOptions) error {
		redis, err := cli.GetEnterpriseRedisCaches()
		if err != nil {
			return err
		}
		printList(redis, 0, 0, 0, nil)
		return nil
	})

	type RedisIdOptions struct {
		ID string
	}
	shellutils.R(&RedisIdOptions{}, "redis-acl-list", "List redis acls", func(cli *azure.SRegion, args *RedisIdOptions) error {
		acls, err := cli.GetRedisAcls(args.ID)
		if err != nil {
			return err
		}
		printList(acls, 0, 0, 0, nil)
		return nil
	})

}

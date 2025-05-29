
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ElkListOptions struct {
		Page int
		Size int
	}
	shellutils.R(&ElkListOptions{}, "elastic-search-list", "List elastic searchs", func(cli *aliyun.SRegion, args *ElkListOptions) error {
		elks, _, err := cli.GetElasticSearchs(args.Size, args.Page)
		if err != nil {
			return err
		}
		printList(elks, 0, 0, 0, nil)
		return nil
	})

	type ElkIdOptions struct {
		ID string
	}

	shellutils.R(&ElkIdOptions{}, "elastic-search-show", "Show elasitc search", func(cli *aliyun.SRegion, args *ElkIdOptions) error {
		elk, err := cli.GetElasitcSearch(args.ID)
		if err != nil {
			return err
		}
		printObject(elk)
		return nil
	})

	shellutils.R(&ElkIdOptions{}, "elastic-search-delete", "Delete elasitc search", func(cli *aliyun.SRegion, args *ElkIdOptions) error {
		return cli.DeleteElasticSearch(args.ID)
	})

}

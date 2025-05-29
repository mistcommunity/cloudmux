
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ContainerGroupListOptions struct {
	}
	shellutils.R(&ContainerGroupListOptions{}, "container-group-list", "list container groups", func(cli *aliyun.SRegion, args *ContainerGroupListOptions) error {
		ret, err := cli.GetContainerGroups()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

}

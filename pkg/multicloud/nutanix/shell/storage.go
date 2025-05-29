
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type StorageListOptions struct {
	}
	shellutils.R(&StorageListOptions{}, "storage-list", "list storage", func(cli *nutanix.SRegion, args *StorageListOptions) error {
		storages, err := cli.GetStorages()
		if err != nil {
			return err
		}
		printList(storages, 0, 0, 0, []string{})
		return nil
	})

	type StorageIdOptions struct {
		ID string
	}

	shellutils.R(&StorageIdOptions{}, "storage-show", "show storage", func(cli *nutanix.SRegion, args *StorageIdOptions) error {
		storage, err := cli.GetStorage(args.ID)
		if err != nil {
			return err
		}
		printObject(storage)
		return nil
	})

}

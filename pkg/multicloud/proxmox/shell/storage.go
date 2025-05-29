
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type StorageListOptions struct {
	}
	shellutils.R(&StorageListOptions{}, "storage-list", "list storage", func(cli *proxmox.SRegion, args *StorageListOptions) error {
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

	shellutils.R(&StorageIdOptions{}, "storage-show", "show storage", func(cli *proxmox.SRegion, args *StorageIdOptions) error {
		storage, err := cli.GetStorage(args.ID)
		if err != nil {
			return err
		}
		printObject(storage)
		return nil
	})

}

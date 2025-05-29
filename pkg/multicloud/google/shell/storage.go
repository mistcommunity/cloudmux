
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type StorageListOptions struct {
		ZONE       string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&StorageListOptions{}, "storage-list", "List storages", func(cli *google.SRegion, args *StorageListOptions) error {
		storages, err := cli.GetStorages(args.ZONE, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(storages, 0, 0, 0, nil)
		return nil
	})

	type StorageShowOptions struct {
		ID string
	}
	shellutils.R(&StorageShowOptions{}, "storage-show", "Show storage", func(cli *google.SRegion, args *StorageShowOptions) error {
		storage, err := cli.GetStorage(args.ID)
		if err != nil {
			return err
		}
		printObject(storage)
		return nil
	})

	type RegionStorageListOptions struct {
		MaxResults int
		PageToken  string
	}
	shellutils.R(&RegionStorageListOptions{}, "region-storage-list", "List region storages", func(cli *google.SRegion, args *RegionStorageListOptions) error {
		storages, err := cli.GetRegionStorages(args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(storages, 0, 0, 0, nil)
		return nil
	})

	type RegionStorageShowOptions struct {
		ID string
	}
	shellutils.R(&RegionStorageShowOptions{}, "region-storage-show", "Show region storage", func(cli *google.SRegion, args *RegionStorageShowOptions) error {
		storage, err := cli.GetRegionStorage(args.ID)
		if err != nil {
			return err
		}
		printObject(storage)
		return nil
	})

}

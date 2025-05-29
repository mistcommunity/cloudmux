
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type FileSystemListOptions struct {
		Id       string `help:"FileSystem Id"`
		PageSize int    `help:"page size"`
		PageNum  int    `help:"page num"`
	}
	shellutils.R(&FileSystemListOptions{}, "file-system-list", "List FileSystem", func(cli *aliyun.SRegion, args *FileSystemListOptions) error {
		nas, _, err := cli.GetFileSystems(args.Id, args.PageSize, args.PageNum)
		if err != nil {
			return err
		}
		printList(nas, 0, 0, 0, []string{})
		return nil
	})

	type FileSystemDeleteOptions struct {
		ID string `help:"File System ID"`
	}
	shellutils.R(&FileSystemDeleteOptions{}, "file-system-delete", "Delete filesystem", func(cli *aliyun.SRegion, args *FileSystemDeleteOptions) error {
		return cli.DeleteFileSystem(args.ID)
	})

	shellutils.R(&cloudprovider.FileSystemCraeteOptions{}, "file-system-create", "Create filesystem", func(cli *aliyun.SRegion, args *cloudprovider.FileSystemCraeteOptions) error {
		fs, err := cli.CreateFileSystem(args)
		if err != nil {
			return err
		}
		printObject(fs)
		return nil
	})

}

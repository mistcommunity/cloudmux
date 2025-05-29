
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type DiskListOptions struct {
		NODE    string
		STORAGE string
	}
	shellutils.R(&DiskListOptions{}, "disk-list", "list disks", func(cli *proxmox.SRegion, args *DiskListOptions) error {
		disks, err := cli.GetDisks(args.NODE, args.STORAGE)
		if err != nil {
			return err
		}
		printList(disks, 0, 0, 0, []string{})
		return nil
	})

	type DiskResizeOptions struct {
		NODE    string
		VM_ID   string
		DRIVER  string
		SIZE_GB int
	}

	shellutils.R(&DiskResizeOptions{}, "disk-resize", "resize disk size", func(cli *proxmox.SRegion, args *DiskResizeOptions) error {
		return cli.ResizeDisk(args.NODE, args.VM_ID, args.DRIVER, args.SIZE_GB)
	})

}

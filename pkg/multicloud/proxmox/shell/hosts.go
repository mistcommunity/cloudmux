
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type HostListOptions struct {
	}
	shellutils.R(&HostListOptions{}, "host-list", "list hosts", func(cli *proxmox.SRegion, args *HostListOptions) error {
		hosts, err := cli.GetHosts()
		if err != nil {
			return err
		}
		printList(hosts, 0, 0, 0, []string{})
		return nil
	})

	type HostIdOptions struct {
		ID string
	}

	shellutils.R(&HostIdOptions{}, "host-show", "show host", func(cli *proxmox.SRegion, args *HostIdOptions) error {
		host, err := cli.GetHost(args.ID)
		if err != nil {
			return err
		}
		printObject(host)
		printObject(host.Cpuinfo)
		return nil
	})

}

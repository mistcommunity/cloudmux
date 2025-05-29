
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type HostListOptions struct {
	}
	shellutils.R(&HostListOptions{}, "host-list", "list hosts", func(cli *nutanix.SRegion, args *HostListOptions) error {
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

	shellutils.R(&HostIdOptions{}, "host-show", "show host", func(cli *nutanix.SRegion, args *HostIdOptions) error {
		host, err := cli.GetHost(args.ID)
		if err != nil {
			return err
		}
		printObject(host)
		return nil
	})

}

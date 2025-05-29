
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type ZoneListOptions struct {
	}
	shellutils.R(&ZoneListOptions{}, "zone-list", "list zones", func(cli *proxmox.SRegion, args *ZoneListOptions) error {
		zone, err := cli.GetZone()
		if err != nil {
			return err
		}
		printObject(zone)
		return nil
	})

}

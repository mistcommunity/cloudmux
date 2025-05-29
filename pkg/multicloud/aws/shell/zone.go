package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type ZoneListOptions struct {
		Id string
	}
	shellutils.R(&ZoneListOptions{}, "zone-list", "List zones", func(cli *aws.SRegion, args *ZoneListOptions) error {
		zones, err := cli.GetZones(args.Id)
		if err != nil {
			return err
		}
		printList(zones, 0, 0, 0, nil)
		return nil
	})
}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type ZoneListOptions struct {
		RegionId   string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&ZoneListOptions{}, "zone-list", "List zones", func(cli *google.SRegion, args *ZoneListOptions) error {
		zones, err := cli.GetZones(args.RegionId, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(zones, 0, 0, 0, nil)
		return nil
	})

	type ZoneShowOptions struct {
		ID string
	}
	shellutils.R(&ZoneShowOptions{}, "zone-show", "Show zones", func(cli *google.SRegion, args *ZoneShowOptions) error {
		zone, err := cli.GetZone(args.ID)
		if err != nil {
			return err
		}
		printObject(zone)
		return nil
	})

}

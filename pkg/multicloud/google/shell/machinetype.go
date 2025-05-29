
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type MachineTypeListOptions struct {
		ZONE       string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&MachineTypeListOptions{}, "machine-type-list", "List machinetypes", func(cli *google.SRegion, args *MachineTypeListOptions) error {
		machinetypes, err := cli.GetMachineTypes(args.ZONE, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(machinetypes, 0, 0, 0, nil)
		return nil
	})

	type MachineTypeShowOptions struct {
		ID string
	}
	shellutils.R(&MachineTypeShowOptions{}, "machine-type-show", "Show machinetype", func(cli *google.SRegion, args *MachineTypeShowOptions) error {
		machinetype, err := cli.GetMachineType(args.ID)
		if err != nil {
			return err
		}
		printObject(machinetype)
		return nil
	})

}

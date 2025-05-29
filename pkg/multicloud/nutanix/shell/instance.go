
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type InstanceListOptions struct {
	}
	shellutils.R(&InstanceListOptions{}, "instance-list", "list instances", func(cli *nutanix.SRegion, args *InstanceListOptions) error {
		vms, err := cli.GetInstances()
		if err != nil {
			return err
		}
		printList(vms, 0, 0, 0, []string{})
		return nil
	})

	type InstanceIdOptions struct {
		ID string
	}

	shellutils.R(&InstanceIdOptions{}, "instance-show", "show instance", func(cli *nutanix.SRegion, args *InstanceIdOptions) error {
		vm, err := cli.GetInstance(args.ID)
		if err != nil {
			return err
		}
		printObject(vm)
		return nil
	})

}

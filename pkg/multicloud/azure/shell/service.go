
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type ServiceListOptions struct {
	}
	shellutils.R(&ServiceListOptions{}, "service-list", "List providers", func(cli *azure.SRegion, args *ServiceListOptions) error {
		services, err := cli.GetClient().ListServices()
		if err != nil {
			return err
		}
		printList(services, len(services), 0, 0, []string{})
		return nil
	})

	type ServiceOptions struct {
		NAME string `help:"Name for service register"`
	}

	shellutils.R(&ServiceOptions{}, "service-register", "Register service", func(cli *azure.SRegion, args *ServiceOptions) error {
		return cli.GetClient().ServiceRegister(args.NAME)
	})

	shellutils.R(&ServiceOptions{}, "service-unregister", "Unregister service", func(cli *azure.SRegion, args *ServiceOptions) error {
		return cli.GetClient().ServiceUnRegister(args.NAME)
	})

	shellutils.R(&ServiceOptions{}, "service-show", "Show service detail", func(cli *azure.SRegion, args *ServiceOptions) error {
		service, err := cli.GetClient().GetSercice(args.NAME)
		if err != nil {
			return err
		}
		printObject(service)
		return nil
	})

}

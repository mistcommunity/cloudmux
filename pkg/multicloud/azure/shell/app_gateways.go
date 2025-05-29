
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type AppGatewayListOptions struct {
	}
	shellutils.R(&AppGatewayListOptions{}, "app-gateway-list", "List app gateways", func(cli *azure.SRegion, args *AppGatewayListOptions) error {
		apps, err := cli.ListAppGateways()
		if err != nil {
			return err
		}
		printList(apps, len(apps), 0, 0, []string{})
		return nil
	})
}

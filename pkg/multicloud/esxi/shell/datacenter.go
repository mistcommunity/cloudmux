
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/esxi"
)

func init() {
	type DatacenterListOptions struct {
	}
	shellutils.R(&DatacenterListOptions{}, "dc-list", "List all datacenters", func(cli *esxi.SESXiClient, args *DatacenterListOptions) error {
		dcs, err := cli.GetDatacenters()
		if err != nil {
			return err
		}
		printList(dcs, nil)
		return nil
	})
}

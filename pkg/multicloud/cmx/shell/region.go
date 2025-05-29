package shell

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func init() {
	r := EmptyOptionProviderR("region")

	r.List("list", "List regions", func(cli cloudprovider.ICloudProvider) (any, error) {
		return cli.GetIRegions()
	})
}

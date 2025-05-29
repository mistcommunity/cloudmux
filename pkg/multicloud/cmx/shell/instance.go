package shell

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func init() {
	cmd := NewCommand("instance")

	type InstanceListOptions struct {
		ListBaseOptions
		HostBaseOptions

		Id []string `help:"IDs of instances to show"`
	}
	HostR[InstanceListOptions](cmd).List("list", "List instances", func(cli cloudprovider.ICloudHost, _ *InstanceListOptions) (any, error) {
		return cli.GetIVMs()
	})
}

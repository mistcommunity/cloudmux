package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type LbbgIdOptions struct {
		ID string
	}
	shellutils.R(&LbbgIdOptions{}, "elb-lbbg-show", "Show loadbalancer backendgroup", func(cli *aws.SRegion, args *LbbgIdOptions) error {
		ret, err := cli.GetElbBackendgroup(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&LbbgIdOptions{}, "elb-lbbg-delete", "Delete loadbalancer", func(cli *aws.SRegion, args *LbbgIdOptions) error {
		return cli.DeleteElbBackendGroup(args.ID)
	})

}

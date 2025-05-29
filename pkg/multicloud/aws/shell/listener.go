package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type LbIdOptions struct {
		ID string
	}
	shellutils.R(&LbIdOptions{}, "elb-listener-show", "Show loadbalancer listener", func(cli *aws.SRegion, args *LbIdOptions) error {
		ret, err := cli.GetElbListener(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&LbIdOptions{}, "elb-listener-delete", "Delete loadbalancer listener", func(cli *aws.SRegion, args *LbIdOptions) error {
		return cli.DeleteElbListener(args.ID)
	})

}

package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type CallerIdentityOptions struct {
	}
	shellutils.R(&CallerIdentityOptions{}, "caller-identity-show", "Show caller identity", func(cli *aws.SRegion, args *CallerIdentityOptions) error {
		identity, err := cli.GetClient().GetCallerIdentity()
		if err != nil {
			return err
		}
		printObject(identity)
		return nil
	})
}

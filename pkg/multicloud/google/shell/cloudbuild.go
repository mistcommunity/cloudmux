
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type CloudbuildOperationShowOptions struct {
		NAME string
	}

	shellutils.R(&CloudbuildOperationShowOptions{}, "cloud-build-operation-show", "Show cloudbuild operation", func(cli *google.SRegion, args *CloudbuildOperationShowOptions) error {
		operation, err := cli.GetCloudbuildOperation(args.NAME)
		if err != nil {
			return err
		}
		printObject(operation)
		return nil
	})
}

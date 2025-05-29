package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type DBInstanceParameterOptions struct {
		NAME string
	}

	shellutils.R(&DBInstanceParameterOptions{}, "dbinstance-parameter-show", "Show rds intance parameter", func(cli *aws.SRegion, args *DBInstanceParameterOptions) error {
		parameters, err := cli.GetDBInstanceParameters(args.NAME)
		if err != nil {
			return err
		}
		printList(parameters, 0, 0, 0, []string{})
		return nil
	})

}

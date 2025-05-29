package shell

import (
	"yunion.io/x/log"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type DBInstanceListOptions struct {
		Id     string
		Marker string
	}
	shellutils.R(&DBInstanceListOptions{}, "dbinstance-list", "List rds intances", func(cli *aws.SRegion, args *DBInstanceListOptions) error {
		instances, _, err := cli.GetDBInstances(args.Id, args.Marker)
		if err != nil {
			return err
		}
		printList(instances, 0, 0, 0, []string{})
		return nil
	})

	type DBInstanceIdOptions struct {
		ID string
	}

	shellutils.R(&DBInstanceIdOptions{}, "dbinstance-show", "Show rds intance", func(cli *aws.SRegion, args *DBInstanceIdOptions) error {
		instance, err := cli.GetDBInstance(args.ID)
		if err != nil {
			return err
		}
		printObject(instance)
		return nil
	})

	shellutils.R(&DBInstanceIdOptions{}, "dbinstance-tags-list", "Show rds intance tags", func(cli *aws.SRegion, args *DBInstanceIdOptions) error {
		instance, err := cli.ListRdsResourceTags(args.ID)
		if err != nil {
			return err
		}
		printObject(instance)
		return nil
	})

	type SDBInstanceUpdateOptions struct {
		ID string
		cloudprovider.SDBInstanceUpdateOptions
	}
	shellutils.R(&SDBInstanceUpdateOptions{}, "dbinstance-update", "Show rds intance tags", func(cli *aws.SRegion, args *SDBInstanceUpdateOptions) error {
		err := cli.Update(args.ID, args.SDBInstanceUpdateOptions)
		if err != nil {
			log.Errorln(err)
		}
		return err
	})

}

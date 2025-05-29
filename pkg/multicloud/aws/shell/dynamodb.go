package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type DynamodbListOptions struct {
	}
	shellutils.R(&DynamodbListOptions{}, "dynamodb-table-list", "List kinesis stream", func(cli *aws.SRegion, args *DynamodbListOptions) error {
		tables, err := cli.ListTables()
		if err != nil {
			return err
		}
		fmt.Println(tables)
		return nil
	})

	type DynamodbNameOptions struct {
		NAME string
	}
	shellutils.R(&DynamodbNameOptions{}, "dynamodb-table-show", "Show kinesis stream", func(cli *aws.SRegion, args *DynamodbNameOptions) error {
		table, err := cli.DescribeTable(args.NAME)
		if err != nil {
			return err
		}
		printObject(table)
		return nil
	})

}

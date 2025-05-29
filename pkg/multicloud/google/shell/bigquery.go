
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type BigQueryOptions struct {
		SQL string `help:"sql to query"`
	}
	shellutils.R(&BigQueryOptions{}, "gcp-bigquery", "Exec bigquery", func(cli *google.SRegion, args *BigQueryOptions) error {
		results, err := cli.BigQuery(args.SQL)
		if err != nil {
			return err
		}

		printList(results, 0, 0, 0, nil)
		return nil
	})
}

package shell

import (
	"strings"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type ProductListOptions struct {
		ServiceCode string `default:"AmazonEC2" choices:"AmazonEC2|AmazonElastiCache"`
		NextToken   string
		Filters     []string
	}
	shellutils.R(&ProductListOptions{}, "product-list", "List product", func(cli *aws.SRegion, args *ProductListOptions) error {
		filters := []aws.ProductFilter{}
		for _, filter := range args.Filters {
			info := strings.Split(filter, "=")
			if len(info) == 2 {
				filters = append(filters, aws.ProductFilter{
					Type:  "TERM_MATCH",
					Field: info[0],
					Value: info[1],
				})
			}
		}
		products, _, err := cli.GetProducts(args.ServiceCode, filters, args.NextToken)
		if err != nil {
			return err
		}
		printList(products, 0, 0, 0, []string{})
		return nil
	})
}

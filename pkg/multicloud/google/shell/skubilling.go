
package shell

import (
	"fmt"
	"io/ioutil"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type SkuBillingListOptions struct {
		PageSize  int
		PageToken string
	}
	shellutils.R(&SkuBillingListOptions{}, "sku-billing-list", "List sku billing", func(cli *google.SRegion, args *SkuBillingListOptions) error {
		billings, err := cli.ListSkuBilling(args.PageSize, args.PageToken)
		if err != nil {
			return err
		}
		printList(billings, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&SkuBillingListOptions{}, "compute-sku-billing-list", "List sku billing", func(cli *google.SRegion, args *SkuBillingListOptions) error {
		billings, err := cli.ListSkuBilling(args.PageSize, args.PageToken)
		if err != nil {
			return err
		}
		info := cli.GetSkuRateInfo(billings)
		fmt.Println(jsonutils.Marshal(info).PrettyString())
		return nil
	})

	type SkuEstimate struct {
		RATE_FAILE string
		SKU        string
		REGION     string
		CPU        int
		MEMORY_MB  int
	}

	shellutils.R(&SkuEstimate{}, "sku-estimate", "Estimate sku price", func(cli *google.SRegion, args *SkuEstimate) error {
		data, err := ioutil.ReadFile(args.RATE_FAILE)
		if err != nil {
			return errors.Wrap(err, "ioutil.ReadFile")
		}
		rate := google.SRateInfo{}
		j, err := jsonutils.Parse(data)
		if err != nil {
			return errors.Wrap(err, "jsonutils.Parse")
		}
		err = jsonutils.Update(&rate, j)
		if err != nil {
			return errors.Wrap(err, "jsonutils.Update")
		}
		result, err := rate.GetSkuPrice(args.REGION, args.SKU, args.CPU, args.MEMORY_MB)
		if err != nil {
			return errors.Wrap(err, "GetSkuPrice")
		}
		fmt.Printf("result: %s\n", jsonutils.Marshal(result).PrettyString())
		return nil
	})

}

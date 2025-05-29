
package shell

import (
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type WafShowOptions struct {
	}
	shellutils.R(&WafShowOptions{}, "waf-v2-instance-show", "Show waf instance", func(cli *aliyun.SRegion, args *WafShowOptions) error {
		waf, err := cli.DescribeWafInstance()
		if err != nil {
			return err
		}
		printObject(waf)
		return nil
	})

	type WafIdOptions struct {
		ID string
	}

	shellutils.R(&WafIdOptions{}, "waf-v2-domain-list", "List waf instance domains", func(cli *aliyun.SRegion, args *WafIdOptions) error {
		domains, err := cli.DescribeWafDomains(args.ID)
		if err != nil {
			return errors.Wrapf(err, "DescribeDomainNames")
		}
		printList(domains, 0, 0, 0, nil)
		return nil
	})

	type WafDomainIdOptions struct {
		ID     string
		DOMAIN string
	}

	shellutils.R(&WafDomainIdOptions{}, "waf-v2-domain-show", "Show waf domain", func(cli *aliyun.SRegion, args *WafDomainIdOptions) error {
		domain, err := cli.DescribeDomainV2(args.ID, args.DOMAIN)
		if err != nil {
			return err
		}
		printObject(domain)
		return nil
	})
}

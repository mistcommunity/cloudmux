
package shell

import (
	"fmt"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type WafShowOptions struct {
	}
	shellutils.R(&WafShowOptions{}, "waf-instance-show", "Show waf instance", func(cli *aliyun.SRegion, args *WafShowOptions) error {
		waf, err := cli.DescribeInstanceSpecInfo()
		if err != nil {
			return err
		}
		printObject(waf)
		return nil
	})

	type WafIdOptions struct {
		ID string
	}

	shellutils.R(&WafIdOptions{}, "waf-instance-delete", "Delete waf instance", func(cli *aliyun.SRegion, args *WafIdOptions) error {
		return cli.DeleteInstance(args.ID)
	})

	shellutils.R(&WafIdOptions{}, "waf-domain-list", "List waf instance domains", func(cli *aliyun.SRegion, args *WafIdOptions) error {
		domains, err := cli.DescribeDomainNames(args.ID)
		if err != nil {
			return errors.Wrapf(err, "DescribeDomainNames")
		}
		fmt.Println("domains: ", domains)
		return nil
	})

	type WafDomainIdOptions struct {
		ID     string
		DOMAIN string
	}

	shellutils.R(&WafDomainIdOptions{}, "waf-domain-show", "Show waf domain", func(cli *aliyun.SRegion, args *WafDomainIdOptions) error {
		domain, err := cli.DescribeDomain(args.ID, args.DOMAIN)
		if err != nil {
			return err
		}
		printObject(domain)
		return nil
	})

	shellutils.R(&WafDomainIdOptions{}, "waf-domain-delete", "Delete waf domain", func(cli *aliyun.SRegion, args *WafDomainIdOptions) error {
		return cli.DeleteDomain(args.ID, args.DOMAIN)
	})

}

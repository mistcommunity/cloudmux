package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
	"yunion.io/x/pkg/util/shellutils"
)

func init() {
	type CdnDomainListOptions struct {
		Domain     string
		PageSize   int
		PageNumber int
	}
	shellutils.R(&CdnDomainListOptions{}, "cdn-domain-list", "list cdn domain", func(cli *aws.SRegion, args *CdnDomainListOptions) error {
		domains, _, err := cli.GetClient().DescribeUserDomains("", int64(args.PageSize))
		if err != nil {
			return err
		}
		printList(domains, 0, 0, 0, []string{})
		return nil
	})

	type CdnDomainShowOptions struct {
		DOMAIN string
	}
	shellutils.R(&CdnDomainShowOptions{}, "cdn-domain-show", "Show cdn domain", func(cli *aws.SRegion, args *CdnDomainShowOptions) error {
		domain, err := cli.GetClient().GetCdnDomain(args.DOMAIN)
		if err != nil {
			return err
		}
		printObject(domain)
		return nil
	})

}

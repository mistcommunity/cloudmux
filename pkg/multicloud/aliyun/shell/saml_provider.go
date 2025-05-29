
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type SamlProviderListOptions struct {
		Marker   string `help:"Marker"`
		MaxItems int    `help:"Max items"`
	}
	shellutils.R(&SamlProviderListOptions{}, "saml-provider-list", "List saml provider", func(cli *aliyun.SRegion, args *SamlProviderListOptions) error {
		result, _, err := cli.GetClient().ListSAMLProviders(args.Marker, args.MaxItems)
		if err != nil {
			return err
		}
		printList(result, 0, 0, 0, []string{})
		return nil
	})

	type SamlProviderDeleteOptions struct {
		NAME string `help:"SAML Provider Name"`
	}

	shellutils.R(&SamlProviderDeleteOptions{}, "saml-provider-delete", "Delete saml provider", func(cli *aliyun.SRegion, args *SamlProviderDeleteOptions) error {
		return cli.GetClient().DeleteSAMLProvider(args.NAME)
	})

	type SAMLProviderCreateOptions struct {
		NAME    string
		METADAT string
		Desc    string
	}

	shellutils.R(&SAMLProviderCreateOptions{}, "saml-provider-create", "Create saml provider", func(cli *aliyun.SRegion, args *SAMLProviderCreateOptions) error {
		sp, err := cli.GetClient().CreateSAMLProvider(args.NAME, args.METADAT, args.Desc)
		if err != nil {
			return err
		}
		printObject(sp)
		return nil
	})

}

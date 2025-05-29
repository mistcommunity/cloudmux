package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type SAMLProviderListOptions struct {
	}
	shellutils.R(&SAMLProviderListOptions{}, "saml-provider-list", "List saml providers", func(cli *aws.SRegion, args *SAMLProviderListOptions) error {
		samls, err := cli.GetClient().ListSAMLProviders()
		if err != nil {
			return err
		}
		printList(samls, 0, 0, 0, []string{})
		return nil
	})

	type SAMLProviderArnOptions struct {
		ARN string
	}

	shellutils.R(&SAMLProviderArnOptions{}, "saml-provider-show", "Show saml provider", func(cli *aws.SRegion, args *SAMLProviderArnOptions) error {
		saml, err := cli.GetClient().GetSAMLProvider(args.ARN)
		if err != nil {
			return err
		}
		printObject(saml)
		return nil
	})

	shellutils.R(&SAMLProviderArnOptions{}, "saml-provider-delete", "Delete saml provider", func(cli *aws.SRegion, args *SAMLProviderArnOptions) error {
		return cli.GetClient().DeleteSAMLProvider(args.ARN)
	})

	type SAMLProviderCreateOptions struct {
		NAME     string
		METADATA string
	}

	shellutils.R(&SAMLProviderCreateOptions{}, "saml-provider-create", "Create saml provider", func(cli *aws.SRegion, args *SAMLProviderCreateOptions) error {
		saml, err := cli.GetClient().CreateSAMLProvider(args.NAME, args.METADATA)
		if err != nil {
			return err
		}
		printObject(saml)
		return nil
	})

	type SAMLProviderUpdateOptions struct {
		ARN      string
		METADATA string
	}

	shellutils.R(&SAMLProviderUpdateOptions{}, "saml-provider-update", "Update saml provider", func(cli *aws.SRegion, args *SAMLProviderUpdateOptions) error {
		saml, err := cli.GetClient().UpdateSAMLProvider(args.ARN, args.METADATA)
		if err != nil {
			return err
		}
		printObject(saml)
		return nil
	})

}


package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type SAMLProviderListOptions struct {
	}
	shellutils.R(&SAMLProviderListOptions{}, "saml-provider-list", "List regions", func(cli *azure.SRegion, args *SAMLProviderListOptions) error {
		sps, err := cli.GetClient().ListSAMLProviders()
		if err != nil {
			return err
		}
		printList(sps, 0, 0, 0, nil)
		return nil
	})

	type SInvitateUser struct {
		EMAIL string
	}

	shellutils.R(&SInvitateUser{}, "invite-user", "Invitate user", func(cli *azure.SRegion, args *SInvitateUser) error {
		user, err := cli.GetClient().InviteUser(args.EMAIL)
		if err != nil {
			return err
		}
		printObject(user)
		fmt.Println("invite url: ", user.GetInviteUrl())
		return nil
	})

}

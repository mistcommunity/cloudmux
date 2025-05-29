
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type SubscriptionListOptions struct {
	}
	shellutils.R(&SubscriptionListOptions{}, "subscription-list", "List subscriptions", func(cli *azure.SRegion, args *SubscriptionListOptions) error {
		subscriptions, err := cli.GetClient().ListSubscriptions()
		if err != nil {
			return err
		}
		printList(subscriptions, 0, 0, 0, nil)
		return nil
	})

	type SubscriptionCreateOptions struct {
		NAME      string
		OfferType string `choices:"MS-AZR-0017P|MS-AZR-0148P" default:"MS-AZR-0017P"`
		EAID      string `help:"Enrollment accounts id"`
	}

	shellutils.R(&SubscriptionCreateOptions{}, "subscription-create", "Create subscription", func(cli *azure.SRegion, args *SubscriptionCreateOptions) error {
		return cli.GetClient().CreateSubscription(args.NAME, args.EAID, args.OfferType)
	})

	type ServicePrincipalListOptions struct {
		AppId string
	}

	shellutils.R(&ServicePrincipalListOptions{}, "sp-list", "List service principal", func(cli *azure.SRegion, args *ServicePrincipalListOptions) error {
		sp, err := cli.GetClient().ListServicePrincipal(args.AppId)
		if err != nil {
			return err
		}
		printList(sp, 0, 0, 0, nil)
		return nil
	})

}

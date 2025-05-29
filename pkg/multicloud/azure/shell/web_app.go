
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type AppSiteListOptions struct {
	}
	shellutils.R(&AppSiteListOptions{}, "web-app-list", "List app service plan", func(cli *azure.SRegion, args *AppSiteListOptions) error {
		ass, err := cli.GetAppSites()
		if err != nil {
			return err
		}
		printList(ass, len(ass), 0, 0, []string{})
		return nil
	})
	type AppSiteShowOptions struct {
		ID string
	}
	shellutils.R(&AppSiteShowOptions{}, "web-app-show", "Show app service plan", func(cli *azure.SRegion, args *AppSiteShowOptions) error {
		as, err := cli.GetAppSite(args.ID)
		if err != nil {
			return err
		}
		printObject(as)
		return nil
	})

	shellutils.R(&AppSiteShowOptions{}, "web-app-backup-list", "List web app backups", func(cli *azure.SRegion, args *AppSiteShowOptions) error {
		backups, err := cli.GetAppBackups(args.ID)
		if err != nil {
			return err
		}
		printList(backups, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&AppSiteShowOptions{}, "web-app-backup-config-show", "Show web app backup config", func(cli *azure.SRegion, args *AppSiteShowOptions) error {
		opts, err := cli.GetAppBackupConfig(args.ID)
		if err != nil {
			return err
		}
		printObject(opts)
		return nil
	})

	shellutils.R(&AppSiteShowOptions{}, "web-app-hybird-connection-list", "List web app hybird connections", func(cli *azure.SRegion, args *AppSiteShowOptions) error {
		ret, err := cli.GetAppHybirConnections(args.ID)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&AppSiteShowOptions{}, "web-app-certificate-list", "List web app certificates", func(cli *azure.SRegion, args *AppSiteShowOptions) error {
		ret, err := cli.GetAppCertificates(args.ID)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&AppSiteShowOptions{}, "web-app-slot-list", "List slots ofr App site", func(cli *azure.SRegion, args *AppSiteShowOptions) error {
		slots, err := cli.GetSlots(args.ID)
		if err != nil {
			return err
		}
		printList(slots, len(slots), 0, 0, []string{})
		return nil
	})
}

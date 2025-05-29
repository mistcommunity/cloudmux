
package shell

import (
	"time"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type LoganalyticsWorkspaceListOptions struct {
	}
	shellutils.R(&LoganalyticsWorkspaceListOptions{}, "loganalytics-workspace-list", "List loganalytics workspaces", func(cli *azure.SRegion, args *LoganalyticsWorkspaceListOptions) error {
		workspaces, err := cli.GetClient().GetLoganalyticsWorkspaces()
		if err != nil {
			return err
		}
		printList(workspaces, len(workspaces), 0, 0, []string{})
		return nil
	})

	type DiskUsageListOptions struct {
		WORKSPACE_ID string
		NAME         string
		START        time.Time
		END          time.Time
	}

	shellutils.R(&DiskUsageListOptions{}, "disk-usage-list", "List disk usage", func(cli *azure.SRegion, args *DiskUsageListOptions) error {
		usage, err := cli.GetClient().GetInstanceDiskUsage(args.WORKSPACE_ID, args.NAME, args.START, args.END)
		if err != nil {
			return err
		}
		printList(usage, len(usage), 0, 0, []string{})
		return nil
	})

}

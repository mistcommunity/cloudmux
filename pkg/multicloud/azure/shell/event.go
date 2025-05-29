
package shell

import (
	"time"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type EventListOptions struct {
		Start time.Time
		End   time.Time
	}
	shellutils.R(&EventListOptions{}, "event-list", "List events", func(cli *azure.SRegion, args *EventListOptions) error {
		events, err := cli.GetEvents(args.Start, args.End)
		if err != nil {
			return err
		}
		printList(events, len(events), 0, 0, []string{})
		return nil
	})
}

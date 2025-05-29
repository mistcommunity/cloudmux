package shell

import (
	"time"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type EventListOptions struct {
		Start        time.Time
		End          time.Time
		WithReadOnly bool
	}
	shellutils.R(&EventListOptions{}, "event-list", "List events", func(cli *aws.SRegion, args *EventListOptions) error {
		events, err := cli.LookupEvents(args.Start, args.End, args.WithReadOnly)
		if err != nil {
			return err
		}
		printList(events, 0, 0, 0, []string{})
		return nil
	})
}

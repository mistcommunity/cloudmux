
package shell

import (
	"fmt"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"
	"yunion.io/x/pkg/util/timeutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type EventListOptions struct {
		START     string
		END       string
		Token     string
		EventRW   string `choices:"Read|Write|All"`
		RequestId string
	}
	shellutils.R(&EventListOptions{}, "event-list", "List event", func(cli *aliyun.SRegion, args *EventListOptions) error {
		start, err := timeutils.ParseTimeStr(args.START)
		if err != nil {
			return errors.Wrap(err, "timeutils.ParseTimeStr.Start")
		}
		end, err := timeutils.ParseTimeStr(args.END)
		if err != nil {
			return errors.Wrap(err, "timeutils.ParseTimeStr.End")
		}
		events, token, err := cli.GetEvents(start, end, args.Token, args.EventRW, args.RequestId)
		if err != nil {
			return err
		}
		fmt.Printf("token: %s", token)
		printList(events, 0, 0, 0, []string{})
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type ReservationOrderListOptions struct {
	}
	shellutils.R(&ReservationOrderListOptions{}, "reservation-order-list", "List reservation orders", func(cli *azure.SRegion, args *ReservationOrderListOptions) error {
		ret, err := cli.GetClient().ListReservationOrders()
		if err != nil {
			return err
		}
		printList(ret, len(ret), 0, 0, []string{})
		return nil
	})

	type ReservationListOptions struct {
	}
	shellutils.R(&ReservationListOptions{}, "reservation-list", "List reservations", func(cli *azure.SRegion, args *ReservationListOptions) error {
		ret, err := cli.GetClient().ListReservations()
		if err != nil {
			return err
		}
		printList(ret, len(ret), 0, 0, []string{})
		return nil
	})

}

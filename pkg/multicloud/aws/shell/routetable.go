package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type RouteTableListOptions struct {
		VpcId        string `help:"vpc id"`
		SubnetId     string
		RouteTableId string
		VpcPeerId    string
		MainOnly     bool
	}
	shellutils.R(&RouteTableListOptions{}, "route-table-list", "List route tables", func(cli *aws.SRegion, args *RouteTableListOptions) error {
		routetables, err := cli.GetRouteTables(args.VpcId, args.SubnetId, args.VpcPeerId, args.RouteTableId, args.MainOnly)
		if err != nil {
			printObject(err)
			return nil
		}

		printList(routetables, 0, 0, 0, nil)
		return nil
	})

	type RouteCreateOptions struct {
		ROUTETABLEID string `help:"routetable id"`
		CIDRBLOCK    string
		TARGETID     string
	}
	shellutils.R(&RouteCreateOptions{}, "route-create", "create route", func(cli *aws.SRegion, args *RouteCreateOptions) error {
		err := cli.CreateRoute(args.ROUTETABLEID, args.CIDRBLOCK, args.TARGETID)
		if err != nil {
			return err
		}
		return nil
	})

	type RouteReplaceOptions struct {
		ROUTETABLEID string `help:"routetable id"`
		CIDRBLOCK    string
		TARGETID     string
	}
	shellutils.R(&RouteReplaceOptions{}, "route-replace", "replace route", func(cli *aws.SRegion, args *RouteReplaceOptions) error {
		err := cli.ReplaceRoute(args.ROUTETABLEID, args.CIDRBLOCK, args.TARGETID)
		if err != nil {
			return err
		}
		return nil
	})
}

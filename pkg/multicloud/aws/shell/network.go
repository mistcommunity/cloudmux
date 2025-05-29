package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type VSwitchListOptions struct {
		Ids    []string
		ZoneId string
		VpcId  string
	}
	shellutils.R(&VSwitchListOptions{}, "network-list", "List vswitches", func(cli *aws.SRegion, args *VSwitchListOptions) error {
		networks, err := cli.GetNetwroks(args.Ids, args.ZoneId, args.VpcId)
		if err != nil {
			return err
		}
		printList(networks, 0, 0, 0, []string{})
		return nil
	})

	type NetworkCreateOption struct {
		ZoneId string
		VpcId  string
		Name   string
		Desc   string
		CIDR   string
	}

	shellutils.R(&NetworkCreateOption{}, "network-create", "create network", func(cli *aws.SRegion, args *NetworkCreateOption) error {
		network, err := cli.CreateNetwork(args.ZoneId, args.VpcId, args.Name, args.CIDR, args.Desc)
		if err != nil {
			return err
		}
		printObject(network)
		return nil
	})

	type NetworkAttribute struct {
		ID             string
		AssignPublicIp bool
	}

	shellutils.R(&NetworkAttribute{}, "network-assign-public-ip", "assign network public ip", func(cli *aws.SRegion, args *NetworkAttribute) error {
		return cli.ModifySubnetAttribute(args.ID, args.AssignPublicIp)
	})

}

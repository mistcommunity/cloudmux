package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type VpcListOptions struct {
		VpcIds []string
	}
	shellutils.R(&VpcListOptions{}, "vpc-list", "List vpcs", func(cli *aws.SRegion, args *VpcListOptions) error {
		vpcs, err := cli.GetVpcs(args.VpcIds)
		if err != nil {
			return err
		}
		printList(vpcs, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&cloudprovider.VpcCreateOptions{}, "vpc-create", "Create vpc", func(cli *aws.SRegion, args *cloudprovider.VpcCreateOptions) error {
		vpc, err := cli.CreateVpc(args)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

	type VpcIdOptions struct {
		ID string
	}

	shellutils.R(&VpcIdOptions{}, "vpc-delete", "Delete vpc", func(cli *aws.SRegion, args *VpcIdOptions) error {
		return cli.DeleteVpc(args.ID)
	})

	type VpcPeeringConnectionListOptions struct {
		Id        string
		VpcId     string
		PeerVpcId string
	}
	shellutils.R(&VpcPeeringConnectionListOptions{}, "vpc-peering-connection-list", "List vpcPeeringConnections", func(cli *aws.SRegion, args *VpcPeeringConnectionListOptions) error {
		peers, err := cli.DescribeVpcPeeringConnections(args.Id, args.VpcId, args.PeerVpcId)
		if err != nil {
			return err
		}
		printList(peers, 0, 0, 0, []string{})
		return nil
	})

	type VpcPeeringConnectionShowOptions struct {
		ID string
	}
	shellutils.R(&VpcPeeringConnectionShowOptions{}, "vpc-peering-connection-show", "show vpcPeeringConnections", func(cli *aws.SRegion, args *VpcPeeringConnectionShowOptions) error {
		vpcPC, err := cli.GetVpcPeeringConnectionById(args.ID)
		if err != nil {
			return err
		}
		printObject(vpcPC)
		return nil
	})

	type VpcPeeringConnectionCreateOptions struct {
		NAME          string
		VPCID         string
		PEERVPCID     string
		PEERACCOUNTID string
		PEERREGIONID  string
		Desc          string
	}
	shellutils.R(&VpcPeeringConnectionCreateOptions{}, "vpc-peering-connection-create", "create vpcPeeringConnection", func(cli *aws.SRegion, args *VpcPeeringConnectionCreateOptions) error {
		opts := cloudprovider.VpcPeeringConnectionCreateOptions{}
		opts.Desc = args.Desc
		opts.Name = args.NAME
		opts.PeerAccountId = args.PEERACCOUNTID
		opts.PeerRegionId = args.PEERREGIONID
		opts.PeerVpcId = args.PEERVPCID

		vpcPC, err := cli.CreateVpcPeeringConnection(args.VPCID, &opts)
		if err != nil {
			return err
		}
		printObject(vpcPC)
		return nil
	})

	type VpcPeeringConnectionAcceptOptions struct {
		ID string
	}
	shellutils.R(&VpcPeeringConnectionAcceptOptions{}, "vpc-peering-connection-accept", "accept vpcPeeringConnection", func(cli *aws.SRegion, args *VpcPeeringConnectionAcceptOptions) error {
		vpcPC, err := cli.AcceptVpcPeeringConnection(args.ID)
		if err != nil {
			return err
		}
		printObject(vpcPC)
		return nil
	})

	type VpcPeeringConnectionDeleteOptions struct {
		ID string
	}
	shellutils.R(&VpcPeeringConnectionDeleteOptions{}, "vpc-peering-connection-delete", "delete vpcPeeringConnection", func(cli *aws.SRegion, args *VpcPeeringConnectionDeleteOptions) error {
		err := cli.DeleteVpcPeeringConnection(args.ID)
		if err != nil {
			return err
		}
		return nil
	})

	type VpcPeeringConnectionRouteDeleteOptions struct {
		ID string
	}
	shellutils.R(&VpcPeeringConnectionAcceptOptions{}, "vpc-peering-connection-route-delete", "delete vpc-peering-connection route", func(cli *aws.SRegion, args *VpcPeeringConnectionAcceptOptions) error {
		err := cli.DeleteVpcPeeringConnectionRoute(args.ID)
		if err != nil {
			printObject(err)
			return err
		}
		return nil
	})
}

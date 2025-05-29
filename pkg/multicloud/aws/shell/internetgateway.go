package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type InternetGatewayCreateOptions struct {
	}
	shellutils.R(&InternetGatewayCreateOptions{}, "igw-create", "Create igw", func(cli *aws.SRegion, args *InternetGatewayCreateOptions) error {
		ret, err := cli.CreateIgw()
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	type InternetGatewayListOptions struct {
		VpcId string
	}

	shellutils.R(&InternetGatewayListOptions{}, "igw-list", "List igw", func(cli *aws.SRegion, args *InternetGatewayListOptions) error {
		igws, err := cli.GetInternetGateways(args.VpcId)
		if err != nil {
			return err
		}
		printList(igws, 0, 0, 0, nil)
		return nil
	})

	type InternetGatewayIdOptions struct {
		ID string
	}

	shellutils.R(&InternetGatewayIdOptions{}, "igw-delete", "Delete igw", func(cli *aws.SRegion, args *InternetGatewayIdOptions) error {
		return cli.DeleteInternetGateway(args.ID)
	})

	type InternetGatewayDetachOptions struct {
		VPC_ID string
		IGW_ID string
	}

	shellutils.R(&InternetGatewayDetachOptions{}, "igw-detach", "Detach igw", func(cli *aws.SRegion, args *InternetGatewayDetachOptions) error {
		return cli.DetachInternetGateway(args.VPC_ID, args.IGW_ID)
	})

}

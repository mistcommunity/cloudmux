package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type NatListOptions struct {
		Ids      []string
		VpcId    string
		SubnetId string
	}
	shellutils.R(&NatListOptions{}, "natgateway-list", "List natgateway", func(cli *aws.SRegion, args *NatListOptions) error {
		nats, err := cli.GetNatGateways(args.Ids, args.VpcId, args.SubnetId)
		if err != nil {
			return err
		}
		printList(nats, 0, 0, 0, nil)
		return nil
	})

	type NatIdOptions struct {
		ID string
	}

	shellutils.R(&NatIdOptions{}, "natgateway-delete", "Delete natgateway", func(cli *aws.SRegion, args *NatIdOptions) error {
		return cli.DeleteNatgateway(args.ID)
	})

}

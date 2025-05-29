package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type LbListOptions struct {
		Id     string
		Marker string
	}
	shellutils.R(&LbListOptions{}, "lb-list", "List loadbalancer", func(cli *aws.SRegion, args *LbListOptions) error {
		ret, _, e := cli.GetLoadbalancers(args.Id, args.Marker)
		if e != nil {
			return e
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type LbIdOptions struct {
		ID string
	}
	shellutils.R(&LbIdOptions{}, "lb-show", "Show loadbalancer attribute", func(cli *aws.SRegion, args *LbIdOptions) error {
		ret, err := cli.GetElbAttributes(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&cloudprovider.SLoadbalancerCreateOptions{}, "lb-create", "Create loadbalancer", func(cli *aws.SRegion, args *cloudprovider.SLoadbalancerCreateOptions) error {
		ret, err := cli.CreateLoadbalancer(args)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&LbIdOptions{}, "lb-delete", "Delete loadbalancer", func(cli *aws.SRegion, args *LbIdOptions) error {
		return cli.DeleteElb(args.ID)
	})

	shellutils.R(&LbIdOptions{}, "lb-tag-list", "Show loadbalancer tags", func(cli *aws.SRegion, args *LbIdOptions) error {
		ret, err := cli.DescribeElbTags(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	type LbBackendGroupListOptions struct {
		ElbId string
		Id    string
	}

	shellutils.R(&LbBackendGroupListOptions{}, "lb-backend-group-list", "List loadbalancer backend groups", func(cli *aws.SRegion, args *LbBackendGroupListOptions) error {
		ret, err := cli.GetElbBackendgroups(args.ElbId, args.Id)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&cloudprovider.SLoadbalancerCertificate{}, "lb-cert-create", "Create loadbalancer cert", func(cli *aws.SRegion, args *cloudprovider.SLoadbalancerCertificate) error {
		arn, err := cli.CreateLoadbalancerCertifacate(args)
		if err != nil {
			return err
		}
		fmt.Println(arn)
		return nil
	})

	type LbCertListOption struct {
	}

	shellutils.R(&LbCertListOption{}, "lb-cert-list", "Create loadbalancer cert", func(cli *aws.SRegion, args *LbCertListOption) error {
		certs, err := cli.ListServerCertificates()
		if err != nil {
			return err
		}
		printList(certs, 0, 0, 0, nil)
		return nil
	})

}

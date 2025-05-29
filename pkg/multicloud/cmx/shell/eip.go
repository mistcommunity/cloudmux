package shell

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func init() {
	cmd := NewCommand("eip")

	type EipListOptions struct {
		ListBaseOptions
		Id          string `help:"Eip id"`
		AssociateId string `help:"Id of associate resource"`
		Addr        string `help:"Eip address"`
	}
	RegionR[EipListOptions](cmd).List("list", "List eips", func(cli cloudprovider.ICloudRegion, args *EipListOptions) (any, error) {
		return cli.GetIEips()
	})

	// type EipAllocateOptions struct {
	// 	Name            string
	// 	BW              int    `help:"Bandwidth limit in Mbps"`
	// 	ResourceGroupId string `help:"Resource group Id"`
	// }
	// RegionRunner[EipAllocateOptions](cmd).Run("create", "Allocate an EIP", func(cli cloudprovider.ICloudRegion, args *EipAllocateOptions) (any, error) {
	// 	eip, err := cli.AllocateEIP(args.Name, args.BW, aliyun.InternetChargeByTraffic, args.ResourceGroupId)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	PrintObject(eip)
	// 	return nil
	// })

	// type EipReleaseOptions struct {
	// 	ID string `help:"EIP allocation ID"`
	// }
	// RegionRunner[EipReleaseOptions](cmd).Run("delete", "Release an EIP", func(cli cloudprovider.ICloudRegion, args *EipReleaseOptions) (any, error) {
	// 	err := cli.DeallocateEIP(args.ID)
	// 	return nil, err
	// })
	//
	// type EipAssociateOptions struct {
	// 	ID       string `help:"EIP allocation ID"`
	// 	INSTANCE string `help:"Instance ID"`
	// }
	// RegionR(&EipAssociateOptions{}, "eip-associate", "Associate an EIP", func(cli cloudprovider.ICloudRegion, args *EipAssociateOptions) error {
	// 	err := cli.AssociateEip(args.ID, args.INSTANCE)
	// 	return err
	// })
	// RegionR(&EipAssociateOptions{}, "eip-dissociate", "Dissociate an EIP", func(cli cloudprovider.ICloudRegion, args *EipAssociateOptions) error {
	// 	err := cli.DissociateEip(args.ID, args.INSTANCE)
	// 	return err
	// })
}

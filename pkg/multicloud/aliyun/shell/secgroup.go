
package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type SecurityGroupListOptions struct {
		VpcId            string   `help:"VPC ID"`
		Name             string   `help:"Secgroup Name"`
		SecurityGroupIds []string `help:"SecurityGroup ids"`
	}
	shellutils.R(&SecurityGroupListOptions{}, "security-group-list", "List security group", func(cli *aliyun.SRegion, args *SecurityGroupListOptions) error {
		secgrps, err := cli.GetSecurityGroups(args.VpcId, args.Name, args.SecurityGroupIds)
		if err != nil {
			return err
		}
		printList(secgrps, 0, 0, 0, []string{})
		return nil
	})

	type SecurityGroupIdOptions struct {
		ID string `help:"ID or name of security group"`
	}
	shellutils.R(&SecurityGroupIdOptions{}, "security-group-rule-list", "Show details of a security group", func(cli *aliyun.SRegion, args *SecurityGroupIdOptions) error {
		rules, err := cli.GetSecurityGroupRules(args.ID)
		if err != nil {
			return err
		}
		printList(rules, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&SecurityGroupIdOptions{}, "security-group-references", "Show references of a security group", func(cli *aliyun.SRegion, args *SecurityGroupIdOptions) error {
		references, err := cli.DescribeSecurityGroupReferences(args.ID)
		if err != nil {
			return err
		}
		printList(references, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&cloudprovider.SecurityGroupCreateInput{}, "security-group-create", "Create details of a security group", func(cli *aliyun.SRegion, args *cloudprovider.SecurityGroupCreateInput) error {
		secgroupId, err := cli.CreateSecurityGroup(args)
		if err != nil {
			return err
		}
		fmt.Printf("secgroupId: %s", secgroupId)
		return nil
	})

}

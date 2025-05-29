package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type SecurityGroupListOptions struct {
		VpcId string `help:"VPC ID"`
		Id    string
		Name  string `help:"Secgroup name"`
	}
	shellutils.R(&SecurityGroupListOptions{}, "security-group-list", "List security group", func(cli *aws.SRegion, args *SecurityGroupListOptions) error {
		secgrps, err := cli.GetSecurityGroups(args.VpcId, args.Name, args.Id)
		if err != nil {
			return err
		}
		printList(secgrps, 0, 0, 0, []string{})
		return nil
	})

	type SecurityGroupIdOptions struct {
		ID string
	}

	shellutils.R(&SecurityGroupIdOptions{}, "security-group-delete", "Show security group", func(cli *aws.SRegion, args *SecurityGroupIdOptions) error {
		return cli.DeleteSecurityGroup(args.ID)
	})

	shellutils.R(&SecurityGroupIdOptions{}, "security-group-rule-list", "Show security group rules", func(cli *aws.SRegion, args *SecurityGroupIdOptions) error {
		rules, err := cli.GetSecurityGroupRules(args.ID)
		if err != nil {
			return err
		}
		printList(rules, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&SecurityGroupIdOptions{}, "security-group-show", "Show security group", func(cli *aws.SRegion, args *SecurityGroupIdOptions) error {
		group, err := cli.GetSecurityGroup(args.ID)
		if err != nil {
			return err
		}
		printObject(group)
		rules, err := group.GetRules()
		if err != nil {
			return err
		}
		printList(rules, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&cloudprovider.SecurityGroupCreateInput{}, "security-group-create", "Create  security group", func(cli *aws.SRegion, args *cloudprovider.SecurityGroupCreateInput) error {
		id, err := cli.CreateSecurityGroup(args)
		if err != nil {
			return err
		}
		fmt.Println(id)
		return nil
	})

	type SecurityGroupRuleCreateOptions struct {
		GROUP_ID string
		cloudprovider.SecurityGroupRuleCreateOptions
	}

	shellutils.R(&SecurityGroupRuleCreateOptions{}, "security-group-rule-create", "Create security group rule", func(cli *aws.SRegion, args *SecurityGroupRuleCreateOptions) error {
		rule, err := cli.CreateSecurityGroupRule(args.GROUP_ID, &args.SecurityGroupRuleCreateOptions)
		if err != nil {
			return err
		}
		printObject(rule)
		return nil
	})

	type SecurityGroupRuleDeleteOptions struct {
		GROUP_ID  string
		ID        string
		Direction string `choices:"in|out" default:"in"`
	}

	shellutils.R(&SecurityGroupRuleDeleteOptions{}, "security-group-rule-delete", "Delete security group rule", func(cli *aws.SRegion, args *SecurityGroupRuleDeleteOptions) error {
		return cli.DeleteSecurityGroupRule(args.GROUP_ID, args.Direction, args.ID)
	})

}

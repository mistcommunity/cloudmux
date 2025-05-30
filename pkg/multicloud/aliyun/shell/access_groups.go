
package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type NasAccessGroupListOptions struct {
		FileSystemType string `choices:"standard|extreme" default:"standard"`
	}
	shellutils.R(&NasAccessGroupListOptions{}, "access-group-list", "List Nas AccessGroups", func(cli *aliyun.SRegion, args *NasAccessGroupListOptions) error {
		ags, err := cli.GetAccessGroups(args.FileSystemType)
		if err != nil {
			return err
		}
		printList(ags, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&cloudprovider.SAccessGroup{}, "access-group-create", "Create Nas AccessGroup", func(cli *aliyun.SRegion, args *cloudprovider.SAccessGroup) error {
		return cli.CreateAccessGroup(args)
	})

	type AccessGroupDeleteOptions struct {
		FileSystemType string
		NAME           string
	}

	shellutils.R(&AccessGroupDeleteOptions{}, "access-group-delete", "Delete AccessGroup", func(cli *aliyun.SRegion, args *AccessGroupDeleteOptions) error {
		return cli.DeleteAccessGroup(args.FileSystemType, args.NAME)
	})

	type NasAccessGroupRuleListOptions struct {
		GROUP string
	}

	shellutils.R(&NasAccessGroupRuleListOptions{}, "access-group-rule-list", "List Nas AccessGroup Rules", func(cli *aliyun.SRegion, args *NasAccessGroupRuleListOptions) error {
		rules, err := cli.GetAccessGroupRules(args.GROUP)
		if err != nil {
			return err
		}
		printList(rules, 0, 0, 0, []string{})
		return nil
	})

	type AccessRuleDeleteOptions struct {
		GROUP   string
		RULE_ID string
	}

	shellutils.R(&AccessRuleDeleteOptions{}, "access-group-rule-delete", "Delete AccessGroup Rule", func(cli *aliyun.SRegion, args *AccessRuleDeleteOptions) error {
		return cli.DeleteAccessGroupRule(args.GROUP, args.RULE_ID)
	})

	type AccessRuleCreateOptions struct {
		SOURCE         string
		FileSystemType string
		GroupName      string
		RwType         cloudprovider.TRWAccessType
		UserType       cloudprovider.TUserAccessType
		Priority       int
	}

	shellutils.R(&AccessRuleCreateOptions{}, "access-group-rule-create", "Delete AccessGroup Rule", func(cli *aliyun.SRegion, args *AccessRuleCreateOptions) error {
		ruleId, err := cli.CreateAccessGroupRule(args.SOURCE, args.FileSystemType, args.GroupName, args.RwType, args.UserType, args.Priority)
		if err != nil {
			return err
		}
		fmt.Println(ruleId)
		return nil
	})

}

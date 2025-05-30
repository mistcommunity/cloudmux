
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type CloudgroupCreateOptions struct {
		NAME     string
		Comments string
	}

	shellutils.R(&CloudgroupCreateOptions{}, "cloud-group-create", "Create Cloud group", func(cli *aliyun.SRegion, args *CloudgroupCreateOptions) error {
		group, err := cli.GetClient().CreateGroup(args.NAME, args.Comments)
		if err != nil {
			return err
		}
		printObject(group)
		return nil
	})

	type CloudgroupListOptions struct {
		Offset string
		Limit  int
	}

	shellutils.R(&CloudgroupListOptions{}, "cloud-group-list", "List Cloud groups", func(cli *aliyun.SRegion, args *CloudgroupListOptions) error {
		groups, err := cli.GetClient().ListGroups(args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(groups.Groups.Group, 0, 0, 0, nil)
		return nil
	})

	type CloudgroupDeleteOptions struct {
		NAME string
	}

	shellutils.R(&CloudgroupDeleteOptions{}, "cloud-group-delete", "Delete Cloud group", func(cli *aliyun.SRegion, args *CloudgroupDeleteOptions) error {
		return cli.GetClient().DeleteGroup(args.NAME)
	})

	type GroupExtListOptions struct {
		GROUP  string
		Offset string
		Limit  int
	}
	shellutils.R(&GroupExtListOptions{}, "cloud-group-user-list", "List Cloud group users", func(cli *aliyun.SRegion, args *GroupExtListOptions) error {
		users, err := cli.GetClient().ListUsersForGroup(args.GROUP, args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(users.Users.User, 0, 0, 0, nil)
		return nil
	})

	type GroupUserOptions struct {
		GROUP string
		USER  string
	}

	shellutils.R(&GroupUserOptions{}, "cloud-group-remove-user", "Remove user from group", func(cli *aliyun.SRegion, args *GroupUserOptions) error {
		return cli.GetClient().RemoveUserFromGroup(args.GROUP, args.USER)
	})

	shellutils.R(&GroupUserOptions{}, "cloud-group-add-user", "Add user to group", func(cli *aliyun.SRegion, args *GroupUserOptions) error {
		return cli.GetClient().AddUserToGroup(args.GROUP, args.USER)
	})

	shellutils.R(&GroupExtListOptions{}, "cloud-group-policy-list", "List Cloud group policies", func(cli *aliyun.SRegion, args *GroupExtListOptions) error {
		policies, err := cli.GetClient().ListPoliciesForGroup(args.GROUP)
		if err != nil {
			return err
		}
		printList(policies, 0, 0, 0, nil)
		return nil
	})

	type GroupPolicyOptions struct {
		GROUP      string
		PolicyType string `default:"System" choices:"System|Custom"`
		POLICY     string
	}

	shellutils.R(&GroupPolicyOptions{}, "cloud-group-attach-policy", "Attach policy for group", func(cli *aliyun.SRegion, args *GroupPolicyOptions) error {
		return cli.GetClient().AttachPolicyToGroup(args.PolicyType, args.POLICY, args.GROUP)
	})

	shellutils.R(&GroupPolicyOptions{}, "cloud-group-detach-policy", "Detach policy from group", func(cli *aliyun.SRegion, args *GroupPolicyOptions) error {
		return cli.GetClient().DetachPolicyFromGroup(args.PolicyType, args.POLICY, args.GROUP)
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ClouduserCreateOptions struct {
		NAME        string
		MobilePhone string
		Comments    string
		Email       string
	}

	shellutils.R(&ClouduserCreateOptions{}, "cloud-user-create", "Create Cloud user", func(cli *aliyun.SRegion, args *ClouduserCreateOptions) error {
		user, err := cli.GetClient().CreateUser(args.NAME, args.MobilePhone, args.Email, args.Comments)
		if err != nil {
			return err
		}
		printObject(user)
		return nil
	})

	type ClouduserListOptions struct {
		Offset string
		Limit  int
	}

	shellutils.R(&ClouduserListOptions{}, "cloud-user-list", "List Cloud users", func(cli *aliyun.SRegion, args *ClouduserListOptions) error {
		users, err := cli.GetClient().ListUsers(args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(users.Users.User, 0, 0, 0, nil)
		return nil
	})

	type UserPolicyListOptions struct {
		USER string
	}

	shellutils.R(&UserPolicyListOptions{}, "cloud-user-policy-list", "List Cloud user policies", func(cli *aliyun.SRegion, args *UserPolicyListOptions) error {
		policies, err := cli.GetClient().ListPoliciesForUser(args.USER)
		if err != nil {
			return err
		}
		printList(policies, 0, 0, 0, nil)
		return nil
	})

	type ClouduserOptions struct {
		NAME string
	}

	shellutils.R(&ClouduserOptions{}, "cloud-user-group-list", "List Cloud user groups", func(cli *aliyun.SRegion, args *ClouduserOptions) error {
		groups, err := cli.GetClient().ListGroupsForUser(args.NAME)
		if err != nil {
			return err
		}
		printList(groups, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&ClouduserOptions{}, "cloud-user-delete", "Delete Cloud user", func(cli *aliyun.SRegion, args *ClouduserOptions) error {
		return cli.GetClient().DeleteClouduser(args.NAME)
	})

	shellutils.R(&ClouduserOptions{}, "cloud-user-loginprofile", "Get Cloud user loginprofile", func(cli *aliyun.SRegion, args *ClouduserOptions) error {
		profile, err := cli.GetClient().GetLoginProfile(args.NAME)
		if err != nil {
			return err
		}
		printObject(profile)
		return nil
	})

	shellutils.R(&ClouduserOptions{}, "cloud-user-loginprofile-delete", "Delete Cloud user loginprofile", func(cli *aliyun.SRegion, args *ClouduserOptions) error {
		return cli.GetClient().DeleteLoginProfile(args.NAME)
	})

	type LoginProfileCreateOptions struct {
		NAME     string
		PASSWORD string
	}

	shellutils.R(&LoginProfileCreateOptions{}, "cloud-user-loginprofile-create", "Create Cloud user loginprofile", func(cli *aliyun.SRegion, args *LoginProfileCreateOptions) error {
		profile, err := cli.GetClient().CreateLoginProfile(args.NAME, args.PASSWORD, false, true)
		if err != nil {
			return err
		}
		printObject(profile)
		return nil
	})

	shellutils.R(&LoginProfileCreateOptions{}, "cloud-user-rest-password", "Reset Cloud user password", func(cli *aliyun.SRegion, args *LoginProfileCreateOptions) error {
		return cli.GetClient().ResetClouduserPassword(args.NAME, args.PASSWORD)
	})

	type ClouduserPolicyOptions struct {
		PolicyType string `default:"System" choices:"System|Custom"`
		POLICY     string
		USER       string
	}

	shellutils.R(&ClouduserPolicyOptions{}, "cloud-user-attach-policy", "Attach policy for user", func(cli *aliyun.SRegion, args *ClouduserPolicyOptions) error {
		return cli.GetClient().AttachPolicyToUser(args.POLICY, args.PolicyType, args.USER)
	})

	shellutils.R(&ClouduserPolicyOptions{}, "cloud-user-detach-policy", "Detach policy from user", func(cli *aliyun.SRegion, args *ClouduserPolicyOptions) error {
		return cli.GetClient().DetachPolicyFromUser(args.POLICY, args.PolicyType, args.USER)
	})

}

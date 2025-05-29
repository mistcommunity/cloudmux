
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ListRolesOptions struct {
		Offset string
		Limit  int
	}
	shellutils.R(&ListRolesOptions{}, "cloud-role-list", "List ram roles", func(cli *aliyun.SRegion, args *ListRolesOptions) error {
		roles, err := cli.GetClient().ListRoles(args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(roles.Roles.Role, 0, 0, 0, []string{})
		return nil
	})

	type GetRoleOptions struct {
		ROLENAME string
	}
	shellutils.R(&GetRoleOptions{}, "cloud-role-show", "Show ram role", func(cli *aliyun.SRegion, args *GetRoleOptions) error {
		role, err := cli.GetClient().GetRole(args.ROLENAME)
		if err != nil {
			return err
		}
		printObject(role)
		return nil
	})

	type RoleCreateOptions struct {
		NAME     string
		DOCUMENT string
		Desc     string
	}

	shellutils.R(&RoleCreateOptions{}, "cloud-role-create", "Create ram role", func(cli *aliyun.SRegion, args *RoleCreateOptions) error {
		role, err := cli.GetClient().CreateRole(args.NAME, args.DOCUMENT, args.Desc)
		if err != nil {
			return err
		}
		printObject(role)
		return nil
	})

	type RolePolicyOptions struct {
		ROLENAME   string
		POLICYNAME string
		POLICYTYPE string `choices:"Custom|System"`
	}

	shellutils.R(&RolePolicyOptions{}, "cloud-role-attach-policy", "Attach policy for role", func(cli *aliyun.SRegion, args *RolePolicyOptions) error {
		return cli.GetClient().AttachPolicy2Role(args.POLICYTYPE, args.POLICYNAME, args.ROLENAME)
	})

	shellutils.R(&RolePolicyOptions{}, "cloud-role-detach-policy", "Detach policy from role", func(cli *aliyun.SRegion, args *RolePolicyOptions) error {
		return cli.GetClient().DetachPolicyFromRole(args.POLICYTYPE, args.POLICYNAME, args.ROLENAME)
	})

	type RolePolicyListOptions struct {
		ROLE string
	}

	shellutils.R(&RolePolicyListOptions{}, "cloud-role-policy-list", "List cloud role policies", func(cli *aliyun.SRegion, args *RolePolicyListOptions) error {
		policies, err := cli.GetClient().ListPoliciesForRole(args.ROLE)
		if err != nil {
			return err
		}
		printList(policies, 0, 0, 0, nil)
		return nil
	})

	type DeleteRoleOptions struct {
		NAME string
	}
	shellutils.R(&DeleteRoleOptions{}, "cloud-role-delete", "Delete role", func(cli *aliyun.SRegion, args *DeleteRoleOptions) error {
		return cli.GetClient().DeleteRole(args.NAME)
	})

	shellutils.R(&ListRolesOptions{}, "enable-image-import", "Enable image import privilege", func(cli *aliyun.SRegion, args *ListRolesOptions) error {
		return cli.GetClient().EnableImageImport()
	})

	shellutils.R(&ListRolesOptions{}, "enable-image-export", "Enable image export privilege", func(cli *aliyun.SRegion, args *ListRolesOptions) error {
		return cli.GetClient().EnableImageExport()
	})

	type CallerShowOptions struct {
	}

	shellutils.R(&CallerShowOptions{}, "caller-show", "Show caller info", func(cli *aliyun.SRegion, args *CallerShowOptions) error {
		caller, err := cli.GetClient().GetCallerIdentity()
		if err != nil {
			return err
		}
		printObject(caller)
		return nil
	})

}

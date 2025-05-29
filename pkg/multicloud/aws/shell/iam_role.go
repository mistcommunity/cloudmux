package shell

import (
	"fmt"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type RoleListOptions struct {
		Offset     string
		Limit      int
		PathPrefix string
	}
	shellutils.R(&RoleListOptions{}, "cloud-role-list", "List roles", func(cli *aws.SRegion, args *RoleListOptions) error {
		roles, err := cli.GetClient().ListRoles(args.Offset, args.Limit, args.PathPrefix)
		if err != nil {
			return err
		}
		printList(roles.Roles, 0, 0, 0, []string{})
		return nil
	})

	type RoleNameOptions struct {
		ROLE string
	}

	shellutils.R(&RoleNameOptions{}, "cloud-role-show", "Show role", func(cli *aws.SRegion, args *RoleNameOptions) error {
		role, err := cli.GetClient().GetRole(args.ROLE)
		if err != nil {
			return err
		}
		printObject(role)
		document := role.GetDocument()
		if document != nil {
			printObject(document)
		}
		return nil
	})

	type RoleAttachPolicyListOptions struct {
		ROLE       string
		Marker     string
		MaxItems   int
		PathPrefix string
	}

	shellutils.R(&RoleAttachPolicyListOptions{}, "cloud-role-attach-policy-list", "List Role attach policy", func(cli *aws.SRegion, args *RoleAttachPolicyListOptions) error {
		policy, err := cli.GetClient().ListAttachedRolePolicies(args.ROLE, args.Marker, args.MaxItems, args.PathPrefix)
		if err != nil {
			return errors.Wrapf(err, "ListAttachedRolePolicies")
		}
		printList(policy.AttachedPolicies, 0, 0, 0, nil)
		if len(policy.Marker) > 0 {
			fmt.Println("marker: ", policy.Marker)
		}
		return nil
	})

	shellutils.R(&RoleNameOptions{}, "cloud-role-delete", "Delete role", func(cli *aws.SRegion, args *RoleNameOptions) error {
		return cli.GetClient().DeleteRole(args.ROLE)
	})

	shellutils.R(&cloudprovider.SRoleCreateOptions{}, "cloud-role-create", "Create role", func(cli *aws.SRegion, args *cloudprovider.SRoleCreateOptions) error {
		role, err := cli.GetClient().CreateRole(args)
		if err != nil {
			return err
		}
		printObject(role)
		return nil
	})

	type RoleAttachPolicyOptions struct {
		ROLE   string
		POLICY string
	}

	shellutils.R(&RoleAttachPolicyOptions{}, "cloud-role-attach-policy", "Attach policy for role", func(cli *aws.SRegion, args *RoleAttachPolicyOptions) error {
		return cli.GetClient().AttachRolePolicy(args.ROLE, args.POLICY)
	})

}

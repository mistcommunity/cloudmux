
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type CloudpolicyListOptions struct {
		Name string
	}
	shellutils.R(&CloudpolicyListOptions{}, "cloud-policy-list", "List cloudpolicies", func(cli *azure.SRegion, args *CloudpolicyListOptions) error {
		roles, err := cli.GetClient().GetRoles(args.Name)
		if err != nil {
			return err
		}
		printList(roles, 0, 0, 0, nil)
		return nil
	})

	type CloudpolicyAssignOption struct {
		OBJECT string
		ROLE   string
	}

	shellutils.R(&CloudpolicyAssignOption{}, "cloud-policy-assign-object", "Assign cloudpolicy for object", func(cli *azure.SRegion, args *CloudpolicyAssignOption) error {
		return cli.GetClient().AssignPolicy(args.OBJECT, args.ROLE)
	})

	type AssignmentListOption struct {
		ObjectId string
	}

	type CloudpolicyAssignListOptions struct {
		ID string
	}

	shellutils.R(&CloudpolicyAssignListOptions{}, "cloud-user-policy-list", "Assign cloudpolicy for object", func(cli *azure.SRegion, args *CloudpolicyAssignListOptions) error {
		ret, err := cli.GetClient().GetPrincipalPolicy(args.ID)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&CloudpolicyAssignListOptions{}, "cloud-group-policy-list", "Assign cloudpolicy for object", func(cli *azure.SRegion, args *CloudpolicyAssignListOptions) error {
		ret, err := cli.GetClient().GetPrincipalPolicy(args.ID)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, nil)
		return nil
	})

}

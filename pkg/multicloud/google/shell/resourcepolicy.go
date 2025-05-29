
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type ResourcePolicyListOptions struct {
		Disk       string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&ResourcePolicyListOptions{}, "resource-policy-list", "List resourcepolicys", func(cli *google.SRegion, args *ResourcePolicyListOptions) error {
		resourcepolicys, err := cli.GetResourcePolicies(args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(resourcepolicys, 0, 0, 0, nil)
		return nil
	})

	type ResourcePolicyShowOptions struct {
		ID string
	}
	shellutils.R(&ResourcePolicyShowOptions{}, "resource-policy-show", "Show resourcepolicy", func(cli *google.SRegion, args *ResourcePolicyShowOptions) error {
		resourcepolicy, err := cli.GetResourcePolicy(args.ID)
		if err != nil {
			return err
		}
		printObject(resourcepolicy)
		return nil
	})

	type ListOrganizationOptions struct{}
	shellutils.R(&ListOrganizationOptions{}, "organization-list", "List organizaitons", func(cli *google.SRegion, args *ListOrganizationOptions) error {
		orgs, err := cli.GetClient().ListOrganizations()
		if err != nil {
			return err
		}
		printList(orgs, 0, 0, 0, nil)
		return nil
	})

}


package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type GetPolicyOptions struct {
		POLICYTYPE string
		POLICYNAME string
	}
	shellutils.R(&GetPolicyOptions{}, "cloud-policy-show", "Show ram policy", func(cli *aliyun.SRegion, args *GetPolicyOptions) error {
		policy, err := cli.GetClient().GetPolicy(args.POLICYTYPE, args.POLICYNAME)
		if err != nil {
			return err
		}
		printObject(policy)
		return nil
	})

	type DeletePolicyOptions struct {
		POLICYTYPE string
		POLICYNAME string
	}
	shellutils.R(&DeletePolicyOptions{}, "cloud-policy-delete", "Delete policy", func(cli *aliyun.SRegion, args *DeletePolicyOptions) error {
		return cli.GetClient().DeletePolicy(args.POLICYTYPE, args.POLICYNAME)
	})

	type PolicyListOptions struct {
		PolicyType string `choices:"System|Custom"`
		Offset     string
		Limit      int
	}

	shellutils.R(&PolicyListOptions{}, "cloud-policy-list", "List cloud policies", func(cli *aliyun.SRegion, args *PolicyListOptions) error {
		policies, err := cli.GetClient().ListPolicies(args.PolicyType, args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(policies.Policies.Policy, 0, 0, 0, nil)
		return nil
	})

	type PolicyCreateOptions struct {
		NAME     string
		DOCUMENT string
		Desc     string
	}

	shellutils.R(&PolicyCreateOptions{}, "cloud-policy-create", "Create ram policy", func(cli *aliyun.SRegion, args *PolicyCreateOptions) error {
		policy, err := cli.GetClient().CreatePolicy(args.NAME, args.DOCUMENT, args.Desc)
		if err != nil {
			return err
		}
		printObject(policy)
		return nil
	})

	type PolicyCreateVersionOptions struct {
		NAME      string
		DOCUMENT  string
		IsDefault bool
	}

	shellutils.R(&PolicyCreateVersionOptions{}, "cloud-policy-version-create", "Create ram policy version", func(cli *aliyun.SRegion, args *PolicyCreateVersionOptions) error {
		return cli.GetClient().CreatePolicyVersion(args.NAME, args.DOCUMENT, args.IsDefault)
	})

}

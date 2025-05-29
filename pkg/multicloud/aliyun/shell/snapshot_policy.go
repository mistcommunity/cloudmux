
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type SSnapshotPolicyListOptions struct {
		PolicyId string `help:"snapshot policy id"`
	}

	shellutils.R(&SSnapshotPolicyListOptions{}, "snapshot-policy-list", "list snapshot policy",
		func(cli *aliyun.SRegion, args *SSnapshotPolicyListOptions) error {
			snapshotPolicis, err := cli.GetSnapshotPolicies(args.PolicyId)
			if err != nil {
				return err
			}
			printList(snapshotPolicis, 0, 0, 0, []string{})
			return nil
		},
	)

	type SSnapshotPolicyDeleteOptions struct {
		ID string `help:"snapshot id"`
	}
	shellutils.R(&SSnapshotPolicyDeleteOptions{}, "snapshot-policy-delete", "delete snapshot policy",
		func(cli *aliyun.SRegion, args *SSnapshotPolicyDeleteOptions) error {
			err := cli.DeleteSnapshotPolicy(args.ID)
			return err
		},
	)

	type SSnapshotPolicyCreateOptions struct {
		Name string `help:"snapshot name"`

		RetentionDays  int   `help:"retention days"`
		RepeatWeekdays []int `help:"auto snapshot which days of the week"`
		TimePoints     []int `help:"auto snapshot which hours of the day"`
	}
	shellutils.R(&SSnapshotPolicyCreateOptions{}, "snapshot-policy-create", "create snapshot policy",
		func(cli *aliyun.SRegion, args *SSnapshotPolicyCreateOptions) error {
			input := cloudprovider.SnapshotPolicyInput{
				RetentionDays:  args.RetentionDays,
				RepeatWeekdays: args.RepeatWeekdays,
				TimePoints:     args.TimePoints,
				Name:           args.Name,
			}
			_, err := cli.CreateSnapshotPolicy(&input)
			if err != nil {
				return err
			}
			return nil
		},
	)

	type SSnapshotPolicyApplyOptions struct {
		SNAPSHOTPOLICYID string   `help:"snapshot policy id"`
		DISKID           []string `help:"disk id"`
	}
	shellutils.R(&SSnapshotPolicyApplyOptions{}, "snapshot-policy-apply", "apply snapshot policy",
		func(cli *aliyun.SRegion, args *SSnapshotPolicyApplyOptions) error {
			err := cli.ApplySnapshotPolicyToDisks(args.SNAPSHOTPOLICYID, args.DISKID)
			return err
		},
	)

	type SSnapshotPolicyCancelOptions struct {
		SNAPSHOTPOLICYID string   `help:"snapshot policy id"`
		DISKID           []string `help:"disk id"`
	}
	shellutils.R(&SSnapshotPolicyCancelOptions{}, "snapshot-policy-cancel", "cancel snapshot policy",
		func(cli *aliyun.SRegion, args *SSnapshotPolicyCancelOptions) error {
			err := cli.CancelSnapshotPolicyToDisks(args.SNAPSHOTPOLICYID, args.DISKID)
			return err
		},
	)
}

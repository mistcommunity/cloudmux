
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type LoadbalancerACLListOptions struct {
	}
	shellutils.R(&LoadbalancerACLListOptions{}, "lb-acl-list", "List loadbalanceAcls", func(cli *aliyun.SRegion, args *LoadbalancerACLListOptions) error {
		acls, err := cli.GetLoadBalancerAcls()
		if err != nil {
			return err
		}
		printList(acls, len(acls), 0, 0, []string{})
		return nil
	})
}

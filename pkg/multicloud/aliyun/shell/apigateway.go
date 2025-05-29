
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type ApigatewayListOptions struct {
	}
	shellutils.R(&ApigatewayListOptions{}, "apigateway-list", "list apigateways", func(cli *aliyun.SRegion, args *ApigatewayListOptions) error {
		ret, err := cli.GetApigateways()
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

}

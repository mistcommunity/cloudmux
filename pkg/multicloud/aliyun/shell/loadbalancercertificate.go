
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type LoadbalancerServerCertificateListOptions struct {
	}
	shellutils.R(&LoadbalancerServerCertificateListOptions{}, "lb-server-certificate-list", "List ServerCertificates", func(cli *aliyun.SRegion, args *LoadbalancerServerCertificateListOptions) error {
		serverCertificate, err := cli.GetLoadbalancerServerCertificates()
		if err != nil {
			return err
		}
		printList(serverCertificate, len(serverCertificate), 0, 0, []string{})
		return nil
	})
}

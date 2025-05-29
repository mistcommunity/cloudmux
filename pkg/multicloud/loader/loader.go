package loader

import (
	"yunion.io/x/log" // on-premise virtualization technologies

	_ "yunion.io/x/cloudmux/pkg/multicloud/aliyun/provider"
	_ "yunion.io/x/cloudmux/pkg/multicloud/aws/provider"
	_ "yunion.io/x/cloudmux/pkg/multicloud/azure/provider"
	_ "yunion.io/x/cloudmux/pkg/multicloud/esxi/provider"   // private clouds
	_ "yunion.io/x/cloudmux/pkg/multicloud/google/provider" // public clouds
	_ "yunion.io/x/cloudmux/pkg/multicloud/nutanix/provider" // private clouds
	_ "yunion.io/x/cloudmux/pkg/multicloud/proxmox/provider" // private clouds
)

func init() {
	log.Infof("Loading cloud providers ...")
}

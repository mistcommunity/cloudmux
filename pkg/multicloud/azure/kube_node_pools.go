
package azure

import (
	"strings"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SKubeNodePool struct {
	multicloud.SResourceBase
	AzureTags

	cluster *SKubeCluster

	Name              string `json:"name"`
	Count             int    `json:"count"`
	VMSize            string `json:"vmSize"`
	OsDiskSizeGB      int    `json:"osDiskSizeGB"`
	OsDiskType        string `json:"osDiskType"`
	KubeletDiskType   string `json:"kubeletDiskType"`
	MaxPods           int    `json:"maxPods"`
	Type              string `json:"type"`
	EnableAutoScaling bool   `json:"enableAutoScaling"`
	ProvisioningState string `json:"provisioningState"`
	PowerState        struct {
		Code string `json:"code"`
	} `json:"powerState"`
	OrchestratorVersion string `json:"orchestratorVersion"`
	Mode                string `json:"mode"`
	OsType              string `json:"osType"`
	OsSKU               string `json:"osSKU"`
	NodeImageVersion    string `json:"nodeImageVersion"`
	EnableFIPS          bool   `json:"enableFIPS"`
}

func (self *SKubeNodePool) GetName() string {
	return self.Name
}

func (self *SKubeNodePool) GetId() string {
	return self.Name
}

func (self *SKubeNodePool) GetGlobalId() string {
	return self.Name
}

func (self *SKubeNodePool) GetStatus() string {
	return strings.ToLower(self.PowerState.Code)
}

func (self *SKubeNodePool) GetMinInstanceCount() int {
	return 0
}

func (self *SKubeNodePool) GetMaxInstanceCount() int {
	return 0
}

func (self *SKubeNodePool) GetDesiredInstanceCount() int {
	return 0
}

func (self *SKubeNodePool) GetRootDiskSizeGb() int {
	return self.OsDiskSizeGB
}

func (self *SKubeNodePool) Delete() error {
	return cloudprovider.ErrNotImplemented
}

func (self *SKubeNodePool) GetInstanceTypes() []string {
	return []string{}
}

func (self *SKubeNodePool) GetNetworkIds() []string {
	return []string{}
}

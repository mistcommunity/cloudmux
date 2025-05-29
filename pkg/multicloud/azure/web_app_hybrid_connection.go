
package azure

import (
	"fmt"
	"strings"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SAppHybirdConnection struct {
	Id         string
	Name       string
	Type       string
	Properties struct {
		ServiceBusNamespace string
		RelayName           string
		Hostname            string
		Port                int
	}
}

func (self *SAppHybirdConnection) GetGlobalId() string {
	return strings.ToLower(self.Id)
}

func (self *SAppHybirdConnection) GetName() string {
	return self.Name
}

func (self *SAppHybirdConnection) GetHostname() string {
	return self.Properties.Hostname
}

func (self *SAppHybirdConnection) GetNamespace() string {
	return self.Properties.ServiceBusNamespace
}

func (self *SAppHybirdConnection) GetPort() int {
	return self.Properties.Port
}

func (self *SRegion) GetAppHybirConnections(appId string) ([]SAppHybirdConnection, error) {
	res := fmt.Sprintf("%s/hybridConnectionRelays", appId)
	resp, err := self.list_v2(res, "2023-12-01", nil)
	if err != nil {
		return nil, err
	}
	ret := []SAppHybirdConnection{}
	err = resp.Unmarshal(&ret, "value")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *SAppSite) GetHybirdConnections() ([]cloudprovider.IAppHybirdConnection, error) {
	connections, err := self.region.GetAppHybirConnections(self.Id)
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.IAppHybirdConnection{}
	for i := range connections {
		ret = append(ret, &connections[i])
	}
	return ret, nil
}

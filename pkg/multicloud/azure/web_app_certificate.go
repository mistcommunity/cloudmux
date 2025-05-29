
package azure

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SAppCertificate struct {
	Id         string
	Name       string
	Type       string
	Properties struct {
		SubjectName    string
		Issuer         string
		IssueDate      time.Time
		Thumbprint     string
		ExpirationDate time.Time
	}
}

func (self *SAppCertificate) GetGlobalId() string {
	return strings.ToLower(self.Id)
}

func (self *SAppCertificate) GetName() string {
	return self.Name
}

func (self *SAppCertificate) GetSubjectName() string {
	return self.Properties.SubjectName
}

func (self *SAppCertificate) GetIssuer() string {
	return self.Properties.Issuer
}

func (self *SAppCertificate) GetIssueDate() time.Time {
	return self.Properties.IssueDate
}

func (self *SAppCertificate) GetThumbprint() string {
	return self.Properties.Thumbprint
}

func (self *SAppCertificate) GetExpireTime() time.Time {
	return self.Properties.ExpirationDate
}

func (self *SRegion) GetAppCertificates(farmId string) ([]SAppCertificate, error) {
	res := fmt.Sprintf("Microsoft.Web/certificates")
	params := url.Values{}
	params.Set("$filter", fmt.Sprintf("ServerFarmId eq '%s'", farmId))
	resp, err := self.list_v2(res, "2023-12-01", params)
	if err != nil {
		return nil, err
	}
	ret := []SAppCertificate{}
	err = resp.Unmarshal(&ret, "value")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *SAppSite) GetCertificates() ([]cloudprovider.IAppCertificate, error) {
	properties, err := self.GetProperties()
	if err != nil {
		return nil, err
	}
	certs, err := self.region.GetAppCertificates(properties.ServerFarmId)
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.IAppCertificate{}
	for i := range certs {
		ret = append(ret, &certs[i])
	}
	return ret, nil
}


package google

import (
	"time"

	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
)

type SOrganizationOwner struct {
	DirectoryCustomerId string `json:"directoryCustomerId"`
}

type SOrganization struct {
	OrganizationId string             `json:"organizationId"`
	DisplayName    string             `json:"displayName"`
	Owner          SOrganizationOwner `json:"owner"`
	CreationTime   time.Time          `json:"creationTime"`
	LifecycleState string             `json:"lifecycleState"`
	Name           string             `json:"name"`
}

// https://cloud.google.com/resource-manager/reference/rest/v1/organizations/search
// require Organization Viewer privilege
func (self *SGoogleClient) ListOrganizations() ([]SOrganization, error) {
	resource := "organizations"
	params := map[string]string{
		"pageSize": "1000",
	}
	resp, err := jsonRequest(self.client, "GET", GOOGLE_MANAGER_DOMAIN, "v1beta1", resource, params, nil, self.debug)
	if err != nil {
		return nil, errors.Wrap(err, "ListOrganizations")
	}
	log.Debugf("ListOrganization: %s", resp)
	ret := make([]SOrganization, 0)
	err = resp.Unmarshal(&ret, "organizations")
	if err != nil {
		return nil, errors.Wrap(err, "resp.Unmarshal")
	}
	return ret, nil
}

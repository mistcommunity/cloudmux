
package google

import (
	"fmt"
	"net/url"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SSqlserverUserDetails struct {
	ServerRoles []string
}

type SDBInstanceAccount struct {
	multicloud.SDBInstanceAccountBase
	GoogleTags
	rds *SDBInstance

	Kind                 string
	Etag                 string
	Name                 string
	Host                 string
	Instance             string
	SelfLink             string
	Project              string
	SqlserverUserDetails SSqlserverUserDetails
}

func (region *SRegion) GetDBInstanceAccounts(instance string) ([]SDBInstanceAccount, error) {
	accounts := []SDBInstanceAccount{}
	params := map[string]string{}
	resource := fmt.Sprintf("instances/%s/users", instance)
	err := region.RdsListAll(resource, params, &accounts)
	if err != nil {
		return nil, errors.Wrap(err, "RdsListAll")
	}
	return accounts, nil
}

func (region *SRegion) GetDBInstanceAccount(id string) (*SDBInstanceAccount, error) {
	account := &SDBInstanceAccount{}
	err := region.rdsGet(id, account)
	if err != nil {
		return nil, errors.Wrap(err, "rdsGet")
	}
	return account, nil
}

func (region *SRegion) DeleteDBInstanceAccount(instanceName, user, host string) error {
	resource := fmt.Sprintf("projects/%s/instances/%s/users?name=%s&host=%s", region.GetProjectId(), instanceName, url.PathEscape(user), url.PathEscape(host))
	return region.rdsDelete(resource)
}

func (account *SDBInstanceAccount) Delete() error {
	return account.rds.region.DeleteDBInstanceAccount(account.rds.Name, account.Name, account.Host)
}

func (account *SDBInstanceAccount) GetHost() string {
	return account.Host
}

func (account *SDBInstanceAccount) GetName() string {
	return account.Name
}

func (account *SDBInstanceAccount) Refresh() error {
	_account, err := account.rds.region.GetDBInstanceAccount(account.SelfLink)
	if err != nil {
		return errors.Wrap(err, "GetDBInstanceAccount")
	}
	return jsonutils.Update(account, _account)
}

func (account *SDBInstanceAccount) GetIDBInstanceAccountPrivileges() ([]cloudprovider.ICloudDBInstanceAccountPrivilege, error) {
	return []cloudprovider.ICloudDBInstanceAccountPrivilege{}, nil
}

func (account *SDBInstanceAccount) ResetPassword(password string) error {
	params := map[string]string{
		"host": account.Host,
		"name": account.Name,
	}
	resource := fmt.Sprintf("instances/%s/users", account.rds.Name)
	body := map[string]string{
		"password": password,
	}
	return account.rds.region.rdsUpdate(resource, params, jsonutils.Marshal(body))
}

func (account *SDBInstanceAccount) GrantPrivilege(database, privilege string) error {
	return cloudprovider.ErrNotSupported
}

func (account *SDBInstanceAccount) RevokePrivilege(database string) error {
	return cloudprovider.ErrNotSupported
}

func (region *SRegion) CreateDBInstanceAccount(instanceId string, name, password, host string) error {
	body := map[string]interface{}{
		"name":     name,
		"password": password,
	}
	if len(host) > 0 {
		body["host"] = host
	}
	return region.rdsDo(instanceId, "users", nil, jsonutils.Marshal(body))
}

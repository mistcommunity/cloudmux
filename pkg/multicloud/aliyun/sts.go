
package aliyun

import (
	"yunion.io/x/jsonutils"
)

func (self *SAliyunClient) stsRequest(apiName string, params map[string]string) (jsonutils.JSONObject, error) {
	cli, err := self.getDefaultClient()
	if err != nil {
		return nil, err
	}
	return jsonRequest(cli, "sts.aliyuncs.com", ALIYUN_STS_API_VERSION, apiName, params, self.debug)
}

type SCallerIdentity struct {
	Arn          string
	AccountId    string
	UserId       string
	RoleId       string
	PrincipalId  string
	IdentityType string
}

func (self *SAliyunClient) GetCallerIdentity() (*SCallerIdentity, error) {
	params := map[string]string{}
	resp, err := self.stsRequest("GetCallerIdentity", params)
	if err != nil {
		return nil, err
	}
	id := &SCallerIdentity{}
	return id, resp.Unmarshal(id)
}

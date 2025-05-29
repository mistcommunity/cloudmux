
package aliyun

import (
	"fmt"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
)

type SAliyunAccount struct {
	Status                string
	JoinMethod            string
	ModifyTime            time.Time
	Type                  string
	ResourceDirectoryId   string
	AccountId             string
	DisplayName           string
	JoinTime              time.Time
	FolderId              string
	ResourceDirectoryPath string
}

func (self *SAliyunClient) rdRequest(apiName string, params map[string]string) (jsonutils.JSONObject, error) {
	cli, err := self.getDefaultClient()
	if err != nil {
		return nil, err
	}
	return jsonRequest(cli, "resourcedirectory.aliyuncs.com", ALIYUN_RD_API_VERSION, apiName, params, self.debug)
}

func (self *SAliyunClient) ListAccounts() ([]SAliyunAccount, error) {
	params := map[string]string{
		"IncludeTags": "true",
		"PageSize":    "100",
	}
	pageNumber := 1
	ret := []SAliyunAccount{}
	for {
		resp, err := self.rdRequest("ListAccounts", params)
		if err != nil {
			return nil, errors.Wrapf(err, "ListAccounts")
		}
		part := struct {
			Accounts struct {
				Account []SAliyunAccount
			}
			TotalCount int
		}{}
		err = resp.Unmarshal(&part)
		if err != nil {
			return nil, errors.Wrapf(err, "Unmarshal")
		}
		ret = append(ret, part.Accounts.Account...)
		if len(part.Accounts.Account) == 0 || len(ret) >= part.TotalCount {
			break
		}
		pageNumber++
		params["PageNumber"] = fmt.Sprintf("%d", pageNumber)
	}
	return ret, nil
}

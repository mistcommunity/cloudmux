
package aliyun

import (
	"yunion.io/x/jsonutils"
)

func (self *SAliyunClient) ramRequest(apiName string, params map[string]string) (jsonutils.JSONObject, error) {
	cli, err := self.getDefaultClient()
	if err != nil {
		return nil, err
	}
	return jsonRequest(cli, "ram.aliyuncs.com", ALIYUN_RAM_API_VERSION, apiName, params, self.debug)
}

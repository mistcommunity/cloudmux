
package aliyun

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
)

type jsonRequestFunc func(action string, params map[string]string) (jsonutils.JSONObject, error)

func unmarshalResult(resp jsonutils.JSONObject, respErr error, resultKey []string, result interface{}) error {
	if respErr != nil {
		return respErr
	}

	if result == nil {
		return nil
	}

	if resultKey != nil && len(resultKey) > 0 {
		respErr = resp.Unmarshal(result, resultKey...)
	} else {
		respErr = resp.Unmarshal(result)
	}

	if respErr != nil {
		log.Errorf("unmarshal json error %s", respErr)
	}

	return nil
}

// 执行操作
func DoAction(client jsonRequestFunc, action string, params map[string]string, resultKey []string, result interface{}) error {
	resp, err := client(action, params)
	return unmarshalResult(resp, err, resultKey, result)
}

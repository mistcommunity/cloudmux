
package aliyun

import (
	"runtime/debug"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
)

func processCommonRequest(client *sdk.Client, req *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("client.ProcessCommonRequest error: %s", r)
			debug.PrintStack()
			response = nil
			jsonError := jsonutils.NewDict()
			jsonError.Add(jsonutils.NewString("SignatureNonceUsed"), "Code")
			err = errors.Error(jsonError.String())
		}
	}()
	return client.ProcessCommonRequest(req)
}

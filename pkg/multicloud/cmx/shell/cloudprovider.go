package shell

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/errors"
)

func init() {
	r := EmptyOptionProviderR("cloudprovider")

	r.Run("system-info", "Get cloudprovider system info", func(cli cloudprovider.ICloudProvider) (any, error) {
		return cli.GetSysInfo()
	})

	r.Run("api-version", "Get cloudprovider api version", func(cli cloudprovider.ICloudProvider) (any, error) {
		version := cli.GetVersion()
		return map[string]string{
			"api_version": version,
		}, nil
	})

	r.Run("iam-login-url", "Get IAM login URL", func(cli cloudprovider.ICloudProvider) (any, error) {
		url := cli.GetIamLoginUrl()
		return map[string]string{
			"url": url,
		}, nil
	})

	r.List("region-list", "List regions of a cloudprovider", func(cli cloudprovider.ICloudProvider) (any, error) {
		return cli.GetIRegions()
	})

	r.List("subaccount-list", "List subaccounts of a cloudprovider", func(cli cloudprovider.ICloudProvider) (any, error) {
		accounts, err := cli.GetSubAccounts()
		if err != nil {
			return nil, errors.Wrap(err, "GetSubAccounts")
		}
		return accounts, nil
	})
}

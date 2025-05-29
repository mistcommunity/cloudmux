package shell

import "yunion.io/x/cloudmux/pkg/cloudprovider"

func init() {
	r := EmptyOptionProviderR("balance")

	r.Run("show", "Show balance", func(cli cloudprovider.ICloudProvider) (any, error) {
		balance, err := cli.GetBalance()
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"available_amount": balance.Amount,
			"status":           balance.Status,
		}, nil
	})
}


package shell

import (
	"yunion.io/x/log"
	"yunion.io/x/pkg/util/printutils"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/esxi"
)

func init() {
	shellutils.R(&cloudprovider.MetricListOptions{}, "metric-list", "List metrics in a namespace", func(cli *esxi.SESXiClient, args *cloudprovider.MetricListOptions) error {
		metrics, err := cli.GetMetrics(args)
		if err != nil {
			return err
		}
		for i := range metrics {
			log.Infof("metric %s %s", metrics[i].Id, metrics[i].MetricType)
			printutils.PrintInterfaceList(metrics[i].Values, 0, 0, 0, nil)
		}
		return nil
	})

	type SMetricTypeShowOptions struct {
	}

	shellutils.R(&SMetricTypeShowOptions{}, "metric-type-list", "List metrics", func(cli *esxi.SESXiClient, args *SMetricTypeShowOptions) error {
		metrics, err := cli.GetMetricTypes()
		if err != nil {
			return err
		}
		printutils.PrintInterfaceList(metrics, 0, 0, 0, nil)
		return nil
	})

}

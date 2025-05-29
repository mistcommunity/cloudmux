package shell

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type ZoneListOptions struct {
	ListBaseOptions
	// ChargeType   string `help:"charge type" choices:"PrePaid|PostPaid" default:"PrePaid"`
	// SpotStrategy string `help:"Spot strategy, NoSpot|SpotWithPriceLimit|SpotAsPriceGo" choices:"NoSpot|SpotWithPriceLimit|SpotAsPriceGo" default:"NoSpot"`
}

func (o ZoneListOptions) GetColumns() []string {
	return []string{"name", "zone_id", "local_name", "available_resource_creation", "available_disk_categories"}
}

func init() {
	cmd := NewCommand("zone")

	NewCO[ZoneListOptions](cmd).UseList().RunByRegion("list", "List zones", func(region cloudprovider.ICloudRegion, _ *ZoneListOptions) (any, error) {
		zones, err := region.GetIZones()
		if err != nil {
			return nil, err
		}
		return zones, nil
	})

}

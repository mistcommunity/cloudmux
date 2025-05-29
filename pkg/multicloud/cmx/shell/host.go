package shell

import "yunion.io/x/cloudmux/pkg/cloudprovider"

func init() {
	cmd := NewCommand("host")

	type HostListOptions struct {
		ListBaseOptions
		ZoneBaseOptions
	}

	ZoneR[HostListOptions](cmd).List("list", "List hosts", func(cli cloudprovider.ICloudZone, args *HostListOptions) (any, error) {
		hosts, err := cli.GetIHosts()
		if err != nil {
			return nil, err
		}
		objs := make([]interface{}, len(hosts))
		for i := range hosts {
			host := hosts[i]
			objs[i] = map[string]interface{}{
				"id":          host.GetId(),
				"global_id":   host.GetGlobalId(),
				"name":        host.GetName(),
				"is_emulated": host.IsEmulated(),
				"status":      host.GetStatus(),
				"host_status": host.GetHostStatus(),
			}
		}
		return objs, nil
	})
}


package esxi

import (
	"fmt"
)

func (wire *sWire) GetSysTags() map[string]string {
	tags := make(map[string]string)
	dc, _ := wire.network.GetDatacenter()
	if dc != nil {
		tags["datacenter"] = dc.GetName()
	}
	ips, macs, _ := wire.getAvailableIpsMacs()
	if len(ips) > 0 {
		tags["vm_ips"] = compactIPs(ips)
	}
	if len(macs) > 0 {
		tags["vm_macs"] = compactMacs(macs)
	}
	paths := wire.network.GetPath()
	for i := 3; i < len(paths); i++ {
		tags[fmt.Sprintf("folder_%d", i-3)] = paths[i]
	}
	return tags
}

func (host *SHost) GetSysTags() map[string]string {
	tags := make(map[string]string)
	dc, _ := host.GetDatacenter()
	if dc != nil {
		tags["datacenter"] = dc.GetName()
	}
	cluster, _ := host.GetCluster()
	if cluster != nil {
		tags["cluster"] = cluster.GetName()
	}
	resourcePool, _ := host.GetResourcePool()
	if resourcePool != nil {
		tags["resource_pool"] = resourcePool.Name()
	}
	paths := host.GetPath()
	for i := 3; i < len(paths); i++ {
		tags[fmt.Sprintf("folder_%d", i-3)] = paths[i]
	}
	return tags
}

func (svm *SVirtualMachine) GetSysTags() map[string]string {
	meta := map[string]string{}
	dc, _ := svm.GetDatacenter()
	if dc != nil {
		meta["datacenter"] = dc.GetName()
	}
	paths := svm.GetPath()
	for i := 3; i < len(paths); i++ {
		meta[fmt.Sprintf("folder_%d", i-3)] = paths[i]
	}
	meta["networks"] = svm.getNetTags()
	// meta["datacenter"] = svm.GetDatacenterPathString()
	rp, _ := svm.getResourcePool()
	if rp != nil {
		rpPath := rp.GetPath()
		rpOffset := -1
		for i := range rpPath {
			if rpPath[i] == "Resources" {
				if i > 0 {
					meta["cluster"] = rpPath[i-1]
					rpOffset = i
				}
			} else if rpOffset >= 0 && i > rpOffset {
				meta[fmt.Sprintf("pool%d", i-rpOffset-1)] = rpPath[i]
			}
		}
	}
	return meta
}


package esxi

import (
	"net/url"
	"strings"

	"github.com/vmware/govmomi/vim25/mo"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

var RESOURCEPOOL_PROPS = []string{"name", "parent", "host"}

type SResourcePool struct {
	multicloud.SProjectBase
	multicloud.STagBase
	SManagedObject
}

func (pool *SResourcePool) GetGlobalId() string {
	return pool.GetId()
}

func (pool *SResourcePool) GetStatus() string {
	return api.EXTERNAL_PROJECT_STATUS_AVAILABLE
}

func (pool *SResourcePool) GetName() string {
	path := pool.GetPath()
	if len(path) > 5 {
		path = path[5:]
	}
	name := []string{}
	for _, _name := range path {
		p, _ := url.PathUnescape(_name)
		name = append([]string{p}, name...)
	}
	return strings.Join(name, "/")
}

func NewResourcePool(manager *SESXiClient, rp *mo.ResourcePool, dc *SDatacenter) *SResourcePool {
	return &SResourcePool{SManagedObject: newManagedObject(manager, rp, dc)}
}

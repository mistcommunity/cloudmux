
package google

import (
	"fmt"
	"strings"
)

type SResourceBase struct {
	Name     string
	SelfLink string
	Id       string
}

func (r *SResourceBase) GetId() string {
	if len(r.Id) > 0 {
		return r.Id
	}
	return r.SelfLink
}

func getGlobalId(selfLink string) string {
	return strings.TrimPrefix(selfLink, fmt.Sprintf("%s/%s/", GOOGLE_COMPUTE_DOMAIN, GOOGLE_API_VERSION))
}

func (r *SResourceBase) GetGlobalId() string {
	if len(r.Id) > 0 {
		return r.Id
	}
	return getGlobalId(r.SelfLink)
}

func (r *SResourceBase) GetName() string {
	return r.Name
}

func (r *SResourceBase) GetDescription() string {
	return ""
}

func (r *SResourceBase) IsEmulated() bool {
	return false
}

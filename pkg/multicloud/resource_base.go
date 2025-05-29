
package multicloud

import "time"

type SResourceBase struct {
}

func (resource *SResourceBase) IsEmulated() bool {
	return false
}

func (resource *SResourceBase) Refresh() error {
	return nil
}

func (resource *SResourceBase) GetCreatedAt() time.Time {
	return time.Time{}
}

func (resource *SResourceBase) GetDescription() string {
	return ""
}

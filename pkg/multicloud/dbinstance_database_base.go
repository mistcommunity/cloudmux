
package multicloud

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SDBInstanceDatabaseBase struct {
	SResourceBase
}

func (db *SDBInstanceDatabaseBase) Delete() error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "Delete")
}

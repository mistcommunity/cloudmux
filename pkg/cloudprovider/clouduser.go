package cloudprovider

import (
	"time"

	"yunion.io/x/jsonutils"
)

type SClouduserCreateConfig struct {
	Name                  string
	Desc                  string
	Password              string
	IsConsoleLogin        bool
	Email                 string
	MobilePhone           string
	UserType              string
	EnableMfa             bool
	PasswordResetRequired bool
}

type SCloudpolicyPermission struct {
	Name     string
	Action   string
	Category string
}

type SCloudpolicyCreateOptions struct {
	Name     string
	Desc     string
	Document *jsonutils.JSONDict
}

type SAccessKey struct {
	Name      string
	AccessKey string
	Secret    string
	Status    string
	CreatedAt time.Time
}

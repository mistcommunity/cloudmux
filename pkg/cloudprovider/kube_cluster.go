package cloudprovider

import (
	"time"
)

type SKubeconfig struct {
	Config     string    `json:"config"`
	Expiration time.Time `json:"expiration"`
}

type KubeClusterCreateOptions struct {
	NAME          string
	Desc          string
	VpcId         string
	Version       string
	NetworkIds    []string
	Tags          map[string]string
	ServiceCIDR   string
	PrivateAccess bool
	PublicAccess  bool
	RoleName      string

	PublicKey string
}

type KubeNodePoolCreateOptions struct {
	NAME string
	Desc string

	MinInstanceCount     int
	MaxInstanceCount     int
	DesiredInstanceCount int

	RootDiskSizeGb int

	PublicKey string

	InstanceTypes []string
	NetworkIds    []string

	Tags map[string]string
}

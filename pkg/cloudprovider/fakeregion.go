package cloudprovider

import "yunion.io/x/pkg/errors"

type SFakeOnPremiseRegion struct {
}

func (region *SFakeOnPremiseRegion) GetId() string {
	return "default"
}

func (region *SFakeOnPremiseRegion) GetName() string {
	return "Default"
}

func (region *SFakeOnPremiseRegion) GetGlobalId() string {
	return "default"
}

func (region *SFakeOnPremiseRegion) GetStatus() string {
	return "available"
}

func (region *SFakeOnPremiseRegion) GetCloudEnv() string {
	return ""
}

func (region *SFakeOnPremiseRegion) Refresh() error {
	return nil
}

func (region *SFakeOnPremiseRegion) IsEmulated() bool {
	return true
}

func (region *SFakeOnPremiseRegion) GetSysTags() map[string]string {
	return nil
}

func (region *SFakeOnPremiseRegion) GetTags() (map[string]string, error) {
	return nil, errors.Wrap(ErrNotImplemented, "GetTags")
}

func (region *SFakeOnPremiseRegion) SetTags(tags map[string]string, replace bool) error {
	return ErrNotImplemented
}

func (region *SFakeOnPremiseRegion) GetGeographicInfo() SGeographicInfo {
	return SGeographicInfo{}
}

func (region *SFakeOnPremiseRegion) GetIZones() ([]ICloudZone, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) GetIZoneById(id string) (ICloudZone, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) GetIVpcById(id string) (ICloudVpc, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) GetIVpcs() ([]ICloudVpc, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) GetIEips() ([]ICloudEIP, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) GetIEipById(id string) (ICloudEIP, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) CreateIVpc(opts *VpcCreateOptions) (ICloudVpc, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) CreateEIP(eip *SEip) (ICloudEIP, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) GetISecurityGroupById(id string) (ICloudSecurityGroup, error) {
	return nil, ErrNotSupported
}

func (region *SFakeOnPremiseRegion) CreateISecurityGroup(conf *SecurityGroupCreateInput) (ICloudSecurityGroup, error) {
	return nil, ErrNotSupported
}

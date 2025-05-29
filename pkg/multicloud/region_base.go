package multicloud

import (
	"time"

	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SRegion struct {
	SResourceBase
	STagBase
}

func (r *SRegion) GetIDiskById(id string) (cloudprovider.ICloudDisk, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIDiskById")
}

func (r *SRegion) GetIHostById(id string) (cloudprovider.ICloudHost, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIHostById")
}

func (r *SRegion) GetIHosts() ([]cloudprovider.ICloudHost, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIHosts")
}

func (r *SRegion) GetISnapshotById(snapshotId string) (cloudprovider.ICloudSnapshot, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISnapshotById")
}

func (r *SRegion) GetISnapshots() ([]cloudprovider.ICloudSnapshot, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISnapshots")
}

func (r *SRegion) GetISecurityGroups() ([]cloudprovider.ICloudSecurityGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISecurityGroups")
}

func (r *SRegion) GetIStorageById(id string) (cloudprovider.ICloudStorage, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIStorageById")
}

func (r *SRegion) GetIStoragecacheById(id string) (cloudprovider.ICloudStoragecache, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIStoragecacheById")
}

func (r *SRegion) GetIStoragecaches() ([]cloudprovider.ICloudStoragecache, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIStoragecaches")
}

func (r *SRegion) GetIStorages() ([]cloudprovider.ICloudStorage, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIStorages")
}

func (r *SRegion) GetIVMs() ([]cloudprovider.ICloudVM, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIVMs")
}

func (r *SRegion) GetIVMById(id string) (cloudprovider.ICloudVM, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIVMById")
}

func (r *SRegion) CreateSnapshotPolicy(input *cloudprovider.SnapshotPolicyInput) (string, error) {
	return "", errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateSnapshotPolicy")
}

func (r *SRegion) GetISnapshotPolicyById(snapshotPolicyId string) (cloudprovider.ICloudSnapshotPolicy, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISnapshotPolicyById")
}

func (self *SRegion) GetISnapshotPolicies() ([]cloudprovider.ICloudSnapshotPolicy, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISnapshotPolicies")
}

func (self *SRegion) GetISkus() ([]cloudprovider.ICloudSku, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotSupported, "GetISkus")
}

func (self *SRegion) CreateISku(opts *cloudprovider.SServerSkuCreateOption) (cloudprovider.ICloudSku, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateISku")
}

func (self *SRegion) GetINetworkInterfaces() ([]cloudprovider.ICloudNetworkInterface, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetINetworkInterfaces")
}

func (self *SRegion) GetICloudEvents(start time.Time, end time.Time, withReadEvent bool) ([]cloudprovider.ICloudEvent, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudEvents")
}

func (self *SRegion) GetICloudQuotas() ([]cloudprovider.ICloudQuota, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudQuotas")
}

func (self *SRegion) CreateInternetGateway() (cloudprovider.ICloudInternetGateway, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotSupported, "CreateInternetGateway")
}

func (self *SRegion) GetICloudFileSystems() ([]cloudprovider.ICloudFileSystem, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudFileSystems")
}

func (self *SRegion) GetICloudFileSystemById(id string) (cloudprovider.ICloudFileSystem, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudFileSystemById")
}

func (self *SRegion) GetICloudAccessGroups() ([]cloudprovider.ICloudAccessGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudAccessGroups")
}

func (self *SRegion) GetICloudAccessGroupById(id string) (cloudprovider.ICloudAccessGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudAccessGroupById")
}

func (self *SRegion) CreateICloudAccessGroup(opts *cloudprovider.SAccessGroup) (cloudprovider.ICloudAccessGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateICloudAccessGroup")
}

func (self *SRegion) CreateICloudFileSystem(opts *cloudprovider.FileSystemCraeteOptions) (cloudprovider.ICloudFileSystem, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateICloudFileSystem")
}

func (self *SRegion) GetICloudWafIPSets() ([]cloudprovider.ICloudWafIPSet, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudWafIPSets")
}

func (self *SRegion) GetICloudWafRegexSets() ([]cloudprovider.ICloudWafRegexSet, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudWafRegexSets")
}

func (self *SRegion) GetICloudWafInstances() ([]cloudprovider.ICloudWafInstance, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudWafInstances")
}

func (self *SRegion) GetICloudWafInstanceById(id string) (cloudprovider.ICloudWafInstance, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudWafInstanceById")
}

func (self *SRegion) CreateICloudWafInstance(opts *cloudprovider.WafCreateOptions) (cloudprovider.ICloudWafInstance, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateICloudWafInstance")
}

func (self *SRegion) GetICloudWafRuleGroups() ([]cloudprovider.ICloudWafRuleGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudWafRuleGroups")
}

func (self *SRegion) GetICloudMongoDBs() ([]cloudprovider.ICloudMongoDB, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudMongoDBs")
}

func (self *SRegion) GetICloudMongoDBById(id string) (cloudprovider.ICloudMongoDB, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudMongoDBById")
}

func (self *SRegion) GetIElasticSearchs() ([]cloudprovider.ICloudElasticSearch, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIElasticSearchs")
}

func (self *SRegion) GetIElasticSearchById(id string) (cloudprovider.ICloudElasticSearch, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIElasticSearchById")
}

func (self *SRegion) GetICloudKafkas() ([]cloudprovider.ICloudKafka, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudKafkas")
}

func (self *SRegion) GetICloudKafkaById(id string) (cloudprovider.ICloudKafka, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudKafkaById")
}

func (self *SRegion) GetICloudApps() ([]cloudprovider.ICloudApp, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudApps")
}

func (self *SRegion) GetICloudAppById(id string) (cloudprovider.ICloudApp, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudAppById")
}

func (self *SRegion) GetICloudNatSkus() ([]cloudprovider.ICloudNatSku, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudNatSkus")
}

func (self *SRegion) GetICloudTablestores() ([]cloudprovider.ICloudTablestore, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudTablestores")
}

type SRegionZoneBase struct {
}

func (self *SRegionZoneBase) GetIZones() ([]cloudprovider.ICloudZone, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIZones")
}

func (self *SRegionZoneBase) GetIZoneById(id string) (cloudprovider.ICloudZone, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIZoneById")
}

type SRegionVpcBase struct {
}

func (self *SRegionVpcBase) CreateIVpc(opts *cloudprovider.VpcCreateOptions) (cloudprovider.ICloudVpc, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateIVpc")
}

func (self *SRegionVpcBase) GetIVpcs() ([]cloudprovider.ICloudVpc, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIVpcs")
}

func (self *SRegionVpcBase) GetIVpcById(id string) (cloudprovider.ICloudVpc, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIVpcById")
}

type SRegionOssBase struct {
}

type SRegionSecurityGroupBase struct {
}

func (self *SRegionSecurityGroupBase) CreateISecurityGroup(conf *cloudprovider.SecurityGroupCreateInput) (cloudprovider.ICloudSecurityGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateISecurityGroup")
}

func (self *SRegionSecurityGroupBase) GetISecurityGroupById(secgroupId string) (cloudprovider.ICloudSecurityGroup, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISecurityGroupById")
}

type SRegionEipBase struct {
}

func (self *SRegionEipBase) GetIEipById(id string) (cloudprovider.ICloudEIP, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIEipById")
}

func (self *SRegionEipBase) GetIEips() ([]cloudprovider.ICloudEIP, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIEips")
}

func (self *SRegionEipBase) CreateEIP(eip *cloudprovider.SEip) (cloudprovider.ICloudEIP, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateEIP")
}

func (self *SRegion) GetIModelartsPools() ([]cloudprovider.ICloudModelartsPool, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIModelartsPools")
}

func (self *SRegion) GetIModelartsPoolById(id string) (cloudprovider.ICloudModelartsPool, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIModelartsPoolDetail")
}

func (self *SRegion) CreateIModelartsPool(pool *cloudprovider.ModelartsPoolCreateOption, callback func(id string)) (cloudprovider.ICloudModelartsPool, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateIModelartsPool")
}

func (self *SRegion) GetStatusMessage() string {
	return ""
}

func (self *SRegion) GetIModelartsPoolSku() ([]cloudprovider.ICloudModelartsPoolSku, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIModelartsPoolSku")
}

func (self *SRegion) GetIMiscResources() ([]cloudprovider.ICloudMiscResource, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetIMiscResources")
}

func (self *SRegion) GetISSLCertificates() ([]cloudprovider.ICloudSSLCertificate, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetISSLCertificate")
}

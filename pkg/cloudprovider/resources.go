package cloudprovider

import (
	"context"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/tristate"
	"yunion.io/x/pkg/util/billing"
	"yunion.io/x/pkg/util/rbacscope"
	"yunion.io/x/pkg/util/samlutils"
	"yunion.io/x/pkg/util/secrules"

	api "yunion.io/x/cloudmux/pkg/apis/cloudid"
)

type ICloudResource interface {
	GetId() string
	GetName() string
	GetGlobalId() string
	GetCreatedAt() time.Time
	GetDescription() string

	GetStatus() string

	Refresh() error

	IsEmulated() bool

	GetSysTags() map[string]string
	GetTags() (map[string]string, error)
	SetTags(tags map[string]string, replace bool) error
}

type ICloudEnabledResource interface {
	ICloudResource
	GetEnabled() bool
}

type IVirtualResource interface {
	ICloudResource

	GetProjectId() string
}

type IBillingResource interface {
	GetBillingType() string
	GetExpiredAt() time.Time
	SetAutoRenew(bc billing.SBillingCycle) error
	Renew(bc billing.SBillingCycle) error
	ChangeBillingType(billType string) error
	IsAutoRenew() bool
}

type ICloudI18nResource interface {
	GetI18n() SModelI18nTable
}

type ICloudRegion interface {
	ICloudResource
	ICloudI18nResource

	GetGeographicInfo() SGeographicInfo

	GetIZones() ([]ICloudZone, error)
	GetIVpcs() ([]ICloudVpc, error)
	GetIEips() ([]ICloudEIP, error)
	GetIVpcById(id string) (ICloudVpc, error)
	GetIZoneById(id string) (ICloudZone, error)
	GetIEipById(id string) (ICloudEIP, error)
	// ICoudVM 的 GetGlobalId 接口不能panic
	GetIVMs() ([]ICloudVM, error)
	// Esxi没有zone，需要通过region确认vm是否被删除
	GetIVMById(id string) (ICloudVM, error)
	GetIDiskById(id string) (ICloudDisk, error)

	// 仅返回region级别的安全组, vpc下面的安全组需要在ICloudVpc底下返回
	GetISecurityGroups() ([]ICloudSecurityGroup, error)
	GetISecurityGroupById(secgroupId string) (ICloudSecurityGroup, error)
	CreateISecurityGroup(opts *SecurityGroupCreateInput) (ICloudSecurityGroup, error)

	CreateIVpc(opts *VpcCreateOptions) (ICloudVpc, error)
	CreateInternetGateway() (ICloudInternetGateway, error)
	CreateEIP(eip *SEip) (ICloudEIP, error)

	GetISnapshots() ([]ICloudSnapshot, error)
	GetISnapshotById(snapshotId string) (ICloudSnapshot, error)

	CreateSnapshotPolicy(*SnapshotPolicyInput) (string, error)
	GetISnapshotPolicies() ([]ICloudSnapshotPolicy, error)
	GetISnapshotPolicyById(id string) (ICloudSnapshotPolicy, error)

	GetIHosts() ([]ICloudHost, error)
	GetIHostById(id string) (ICloudHost, error)

	GetIStorages() ([]ICloudStorage, error)
	GetIStorageById(id string) (ICloudStorage, error)

	GetIStoragecaches() ([]ICloudStoragecache, error)
	GetIStoragecacheById(id string) (ICloudStoragecache, error)

	GetISkus() ([]ICloudSku, error)
	CreateISku(opts *SServerSkuCreateOption) (ICloudSku, error)

	GetICloudNatSkus() ([]ICloudNatSku, error)

	GetINetworkInterfaces() ([]ICloudNetworkInterface, error)

	GetIBuckets() ([]ICloudBucket, error)
	CreateIBucket(name string, storageClassStr string, acl string) error
	DeleteIBucket(name string) error
	IBucketExist(name string) (bool, error)
	GetIBucketById(name string) (ICloudBucket, error)
	GetIBucketByName(name string) (ICloudBucket, error)

	GetCloudEnv() string
	GetProvider() string

	GetICloudEvents(start time.Time, end time.Time, withReadEvent bool) ([]ICloudEvent, error) //获取公有云操作日志接口

	GetCapabilities() []string

	GetICloudQuotas() ([]ICloudQuota, error)

	GetICloudFileSystems() ([]ICloudFileSystem, error)
	GetICloudFileSystemById(id string) (ICloudFileSystem, error)

	CreateICloudFileSystem(opts *FileSystemCraeteOptions) (ICloudFileSystem, error)

	GetICloudAccessGroups() ([]ICloudAccessGroup, error)
	CreateICloudAccessGroup(opts *SAccessGroup) (ICloudAccessGroup, error)
	GetICloudAccessGroupById(id string) (ICloudAccessGroup, error)

	GetICloudWafIPSets() ([]ICloudWafIPSet, error)
	GetICloudWafRegexSets() ([]ICloudWafRegexSet, error)
	GetICloudWafInstances() ([]ICloudWafInstance, error)
	GetICloudWafInstanceById(id string) (ICloudWafInstance, error)
	CreateICloudWafInstance(opts *WafCreateOptions) (ICloudWafInstance, error)
	GetICloudWafRuleGroups() ([]ICloudWafRuleGroup, error)

	GetICloudMongoDBs() ([]ICloudMongoDB, error)
	GetICloudMongoDBById(id string) (ICloudMongoDB, error)

	GetIElasticSearchs() ([]ICloudElasticSearch, error)
	GetIElasticSearchById(id string) (ICloudElasticSearch, error)

	GetICloudKafkas() ([]ICloudKafka, error)
	GetICloudKafkaById(id string) (ICloudKafka, error)

	GetICloudApps() ([]ICloudApp, error)
	GetICloudAppById(id string) (ICloudApp, error)

	GetICloudTablestores() ([]ICloudTablestore, error)

	GetIModelartsPools() ([]ICloudModelartsPool, error)
	GetIModelartsPoolById(id string) (ICloudModelartsPool, error)
	CreateIModelartsPool(pool *ModelartsPoolCreateOption, callback func(externalId string)) (ICloudModelartsPool, error)
	GetIModelartsPoolSku() ([]ICloudModelartsPoolSku, error)

	GetIMiscResources() ([]ICloudMiscResource, error)
}

type ICloudZone interface {
	ICloudResource
	ICloudI18nResource

	GetIRegion() ICloudRegion

	GetIHosts() ([]ICloudHost, error)
	GetIHostById(id string) (ICloudHost, error)

	GetIStorages() ([]ICloudStorage, error)
	GetIStorageById(id string) (ICloudStorage, error)
}

type ICloudImage interface {
	IVirtualResource

	IOSInfo

	Delete(ctx context.Context) error
	GetIStoragecache() ICloudStoragecache

	GetSizeByte() int64
	GetImageType() TImageType
	GetImageStatus() string

	GetMinOsDiskSizeGb() int
	GetMinRamSizeMb() int
	GetImageFormat() string

	GetPublicScope() rbacscope.TRbacScope
	GetSubImages() []SSubImage

	Export(opts *SImageExportOptions) ([]SImageExportInfo, error)
}

type ICloudStoragecache interface {
	ICloudResource

	// 私有云需要实现
	GetICloudImages() ([]ICloudImage, error)
	// 公有云需要实现
	GetICustomizedCloudImages() ([]ICloudImage, error)
	GetIImageById(extId string) (ICloudImage, error)

	GetPath() string

	UploadImage(ctx context.Context, image *SImageCreateOption, callback func(float32)) (string, error)
}

type ICloudStorage interface {
	ICloudResource

	GetIStoragecache() ICloudStoragecache

	GetIZone() ICloudZone
	GetIDisks() ([]ICloudDisk, error)

	GetStorageType() string
	GetMediumType() string
	GetCapacityMB() int64 // MB
	GetCapacityUsedMB() int64
	GetStorageConf() jsonutils.JSONObject
	GetEnabled() bool

	CreateIDisk(conf *DiskCreateConfig) (ICloudDisk, error)

	GetIDiskById(idStr string) (ICloudDisk, error)

	GetMountPoint() string

	IsSysDiskStore() bool

	DisableSync() bool
}

type ICloudHost interface {
	ICloudResource

	GetIVMs() ([]ICloudVM, error)
	GetIVMById(id string) (ICloudVM, error)

	GetIStorages() ([]ICloudStorage, error)
	GetIStorageById(id string) (ICloudStorage, error)

	GetEnabled() bool      // is enabled
	GetHostStatus() string // service status
	GetAccessIp() string   //
	GetAccessMac() string  //
	GetSysInfo() jsonutils.JSONObject
	GetSN() string
	GetCpuCount() int
	GetNodeCount() int8
	GetCpuDesc() string
	GetCpuMhz() int
	GetCpuCmtbound() float32
	GetCpuArchitecture() string

	GetMemSizeMB() int
	GetMemCmtbound() float32
	GetReservedMemoryMb() int
	GetStorageSizeMB() int64
	GetStorageType() string
	GetHostType() string
	GetStorageDriver() string
	GetStorageInfo() jsonutils.JSONObject

	GetIsMaintenance() bool
	GetVersion() string

	CreateVM(desc *SManagedVMCreateConfig) (ICloudVM, error)
	GetIHostNics() ([]ICloudHostNetInterface, error)

	GetSchedtags() ([]string, error)

	GetOvnVersion() string // just for cloudpods host

	GetIsolateDevices() ([]IsolateDevice, error)
}

type IsolateDevice interface {
	GetName() string
	GetGlobalId() string
	GetModel() string
	GetAddr() string
	GetDevType() string
	GetNumaNode() int8
	GetVendorDeviceId() string
	GetSharedProjectIds() ([]string, error)
}

type ICloudVM interface {
	IBillingResource
	IVirtualResource

	IOSInfo

	ConvertPublicIpToEip() error

	GetHostname() string
	GetIHost() ICloudHost
	GetIHostId() string

	GetIDisks() ([]ICloudDisk, error)
	GetINics() ([]ICloudNic, error)

	GetIEIP() (ICloudEIP, error)

	GetInternetMaxBandwidthOut() int
	GetThroughput() int
	// GetStatus() string
	// GetRemoteStatus() string

	GetSerialOutput(port int) (string, error) // 目前仅谷歌云windows机器会使用到此接口

	GetCpuSockets() int
	GetVcpuCount() int
	GetVmemSizeMB() int //MB
	GetBootOrder() string
	GetVga() string
	GetVdi() string

	// GetOSArch() string
	// GetOsType() TOsType
	// GetOSName() string
	// GetBios() string

	GetMachine() string
	GetInstanceType() string

	GetSecurityGroupIds() ([]string, error)
	SetSecurityGroups(secgroupIds []string) error

	GetHypervisor() string

	StartVM(ctx context.Context) error
	StopVM(ctx context.Context, opts *ServerStopOptions) error
	// 需要删除挂载的磁盘
	DeleteVM(ctx context.Context) error

	UpdateVM(ctx context.Context, input SInstanceUpdateOptions) error

	UpdateUserData(userData string) error

	RebuildRoot(ctx context.Context, config *SManagedVMRebuildRootConfig) (string, error)

	DeployVM(ctx context.Context, opts *SInstanceDeployOptions) error

	ChangeConfig(ctx context.Context, config *SManagedVMChangeConfig) error

	GetVNCInfo(input *ServerVncInput) (*ServerVncOutput, error)
	// 若有跟随主机删除的选项，需要设置为True
	AttachDisk(ctx context.Context, diskId string) error
	DetachDisk(ctx context.Context, diskId string) error

	CreateDisk(ctx context.Context, opts *GuestDiskCreateOptions) (string, error)

	MigrateVM(hostid string) error
	LiveMigrateVM(hostid string) error

	GetError() error

	CreateInstanceSnapshot(ctx context.Context, name string, desc string) (ICloudInstanceSnapshot, error)
	GetInstanceSnapshot(idStr string) (ICloudInstanceSnapshot, error)
	GetInstanceSnapshots() ([]ICloudInstanceSnapshot, error)
	ResetToInstanceSnapshot(ctx context.Context, idStr string) error

	SaveImage(opts *SaveImageOptions) (ICloudImage, error)

	AllocatePublicIpAddress() (string, error)
	GetPowerStates() string
	GetIsolateDeviceIds() ([]string, error)
}

type ICloudNic interface {
	GetId() string
	GetIP() string
	GetIP6() string
	GetMAC() string
	InClassicNetwork() bool
	GetDriver() string
	GetINetworkId() string

	// GetSubAddress returns non-primary/secondary/alias ipv4 addresses of
	// the network interface
	//
	// Implement it when any AssignXx ops methods are implemented
	GetSubAddress() ([]string, error)
	AssignNAddress(count int) ([]string, error)
	AssignAddress(ipAddrs []string) error
	// UnassignAddress should not return error if the network interface is
	// now not present, or the addresses is not assigned to the network
	// interface in the first place
	UnassignAddress(ipAddrs []string) error
}

const ErrAddressCountExceed = errors.Error("ErrAddressCountExceed")

type DummyICloudNic struct{}

var _ ICloudNic = DummyICloudNic{}

func (d DummyICloudNic) GetId() string          { panic(errors.ErrNotImplemented) }
func (d DummyICloudNic) GetIP() string          { panic(errors.ErrNotImplemented) }
func (d DummyICloudNic) GetIP6() string         { return "" }
func (d DummyICloudNic) GetMAC() string         { panic(errors.ErrNotImplemented) }
func (d DummyICloudNic) InClassicNetwork() bool { panic(errors.ErrNotImplemented) }
func (d DummyICloudNic) GetDriver() string      { panic(errors.ErrNotImplemented) }
func (d DummyICloudNic) GetINetworkId() string  { panic(errors.ErrNotImplemented) }
func (d DummyICloudNic) GetSubAddress() ([]string, error) {
	return nil, nil
}
func (d DummyICloudNic) AssignNAddress(count int) ([]string, error) {
	return nil, errors.ErrNotImplemented
}
func (d DummyICloudNic) AssignAddress(ipAddrs []string) error   { return errors.ErrNotImplemented }
func (d DummyICloudNic) UnassignAddress(ipAddrs []string) error { return errors.ErrNotImplemented }

type ICloudEIP interface {
	IBillingResource
	IVirtualResource

	GetIpAddr() string
	GetMode() string
	GetINetworkId() string
	GetAssociationType() string
	GetAssociationExternalId() string

	GetBandwidth() int

	GetInternetChargeType() string

	Delete() error

	Associate(conf *AssociateConfig) error
	Dissociate() error

	ChangeBandwidth(bw int) error
}

type ICloudSecurityGroup interface {
	IVirtualResource

	GetDescription() string
	// 返回的优先级字段(priority)要求数字越大优先级越高, 若有默认不可修改的allow规则依然需要返回
	GetRules() ([]ISecurityGroupRule, error)
	GetVpcId() string

	CreateRule(opts *SecurityGroupRuleCreateOptions) (ISecurityGroupRule, error)

	GetReferences() ([]SecurityGroupReference, error)
	Delete() error
}

type ISecurityGroupRule interface {
	GetGlobalId() string
	GetDirection() secrules.TSecurityRuleDirection
	GetPriority() int
	GetAction() secrules.TSecurityRuleAction
	GetProtocol() string
	GetPorts() string
	GetDescription() string
	GetCIDRs() []string

	Update(opts *SecurityGroupRuleUpdateOptions) error
	Delete() error
}

type ICloudRouteTable interface {
	ICloudResource

	GetAssociations() []RouteTableAssociation
	GetDescription() string
	GetRegionId() string
	GetVpcId() string
	GetType() RouteTableType
	GetIRoutes() ([]ICloudRoute, error)

	CreateRoute(route RouteSet) error
	UpdateRoute(route RouteSet) error
	RemoveRoute(route RouteSet) error
}

type ICloudRoute interface {
	ICloudResource
	GetType() string
	GetCidr() string
	GetNextHopType() string
	GetNextHop() string
}

type ICloudDisk interface {
	IBillingResource
	IVirtualResource

	GetIStorage() (ICloudStorage, error)
	GetIStorageId() string

	// GetStatus() string
	GetDiskFormat() string
	GetDiskSizeMB() int // MB
	GetIsAutoDelete() bool
	GetTemplateId() string
	GetDiskType() string
	GetFsFormat() string
	GetIsNonPersistent() bool
	GetIops() int

	GetDriver() string
	GetCacheMode() string
	GetMountpoint() string

	GetAccessPath() string

	Delete(ctx context.Context) error

	CreateISnapshot(ctx context.Context, name string, desc string) (ICloudSnapshot, error)
	GetISnapshots() ([]ICloudSnapshot, error)

	Resize(ctx context.Context, newSizeMB int64) error
	Reset(ctx context.Context, snapshotId string) (string, error)

	Rebuild(ctx context.Context) error

	GetPreallocation() string
}

type ICloudSnapshot interface {
	IVirtualResource

	GetSizeMb() int32
	GetDiskId() string
	GetDiskType() string
	Delete() error
}

type ICloudInstanceSnapshot interface {
	IVirtualResource

	GetDescription() string
	Delete() error
}

type ICloudSnapshotPolicy interface {
	IVirtualResource

	GetRetentionDays() int
	GetRepeatWeekdays() ([]int, error)
	GetTimePoints() ([]int, error)
	Delete() error
	ApplyDisks(ids []string) error
	CancelDisks(ids []string) error
	GetApplyDiskIds() ([]string, error)
}

type ICloudGlobalVpc interface {
	ICloudResource

	GetISecurityGroups() ([]ICloudSecurityGroup, error)
	CreateISecurityGroup(opts *SecurityGroupCreateInput) (ICloudSecurityGroup, error)

	Delete() error
}

type ICloudIPv6Gateway interface {
	IVirtualResource

	GetInstanceType() string
}

type ICloudVpc interface {
	ICloudResource

	GetGlobalVpcId() string
	IsSupportSetExternalAccess() bool // 是否支持Attach互联网网关.
	GetExternalAccessMode() string
	AttachInternetGateway(igwId string) error

	GetRegion() ICloudRegion
	GetIsDefault() bool
	GetCidrBlock() string
	GetCidrBlock6() string
	GetIWires() ([]ICloudWire, error)
	CreateIWire(opts *SWireCreateOptions) (ICloudWire, error)
	GetISecurityGroups() ([]ICloudSecurityGroup, error)
	GetIRouteTables() ([]ICloudRouteTable, error)
	GetIRouteTableById(routeTableId string) (ICloudRouteTable, error)

	Delete() error

	GetIWireById(wireId string) (ICloudWire, error)
	GetINatGateways() ([]ICloudNatGateway, error)
	CreateINatGateway(opts *NatGatewayCreateOptions) (ICloudNatGateway, error)

	GetICloudVpcPeeringConnections() ([]ICloudVpcPeeringConnection, error)
	GetICloudAccepterVpcPeeringConnections() ([]ICloudVpcPeeringConnection, error)
	GetICloudVpcPeeringConnectionById(id string) (ICloudVpcPeeringConnection, error)
	CreateICloudVpcPeeringConnection(opts *VpcPeeringConnectionCreateOptions) (ICloudVpcPeeringConnection, error)
	AcceptICloudVpcPeeringConnection(id string) error

	GetAuthorityOwnerId() string

	ProposeJoinICloudInterVpcNetwork(opts *SVpcJointInterVpcNetworkOption) error

	GetICloudIPv6Gateways() ([]ICloudIPv6Gateway, error)
}

type ICloudInternetGateway interface {
	ICloudResource
}

type ICloudWire interface {
	ICloudResource
	GetIVpc() ICloudVpc
	GetIZone() ICloudZone
	GetINetworks() ([]ICloudNetwork, error)
	GetBandwidth() int

	GetINetworkById(netid string) (ICloudNetwork, error)

	CreateINetwork(opts *SNetworkCreateOptions) (ICloudNetwork, error)
}

type ICloudNetwork interface {
	IVirtualResource

	GetIWire() ICloudWire

	GetIpStart() string
	GetIpEnd() string
	GetIpMask() int8
	GetGateway() string

	// IPv6
	GetIp6Start() string
	GetIp6End() string
	GetIp6Mask() uint8
	GetGateway6() string

	GetServerType() string
	//GetIsPublic() bool
	// 仅私有云有用，公有云无效
	// 1. scope = none 非共享, network仅会属于一个项目,并且私有
	// 2. scope = system 系统共享 云账号共享会跟随云账号共享，云账号非共享,会共享到network所在域
	GetPublicScope() rbacscope.TRbacScope

	Delete() error

	GetAllocTimeoutSeconds() int
}

type ICloudHostNetInterface interface {
	GetDevice() string
	GetDriver() string
	GetMac() string
	GetVlanId() int
	GetIndex() int8
	IsLinkUp() tristate.TriState
	GetIpAddr() string
	GetMtu() int32
	GetNicType() string
	GetBridge() string
	GetIWire() ICloudWire
}

type ICloudSku interface {
	ICloudResource

	GetInstanceTypeFamily() string
	GetInstanceTypeCategory() string

	GetPrepaidStatus() string
	GetPostpaidStatus() string

	GetCpuArch() string
	GetCpuCoreCount() int
	GetMemorySizeMB() int

	GetOsName() string

	GetSysDiskResizable() bool
	GetSysDiskType() string
	GetSysDiskMinSizeGB() int
	GetSysDiskMaxSizeGB() int

	GetAttachedDiskType() string
	GetAttachedDiskSizeGB() int
	GetAttachedDiskCount() int

	GetDataDiskTypes() string
	GetDataDiskMaxCount() int

	GetNicType() string
	GetNicMaxCount() int

	GetGpuAttachable() bool
	GetGpuSpec() string
	GetGpuCount() string
	GetGpuMaxCount() int

	Delete() error
}

type ICloudProject interface {
	ICloudResource
}

type ICloudNatGateway interface {
	ICloudResource
	IBillingResource

	// 获取 NAT 规格
	GetNatSpec() string
	GetIEips() ([]ICloudEIP, error)
	GetINatDTable() ([]ICloudNatDEntry, error)
	GetINatSTable() ([]ICloudNatSEntry, error)

	// ID is the ID of snat entry/rule or dnat entry/rule.
	GetINatDEntryById(id string) (ICloudNatDEntry, error)
	GetINatSEntryById(id string) (ICloudNatSEntry, error)

	// Read the description of these two structures before using.
	CreateINatDEntry(rule SNatDRule) (ICloudNatDEntry, error)
	CreateINatSEntry(rule SNatSRule) (ICloudNatSEntry, error)

	GetINetworkId() string
	// internet(公网) or intranet(VPC)
	GetNetworkType() string
	GetBandwidthMb() int
	GetIpAddr() string

	Delete() error
}

// ICloudNatDEntry describe a DNat rule which transfer externalIp:externalPort to
// internalIp:internalPort with IpProtocol(tcp/udp)
type ICloudNatDEntry interface {
	ICloudResource

	GetIpProtocol() string
	GetExternalIp() string
	GetExternalPort() int

	GetInternalIp() string
	GetInternalPort() int

	Delete() error
}

// ICloudNatSEntry describe a SNat rule which transfer internalIp(GetIP()) to externalIp which from sourceCIDR
type ICloudNatSEntry interface {
	ICloudResource

	GetIP() string
	GetSourceCIDR() string
	GetNetworkId() string

	Delete() error
}

type ICloudNetworkInterface interface {
	ICloudResource

	GetMacAddress() string
	GetAssociateType() string
	GetAssociateId() string

	GetICloudInterfaceAddresses() ([]ICloudInterfaceAddress, error)
}

type ICloudInterfaceAddress interface {
	GetGlobalId() string //返回IP即可

	GetINetworkId() string
	GetIP() string
	IsPrimary() bool
}

type ICloudEvent interface {
	GetName() string
	GetService() string
	GetAction() string
	GetResourceType() string
	GetRequestId() string
	GetRequest() jsonutils.JSONObject
	GetAccount() string
	IsSuccess() bool

	GetCreatedAt() time.Time
}

type ICloudQuota interface {
	GetGlobalId() string
	GetDesc() string
	GetQuotaType() string
	GetMaxQuotaCount() int
	GetCurrentQuotaUsedCount() int
}

type SClouduserEnableOptions struct {
	Password              string
	EnableMfa             bool
	PasswordResetRequired bool
}

// 公有云子账号
type IClouduser interface {
	GetGlobalId() string
	GetName() string

	GetEmailAddr() string
	GetInviteUrl() string

	GetICloudgroups() ([]ICloudgroup, error)

	GetICloudpolicies() ([]ICloudpolicy, error)

	AttachPolicy(policyName string, policyType api.TPolicyType) error
	DetachPolicy(policyName string, policyType api.TPolicyType) error

	SetEnable(opts *SClouduserEnableOptions) error
	SetDisable() error

	Delete() error

	ResetPassword(password string) error
	IsConsoleLogin() bool

	CreateAccessKey(name string) (*SAccessKey, error)
	DeleteAccessKey(accessKey string) error
	GetAccessKeys() ([]SAccessKey, error)
}

// 公有云子账号权限
type ICloudpolicy interface {
	GetGlobalId() string
	GetName() string
	GetDescription() string
	GetPolicyType() api.TPolicyType

	GetDocument() (*jsonutils.JSONDict, error)
	UpdateDocument(*jsonutils.JSONDict) error

	Delete() error
}

// 公有云用户组
type ICloudgroup interface {
	GetGlobalId() string
	GetName() string
	GetDescription() string
	GetICloudpolicies() ([]ICloudpolicy, error)
	GetICloudusers() ([]IClouduser, error)

	AddUser(name string) error
	RemoveUser(name string) error

	AttachPolicy(policyName string, policyType api.TPolicyType) error
	DetachPolicy(policyName string, policyType api.TPolicyType) error

	Delete() error
}

type ICloudVpcPeeringConnection interface {
	ICloudResource

	GetPeerVpcId() string
	GetPeerAccountId() string
	GetEnabled() bool
	Delete() error
}

type ICloudSAMLProvider interface {
	ICloudResource

	GetMetadataDocument() (*samlutils.EntityDescriptor, error)
	UpdateMetadata(samlutils.EntityDescriptor) error

	GetAuthUrl(apiServer string) string
	Delete() error
}

type ICloudrole interface {
	GetGlobalId() string
	GetName() string

	GetDocument() *jsonutils.JSONDict
	GetSAMLProvider() string

	GetICloudpolicies() ([]ICloudpolicy, error)
	AttachPolicy(policyName string, policyType api.TPolicyType) error
	DetachPolicy(policyName string, policyType api.TPolicyType) error

	Delete() error
}

type ICloudInterVpcNetwork interface {
	ICloudResource
	GetAuthorityOwnerId() string
	GetICloudVpcIds() ([]string, error)
	AttachVpc(opts *SInterVpcNetworkAttachVpcOption) error
	DetachVpc(opts *SInterVpcNetworkDetachVpcOption) error
	Delete() error
	GetIRoutes() ([]ICloudInterVpcNetworkRoute, error)
	EnableRouteEntry(routeId string) error
	DisableRouteEntry(routeId string) error
}

type ICloudInterVpcNetworkRoute interface {
	ICloudResource
	GetInstanceId() string
	GetInstanceType() string
	GetInstanceRegionId() string

	GetEnabled() bool
	GetCidr() string
}

type ICloudFileSystem interface {
	IVirtualResource
	IBillingResource

	GetFileSystemType() string
	GetStorageType() string
	GetProtocol() string
	GetCapacityGb() int64
	GetUsedCapacityGb() int64
	GetMountTargetCountLimit() int

	GetZoneId() string

	GetMountTargets() ([]ICloudMountTarget, error)
	CreateMountTarget(opts *SMountTargetCreateOptions) (ICloudMountTarget, error)

	SetQuota(input *SFileSystemSetQuotaInput) error

	Delete() error
}

type ICloudMountTarget interface {
	GetGlobalId() string
	GetName() string
	GetAccessGroupId() string
	GetDomainName() string
	GetNetworkType() string
	GetVpcId() string
	GetNetworkId() string
	GetStatus() string

	Delete() error
}

type ICloudAccessGroup interface {
	GetGlobalId() string
	GetName() string
	GetDesc() string
	GetSupporedUserAccessTypes() []TUserAccessType
	GetNetworkType() string
	GetFileSystemType() string
	GetMountTargetCount() int

	GetRules() ([]IAccessGroupRule, error)
	CreateRule(opts *AccessGroupRule) (IAccessGroupRule, error)

	Delete() error
}

type IAccessGroupRule interface {
	GetGlobalId() string
	GetPriority() int
	GetRWAccessType() TRWAccessType
	GetUserAccessType() TUserAccessType
	GetSource() string

	Delete() error
}

type ICloudWafIPSet interface {
	GetName() string
	GetDesc() string
	GetType() TWafType
	GetGlobalId() string
	GetAddresses() WafAddresses

	Delete() error
}

type ICloudWafRegexSet interface {
	GetName() string
	GetDesc() string
	GetType() TWafType
	GetGlobalId() string
	GetRegexPatterns() WafRegexPatterns

	Delete() error
}

type ICloudWafInstance interface {
	ICloudEnabledResource

	GetWafType() TWafType
	GetDefaultAction() *DefaultAction
	GetRules() ([]ICloudWafRule, error)
	AddRule(opts *SWafRule) (ICloudWafRule, error)

	// 绑定的资源列表
	GetCloudResources() ([]SCloudResource, error)

	// 前面是否有代理服务
	GetIsAccessProduct() bool
	GetAccessHeaders() []string
	GetHttpPorts() []int
	GetHttpsPorts() []int
	GetCname() string
	// 源站地址
	GetSourceIps() []string
	// 回源地址
	GetCcList() []string
	GetCertId() string
	GetCertName() string
	GetUpstreamScheme() string
	GetUpstreamPort() int

	Delete() error
}

type ICloudWafRuleGroup interface {
	GetName() string
	GetDesc() string
	GetGlobalId() string
	GetWafType() TWafType
	GetRules() ([]ICloudWafRule, error)
}

type ICloudWafRule interface {
	GetName() string
	GetDesc() string
	GetGlobalId() string
	GetPriority() int
	GetType() string
	GetAction() *DefaultAction
	GetStatementCondition() TWafStatementCondition
	GetStatements() ([]SWafStatement, error)

	Update(opts *SWafRule) error
	Delete() error
}

type ICloudMongoDB interface {
	IVirtualResource
	IBillingResource

	GetVpcId() string
	GetNetworkId() string
	GetIpAddr() string
	GetVcpuCount() int
	GetVmemSizeMb() int
	GetDiskSizeMb() int
	GetZoneId() string
	GetReplicationNum() int
	GetCategory() string
	GetEngine() string
	GetEngineVersion() string
	GetInstanceType() string
	GetMaintainTime() string
	GetPort() int
	GetIops() int

	GetMaxConnections() int

	GetNetworkAddress() string

	GetIBackups() ([]SMongoDBBackup, error)

	Delete() error
}

type ICloudElasticSearch interface {
	IVirtualResource
	IBillingResource

	GetVersion() string
	GetStorageType() string
	GetDiskSizeGb() int
	GetCategory() string

	GetInstanceType() string
	GetVcpuCount() int
	GetVmemSizeGb() int

	GetVpcId() string
	GetNetworkId() string
	GetZoneId() string
	IsMultiAz() bool

	Delete() error
}

type ICloudKafka interface {
	IVirtualResource
	IBillingResource

	GetNetworkId() string
	GetVpcId() string
	GetZoneId() string
	GetInstanceType() string

	GetVersion() string
	GetDiskSizeGb() int
	GetStorageType() string
	GetBandwidthMb() int
	GetEndpoint() string
	GetMsgRetentionMinute() int

	IsMultiAz() bool

	Delete() error
}

type AppBackupConfig struct {
	Enabled               bool
	FrequencyInterval     int
	FrequencyUnit         string
	RetentionPeriodInDays int
}

type ICloudApp interface {
	IVirtualResource
	GetEnvironments() ([]ICloudAppEnvironment, error)
	GetTechStack() string
	GetOsType() TOsType
	GetIpAddress() string
	GetHostname() string
	GetServerFarm() string
	GetBackups() ([]IAppBackup, error)
	GetPublicNetworkAccess() string
	GetNetworkId() string
	GetHybirdConnections() ([]IAppHybirdConnection, error)
	GetCertificates() ([]IAppCertificate, error)
	GetBackupConfig() AppBackupConfig
	GetDomains() ([]IAppDomain, error)
}

type IAppDomain interface {
	GetGlobalId() string
	GetName() string
	GetStatus() string
	GetSslState() string
}

type IAppCertificate interface {
	GetGlobalId() string
	GetName() string
	GetSubjectName() string
	GetIssuer() string
	GetIssueDate() time.Time
	GetThumbprint() string
	GetExpireTime() time.Time
}

type IAppHybirdConnection interface {
	GetGlobalId() string
	GetName() string
	GetHostname() string
	GetNamespace() string
	GetPort() int
}

type IAppBackup interface {
	GetGlobalId() string
	GetName() string
	GetType() string
}

type ICloudAppEnvironment interface {
	IVirtualResource
}

type ICloudNatSku interface {
	GetName() string
	GetDesc() string
	GetGlobalId() string
	GetPrepaidStatus() string
	GetPostpaidStatus() string
}

type ICloudCDNDomain interface {
	IVirtualResource
	GetEnabled() bool

	GetArea() string
	GetServiceType() string
	GetCname() string
	GetOrigins() *SCdnOrigins

	// 是否忽略参数
	GetCacheKeys() (*SCDNCacheKeys, error)
	// 是否分片回源
	GetRangeOriginPull() (*SCDNRangeOriginPull, error)
	// 缓存配置
	GetCache() (*SCDNCache, error)
	// https配置
	GetHTTPS() (*SCDNHttps, error)
	// 强制跳转
	GetForceRedirect() (*SCDNForceRedirect, error)
	// 防盗链配置
	GetReferer() (*SCDNReferer, error)
	// 浏览器缓存配置
	GetMaxAge() (*SCDNMaxAge, error)

	Delete() error
}

type ICloudTablestore interface {
	IVirtualResource
}

type ICloudMiscResource interface {
	IVirtualResource

	GetResourceType() string

	GetConfig() jsonutils.JSONObject
}

type ICloudSSLCertificate interface {
	IVirtualResource

	GetSans() string
	GetStartDate() time.Time
	GetProvince() string
	GetCommon() string
	GetCountry() string
	GetIssuer() string
	GetExpired() bool
	GetEndDate() time.Time
	GetFingerprint() string
	GetCity() string
	GetOrgName() string
	GetIsUpload() bool
	GetCert() string
	GetKey() string
}

package cloudprovider

import (
	"context"
	"fmt"
	"net/http"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/httputils"
	"yunion.io/x/pkg/utils"
)

type SCloudaccountCredential struct {
	// Username (openstack, esxi)
	Username string `json:"username"`

	// Password (openstack, esxi)
	Password string `json:"password"`

	// Key id (Aliyun, Aws)
	AccessKeyId string `json:"access_key_id"`

	// Key (Aliyun, Aws)
	AccessKeySecret string `json:"access_key_secret"`

	// Environment (Azure, Aws, aliyun)
	Environment string `json:"environment"`

	// Directory id (Azure)
	DirectoryId string `json:"directory_id"`

	// Client id (Azure)
	ClientId string `json:"client_id"`

	// Client secret key (Azure)
	ClientSecret string `json:"client_secret"`

	// Host IP (esxi)
	Host string `json:"host"`
	// Host port (esxi)
	Port int `json:"port"`

	// The highest organization id allowed by FeiTian, ​​default is 1
	OrganizationId int `json:"organization_id"`

	// Google service account email (gcp)
	GCPClientEmail string `json:"gcp_client_email"`
	// Google service account project id (gcp)
	GCPProjectId string `json:"gcp_project_id"`
	// Google service account key id (gcp)
	GCPPrivateKeyId string `json:"gcp_private_key_id"`
	// Google service account key (gcp)
	GCPPrivateKey string `json:"gcp_private_key"`

	RegionId string
}

type SCloudaccount struct {
	Account string `json:"account"`

	// swagger:ignore
	Secret string

	// Authentication address
	AccessUrl string `json:"access_url"`
}

type ProviderConfig struct {
	// Id, Name are properties of Cloudprovider object
	Id   string
	Name string

	// Vendor are names like Aliyun, etc.
	Vendor  string
	URL     string
	Account string
	Secret  string

	ReadOnly bool

	AccountId string

	Options *jsonutils.JSONDict

	RegionId  string
	ProxyFunc httputils.TransportProxyFunc
	Debug     bool

	// Only used to detect whether cloudpods manages its own environment (system project id)
	AdminProjectId string

	AliyunResourceGroupIds []string

	UpdatePermission func(service, permission string)
}

func (cp *ProviderConfig) AdaptiveTimeoutHttpClient() *http.Client {
	client := httputils.GetAdaptiveTimeoutClient()
	httputils.SetClientProxyFunc(client, cp.ProxyFunc)
	return client
}

type SProviderInfo struct {
	Name    string
	Url     string
	Account string
	Secret  string
	Region  string
	Options *jsonutils.JSONDict
}

type ICloudProviderFactory interface {
	GetProvider(cfg ProviderConfig) (ICloudProvider, error)

	GetClientRC(SProviderInfo) (map[string]string, error)

	GetId() string
	GetName() string

	ValidateChangeBandwidth(instanceId string, bandwidth int64) error
	ValidateCreateCloudaccountData(ctx context.Context, input SCloudaccountCredential) (SCloudaccount, error)
	IsReadOnly() bool
	ValidateUpdateCloudaccountCredential(ctx context.Context, input SCloudaccountCredential, cloudaccount string) (SCloudaccount, error)
	GetSupportedBrands() []string

	IsPublicCloud() bool
	IsOnPremise() bool
	IsMultiTenant() bool
	IsSupportPrepaidResources() bool
	NeedSyncSkuFromCloud() bool

	IsCloudeventRegional() bool
	GetMaxCloudEventSyncDays() int
	GetMaxCloudEventKeepDays() int

	IsNeedForceAutoCreateProject() bool

	IsSupportCrossCloudEnvVpcPeering() bool
	IsSupportCrossRegionVpcPeering() bool
	IsSupportVpcPeeringVpcCidrOverlap() bool
	ValidateCrossRegionVpcPeeringBandWidth(bandwidth int) error

	IsSupportModifyRouteTable() bool

	GetSupportedDnsZoneTypes() []TDnsZoneType
	GetSupportedDnsTypes() map[TDnsZoneType][]TDnsType
	GetSupportedDnsPolicyTypes() map[TDnsZoneType][]TDnsPolicyType
	GetSupportedDnsPolicyValues() map[TDnsPolicyType][]TDnsPolicyValue
	GetTTLRange(zoneType TDnsZoneType, productType TDnsProductType) TTlRange

	IsSupportSAMLAuth() bool

	GetAccountIdEqualizer() func(origin, now string) bool
}

type SBalanceInfo struct {
	Currency string
	Amount   float64
	Status   string
}

type ICloudProvider interface {
	GetFactory() ICloudProviderFactory

	GetSysInfo() (jsonutils.JSONObject, error)
	GetVersion() string
	GetIamLoginUrl() string

	GetIRegions() ([]ICloudRegion, error)
	GetIProjects() ([]ICloudProject, error)
	CreateIProject(name string) (ICloudProject, error)
	GetIRegionById(id string) (ICloudRegion, error)

	GetOnPremiseIRegion() (ICloudRegion, error)

	GetBalance() (*SBalanceInfo, error)

	GetSubAccounts() ([]SSubAccount, error)
	GetAccountId() string

	// The region external id is prefixed with provider. Therefore, this judgment condition can be used to filter out the regions list of the same provider.
	// However, Huawei Cloud is a bit special. A provider only corresponds to one region. Therefore, it is necessary to further specify the region name to find the region corresponding to the provider.
	GetCloudRegionExternalIdPrefix() string

	GetStorageClasses(regionId string) []string
	GetBucketCannedAcls(regionId string) []string
	GetObjectCannedAcls(regionId string) []string

	GetCapabilities() []string

	IsClouduserSupportPassword() bool
	GetICloudusers() ([]IClouduser, error)
	GetICloudpolicies() ([]ICloudpolicy, error)
	GetICloudgroups() ([]ICloudgroup, error)
	GetICloudgroupByName(name string) (ICloudgroup, error)
	CreateICloudgroup(name, desc string) (ICloudgroup, error)
	GetIClouduserByName(name string) (IClouduser, error)
	CreateIClouduser(conf *SClouduserCreateConfig) (IClouduser, error)
	CreateICloudSAMLProvider(opts *SAMLProviderCreateOptions) (ICloudSAMLProvider, error)
	GetICloudroles() ([]ICloudrole, error)
	GetICloudroleById(id string) (ICloudrole, error)
	GetICloudroleByName(name string) (ICloudrole, error)
	CreateICloudrole(opts *SRoleCreateOptions) (ICloudrole, error)

	CreateICloudpolicy(opts *SCloudpolicyCreateOptions) (ICloudpolicy, error)

	GetEnrollmentAccounts() ([]SEnrollmentAccount, error)
	CreateSubscription(SubscriptionCreateInput) error

	GetSamlEntityId() string

	GetICloudDnsZones() ([]ICloudDnsZone, error)
	GetICloudDnsZoneById(id string) (ICloudDnsZone, error)
	CreateICloudDnsZone(opts *SDnsZoneCreateOptions) (ICloudDnsZone, error)

	GetICloudGlobalVpcs() ([]ICloudGlobalVpc, error)
	CreateICloudGlobalVpc(opts *GlobalVpcCreateOptions) (ICloudGlobalVpc, error)
	GetICloudGlobalVpcById(id string) (ICloudGlobalVpc, error)

	GetICloudInterVpcNetworks() ([]ICloudInterVpcNetwork, error)
	GetICloudInterVpcNetworkById(id string) (ICloudInterVpcNetwork, error)
	CreateICloudInterVpcNetwork(opts *SInterVpcNetworkCreateOptions) (ICloudInterVpcNetwork, error)

	GetICloudCDNDomains() ([]ICloudCDNDomain, error)
	GetICloudCDNDomainByName(name string) (ICloudCDNDomain, error)
	CreateICloudCDNDomain(opts *CdnCreateOptions) (ICloudCDNDomain, error)

	GetMetrics(opts *MetricListOptions) ([]MetricValues, error)

	GetISSLCertificates() ([]ICloudSSLCertificate, error)
}

func IsSupportCapability(prod ICloudProvider, capa string) bool {
	return utils.IsInStringArray(capa, prod.GetCapabilities()) || utils.IsInStringArray(capa+READ_ONLY_SUFFIX, prod.GetCapabilities())
}

func IsSupportCDN(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_CDN)
}

func IsSupportProject(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_PROJECT)
}

func IsSupportQuota(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_QUOTA)
}

func IsSupportDnsZone(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_DNSZONE)
}

func IsSupportInterVpcNetwork(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_INTERVPCNETWORK)
}

func IsSupportCompute(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_COMPUTE)
}

func IsSupportLoadbalancer(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_LOADBALANCER)
}

func IsSupportObjectstore(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_OBJECTSTORE)
}

func IsSupportRds(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_RDS)
}

func IsSupportNAS(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_NAS)
}

func IsSupportNAT(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_NAT)
}

func IsSupportElasticCache(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_CACHE)
}

func IsSupportWaf(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_WAF)
}

func IsSupportMongoDB(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_MONGO_DB)
}

func IsSupportElasticSearch(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_ES)
}

func IsSupportSSLCertificate(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_CERT)
}

func IsSupportKafka(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_KAFKA)
}

func IsSupportApp(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_APP)
}

func IsSupportContainer(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_CONTAINER)
}

func IsSupportTablestore(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_TABLESTORE)
}

func IsSupportModelartsPool(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_MODELARTES)
}

func IsSupportMiscResources(prod ICloudProvider) bool {
	return IsSupportCapability(prod, CLOUD_CAPABILITY_MISC)
}

var providerTable map[string]ICloudProviderFactory

func init() {
	providerTable = make(map[string]ICloudProviderFactory)
}

func RegisterFactory(factory ICloudProviderFactory) {
	providerTable[factory.GetId()] = factory
}

func GetProviderFactory(provider string) (ICloudProviderFactory, error) {
	factory, ok := providerTable[provider]
	if ok {
		return factory, nil
	}
	return nil, fmt.Errorf("No such provider %s", provider)
}

func GetRegistedProviderIds() []string {
	providers := []string{}
	for id := range providerTable {
		providers = append(providers, id)
	}
	return providers
}

func GetProvider(cfg ProviderConfig) (ICloudProvider, error) {
	driver, err := GetProviderFactory(cfg.Vendor)
	if err != nil {
		return nil, errors.Wrap(err, "GetProviderFactory")
	}
	return driver.GetProvider(cfg)
}

func GetClientRC(provider string, info SProviderInfo) (map[string]string, error) {
	driver, err := GetProviderFactory(provider)
	if err != nil {
		return nil, errors.Wrap(err, "GetProviderFactory")
	}
	return driver.GetClientRC(info)
}

func IsSupported(provider string) bool {
	_, ok := providerTable[provider]
	return ok
}

func IsValidCloudAccount(cfg ProviderConfig) (ICloudProvider, string, error) {
	factory, ok := providerTable[cfg.Vendor]
	if ok {
		provider, err := factory.GetProvider(cfg)
		if err != nil {
			return nil, "", err
		}
		return provider, provider.GetAccountId(), nil
	}
	return nil, "", ErrNoSuchProvder
}

type SBaseProvider struct {
	factory ICloudProviderFactory
}

func (provider *SBaseProvider) GetFactory() ICloudProviderFactory {
	return provider.factory
}

func (self *SBaseProvider) GetOnPremiseIRegion() (ICloudRegion, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetIamLoginUrl() string {
	return ""
}

func (self *SBaseProvider) IsClouduserSupportPassword() bool {
	return true
}

func (self *SBaseProvider) GetICloudusers() ([]IClouduser, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudgroups() ([]ICloudgroup, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudgroupByName(name string) (ICloudgroup, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) CreateICloudgroup(name, desc string) (ICloudgroup, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudpolicies() ([]ICloudpolicy, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetIClouduserByName(name string) (IClouduser, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) CreateIClouduser(conf *SClouduserCreateConfig) (IClouduser, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudroles() ([]ICloudrole, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudroles")
}

func (self *SBaseProvider) GetICloudroleById(id string) (ICloudrole, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudroleById")
}

func (self *SBaseProvider) GetICloudroleByName(name string) (ICloudrole, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudroleByName")
}

func (self *SBaseProvider) CreateICloudrole(opts *SRoleCreateOptions) (ICloudrole, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "CreateICloudrole")
}

func (self *SBaseProvider) CreateICloudSAMLProvider(opts *SAMLProviderCreateOptions) (ICloudSAMLProvider, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "CreateICloudSAMLProvider")
}

func (self *SBaseProvider) CreateICloudpolicy(opts *SCloudpolicyCreateOptions) (ICloudpolicy, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetEnrollmentAccounts() ([]SEnrollmentAccount, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) CreateSubscription(SubscriptionCreateInput) error {
	return ErrNotImplemented
}

func (self *SBaseProvider) GetICloudDnsZones() ([]ICloudDnsZone, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudDnsZoneById(id string) (ICloudDnsZone, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) CreateICloudDnsZone(opts *SDnsZoneCreateOptions) (ICloudDnsZone, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetCloudRegionExternalIdPrefix() string {
	return self.factory.GetId()
}

func (self *SBaseProvider) CreateIProject(name string) (ICloudProject, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetSamlEntityId() string {
	return ""
}

func (self *SBaseProvider) GetSamlSpInitiatedLoginUrl(idpName string) string {
	return ""
}

func (self *SBaseProvider) GetICloudInterVpcNetworks() ([]ICloudInterVpcNetwork, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudInterVpcNetworkById(id string) (ICloudInterVpcNetwork, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) CreateICloudInterVpcNetwork(opts *SInterVpcNetworkCreateOptions) (ICloudInterVpcNetwork, error) {
	return nil, ErrNotImplemented
}

func (self *SBaseProvider) GetICloudGlobalVpcs() ([]ICloudGlobalVpc, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudGlobalVpcs")
}

func (self *SBaseProvider) GetICloudGlobalVpcById(id string) (ICloudGlobalVpc, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudGlobalVpcById")
}

func (self *SBaseProvider) CreateICloudGlobalVpc(opts *GlobalVpcCreateOptions) (ICloudGlobalVpc, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "CreateICloudGlobalVpc")
}

func (self *SBaseProvider) GetICloudCDNDomains() ([]ICloudCDNDomain, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudCDNDomains")
}

func (self *SBaseProvider) GetICloudCDNDomainByName(name string) (ICloudCDNDomain, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetICloudCDNDomainByName")
}

func (self *SBaseProvider) CreateICloudCDNDomain(opts *CdnCreateOptions) (ICloudCDNDomain, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "CreateICloudCDNDomain")
}

func (self *SBaseProvider) GetMetrics(opts *MetricListOptions) ([]MetricValues, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetMetric")
}

func (self *SBaseProvider) GetIModelartsPools() ([]ICloudModelartsPool, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetIModelartsPools")
}

func (self *SBaseProvider) GetIModelartsPoolById(id string) (ICloudModelartsPool, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetIModelartsPoolDetail")
}

func (self *SBaseProvider) CreateIModelartsPool(pool *ModelartsPoolCreateOption, callback func(id string)) (ICloudModelartsPool, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "CreateIModelartsPool")
}

func (self *SBaseProvider) GetIModelartsPoolSku() ([]ICloudModelartsPoolSku, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetIModelartsPoolSku")
}

func (self *SBaseProvider) GetISSLCertificates() ([]ICloudSSLCertificate, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetISSLCertificates")
}

func NewBaseProvider(factory ICloudProviderFactory) SBaseProvider {
	return SBaseProvider{factory: factory}
}

func GetPublicProviders() []string {
	providers := make([]string, 0)
	for p, d := range providerTable {
		if d.IsPublicCloud() {
			providers = append(providers, p)
		}
	}
	return providers
}

func GetPrivateProviders() []string {
	providers := make([]string, 0)
	for p, d := range providerTable {
		if !d.IsPublicCloud() && !d.IsOnPremise() {
			providers = append(providers, p)
		}
	}
	return providers
}

func GetOnPremiseProviders() []string {
	providers := make([]string, 0)
	for p, d := range providerTable {
		if !d.IsPublicCloud() && d.IsOnPremise() {
			providers = append(providers, p)
		}
	}
	return providers
}

func GetProviderCloudEnv(provider string) string {
	p, err := GetProviderFactory(provider)
	if err != nil {
		return ""
	}
	if p.IsPublicCloud() {
		return CLOUD_ENV_PUBLIC_CLOUD
	}
	if p.IsOnPremise() {
		return CLOUD_ENV_ON_PREMISE
	}
	return CLOUD_ENV_PRIVATE_CLOUD
}

type baseProviderFactory struct {
}

func (factory *baseProviderFactory) ValidateChangeBandwidth(instanceId string, bandwidth int64) error {
	return nil
}

func (factory *baseProviderFactory) GetSupportedBrands() []string {
	return []string{}
}

func (factory *baseProviderFactory) IsSupportSAMLAuth() bool {
	return false
}

func (factory *baseProviderFactory) GetProvider(providerId, providerName, url, username, password string) (ICloudProvider, error) {
	return nil, errors.Wrapf(ErrNotImplemented, "GetProvider")
}

func (factory *baseProviderFactory) IsOnPremise() bool {
	return false
}

func (factory *baseProviderFactory) IsMultiTenant() bool {
	return false
}

func (factory *baseProviderFactory) IsReadOnly() bool {
	return false
}

func (factory *baseProviderFactory) IsCloudeventRegional() bool {
	return false
}

func (factory *baseProviderFactory) GetMaxCloudEventSyncDays() int {
	return 7
}

func (factory *baseProviderFactory) GetMaxCloudEventKeepDays() int {
	return 7
}

func (factory *baseProviderFactory) IsNeedForceAutoCreateProject() bool {
	return false
}

func (factory *baseProviderFactory) IsSupportCrossCloudEnvVpcPeering() bool {
	return false
}

func (factory *baseProviderFactory) IsSupportCrossRegionVpcPeering() bool {
	return false
}

func (factory *baseProviderFactory) IsSupportVpcPeeringVpcCidrOverlap() bool {
	return false
}

func (factory *baseProviderFactory) ValidateCrossRegionVpcPeeringBandWidth(bandwidth int) error {
	return nil
}

func (factory *baseProviderFactory) IsSupportModifyRouteTable() bool {
	return false
}

func (factory *baseProviderFactory) GetSupportedDnsZoneTypes() []TDnsZoneType {
	return []TDnsZoneType{}
}

func (factory *baseProviderFactory) GetSupportedDnsTypes() map[TDnsZoneType][]TDnsType {
	return map[TDnsZoneType][]TDnsType{}
}

func (factory *baseProviderFactory) GetSupportedDnsPolicyTypes() map[TDnsZoneType][]TDnsPolicyType {
	return map[TDnsZoneType][]TDnsPolicyType{}
}

func (factory *baseProviderFactory) GetSupportedDnsPolicyValues() map[TDnsPolicyType][]TDnsPolicyValue {
	return map[TDnsPolicyType][]TDnsPolicyValue{}
}

func (factory *baseProviderFactory) GetTTLRange(zoneType TDnsZoneType, productType TDnsProductType) TTlRange {
	return TTlRange{}
}

func (factory *baseProviderFactory) GetAccountIdEqualizer() func(origin, now string) bool {
	return func(origin, now string) bool {
		if len(now) > 0 && now != origin {
			return false
		}
		return true
	}
}

type SDnsCapability struct {
	ZoneTypes    []TDnsZoneType
	DnsTypes     map[TDnsZoneType][]TDnsType
	PolicyTypes  map[TDnsZoneType][]TDnsPolicyType
	PolicyValues map[TDnsPolicyType][]TDnsPolicyValue
}

func GetDnsCapabilities() map[string]SDnsCapability {
	capabilities := map[string]SDnsCapability{}
	for provider, driver := range providerTable {
		capabilities[provider] = SDnsCapability{
			ZoneTypes:    driver.GetSupportedDnsZoneTypes(),
			DnsTypes:     driver.GetSupportedDnsTypes(),
			PolicyTypes:  driver.GetSupportedDnsPolicyTypes(),
			PolicyValues: driver.GetSupportedDnsPolicyValues(),
		}
	}
	return capabilities
}

type SPremiseBaseProviderFactory struct {
	baseProviderFactory
}

func (factory *SPremiseBaseProviderFactory) IsPublicCloud() bool {
	return false
}

func (factory *SPremiseBaseProviderFactory) IsSupportPrepaidResources() bool {
	return false
}

func (factory *SPremiseBaseProviderFactory) IsOnPremise() bool {
	return true
}

func (factory *SPremiseBaseProviderFactory) IsMultiTenant() bool {
	return false
}

func (factory *SPremiseBaseProviderFactory) NeedSyncSkuFromCloud() bool {
	return false
}

type SPublicCloudBaseProviderFactory struct {
	baseProviderFactory
}

func (factory *SPublicCloudBaseProviderFactory) IsMultiTenant() bool {
	return true
}

func (factory *SPublicCloudBaseProviderFactory) IsPublicCloud() bool {
	return true
}

func (factory *SPublicCloudBaseProviderFactory) IsSupportPrepaidResources() bool {
	return true
}

func (factory *SPublicCloudBaseProviderFactory) NeedSyncSkuFromCloud() bool {
	return false
}

type SPrivateCloudBaseProviderFactory struct {
	baseProviderFactory
}

func (factory *SPrivateCloudBaseProviderFactory) IsMultiTenant() bool {
	return false
}

func (factory *SPrivateCloudBaseProviderFactory) IsPublicCloud() bool {
	return false
}

func (factory *SPrivateCloudBaseProviderFactory) IsSupportPrepaidResources() bool {
	return false
}

func (factory *SPrivateCloudBaseProviderFactory) NeedSyncSkuFromCloud() bool {
	return true
}

type ICloudModelartsPool interface {
	ICloudResource
	IBillingResource

	Delete() error
	GetProjectId() string
	GetInstanceType() string
	GetWorkType() string
	GetNodeCount() int
	ChangeConfig(opts *ModelartsPoolChangeConfigOptions) error
	GetStatusMessage() string
}

type ICloudModelartsPoolSku interface {
	ICloudResource

	GetCpuCoreCount() int
	GetCpuArch() string
	GetStatus() string
	GetMemorySizeMB() int
	GetPoolType() string
	GetGpuSize() int
	GetGpuType() string
	GetNpuSize() int
	GetNpuType() string
	GetProcessorType() string
}

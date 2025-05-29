package cloudprovider

const (
	CloudVMStatusRunning      = "running"
	CloudVMStatusStopping     = "stopping"
	CloudVMStatusSuspend      = "suspend"
	CloudVMStatusStopped      = "stopped"
	CloudVMStatusChangeFlavor = "change_flavor"
	CloudVMStatusDeploying    = "deploying"
	CloudVMStatusOther        = "other"
)

const (
	READ_ONLY_SUFFIX = ".readonly"

	CLOUD_CAPABILITY_PROJECT         = "project"
	CLOUD_CAPABILITY_COMPUTE         = "compute"
	CLOUD_CAPABILITY_NETWORK         = "network"
	CLOUD_CAPABILITY_SECURITY_GROUP  = "security_group"
	CLOUD_CAPABILITY_EIP             = "eip"
	CLOUD_CAPABILITY_SNAPSHOT_POLICY = "snapshot_policy"
	CLOUD_CAPABILITY_LOADBALANCER    = "loadbalancer"
	CLOUD_CAPABILITY_OBJECTSTORE     = "objectstore"
	CLOUD_CAPABILITY_RDS             = "rds"
	CLOUD_CAPABILITY_CACHE           = "cache" // 弹性缓存包含redis、memcached
	CLOUD_CAPABILITY_EVENT           = "event"
	CLOUD_CAPABILITY_CLOUDID         = "cloudid"
	CLOUD_CAPABILITY_DNSZONE         = "dnszone"
	CLOUD_CAPABILITY_PUBLIC_IP       = "public_ip"
	CLOUD_CAPABILITY_INTERVPCNETWORK = "intervpcnetwork"
	CLOUD_CAPABILITY_SAML_AUTH       = "saml_auth"    // 是否支持SAML 2.0
	CLOUD_CAPABILITY_QUOTA           = "quota"        // 配额
	CLOUD_CAPABILITY_NAT             = "nat"          // NAT网关
	CLOUD_CAPABILITY_NAS             = "nas"          // NAS
	CLOUD_CAPABILITY_WAF             = "waf"          // WAF
	CLOUD_CAPABILITY_MONGO_DB        = "mongodb"      // MongoDB
	CLOUD_CAPABILITY_ES              = "es"           // ElasticSearch
	CLOUD_CAPABILITY_KAFKA           = "kafka"        // Kafka
	CLOUD_CAPABILITY_APP             = "app"          // App
	CLOUD_CAPABILITY_CDN             = "cdn"          // CDN
	CLOUD_CAPABILITY_CONTAINER       = "container"    // 容器
	CLOUD_CAPABILITY_IPV6_GATEWAY    = "ipv6_gateway" // IPv6网关
	CLOUD_CAPABILITY_TABLESTORE      = "tablestore"   // 表格存储
	CLOUD_CAPABILITY_MODELARTES      = "modelarts"
	CLOUD_CAPABILITY_VPC_PEER        = "vpcpeer" // 对等连接
	CLOUD_CAPABILITY_MISC            = "misc"
	CLOUD_CAPABILITY_CERT            = "sslcertificates" // 证书
	CLOUD_CAPABILITY_IMAGE           = "image"           // 镜像
)

const (
	CLOUD_ENV_PUBLIC_CLOUD  = "public"
	CLOUD_ENV_PRIVATE_CLOUD = "private"
	CLOUD_ENV_ON_PREMISE    = "onpremise"

	CLOUD_ENV_PRIVATE_ON_PREMISE = "private_or_onpremise"
)

package compute

const (
	CLOUD_PROVIDER_ALIYUN  = "Aliyun"
	CLOUD_PROVIDER_AZURE   = "Azure"
	CLOUD_PROVIDER_AWS     = "Aws"
	CLOUD_PROVIDER_GOOGLE  = "Google"
	CLOUD_PROVIDER_PROXMOX = "Proxmox"
	CLOUD_PROVIDER_NUTANIX = "Nutanix"
	CLOUD_PROVIDER_VMWARE  = "VMware"

	CLOUD_PROVIDER_HEALTH_NORMAL        = "normal"        // The remote end is in a healthy state
	CLOUD_PROVIDER_HEALTH_INSUFFICIENT  = "insufficient"  // Insufficient on-demand resource balance
	CLOUD_PROVIDER_HEALTH_SUSPENDED     = "suspended"     // The remote end is in a frozen state
	CLOUD_PROVIDER_HEALTH_ARREARS       = "arrears"       // The remote end is in arrears
	CLOUD_PROVIDER_HEALTH_UNKNOWN       = "unknown"       // Unknown state, query failed
	CLOUD_PROVIDER_HEALTH_NO_PERMISSION = "no permission" // No permission to obtain billing information
)

const (
	CLOUD_ACCESS_ENV_AWS_GLOBAL          = CLOUD_PROVIDER_AWS + "-int"
	CLOUD_ACCESS_ENV_AWS_CHINA           = CLOUD_PROVIDER_AWS
	CLOUD_ACCESS_ENV_AZURE_GLOBAL        = CLOUD_PROVIDER_AZURE + "-int"
	CLOUD_ACCESS_ENV_AZURE_GERMAN        = CLOUD_PROVIDER_AZURE + "-de"
	CLOUD_ACCESS_ENV_AZURE_US_GOVERNMENT = CLOUD_PROVIDER_AZURE + "-us-gov"
	CLOUD_ACCESS_ENV_AZURE_CHINA         = CLOUD_PROVIDER_AZURE
	CLOUD_ACCESS_ENV_ALIYUN_GLOBAL       = CLOUD_PROVIDER_ALIYUN
	CLOUD_ACCESS_ENV_ALIYUN_FINANCE      = CLOUD_PROVIDER_ALIYUN + "-fin"
)

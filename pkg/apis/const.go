package apis

const (
	STATUS_DELETING      = "deleting"
	STATUS_DELETE_FAILED = "delete_failed"
	STATUS_CREATING      = "creating"
	STATUS_CREATE_FAILED = "create_failed"
	STATUS_AVAILABLE     = "available"
	STATUS_UNKNOWN       = "unknown"

	USER_TAG_PREFIX = "user:"

	SKU_STATUS_AVAILABLE = "available"
	SKU_STATUS_SOLDOUT   = "soldout"
)

const (
	OS_ARCH_X86 = "x86"
	OS_ARCH_ARM = "arm"

	OS_ARCH_I386    = "i386"
	OS_ARCH_X86_32  = "x86_32"
	OS_ARCH_X86_64  = "x86_64"
	OS_ARCH_AARCH32 = "aarch32"
	OS_ARCH_AARCH64 = "aarch64"
)

const (
	PUBLIC_CLOUD_ANSIBLE_USER = "cloudroot"
)

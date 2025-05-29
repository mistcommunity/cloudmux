package compute

const (
	VPC_STATUS_PENDING     = "pending"
	VPC_STATUS_AVAILABLE   = "available"
	VPC_STATUS_UNAVAILABLE = "unavailable"
	VPC_STATUS_FAILED      = "failed"
	VPC_STATUS_DELETING    = "deleting"
	VPC_STATUS_UNKNOWN     = "unknown"

	DEFAULT_VPC_ID = "default"
	NORMAL_VPC_ID  = "normal" // 没有关联VPC的安全组，统一使用normal
)

package compute

const (
	EIP_MODE_INSTANCE_PUBLICIP = "public_ip"
	EIP_MODE_STANDALONE_EIP    = "elastic_ip"

	EIP_ASSOCIATE_TYPE_SERVER       = "server"
	EIP_ASSOCIATE_TYPE_NAT_GATEWAY  = "natgateway"
	EIP_ASSOCIATE_TYPE_LOADBALANCER = "loadbalancer"
	EIP_ASSOCIATE_TYPE_UNKNOWN      = "unknown"

	EIP_ASSOCIATE_TYPE_INSTANCE_GROUP = "instancegroup"

	EIP_STATUS_READY         = "ready"
	EIP_STATUS_UNKNOWN       = "unknown"
	EIP_STATUS_ALLOCATE      = "allocate"
	EIP_STATUS_ALLOCATE_FAIL = "allocate_fail"
	EIP_STATUS_DEALLOCATE    = "deallocate"
	EIP_STATUS_ASSOCIATE     = "associate"
	EIP_STATUS_DISSOCIATE    = "dissociate"

	EIP_CHARGE_TYPE_BY_TRAFFIC   = "traffic"
	EIP_CHARGE_TYPE_BY_BANDWIDTH = "bandwidth"
)

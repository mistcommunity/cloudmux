package compute

const (
	NAT_STAUTS_AVAILABLE     = "available"     // 可用
	NAT_STATUS_ALLOCATE      = "allocate"      // 创建中
	NAT_STATUS_DEPLOYING     = "deploying"     // 配置中
	NAT_STATUS_UNKNOWN       = "unknown"       // 未知状态
	NAT_STATUS_CREATE_FAILED = "create_failed" // 创建失败
	NAT_STATUS_DELETING      = "deleting"      // 删除中

	NAT_SPEC_SMALL  = "small"  //小型
	NAT_SPEC_MIDDLE = "middle" //中型
	NAT_SPEC_LARGE  = "large"  //大型
	NAT_SPEC_XLARGE = "xlarge" //超大型

	QCLOUD_NAT_SPEC_SMALL  = "small"
	QCLOUD_NAT_SPEC_MIDDLE = "middle"
	QCLOUD_NAT_SPEC_LARGE  = "large"

	// 公网
	NAT_NETWORK_TYPE_INTERNET = "internet"
	// VPC
	NAT_NETWORK_TYPE_INTRANET = "intranet"
)

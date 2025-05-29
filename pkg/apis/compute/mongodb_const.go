package compute

const (
	MONGO_DB_STATUS_CREATING      = "creating"
	MONGO_DB_STATUS_RUNNING       = "running"
	MONGO_DB_STATUS_PROCESSING    = "processing"
	MONGO_DB_STATUS_DEPLOY        = "deploy"
	MONGO_DB_STATUS_CHANGE_CONFIG = "change_config"
	MONGO_DB_STATUS_DELETING      = "deleting"
	MONGO_DB_STATUS_REBOOTING     = "rebooting"

	MONGO_DB_ENGINE_WIRED_TIGER = "WiredTiger"
	MONGO_DB_ENGINE_ROCKS       = "Rocks"

	// 分片
	MONGO_DB_CATEGORY_SHARDING = "sharding"
	// 副本集
	MONGO_DB_CATEGORY_REPLICATE = "replicate"
)

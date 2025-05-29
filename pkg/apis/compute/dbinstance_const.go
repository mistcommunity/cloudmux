package compute

const (
	//实例状态
	DBINSTANCE_INIT        = "init"        //初始化
	DBINSTANCE_DEPLOYING   = "deploying"   //部署中
	DBINSTANCE_RUNNING     = "running"     //运行中
	DBINSTANCE_REBOOTING   = "rebooting"   //重启中
	DBINSTANCE_MIGRATING   = "migrating"   //迁移中
	DBINSTANCE_BACKING_UP  = "backing_up"  //备份中
	DBINSTANCE_RESTORING   = "restoring"   //备份恢复中
	DBINSTANCE_IMPORTING   = "importing"   //数据导入中
	DBINSTANCE_CLONING     = "cloning"     //克隆中
	DBINSTANCE_DELETING    = "deleting"    //删除中
	DBINSTANCE_MAINTENANCE = "maintenance" //维护中
	DBINSTANCE_ISOLATING   = "isolating"   //隔离中
	DBINSTANCE_ISOLATE     = "isolate"     //已隔离
	DBINSTANCE_UPGRADING   = "upgrading"   //升级中
	DBINSTANCE_UNKNOWN     = "unknown"

	DBINSTANCE_CHANGE_CONFIG = "change_config" //调整配置

	DBINSTANCE_CREATE_FAILED = "create_failed" //创建失败

	//备份状态
	DBINSTANCE_BACKUP_READY         = "ready"         //正常
	DBINSTANCE_BACKUP_CREATING      = "creating"      //创建中
	DBINSTANCE_BACKUP_CREATE_FAILED = "create_failed" //创建失败
	DBINSTANCE_BACKUP_DELETING      = "deleting"      //删除中
	DBINSTANCE_BACKUP_FAILED        = "failed"        //异常
	DBINSTANCE_BACKUP_UNKNOWN       = "unknown"       //未知

	//备份模式
	BACKUP_MODE_AUTOMATED = "automated" //自动
	BACKUP_MODE_MANUAL    = "manual"    //手动

	//实例数据库状态
	DBINSTANCE_DATABASE_CREATING = "creating" //创建中
	DBINSTANCE_DATABASE_RUNNING  = "running"  //正常
	DBINSTANCE_DATABASE_DELETING = "deleting" //删除中

	//实例用户状态
	DBINSTANCE_USER_UNAVAILABLE = "unavailable" //不可用
	DBINSTANCE_USER_AVAILABLE   = "available"   //正常
	DBINSTANCE_USER_CREATING    = "creating"    //创建中
	DBINSTANCE_USER_DELETING    = "deleting"    //删除中

	//数据库权限
	DATABASE_PRIVILEGE_RW     = "rw" //读写
	DATABASE_PRIVILEGE_R      = "r"  //只读
	DATABASE_PRIVILEGE_DDL    = "ddl"
	DATABASE_PRIVILEGE_DML    = "dml"
	DATABASE_PRIVILEGE_OWNER  = "owner"
	DATABASE_PRIVILEGE_CUSTOM = "custom" //自定义

	DBINSTANCE_TYPE_MYSQL      = "MySQL"
	DBINSTANCE_TYPE_SQLSERVER  = "SQLServer"
	DBINSTANCE_TYPE_POSTGRESQL = "PostgreSQL"
	DBINSTANCE_TYPE_MARIADB    = "MariaDB"
	DBINSTANCE_TYPE_ORACLE     = "Oracle"
	DBINSTANCE_TYPE_PPAS       = "PPAS"
	DBINSTANCE_TYPE_PERCONA    = "Percona"
	DBINSTANCE_TYPE_AURORA     = "Aurora"

	//阿里云实例类型
	ALIYUN_DBINSTANCE_CATEGORY_BASIC    = "basic"             //基础版
	ALIYUN_DBINSTANCE_CATEGORY_HA       = "high_availability" //高可用
	ALIYUN_DBINSTANCE_CATEGORY_ALWAYSON = "always_on"         //集群版
	ALIYUN_DBINSTANCE_CATEGORY_FINANCE  = "finance"           //金融版

	//谷歌云实例类型
	GOOGLE_DBINSTANCE_CATEGORY_REGIONAL = "Regional" // 高可用性（区域级）
	GOOGLE_DBINSTANCE_CATEGORY_ZONAL    = "Zonal"    // 单个地区

	// Azure
	AZURE_DBINSTANCE_CATEGORY_BASIC = "basic"

	// Aws
	// SQLServer
	AWS_DBINSTANCE_CATEGORY_ENTERPRISE_EDITION = "Enterprise Edition"
	AWS_DBINSTANCE_CATEGORY_EXPRESS_EDITION    = "Express Edition"
	AWS_DBINSTANCE_CATEGORY_STANDARD_EDITION   = "Standard Edition"
	AWS_DBINSTANCE_CATEGORY_WEB_EDITION        = "Web Edition"
	// Oracle
	AWS_DBINSTANCE_CATEGORY_STANDARD_EDITION_TWO = "Standard Edition Two"

	AWS_DBINSTANCE_CATEGORY_GENERAL_PURPOSE  = "General Purpose"
	AWS_DBINSTANCE_CATEGORY_MEMORY_OPTIMIZED = "Memory Optimized"

	//阿里云存储类型
	ALIYUN_DBINSTANCE_STORAGE_TYPE_LOCAL_SSD  = "local_ssd"  //本地盘SSD盘
	ALIYUN_DBINSTANCE_STORAGE_TYPE_CLOUD_ESSD = "cloud_essd" //ESSD云盘
	ALIYUN_DBINSTANCE_STORAGE_TYPE_CLOUD_SSD  = "cloud_ssd"  //SSD云盘

	// Azure
	AZURE_DBINSTANCE_STORAGE_TYPE_DEFAULT = "default"
)

var (
	QCLOUD_RW_PRIVILEGE_SET = []string{
		"SELECT", "INSERT", "UPDATE", "DELETE", "CREATE",
		"DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES",
		"LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE",
		"ALTER ROUTINE", "EVENT", "TRIGGER",
	}
	QCLOUD_R_PRIVILEGE_SET = []string{"SELECT", "LOCK TABLES", "SHOW VIEW"}
)

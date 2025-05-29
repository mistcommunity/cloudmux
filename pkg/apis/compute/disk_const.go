package compute

const (
	DISK_INIT                = "init"
	DISK_REBUILD             = "rebuild"
	DISK_ALLOC_FAILED        = "alloc_failed"
	DISK_BACKUP_STARTALLOC   = "backup_start_alloc"
	DISK_BACKUP_ALLOC_FAILED = "backup_alloc_failed"
	DISK_ALLOCATING          = "allocating"
	DISK_READY               = "ready"
	DISK_RESET               = "reset"
	DISK_RESET_FAILED        = "reset_failed"
	DISK_DEALLOC             = "deallocating"
	DISK_DEALLOC_FAILED      = "dealloc_failed"
	DISK_UNKNOWN             = "unknown"
	DISK_DETACHING           = "detaching"
	DISK_ATTACHING           = "attaching"
	DISK_CLONING             = "cloning" // 硬盘克隆

	DISK_SAVING = "saving"

	DISK_RESIZING      = "resizing"
	DISK_RESIZE_FAILED = "resize_failed"

	DISK_TYPE_SYS  = "sys"
	DISK_TYPE_SWAP = "swap"
	DISK_TYPE_DATA = "data"

	DISK_PREALLOCATION_OFF = "off"
	// 精简置备
	DISK_PREALLOCATION_METADATA = "metadata"
	// 厚置备延迟归零
	DISK_PREALLOCATION_FALLOC = "falloc"
	// 厚置备快速归零
	DISK_PREALLOCATION_FULL = "full"
)

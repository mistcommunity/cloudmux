package cloudprovider

type SServerSku struct {
	Name         string
	CpuCoreCount int
	MemorySizeMB int
}

type SServerSkuCreateOption struct {
	Name             string
	CpuCount         int
	VmemSizeMb       int
	SysDiskMinSizeGb int
	SysDiskMaxSizeGb int
}

package cloudprovider

type DiskCreateConfig struct {
	Name       string
	SizeGb     int
	Desc       string
	Iops       int
	Throughput int
	ProjectId  string
}

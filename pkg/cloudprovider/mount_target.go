package cloudprovider

type SMountTargetCreateOptions struct {
	AccessGroupId string
	NetworkType   string
	VpcId         string
	NetworkId     string
	FileSystemId  string
}

type SFileSystemSetQuotaInput struct {
	MaxFiles int64
	MaxGb    int64
}

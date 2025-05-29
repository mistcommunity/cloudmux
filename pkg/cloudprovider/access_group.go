package cloudprovider

type SAccessGroup struct {
	Name           string
	NetworkType    string
	FileSystemType string
	Desc           string
}

type TRWAccessType string
type TUserAccessType string

const (
	RWAccessTypeRW = TRWAccessType("RW")
	RWAccessTypeR  = TRWAccessType("R")

	UserAccessTypeNoRootSquash = TUserAccessType("no_root_squash")
	UserAccessTypeRootSquash   = TUserAccessType("root_squash")
	UserAccessTypeAllSquash    = TUserAccessType("all_squash")
)

type AccessGroupRule struct {
	Priority       int
	RWAccessType   TRWAccessType
	UserAccessType TUserAccessType
	Source         string
}

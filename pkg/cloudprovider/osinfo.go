package cloudprovider

type IOSInfo interface {
	GetFullOsName() string
	GetOsType() TOsType
	GetOsDist() string
	GetOsVersion() string
	GetOsArch() string
	GetOsLang() string
	GetBios() TBiosType
}

func IsUEFI(ios IOSInfo) bool {
	return ios.GetBios() == UEFI
}

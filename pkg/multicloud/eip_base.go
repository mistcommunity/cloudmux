
package multicloud

type SEipBase struct {
	SResourceBase
	SBillingBase
}

func (self *SEipBase) GetINetworkId() string {
	return ""
}

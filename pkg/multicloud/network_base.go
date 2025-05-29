
package multicloud

type SNetworkBase struct {
	SResourceBase
}

func (net *SNetworkBase) GetIp6Start() string {
	return ""
}

func (net *SNetworkBase) GetIp6End() string {
	return ""
}

func (net *SNetworkBase) GetIp6Mask() uint8 {
	return 0
}

func (net *SNetworkBase) GetGateway6() string {
	return ""
}

package cloudprovider

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"

	"yunion.io/x/pkg/tristate"
)

type TGlobalIdResource interface {
	GetGlobalId() string
}

type TCloudResources[T TGlobalIdResource] []T

func (a TCloudResources[T]) Len() int           { return len(a) }
func (a TCloudResources[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TCloudResources[T]) Less(i, j int) bool { return a[i].GetGlobalId() < a[j].GetGlobalId() }

func GetHostNetifs(host ICloudHost, wires []ICloudWire) []ICloudHostNetInterface {
	localWires := make(TCloudResources[ICloudWire], len(wires))
	copy(localWires, wires)
	sort.Sort(localWires)

	ret := make([]ICloudHostNetInterface, 0, len(localWires))
	for i := range localWires {
		hw := sFakeHostNics{
			host:  host,
			wire:  localWires[i],
			index: i,
		}
		ret = append(ret, &hw)
	}
	return ret
}

type sFakeHostNics struct {
	host  ICloudHost
	wire  ICloudWire
	index int
}

func (hn *sFakeHostNics) GetDevice() string {
	return fmt.Sprintf("eth%d", hn.index)
}

func (hn *sFakeHostNics) GetDriver() string {
	return "e1000"
}

func (hn *sFakeHostNics) GetMac() string {
	return HashIdsMac(hn.host.GetGlobalId(), hn.wire.GetGlobalId())
}

func (hn *sFakeHostNics) GetVlanId() int {
	return 1
}

func (hn *sFakeHostNics) GetIndex() int8 {
	return int8(hn.index)
}

func (hn *sFakeHostNics) IsLinkUp() tristate.TriState {
	return tristate.True
}

func (hn *sFakeHostNics) GetIpAddr() string {
	return ""
}

func (hn *sFakeHostNics) GetMtu() int32 {
	return 1500
}

func (hn *sFakeHostNics) GetNicType() string {
	return ""
}

func (hn *sFakeHostNics) GetBridge() string {
	return fmt.Sprintf("br%d", hn.index)
}

func (hn *sFakeHostNics) GetIWire() ICloudWire {
	return hn.wire
}

func HashIdsMac(ids ...string) string {
	h := sha256.New()
	for _, id := range ids {
		h.Write([]byte(id))
	}
	sum := h.Sum(nil)
	hexStr := make([]string, 6)
	hexStr[0] = "ff"
	for i := 1; i < 6; i++ {
		hexStr[i] = fmt.Sprintf("%02x", sum[i])
	}
	return strings.Join(hexStr, ":")
}

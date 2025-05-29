
package azure

/*
{
	"capabilities":[
		{"name":"MaxResourceVolumeMB","value":"286720"},
		{"name":"OSVhdSizeMB","value":"1047552"},
		{"name":"vCPUs","value":"20"},
		{"name":"MemoryGB","value":"140"},
		{"name":"MaxDataDiskCount","value":"64"},
		{"name":"LowPriorityCapable","value":"True"},
		{"name":"PremiumIO","value":"True"},
		{"name":"EphemeralOSDiskSupported","value":"True"}
	],
	"family":"standardDSv2Family",
	"locations":["CentralUSEUAP"],
	"name":"Standard_DS15_v2",
	"resourceType":"virtualMachines",
	"restrictions":[],
	"size":"DS15_v2",
	"tier":"Standard"
}
*/

type SResourceSkuCapability struct {
	Name  string
	Value string
}

type TResourceSkuCapacityScaleType string

const (
	ResourceSkuCapacityScaleTypeAutomatic = TResourceSkuCapacityScaleType("Automatic")
	ResourceSkuCapacityScaleTypeManual    = TResourceSkuCapacityScaleType("Manual")
	ResourceSkuCapacityScaleTypeNone      = TResourceSkuCapacityScaleType("None")
)

type SResourceSkuCapacity struct {
	Default   int
	Maximum   int
	Minimum   int
	ScaleType TResourceSkuCapacityScaleType
}

type SResourceSkuLocationInfo struct {
	Location string
	Zones    []string
}

type TResourceSkuRestrictionsType string

const (
	ResourceSkuRestrictionsTypeLocation = TResourceSkuRestrictionsType("Location")
	ResourceSkuRestrictionsTypeZone     = TResourceSkuRestrictionsType("Zone")
)

type TResourceSkuRestrictionsReasonCode string

const (
	ResourceSkuRestrictionsReasonCodeNotAvailable = TResourceSkuRestrictionsReasonCode("NotAvailableForSubscription")
	ResourceSkuRestrictionsReasonCodeQuotaId      = TResourceSkuRestrictionsReasonCode("QuotaId")
)

type SResourceSkuRestrictionInfo struct {
	Locations []string
	Zones     []string
}

type SResourceSkuRestrictions struct {
	ReasonCode      TResourceSkuRestrictionsReasonCode
	RestrictionInfo SResourceSkuRestrictionInfo
	Type            TResourceSkuRestrictionsType
	Values          []string
}

type SResourceSku struct {
	Capabilities []SResourceSkuCapability
	Capacity     *SResourceSkuCapacity
	Family       string
	Kind         string
	LocationInfo []SResourceSkuLocationInfo
	Locations    []string
	Name         string
	ResourceType string
	Restrictions []SResourceSkuRestrictions
	Size         string
	Tier         string
}

type SResourceSkusResult struct {
	NextLink string
	Value    []SResourceSku
}

func (self *SAzureClient) ListResourceSkus() ([]SResourceSku, error) {
	skus := []SResourceSku{}
	resource := "Microsoft.Compute/skus"
	return skus, self.list(resource, nil, &skus)
}

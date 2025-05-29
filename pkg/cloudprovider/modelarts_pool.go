package cloudprovider

type ModelartsPoolCreateOption struct {
	Name         string
	PoolDesc     string
	BillingMode  uint
	PeriodType   uint
	PeriodNum    uint
	AutoRenew    uint
	InstanceType string
	NodeCount    int
	NetworkId    string
	Cidr         string

	WorkType string
}

type Azs struct {
	Az    string `json:"az"`
	Count int    `json:"count"`
}

type ModelartsPoolChangeConfigOptions struct {
	NodeCount int
}

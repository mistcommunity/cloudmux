package cloudprovider

// +onecloud:model-api-gen
type ElasticSearchAccessInfo struct {
	Domain           string
	PrivateDomain    string
	Vip              string
	Port             int
	PrivatePort      int
	KibanaUrl        string
	KibanaPrivateUrl string
}

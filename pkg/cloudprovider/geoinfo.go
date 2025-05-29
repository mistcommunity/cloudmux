package cloudprovider

// +onecloud:model-api-gen
type SGeographicInfo struct {
	// 纬度
	// example: 26.647003
	Latitude float32 `list:"user" update:"admin" create:"admin_optional"`
	// 经度
	// example: 106.630211
	Longitude float32 `list:"user" update:"admin" create:"admin_optional"`

	// 城市
	// example: Guiyang
	City string `list:"user" width:"32" update:"admin" create:"admin_optional"`
	// 国家代码
	// example: CN
	CountryCode string `list:"user" width:"4" update:"admin" create:"admin_optional"`
}

func (self SGeographicInfo) IsEquals(geo SGeographicInfo) bool {
	return self.City == geo.City && self.CountryCode == geo.CountryCode &&
		self.Latitude-geo.Latitude < 0.01 && self.Longitude-geo.Longitude < 0.01
}

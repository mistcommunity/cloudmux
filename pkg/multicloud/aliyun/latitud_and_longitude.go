
package aliyun

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

var LatitudeAndLongitude = map[string]cloudprovider.SGeographicInfo{
	"cn-qingdao":            api.RegionQingdao,
	"cn-beijing":            api.RegionBeijing,
	"cn-zhangjiakou":        api.RegionZhangjiakou,
	"cn-huhehaote":          api.RegionHuhehaote,
	"cn-huhehaote-nebula-1": api.RegionHuhehaote,
	"cn-hangzhou":           api.RegionHangzhou,
	"cn-shanghai":           api.RegionShanghai,
	"cn-shanghai-finance-1": api.RegionShanghai,
	"cn-shenzhen":           api.RegionShenzhen,
	"cn-shenzhen-finance-1": api.RegionShenzhen,
	"cn-hongkong":           api.RegionHongkong,
	"cn-chengdu":            api.RegionChengdu,
	"cn-heyuan":             api.RegionHeyuan,
	"ap-northeast-1":        api.RegionTokyo,
	"ap-southeast-1":        api.RegionSingapore,
	"ap-southeast-2":        api.RegionSydney,
	"ap-southeast-3":        api.RegionKualaLumpur,
	"ap-southeast-5":        api.RegionJakarta,
	"ap-south-1":            api.RegionMumbai,
	"us-east-1":             api.RegionVirginia,
	"us-west-1":             api.RegionSiliconValley,
	"eu-west-1":             api.RegionLondon,
	"me-east-1":             api.RegionDubai,
	"eu-central-1":          api.RegionFrankfurt,
	"cn-wulanchabu":         api.RegionWulanchabu,
	"cn-guangzhou":          api.RegionGuangzhou,
}

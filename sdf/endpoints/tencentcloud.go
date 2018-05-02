package endpoints

import "regexp"

const TencentCloudPartitionID = "tencent-cloud" // tencent cloud partition.

// tencent cloud partition's regions.
const (
	ApBeijingRegionID       = "ap-beijing"        // 华北地区(北京).
	ApChengduRegionID       = "ap-chengdu"        // 西南地区(成都).
	ApGuangzhouRegionID     = "ap-guangzhou"      // 华南地区(广州).
	ApGuangzhouOpenRegionID = "ap-guangzhou-open" // 华南地区(广州Open).
	ApHongkongRegionID      = "ap-hongkong"       // 东南亚地区(香港).
	ApShanghaiRegionID      = "ap-shanghai"       // 华东地区(上海).
	ApShanghaiFsiRegionID   = "ap-shanghai-fsi"   // 华东地区(上海金融).
	ApShenzhenFsiRegionID   = "ap-shenzhen-fsi"   // 华南地区(深圳金融).
	ApSingaporeRegionID     = "ap-singapore"      // 东南亚地区(新加坡).
	EuFrankfurtRegionID     = "eu-frankfurt"      // 欧洲地区(德国).
	NaSiliconvalleyRegionID = "na-siliconvalley"  // 美国西部(硅谷).
	NaTorontoRegionID       = "na-toronto"        // 北美地区(多伦多).
)

// TencentcloudPartition returns the Resolver for tencentcloud standard.
func TencentcloudPartition() Partition {
	return tencentcloudPartition.Partition()
}

var tencentcloudPartition = partition{
	ID:        "tencentcloud",
	Name:      "tencent cloud",
	DNSSuffix: "api.qcloud.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(sh|gz|bj|hk|ca|sg|shjr|szjr|gzppen)$")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{service}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"tencent"},
	},
	Regions: regions{
		"ap-beijing": region{
			Description: "华北地区(北京)",
		},
		"ap-chengdu": region{
			Description: "西南地区(成都)",
		},
		"ap-guangzhou": region{
			Description: "华南地区(广州)",
		},
		"ap-guangzhou-open": region{
			Description: "华南地区(广州Open)",
		},
		"ap-hongkong": region{
			Description: "东南亚地区(香港)",
		},
		"ap-shanghai": region{
			Description: "华东地区(上海)",
		},
		"ap-shanghai-fsi": region{
			Description: "华东地区(上海金融)",
		},
		"ap-shenzhen-fsi": region{
			Description: "华南地区(深圳金融)",
		},
		"ap-singapore": region{
			Description: "东南亚地区(新加坡)",
		},
		"eu-frankfurt": region{
			Description: "欧洲地区(德国)",
		},
		"na-siliconvalley": region{
			Description: "美国西部(硅谷)",
		},
		"na-toronto": region{
			Description: "北美地区(多伦多)",
		},
	},
	Services: services{
		"cbs": service{

			Endpoints: endpoints{
				"ap-beijing":        endpoint{},
				"ap-chengdu":        endpoint{},
				"ap-guangzhou":      endpoint{},
				"ap-guangzhou-open": endpoint{},
				"ap-hongkong":       endpoint{},
				"ap-shanghai":       endpoint{},
				"ap-shanghai-fsi":   endpoint{},
				"ap-shenzhen-fsi":   endpoint{},
				"ap-singapore":      endpoint{},
				"eu-frankfurt":      endpoint{},
				"na-siliconvalley":  endpoint{},
				"na-toronto":        endpoint{},
			},
		},
		"cvm": service{

			Endpoints: endpoints{
				"ap-beijing":        endpoint{},
				"ap-chengdu":        endpoint{},
				"ap-guangzhou":      endpoint{},
				"ap-guangzhou-open": endpoint{},
				"ap-hongkong":       endpoint{},
				"ap-shanghai":       endpoint{},
				"ap-shanghai-fsi":   endpoint{},
				"ap-shenzhen-fsi":   endpoint{},
				"ap-singapore":      endpoint{},
				"eu-frankfurt":      endpoint{},
				"na-siliconvalley":  endpoint{},
				"na-toronto":        endpoint{},
			},
		},
		"dfw": service{
			Endpoints: endpoints{
				"ap-beijing":        endpoint{},
				"ap-chengdu":        endpoint{},
				"ap-guangzhou":      endpoint{},
				"ap-guangzhou-open": endpoint{},
				"ap-hongkong":       endpoint{},
				"ap-shanghai":       endpoint{},
				"ap-shanghai-fsi":   endpoint{},
				"ap-shenzhen-fsi":   endpoint{},
				"ap-singapore":      endpoint{},
				"eu-frankfurt":      endpoint{},
				"na-siliconvalley":  endpoint{},
				"na-toronto":        endpoint{},
			},
		},
		"vpc": service{
			Endpoints: endpoints{
				"ap-beijing":        endpoint{},
				"ap-chengdu":        endpoint{},
				"ap-guangzhou":      endpoint{},
				"ap-guangzhou-open": endpoint{},
				"ap-hongkong":       endpoint{},
				"ap-shanghai":       endpoint{},
				"ap-shanghai-fsi":   endpoint{},
				"ap-shenzhen-fsi":   endpoint{},
				"ap-singapore":      endpoint{},
				"eu-frankfurt":      endpoint{},
				"na-siliconvalley":  endpoint{},
				"na-toronto":        endpoint{},
			},
		},
		"ec2metadata": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "169.254.169.254/latest",
					Protocols: []string{"http"},
				},
			},
		},
		"image": service{

			Endpoints: endpoints{
				"ap-beijing":        endpoint{},
				"ap-chengdu":        endpoint{},
				"ap-guangzhou":      endpoint{},
				"ap-guangzhou-open": endpoint{},
				"ap-hongkong":       endpoint{},
				"ap-shanghai":       endpoint{},
				"ap-shanghai-fsi":   endpoint{},
				"ap-shenzhen-fsi":   endpoint{},
				"ap-singapore":      endpoint{},
				"eu-frankfurt":      endpoint{},
				"na-siliconvalley":  endpoint{},
				"na-toronto":        endpoint{},
			},
		},
		"monitor": service{

			Endpoints: endpoints{
				"ap-beijing":        endpoint{},
				"ap-chengdu":        endpoint{},
				"ap-guangzhou":      endpoint{},
				"ap-guangzhou-open": endpoint{},
				"ap-hongkong":       endpoint{},
				"ap-shanghai":       endpoint{},
				"ap-shanghai-fsi":   endpoint{},
				"ap-shenzhen-fsi":   endpoint{},
				"ap-singapore":      endpoint{},
				"eu-frankfurt":      endpoint{},
				"na-siliconvalley":  endpoint{},
				"na-toronto":        endpoint{},
			},
		},
	},
}

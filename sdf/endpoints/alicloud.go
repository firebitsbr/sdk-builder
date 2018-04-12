package endpoints

import "regexp"

const AlicloudPartitionID = "aliyun" // alicloud standard partition.

// alicloud standard partition's regions.
const (
	CnBeijingRegionID          = "cn-beijing"            // 华北2.
	CnHangzhouRegionID         = "cn-hangzhou"           // 华东1.
	CnHongkongRegionID         = "cn-hongkong"           // 香港.
	CnQingdaoRegionID          = "cn-qingdao"            // 华北1.
	CnShanghaiRegionID         = "cn-shanghai"           // 华东2.
	CnShanghaiFinance1RegionID = "cn-shanghai-finance-1" // 华东2金融云.
	CnShenzhenRegionID         = "cn-shenzhen"           // 华南1.
	CnShenzhenFinance1RegionID = "cn-shenzhen-finance-1" // 华南1金融云.
	CnZhangjiakouRegionID      = "cn-zhangjiakou"        // 华北3.
	MeEast1RegionID            = "me-east-1"             // 中东东部1.
)

// AlicloudPartition returns the Resolver for alicloud standard.
func AlicloudPartition() Partition {
	return alicloudPartition.Partition()
}

var alicloudPartition = partition{
	ID:        "alicloud",
	Name:      "alicloud standard",
	DNSSuffix: "aliyuncs.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(cn|eu|ap|me)\\-\\w+")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{service}-{region}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"aliyun"},
	},
	Regions: regions{
		"cn-beijing": region{
			Description: "华北2",
		},
		"cn-hangzhou": region{
			Description: "华东1",
		},
		"cn-hongkong": region{
			Description: "香港",
		},
		"cn-qingdao": region{
			Description: "华北1",
		},
		"cn-shanghai": region{
			Description: "华东2",
		},
		"cn-shanghai-finance-1": region{
			Description: "华东2金融云",
		},
		"cn-shenzhen": region{
			Description: "华南1",
		},
		"cn-shenzhen-finance-1": region{
			Description: "华南1金融云",
		},
		"cn-zhangjiakou": region{
			Description: "华北3",
		},
		"me-east-1": region{
			Description: "中东东部1",
		},
	},
	Services: services{
		"ecs": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{
					Hostname: "ecs.ap-northeast-1.aliyuncs.com",
				},
				"ap-southeast-1": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "ap-southeast-1",
					},
				},
				"ap-southeast-2": endpoint{
					Hostname: "ecs.ap-southeast-2.aliyuncs.com",
				},
				"cn-beijing": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "cn-beijing",
					},
				},
				"cn-hangzhou": endpoint{},
				"cn-hongkong": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "cn-hongkong",
					},
				},
				"cn-qingdao": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "cn-qingdao",
					},
				},
				"cn-shanghai": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "cn-shanghai",
					},
				},
				"cn-shanghai-finance-1": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "cn-shanghai-finance-1",
					},
				},
				"cn-shenzhen": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "cn-shenzhen",
					},
				},
				"cn-shenzhen-finance-1": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
				},
				"cn-zhangjiakou": endpoint{
					Hostname: "ecs.cn-zhangjiakou.aliyuncs.com",
				},
				"eu-central-1": endpoint{
					Hostname: "ecs.eu-central-1.aliyuncs.com",
				},
				"me-east-1": endpoint{
					Hostname: "ecs.me-east-1.aliyuncs.com",
				},
				"us-east-1": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"us-west-1": endpoint{
					Hostname: "ecs-cn-hangzhou.aliyuncs.com",
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
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
	},
}

gen-services:
	go generate ./service

gen-services-aws:
	go run -tags codegen ./private/model/cli/gen-api/main.go -path=./service/aws ./models/apis/aws/*/*/api-2.json

gen-services-alicloud:
	go run -tags codegen ./private/model/cli/gen-api/main.go -path=./service/alicloud ./models/apis/alicloud/*/*/api-2.json

gen-services-tencentcloud:
	go run -tags codegen ./private/model/cli/gen-api/main.go -path=./service/tencentcloud ./models/apis/tencentcloud/*/*/api-2.json

gen-endpoints:
	go generate ./models/endpoints/
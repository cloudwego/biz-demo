module github.com/cloudwego/biz-demo/mall

go 1.16

require (
	github.com/apache/thrift v0.16.0
	github.com/cloudwego/kitex v0.2.1
	github.com/cloudwego/netpoll v0.2.4 // indirect
	github.com/cloudwego/thriftgo v0.1.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kitex-contrib/registry-etcd v0.0.0-20220110034026-b1c94979cea3
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/tidwall/gjson v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

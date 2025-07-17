module github.com/cloudwego/biz-demo/gomall/app/frontend

go 1.23.0

toolchain go1.23.11

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/cloudwego/hertz v0.8.1
	github.com/hertz-contrib/cors v0.1.0
	github.com/hertz-contrib/gzip v0.0.3
	github.com/hertz-contrib/logger/accesslog v0.0.0-20240128134225-6b18af47a115
	github.com/hertz-contrib/logger/logrus v1.0.1
	github.com/hertz-contrib/pprof v0.1.2
	github.com/kr/pretty v0.3.1
	github.com/redis/go-redis/v9 v9.5.1
	go.uber.org/zap v1.27.0
	google.golang.org/protobuf v1.34.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.6
	gorm.io/gorm v1.25.10
)

require (
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20220413063733-65bf48ffb3a7 // indirect
	github.com/bytedance/sonic v1.8.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cloudwego/netpoll v0.5.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/pprof v0.0.0-20211214055906-6f57359322fd // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/sys v0.34.0 // indirect
)

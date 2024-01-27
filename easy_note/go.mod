module github.com/cloudwego/biz-demo/easy_note

go 1.18

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/hertz v0.7.3
	github.com/cloudwego/kitex v0.8.0
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/hertz-contrib/gzip v0.0.3
	github.com/hertz-contrib/jwt v1.0.2
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus v0.1.1
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.2.0
	github.com/hertz-contrib/pprof v0.1.1
	github.com/hertz-contrib/requestid v1.1.0
	github.com/kitex-contrib/obs-opentelemetry v0.2.1
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230219092456-5e6c84962323
	github.com/kitex-contrib/registry-etcd v0.1.0
	go.opentelemetry.io/otel/trace v1.13.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.6
	gorm.io/plugin/opentelemetry v0.1.1
)

require (
	github.com/andeya/goutil v1.0.0 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.6 // indirect
	github.com/bytedance/gopkg v0.0.0-20230728082804-614d0af6619b // indirect
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/choleraehyq/pid v0.0.17 // indirect
	github.com/cloudwego/configmanager v0.2.0 // indirect
	github.com/cloudwego/dynamicgo v0.1.6 // indirect
	github.com/cloudwego/fastpb v0.0.4 // indirect
	github.com/cloudwego/frugal v0.1.12 // indirect
	github.com/cloudwego/localsession v0.0.2 // indirect
	github.com/cloudwego/netpoll v0.5.1 // indirect
	github.com/cloudwego/thriftgo v0.3.3 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/pprof v0.0.0-20220608213341-c488b8fa1db3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // indirect
	github.com/henrylee2cn/ameda v1.5.1 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nyaruka/phonenumbers v1.1.6 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/v3 v3.5.5 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.39.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.14.0 // indirect
	go.opentelemetry.io/contrib/propagators/ot v1.14.0 // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.36.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.36.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.13.0 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.13.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.36.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230216225411-c8e22ba71e44 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

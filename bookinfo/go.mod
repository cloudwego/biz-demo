module github.com/cloudwego/biz-demo/bookinfo

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/hertz v0.3.0
	github.com/cloudwego/kitex v0.4.2
	github.com/google/wire v0.5.0
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.1.1
	github.com/kitex-contrib/obs-opentelemetry v0.1.0
	github.com/kitex-contrib/xds v0.0.0-20220825123131-e4b331610f97
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.13.0
	go.opentelemetry.io/otel v1.9.0
)

package main

import (
	"net"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/mtl"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/kitex-contrib/config-etcd/etcd"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	_ = godotenv.Load()
	mtl.InitMtl()
	rpc.InitClient()
	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	serviceName := conf.GetConf().Kitex.Service
	etcdClient, err := etcd.NewClient(etcd.Options{})
	if err != nil {
		panic(err)
	}

	// service info
	opts = append(opts,
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
		server.WithSuite(etcdServer.NewSuite(serviceName, etcdClient)),
	)

	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulRegister(os.Getenv("REGISTRY_ADDR"))
		if err != nil {
			klog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(r))
	}
	_ = provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)
	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))

	// klog
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(os.Stdout)
	return
}

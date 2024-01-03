package rpc

import (
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"os"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	"github.com/cloudwego/biz-demo/gomall/app/cart/infra/mtl"
	"github.com/cloudwego/biz-demo/gomall/app/cart/kitex_gen/product/productcatalogservice"
	cartutils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		cartutils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}
	_ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider), provider.WithEnableMetrics(false))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithSuite(tracing.NewClientSuite()),
	)

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}

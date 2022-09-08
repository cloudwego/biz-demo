package injectors

import (
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds/xdssuite"
)

type DetailsClientOptions struct {
	Endpoint  string `mapstructure:"endpoint"`
	EnableXDS bool   `mapstructure:"enableXDS"`
}

func DefaultDetailsClientOptions() *DetailsClientOptions {
	return &DetailsClientOptions{
		Endpoint:  ":8084",
		EnableXDS: false,
	}
}

func ProvideDetailsClient(opts *DetailsClientOptions) (detailsservice.Client, error) {
	if opts.EnableXDS {
		return detailsservice.NewClient(
			constants.DetailsServiceName,
			kclient.WithHostPorts(opts.Endpoint),
			kclient.WithSuite(tracing.NewClientSuite()),
			kclient.WithXDSSuite(xds.ClientSuite{
				RouterMiddleware: xdssuite.NewXDSRouterMiddleware(),
				Resolver:         xdssuite.NewXDSResolver(),
			}),
		)
	}

	return detailsservice.NewClient(
		constants.DetailsServiceName,
		kclient.WithHostPorts(opts.Endpoint),
		kclient.WithSuite(tracing.NewClientSuite()),
	)
}

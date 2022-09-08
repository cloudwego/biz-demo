package injectors

import (
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings/ratingservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds/xdssuite"
)

type RatingsClientOptions struct {
	Endpoint  string `mapstructure:"endpoint"`
	EnableXDS bool   `mapstructure:"enableXDS"`
}

func DefaultRatingsClientOptions() *RatingsClientOptions {
	return &RatingsClientOptions{
		Endpoint:  ":8083",
		EnableXDS: false,
	}
}

func ProvideRatingsClient(opts *RatingsClientOptions) (ratingservice.Client, error) {
	if opts.EnableXDS {
		return ratingservice.NewClient(
			constants.RatingsServiceName,
			kclient.WithHostPorts(opts.Endpoint),
			kclient.WithSuite(tracing.NewClientSuite()),
			kclient.WithXDSSuite(xds.ClientSuite{
				RouterMiddleware: xdssuite.NewXDSRouterMiddleware(),
				Resolver:         xdssuite.NewXDSResolver(),
			}),
		)
	}

	return ratingservice.NewClient(
		constants.RatingsServiceName,
		kclient.WithHostPorts(opts.Endpoint),
		kclient.WithSuite(tracing.NewClientSuite()),
	)
}

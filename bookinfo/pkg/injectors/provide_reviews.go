package injectors

import (
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews/reviewsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds/xdssuite"
)

type ReviewClientOptions struct {
	Endpoint  string `mapstructure:"endpoint"`
	EnableXDS bool   `mapstructure:"enableXDS"`
}

func DefaultReviewClientOptions() *ReviewClientOptions {
	return &ReviewClientOptions{
		Endpoint:  ":8082",
		EnableXDS: false,
	}
}

func ProvideReviewClient(opts *ReviewClientOptions) (reviewsservice.Client, error) {
	if opts.EnableXDS {
		return reviewsservice.NewClient(
			constants.ReviewsServiceName,
			kclient.WithHostPorts(opts.Endpoint),
			kclient.WithSuite(tracing.NewClientSuite()),
			kclient.WithXDSSuite(xds.ClientSuite{
				RouterMiddleware: xdssuite.NewXDSRouterMiddleware(),
				Resolver:         xdssuite.NewXDSResolver(),
			}),
		)
	}

	return reviewsservice.NewClient(
		constants.ReviewsServiceName,
		kclient.WithHostPorts(opts.Endpoint),
		kclient.WithSuite(tracing.NewClientSuite()),
	)
}

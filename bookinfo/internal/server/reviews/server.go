package reviews

import (
	"context"
	"net"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings/ratingservice"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews/reviewsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds"
)

type Server struct {
	opts          *ServerOptions
	svc           reviews.ReviewsService
	ratingsClient ratingservice.Client
}

type ServerOptions struct {
	Addr      string `mapstructure:"addr"`
	EnableXDS bool   `mapstructure:"enableXDS"`
}

func DefaultServerOptions() *ServerOptions {
	return &ServerOptions{
		Addr:      ":8082",
		EnableXDS: false,
	}
}

func (s *Server) Run(ctx context.Context) error {
	if s.opts.EnableXDS {
		if err := xds.Init(); err != nil {
			klog.Fatal(err)
		}
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ReviewsServiceName),
		provider.WithInsecure(),
	)
	defer p.Shutdown(ctx)

	addr, err := net.ResolveTCPAddr("tcp", s.opts.Addr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := reviewsservice.NewServer(
		s.svc,
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)
	if err := svr.Run(); err != nil {
		klog.Fatal(err)
	}

	return nil
}
